
package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)
var DB *gorm.DB



func InitDB() *gorm.DB {
    dsn := "user=postgres password=kanaka dbname=raju sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

	return db
}

