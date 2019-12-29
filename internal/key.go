package internal

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func ReadPassword() (key []byte, err error) {
	fd := int(os.Stdin.Fd())
	if !terminal.IsTerminal(fd) {
		return nil, errors.New("Cannot read passwrod from redirected stdin")
	}
	fmt.Fprint(os.Stderr, "password: ")
	key, err = terminal.ReadPassword(fd)
	fmt.Fprintln(os.Stderr)
	return
}
