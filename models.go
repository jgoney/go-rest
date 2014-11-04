package main

import (
	"fmt"
	"net/url"
	"strings"
)

type model struct {
	id int
	firstname string
	lastname string
	email string
	gender string
}

// Sets fields on a model from POST URL values
func (m *model) setFields(urlv url.Values) {
	m.firstname = strings.Join(urlv["firstname"], " ")
	m.lastname = strings.Join(urlv["lastname"], " ")
	m.email = strings.Join(urlv["email"], " ")
	m.gender = strings.Join(urlv["gender"], " ")
}


func modelPrint() {
	fmt.Println("This is a model.")
}