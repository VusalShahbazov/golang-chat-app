package models

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"time"
)
import "gorm.io/driver/mysql"

type Model struct {
	Id uint64 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

var DB *gorm.DB

func SetUp()  {
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?multiStatements=true&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	connection , err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil  {
		panic("cannot connect to database " + err.Error())
	}

	DB = connection
}