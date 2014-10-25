package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const VERSION = "1"
const APISTATUS = "unstable"

func main() {
	router := mux.NewRouter()

	GET := router.Methods("GET").Subrouter()
	GET.HandleFunc("/", helloHTTP)

	http.Handle("/", router)

	listenAddr := ":" + os.Getenv("PORT")
	fmt.Println("Listening on " + listenAddr)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
