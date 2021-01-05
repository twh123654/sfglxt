package model

import (
	orm "sfglxt/app/database"
	"fmt"
)

func InitTable() {
	orm.Init()
	fmt.Println(orm.Db.HasTable(&User{}))
	orm.Db.AutoMigrate(&User{})
}