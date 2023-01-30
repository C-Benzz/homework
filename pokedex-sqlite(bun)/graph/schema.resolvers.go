package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"pokedex-bun/graph/model"
)

// Create is the resolver for the Create field.
func (r *mutationResolver) Create(ctx context.Context, input model.NewPokemon) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented: Create - Create"))
}

// Update is the resolver for the Update field.
func (r *mutationResolver) Update(ctx context.Context, id int, input model.NewPokemon) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented: Update - Update"))
}

// Delete is the resolver for the Delete field.
func (r *mutationResolver) Delete(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented: Delete - Delete"))
}

// AllPokemon is the resolver for the AllPokemon field.
func (r *queryResolver) AllPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented: AllPokemon - AllPokemon"))
}

// GetPokemonByID is the resolver for the GetPokemonByID field.
func (r *queryResolver) GetPokemonByID(ctx context.Context, id string) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented: GetPokemonByID - GetPokemonByID"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
