package mssql

import (
	"testing"
)

func TestConnect(t *testing.T) {

	db := Connect()
	defer db.Close()

	if err := db.Ping(); err != nil {

		t.Errorf(err.Error())
	}
	return
}
