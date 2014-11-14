package orm

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB_NAME string = "./sqlite3.db"

// InitDB initializes a new sqlite3 database; creates file and adds tables based on the Model objects passed as arguments.
func InitDB(m ...*Model) {
	// TODO: refactor this to create the database file ONLY. Table creation should be done by another function.

	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Loop through the model args, and create a table for each model
	for _, v := range m {
		command := v.GenCreateTable()
		_, err = db.Exec(command)
		if err != nil {
			log.Printf("%q: %s\n", err, command)
			return
		}
	}
}

// InsertDB inserts data into the database based on the contents of the Models passed as arguments.
func InsertDB(m ...*Model) error {

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

	command := m[0].GenInsertInto()
	log.Println(command)
	stmt, err := tx.Prepare(command)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	for _, v := range m {
		log.Println("Inserting...", v.String())

		_, err = stmt.Exec(v.GenValues()...)
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Println(v.String(), " inserted!")
	}

	tx.Commit()
	return nil
}

// GetResultsDB returns an array of Model objects based on a Model object passed as an argument.
// Currently, this only returns all the objects of the type passed to args. No filtering or joins are possible yet.
// Also, this function currently doesn't work, due to SetFromByteArray() not working properly.
func GetResultsDB(m *Model) []*Model {
	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	selectStmt := fmt.Sprintf("select * from %ss", m.Meta.modelName)
	rows, err := db.Query(selectStmt)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	models := make([]*Model, 2)

	for rows.Next() {

		model := NewModel(ExampleModel{})

		data := make([]interface{}, model.Meta.numFields)
		dataPtrs := make([]interface{}, model.Meta.numFields)

		for i, _ := range data {
			dataPtrs[i] = &data[i]
		}

		err := rows.Scan(dataPtrs...)

		model.SetFromByteArray(data)
		models = append(models, model)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	}

	return models
}
