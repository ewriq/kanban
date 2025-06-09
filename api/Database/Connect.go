package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/auth")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Print("Damm, In connected !")
	}
}