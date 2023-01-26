package graph

import (
	"fmt"
	"pokedex-graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Database struct {
	pokemon map[string]model.Pokemon
}

func (db *Database) CreatePokemon(input *model.Pokemon) error {
	db.pokemon[input.Name] = *input
	return nil
}

func (db *Database) UpdatePokemon(input *model.Pokemon) error {
	if _, key := db.pokemon[input.Name]; !key {
		return fmt.Errorf("pokemon: %s was not found", input.Name)
	}
	db.pokemon[input.Name] = *input
	return nil
}

func (db *Database) DeletePokemon(name string) error {
	if _, key := db.pokemon[name]; !key {
		return fmt.Errorf("pokemon: %s was not found", name)
	}
	delete(db.pokemon, name)
	return nil
}

func (db *Database) AllPokemon() []*model.Pokemon {
	listPokemon := []*model.Pokemon{}
	for _, n := range db.pokemon {
		listPokemon = append(listPokemon, &n)
	}
	return listPokemon
}

type Resolver struct {
	DB Database
}
