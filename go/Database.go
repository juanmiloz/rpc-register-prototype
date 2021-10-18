package main

import(
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "rpc"
)

//get connection obtiene la conexión con la base de datos
func getConnection() * sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		log.Fatal(err)
	}
	return db
}