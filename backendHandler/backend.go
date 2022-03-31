package database

import (
	"angelos/goserver/orm"
	"database/sql"
	"fmt"
	"log"
)

const (
	username = "docker"
	password = "1234"
	host     = "project_database"
	port     = "3306"
	db_name  = "golang_test"
)


/*Make initial connection to backend*/
func ConnectToDb() *sql.DB {
	fmt.Println("Opening database connection")
	info := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, db_name)

	db, err := sql.Open("mysql", info)
	if err != nil {
		panic("Couldnt open database for connection.")
	}

	return db
}

/*Get table info*/
func GetData(db *sql.DB) {
	defer recoverBackend()
	var person orm.Person

	rows, err := db.Query("select * from Person")
	if err != nil {
		panic("Couldnt perform operation 'SELECT' on database tables.")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&person.ID, &person.Name, &person.Email, &person.Phone, &person.City)
		if err != nil {
			panic("Something went wrong while scanning the table's rows.")
		}
		log.Println(person.ID, person.Name, person.Email, person.Phone, person.City)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func recoverBackend() {
	if r := recover(); r != nil {
		fmt.Println("Phew... Recovered from:", r)
	}
}

/*Puts Person's details to backend database*/
func PutData(db *sql.DB, person *orm.Person) {
	defer recoverBackend()
	stmt, err := db.Prepare("INSERT INTO Person(fname,email,phone,city) VALUES(?,?,?,?)")
	if err != nil {
		panic("Error with opening database. Consider creating desired database w/ 'CREATE DATABASE <db_name> and rerun the program.")
	}
	res, err := stmt.Exec(person.Name, person.Email, person.Phone, person.City)
	if err != nil {
		panic("Error with executing SQL commands.")
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}
