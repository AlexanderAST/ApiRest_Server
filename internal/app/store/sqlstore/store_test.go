package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=postgres_test user=postgres port=5436 password=secret sslmode=disable"
	}
	os.Exit(m.Run())
}
