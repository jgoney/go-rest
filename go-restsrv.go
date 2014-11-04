package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Method + "\n")
	fmt.Fprint(w, r.URL)
}

// Utility function to view available header members
func enumHeader(w *http.ResponseWriter, r *http.Request) {
	for k, v := range(r.Header) {
		fmt.Fprint(*w, k, v, "\n")
	}
}

func createEntry(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Creating an entry:\n\n")
	fmt.Fprint(w, "\n")

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
		   log.Fatal(err)
		}

		fmt.Fprint(w, r.PostForm)

		model := model{}
		model.setFields(r.PostForm)

		insertDB(model)
	}
}

func viewEntry(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path[len("/view/"):])
	fmt.Fprint(w, "Viewing entry:", r.URL.Path[len("/edit/"):])
}

func siteRoot(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadFile("html/index.html")
    if err != nil {
		log.Fatal(err)
    }

	w.Write(body)
}

func main() {
	initDB()

	m := model{firstname: "Justin", 
			   lastname: "Goney", 
			   email: "goulash@gmail.com", 
			   gender: "Male"}

	a := []model{model{firstname: "Bernice", lastname: "Smith", email: "someone@gmail.com", gender: "Female"}, 
		model{firstname: "McLovin", lastname: "", email: "mclovin@gmail.com", gender: "Male"}}

	insertDB(m)
	insertDB(a...)
	getResultsDB()

	http.HandleFunc("/create", createEntry)
	http.HandleFunc("/view/", viewEntry)
	http.HandleFunc("/", siteRoot)
	http.ListenAndServe("localhost:4000", nil)
}
