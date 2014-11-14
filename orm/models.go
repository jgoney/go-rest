package orm

// ExampleModel represents an example Modeler type to be associated with a Model object.
// Basically, user data should be defined by defining custom structs here, which are then passed to the NewModel constructor.
type ExampleModel struct {
	Id        int
	Firstname string
	Lastname  string
	Email     string
	Gender    string
}

// See ExampleModel for more information.
type AnotherModel struct {
	Id  int
	Fee string
	Fi  string
	Fo  string
	Fum float64
}
