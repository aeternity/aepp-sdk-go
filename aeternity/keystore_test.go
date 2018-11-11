package aeternity

import (
  "encoding/hex"
  "fmt"
  "testing"
)

func TestWrite(t *testing.T) {

  prv := []byte{225, 128, 98, 185, 25, 78, 104, 215, 238, 158, 73, 59, 202, 121, 33, 211, 236, 62, 1, 121, 152, 198, 219, 177, 15, 248, 248, 172, 85, 22, 105, 133, 201, 101, 152, 238, 118, 129, 16, 165, 224, 51, 1, 186, 46, 47, 63, 47, 70, 67, 232, 228, 202, 93, 46, 182, 144, 182, 8, 152, 185, 3, 23, 233}
  // prvHex := "e18062b9194e68d7ee9e493bca7921d3ec3e017998c6dbb10ff8f8ac55166985c96598ee768110a5e03301ba2e2f3f2f4643e8e4ca5d2eb690b60898b90317e9"
  // pub := []byte{201, 101, 152, 238, 118, 129, 16, 165, 224, 51, 1, 186, 46, 47, 63, 47, 70, 67, 232, 228, 202, 93, 46, 182, 144, 182, 8, 152, 185, 3, 23, 233}
  // pubAdd := "ak_2XhQw1o9UwvHNFTe1vCaLEDfUQv9Y4APSVRomFgQtTjHukMbdH"

  // ac, _ := NewAccount()
  // fmt.Println(ac.SigningKey)
  // fmt.Println(ac.SigningKeyToHexString())
  // fmt.Println(ac.SigningKey.Public())
  // fmt.Println(ac.Address)

  ac, _ := loadAccountFromPrivateKeyRaw(prv)

  type args struct {
    account  *Account
    password string
  }
  tests := []struct {
    name  string
    args  args
    match bool
  }{
    {
      "one",
      args{ac, "test"},
      true,
    },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      encryped, err := KeystoreCreate(tt.args.account, tt.args.password)
      if err != nil {
        t.Errorf("Error %s", err)
        return
      }
      fmt.Printf("%s", encryped)
      decrypted, err := KeystoreLoad(encryped, tt.args.password)
      if err != nil {
        t.Errorf("Error %s", err)
        return
      }
      a := hex.EncodeToString(decrypted.SigningKey)
      b := hex.EncodeToString(tt.args.account.SigningKey)
      if !tt.match && a == b {
        t.Errorf("Wanted match but no")
      }
    })

  }
}
