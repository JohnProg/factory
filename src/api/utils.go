package api

import (
	"net/http"
)

var VERSION, APISTATUS string

func SetVersion(version, status string) {
	VERSION = version
	APISTATUS = status
}

func printHeaders(rw http.ResponseWriter) {
	rw.Header().Set("X-oo-version", VERSION)
	rw.Header().Set("X-api-status", APISTATUS)
}
