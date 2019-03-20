package aeternity

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aeternity/aepp-sdk-go/utils"

	apiclient "github.com/aeternity/aepp-sdk-go/generated/client"
	"github.com/aeternity/aepp-sdk-go/generated/client/external"
	"github.com/aeternity/aepp-sdk-go/generated/models"
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

// GetNextNonce retrieves the current nonce for an account + 1
func (ae *Ae) GetNextNonce(accountID string) (nextNonce uint64, err error) {
	return getNextNonce(ae.Node, accountID)
}

// GetTTLNonce is a convenience function that combines GetTTL() and GetNextNonce()
func (ae *Ae) GetTTLNonce(accountID string, offset uint64) (ttl, nonce uint64, err error) {
	ttl, err = ae.GetTTL(offset)
	if err != nil {
		return 0, 0, err
	}

	nonce, err = ae.GetNextNonce(accountID)
	if err != nil {
		return 0, 0, err
	}

	return ttl, nonce, nil
}

func getTTL(node *apiclient.Node, offset uint64) (height uint64, err error) {
	kb, err := getTopBlock(node)
	if err != nil {
		return
	}

	if kb.KeyBlock == nil {
		height = *kb.MicroBlock.Height + offset
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

// SpendTxStr creates an unsigned SpendTx but returns the base64 representation instead of an RLP bytestring
func SpendTxStr(sender, recipient string, amount, fee utils.BigInt, message string, ttl, nonce uint64) (base64Tx string, err error) {
	rlpUnsignedTx, err := SpendTx(sender, recipient, amount, fee, message, ttl, nonce)
	if err != nil {
		return
	}

	base64Tx = Encode(PrefixTransaction, rlpUnsignedTx)

	return base64Tx, err
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

// NamePreclaimTxStr creates a name preclaim transaction and nameSalt (required for claiming)
func (n *Aens) NamePreclaimTxStr(name string, ttl, nonce uint64) (tx string, nameSalt uint64, err error) {
	// calculate the commitment and get the preclaim salt
	cm, salt, err := computeCommitmentID(name)
	if err != nil {
		return "", 0, err
	}
	// convert the salt back into uint64 from binary
	nameSalt = binary.BigEndian.Uint64(salt)

	// build the transaction
	txRaw, err := NamePreclaimTx(n.owner.Address, cm, Config.Client.Names.PreClaimFee, ttl, nonce)
	if err != nil {
		return "", 0, err
	}

	tx = Encode(PrefixTransaction, txRaw)
	return
}

// NameClaimTxStr creates a claim transaction
func (n *Aens) NameClaimTxStr(name string, nameSalt, ttl, nonce uint64) (tx string, err error) {
	//TODO: do we need the encoded name here?
	// encodedName := encodeP(PrefixNameHash, []byte(name))
	prefix := HashPrefix(name[0:3])
	encodedName := Encode(prefix, []byte(name))
	// create the transaction
	txRaw, err := NameClaimTx(n.owner.Address, encodedName, nameSalt, Config.Client.Names.ClaimFee, ttl, nonce)
	if err != nil {
		return "", err
	}

	tx = Encode(PrefixTransaction, txRaw)
	return
}

// NameUpdateTxStr perform a name update
func (n *Aens) NameUpdateTxStr(name string, targetAddress string, ttl, nonce uint64) (tx string, err error) {
	encodedNameHash := Encode(PrefixName, namehash(name))
	absNameTTL, err := getTTL(n.nodeClient, Config.Client.Names.NameTTL)
	if err != nil {
		return "", err
	}
	// create and sign the transaction
	txRaw, err := NameUpdateTx(n.owner.Address, encodedNameHash, []string{targetAddress}, absNameTTL, Config.Client.Names.ClientTTL, Config.Client.Names.UpdateFee, ttl, nonce)
	if err != nil {
		return "", err
	}

	tx = Encode(PrefixTransaction, txRaw)
	return
}

// OracleRegister register an oracle
// TODO: not implemented
// func (o *Oracle) OracleRegisterTxStr(querySpec, responseSpec string, queryFee, ttlType, ttlValue, abiVersion, ttl, nonce uint64) (txRaw []byte, err error) {
// 	// TODO: specs incomplete
// 	txRaw, err := OracleRegisterTx(o.owner.Address, o.owner.Nonce, querySpec, responseSpec, queryFee, ttlType, ttlValue, )
// 	return
// }

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
	txRLP := decodeRLPMessage(txRawSigned)

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
