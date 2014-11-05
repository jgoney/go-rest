package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB_NAME string = "./sqlite3.db"

// Initializes a new sqlite3 database; creates file and adds table
func initDB(m ...model) {

	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Loop through the model args, and create a table for each model
	for _, v := range m {
		command := v.genCreateTable()
		_, err = db.Exec(command)
		if err != nil {
			log.Printf("%q: %s\n", err, command)
			return
		}
	}
}

func insertDB(m ...model) error {
	log.Println("Inserting...", m)

	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}

	stmt, err := tx.Prepare("insert into names(first_name, last_name, email, gender) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	for _, v := range m {
		_, err = stmt.Exec(v.firstname, v.lastname, v.email, v.gender)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	tx.Commit()
	log.Println(m, " inserted!")
	return nil
}

func getResultsDB() {
	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// var (
	// 	id   int
	// 	name string
	// )

	// rows, err := db.Query("select id, name from foo where id = ?", 1)
	rows, err := db.Query("select * from names")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		model := model{}
		err := rows.Scan(&model.id,
			&model.firstname,
			&model.lastname,
			&model.email,
			&model.gender)

		if err != nil {
			log.Fatal(err)
		}

		log.Println(model)
	}

	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	}
}
