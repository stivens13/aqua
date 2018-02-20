package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func handlerContact(c *gin.Context) {
	var body struct {
		Name    string `form:"name" json:"name" valid:"required"`
		Email   string `form:"email" json:"email" valid:"email,required"`
		Message string `form:"message" json:"message" valid:"required"`
	}
	err := c.Bind(&body)
	if err != nil {
		EXCEPTION(err)
	}
	valid, resErr := validate(body)
	if !valid {
		c.JSON(http.StatusBadRequest, resErr)
		return
	}
	mailConf := MailConfig{}
	mailConf.Data = body
	mailConf.From = config.SMTP.From.Name + " <" + config.SMTP.From.Address + ">"
	//mailConf.To = config.SystemEmail
	mailConf.To = "im7mortal@gmail.com"
	mailConf.Subject = config.CompanyName + " contact form"
	//mailConf.ReplyTo = body.Email
	mailConf.ReplyTo = "im7mortal@gmail.com"
	mailConf.HtmlPath = "templates/email-html.html"

	if err := mailConf.SendMail(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message" : "Email wasn't send. Please try another time or later.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Success.",
	})
}

func EXCEPTION(i interface{}) {
	log.Panicln(i)
}
