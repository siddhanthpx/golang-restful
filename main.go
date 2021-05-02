package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://pdzenyaibpistf:00e77e74cf97e9e4fd9d647900a8e380a3bf132fcc48ca8927e2ca8b9680d9b4@ec2-54-166-167-192.compute-1.amazonaws.com:5432/db2vmap0n2c4of"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to db", db.Name())
}
