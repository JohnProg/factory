package main

import (
	"./api"
	"./data"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

const VERSION = "1"
const APISTATUS = "unstable"

var factory data.Factory

func main() {
	configraw, _ := ioutil.ReadFile("config.json")
	var config data.Config
	json.Unmarshal(configraw, &config)

	factory.Domain = config.Domain
	factory.Version = VERSION

	api.SetVersion(VERSION, APISTATUS)

	// Initialize database
	database := data.InitDB(config.Database.Hosts, config.Database.Name)
	defer database.Close()

	router := mux.NewRouter()

	GET := router.Methods("GET").Subrouter()
	GET.HandleFunc("/", api.Hello)

	POST := router.Methods("POST").Subrouter()
	POST.HandleFunc("/user/register", api.UserRegister(database))
	POST.HandleFunc("/user/get", api.UserGet(database))
	POST.HandleFunc("/app/auth", api.AppAuth(database))

	http.Handle("/", router)

	fmt.Println("Listening on " + config.BindAddr)
	if config.UseTLS {
		http.ListenAndServeTLS(config.BindAddr, config.CertFile, config.PrivKeyFile, nil)
	} else {
		http.ListenAndServe(config.BindAddr, nil)
	}
}
