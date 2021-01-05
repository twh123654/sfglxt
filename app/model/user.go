package model

import (
	// orm "sfglxt/app/database"
	"github.com/jinzhu/gorm"
	// "time"
)

type User struct {
	gorm.Model
	Code string `json: "code"`
	Name string `json: "name"`
	Password string `json: "password"`
	Description string `json: "desc"`
}

