package aeternity

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	apiclient "github.com/aeternity/aepp-sdk-go/generated/client"
	"github.com/aeternity/aepp-sdk-go/generated/client/external"
	"github.com/aeternity/aepp-sdk-go/generated/models"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func urlComponents(url string) (host string, schemas []string) {
	p := strings.Split(url, "://")
	if len(p) == 1 {
		host = p[0]
		schemas = []string{"http"}
		return
	}
	host = p[1]
	schemas = []string{p[0]}
	return
}

// GetTTL returns the chain height + offset
func (ae *Ae) GetTTL(offset uint64) (height uint64, err error) {
	return getTTL(ae.Node, offset)
}

// GetNextNonce retrieves the current accountNonce for an account + 1
func (ae *Ae) GetNextNonce(accountID string) (nextNonce uint64, err error) {
	return getNextNonce(ae.Node, accountID)
}

// GetTTLNonce is a convenience function that combines GetTTL() and GetNextNonce()
func (ae *Ae) GetTTLNonce(accountID string, offset uint64) (txTTL, accountNonce uint64, err error) {
	return getTTLNonce(ae.Node, accountID, offset)
}

func getTTL(node *apiclient.Node, offset uint64) (height uint64, err error) {
	kb, err := getTopBlock(node)
	if err != nil {
		return
	}

	if kb.KeyBlock == nil {
		height = *kb.MicroBlock.Height + offset
	} else {
		height = *kb.KeyBlock.Height + offset
	}

	return
}

func getNextNonce(node *apiclient.Node, accountID string) (nextNonce uint64, err error) {
	a, err := getAccount(node, accountID)
	if err != nil {
		return
	}
	nextNonce = *a.Nonce + 1
	return
}

func getTTLNonce(node *apiclient.Node, accountID string, offset uint64) (height uint64, nonce uint64, err error) {
	height, err = getTTL(node, offset)
	if err != nil {
		return
	}

	nonce, err = getNextNonce(node, accountID)
	if err != nil {
		return
	}
	return
}

// waitForTransaction to appear on the chain
func waitForTransaction(nodeClient *apiclient.Node, txHash string) (blockHeight uint64, blockHash string, err error) {
	// caclulate the date for the timeout
	ctm := Config.Tuning.ChainTimeout
	tm := time.Now().Add(time.Millisecond * time.Duration(ctm))
	// start querying the transaction
	for {
		if time.Now().After(tm) {
			// TODO: should use the chain height instead of a timeout
			err = fmt.Errorf("Timeout waiting for the transaction to appear")
			break // timeout execed
		}
		tx, err := getTransactionByHash(nodeClient, txHash)
		if err != nil {
			break
		}
		if len(tx.BlockHash) > 0 {
			blockHeight = *tx.BlockHeight
			blockHash = fmt.Sprint(tx.BlockHash)
			break
		}
		time.Sleep(time.Millisecond * time.Duration(Config.Tuning.ChainPollInteval))
	}
	return
}

// BroadcastTransaction recalculates the transaction hash and sends the transaction to the node.
func (ae *Ae) BroadcastTransaction(txSignedBase64 string) (err error) {
	// Get back to RLP to calculate txhash
	txRLP, _ := Decode(txSignedBase64)

	// calculate the hash of the decoded txRLP
	rlpTxHashRaw, _ := hash(txRLP)
	// base58/64 encode the hash with the th_ prefix
	signedEncodedTxHash := Encode(PrefixTransactionHash, rlpTxHashRaw)

	// send it to the network
	err = postTransaction(ae.Node, txSignedBase64, signedEncodedTxHash)
	return
}

