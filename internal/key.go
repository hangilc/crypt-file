package internal

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

// func ReadPassword() (key []byte, err error) {
// 	fd := int(os.Stdin.Fd())
// 	if !terminal.IsTerminal(fd) {
// 		return nil, errors.New("Cannot read passwrod from redirected stdin")
// 	}
// 	fmt.Fprint(os.Stderr, "password: ")
// 	key, err = terminal.ReadPassword(fd)
// 	fmt.Fprintln(os.Stderr)
// 	return
// }
