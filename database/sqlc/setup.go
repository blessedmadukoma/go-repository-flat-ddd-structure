package database

import (
	"context"

	"goRepositoryPattern/util"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

// ConnectDataBase handles the connection to the database
func ConnectDataBase(config util.Config) (Store, error) {
	var err error

	// connect to database
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		// log.Fatal().Msg("cannot connect to db")
		log.Fatal("cannot connect to db:", err)
	}

	// run db migration
	runDBMigration(config.MigrationUrl, config.DBSource)

	store := NewStore(connPool)

	return store, nil
}

// runDBMigration runs the database migration
func runDBMigration(migrationURL string, dbSource string) {
	log.Println("migrating db...")
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Println("cannot create migration instance:", err)
	}

	if err = migration.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("no migration changes...")
			return
		}
		log.Println("cannot migrate db:", err)
	}

	log.Println("migration successful...")
}
