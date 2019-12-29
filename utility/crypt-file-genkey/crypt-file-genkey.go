package main

import (
	"crypto/rand"
	"fmt"
)

var keybytes = 16

func main() {
	key := make([]byte, keybytes)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", key)
}
