package models

import (
	"fmt"
	"os"

	//this is the postgress go orm
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

//the database
var db *gorm.DB

func init() {

	//load the .env file
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	//build connection string
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	//db migrations
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

//returns a handle to the DB Object

func GetDB() *gorm.DB {
	return db
}
