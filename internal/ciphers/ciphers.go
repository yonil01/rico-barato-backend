package ciphers

import (
	"fmt"
	openssl "github.com/Luzifer/go-openssl/v4"
)

var secretKey string

func init() {
	secretKey = "204812730425442A472D2F423F452847"
}

func Encrypt(strToEncrypt string) string {

	o := openssl.New()

	enc, err := o.EncryptBytes(secretKey, []byte(strToEncrypt), openssl.BytesToKeyMD5)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		return ""
	}

	return string(enc)
}

func Decrypt(strToDecrypt string) string {
	o := openssl.New()
	dec, err := o.DecryptBytes(secretKey, []byte(strToDecrypt), openssl.BytesToKeyMD5)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		return ""
	}
	return string(dec)
}

func GetSecret() string {
	return secretKey
}
