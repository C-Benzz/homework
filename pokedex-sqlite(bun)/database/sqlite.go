package database

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"pokedex-bun/graph/model"
	"strconv"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type DatabaseBun struct {
	db *bun.DB
}

type Pokemon struct {
	bun.BaseModel `bun:"table:pokemon,alias:u"`
	ID            int           `bun:"id,pk,autoincrement"`
	Name          string        `bun:"name,notnull"`
	Description   string        `bun:"description,notnull"`
	Category      string        `bun:"category,notnull"`
	Type          []PokemonType `bun:"type,notnull"`
	Abilities     []string      `bun:"abilities,notnull"`
}
type PokemonType string

const (
	PokemonTypeBug      PokemonType = "Bug"
	PokemonTypeFlying   PokemonType = "Flying"
	PokemonTypeFire     PokemonType = "Fire"
	PokemonTypeGrass    PokemonType = "Grass"
	PokemonTypeWater    PokemonType = "Water"
	PokemonTypePoison   PokemonType = "Poison"
	PokemonTypeElectric PokemonType = "Electric"
)

var AllPokemonType = []PokemonType{
	PokemonTypeBug,
	PokemonTypeFlying,
	PokemonTypeFire,
	PokemonTypeGrass,
	PokemonTypeWater,
	PokemonTypePoison,
	PokemonTypeElectric,
}

func (e PokemonType) IsValid() bool {
	switch e {
	case PokemonTypeBug, PokemonTypeFlying, PokemonTypeFire, PokemonTypeGrass, PokemonTypeWater, PokemonTypePoison, PokemonTypeElectric:
		return true
	}
	return false
}

func (e PokemonType) String() string {
	return string(e)
}

func (e *PokemonType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PokemonType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PokemonType", str)
	}
	return nil
}

func (e PokemonType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

var ctx = context.Background()

// var _ bun.AfterCreateTableHook = (*Pokemon)(nil)

// func (*Pokemon) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
// 	_, err := query.DB().NewCreateIndex().
// 		Model((*Pokemon)(nil)).
// 		Index("pokemon_id_idx").
// 		Column("pokemon_id").
// 		Exec(ctx)
// 	return err
// }

func ConnectDatabase() DatabaseBun {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:pokemon.db")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	if _, err := db.NewCreateTable().Model((*Pokemon)(nil)).Exec(ctx); err != nil {
		panic(err)
	}
	return DatabaseBun{db}
}

func (d DatabaseBun) CreatePokemon(ctx context.Context, p model.Pokemon) (*model.Pokemon, error) {
	res, err := d.db.NewInsert().Model(&p).Exec(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("create pokemon: %v", res)
	return &p, nil
}

func (d DatabaseBun) UpdatePokemon(ctx context.Context, p model.Pokemon, index int) {
	res, err := d.db.NewUpdate().Model(&p).Where("id = ?", index).Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("update pokemon: %v", res)
}

func (d *DatabaseBun) DeletePokemon(ctx context.Context, index int) (bool, error) {
	res, err := d.db.NewDelete().Model((*Pokemon)(nil)).Where("id = ?", index).Exec(ctx)
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
