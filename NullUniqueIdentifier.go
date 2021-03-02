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

	if err := nui.UniqueIdentifier.Scan(v); err != nil {

		nui.Valid = false
		return nil
	}

	nui.Valid = true
	return nil
}

// String regresa el valor del unique identifier
func (nui NullUniqueIdentifier) String() string {

	if !nui.Valid {

		return ""
	}

	return nui.UniqueIdentifier.String()
}
