/*
   En este documento se escribe la función necesaria para establecer una
   conexión con la base de datos MSSQL,
*/
package mssql

import (
	// Se importa en blanco para el uso como driver
	_ "github.com/denisenkom/go-mssqldb"

	"context"
	"database/sql"
	"log"
	"os"
)

// connString - Variable para guardar la cadena de conexión con la BD
var connString = os.Getenv("MSSQL_STRING")

var (
	ctx context.Context
	db  *sql.DB
)

// Connect - Se encarga de realizar la conexión con SQL, si no existe alguna ya
// iniciada, crea otra conexión.
// Cada interacción con MSSQL deberá llamar primero a esta función para aegurar
// que la interación con la BD sea exitosa.
func Connect() *sql.DB {

	if db == nil {
		open("db nil")
	}

	if err := db.Ping(); err != nil {
		open("Ping failed")
	}

	return db
}

// open establece la conexión con MSSQL
func open(trigger string) {

	log.Printf("****** Generando una nueva conexión a MSSQL - %s ******\n", trigger)
	var err error
	db, err = sql.Open("mssql", connString)
	if err != nil {
		log.Println("****** Error generado al conectar a MSSQL ******")
		log.Println(err.Error())
		log.Println("------ Error generado al conectar a MSSQL ------")
		return
	}
}
