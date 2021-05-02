package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://izrblupa:wnPCFvuddXqtYVEl1Vaxa0APThO_OZ1x@john.db.elephantsql.com:5432/izrblupa"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to db", db.Name())
}
