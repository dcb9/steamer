package database

import (
	"testing"
	"os"
)

func TestConn(t *testing.T) {
	db := Conn(
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
		"tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")",
		os.Getenv("DB_NAME"), "utf8mb4,utf8",
	)
	err := db.Ping()
	if err != nil {
		t.Errorf("Connect to mysql error: %q %s", err, os.Getenv("DB_USER"))
	}
}
