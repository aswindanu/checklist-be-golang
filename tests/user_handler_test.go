package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-starterrouter"

	"github.com/gavv/httpexpect"
)

func TestGetUsers(t *testing.T) {
	// create http.Handler
	handler := router.New()

	// run server using httptest
	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	// without query string
	obj := e.GET("/users").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Keys().Contains("data")
	obj.Value("data").Array().Element(0).Object()

	// with rp & p set
	obj = e.GET("/users").WithQuery("rp", "2").WithQuery("p", "2").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Keys().Contains("data")
	obj.Value("data").Array().Element(0).Object().Value("id").Equal(3)

	// with filter name
	obj = e.GET("/users").WithQuery("name", "iman").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Keys().Contains("data")
	obj.Value("data").Array().Element(0).Object().Value("name").Equal("iman")
}

func TestAddUser(t *testing.T) {
	// create http.Handler
	handler := router.New()

	// run server using httptest
	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	payload := make(map[string]interface{})

	// normal add new
	payload = map[string]interface{}{
		"name":  "New User",
		"email": "new@email.com",
	}
	obj := e.POST("/users").
		WithJSON(payload).
		Expect().
		Status(http.StatusCreated).JSON().Object()
	obj.ContainsKey("name").ValueEqual("name", "New User")
	obj.ContainsKey("email").ValueEqual("email", "new@email.com")

	// failed with 422
	payload = map[string]interface{}{
		"email": "new@email.com",
	}
	obj = e.POST("/users").
		WithJSON(payload).
		Expect().
		Status(http.StatusUnprocessableEntity).JSON().Object()
	obj.Value("validation_errors").Object().Value("name").Array().Element(0).String().Equal("The name field is required")
}
