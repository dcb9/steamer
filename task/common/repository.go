package common

import "github.com/dcb9/steamer/database"

var conn = database.Db

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
