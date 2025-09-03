package adifield

// ADIField is the ADIF field name in and ADI file.
// By convention in the adifield package, field name constants are always UPPERCASE.
// This is a departure from the ADIF specification, which allows field names to include lowercase letters.
type ADIField string

func (f ADIField) String() string {
	return string(f)
}
