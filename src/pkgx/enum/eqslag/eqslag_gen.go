// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package eqslag

var (
	N EQSL_AG = "N" // the QSO is confirmed but not "Authenticity Guaranteed" by eQSL
	U EQSL_AG = "U" // unknown
	Y EQSL_AG = "Y" // the QSO is confirmed and "Authenticity Guaranteed" by eQSL
)

// IsValid returns true if the EQSL_AG exists in the ADIF specification. This includes Import Only and Deleted values.
func (value EQSL_AG) IsValid() bool {
	switch value {
	case N:
		return true
	case U:
		return true
	case Y:
		return true
	}
	return false
}
