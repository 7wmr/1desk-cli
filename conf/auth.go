package conf

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	b64 "encoding/base64"
	"encoding/json"
)

// Auth contains username and password (decoded)
type Auth struct {
	Username string
	Password string
	Encoded  string
}

// Token returned from API request (valid approx 5 mins idle time).
type Token struct {
	Access string `json:"access_token"`
	Type   string `json:"token_type"`
}

// GetHeader return the header string for the token.
func (t *Token) GetHeader() (string, string) {
	return "Authorization", (t.Type + " " + t.Access)
}

// Encode will encode Username and Password to base64 string.
func (a *Auth) Encode() string {
	a.Encoded = b64.StdEncoding.EncodeToString(
		[]byte((a.Username + ":" + a.Password)))
	return a.Encoded
}

// Decode will decode Encoded to Username and Password strings.
func (a *Auth) Decode() {
	data, err := b64.StdEncoding.DecodeString(a.Encoded)
	if err != nil {
		fmt.Println("Error: Issue when decoding context string")
		return
	}
	creds := strings.Split(string(data), ":")
	a.Username = creds[0]
	a.Password = creds[1]
}

// GetToken will return the token provided for the specified user.
func (a *Auth) GetToken(context *Context) (Token, error) {
	var token Token

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// Decode username and password.
	a.Decode()

	var url = fmt.Sprintf("https://%s/api/auth-service/token/", context.Domain)
	var post = fmt.Sprintf("{ \"username\": \"%s\", \"password\": \"%s\", \"grant_type\": \"password\" }", a.Username, a.Password)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(post)))
	if err != nil {
		return token, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return token, err
	}

	if resp.StatusCode != 200 {
		return token, errors.New("Error performing API call: " + string(resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &token)

	return token, nil
}
