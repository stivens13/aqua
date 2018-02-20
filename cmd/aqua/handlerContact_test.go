package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/url"
)


/*
func TestContact(t *testing.T) {

	// RUN
	router.POST("/contact", handlerContact)
	req, err := http.NewRequest("POST", "/contact", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, resp.Body.Bytes(), []byte("{\"message\":\"Success.\"}\n"))

	return
}
*/


func TestContactValidation(t *testing.T) {
	router.POST("/contact", handlerContact)
	form := url.Values{}

	// RUN
	resp := getPostResponse("/contact", &form)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(
		t,
		`{"statusCode":400,"error":"Bad Request","message":"child \"name\" fails because [\"name\" is required]","validation":{"keys":["name"],"source":"payload"}}` + "\n",
		string(resp.Body.Bytes()),
	)

	form.Set("name", "name")
	resp = getPostResponse("/contact", &form)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(
		t,
		`{"statusCode":400,"error":"Bad Request","message":"child \"email\" fails because [\"email\" is required]","validation":{"keys":["email"],"source":"payload"}}` + "\n",
		string(resp.Body.Bytes()),
	)

	form.Set("email", "invalidEmail")
	resp = getPostResponse("/contact", &form)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(
		t,
		`{"statusCode":400,"error":"Bad Request","message":"child \"email\" fails because [\"email\" must be a valid email]","validation":{"keys":["email"],"source":"payload"}}` + "\n",
		string(resp.Body.Bytes()),
	)

	form.Set("email", "valid@email.com")
	resp = getPostResponse("/contact", &form)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(
		t,
		`{"statusCode":400,"error":"Bad Request","message":"child \"message\" fails because [\"message\" is required]","validation":{"keys":["message"],"source":"payload"}}` + "\n",
		string(resp.Body.Bytes()),
	)
	return
}
