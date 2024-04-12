package main

import (
	"fmt"

	db "github.com/suleiman-oss/Design_Patterns_in_Go/Creational/Singleton/database"
)

func main() {
	db1 := db.NewDatabase()
	db2 := db.NewDatabase()
	fmt.Println(db1 == db2) //false

	db1 = db.GetNewDatabase()
	db2 = db.GetNewDatabase()
	fmt.Println(db1 == db2) //true
}
