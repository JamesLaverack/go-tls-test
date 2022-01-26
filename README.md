# go-tls-test

An example containerised application, written in Go, that serves its own TLS. This is intended for testing purposes.

# Usage

Mount in a TLS certificate and key, provide the paths to these files as environment variables `CERT_FILE` and `KEY_FILE`.