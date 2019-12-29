package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/hangilc/crypt-file/internal"
)

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
	key, err := internal.ReadPassword()
	if err != nil {
		panic(nil)
	}
	fmt.Printf("%s\n", string(key))
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
	ver := 1
	_, err = output.Write(makeVersionHeader(ver))
	if err != nil {
		panic(err)
	}
}
