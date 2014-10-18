package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var Conn *gorp.DbMap
var db gorm.DB
var err error
var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

func Setup() {
	db, err := gorm.Open("sqlite3", "/tmp/data.db")
	db.LogMode(false)
	db.SetLogger(Logger)
	if err != nil {
		Logger.Println(err)
		return err
	}
	_, err = os.Stat(config.Conf.DBPath)
	if err != nil {
		// configure table mapping e.g.
		db.CreateTable(Application{})
	}

	return nil
}
