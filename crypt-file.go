package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/hangilc/crypt-file/internal"
)

var decrypt = flag.Bool("d", false, "decrypt")
var outfile = flag.String("o", "", "output file")

func makeVersionHeader(ver int) []byte {
	buf := make([]byte, 3)
	buf[0] = 'C'
	buf[1] = 'F'
	buf[2] = byte(ver)
	return buf
}

func readVersion(buf []byte) (ver int, rem []byte, err error) {
	if len(buf) < 3 {
		err = errors.New("Cannot read version header")
		return
	}
	if !(buf[0] == 'C' && buf[1] == 'F') {
		err = errors.New("It is not crypt-file encoded file")
		return
	}
	return int(buf[2]), buf[3:], nil
}

func main() {
	flag.Parse()
	var err error
	key := []byte("plaintextpassword")[:16]
	if *decrypt {
		enc, err := ioutil.ReadFile(flag.Args()[0])
		if err != nil {
			panic(err)
		}
		plain, err := internal.Decrypt(key, enc)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", plain)
	} else {
		plaintext := []byte("hello, world")
		nonce := make([]byte, 12)
		_, err = rand.Read(nonce)
		if err != nil {
			panic(err)
		}
		enc, err := internal.Encrypt(internal.DefaultVersion, key, nonce, plaintext)
		var output io.Writer
		if *outfile == "" {
			output = os.Stdout
		} else {
			f, err := os.Create(*outfile)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			output = f
		}
		_, err = output.Write(enc)
		if err != nil {
			panic(err)
		}
	}
}
