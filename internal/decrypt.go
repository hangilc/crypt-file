package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Decrypt(key []byte, encrypted []byte) ([]byte, error) {
	nonce, rest, err := ExtractHeader(encrypted)
	fmt.Printf("%v\n", nonce)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return aead.Open(rest[:0], nonce, rest, nil)
}
