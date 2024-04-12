package main

import service "github.com/suleiman-oss/Design_Patterns_in_Go/Creational/Factory/Service"

func main() {
	//Using MySQL DB
	m := &service.MySQL{}
	u := service.UserService{Db: m}
	u.Add()
	//Using PostgreSQL DB
	p := &service.PostgreSQL{}
	up := service.UserService{Db: p}
	up.Add()
}
