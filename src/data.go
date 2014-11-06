package main

import (
	"gopkg.in/mgo.v2/bson"
	"math/big"
)

type Config struct {
	Domain      string
	BindAddr    string
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
	Id         bson.ObjectId "_id"
	Name       string
	Email      string
	Password   string
	Registered int64
	Credits    big.Int
	Inventory  map[string]int
	Data       map[string]interface{}
}

type Oo struct {
	Id        bson.ObjectId "_id"
	Name      string
	Owner     bson.ObjectId
	Inventory []Item
	Data      map[string]interface{}
}

type Item struct {
	AppOwner bson.ObjectId
	Name     string
	Qty      int64
}

type App struct {
	Id          bson.ObjectId "_id"
	Name        string
	Secret      string
	Version     string
	Permissions []string
}
