package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

func handlerSignup(c *gin.Context) {
	var body struct {
		Name    string `form:"name" json:"name" valid:"required"`
		Email   string `form:"email" json:"email" valid:"required,email,lowercase"`
		Username string `form:"username" json:"username" valid:"required,lowercase"`
		Password string `form:"password" json:"password" valid:"required"`
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
	user := User{}
	err = user.findOne(&bson.M{"$or": []bson.M{
		{"username": body.Username},
		{"email": body.Email},
	}})

	// we expect mgo.ErrNotFound
	if err == nil {
		if user.Username == body.Username {
			c.JSON(http.StatusOK, gin.H{
				"message" : "Username already in use.",
			})
		} else if user.Email == body.Email {
			c.JSON(http.StatusOK, gin.H{
				"message" : "Email already in use.",
			})
		}
		return
	}

	user.Create(body.Username, body.Email, body.Password)
	account := Account{}
	account.Create(body.Name)
	user.linkAccount(&account)





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

