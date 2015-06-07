package datastore

import (
	"github.com/jmoiron/modl"
	"github.com/wolfeidau/assembly/assembly"
)

func init() {
	DB.AddTableWithName(assembly.User{}, "user").SetKeys(true, "ID")
}

type usersStore struct{ *Datastore }

func (s *usersStore) Get(login string) (*assembly.User, assembly.Response, error) {
	var users []*assembly.User
	if err := s.dbh.Select(&users, `SELECT * FROM public.user WHERE login=$1;`, login); err != nil {
		return nil, nil, err
	}
	if len(users) == 0 {
		return nil, nil, assembly.ErrUserNotFound
	}
	return users[0], nil, nil
}

func (s *usersStore) Create(user *assembly.User) (*assembly.User, assembly.Response, error) {

	err := transact(s.dbh, func(tx modl.SqlExecutor) error {
		var existing []*assembly.User

		if err := s.dbh.Select(&existing, `SELECT * FROM public.user WHERE login=$1;`, user.Login); err != nil {
			return err
		}

		if len(existing) > 0 {
			return assembly.ErrUserLoginExists
		}

		if err := tx.Insert(user); err != nil {
			return err
		}

		return nil
	})

	return user, nil, err
}

func (s *usersStore) Update(user *assembly.User) (*assembly.User, assembly.Response, error) {

	err := transact(s.dbh, func(tx modl.SqlExecutor) error {

		n, err := tx.Update(user)

		if err != nil {
			return err
		}

		if n != 1 {
			return assembly.ErrUserNotFound
		}

		return nil
	})

	return user, nil, err
}

func (s *usersStore) List(opt *assembly.UsersListOptions) ([]*assembly.User, assembly.Response, error) {
	if opt == nil {
		opt = &assembly.UsersListOptions{}
	}

	sql := `SELECT * FROM public.user`

	var users []*assembly.User

	return users, nil, s.dbh.Select(&users, sql)
}
