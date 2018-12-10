package aeternity

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

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

// getAbsoluteHeight return the chain height adding the offset
func getAbsoluteHeight(epochCli *apiclient.Epoch, offset uint64) (height uint64, err error) {
	kb, err := getTopBlock(epochCli)
	if err != nil {
		return
	}

	if kb.KeyBlock == nil {
		height = *kb.MicroBlock.Height + offset
	}

	return
}

// getNextNonce retrieve the next nonce for an account
// it has to query the chain to do so
func getNextNonce(epochCli *apiclient.Epoch, acccount *Account) (nextNonce uint64, err error) {
	a, err := getAccount(epochCli, acccount.Address)
	if err != nil {
		return
	}
	nextNonce = *a.Nonce + 1
	return
}

// waitForTransaction to appear on the chain
func waitForTransaction(epochCli *apiclient.Epoch, txHash string) (blockHeight uint64, blockHash string, err error) {
	// caclulate the date for the timeout
	ctm := Config.P.Tuning.ChainTimeout
	tm := time.Now().Add(time.Millisecond * time.Duration(ctm))
	// start querying the transaction
	for {
		if time.Now().After(tm) {
			// TODO: should use the chain height instead of a timeout
			err = fmt.Errorf("Timeout waiting for the transaction to appear")
			break // timeout execed
		}
		tx, err := getTransactionByHash(epochCli, txHash)
		if err != nil {
			break
		}
		if len(tx.BlockHash) > 0 {
			blockHeight = *tx.BlockHeight
			blockHash = fmt.Sprint(tx.BlockHash)
			break
		}
		time.Sleep(time.Millisecond * time.Duration(Config.P.Tuning.ChainPollInteval))
	}
	return
}

// Spend transfer tokens from an account to another
func (w *Wallet) Spend(recipientAddress string, amount int64, message string) (tx, txHash, signature string, ttl uint64, nonce uint64, err error) {
	// calculate the absolute ttl for the transaction
	ttl, err = getAbsoluteHeight(w.epochCli, Config.P.Client.TTL)
	if err != nil {
		return
	}
	// create the spend transaction

	nonce, err = getNextNonce(w.epochCli, w.owner)
	if err != nil {
		return
	}
	spendTxRaw, err := createSpendTransaction(w.owner.Address, recipientAddress, message, amount, Config.P.Client.Fee, ttl, nonce)
	if err != nil {
		return
	}
	// sign the transaction
	tx, txHash, signature, err = SignEncodeTx(w.owner, spendTxRaw)
	if err != nil {
		return
	}
	// post the transaction
	err = postTransaction(w.epochCli, tx, txHash)
	return
}

// naming
func computeCommitmentID(name string) (ch string, salt []byte, err error) {
	salt, err = randomBytes(32)
	if err != nil {
		return
	}
	// TODO: this is done using the api (concatenating )
	nh := append(namehash(name), salt...)
	nh, _ = hash(nh)
	// nh := namehash(name)
	ch = encode(PrefixCommitment, nh)
	return
}

// NamePreclaim post a preclaim transaction to the chain
func (n *Aens) NamePreclaim(name string) (tx, txHash, signature string, ttl uint64, nonce uint64, nameSalt int64, err error) {
	// get the ttl offset
	ttl, err = getAbsoluteHeight(n.epochCli, Config.P.Client.TTL)
	if err != nil {
		return
	}
	// calculate the commitment and get the preclaim salt
	cm, salt, err := computeCommitmentID(name)
	if err != nil {
		return
	}
	// convert the stalt to a int64
	nameSalt = int64(binary.BigEndian.Uint64(salt))
	// get the account nonce
	nonce, err = getNextNonce(n.epochCli, n.owner)
	if err != nil {
		return
	}
	// build the transaction
	tx, txHash, signature, err = namePreclaimTxSigned(n.owner, cm, Config.P.Client.Names.PreClaimFee, ttl, nonce)
	if err != nil {
		return
	}
	// post transaction to the chain
	err = postTransaction(n.epochCli, tx, txHash)
	return
}

