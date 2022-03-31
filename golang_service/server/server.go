package main

import (
	routes "angelos/goserver/routehandling"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

func main() {
	Wg := new(sync.WaitGroup)
	Wg.Add(2)
	routes.ServerSetup(8080)
	Wg.Wait()
}
