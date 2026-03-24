package main

import (
	"bufio"
	"database/sql"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./data.db")
	check("sql.DB.Open()", err)
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user_input (id INTEGER PRIMARY KEY AUTOINCREMENT, data TEXT)")
	check("sql.DB.Exec()", err)

	rows, err := db.Query("SELECT id, data FROM user_input")
	check("sql.DB.Query()", err)
	defer rows.Close()

	tables := getRows(rows)

	sc := bufio.NewScanner(os.Stdin)

	listRows(tables)

	if len(tables) != 0 {
		askDelete(db, sc)
	}

	inputs := getInputs(sc)
	insertInputs(inputs, db)

}