// NamePreclaimTx creates a name preclaim transaction and salt (required for claiming)
// It should return the Tx struct, not the base64 encoded RLP, to ease subsequent inspection.
func (n *Aens) NamePreclaimTx(name string, fee utils.BigInt) (tx NamePreclaimTx, nameSalt *utils.BigInt, err error) {
	txTTL, accountNonce, err := getTTLNonce(n.nodeClient, n.owner.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	// calculate the commitment and get the preclaim salt
	// since the salt is 32 bytes long, you must use a big.Int to convert it into an integer
	cm, nameSalt, err := generateCommitmentID(name)
	if err != nil {
		return NamePreclaimTx{}, utils.NewBigInt(), err
	}

	// build the transaction
	tx = NewNamePreclaimTx(n.owner.Address, cm, fee, txTTL, accountNonce)
	if err != nil {
		return NamePreclaimTx{}, utils.NewBigInt(), err
	}

	return
}

// NameClaimTx creates a claim transaction
func (n *Aens) NameClaimTx(name string, nameSalt utils.BigInt, fee utils.BigInt) (tx NameClaimTx, err error) {
	txTTL, accountNonce, err := getTTLNonce(n.nodeClient, n.owner.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	// create the transaction
	tx = NewNameClaimTx(n.owner.Address, name, nameSalt, fee, txTTL, accountNonce)

	return tx, err
}

// NameUpdateTx perform a name update
func (n *Aens) NameUpdateTx(name string, targetAddress string) (tx NameUpdateTx, err error) {
	txTTL, accountNonce, err := getTTLNonce(n.nodeClient, n.owner.Address, Config.Client.TTL)
	if err != nil {
		return
	}

	encodedNameHash := Encode(PrefixName, Namehash(name))
	absNameTTL, err := getTTL(n.nodeClient, Config.Client.Names.NameTTL)
	if err != nil {
		return NameUpdateTx{}, err
	}
	// create and sign the transaction
	tx = NewNameUpdateTx(n.owner.Address, encodedNameHash, []string{targetAddress}, absNameTTL, Config.Client.Names.ClientTTL, Config.Client.Names.UpdateFee, txTTL, accountNonce)

	return
}

// PrintGenerationByHeight utility function to print a generation by it's height
func (ae *Ae) PrintGenerationByHeight(height uint64) {
	p := external.NewGetGenerationByHeightParams().WithHeight(height)
	if r, err := ae.External.GetGenerationByHeight(p); err == nil {
		PrintObject("generation", r.Payload)
		// search for transaction in the microblocks
		for _, mbh := range r.Payload.MicroBlocks {
			// get the microblok
			mbhs := fmt.Sprint(mbh)
			p := external.NewGetMicroBlockTransactionsByHashParams().WithHash(mbhs)
			r, err := ae.External.GetMicroBlockTransactionsByHash(p)
			if err != nil {
				Pp("Error:", err)
			}
			// go through all the hashes
			for _, btx := range r.Payload.Transactions {
				p := external.NewGetTransactionByHashParams().WithHash(fmt.Sprint(btx.Hash))
				if r, err := ae.External.GetTransactionByHash(p); err == nil {
					PrintObject("transaction", r.Payload)
				}
			}
		}
	} else {
		switch err.(type) {
		case *external.GetGenerationByHashBadRequest:
			PrintError("Bad request:", err.(*external.GetGenerationByHashBadRequest).Payload)
		case *external.GetGenerationByHashNotFound:
			PrintError("Block not found:", err.(*external.GetGenerationByHashNotFound).Payload)
		}
	}
}

// WaitForTransactionUntilHeight waits for a transaction until heightLimit (inclusive) is reached
func (ae *Ae) WaitForTransactionUntilHeight(height uint64, txHash string) (blockHeight uint64, blockHash, microBlockHash string, tx *models.GenericSignedTx, err error) {
	kb, err := getCurrentKeyBlock(ae.Node)
	if err != nil {
		return
	}
	// current height
	targetHeight := *kb.Height
	nextHeight := *kb.Height
	// hold the generation
	var g *models.Generation

Main:
	for {
		// check the top height
		if targetHeight > height {
			err = fmt.Errorf("Transaction %s not found, height %d", txHash, height)
			break
		}
		// get the generation of the targetHeight
		p := external.NewGetGenerationByHeightParams().WithHeight(targetHeight)
		r, err := ae.External.GetGenerationByHeight(p)
		if err != nil {
			break
		}
		g = r.Payload
		// search for transaction in the microblocks
		for _, mbh := range g.MicroBlocks {
			// get the microblok
			mbhs := fmt.Sprint(mbh)
			p := external.NewGetMicroBlockTransactionsByHashParams().WithHash(mbhs)
			r, mErr := ae.External.GetMicroBlockTransactionsByHash(p)
			if mErr != nil {
				// TODO: err will still be nil outside this scope. Consider refactoring whole function.
				err = mErr
				break Main
			}
			// go through all the hashes
			for _, btx := range r.Payload.Transactions {
				if fmt.Sprint(btx.Hash) == txHash {
					// transaction found !!
					blockHash = fmt.Sprint(g.KeyBlock.Hash)
					blockHeight = *g.KeyBlock.Height
					microBlockHash = mbhs
					tx = btx
					break Main
				}
			}
		}
		// here we want to query one more time the current generation
		// before switching to the next one in case microblocks have been added meanwhile
		// update targetHeight
		if nextHeight > targetHeight {
			targetHeight = nextHeight
		}
		// update nextHeight
		kb, err = getCurrentKeyBlock(ae.Node)
		if err != nil {
			break
		}
		nextHeight = *kb.Height
	}

	return
}

// StoreAccountToKeyStoreFile store an account to a json file
func StoreAccountToKeyStoreFile(account *Account, password, walletName string) (filePath string, err error) {
	// keystore will be saved in current directory
	basePath, _ := os.Getwd()

	// generate the keystore file
	jks, err := KeystoreSeal(account, password)
	if err != nil {
		return
	}
	// build the wallet path
	filePath = filepath.Join(basePath, keyFileName(account.Address))
	if len(walletName) > 0 {
		filePath = filepath.Join(basePath, walletName)
	}
	// write the file to disk
	err = ioutil.WriteFile(filePath, jks, 0600)
	return
}

// LoadAccountFromKeyStoreFile load file from the keystore
func LoadAccountFromKeyStoreFile(keyFile, password string) (account *Account, err error) {
	// find out the real path of the wallet
	filePath, err := GetWalletPath(keyFile)
	if err != nil {
		return
	}
	// load the json file
	jks, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	// decrypt keystore
	account, err = KeystoreOpen(jks, password)
	return
}

// GetWalletPath try to locate a wallet
func GetWalletPath(path string) (walletPath string, err error) {
	// if file exists then load the file
	if _, err = os.Stat(path); !os.IsNotExist(err) {
		walletPath = path
		return
	}
	return
}

// SignEncodeTxStr sign and encode a transaction format as string (ex. tx_xyz)
func SignEncodeTxStr(kp *Account, tx string, networkID string) (signedEncodedTx, signedEncodedTxHash, signature string, err error) {
	txRaw, err := Decode(tx)
	if err != nil {
		fmt.Println("Error decoding tx from base64")
		os.Exit(1)
	}

	signedEncodedTx, signedEncodedTxHash, signature, err = SignEncodeTx(kp, txRaw, networkID)
	return
}

// VerifySignedTx verifies a tx_ with signature
func VerifySignedTx(accountID string, txSigned string, networkID string) (valid bool, err error) {
	txRawSigned, _ := Decode(txSigned)
	txRLP := DecodeRLPMessage(txRawSigned)

	// RLP format of signed signature: [[Tag], [Version], [Signatures...], [Transaction]]
	tx := txRLP[3].([]byte)
	txSignature := txRLP[2].([]interface{})[0].([]byte)

	msg := append([]byte(networkID), tx...)

	valid, err = Verify(accountID, msg, txSignature)
	if err != nil {
		return
	}
	return
}
