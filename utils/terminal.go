package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/aeternity/aepp-sdk-go/models"
)

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

// Left left-pads the string with pad up to len runes
// len may be exceeded if
func left(str string, length int, pad string) string {
	return times(pad, length-len(str)) + str
}

// Right right-pads the string with pad up to len runes
func right(str string, length int, pad string) string {
	return str + times(pad, length-len(str))
}

// Pp pretty print
func Pp(data ...interface{}) {
	for i := 0; i < len(data); i += 2 {
		fmt.Println(right(fmt.Sprintf("%v", data[i]), 50, "_"), data[i+1])
	}
}

func printBlock(hash models.EncodedHash, o models.Header) {
	Pp(
		"Block Hash", hash,
		"Block Height", o.Height,
		"Previous block hash", o.PrevHash,
		"Miner", o.Miner,
		"Beneficiary", o.Beneficiary,
		"State hash", o.StateHash,
		"Time", time.Unix(0, o.Time*int64(time.Millisecond)).Format(time.RFC3339),
		"Transactions", o.TxsHash,
	)
}

func printNodeVersion(o *models.Version) {
	Pp(
		"Node version", o.Version,
		"Genesis hash", o.GenesisHash,
		"Revision", o.Revision,
	)
}

func printTx(o models.SingleTxObject) {
	switch o.(type) {
	case *models.SingleTxJSON:
		Pp(
			"Transaction hash", o.(*models.SingleTxJSON).Transaction.Hash,
			"Signatures", strings.Join(o.(*models.SingleTxJSON).Transaction.Signatures, ", "),
			"Message encoding", "json",
		)
	case *models.SingleTxMsgPack:
		Pp(
			"Transaction hash", o.(*models.SingleTxMsgPack).Transaction.Tx,
			"Encoding", "msgpack",
		)
	}
}

func printName(o *models.NameEntry) {
	Pp(
		"Name", o.Name,
		"Name hash", o.NameHash,
		"Name TTL", o.NameTTL,
	)
}

// PrintObject pretty print an object obtained from the api
func PrintObject(i interface{}) {
	PrintObjectT("", i)
}

// PrintError print error
func PrintError(code string, e *models.Error) {
	Pp(code, e.Reason)
}

// PrintObjectT pretty print an object obtained from the api with a title
func PrintObjectT(title string, i interface{}) {
	if len(title) > 0 {
		fmt.Println(title)
	}
	switch i.(type) {
	case *models.GenericBlock:
		printBlock(i.(*models.GenericBlock).Hash, i.(*models.GenericBlock).Header)
	case *models.Top:
		printBlock(models.EncodedHash(i.(*models.Top).Hash), i.(*models.Top).Header)
	case *models.Version:
		printNodeVersion(i.(*models.Version))
	case models.SingleTxObject:
		printTx(i.(models.SingleTxObject))
	case *models.NameEntry:
		printName(i.(*models.NameEntry))
	default:
		fmt.Printf("Pretty printer not available for type %v", i)
	}
}
