package db

import (
	"github.com/zufaralam02/first-go/model"
)

func Migrate() {
	db := InitDb()

	db.AutoMigrate(&model.User{})
}
