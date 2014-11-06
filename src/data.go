package main

import (
	"gopkg.in/mgo.v2/bson"
	"math/big"
)

type Config struct {
	Domain   string
	BindAddr string
}

type Factory struct {
	Domain  string
	Version string
}

type User struct {
	Id         bson.ObjectId "_id"
	Name       string
	Email      string
	Password   string
	Registered int64
	Credits    big.Int
	Inventory  map[string]int
	Properties map[string]interface{}
}

type Oo struct {
	Id         bson.ObjectId "_id"
	Name       string
	Owner      bson.ObjectId
	Inventory  map[string]int
	Properties map[string]interface{}
}
