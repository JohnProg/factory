package data

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type Database struct {
	session *mgo.Session
	db      *mgo.Database
}

func InitDB(servers []string, dbname string) Database {
	var database Database
	var err error
	database.session, err = mgo.Dial(strings.Join(servers, ","))
	if err != nil {
		panic(err)
	}
	database.db = database.session.DB(dbname)
	return database
}

func (d Database) Close() {
	d.session.Close()
}

func (d Database) AuthApp(realm, secret string) (string, error) {
	token := GenToken()
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{"token": token}},
	}
	_, err := d.db.C("app").Find(bson.M{"Name": realm}).Apply(change, nil)
	return token, err
}

func (d Database) GetApp(realm, token string) error {
	var app App
	err := d.db.C("app").Find(bson.M{"Name": realm}).One(&app)
	if err != nil {
		return err
	}

	if app.Token != token {
		return errors.New("Invalid token")
	}

	return nil
}

func (d Database) CreateUser(realm, name, pass, email string) (string, error) {
	token := GenToken()
	user := User{
		Realm:        realm,
		Id:           bson.NewObjectId(),
		Name:         name,
		Email:        email,
		Password:     EncryptPassword(pass),
		Registered:   time.Now().UTC().Unix(),
		ActiveTokens: []string{token},
	}
	err := d.db.C("users").Insert(user)
	return token, err
}

func (d Database) GetUser(name string) (User, error) {
	var user User
	err := d.db.C("users").Find(bson.M{"Name": name}).One(&user)
	return user, err
}
