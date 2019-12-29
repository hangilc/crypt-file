package internal

import (
	"crypto/aes"
	"crypto/cipher"
)

func Encrypt(ver int, key []byte, nonce []byte, plain []byte) ([]byte, error) {
	head, err := CreateHeader(ver, nonce)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	aead, err := cipher.NewGCM(block)
	ciphertext := aead.Seal(plain[:0], nonce, plain, nil)
	return append(head, ciphertext...), nil
}
