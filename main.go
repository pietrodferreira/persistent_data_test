package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/glebarez/go-sqlite"
)

//TODO: Extract functionality into function, upgrade the delete, add edit (update)

func main() {
	db, err := sql.Open("sqlite", "./data.db")
	check("sql.DB.Open", err)
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user_input (id INTEGER PRIMARY KEY AUTOINCREMENT, data TEXT)")
	check("sql.DB.Exec()", err)

	rows, err := db.Query("SELECT id, data FROM user_input")
	check("sql.DB.Query()", err)
	defer rows.Close()

	var tables []UserInput
	for rows.Next() {
		var userInput UserInput
		err = rows.Scan(&userInput.id, &userInput.data)
		check("sql.Rows.Scan()", err)
		tables = append(tables, userInput)
	}

	sc := bufio.NewScanner(os.Stdin)

	if len(tables) == 0 {
		fmt.Println("NO DATA!")
	} else {
		for _, v := range tables {
			fmt.Println(v.String())
		}
		fmt.Print("delete all? [y/n] >> ")
		if !sc.Scan() {
			check("bufio.Scanner.Scan()", err)
		}
		input := sc.Text()
		if input == "y" {
			_, err = db.Exec("DELETE FROM user_input")
			check("sql.DB.Exec()", err)
			_, err = db.Exec("DELETE FROM SQLITE_SEQUENCE WHERE name='user_input'")
		}
	}

	fmt.Print(">> ")
	if !sc.Scan() {
		check("bufio.Scanner().Scan()", err)
	}
	numberOfInputs, err := strconv.Atoi(sc.Text())
	check("strconv.Atoi", err)

	var inputs []string
	for i := range numberOfInputs {
		fmt.Printf("input[%d] = ", i)
		if !sc.Scan() {
			check("bufio.Scanner.Scan()", err)
		}
		inputs = append(inputs, sc.Text())
	}
	for _, v := range inputs {
		_, err := db.Exec("INSERT INTO user_input (data) VALUES (?)", v)
		check("sql.DB.Exec()", err)
	}
}

type UserInput struct {
	id   int
	data string
}

func (u UserInput) String() string {
	return fmt.Sprintf("[%d] = %s", u.id, u.data)
}

func check(where string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", where, err)
	}
}
