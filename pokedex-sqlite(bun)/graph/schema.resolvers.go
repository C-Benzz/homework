package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"pokedex-bun/database"
)

// Create is the resolver for the Create field.
func (r *mutationResolver) Create(ctx context.Context, input database.NewPokemon) (*database.Pokemon, error) {
	Pokemon := &database.Pokemon{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Type:        input.Type,
		Abilities:   input.Abilities,
	}
	return r.DB.CreatePokemon(ctx, *Pokemon)
}

// Update is the resolver for the Update field.
func (r *mutationResolver) Update(ctx context.Context, id int, input database.NewPokemon) (*database.Pokemon, error) {
	UpdatePokemon := &database.Pokemon{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Type:        input.Type,
		Abilities:   input.Abilities,
	}
	return r.DB.UpdatePokemon(ctx, *UpdatePokemon, id)
}

// Delete is the resolver for the Delete field.
func (r *mutationResolver) Delete(ctx context.Context, id int) (bool, error) {
	return r.DB.DeletePokemon(ctx, id)
}

// AllPokemon is the resolver for the AllPokemon field.
func (r *queryResolver) AllPokemon(ctx context.Context) ([]*database.Pokemon, error) {
	return r.DB.AllPokemon(ctx)
}

// GetPokemonByID is the resolver for the GetPokemonByID field.
func (r *queryResolver) GetPokemonByID(ctx context.Context, id int) (*database.Pokemon, error) {
	return r.DB.GetPokemonByID(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
