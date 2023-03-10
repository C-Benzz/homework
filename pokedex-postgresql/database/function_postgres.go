package database

import (
	"context"
	"database/sql"
	"fmt"
	"pokedex-bun/graph/model"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type DatabaseBun struct {
	db *bun.DB
}

var ctx = context.Background()

func ConnectDatabase() DatabaseBun {
	dsn := "postgres://postgres:mysecretpassword@postgres:5432/postgres?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	if _, err := db.NewCreateTable().IfNotExists().Model((*model.Pokemon)(nil)).Exec(ctx); err != nil {
		panic(err)
	}
	return DatabaseBun{db}
}

func (d *DatabaseBun) CreatePokemon(ctx context.Context, p model.Pokemon) (*model.Pokemon, error) {
	res, err := d.db.NewInsert().Model(&p).Exec(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("create pokemon: %v", res)
	return &p, nil
}

func (d *DatabaseBun) UpdatePokemon(ctx context.Context, p model.Pokemon, id int) (*model.Pokemon, error) {
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
	p := &model.Pokemon{
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

func (d *DatabaseBun) AllPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	pokemon := []*model.Pokemon{}
	err := d.db.NewSelect().Model(&pokemon).Scan(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("all pokemon: %v", pokemon)
	return pokemon, nil
}

func (d *DatabaseBun) GetPokemonByID(ctx context.Context, id int) (*model.Pokemon, error) {
	pokemon := new(model.Pokemon)
	if err := d.db.NewSelect().Model(pokemon).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return pokemon, nil
}
