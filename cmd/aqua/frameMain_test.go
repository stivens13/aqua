package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"os"
	"flag"
	"net/url"
	"net/http/httptest"
	"bytes"
	"fmt"
	"net/http"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	flag.Parse()

	// SETUP
	gin.SetMode(gin.TestMode)
	router = gin.Default()

	os.Exit(m.Run())
}

func getPostResponse(path string, body *url.Values) *httptest.ResponseRecorder {
	req, err := http.NewRequest("POST", path, bytes.NewBufferString(body.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	return resp
}
