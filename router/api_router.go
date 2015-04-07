package router

import "github.com/gorilla/mux"

var (
	Users      = "users"
	User       = "user"
	UserEmails = "user.emails"
)

// NewAPIRouter creates a new API router with route URL pattern definitions but
// no handlers attached to the routes.
func NewAPIRouter(base *mux.Router) *mux.Router {

	if base == nil {
		base = mux.NewRouter()
	}

	base.StrictSlash(true)

	base.Path("/users").Methods("GET").Name(Users)

	userPath := "/users/{Login:.*}"

	base.Path(userPath).Methods("GET").Name(User)

	user := base.PathPrefix(userPath).Subrouter()
	user.Path("/emails").Methods("GET").Name(UserEmails)

	return base
}
