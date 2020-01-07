package integrationtest

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/v8/aeternity"
	"github.com/aeternity/aepp-sdk-go/v8/config"
	"github.com/aeternity/aepp-sdk-go/v8/naet"
	"github.com/aeternity/aepp-sdk-go/v8/transactions"
)

func getNameEntry(t *testing.T, node *naet.Node, name string) (responseJSON string) {
	response, err := node.GetNameEntryByName(name)
	if err != nil {
		t.Fatal(err)
	}
	r, _ := response.MarshalBinary()
	responseJSON = string(r)
	return responseJSON
}

func randomName(length int) string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		r := rand.Intn(len(letters))
		b[i] = letters[r]
	}

	ans := fmt.Sprintf("%s.chain", string(b))
	return ans
}

func TestAENSWorkflow(t *testing.T) {
	node := setupNetwork(t, privatenetURL, false)
	alice, bob := setupAccounts(t)
	ctxAlice := aeternity.NewContext(alice, node)
	ctxBob := aeternity.NewContext(bob, node)

	name := randomName(int(config.Client.Names.NameAuctionMaxLength + 1))
	// Preclaim the name
	preclaimTx, nameSalt, err := transactions.NewNamePreclaimTx(alice.Address, name, ctxAlice.TTLNoncer())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Preclaim", name)
	r, err := ctxAlice.SignBroadcastWait(preclaimTx, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)

	// Claim the name
	nameFee := transactions.CalculateMinNameFee(name)
	claimTx, err := transactions.NewNameClaimTx(alice.Address, name, nameSalt, nameFee, ctxAlice.TTLNoncer())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Claim")
	r, err = ctxAlice.SignBroadcastWait(claimTx, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)

	// Verify that the name exists
	var nameEntry string
	printNameEntry := func() {
		nameEntry = getNameEntry(t, node, name)
		fmt.Println(nameEntry)
	}
	delay(printNameEntry)

	// Update the name, make it point to something
	alicesAddress, err := transactions.NewNamePointer("account_pubkey", alice.Address)
	if err != nil {
		t.Fatal(err)
	}
	updateTx, err := transactions.NewNameUpdateTx(alice.Address, name, []*transactions.NamePointer{alicesAddress}, config.Client.Names.ClientTTL, ctxAlice.TTLNoncer())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Update")
	r, err = ctxAlice.SignBroadcastWait(updateTx, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)

	// Verify that the name was updated
	delay(printNameEntry)
	if !strings.Contains(nameEntry, alice.Address) {
		t.Fatalf("The AENS entry should now point to %s but doesn't: %s", alice.Address, nameEntry)
	}

	// Transfer the name to a recipient
	transferTx, err := transactions.NewNameTransferTx(alice.Address, name, bob.Address, ctxAlice.TTLNoncer())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Transfer")
	r, err = ctxAlice.SignBroadcastWait(transferTx, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)

	// Receiver updates the name, makes it point to himself
	bobsAddress, err := transactions.NewNamePointer("account_pubkey", alice.Address)
	if err != nil {
		t.Fatal(err)
	}
	updateTx2, err := transactions.NewNameUpdateTx(bob.Address, name, []*transactions.NamePointer{bobsAddress}, config.Client.Names.ClientTTL, ctxBob.TTLNoncer())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Update Signed By Recipient")
	r, err = ctxBob.SignBroadcastWait(updateTx2, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)

	// Revoke the name - shouldn't work because it is signed by the sender, who no longer owns the address
	revokeTx, err := transactions.NewNameRevokeTx(alice.Address, name, ctxAlice.TTLNoncer())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Revoke")
	_, revokeTxShouldHaveFailed := ctxAlice.SignBroadcastWait(revokeTx, config.Client.WaitBlocks)
	if revokeTxShouldHaveFailed == nil {
		t.Fatal(err)
	}

	// Revoke the name - signed by the recipient
	revokeTx2, err := transactions.NewNameRevokeTx(bob.Address, name, ctxBob.TTLNoncer())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Revoke Signed By Recipient")
	r, err = ctxBob.SignBroadcastWait(revokeTx2, config.Client.WaitBlocks)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
