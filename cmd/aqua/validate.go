package main

import (
	"github.com/asaskevich/govalidator"
	"strings"
)

type resp struct {
	StatusCode int `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
	Validation map[string]interface{} `json:"validation"`
}

func validate(o interface{}) (bool, *resp) {

	valid, err := govalidator.ValidateStruct(o)
	if valid {
		return true, nil
	}
	if err != nil {
		if errors, ok := err.(govalidator.Errors); ok {
			res := resp{}
			res.StatusCode = 400
			res.Error = "Bad Request"
			validDescription := getValidationMessage(errors[0].Error())
			res.Message = validDescription.Message
			res.Validation = map[string]interface{}{
				"source": "payload",
				"keys": []string{
					validDescription.Tag,
				},
			}
			return false, &res
		}
		EXCEPTION(err)
	}
	return false, nil
}

/*

{
  "statusCode": 400,
  "error": "Bad Request",
  "message": "child \"email\" fails because [\"email\" is required]",
  "validation": {
    "source": "payload",
    "keys": [
      "email"
    ]
  }
}

*/


// mapping
// don't like it
// rough place


func getValidationMessage(e string) validMessage {
	if v, ok := validateMessages[e]; ok {
		return v
	}
	switch {
	case strings.Contains(e, "does not validate as email"):
		return validateMessages["Email: ... does not validate as email"]
	}
	return validMessage{
		Message: "weird",
		Tag: "weird",
	}
}

type validMessage struct {
	Message string
	Tag     string
}

var validateMessages map[string]validMessage = map[string]validMessage{
	"Name: non zero value required": {
		Message: `child "name" fails because ["name" is required]`,
		Tag:     "name",
	},
	"Email: non zero value required": {
		Message: `child "email" fails because ["email" is required]`,
		Tag:     "email",
	},
	// this has a special case
	"Email: ... does not validate as email": {
		Message: `child "email" fails because ["email" must be a valid email]`,
		Tag:     "email",
	},
	"Message: non zero value required": {
		Message: `child "message" fails because ["message" is required]`,
		Tag:     "message",
	},
	"Username: non zero value required": {
		Message: `child "username" fails because ["username" is required]`,
		Tag:     "username",
	},
	"Password: non zero value required": {
		Message: `child "password" fails because ["password" is required]`,
		Tag:     "password",
	},
}

