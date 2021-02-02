package mssql

import (
	mssql "github.com/denisenkom/go-mssqldb"
)

// NullUniqueIdentifier es una extend del paquete de denisenkom para que esté
// disponible sin reimportar el pquete
type NullUniqueIdentifier struct {
	UniqueIdentifier mssql.UniqueIdentifier
	Valid            bool
}

// Scan Implementación del método scan
func (nui *NullUniqueIdentifier) Scan(v interface{}) error {

	if v == nil {

		nui.Valid = false
		return nil
	}

	nui.Valid = true
	return nui.UniqueIdentifier.Scan(v)
}
