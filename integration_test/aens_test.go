package integrationtest

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

func getNameEntry(aeClient *aeternity.Client, name string) (responseJSON string, err error) {
	response, err := aeClient.APIGetNameEntryByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, _ := response.MarshalBinary()
	responseJSON = string(r)
	return responseJSON, nil
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
	acc, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	aeClient := aeternity.NewClient(nodeURL, false)
	aensAlice := aeternity.Aens{Client: aeClient, Account: acc}

	aeternity.Config.Node.NetworkID = networkID

	name := randomName(6)
	fmt.Println("Testing with name: ", name)
	// Preclaim the name
	fmt.Println("PreclaimTx")
	preclaimTx, salt, err := aensAlice.NamePreclaimTx(name, aeternity.Config.Client.Fee)
	if err != nil {
		fmt.Println(err)
		return
	}
	preclaimTxStr, _ := aeternity.BaseEncodeTx(&preclaimTx)
	fmt.Println("PreclaimTx and Salt:", preclaimTxStr, salt)

	hash, err := signBroadcast(preclaimTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NamePreclaimTx", hash)

	// Wait for a bit
	_ = waitForTransaction(aeClient, hash)

	// Claim the name
	fmt.Println("NameClaimTx")
	claimTx, err := aensAlice.NameClaimTx(name, *salt, aeternity.Config.Client.Fee)
	if err != nil {
		fmt.Println(err)
		return
	}
	claimTxStr, _ := aeternity.BaseEncodeTx(&claimTx)
	fmt.Println("ClaimTx:", claimTxStr)

	hash, err = signBroadcast(claimTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameClaimTx")

	// Wait for a bit
	_ = waitForTransaction(aeClient, hash)

	// Verify that the name exists
	entryAfterNameClaim, err := getNameEntry(aeClient, name)
	fmt.Println(entryAfterNameClaim)

	// Update the name, make it point to something
	fmt.Println("NameUpdateTx")
	updateTx, err := aensAlice.NameUpdateTx(name, acc.Address)
	updateTxStr, _ := aeternity.BaseEncodeTx(&updateTx)
	fmt.Println("UpdateTx:", updateTxStr)

	_, err = signBroadcast(updateTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameUpdateTx")

	// Verify that the name was updated
	// Sleep a little, it takes time for the entry update to show up
	fmt.Printf("Sleeping a bit before querying /names/%s...\n", name)
	time.Sleep(1000 * time.Millisecond)
	entryAfterNameUpdate, _ := getNameEntry(aeClient, name)
	fmt.Println(entryAfterNameUpdate)

	if !strings.Contains(entryAfterNameUpdate, acc.Address) {
		t.Errorf("The AENS entry should now point to %s but doesn't: %s", acc.Address, entryAfterNameUpdate)
	}

	// Transfer the name to a recipient
	fmt.Println("NameTransferTx")
	acc2, err := aeternity.AccountFromHexString(recipientPrivateKey)
	if err != nil {
		t.Error(err)
	}
	transferTx, err := aensAlice.NameTransferTx(name, acc2.Address)
	transferTxStr, _ := aeternity.BaseEncodeTx(&transferTx)
	fmt.Println("TransferTx:", transferTxStr)
	hash, err = signBroadcast(transferTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameTransferTx")
	// Wait for a bit
	_ = waitForTransaction(aeClient, hash)

	// Receiver updates the name, makes it point to himself
	aeClient2 := aeternity.NewClient(nodeURL, false)
	aensBob := aeternity.Aens{Client: aeClient, Account: acc2}

	fmt.Println("NameUpdateTx Signed By Recipient")
	updateTx2, err := aensBob.NameUpdateTx(name, acc2.Address)
	updateTx2Str, _ := aeternity.BaseEncodeTx(&updateTx2)
	fmt.Println("UpdateTx:", updateTx2Str)

	_, err = signBroadcast(updateTx2Str, acc2, aeClient2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameUpdateTx Signed By Recipient")

	// Revoke the name - shouldn't work because it is signed by the sender, who no longer owns the address
	fmt.Println("NameRevokeTx")
	revokeTx, err := aensAlice.NameRevokeTx(name, acc.Address)
	revokeTxStr, _ := aeternity.BaseEncodeTx(&revokeTx)
	fmt.Println("RevokeTx:", revokeTxStr)
	hash, err = signBroadcast(revokeTxStr, acc, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameRevokeTx")
	// Wait for a bit
	revokeTxShouldHaveFailed := waitForTransaction(aeClient, hash)
	if revokeTxShouldHaveFailed == nil {
		t.Error("After transferring the name to Recipient, the Sender should not have been able to revoke the name")
	} else {
		fmt.Println(revokeTxShouldHaveFailed)
	}

	// Revoke the name - signed by the recipient
	fmt.Println("NameRevokeTx Signed By Recipient")
	revokeTx2, err := aensBob.NameRevokeTx(name, acc2.Address)
	revokeTx2Str, _ := aeternity.BaseEncodeTx(&revokeTx2)
	fmt.Println("RevokeTx Signed By Recipient:", revokeTx2Str)
	hash, err = signBroadcast(revokeTx2Str, acc2, aeClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Signed & Broadcasted NameRevokeTx Signed By Recipient")
	// Wait for a bit
	_ = waitForTransaction(aeClient, hash)

}
