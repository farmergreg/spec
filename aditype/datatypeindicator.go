package aditype

import (
	"encoding/json"
	"errors"
	"unicode"
)

// DataTypeIndicator is a single character that represents the ADIF Data Type Indicator that precedes the data field in an ADI record.
type DataTypeIndicator rune

func (t *DataTypeIndicator) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	for _, r := range val {
		*t = DataTypeIndicator(unicode.ToUpper(r))
		return nil
	}

	return errors.New("invalid DataTypeIndicator: empty string")
}

func NewDataTypeIndicator(value rune) DataTypeIndicator {
	return DataTypeIndicator(unicode.ToUpper(value))
}

func (t DataTypeIndicator) String() string {
	return string(t)
}

func (t DataTypeIndicator) Compare(other DataTypeIndicator) int {
	return int(unicode.ToUpper(rune(t))) - int(unicode.ToUpper(rune(other)))
}

func (t DataTypeIndicator) Equals(other DataTypeIndicator) bool {
	return unicode.ToUpper(rune(t)) == unicode.ToUpper(rune(other))
}
