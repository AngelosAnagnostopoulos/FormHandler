package routes

import (
	vars "angelos/goserver/globals"
	"angelos/goserver/orm"
	"fmt"
	"log"
	"net/http"
)

/*Reads an html form for data and updates a person type*/
func readForm(r *http.Request, person *orm.Person) {
	person.ID = vars.ID + 1
	person.Name = r.FormValue("name")
	person.Email = r.FormValue("email")
	person.City = r.FormValue("city")
	person.Phone = r.FormValue("phone")

	vars.Persons = append(vars.Persons, *person)
	vars.Flag = true
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	var person orm.Person

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	readForm(r, &person)
	http.Redirect(w, r, "/#", http.StatusFound)
}

/*Make a simple http server to handle form submission*/
func ServerSetup(port int) {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port %d\n", port)
	con := fmt.Sprintf(":%d", port)

	go func() {
		if err := http.ListenAndServe(con, nil); err != nil {
			log.Fatal(err)
		}
	}()
}
