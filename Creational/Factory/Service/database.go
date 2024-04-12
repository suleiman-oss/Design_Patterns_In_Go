package service

import "fmt"

type Database interface {
	Connect()
	Query() DBQuery
}

type MySQL struct{}

func (m *MySQL) Connect() {
	fmt.Println("Connected to MySQL DB")
}

func (m *MySQL) Query() DBQuery {
	return &MySQLQuery{}
}

type PostgreSQL struct{}

func (p *PostgreSQL) Connect() {
	fmt.Println("Connected to PostgreSQL DB")
}

func (p *PostgreSQL) Query() DBQuery {
	return &PostgreSQLQuery{}
}
