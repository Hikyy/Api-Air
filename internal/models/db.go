package models

import (
	"App/internal/database"
	"App/internal/modules/hash"
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseProvider struct {
	EntityDB
}

var Db *DatabaseProvider

var InitGorm *DbGorm

func InitDB() (*DbGorm, error) {
	conn, err := gorm.Open(postgres.Open(database.BuildConnectionString()), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	dbase, err := conn.DB()

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database")

	// conn.LogMode(true)

	return &DbGorm{
		db:    conn,
		dbase: dbase,
	}, nil
}

func GetDB() *DatabaseProvider {
	return Db
}

func (ug *DbGorm) Begin() *gorm.DB {
	return ug.db.Begin()
}

func DatabaseServiceProvider() error {
	ug, err := InitDB()

	if err != nil {
		log.Fatal("Erreur lors de la récupération de l'objet DB:", err)
		return err
	}

	// defer sqlDB.Close()

	InitGorm = ug

	// ug.db.AutoMigrate(&User{})

	ug.Ping()

	hmac := hash.NewHMAC(HmacSecret)

	uv := &dbConnectionValidator{
		EntityDB: ug,
		hmac:     hmac,
	}

	Db = &DatabaseProvider{
		EntityDB: *uv,
	}

	return nil
}

func (ug *DbGorm) Close() error {
	return ug.dbase.Close()
}

func (ug *DbGorm) Ping() error {

	if err := ug.dbase.Ping(); err != nil {
		ug.dbase.Close()
		return errors.New("Connection to DB is not available")
	}

	fmt.Println("Connection ok (ping)", ug)

	return nil
}
