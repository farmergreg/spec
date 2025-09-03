package aditype

// ADITypeIndicator is a single character that represents the ADIF Data Type Indicator that precedes the data field in an ADI record.
type ADITypeIndicator rune

// ADIType represents the ADIF data type of a data field
type ADIType string

func (t ADIType) String() string {
	return string(t)
}
