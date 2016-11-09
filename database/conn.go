package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
	"github.com/dcb9/steamer/app"
)

var Db *sql.DB

func Conn(p app.DbParams) (*sql.DB) {
	if Db == nil {
		source := fmt.Sprintf(
			"%s:%s@%s/%s?charset=%s",
			p.User, p.Pass, p.ProtocleAddr, p.Name, p.Charset,
		);
		conn, err := sql.Open("mysql", source)
		if err != nil {
			log.Fatalf("Connect mysql error: %q", err)
		}
		err = conn.Ping()
		if err != nil {
			log.Fatalf("Ping mysql error: %q", err)
		}

		Db = conn
		fmt.Println("======= sql.Open ======")
	}

	return Db
}

func init() {
	Conn(app.MyDbParams)

	Db.SetMaxOpenConns(app.MyDbParams.MaxOpen)
	Db.SetMaxIdleConns(app.MyDbParams.MaxIdle)
}
