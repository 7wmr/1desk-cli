package conf

import (
	"fmt"
	"strings"

	b64 "encoding/base64"
)

// Auth contains username and password (decoded)
type Auth struct {
	Username string
	Password string
	Encoded  string
}

// Encode will encode Username and Password to base64 string.
func (a *Auth) Encode() string {
	a.Encoded = b64.StdEncoding.EncodeToString(
		[]byte((a.Username + ":" + a.Password)))
	return a.Encoded
}

// Decode will decode Encoded to Username and Password strings.
func (a *Auth) Decode() (string, string, error) {
	data, err := b64.StdEncoding.DecodeString(a.Encoded)
	if err != nil {
		fmt.Println("Error: Issue when decoding context string")
		return "", "", err
	}
	creds := strings.Split(string(data), ":")
	a.Username = creds[0]
	a.Password = creds[1]
	return a.Username, a.Password, nil
}
