package orm

import (
	"fmt"
	"reflect"
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

type Meta struct {
	modelName string
	numFields int
	elements  [][]string
	isSet     bool
}

// type Modeler interface {
// 	String() string
// 	SetMetaFields()
// 	SetFieldsFromPOST(url.Values)
// 	GenCreateTable() string
// 	GenInsertInto() string
// 	GenValueString() []interface{}
// 	ExposeFields() []interface{}
// 	SetFromByteArray([]interface{})
// }

type Modeler interface{}

type Model struct {
	Modeler
	Meta Meta
}

func NewModel(modeler Modeler) *Model {
	m := new(Model)
	m.Modeler = modeler
	m.SetMetaFields()
	return m
}

func (m *Model) String() string {
	return fmt.Sprintf("%v : %v\n", m.Modeler, m.Meta)
}

// Sets metadata fields via reflection. Used to convert Go types to SQL syntax
func (m *Model) SetMetaFields() {

	m.Meta.modelName = reflect.TypeOf(m.Modeler).Name()
	m.Meta.numFields = reflect.TypeOf(m.Modeler).NumField()
	m.Meta.elements = make([][]string, m.Meta.numFields)

	for i := 0; i < m.Meta.numFields; i++ {
		m.Meta.elements[i] = []string{reflect.TypeOf(m.Modeler).Field(i).Name, typeMappings[reflect.TypeOf(m.Modeler).Field(i).Type.Name()]}
	}

	m.Meta.isSet = true
}

// Sets fields on a model from POST URL values
// func (m *Model) SetFieldsFromPOST(urlv url.Values) {
// 	m.Firstname = strings.Join(urlv["Firstname"], " ")
// 	m.Lastname = strings.Join(urlv["Lastname"], " ")
// 	m.Email = strings.Join(urlv["Email"], " ")
// 	m.Gender = strings.Join(urlv["Gender"], " ")
// }

// Returns a string used to create a table representing
func (m *Model) GenCreateTable() string {

	// Check to see if the model's metadata has been set, and set it if not:
	if !m.Meta.isSet {
		m.SetMetaFields()
	}

	// "create table names (id integer not null primary key autoincrement, first_name text, last_name text, email text, gender text);"
	modelNames := m.Meta.modelName + "s"
	fields := make([]string, m.Meta.numFields)

	for i, v := range m.Meta.elements {
		fields[i] = strings.Join(v, " ")
	}

	command := fmt.Sprintf("create table %s (%s);", modelNames, strings.Join(fields, ", "))
	return command
}

func (m *Model) GenInsertInto() string {

	// Check to see if the model's metadata has been set, and set it if not:
	if !m.Meta.isSet {
		m.SetMetaFields()
	}

	// "insert into names(first_name, last_name, email, gender) values(?, ?, ?, ?)"
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

func (m *Model) GenValueString() []interface{} {
	// Check to see if the model's metadata has been set, and set it if not:
	if !m.Meta.isSet {
		m.SetMetaFields()
	}
	values := make([]interface{}, m.Meta.numFields)

	s := reflect.ValueOf(m).Elem()
	for i := range m.Meta.elements {
		values[i] = fmt.Sprintf("%v", s.Field(i).Interface())
	}

	return values
}

func (m *Model) ExposeFields() []interface{} {
	// Check to see if the model's metadata has been set, and set it if not:
	if !m.Meta.isSet {
		m.SetMetaFields()
	}
	values := make([]interface{}, m.Meta.numFields)

	s := reflect.ValueOf(m).Elem()
	typeOfT := s.Type()
	for i := 0; i < m.Meta.numFields; i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
		values[i] = f.Interface()
	}
	fmt.Println(values)
	return values
}

func (m *Model) SetFromByteArray(byteArray []interface{}) {
	// Check to see if the model's metadata has been set, and set it if not:
	if !m.Meta.isSet {
		m.SetMetaFields()
	}

	s := reflect.ValueOf(m).Elem()
	for i := 0; i < m.Meta.numFields; i++ {
		f := s.Field(i)
		t := f.Type()
		switch t.String() {
		// TODO: finish this switch with the rest of the supported types.
		case "string":
			f.SetString(fmt.Sprintf("%s", byteArray[i]))
		case "int":
			f.SetInt(byteArray[i].(int64))
		case "int64":
			f.SetInt(byteArray[i].(int64))
		case "bool":
			f.SetBool(byteArray[i].(bool))
			// default:
			//     some default case here...
		}
	}
}
