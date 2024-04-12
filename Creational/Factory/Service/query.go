package service

import "fmt"

type DBQuery interface {
	ExecuteQuery()
}
type MySQLQuery struct{}

func (mq *MySQLQuery) ExecuteQuery() {
	fmt.Println("Executing MySQL Query")
}

type PostgreSQLQuery struct{}

func (pq *PostgreSQLQuery) ExecuteQuery() {
	fmt.Println("Executing PostgreSQL Query")
}
