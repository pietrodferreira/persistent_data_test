package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	"log"
	"strconv"
)

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

func getRows(rows *sql.Rows) []UserInput{
	var tables []UserInput
	for rows.Next() {
		var userInput UserInput
		err := rows.Scan(&userInput.id, &userInput.data)
		check("sql.Rows.Scan()", err)
		tables = append(tables, userInput)
	}
	return tables
}

func listRows(tables []UserInput) {
	if len(tables) == 0 {
		fmt.Println("NO DATA!")
	} else {
		for _, v := range tables {
			fmt.Println(v.String())
		}
	}

}
func insertInputs(inputs []string, db *sql.DB) {
	for _, v := range inputs {
		_, err := db.Exec("INSERT INTO user_input (data) VALUES (?)", v)
		check("sql.DB.Exec()", err)
	}
}

func getInputs(sc *bufio.Scanner) []string {
	fmt.Print("number of inputs >> ")
	if !sc.Scan() {
		check("bufio.Scanner().Scan()", sc.Err())
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
	return inputs
}
func askDelete(db *sql.DB, sc *bufio.Scanner) {
	fmt.Print("delete all? [y/n] >> ")
	if !sc.Scan() {
		check("bufio.Scanner.Scan()", sc.Err())
	}
	input := sc.Text()
	if input == "y" {
		_, err := db.Exec("DELETE FROM user_input")
		check("sql.DB.Exec()", err)
		_, err = db.Exec("DELETE FROM SQLITE_SEQUENCE WHERE name='user_input'")
	} else if input == "n" {
		fmt.Println("Ok")
	} else {
		fmt.Println("Invalid")
	}
}
