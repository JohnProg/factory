package main

import (
	"fmt"
	"net/http"
)

func apiHello(rw http.ResponseWriter, req *http.Request) {
	printApiHeaders(rw)
	fmt.Fprintln(rw, "oo")
}

func apiRegisterUser(rw http.ResponseWriter, req *http.Request) {
	printApiHeaders(rw)
}

func printApiHeaders(rw http.ResponseWriter) {
	rw.Header().Set("X-oo-version", VERSION)
	rw.Header().Set("X-api-status", APISTATUS)
}
