package lib

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"

	"github.com/hangilc/crypt-file/internal"
)

func Encrypt(key []byte, plain []byte) ([]byte, error) {
	nonce, err := internal.CreateNonce()
	if err != nil {
		return nil, err
	}
	head, err := internal.CreateHeader(internal.DefaultVersion, nonce)
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

func CompressAndEncrypt(key []byte, plain []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	w.Write(plain)
	w.Close()
	return Encrypt(key, buf.Bytes())
}
