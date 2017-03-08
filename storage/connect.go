package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// SQLite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// Postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ConnectMySQL is used to connect to a MySQL database
func ConnectMySQL(host, user, password, dbname string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, dbname))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ConnectPostgreSQL is used to connect to a PostgreSQL database
func ConnectPostgreSQL(host, user, password, dbname, sslmode string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
			host, user, dbname, sslmode, password))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ConnectSQLite3 is used to connect to a SQLite3 database
func ConnectSQLite3(host, user, password, dbname, sslmode string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./cronmon.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
