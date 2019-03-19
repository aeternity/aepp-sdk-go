package aeternity

import (
	"fmt"

	"github.com/aeternity/aepp-sdk-go/utils"
)

// SignEncodeTx sign and encode a transaction
func SignEncodeTx(kp *Account, txRaw []byte, networkID string) (signedEncodedTx, signedEncodedTxHash, signature string, err error) {
	// add the network_id to the transaction
	msg := append([]byte(networkID), txRaw...)
	// sign the transaction
	sigRaw := kp.Sign(msg)
	if err != nil {
		return
	}
	// encode the message using rlp
	rlpTxRaw, err := createSignedTransaction(txRaw, [][]byte{sigRaw})
	// encode the rlp message with the prefix
	signedEncodedTx = Encode(PrefixTransaction, rlpTxRaw)
	// compute the hash
	rlpTxHashRaw, err := hash(rlpTxRaw)
	signedEncodedTxHash = Encode(PrefixTransactionHash, rlpTxHashRaw)
	// encode the signature
	signature = Encode(PrefixSignature, sigRaw)
	return
}

func createSignedTransaction(txRaw []byte, signatures [][]byte) (rlpRawMsg []byte, err error) {
	// encode the message using rlp
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagSignedTransaction,
		rlpMessageVersion,
		signatures,
		txRaw,
	)
	return
}

// SpendTx create a spend transaction
func SpendTx(senderID, recipientID string, amount, fee utils.BigInt, payload string, ttl, nonce uint64) (rlpRawMsg []byte, err error) {
	// build id for the sender
	sID, err := buildIDTag(IDTagAccount, senderID)
	if err != nil {
		return
	}
	// build id for the recipient
	rID, err := buildIDTag(IDTagAccount, recipientID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagSpendTransaction,
		rlpMessageVersion,
		sID,
		rID,
		amount.Bytes(),
		fee.Bytes(),
		ttl,
		nonce,
		[]byte(payload))
	return
}

// NamePreclaimTx build a preclaim transaction
func NamePreclaimTx(accountID, commitmentID string, fee uint64, ttl, nonce uint64) (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, accountID)
	if err != nil {
		return
	}
	// build id for the committment
	cID, err := buildIDTag(IDTagCommitment, commitmentID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServicePreclaimTransaction,
		rlpMessageVersion,
		aID,
		nonce,
		cID,
		uint64(fee),
		ttl)
	return
}

// NameClaimTx build a preclaim transaction
func NameClaimTx(accountID, name string, nameSalt, fee uint64, ttl, nonce uint64) (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, accountID)
	if err != nil {
		return
	}
	// build id for the sender
	nID, err := buildIDTag(IDTagName, name)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceClaimTransaction,
		rlpMessageVersion,
		aID,
		nonce,
		nID,
		uint64(nameSalt),
		uint64(fee),
		ttl)
	return
}

func buildPointers(pointers []string) (ptrs [][]uint8, err error) {
	// TODO: handle errors
	ptrs = make([][]uint8, len(pointers))
	for i, p := range pointers {
		switch GetHashPrefix(p) {
		case PrefixAccountPubkey:
			pID, err := buildIDTag(IDTagName, p)
			ptrs[i] = pID
			if err != nil {
				break
			}
		case PrefixOraclePubkey:
			pID, err := buildIDTag(IDTagOracle, p)
			ptrs[i] = pID
			if err != nil {
				break
			}
		default:
			err = fmt.Errorf("Invalid ID %v for pointers", p)
		}
	}
	return
}

// NameUpdateTx build an update name transaction
func NameUpdateTx(accountID, nameID string, pointers []string, nameTTL, clientTTL uint64, fee uint64, ttl, nonce uint64) (rlpRawMsg []byte, err error) {
	// build id for the sender
	aID, err := buildIDTag(IDTagAccount, accountID)
	if err != nil {
		return
	}
	// build id for the sender
	nID, err := buildIDTag(IDTagName, nameID)
	if err != nil {
		return
	}
	// build id for pointers
	ptrs, err := buildPointers(pointers)
	if err != nil {
		return
	}

	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagNameServiceClaimTransaction,
		rlpMessageVersion,
		aID,
		nonce,
		nID,
		uint64(nameTTL),
		ptrs,
		uint64(clientTTL),
		uint64(fee),
		ttl)
	return
}

// OracleRegisterTx register an oracle
func OracleRegisterTx(accountID string, accountNonce uint64, querySpec, responseSpec string, queryFee uint64, ttlType string, ttlValue uint64, fee utils.BigInt, ttl, abiVersion uint64) (rlpRawMsg []byte, err error) {
	// build id for the account
	aID, err := buildIDTag(IDTagAccount, accountID)
	if err != nil {
		return
	}
	// create the transaction
	rlpRawMsg, err = buildRLPMessage(
		ObjectTagOracleRegisterTransaction,
		rlpMessageVersion,
		aID,
		accountNonce,
		[]byte(querySpec),
		[]byte(responseSpec),
		queryFee,
		[]byte(ttlType),
		ttlValue,
		fee.Bytes(),
		ttl,
		abiVersion)
	return
}

// OracleExtendTx extend an oracle's lifetime
func OracleExtendTx(oracleID string, accountNonce, ttlType, ttlValue uint64, fee utils.BigInt, ttl uint64) (rlpRawMsg []byte, err error) {
	aID, err := buildIDTag(IDTagOracle, oracleID)
	if err != nil {
		return
	}

	rlpRawMsg, err = buildRLPMessage(
		ObjectTagOracleExtendTransaction,
		rlpMessageVersion,
		aID,
		accountNonce,
		ttlType,
		ttlValue,
		fee.Bytes(),
		ttl)
	return
}
