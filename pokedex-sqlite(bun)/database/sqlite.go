package database

import (
	"context"
	"database/sql"
	"fmt"
	"pokedex-bun/graph/model"

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
	if _, err := db.NewCreateTable().Model((*model.Pokemon)(nil)).Exec(ctx); err != nil {
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

func (d *DatabaseBun) UpdatePokemon(ctx context.Context, p model.Pokemon, index int) (*model.Pokemon, error) {
	res, err := d.db.NewUpdate().Model(&p).Where("id = ?", index).Exec(ctx)
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
	res, err := d.db.NewDelete().Model(p).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return false, err
	}
	fmt.Printf("delete pokemon: %v", res)
	return true, nil
}

func (d *DatabaseBun) AllPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	pokemon := make([]*model.Pokemon, 0)
	result, err := d.db.NewSelect().Model(&pokemon).Exec(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("all pokemon: %v", result)
	return pokemon, nil
}
