// +build go1.8,dev

package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

// CipherSuites without known attacks or extreme CPU usage
// https://golang.org/src/crypto/tls/cipher_suites.go#L75
var CipherSuites = []uint16{
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
	tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
}

// Curves without known attacks or extreme CPU usage
// https://golang.org/src/crypto/tls/common.go#L542
var Curves = []tls.CurveID{
	tls.CurveP256,
	tls.X25519,
}

// TLSConfig for including autocert manager
func tlsConfig(domain string) *tls.Config {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		panic(err)
	}

	return &tls.Config{
		MinVersion:               tls.VersionTLS12,
		Certificates:             []tls.Certificate{cert},
		PreferServerCipherSuites: true,
		CurvePreferences:         Curves,
		CipherSuites:             CipherSuites,
	}
}

func getHTTPServer(domain, addr string) *http.Server {
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,

		Addr:      addr,
		TLSConfig: tlsConfig(domain),
	}
}
