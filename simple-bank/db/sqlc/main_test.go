package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	testDB, err = pgxpool.New(
		context.Background(),
		"postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable",
	)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	code := m.Run()

	testDB.Close()

	os.Exit(code)
}