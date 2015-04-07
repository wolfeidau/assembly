package assembly

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// User represents an assembly user.
type User struct {
	ID        int        `json:"id,omitempty"`
	Login     *string    `json:"login,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Email     *string    `json:"email,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
}

// UsersService communicates with the users-related endpoints in the
// assembly API.
type UsersService interface {
	// Get fetches a user.
	Get(login string) (*User, Response, error)

	// List fetches all user.
	List(opt *UsersListOptions) ([]*User, Response, error)
}

// UsersListOptions specifies the optional parameters to the
// UsersService.List method.
type UsersListOptions struct {
	Sort      string `url:",omitempty"`
	Direction string `url:",omitempty"`
}

// EmailAddr is an email address associated with a user.
type EmailAddr struct {
	Email string
}

// usersService implements UsersService.
type usersService struct {
	client *Client
}

var _ UsersService = &usersService{}

func (s *usersService) Get(login string) (*User, Response, error) {

	u := fmt.Sprintf("users/%s", login)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var user *User
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (s *usersService) List(opt *UsersListOptions) ([]*User, Response, error) {

	u := fmt.Sprintf("users")

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := s.client.Do(req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

type MockUsersService struct {
	GetFunc  func(login string) (*User, Response, error)
	ListFunc func(opt *UsersListOptions) ([]*User, Response, error)
}

var _ UsersService = &MockUsersService{}

func (s *MockUsersService) Get(login string) (*User, Response, error) {
	return s.GetFunc(login)
}

func (s *MockUsersService) List(opt *UsersListOptions) ([]*User, Response, error) {
	return s.ListFunc(opt)
}
