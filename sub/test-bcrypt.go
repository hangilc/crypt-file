package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	key := "hello, world"
	for i := bcrypt.DefaultCost; i <= bcrypt.MaxCost; i++ {
		start := time.Now()
		hash, _ := bcrypt.GenerateFromPassword([]byte(key), i)
		fmt.Printf("%x\n", hash)
		fmt.Printf("(%d) %f seconds\n", i, time.Since(start).Seconds())
	}
}
