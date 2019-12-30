package lib

import (
	"encoding/hex"
	"io/ioutil"
	"strings"
)

func ReadKeyFile(path string) ([]byte, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	s := string(c)
	s = strings.TrimSpace(s)
	return hex.DecodeString(s)
}
