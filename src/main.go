package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

const VERSION = "1"
const APISTATUS = "unstable"

var factory Factory

func main() {
	configraw, _ := ioutil.ReadFile("config.json")
	var config Config
	json.Unmarshal(configraw, &config)

	factory.Domain = config.Domain
	factory.Version = VERSION

	router := mux.NewRouter()

	GET := router.Methods("GET").Subrouter()
	GET.HandleFunc("/", helloHTTP)

	http.Handle("/", router)

	fmt.Println("Listening on " + config.BindAddr)
	http.ListenAndServe(config.BindAddr, nil)
}
