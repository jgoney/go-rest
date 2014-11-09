package orm

// import (
// 	"fmt"
// 	"net/url"
// 	"reflect"
// 	"strings"
// )

type MyModel struct {
	Id        int
	Firstname string
	Lastname  string
	Email     string
	Gender    string
}

// // TODO: extract these functions out into prototype functions
// func (m MyModel) String() string {
// 	// Check to see if the model's metadata has been set, and set it if not:
// 	if !m.Meta.isSet {
// 		m.SetMetaFields()
// 	}
// 	return fmt.Sprintf("%v: %v, %v, %v, %v, %v", m.Meta.modelName, m.Id, m.Firstname, m.Lastname, m.Email, m.Gender)
// }

// // Sets metadata fields via reflection. Used to convert Go types to SQL syntax
// func (m MyModel) SetMetaFields() {
// 	// Check to see if the model's metadata has already been set, and set it if not:
// 	if m.Meta.isSet {
// 		return
// 	}

// 	m.Meta.modelName = reflect.TypeOf(m).Name()
// 	m.Meta.numFields = reflect.TypeOf(m).NumField() - 1 // Minus 1 since we're ignoring the metadata struct
// 	m.Meta.elements = make([][]string, m.Meta.numFields)

// 	for i := 0; i < m.Meta.numFields; i++ {
// 		m.Meta.elements[i] = []string{reflect.TypeOf(m).Field(i).Name, typeMappings[reflect.TypeOf(m).Field(i).Type.Name()]}
// 	}

// 	m.Meta.isSet = true
// }

// // Sets fields on a model from POST URL values
// func (m MyModel) SetFieldsFromPOST(urlv url.Values) {
// 	m.Firstname = strings.Join(urlv["Firstname"], " ")
// 	m.Lastname = strings.Join(urlv["Lastname"], " ")
// 	m.Email = strings.Join(urlv["Email"], " ")
// 	m.Gender = strings.Join(urlv["Gender"], " ")
// }

// // Returns a string used to create a table representing
// func (m MyModel) GenCreateTable() string {

// 	// Check to see if the model's metadata has been set, and set it if not:
// 	if !m.Meta.isSet {
// 		m.SetMetaFields()
// 	}

// 	// "create table names (id integer not null primary key autoincrement, first_name text, last_name text, email text, gender text);"
// 	modelNames := m.Meta.modelName + "s"
// 	fields := make([]string, m.Meta.numFields)

// 	for i, v := range m.Meta.elements {
// 		fields[i] = strings.Join(v, " ")
// 	}

// 	command := fmt.Sprintf("create table %s (%s);", modelNames, strings.Join(fields, ", "))
// 	return command
// }

// func (m MyModel) GenInsertInto() string {

// 	// Check to see if the model's metadata has been set, and set it if not:
// 	if !m.Meta.isSet {
// 		m.SetMetaFields()
// 	}

// 	// "insert into names(first_name, last_name, email, gender) values(?, ?, ?, ?)"
// 	modelNames := m.Meta.modelName + "s"
// 	fields := make([]string, m.Meta.numFields)
// 	placeholders := make([]string, m.Meta.numFields)

// 	for i, v := range m.Meta.elements {
// 		fields[i] = v[0]
// 		placeholders[i] = "?"
// 	}

// 	command := fmt.Sprintf("insert into %s(%s) values(%s);", modelNames, strings.Join(fields, ", "), strings.Join(placeholders, ", "))
// 	return command
// }

// func (m MyModel) GenValueString() []interface{} {
// 	// Check to see if the model's metadata has been set, and set it if not:
// 	if !m.Meta.isSet {
// 		m.SetMetaFields()
// 	}
// 	values := make([]interface{}, m.Meta.numFields)

// 	s := reflect.ValueOf(m).Elem()
// 	for i := range m.Meta.elements {
// 		values[i] = fmt.Sprintf("%v", s.Field(i).Interface())
// 	}

// 	return values
// }

// func (m MyModel) ExposeFields() []interface{} {
// 	// Check to see if the model's metadata has been set, and set it if not:
// 	if !m.Meta.isSet {
// 		m.SetMetaFields()
// 	}
// 	values := make([]interface{}, m.Meta.numFields)

// 	s := reflect.ValueOf(m).Elem()
// 	typeOfT := s.Type()
// 	for i := 0; i < m.Meta.numFields; i++ {
// 	    f := s.Field(i)
// 	    fmt.Printf("%d: %s %s = %v\n", i,
// 	        typeOfT.Field(i).Name, f.Type(), f.Interface())
// 	    values[i] = f.Interface()
// 	}
// 	fmt.Println(values)
// 	return values
// }

// func (m MyModel) SetFromByteArray(byteArray []interface{}) {
// 	// Check to see if the model's metadata has been set, and set it if not:
// 	if !m.Meta.isSet {
// 		m.SetMetaFields()
// 	}

// 	s := reflect.ValueOf(m).Elem()
// 	for i := 0; i < m.Meta.numFields; i++ {
// 	    f := s.Field(i)
// 	    t := f.Type()
//     	switch t.String() {
//     		// TODO: finish this switch with the rest of the supported types.
// 		    case "string":
// 		        f.SetString(fmt.Sprintf("%s", byteArray[i]))
// 		    case "int":
// 		        f.SetInt(byteArray[i].(int64))
// 		    case "int64":
// 		        f.SetInt(byteArray[i].(int64))
// 		    case "bool":
// 		        f.SetBool(byteArray[i].(bool))
// 		    // default:
// 		    //     some default case here...
// 	    }
// 	}
// }
