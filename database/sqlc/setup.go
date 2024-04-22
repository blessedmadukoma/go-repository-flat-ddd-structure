package database

import (
	"context"

	"goRepositoryPattern/util"
	"log"

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
	// runDBMigration(config.MigrationURL, config.DBSource)

	store := NewStore(connPool)

	return store, nil
}

// runDBMigration runs the database migration
// func runDBMigration(migrationURL string, dbSource string) {
// 	// log.Info().Msg("migrating db...")
// 	migration, err := migrate.New(migrationURL, dbSource)
// 	if err != nil {
// 		// log.Fatal().Msg("cannot create migration instance")
// 	}

// 	if err = migration.Up(); err != nil {
// 		if err == migrate.ErrNoChange {
// 			// log.Info().Msg("no migration changes...")
// 			return
// 		}
// 		// log.Fatal().Msg("cannot migrate db:")
// 	}

// 	// log.Info().Msg("migration successful")
// }

// func RunMigrations() error {
// 	db, err := c.DB()
// 	if err != nil {
// 		return err
// 	}
// 	log.Print("Running migrations")

// 	if err := goose.Up(db, "database/migrations"); err != nil {
// 		return err
// 	}
// 	return nil
// }
