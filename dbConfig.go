package main

import "fmt"

type DBConfig struct {
	driver, user, password, dbname, host, port, sslmode string
}

func (dbc DBConfig) GetConfigString() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbc.user, dbc.password, dbc.dbname, dbc.host, dbc.port, dbc.sslmode,
	)
}

var dbConfig = DBConfig{
	driver:   "postgres",
	user:     "darv",
	password: "groovy",
	dbname:   "goserv",
	host:     "localhost",
	port:     "5432",
	sslmode:  "disable",
}
