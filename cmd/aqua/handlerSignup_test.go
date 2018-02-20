package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/url"
	"gopkg.in/mgo.v2/bson"
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


func TestSignupValidation(t *testing.T) {
	router.POST("/signup", handlerSignup)
	form := url.Values{}

	// RUN
	resp := getPostResponse("/signup", &form)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(
		t,
		`{"statusCode":400,"error":"Bad Request","message":"child \"name\" fails because [\"name\" is required]","validation":{"keys":["name"],"source":"payload"}}` + "\n",
		string(resp.Body.Bytes()),
	)

	form.Set("name", "name")
	resp = getPostResponse("/signup", &form)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(
		t,
		`{"statusCode":400,"error":"Bad Request","message":"child \"email\" fails because [\"email\" is required]","validation":{"keys":["email"],"source":"payload"}}` + "\n",
		string(resp.Body.Bytes()),
	)

	form.Set("email", "invalidEmail")
	resp = getPostResponse("/signup", &form)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(
		t,
		`{"statusCode":400,"error":"Bad Request","message":"child \"email\" fails because [\"email\" must be a valid email]","validation":{"keys":["email"],"source":"payload"}}` + "\n",
		string(resp.Body.Bytes()),
	)

	form.Set("email", "valid@email.com")
	resp = getPostResponse("/signup", &form)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(
		t,
		`{"statusCode":400,"error":"Bad Request","message":"child \"username\" fails because [\"username\" is required]","validation":{"keys":["username"],"source":"payload"}}` + "\n",
		string(resp.Body.Bytes()),
	)

	form.Set("username", "username")
	resp = getPostResponse("/signup", &form)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(
		t,
		`{"statusCode":400,"error":"Bad Request","message":"child \"password\" fails because [\"password\" is required]","validation":{"keys":["password"],"source":"payload"}}` + "\n",
		string(resp.Body.Bytes()),
	)
	return
}

func TestSignupDuplicate(t *testing.T) {

	_id := bson.NewObjectId()
	c := db.C(USERS)
	err := c.Insert(bson.M{
		"_id": _id,
		"username": "username",
		"email": "email",
	})
	FatalTestErr(t, err)
	defer c.RemoveId(_id)

	form := url.Values{}

	form.Set("name", "name lastname")
	form.Set("username", "username")
	form.Set("email", "valid@email.com")
	form.Set("password", "password")
	// RUN
	resp := getPostResponse("/signup", &form)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(
		t,
		`{"statusCode":400,"error":"Bad Request","message":"child \"name\" fails because [\"name\" is required]","validation":{"keys":["name"],"source":"payload"}}` + "\n",
		string(resp.Body.Bytes()),
	)


	return
}

func FatalTestErr(t *testing.T, e error){
	t.Fatal(e)
}
