package api

import (
	"testing"

	"github.com/wolfeidau/assembly/assembly"
)

func TestUser(t *testing.T) {
	setup()

	wantUser := &assembly.User{ID: 1, Login: assembly.String("jdoe")}

	calledGet := false
	store.Users.(*assembly.MockUsersService).GetFunc = func(login string) (*assembly.User, assembly.Response, error) {
		if login != *wantUser.Login {
			t.Errorf("wanted request for user %s but got %s", *wantUser.Login, login)
		}
		calledGet = true
		return wantUser, nil, nil
	}

	gotUser, _, err := apiClient.Users.Get(*wantUser.Login)
	if err != nil {
		t.Fatal(err)
	}

	if !calledGet {
		t.Error("!calledGet")
	}
	if !normalizeDeepEqual(wantUser, gotUser) {
		t.Errorf("got user %+v but wanted user %+v", wantUser, gotUser)
	}
}

func TestUser_List(t *testing.T) {
	setup()

	wantUsers := []*assembly.User{{ID: 1}}
	wantOpt := &assembly.UsersListOptions{}

	calledList := false
	store.Users.(*assembly.MockUsersService).ListFunc = func(opt *assembly.UsersListOptions) ([]*assembly.User, assembly.Response, error) {
		if !normalizeDeepEqual(wantOpt, opt) {
			t.Errorf("wanted list options %+v but got %+v", wantOpt, opt)
		}
		calledList = true
		return wantUsers, nil, nil
	}

	users, _, err := apiClient.Users.List(wantOpt)
	if err != nil {
		t.Fatal(err)
	}

	if !calledList {
		t.Error("!calledList")
	}

	if !normalizeDeepEqual(&wantUsers, &users) {
		t.Errorf("got users %+v but wanted users %+v", users, wantUsers)
	}
}
