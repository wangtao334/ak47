package client

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var (
	httpClient *http.Client
)

func initHttpClient() {
	httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Jar:     &cookiejar.Jar{},
		Timeout: 60 * time.Second,
	}
}

func closeHttpClient() {
	httpClient.CloseIdleConnections()
}

func AcquireHttpClient() *http.Client {
	return httpClient
}

func InitClient() {
	// http
	initHttpClient()
}

func CloseClient() {
	// http
	closeHttpClient()
}
