package main

import "errors"

type Usuario struct {
	nickname  string
	email     string
	firstname string
	lastname  string
	password  string
	country   string
}

func newUsuario(u Usuario) error{
	q :=`INSERT INTO RPC
		(nickname, email, firstname, lastname, password, country) 
		VALUES ($1,$2,$3,$4,$5,$6)`
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil{
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(u.nickname, u.email, u.firstname, u.lastname, u.password, u.country)
	if err != nil{
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Affected row expected")
	}
	return nil
}
