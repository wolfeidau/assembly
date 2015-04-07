package datastore

import "github.com/wolfeidau/assembly/assembly"

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

func (s *usersStore) List(opt *assembly.UsersListOptions) ([]*assembly.User, assembly.Response, error) {
	if opt == nil {
		opt = &assembly.UsersListOptions{}
	}

	sql := `SELECT * FROM public.user`

	var users []*assembly.User

	return users, nil, s.dbh.Select(&users, sql)
}
