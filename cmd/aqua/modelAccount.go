package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
	"strings"
)

type Account struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	User         struct {
		             ID   bson.ObjectId `bson:"id,omitempty" json:"id"`
		             Name string        `bson:"name" json:"name"`
	             } `bson:"user" json:"user"`
	Name         struct {
		             First  string `bson:"first" json:"first"`
		             Middle string `bson:"middle" json:"middle"`
		             Last   string `bson:"last" json:"last"`
		             Full   string `bson:"full" json:"full"`
	             } `bson:"name" json:"name"`
	Status       struct {
		             ID  Status `bson:"current,omitempty" json:"current"`
		             Log []Status        `bson:"log" json:"log"`
	             } `bson:"status" json:"status"`
	Notes        []Note          `bson:"notes" json:"notes"`
	Verification struct {
		             Complete Status `bson:"complete" json:"complete"`
		             Token    []Status `bson:"token" json:"token"`
	             } `bson:"verification" json:"verification"`
	TimeCreated  time.Time `json:"timeCreated" bson:"timeCreated"`
}

func (a *Account)Create(name string) {
	a.ID = bson.NewObjectId()
	a.TimeCreated = time.Now()
	a.Name.Full = name
	first, middle, last := parseName(name)
	a.Name.First = first
	a.Name.Middle = middle
	a.Name.Last = last
	err := db.C(ACCOUNTS).Insert(a)
	if err != nil {
		EXCEPTION(err)
	}
}

func parseName(name string) (first, middle, last string) {
	names := strings.Split(name, " ")
	switch len(names) {
	case 1:
		first = name
	case 2:
		first = names[0]
		last = names[1]
	case 3:
		first = names[0]
		middle = names[1]
		last = names[2]
	default:
		first = name
	}
	return
}


func (a *Account) findOne(q *bson.M) (err error) {
	err = db.C(ACCOUNTS).Find(q).One(a)
	if err != mgo.ErrNotFound {
		EXCEPTION(err)
	}
	return
}
