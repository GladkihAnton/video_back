package models

import (
	"fmt"
)

type User struct {
	Id       int64
	Email    string `pg:",unique"`
	Password string

	tableName struct{} `pg:"user"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Email)
}
