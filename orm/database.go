package orm

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB_NAME string = "./sqlite3.db"

// Initializes a new sqlite3 database; creates file and adds table
func InitDB(m ...*Model) {

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

		_, err = stmt.Exec(v.GenValueString()...)
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Println(v.String(), " inserted!")
	}

	tx.Commit()
	return nil
}

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

		model := NewModel(MyModel{})

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
