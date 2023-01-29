package database

import (
	"fmt"
	"pokedex-graphql/graph/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB
var err error

func ConnectDB() {
	DBInstance, err = gorm.Open(sqlite.Open("pokemon.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Database connection attempt was unsuccessful.....")
	} else {
		fmt.Println("Database Connected successfully.....")
	}
}

func CreateDB() {
	// Create a database
	DBInstance.Exec("CREATE DATABASE IF NOT EXISTS Blog_Posts")
	// make the database available to this connection
	DBInstance.Exec("USE Blog_Posts")
}

func MigrateDB() {
	// migrate and sync the model to create a db table
	DBInstance.AutoMigrate(&model.Pokemon{})
	fmt.Println("Database migration completed....")
}
