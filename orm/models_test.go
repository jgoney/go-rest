package orm

import "reflect"
import "testing"

func SetUp() *Model {
	myModel := MyModel{Id: 0,
		Firstname: "Justin",
		Lastname:  "Goney",
		Email:     "goulash@gmail.com",
		Gender:    "Male"}

	return NewModel(myModel)
}

func TestCreateTable(t *testing.T) {

	m := SetUp()

	command := m.GenCreateTable()
	expectedCommand := "create table MyModels (Id integer, Firstname text, Lastname text, Email text, Gender text);"

	t.Logf("Expected: %s", expectedCommand)
	t.Logf("Obtained: %s", command)

	if command != expectedCommand {
		t.Error("Output string didn't match expected string.")
	}
}

func TestString(t *testing.T) {
	m := SetUp()

	command := m.String()
	expectedCommand := "{0 Justin Goney goulash@gmail.com Male} : {MyModel 5 [[Id integer] [Firstname text] [Lastname text] [Email text] [Gender text]] true}"

	t.Logf("Expected: %s", expectedCommand)
	t.Logf("Obtained: %s", command)

	if command != expectedCommand {
		t.Error("Output string didn't match expected string.")
	}
}

func TestInsertInto(t *testing.T) {
	m := SetUp()

	command := m.GenInsertInto()
	expectedCommand := "insert into MyModels(Id, Firstname, Lastname, Email, Gender) values(?, ?, ?, ?, ?);"

	t.Logf("Expected: %s", expectedCommand)
	t.Logf("Obtained: %s", command)

	if command != expectedCommand {
		t.Error("Output string didn't match expected string.")
	}
}

func TestValueString(t *testing.T) {
	m := SetUp()

	command := m.GenValueString()

	for i := range command {
		t.Log(command[i])
	}
}

// Tests that metadata fields are correctly set by constructor.
func TestMetaIsSet(t *testing.T) {

	m := SetUp()

	expectedModelName := "MyModel"
	if m.Meta.modelName != expectedModelName {
		t.Errorf("Meta.modelName: %s didn't match %s", m.Meta.modelName, expectedModelName)
	}

	expectedNumFields := 5
	if m.Meta.numFields != expectedNumFields {
		t.Errorf("Meta.numFields: %d didn't match %d", m.Meta.numFields, expectedNumFields)
	}

	for i := range m.Meta.elements {

		typeName := reflect.TypeOf(m.Modeler).Field(i).Name
		if m.Meta.elements[i][0] != typeName {
			t.Errorf("Meta.elements: %s didn't match %s", m.Meta.elements[i][0], typeName)
		}

		typeMapping := typeMappings[reflect.TypeOf(m.Modeler).Field(i).Type.Name()]
		if m.Meta.elements[i][1] != typeMapping {
			t.Errorf("Meta.elements: %s didn't match %s", m.Meta.elements[i][1], typeMapping)
		}
	}

	if !m.Meta.isSet {
		t.Error("Meta.isSet was not correctly set")
	}
}
