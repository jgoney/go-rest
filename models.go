package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type meta struct {
	modelName string
	numFields int
	elements  map[string]string
}

type model struct {
	id        int
	firstname string
	lastname  string
	email     string
	gender    string
	meta      meta
}

// Would be nice to get these dynamically from the package
var typeMappings = map[string]string{
	"nil":       "null",
	"int":       "integer",
	"int64":     "integer",
	"float64":   "float",
	"bool":      "integer",
	"[]byte":    "blob",
	"string":    "text",
	"time.Time": "timestamp/datetime",
}

// Sets metadata fields via reflection. Used to convert Go types to SQL syntax
func (m *model) setMetaFields() {
	m.meta.modelName = reflect.TypeOf(*m).Name()
	m.meta.numFields = reflect.TypeOf(*m).NumField() - 1 // Minus 1 since we're ignoring the metadata struct
	m.meta.elements = make(map[string]string)

	for i := 0; i < m.meta.numFields; i++ {
		m.meta.elements[reflect.TypeOf(*m).Field(i).Name] = typeMappings[reflect.TypeOf(*m).Field(i).Type.Name()]
	}
}

// Sets fields on a model from POST URL values
func (m *model) setFieldsFromPOST(urlv url.Values) {
	m.firstname = strings.Join(urlv["firstname"], " ")
	m.lastname = strings.Join(urlv["lastname"], " ")
	m.email = strings.Join(urlv["email"], " ")
	m.gender = strings.Join(urlv["gender"], " ")
}

// Returns a string used to create a table representing
func (m *model) genCreateTable() string {

	// "create table names (id integer not null primary key autoincrement, first_name text, last_name text, email text, gender text);"
	modelNames := m.meta.modelName + "s"
	elements := make([]string, m.meta.numFields)

	i := 0
	for k, v := range m.meta.elements {
		elements[i] = fmt.Sprintf("%s %s", k, v)
		i++
	}

	command := fmt.Sprintf("create table %s (%s);", modelNames, strings.Join(elements, ", "))
	return command
}

func (m *model) genInsertInto() string {
	// "insert into names(first_name, last_name, email, gender) values(?, ?, ?, ?)"
	return "nil"
}
