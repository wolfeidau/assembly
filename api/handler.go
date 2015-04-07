package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/wolfeidau/assembly/datastore"
	"github.com/wolfeidau/assembly/router"
)

var (
	store         = datastore.NewDatastore(nil)
	schemaDecoder = schema.NewDecoder()
)

func Handler(base *mux.Router) *mux.Router {
	m := router.NewAPIRouter(base)
	m.Get(router.Users).Handler(handler(serveUsers))
	m.Get(router.User).Handler(handler(serveUser))
	return m
}

type handler func(http.ResponseWriter, *http.Request) error

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %s", err)
		log.Println(err)
	}
}
