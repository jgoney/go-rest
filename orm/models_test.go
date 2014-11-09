package orm

import "testing"

func SetUp() *Model {
	myModel := MyModel{Firstname: "Justin",
		Lastname: "Goney",
		Email:    "goulash@gmail.com",
		Gender:   "Male"}

	return NewModel(myModel)
}

func TestCreateTable(t *testing.T) {

	m := SetUp()

	command := m.GenCreateTable()
	expected_command := "create table MyModels (Id integer, Firstname text, Lastname text, Email text, Gender text);"

	t.Logf("Expected: %s", expected_command)
	t.Logf("Obtained: %s", command)

	if command != expected_command {
		t.Error("Output string didn't match expected string.")
	}
}

func TestString(t *testing.T) {
	m := SetUp()

	command := m.String()
	expected_command := "{0 Justin Goney goulash@gmail.com Male} : {MyModel 5 [[Id integer] [Firstname text] [Lastname text] [Email text] [Gender text]] true}"

	t.Logf("Expected: %s", expected_command)
	t.Logf("Obtained: %s", command)

	if command != expected_command {
		t.Error("Output string didn't match expected string.")
	}
}
