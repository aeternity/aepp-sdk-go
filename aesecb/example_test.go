
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aesecb

 
import (
 	"os"
 )
 
func ExampleNewECBDecrypter() {
	key := []byte("example key 1234")
	ciphertext, _ := hex.DecodeString("e1cdb90013f76bdf10c3d76b40e5e164")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	// ECB mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewECBDecrypter(block)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	fmt.Printf("%s\n", ciphertext)
	// Output: exampleplaintext
}

func ExampleNewECBEncrypter() {
	key := []byte("example key 1234")
	plaintext := []byte("exampleplaintext")

	// ECB mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewECBEncrypter(block)
	mode.CryptBlocks(ciphertext, plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	fmt.Printf("%x\n", ciphertext)
}

 func ExampleNewCBCDecrypter() {
 	key := []byte("example key 1234")
 	ciphertext, _ := hex.DecodeString("f363f3ccdcb12bb883abf484ba77d9cd7d32b5baecb3d4b1b3e0e4beffdb3ded")