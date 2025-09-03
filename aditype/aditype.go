package aditype

// ADIType represents the ADIF data type of a data field
type ADIType string

func (t ADIType) String() string {
	return string(t)
}
