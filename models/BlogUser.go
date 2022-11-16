package models

import "time"

type BlogUser struct {
	Id                  int
	FirstName           string
	LastName            string
	Username            string
	Email               string
	Password            string
	AccountLocked       bool
	AccountDisabled     bool
	TimeAccountLockedAt time.Time
}

type BlogUserExternal struct {
	Id              int
	FirstName       string
	LastName        string
	Email           string
	AccountLocked   bool
	AccountDisabled bool
}
