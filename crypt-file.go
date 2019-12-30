package main

import (
	"bytes"
	"compress/zlib"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/hangilc/crypt-file/internal"
)

var keyFile = flag.String("k", "", "key file")
var decrypt = flag.Bool("d", false, "decrypt")
var outfile = flag.String("o", "", "output file")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] [INPUT]\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "If INPUT is not specified, stdin becomes INPUT\n")
		fmt.Fprintf(flag.CommandLine.Output(), "[options]\n")
		flag.PrintDefaults()
	}
}

func readInput() ([]byte, error) {
	if len(flag.Args()) == 0 {
		return ioutil.ReadAll(os.Stdin)
	} else {
		return ioutil.ReadFile(flag.Args()[0])
	}
}

func main() {
	flag.Parse()
	var err error
	var key []byte
	if *keyFile != "" {
		key, err = internal.ReadKeyFile(*keyFile)
		if err != nil {
			panic(err)
		}
	}
	if len(key) != 16 {
		switch len(key) {
		case 0:
			fmt.Fprintf(os.Stderr, "Cannot find password")
		default:
			fmt.Fprintf(os.Stderr, "Password size is not 16 (actually %d)\n", len(key))
		}
		os.Exit(1)
	}
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
	if *decrypt {
		enc, err := readInput()
		if err != nil {
			panic(err)
		}
		plain, err := internal.Decrypt(key, enc)
		if err != nil {
			panic(err)
		}
		b := bytes.NewReader(plain)
		r, err := zlib.NewReader(b)
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(output, r)
		r.Close()
		if err != nil {
			panic(err)
		}
	} else {
		plaintext, err := readInput()
		if err != nil {
			panic(err)
		}
		var buf bytes.Buffer
		w := zlib.NewWriter(&buf)
		w.Write(plaintext)
		w.Close()
		nonce := make([]byte, 12)
		_, err = rand.Read(nonce)
		if err != nil {
			panic(err)
		}
		enc, err := internal.Encrypt(internal.DefaultVersion, key, nonce, buf.Bytes())
		if err != nil {
			panic(err)
		}
		_, err = output.Write(enc)
		if err != nil {
			panic(err)
		}
	}
}
