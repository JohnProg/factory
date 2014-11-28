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
	GET.HandleFunc("/", apiHello)

	POST := router.Methods("POST").Subrouter()
	POST.HandleFunc("/register", apiRegisterUser)

	http.Handle("/", router)

	fmt.Println("Listening on " + config.BindAddr)
	if config.UseTLS {
		http.ListenAndServeTLS(config.BindAddr, config.CertFile, config.PrivKeyFile, nil)
	} else {
		http.ListenAndServe(config.BindAddr, nil)
	}
}
