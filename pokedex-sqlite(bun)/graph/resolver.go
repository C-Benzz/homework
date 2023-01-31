package graph

import (
	"pokedex-bun/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// type Database struct {
// 	pokemon map[string]model.Pokemon
// }

// func NewDatabase() Database {
// 	var data = make(map[string]model.Pokemon)
// 	return Database{data}
// }

type Resolver struct {
	DB database.DatabaseBun
	// db Database
}
