package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/jgoney/go-rest/orm"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Method+"\n")
	fmt.Fprint(w, r.URL)
}

// Utility function to view available header members
func enumHeader(w *http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
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

		model := orm.MyModel{}
		m := orm.NewModel(model)
		//model.SetFieldsFromPOST(r.PostForm)

		orm.InsertDB(m)
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

	myModel := orm.MyModel{Firstname: "Justin",
		Lastname: "Goney",
		Email:    "goulash@gmail.com",
		Gender:   "Male"}

	m := orm.NewModel(myModel)

	a := []*orm.Model{orm.NewModel(orm.MyModel{Firstname: "Bernice", Lastname: "Smith", Email: "someone@gmail.com", Gender: "Female"}),
		orm.NewModel(orm.MyModel{Firstname: "McLovin", Lastname: "", Email: "mclovin@gmail.com", Gender: "Male"}),
	}

	aModel := orm.AnotherModel{Fee: "Fee",
		Fi: "Fi",
		Fo:    "Fo",
		Fum:   3.14}

	ma := orm.NewModel(aModel)

	// Create and initialize DB only if it doesn't exist
	if _, err := os.Stat(orm.DB_NAME); err != nil {
		orm.InitDB(m, ma)
	}

	// Insert MyModel and array of MyModels
	orm.InsertDB(m)
	orm.InsertDB(a...)

	// Insert AnotherModel
	orm.InsertDB(ma)

	list := orm.GetResultsDB(m)
	for _, v := range list {
		fmt.Println(v)
	}

	// http.HandleFunc("/create", createEntry)
	// http.HandleFunc("/view/", viewEntry)
	// http.HandleFunc("/", siteRoot)
	// http.ListenAndServe("localhost:4000", nil)
}
