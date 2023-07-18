package models

import (
	"App/internal/database"
	"App/internal/modules/hash"
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseProvider struct {
	EntityDB
	Db *gorm.DB
}

var Db *DatabaseProvider

var InitGorm *DbGorm

func InitDB() (*DbGorm, error) {
	conn, err := gorm.Open(postgres.Open(database.BuildConnectionString()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

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
		Db:    conn,
		Dbase: dbase,
	}, nil
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

	err = ug.Ping()

	if err != nil {
		return err
	}

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

func (ug *DbGorm) Begin() *gorm.DB {
	return ug.Db.Begin()
}

func (ug *DbGorm) Close() error {
	return ug.Dbase.Close()
}

func (ug *DbGorm) Ping() error {

	if err := ug.Dbase.Ping(); err != nil {
		ug.Dbase.Close()
		return errors.New("Connection to DB is not available")
	}

	fmt.Println("Connection ok (ping)", ug)

	return nil
}
