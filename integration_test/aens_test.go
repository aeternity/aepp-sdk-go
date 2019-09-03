package integrationtest

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

func getNameEntry(t *testing.T, node *aeternity.Node, name string) (responseJSON string) {
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

	ans := fmt.Sprintf("%s.test", string(b))
	return ans
}

func TestAENSWorkflow(t *testing.T) {
	node := setupNetwork(t, privatenetURL, false)
	helpers := aeternity.Helpers{Node: node}
	alice, bob := setupAccounts(t)
	aensAlice := aeternity.Context{Helpers: helpers, Address: alice.Address}

	name := randomName(6)
	// Preclaim the name
	preclaimTx, salt, err := aensAlice.NamePreclaimTx(name, aeternity.Config.Client.Fee)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Preclaim %+v with name %s \n", preclaimTx, name)
	_, hash, _, err := aeternity.SignBroadcastTransaction(preclaimTx, alice, node, networkID)
	if err != nil {
		t.Fatal(err)
	}

	// Wait for a bit
	_, _, _ = waitForTransaction(node, hash)

	// Claim the name
	claimTx, err := aensAlice.NameClaimTx(name, *salt, aeternity.Config.Client.Fee)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Claim %+v\n", claimTx)
	_, hash, _, err = aeternity.SignBroadcastTransaction(claimTx, alice, node, networkID)
	if err != nil {
		t.Fatal(err)
	}

	// Wait for a bit
	_, _, _ = waitForTransaction(node, hash)

	// Verify that the name exists
	var nameEntry string
	printNameEntry := func() {
		nameEntry = getNameEntry(t, node, name)
		fmt.Println(nameEntry)
	}
	delay(printNameEntry)

	// Update the name, make it point to something
	updateTx, err := aensAlice.NameUpdateTx(name, alice.Address)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Update %+v\n", updateTx)
	_, _, _, err = aeternity.SignBroadcastTransaction(updateTx, alice, node, networkID)
	if err != nil {
		t.Fatal(err)
	}

	// Verify that the name was updated
	delay(printNameEntry)
	if !strings.Contains(nameEntry, alice.Address) {
		t.Fatalf("The AENS entry should now point to %s but doesn't: %s", alice.Address, nameEntry)
	}

	// Transfer the name to a recipient
	transferTx, err := aensAlice.NameTransferTx(name, bob.Address)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Transfer %+v\n", transferTx)
	_, hash, _, err = aeternity.SignBroadcastTransaction(transferTx, alice, node, networkID)
	if err != nil {
		t.Fatal(err)
	}

	// Wait for a bit
	_, _, _ = waitForTransaction(node, hash)

	// Receiver updates the name, makes it point to himself
	aensBob := aeternity.Context{Helpers: helpers, Address: bob.Address}

	updateTx2, err := aensBob.NameUpdateTx(name, bob.Address)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Update Signed By Recipient %+v\n", updateTx2)
	_, _, _, err = aeternity.SignBroadcastTransaction(updateTx2, bob, node, networkID)
	if err != nil {
		t.Fatal(err)
	}

	// Revoke the name - shouldn't work because it is signed by the sender, who no longer owns the address
	revokeTx, err := aensAlice.NameRevokeTx(name)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Revoke %+v\n", revokeTx)
	_, hash, _, err = aeternity.SignBroadcastTransaction(revokeTx, alice, node, networkID)
	if err != nil {
		t.Fatal(err)
	}

	// Wait for a bit
	_, _, revokeTxShouldHaveFailed := waitForTransaction(node, hash)
	if revokeTxShouldHaveFailed == nil {
		t.Fatal("After transferring the name to Recipient, the Sender should not have been able to revoke the name")
	} else {
		fmt.Println(revokeTxShouldHaveFailed)
	}

	// Revoke the name - signed by the recipient
	revokeTx2, err := aensBob.NameRevokeTx(name)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Revoke Signed By Recipient %+v\n", revokeTx2)
	_, hash, _, err = aeternity.SignBroadcastTransaction(revokeTx2, bob, node, networkID)
	if err != nil {
		t.Fatal(err)
	}
	// Wait for a bit
	_, _, _ = waitForTransaction(node, hash)

}
