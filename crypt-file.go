package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var infile string
var outfile string
var decrypt bool
var passwordEnv = "ENCRYPT_DUMP_PWD"

func init() {
	flag.StringVar(&infile, "i", "", "input file")
	flag.StringVar(&outfile, "o", "", "output file")
	flag.BoolVar(&decrypt, "d", false, "decrypt instead of encryp")
}

func main() {
	flag.Parse()
	password := os.Getenv(passwordEnv)
	if password == "" {
		fmt.Fprintf(os.Stderr, "Cannot read password from: $%s\n", passwordEnv)
		os.Exit(1)
	}
	key := sha256.Sum256([]byte(password))

	var input []byte
	var err error
	if infile == "" {
		input, err = ioutil.ReadAll(os.Stdin)
	} else {
		input, err = ioutil.ReadFile(infile)
		if err != nil {
			log.Fatal(err)
		}
	}

	block, err := aes.NewCipher(key[:])
	if err != nil {
		log.Fatal(err)
	}

	var output []byte
	if decrypt {
		iv := input[:aes.BlockSize]
		output = make([]byte, len(input)-aes.BlockSize)
		ds := cipher.NewCFBDecrypter(block, iv)
		ds.XORKeyStream(output, input[aes.BlockSize:])
	} else {
		ciphertext := make([]byte, aes.BlockSize+len(input))
		iv := ciphertext[:aes.BlockSize]
		_, err = io.ReadFull(rand.Reader, iv)
		if err != nil {
			log.Fatal(err)
		}
		stream := cipher.NewCFBEncrypter(block, iv)
		stream.XORKeyStream(ciphertext[aes.BlockSize:], input)
		output = ciphertext
	}

	if outfile == "" {
		_, err = os.Stdout.Write(output)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = ioutil.WriteFile(outfile, output, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}
