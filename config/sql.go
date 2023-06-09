package config

import (
	"github.com/go-sql-driver/mysql"
	"os"
)

func Sql() (conf mysql.Config) {
	conf = mysql.Config{
		User:                 "root",
		Passwd:               os.Getenv("MARIADB_ROOT_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "database:3306",
		DBName:               os.Getenv("MARIADB_DATABASE"),
		AllowNativePasswords: true,
	}
	return conf
}