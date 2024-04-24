package config

import (
	"database/sql"
	"log"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

func LoadDB() *sql.DB {
	//connectionStr := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_URL, ENV.DB_DATABASE)
	//db, err := sql.Open(mysql.Open(connectionStr), &gorm.Config{})

	db, err := sql.Open("mysql", ENV.DB_USERNAME+"@tcp("+ENV.DB_URL+")/"+ENV.DB_DATABASE)

	if err != nil {
		log.Fatal("Failed to connection database")
	}

	return db

}
