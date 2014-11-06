package main

import (
	"fmt"
	"net/http"
)

func helloHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("X-oo-version", VERSION)
	rw.Header().Set("X-api-status", APISTATUS)
	fmt.Fprintf(rw, "oo")
}
