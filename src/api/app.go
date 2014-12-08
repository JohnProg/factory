package api

import (
	"../data"
	"net/http"
)

func Hello(rw http.ResponseWriter, req *http.Request) {
	printHeaders(rw)
	rw.Write([]byte("oo"))
}

func AppAuth(db data.Database) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		printHeaders(rw)
		realm := req.FormValue("realm")
		secret := req.FormValue("secret")
		token, err := db.AuthApp(realm, secret)
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}
		rw.Write([]byte(token))
	}
}
