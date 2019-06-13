package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func InitDatabase (database string, url string) {
    var err error
    db, err = gorm.Open(database, url)
    if err != nil {
        defer db.Close()
        panic(err)
    }
}

func Database () *gorm.DB {
    return db
}