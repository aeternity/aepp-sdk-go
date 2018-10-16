package aeternity

import (
  "fmt"

  "github.com/aeternity/aepp-sdk-go/rlp"
)

// SignEncodeTx sign and encode a transaction
func SignEncodeTx(kp *Account, txRaw []byte) (signedEncodedTx, signedEncodedTxHash, signature string, err error) {
  // sign the transaction
  sigRaw := kp.Sign(txRaw)
  if err != nil {
    return
  }
  // encode the message using rlp
  rlpTxRaw, err := createSignedTransaction(txRaw, [][]byte{sigRaw})
  // encode the rlp message with the prefix
  signedEncodedTx = encodeP(PrefixTx, rlpTxRaw)
  // compute the hash
  rlpTxHashRaw, err := hash(rlpTxRaw)
  signedEncodedTxHash = encodeP(PrefixTxHash, rlpTxHashRaw)
  // encode the signature
  signature = encodeP(PrefixSignature, sigRaw)
  return
}

func buildRLPMessage(tag uint, version uint, fields ...interface{}) (rlpRawMsg []byte, err error) {
  // create a message of the transaction and signature
  data := []interface{}{tag, version}
  data = append(data, fields...)
  // fmt.Printf("TX %#v\n\n", data)
  // encode the message using rlp
  rlpRawMsg, err = rlp.EncodeToBytes(data)
  // fmt.Printf("ENCODED %#v\n\n", data)
  return
}

// buildIDTag assemble an id() object
// see https://github.com/aeternity/protocol/blob/epoch-v0.22.0/serializations.md#the-id-type
func buildIDTag(IDTag uint8, encodedHash string) (v []uint8, err error) {
  raw, err := decode(encodedHash)
  v = []uint8{IDTag}
  for _, x := range raw {
    v = append(v, uint8(x))
  }
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

// createSpendTransaction
// see https://github.com/aeternity/protocol/blob/epoch-v0.22.0/serializations.md
func createSpendTransaction(senderID, recipientID, payload string, amount, fee int64, ttl, nonce uint64) (rlpRawMsg []byte, err error) {
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
    uint64(amount),
    uint64(fee),
    ttl,
    nonce,
    []byte(payload))
  return
}

// namePreclaimTx build a preclaim transaction
// see https://github.com/aeternity/protocol/blob/epoch-v0.22.0/serializations.md#name-service-preclaim-transaction
func namePreclaimTx(accountID, commitmentID string, fee int64, ttl, nonce uint64) (rlpRawMsg []byte, err error) {
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

func namePreclaimTxSigned(account *Account, commitmentID string, fee int64, ttl, nonce uint64) (signedEncodedTx, signedEncodedTxHash, signature string, err error) {
  // this is the transaction to sign
  rawTx, err := namePreclaimTx(account.Address, commitmentID, fee, ttl, nonce)
  if err != nil {
    return
  }
  // sign the above transaction with the private key
  signedEncodedTx, signedEncodedTxHash, signature, err = SignEncodeTx(account, rawTx)
  return
}

// nameClaimTx build a preclaim transaction
// see https://github.com/aeternity/protocol/blob/epoch-v0.22.0/serializations.md#name-service-claim-transaction
func nameClaimTx(accountID, name string, nameSalt, fee int64, ttl, nonce uint64) (rlpRawMsg []byte, err error) {
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

func nameClaimTxSigned(account *Account, name string, nameSalt, fee int64, ttl, nonce uint64) (signedEncodedTx, signedEncodedTxHash, signature string, err error) {
  // this is the transaction to sign
  rawTx, err := nameClaimTx(account.Address, name, nameSalt, fee, ttl, nonce)
  if err != nil {
    return
  }
  // sign the above transaction with the private key
  signedEncodedTx, signedEncodedTxHash, signature, err = SignEncodeTx(account, rawTx)
  return
}

func buildPointers(pointers []string) (ptrs [][]uint8, err error) {
  // TODO: handle errors
  ptrs = make([][]uint8, len(pointers))
  for i, p := range pointers {
    switch GetHashPrefix(p) {
    case PrefixAccount:
      pID, err := buildIDTag(IDTagName, p)
      ptrs[i] = pID
      if err != nil {
        break
      }
    case PrefixOracle:
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

// nameUpdateTx build an update name transaction
// see https://github.com/aeternity/protocol/blob/epoch-v0.22.0/serializations.md#name-service-update-transaction
func nameUpdateTx(accountID, nameID string, pointers []string, nameTTL, clientTTL uint64, fee int64, ttl, nonce uint64) (rlpRawMsg []byte, err error) {
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

func nameUpdateTxSigned(account *Account, nameID string, pointers []string, nameTTL, clientTTL uint64, fee int64, ttl, nonce uint64) (signedEncodedTx, signedEncodedTxHash, signature string, err error) {
  // this is the transaction to sign
  rawTx, err := nameUpdateTx(account.Address, nameID, pointers, nameTTL, clientTTL, fee, ttl, nonce)
  if err != nil {

  }
  // sign the above transaction with the private key
  signedEncodedTx, signedEncodedTxHash, signature, err = SignEncodeTx(account, rawTx)
  return
}

// txOracleRegister
// see https://github.com/aeternity/protocol/blob/master/serializations.md#oracles
func txOracleRegister(accountID, queryFormat, responseFormat string, queryFee, expires int64) (rlpRawMsg []byte, err error) {
  // build id for the account
  aID, err := buildIDTag(IDTagAccount, accountID)
  if err != nil {
    return
  }
  // create the transaction
  rlpRawMsg, err = buildRLPMessage(
    ObjectTagOracle,
    rlpMessageVersion,
    aID,
    []byte(queryFormat),
    []byte(responseFormat),
    uint64(queryFee),
    uint64(expires))
  return
}
