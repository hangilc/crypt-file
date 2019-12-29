package internal

import "fmt"

var DefaultVersion = 1

func CreateHeader(ver int, nonce []byte) ([]byte, error) {
	switch ver {
	case 1:
		return createHeaderVer1(ver, nonce), nil
	default:
		return nil, fmt.Errorf("Invalid version number: %d", ver)
	}
}

func ExtractHeader(encrypted []byte) (nonce []byte, rest []byte, err error) {
	if len(encrypted) < 3 {
		err = fmt.Errorf("Too small data to extract header")
		return
	}
	if !(encrypted[0] == 'C' && encrypted[1] == 'F') {
		err = fmt.Errorf("Not crypt-file data")
		return
	}
	ver := int(encrypted[2])
	switch ver {
	case 1:
		return extractHeaderVer1(encrypted)
	default:
		err = fmt.Errorf("Invalid version ")
		return
	}
}

func createHeaderVer1(ver int, nonce []byte) []byte {
	head := make([]byte, 3)
	head[0] = 'C'
	head[1] = 'F'
	head[2] = byte(ver)
	if len(nonce) != 12 {
		panic(fmt.Sprintf("size of nonce is not 12"))
	}
	return append(head, nonce...)
}

func extractHeaderVer1(encrypted []byte) (nonce []byte, rest []byte, err error) {
	enc := encrypted[3:]
	nonce = enc[:12]
	rest = enc[12:]
	return
}
