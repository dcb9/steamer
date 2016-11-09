package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
	"github.com/dcb9/steamer/app"
	"github.com/rubenv/sql-migrate"
)

var Db *sql.DB

func Conn(p app.DbParams) (*sql.DB) {
	if Db == nil {
		//  https://bitbucket.org/liamstask/goose/issues/62/scan-error-on-column-index-0-unsupported
		source := fmt.Sprintf(
			"%s:%s@%s/%s?charset=%s&parseTime=true",
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

	n, err := migrate.Exec(Db, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Migrate error: %q\n", err)
	}
	log.Printf("Applied %d migrations\n", n)
	Db.SetMaxOpenConns(app.MyDbParams.MaxOpen)
	Db.SetMaxIdleConns(app.MyDbParams.MaxIdle)
}
