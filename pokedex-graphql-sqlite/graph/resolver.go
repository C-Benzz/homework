package graph

import (
	"fmt"
	"pokedex-graphql/graph/model"

	"github.com/google/uuid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Database struct {
	pokemon map[string]model.Pokemon
}

func NewDatabase() Database {
	var data = make(map[string]model.Pokemon)
	return Database{data}
}

func (db *Database) CreatePokemon(input *model.Pokemon) error {
	newID := uuid.New().String()
	input.ID = newID
	db.pokemon[newID] = *input
	return nil
}

func (db *Database) UpdatePokemon(input model.Pokemon) error {
	if _, key := db.pokemon[input.ID]; !key {
		return fmt.Errorf("pokemon: %s was not found", input.Name)
	}
	db.pokemon[input.ID] = input
	return nil
}

func (db *Database) DeletePokemon(ID string) error {
	if _, key := db.pokemon[ID]; !key {
		return fmt.Errorf("pokemon: %s was not found", ID)
	}
	delete(db.pokemon, ID)
	return nil
}

func (db *Database) AllPokemon() []*model.Pokemon {
	listPokemon := []*model.Pokemon{}
	for _, n := range db.pokemon {
		address := n
		listPokemon = append(listPokemon, &address)
	}
	return listPokemon
}

func (db *Database) PokemonByID(ID string) (*model.Pokemon, error) {
	if ret, ok := db.pokemon[ID]; ok {
		return &ret, nil
	} else {
		return nil, fmt.Errorf("movie id: %s was not found", ID)
	}
}

type Resolver struct {
	DB Database
}
