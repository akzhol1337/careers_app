package model

import "database/sql"

func GetConnection(link string, port string, dbname string, login string, password string) *sql.DB{
	db, err := sql.Open("mysql", login + ":" + password + "@tcp(" + link+ ":" + port + ")/" + dbname)
	if err != nil{
		panic(err)
	}
	return db
}
