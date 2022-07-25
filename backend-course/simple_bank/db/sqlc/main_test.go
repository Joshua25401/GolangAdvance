package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://joshua:capitalX123@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		fmt.Println("Cannot connect to db", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
