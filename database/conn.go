package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
	"os"
)

var Db *sql.DB

func Conn(username, password, protocleAddr, dbname, charset string) (*sql.DB) {
	if Db == nil {
		conn, err := sql.Open("mysql", username + ":" + password + "@" + protocleAddr + "/" + dbname + "?charset=" + charset)
		if err != nil {
			log.Fatalf("Connect mysql error: %q", err)
		}
		Db = conn
		fmt.Println("======= sql.Open ======")
	}

	return Db
}

func init() {
	Conn(
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
		"tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")",
		os.Getenv("DB_NAME"), "utf8mb4,utf8",
	)

	Db.SetMaxOpenConns(200)
	Db.SetMaxIdleConns(100)
}
