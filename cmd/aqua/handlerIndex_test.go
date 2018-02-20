package main

import (
	"testing"
	"fmt"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
)


func TestIndex(t *testing.T) {

	// RUN
	router.GET("/", handlerIndex)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, resp.Body.Bytes(), []byte("{\"message\":\"Welcome to the plot device.\"}\n"))
	return
}

