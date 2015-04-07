package assembly

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUserService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/jdoe", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":123,"login":"jdoe","email":"jane@doe.com"}`)
	})

	user, _, err := client.Users.Get("jdoe")
	if err != nil {
		t.Errorf("User.Get returned error: %v", err)
	}

	want := &User{ID: 123, Login: String("jdoe"), Email: String("jane@doe.com")}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("User.Get returned %+v, want %+v", user, want)
	}
}
