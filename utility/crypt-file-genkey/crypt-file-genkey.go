package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

var keybytes = 16
var outfile = flag.String("o", "", "output file")

func main() {
	flag.Parse()
	key := make([]byte, keybytes)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	rep := hex.EncodeToString(key)
	if *outfile != "" {
		f, err := os.OpenFile(*outfile, os.O_CREATE|os.O_EXCL, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = f.WriteString(rep)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("%s\n", rep)
	}
}
