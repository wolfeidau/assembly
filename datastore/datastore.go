package datastore

import (
	"github.com/jmoiron/modl"
	"github.com/wolfeidau/assembly/assembly"
)

// A Datastore accesses the datastore (in PostgreSQL).
type Datastore struct {
	Users assembly.UsersService

	dbh modl.SqlExecutor
}

// NewDatastore creates a new client for accessing the datastore (in
// PostgreSQL). If dbh is nil, it uses the global DB handle.
func NewDatastore(dbh modl.SqlExecutor) *Datastore {
	if dbh == nil {
		dbh = DBH
	}

	d := &Datastore{dbh: dbh}
	d.Users = &usersStore{d}
	return d
}

func NewMockDatastore() *Datastore {
	return &Datastore{
		Users: &assembly.MockUsersService{},
	}
}
