package db

import (
	"fmt"
	"imbd_goroutine_concurency/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect(config config.TomlConfig) *sqlx.DB {

	strConn := fmt.Sprintf("%s:%s@(%s:%d)/%s", config.Database.User, config.Database.Pass, config.Database.Host, config.Database.Port, config.Database.Name)

	db, err := sqlx.Connect(config.Database.Type, strConn)
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	log.Println("Connect database ok")
	return db
}

func Close(db *sqlx.DB) {
	db.Close()
}
