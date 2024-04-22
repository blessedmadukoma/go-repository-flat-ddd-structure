package db_test

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	db "goRepositoryPattern/db/sqlc"
// 	"log"
// 	"os"
// 	"testing"

// 	"goRepositoryPattern/util"

// 	"github.com/jackc/pgx/v5/pgxpool"
// 	// _ "github.com/lib/pq"
// )

// var testQuery *db.Queries

// func TestMain(m *testing.M) {
// 	config := util.LoadEnvConfig()

// 	connPool, err := pgxpool.New(context.Background(), config.DBSource)
// 	if err != nil {
// 		log.Fatal().Msg("cannot connect to db")
// 	}

// 	conn, err := sql.Open(config.DBDriver, fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", config.DBDriver, config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME))

// 	if err != nil {
// 		log.Fatal("Could not connect to database:", err)
// 	}

// 	testQuery = db.New(conn)

// 	os.Exit(m.Run())
// }
