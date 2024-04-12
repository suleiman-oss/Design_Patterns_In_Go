package database

import "sync"

type database struct {
	username string
	password string
	host     string
}

// Without using singleton
func NewDatabase() *database {
	return &database{username: "postgres", password: "psotgres", host: ":8080"}
}

// With singleton
var Db *database
var once sync.Once
var mu sync.Mutex

func GetNewDatabase() *database {
	mu.Lock()
	defer mu.Unlock()
	once.Do(func() {
		Db = &database{username: "suleiman", password: "suleima", host: ":8080"}
	})
	return Db
}
