package vars

import (
	"angelos/goserver/orm"
	"sync"
)

var (
	Persons = []orm.Person{}
	ID      = -1
	Wg      = new(sync.WaitGroup)
	Flag    = false
	Rcount  = 0 //With hopes of making this a proper service
)
