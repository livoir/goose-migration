package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	db2 "github.com/livoir/goose-migration/db/sqlc"
	"log"
	"os"
	"testing"
)

var userRepository *UserRepositoryImpl

func TestMain(m *testing.M) {
	var err error
	var queries *db2.Queries
	var db *sql.DB
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/goose_migration_test")
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	queries = db2.New(db)

	userRepository = NewUserRepositoryImpl(queries)

	os.Exit(m.Run())
}
