package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(w, "TLS Example\n")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
	log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/")
	err := http.ListenAndServeTLS(":8443", os.Getenv("CERT_FILE"), os.Getenv("KEY_FILE"), nil)
	log.Fatal(err)
}
