package main

import (
	database "angelos/goserver/backendHandler"
	vars "angelos/goserver/globals"
	routes "angelos/goserver/routehandling"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	vars.Wg.Add(2)

	db := database.ConnectToDb()
	routes.ServerSetup(8080)

	// Implement an: event.OnClick("/form.html", "submit"){}
	// This only runs once but it works for now.
	for {
		if vars.Flag {
			database.PutData(db, &vars.Persons[vars.Rcount])
			database.GetData(db)
			vars.Rcount++
			break
		}
	}

	vars.Wg.Wait()
}
