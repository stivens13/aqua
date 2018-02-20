package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
)

type User struct {
	db driver
	ID                   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Username             string        `json:"username" bson:"username"`
	Password             string        `json:"-" bson:"password"`
	Email                string        `json:"email" bson:"email"`
	Roles                struct {
		                     Admin   bson.ObjectId `json:"admin" bson:"admin,omitempty"`
		                     Account bson.ObjectId `json:"account" bson:"account,omitempty"`
	                     } `json:"roles" bson:"roles"`

	IsActive             string    `json:"isActive" bson:"isActive,omitempty"`
	TimeCreated          time.Time `json:"timeCreated" bson:"timeCreated"`
	ResetPassword        struct {
		                     Token   string `json:"token" bson:"token,omitempty"`
		                     Expires time.Time `json:"expires" bson:"expires"`
	                     } `json:"resetPassword" bson:"resetPassword"`

	ResetPasswordToken   string    `json:"-" bson:"resetPasswordToken,omitempty"`
	ResetPasswordExpires time.Time `json:"resetPasswordExpires" bson:"resetPasswordExpires,omitempty"`

	Github               vendorOauth `json:"github" bson:"github"`
	Facebook             vendorOauth `json:"facebook" bson:"facebook"`
	Search               []string    `json:"search" bson:"search"`
}

type driver interface {
	find()
}

type msqld struct {

}

// MongoDB
var UserIndex mgo.Index = mgo.Index{
	Key: []string{"timeCreated", "twitter.id", "github.id", "facebook.id", "google.id", "search"},
}



func newUser(name, email, password string) {

}

func (u *User)Create(username, email, password string) {
	u.ID = bson.NewObjectId()
	u.TimeCreated = time.Now()
	u.Username = username
	u.Email = email
	u.Password = getHash(password)
	err := db.C(USERS).Insert(u)
	if err != nil {
		EXCEPTION(err)
	}
}

func (u *User) findOne(q *bson.M) (err error) {
	err = db.C(USERS).Find(q).One(u)
	if err != mgo.ErrNotFound {
		EXCEPTION(err)
	}
	return
}

func (u *User) linkAccount(a *Account) (err error) {
	err = db.C(USERS).UpdateId(u.ID, bson.M{
		"$set": bson.M{
			"roles.account": a.ID,
		},
	})
	if err != nil {
		EXCEPTION(err)
	}
	err = db.C(ACCOUNTS).UpdateId(a.ID, bson.M{
		"$set": bson.M{
			"user.id": u.ID,
			"user.name": u.Username,
		},
	})
	if err != nil {
		EXCEPTION(err)
	}
	return
}

func (u *User) unlinkAccount(a *Account) (err error) {
	err = db.C(USERS).UpdateId(u.ID, bson.M{
		"$set": bson.M{
			"roles.account": "",
		},
	})
	if err != nil {
		EXCEPTION(err)
	}
	err = db.C(ACCOUNTS).UpdateId(a.ID, bson.M{
		"$set": bson.M{
			"user.id": "",
			"user.name": "",
		},
	})
	if err != nil {
		EXCEPTION(err)
	}
	return
}

type vendorOauth struct {
	ID string `bson:"id"`
}

