package main

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/kingpin"
	"github.com/gorilla/mux"
	"github.com/juju/loggo"
	"github.com/wolfeidau/assembly/api"
	"github.com/wolfeidau/assembly/assembly"
	"github.com/wolfeidau/assembly/datastore"
)

var (
	debug        = kingpin.Flag("debug", "Enable debug mode.").OverrideDefaultFromEnvar("DEBUG").Bool()
	createTables = kingpin.Flag("createTables", "Create tables on startup.").OverrideDefaultFromEnvar("CREATE_TABLES").Bool()
	port         = kingpin.Flag("port", "Port to bind to for the HTTP service.").OverrideDefaultFromEnvar("PORT").Int()

	log = loggo.GetLogger("assembly-server")
)

func init() {

}

func main() {
	kingpin.Version(assembly.Version)
	kingpin.Parse()

	// apply flags
	if *debug {
		loggo.GetLogger("").SetLogLevel(loggo.DEBUG)
	}

	log.Infof("connecting to db")
	datastore.Connect()

	if *createTables {
		datastore.Create()
	}

	mux := mux.NewRouter()

	api.Handler(mux)

	addr := fmt.Sprintf(":%d", *port)

	log.Infof("Listening on %d", *port)

	err := http.ListenAndServe(addr, mux)

	if err != nil {
		log.Errorf("Failed to start listener: %s", err)
	}

}
