package models

// import (
// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// 	"github.com/jinzhu/gorm"
// 	"github.com/joho/godotenv"
// 	"fmt"
// )

// var db *gorm.DB //database

// func init() {

// 	e := godotenv.Load() //Load .env file
// 	if e != nil {
// 		fmt.Print(e)
// 	}

// 	username := "root"
// 	password := "password"
// 	dbName := "db"
// 	dbHost := "localhost:3307"


// 	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
// 	fmt.Println(dbUri)

// 	conn, err := gorm.Open("postgres", dbUri)
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	db = conn
// 	db.Debug().AutoMigrate() //Database migration
// }

// //returns a handle to the DB object
// func GetDB() *gorm.DB {
// 	return db
// }