package database

import (
	"testing"
	"os"
	"github.com/dcb9/steamer/app"
)

func TestConn(t *testing.T) {
	db := Conn(app.MyDbParams)
	err := db.Ping()
	if err != nil {
		t.Errorf("Connect to mysql error: %q %s", err, os.Getenv("DB_USER"))
	}
}
