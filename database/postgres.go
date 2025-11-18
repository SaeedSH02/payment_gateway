package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() error {

	dsn := "host=localhost user=postgres password=123 dbname=goTodo port=54321 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Faild to connect to Database!....")
		return err
	}

	pg, err := db.DB()
	if err != nil {
		log.Fatal("Faild to connect to Database!....")
		return err
	}

	if err := pg.Ping(); err != nil {
		log.Panic("DB Connection is not established")
		return err
	}


	db.AutoMigrate()
	return nil
}
