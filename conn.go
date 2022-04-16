package connection

import (
	"fmt"

	"github.com/FantasyApi/interfaces/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conn() *gorm.db {
	dbURL := "postgres://pg:nashkim@localhost:5432/fantasy"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		fmt.Println("error occured", err)
	}
	db.AutoMigrate(&users.Users{})

	return db
}
