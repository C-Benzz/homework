package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type DatabaseBun struct {
	db *bun.DB
}

var ctx = context.Background()

func ConnectDatabase() DatabaseBun {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:pokemon.db")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	if _, err := db.NewCreateTable().IfNotExists().Model((*Pokemon)(nil)).Exec(ctx); err != nil {
		panic(err)
	}
	return DatabaseBun{db}
}

func (d *DatabaseBun) CreatePokemon(ctx context.Context, p Pokemon) (*Pokemon, error) {
	res, err := d.db.NewInsert().Model(&p).Exec(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("create pokemon: %v", res)
	return &p, nil
}

func (d *DatabaseBun) UpdatePokemon(ctx context.Context, p Pokemon, id int) (*Pokemon, error) {
	// errorid it's feel like this variable is a id of an error
	// you might change it to `err` instead
	if _, errorid := d.GetPokemonByID(ctx, id); errorid != nil {
		return nil, errorid
	}
	res, err := d.db.NewUpdate().Model(&p).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("update pokemon: %v", res)
	return &p, nil
}

func (d *DatabaseBun) DeletePokemon(ctx context.Context, id int) (bool, error) {
	p := &Pokemon{
		ID: id,
	}
	if _, errorid := d.GetPokemonByID(ctx, id); errorid != nil {
		return false, errorid
	}
	res, err := d.db.NewDelete().Model(p).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return false, err
	}
	fmt.Printf("delete pokemon: %v", res)
	return true, nil
}

func (d *DatabaseBun) AllPokemon(ctx context.Context) ([]*Pokemon, error) {
	pokemon := []*Pokemon{}
	err := d.db.NewSelect().Model(&pokemon).Scan(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("all pokemon: %v", pokemon)
	return pokemon, nil
}

func (d *DatabaseBun) GetPokemonByID(ctx context.Context, id int) (*Pokemon, error) {
	pokemon := new(Pokemon)
	if err := d.db.NewSelect().Model(pokemon).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return pokemon, nil
}
