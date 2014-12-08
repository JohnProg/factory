package data

import (
	"gopkg.in/mgo.v2/bson"
	"math/big"
)

type Config struct {
	Domain   string
	BindAddr string
	Database struct {
		Hosts []string
		Name  string
	}
	UseTLS      bool
	CertFile    string
	PrivKeyFile string
}

type Factory struct {
	Domain  string
	Version string
	Apps    []App
}

type User struct {
	Realm        string
	Id           bson.ObjectId "_id"
	Name         string
	Email        string
	Password     string
	Registered   int64
	Credits      big.Int
	Inventory    map[string]Item
	Data         map[string]interface{}
	ActiveTokens []string
}

type Oo struct {
	Id        bson.ObjectId "_id"
	Name      string
	Owner     bson.ObjectId
	Inventory map[string]Item
	Data      map[string]interface{}
}

type Item struct {
	Name string
	Show bool
	Data map[string]interface{}
	Qty  int64
}

type App struct {
	Id          bson.ObjectId "_id"
	Name        string
	Secret      string
	Version     string
	Permissions []string
	Budget      big.Int
	Properties  map[string]interface{}
	Data        map[string]interface{}
	Token       string
}
