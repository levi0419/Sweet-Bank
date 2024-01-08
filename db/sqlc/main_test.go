package sqlc

import (
	"database/sql"
	"testing"
	"log"
	"os"
	_"github.com/lib/pq"
)


var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password1234@localhost:5433/personal_test?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())


}