package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectaBancoPostgres() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=fabio1012 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error)
	} else {
		fmt.Printf("Conectado ao PostgreSQL \n")
	}
	return db
}
