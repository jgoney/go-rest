/* Package orm provides a simple, dynamic mapping between Go types (provided in a struct) and SQL entries.

This package is very immature; it's missing key functionality, and hasn't been tested in any sort of production setting. You'd probably be better off looking at other ORM options for any serious project.
*/

package orm

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

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

// Meta is a simple type representing metadata attached to a Model object.
type Meta struct {
	modelName string
	numFields int
	elements  [][]string
	isSet     bool
}

// Modeler is a placeholder type for the actual user model to be associated with a given Model object.
type Modeler interface{}

// Model is a representation of a user's model data as a Go object.
type Model struct {
	Modeler
	Meta Meta
}

// NewModel is a constructor function that accepts a Modeler type (a struct), sets the associated metadata for that struct, and returns a pointer to a Model object.
func NewModel(modeler Modeler) *Model {
	m := new(Model)
	m.Modeler = modeler
	m.SetMetaFields()
	return m
}

// String implements the Stringer interface for a Model object.
func (m *Model) String() string {
	return fmt.Sprintf("%v : %v", m.Modeler, m.Meta)
}

// SetMetaFields sets metadata fields via reflection. Used to convert Go types to SQL syntax.
func (m *Model) SetMetaFields() {

	m.Meta.modelName = reflect.TypeOf(m.Modeler).Name()
	m.Meta.numFields = reflect.TypeOf(m.Modeler).NumField()
	m.Meta.elements = make([][]string, m.Meta.numFields)

	for i := 0; i < m.Meta.numFields; i++ {
		m.Meta.elements[i] = []string{reflect.TypeOf(m.Modeler).Field(i).Name, typeMappings[reflect.TypeOf(m.Modeler).Field(i).Type.Name()]}
	}

	m.Meta.isSet = true
}

// SetFieldsFromPOST sets fields on a model from POST URL values. Currently not implemented.
// func (m *Model) SetFieldsFromPOST(urlv url.Values) {
// 	m.Firstname = strings.Join(urlv["Firstname"], " ")
// 	m.Lastname = strings.Join(urlv["Lastname"], " ")
// 	m.Email = strings.Join(urlv["Email"], " ")
// 	m.Gender = strings.Join(urlv["Gender"], " ")
// }

// GenCreateTable generates and returns a string used to create a table representing the data fields contained in the Model object.
func (m *Model) GenCreateTable() string {

	modelNames := m.Meta.modelName + "s"
	fields := make([]string, m.Meta.numFields)

	for i, v := range m.Meta.elements {
		fields[i] = strings.Join(v, " ")
	}

	command := fmt.Sprintf("create table %s (%s);", modelNames, strings.Join(fields, ", "))
	return command
}

// GenInsertInto generates and returns a string used to insert data into a table corresponding to the data in the Model object.
func (m *Model) GenInsertInto() string {

	modelNames := m.Meta.modelName + "s"
	fields := make([]string, m.Meta.numFields)
	placeholders := make([]string, m.Meta.numFields)

	for i, v := range m.Meta.elements {
		fields[i] = v[0]
		placeholders[i] = "?"
	}

	command := fmt.Sprintf("insert into %s(%s) values(%s);", modelNames, strings.Join(fields, ", "), strings.Join(placeholders, ", "))
	return command
}

// GenValues generates and returns an array of interface{} objects which are passed as arguments to SQL queries.
// This mechanism is required to avoid SQL injection.
func (m *Model) GenValues() []interface{} {

	values := make([]interface{}, m.Meta.numFields)

	for i := range values {

		r := reflect.ValueOf(m.Modeler)
		f := reflect.Indirect(r).Field(i)

		// TODO: finish this switch with the rest of the supported types.
		switch f.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			values[i] = strconv.FormatInt(f.Int(), 10)
		case reflect.String:
			values[i] = f.String()
			// etc...
		}
	}

	return values
}

// SetFromByteArray is meant to set a new Model object's fields based on a SQL query.
// This method is currently broken.
func (m *Model) SetFromByteArray(byteArray []interface{}) {

	if len(byteArray) != m.Meta.numFields {
		panic("Too many or too few values to unpack.")
	}

	for i := 0; i < m.Meta.numFields; i++ {

		r := reflect.ValueOf(m.Modeler)
		f := reflect.Indirect(r).Field(i).Kind()

		fmt.Println(r, f)
		fmt.Println("type of r:", r.Type())
		fmt.Println("settability of r:", r.CanSet())
		fmt.Println("kind of r:", r.Kind())

		// TODO: finish this switch with these types:
		// "nil":       "null",
		// "[]byte":    "blob",
		// "string":    "text",
		// "time.Time": "timestamp/datetime",

		// switch f.Kind() {
		// case reflect.String:
		// 	f.SetString(fmt.Sprintf("%s", byteArray[i]))
		// case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// 	fmt.Println(byteArray[i].(int64))
		// 	fmt.Println(byteArray[i])
		// 	f.SetInt(f.Int())
		// case reflect.Float32, reflect.Float64:
		// 	f.SetFloat(byteArray[i].(float64))
		// case reflect.Bool:
		// 	f.SetBool(byteArray[i].(bool))
		// default:
		// 	panic(fmt.Sprintf("Not a valid type: %v\n", f.Kind()))
		// }
	}
}
