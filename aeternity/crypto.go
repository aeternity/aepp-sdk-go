package aeternity

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"

	"github.com/aeternity/aepp-sdk-go/aesecb"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ed25519"
)

// KeyPair holds the signing key and the aeternity account address
type KeyPair struct {
	SigningKey ed25519.PrivateKey
	Address    string
}

// Load parse a string into a private key
func Load(priv string) (kp *KeyPair, err error) {
	privb, err := hex.DecodeString(priv)
	if err != nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Invalid private key")
		}
	}()
	kp = &KeyPair{SigningKey: ed25519.PrivateKey(privb)}
	pubb := []byte(fmt.Sprintf("%s", kp.SigningKey.Public()))
	kp.Address = encodeP(PrefixAccount, pubb)
	return
}

func h(data []byte) []byte {
	d := sha256.New()
	d.Write(data)
	return d.Sum(nil)
}

// encode encode a byte array into base58 with chacksum
func encode(in []byte) string {
	c := h(h(in))
	return base58.Encode(append(in, c[0:4]...))
}

// encodeP encode a byte array into base58 with chacksum and a prefix
func encodeP(prefix HashPrefix, data []byte) string {
	return fmt.Sprint(prefix, encode(data))
}

// decode decode a string encoded with base58 + checksum to a byte array
func decode(in string) (out []byte, err error) {
	if len(in) >= 3 && string(in[2]) == "$" {
		in = in[3:]
	}
	raw := base58.Decode(in)
	if len(raw) < 5 {
		err = fmt.Errorf("Invalid input, %s cannot be decoded", in)
		return
	}
	out = raw[:len(raw)-4]
	if chk := encode(out); in != chk {
		err = fmt.Errorf("Invalid checksum, expected %s got %s", chk, in)
	}
	return
}

// hash calculate the blacke2b 32bit hash of the input byte array
func hash(in []byte) (out []byte, err error) {
	h, err := blake2b.New(32, nil)
	if err != nil {
		return
	}
	h.Write(in)
	out = h.Sum(nil)
	return
}

// namehash calculate the namehash of a string
// TODO: link to the
func namehash(name string) []byte {
	buf := make([]byte, 32)
	for _, s := range strings.Split(name, ".") {
		sh, _ := hash([]byte(s))
		buf, _ = hash(append(buf, sh...))
	}
	return buf
}

// Sign a message with a private key
func (k *KeyPair) Sign(message []byte) (signature []byte) {
	signature = ed25519.Sign(k.SigningKey, message)
	return
}

// Verify a message with a private key
func Verify(address string, message, signature []byte) (valid bool, err error) {
	pub, err := decode(address)
	if err != nil {
		return
	}
	valid = ed25519.Verify(ed25519.PublicKey(pub), message, signature)
	return
}

func encrypt(key, data []byte) {
	// ECB mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(data)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, len(data))
	mode := aesecb.NewECBEncrypter(block)
	mode.CryptBlocks(ciphertext, data)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	fmt.Printf("%x\n", ciphertext)
}

func decrypt(key, ciphertext []byte) {
	block, err := aes.NewCipher(h(key))
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

	mode := aesecb.NewECBDecrypter(block)

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
}

// randomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}
