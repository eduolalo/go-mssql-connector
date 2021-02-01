package mssql

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
)

// NullUniqueIdentifier es una extend del paquete de denisenkom para que esté
// disponible sin reimportar el pquete
type NullUniqueIdentifier struct {
	UniqueIdentifier [16]byte
	Valid            bool
}

// Scan Implementación del método scan
func (nui *NullUniqueIdentifier) Scan(v interface{}) error {

	reverse := func(b []byte) {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}

	switch vt := v.(type) {
	case []byte:
		if len(vt) != 16 {

			nui.Valid = false
			return nil
		}

		var raw [16]byte

		copy(raw[:], vt)

		reverse(raw[0:4])
		reverse(raw[4:6])
		reverse(raw[6:8])
		nui.UniqueIdentifier = raw

		return nil
	case string:
		if len(vt) != 36 {

			nui.Valid = false
			return nil
		}

		b := []byte(vt)
		for i, c := range b {
			switch c {
			case '-':
				b = append(b[:i], b[i+1:]...)
			}
		}

		_, err := hex.Decode(nui.UniqueIdentifier[:], []byte(b))
		return err
	default:
		return fmt.Errorf("mssql: cannot convert %T to UniqueIdentifier", v)
	}
}

func (nui NullUniqueIdentifier) Value() (driver.Value, error) {
	reverse := func(b []byte) {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}

	raw := make([]byte, len(nui.UniqueIdentifier))
	copy(raw, nui.UniqueIdentifier[:])

	reverse(raw[0:4])
	reverse(raw[4:6])
	reverse(raw[6:8])

	return raw, nil
}

func (nui NullUniqueIdentifier) String() string {
	return fmt.Sprintf("%X-%X-%X-%X-%X", nui.UniqueIdentifier[0:4], nui.UniqueIdentifier[4:6], nui.UniqueIdentifier[6:8], nui.UniqueIdentifier[8:10], nui.UniqueIdentifier[10:])
}

// MarshalText converts NullUniqueidentifier to bytes corresponding to the stringified hexadecimal representation of the NullUniqueidentifier
// e.g., "AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA" -> [65 65 65 65 65 65 65 65 45 65 65 65 65 45 65 65 65 65 45 65 65 65 65 65 65 65 65 65 65 65 65]
func (nui NullUniqueIdentifier) MarshalText() []byte {
	return []byte(nui.String())
}
