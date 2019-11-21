package clib
/*
struct SpendTx {
	char *SenderID;
	char *RecipientID;
	unsigned long long int Amount;
	unsigned long long int Fee;
	char *Payload;
	unsigned long long int TTL;
	unsigned long long int Nonce;
};
*/
import "C"

//export GetTx
func GetTx() C.struct_SpendTx {
	return C.struct_SpendTx{
		SenderID: C.CString("ak_Hello"),
		RecipientID: C.CString("ak_Hello"),
		Amount: C.ulonglong(64),
		Fee:C.ulonglong(64),
		Payload: C.CString("a payload"),
		TTL:C.ulonglong(64),
		Nonce:C.ulonglong(64),
	}
}