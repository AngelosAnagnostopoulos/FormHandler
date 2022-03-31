package routes

import (
	database "angelos/goserver/backendHandler"
	"angelos/goserver/orm"
	"fmt"
	"log"
	"net/http"
)

/*Reads an html form for data and updates a person type*/
func readForm(r *http.Request, person *orm.Person) {
	fmt.Println("Hello from readForm")
	person.Name = r.FormValue("name")
	person.Email = r.FormValue("email")
	person.City = r.FormValue("city")
	person.Phone = r.FormValue("phone")

	db := database.ConnectToDb()
	fmt.Println("Connecting to db from readForm")
	database.PutData(db, person)
	database.GetData(db)
	db.Close()
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
