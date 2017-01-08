package pg

import (
	"fmt"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/fortytw2/hydrocarbon"
	"github.com/fortytw2/hydrocarbon/internal/pgmigrate"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

//go:generate go-bindata -pkg pg -o migrations_generated.go schema/

// Store provides basic persistence primitives
type Store struct {
	db *sqlx.DB
}

// NewStore creates a primitive persistence layer
func NewStore(dsn string) (hydrocarbon.PrimitiveStore, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	err = Migrate(db)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func Migrate(db *sqlx.DB) error {
	migrations, err := pgmigrate.LoadMigrations(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "schema"})
	if err != nil {
		return err
	}

	ran, err := pgmigrate.DefaultConfig.Migrate(db.DB, migrations)
	if err != nil {
		return err
	}

	fmt.Println(ran)

	return nil
}
