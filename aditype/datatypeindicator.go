package aditype

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

// DataTypeIndicator is a single character that represents the ADIF Data Type Indicator that precedes the data field in an ADI record.
type DataTypeIndicator rune

func (t *DataTypeIndicator) UnmarshalJSON(p []byte) error {
	r, c := utf8.DecodeRune(p)
	if r == utf8.RuneError && c == 0 {
		return errors.New("invalid rune")
	}
	*t = DataTypeIndicator(r)
	return nil
}

func (t DataTypeIndicator) String() string {
	return string(t)
}

func (t DataTypeIndicator) Compare(other DataTypeIndicator) int {
	return int(t) - int(other)
}

func (t DataTypeIndicator) Equals(other DataTypeIndicator) bool {
	return unicode.ToUpper(rune(t)) == unicode.ToUpper(rune(other))
}
