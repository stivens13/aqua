package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func TestR500(t *testing.T) {

	router.Use(checkRecover)
	router.GET("/test500", func(c *gin.Context) {
		panic("test500")
	})
	req, err := http.NewRequest("GET", "/test500", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 500, resp.Code)
	assert.Equal(
		t,
		`{"error":"Internal Server Error","message":"An internal server error occurred","statusCode":500}` + "\n",
		string(resp.Body.Bytes()),
	)
}

func TestStatus404(t *testing.T) {
	router.NoRoute(notFound)
	req, err := http.NewRequest("GET", "/test400", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 404, resp.Code)
	assert.Equal(
		t,
		`{"error":"Not Found","statusCode":404}` + "\n",
		string(resp.Body.Bytes()),
	)
}
