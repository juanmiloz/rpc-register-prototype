package main

import(
	"database/sql"
	_"github.com/lib/pq"
	"log"
)

//get connection obtiene la conexión con la base de datos
func getConnection() * sql.DB{
	dsn := "postgres://postgres:DatabaseJCZ13*@127.0.0.1:5432/RPC?sslmode=disable"
	db, err := sql.Open("postgres",dsn)
	if err != nil{
		log.Fatal(err)
	}

	return db
}