// NameClaim perform a name claiming
func (n *Aens) NameClaim(name string, nameSalt int64) (tx, txHash, signature string, ttl uint64, nonce uint64, err error) {
	// get the ttl offset
	ttl, err = getAbsoluteHeight(n.epochCli, Config.P.Client.TTL)
	if err != nil {
		return
	}
	// get the account nonce
	nonce, err = getNextNonce(n.epochCli, n.owner)
	if err != nil {
		return
	}
	//TODO: do we need the encoded name here?
	// encodedName := encodeP(PrefixNameHash, []byte(name))
	prefix := HashPrefix(name[0:3])
	encodedName := encode(prefix, []byte(name))
	// sign the above transaction with the private key
	tx, txHash, signature, err = nameClaimTxSigned(n.owner, encodedName, nameSalt, Config.P.Client.Names.ClaimFee, ttl, nonce)
	if err != nil {
		return
	}
	// post transaction to the chain
	err = postTransaction(n.epochCli, tx, txHash)
	return
}

// NameUpdate perform a name update
func (n *Aens) NameUpdate(name string, targetAddress string) (tx, txHash, signature string, ttl uint64, nonce uint64, err error) {
	ttl, err = getAbsoluteHeight(n.epochCli, Config.P.Client.TTL)
	if err != nil {
		return
	}
	// get the account nonce
	nonce, err = getNextNonce(n.epochCli, n.owner)
	if err != nil {
		return
	}

	encodedNameHash := encode(PrefixName, namehash(name))
	absClientTTL, err := getAbsoluteHeight(n.epochCli, Config.P.Client.Names.ClientTTL)
	if err != nil {
		return
	}
	absNameTTL, err := getAbsoluteHeight(n.epochCli, Config.P.Client.Names.NameTTL)
	if err != nil {
		return
	}
	// create and sign the transaction
	tx, txHash, signature, err = nameUpdateTxSigned(n.owner, encodedNameHash, []string{targetAddress}, absNameTTL, absClientTTL, Config.P.Client.Names.UpdateFee, ttl, nonce)
	if err != nil {
		return
	}
	// post transaction to the chain
	err = postTransaction(n.epochCli, tx, txHash)
	return
}

// OracleRegister register an oracle
// TODO: not implemented
func (o *Oracle) OracleRegister(queryFormat, responseFormat string) (tx, txHash, signature string, ttl int64, nonce uint64, err error) {
	// TODO: specs incomplete
	//txOracleCreate(o.owner.Address, queryFormat, responseFormat, Config.P.Client.Oracles.QueryFee, Config.P.Client.Oracles.Expires)
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

// WaitForTransactionUntillHeight waits for a transaction until heightLimit (inclusive) is reached
func (ae *Ae) WaitForTransactionUntillHeight(height uint64, txHash string) (blockHeight uint64, blockHash, microBlockHash string, tx *models.GenericSignedTx, err error) {
	kb, err := getCurrentKeyBlock(ae.Epoch)
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
		kb, err = getCurrentKeyBlock(ae.Epoch)
		if err != nil {
			break
		}
		nextHeight = *kb.Height
	}

	return
}

// StoreAccountToKeyStoreFile store an account to a json file
func StoreAccountToKeyStoreFile(account *Account, password, walletName string) (filePath string, err error) {
	// keys are in the same folder as config
	basePath := filepath.Join(filepath.Dir(Config.ConfigPath), "accounts")
	err = os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		return
	}
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
	// check by name in the default location
	path = filepath.Join(Config.KeysFolder, path)
	if _, err = os.Stat(path); !os.IsNotExist(err) {
		walletPath = path
	}
	return
}

// ListWallets in the default wallet paths
func ListWallets() (wallets []string, err error) {
	wallets, err = filepath.Glob(filepath.Join(Config.KeysFolder, "*"))
	return
}
