package main

import "gopkg.in/mgo.v2"


const USERS = "users"
const LOGINATTEMPTS = "loginattempts"
const ACCOUNTS = "accounts"
const ADMINGROUPS = "admingroups"
const CATEGORIES = "categories"
const STATUSES = "status"
const ADMINS = "admins"

var db *mgo.Database

func init() {
	session, err := mgo.Dial(config.MongoDB)
	if err != nil {
		EXCEPTION(err)
	}
	db = session.DB(config.dbName)

	// create collections if not exist and ensure indexes
	c := db.C(USERS)
	//c.EnsureIndex(UserUniqueIndex)
	c.EnsureIndex(UserIndex)

	c = db.C(LOGINATTEMPTS)
	//c.EnsureIndex(LoginAttemptsIndex)

	c = db.C(ACCOUNTS)
	//c.EnsureIndex(AccountIndex)

	c = db.C(ADMINGROUPS)
	//c.EnsureIndex(AdminGroupIndex)

	c = db.C(CATEGORIES)
	//c.EnsureIndex(CategoryIndex)

	c = db.C(STATUSES)
	c.EnsureIndex(StatusesIndex)

	c = db.C(ADMINS)
	//c.EnsureIndex(AdminsIndex)

}
