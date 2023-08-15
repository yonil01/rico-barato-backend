package ws

import (
	"crypto/tls"
	"net/http"
	"sync"
)

var (
	once   sync.Once
	client *http.Client
)

func GetClientConnection() *http.Client {
	once.Do(func() {

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	})

	return client
}
