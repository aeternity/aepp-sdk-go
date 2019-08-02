package aeternity

func getRLPSerialized(tx1 string, tx2 string) ([]interface{}, []interface{}) {
	tx1Bytes, _ := Decode(tx1)
	tx1RLP := DecodeRLPMessage(tx1Bytes)
	tx2Bytes, _ := Decode(tx2)
	tx2RLP := DecodeRLPMessage(tx2Bytes)
	return tx1RLP, tx2RLP
}
