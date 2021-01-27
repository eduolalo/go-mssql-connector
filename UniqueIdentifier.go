package mssql

import mssqldb "github.com/denisenkom/go-mssqldb"

// UniqueIdentifier es una extend del paquete de denisenkompara que esté
// disponible sin reimportar el pquete
type UniqueIdentifier struct {
	mssqldb.UniqueIdentifier
}
