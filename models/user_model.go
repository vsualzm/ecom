package models

import "time"

type User struct {
	Id            int
	Username      string
	Email         string
	Password      string
	Role          string
	CodeGenereate string
	CreatedAt     time.Time
	UpdateAt      time.Time
}
