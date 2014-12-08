package api

import (
	"../data"
	"encoding/json"
	"net/http"
	"strings"
)

func UserRegister(db data.Database) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		printHeaders(rw)
		realm := req.FormValue("realm")
		token := req.FormValue("token")
		user := req.FormValue("user")
		pass := req.FormValue("pass")
		email := req.FormValue("email")

		err := db.GetApp(realm, token)
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}

		token, err = db.CreateUser(realm, user, pass, email)
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}

		rw.Write([]byte(token))
	}
}

func UserGet(db data.Database) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		printHeaders(rw)
		name := req.FormValue("user")

		realm := req.FormValue("realm")
		token := req.FormValue("token")
		err := db.GetApp(realm, token)
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}

		user, err := db.GetUser(name)
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}

		/* Strip things the app shouldn't see */
		props := make(map[string]interface{})
		for prop := range user.Data {
			if strings.HasPrefix(prop, realm+".") {
				props[prop[len(realm)+2:]] = user.Data[prop]
			}
		}

		strval, err := json.Marshal(data.User{
			Name:       user.Name,
			Email:      user.Email,
			Registered: user.Registered,
			Credits:    user.Credits,
			Inventory:  user.Inventory,
			Data:       props,
		})

		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}

		rw.Write(strval)
	}
}
