package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "Welcome to the plot device.",
	})
}
