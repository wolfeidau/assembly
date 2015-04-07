package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wolfeidau/assembly/assembly"
)

func serveUser(w http.ResponseWriter, r *http.Request) error {
	login := mux.Vars(r)["Login"]

	user, _, err := store.Users.Get(login)
	if err != nil {
		return err
	}

	return writeJSON(w, user)
}

func serveUsers(w http.ResponseWriter, r *http.Request) error {
	var opt assembly.UsersListOptions
	if err := schemaDecoder.Decode(&opt, r.URL.Query()); err != nil {
		return err
	}

	users, _, err := store.Users.List(&opt)
	if err != nil {
		return err
	}
	if users == nil {
		users = []*assembly.User{}
	}

	return writeJSON(w, users)
}
