// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package qslmedium

var (
	CARD QSLMedium = "CARD" // QSO confirmation via paper QSL card
	EQSL QSLMedium = "EQSL" // QSO confirmation via eQSL.cc
	LOTW QSLMedium = "LOTW" // QSO confirmation via ARRL Logbook of the World
)

// IsValid returns true if the QSLMedium exists in the ADIF specification. This includes Import Only and Deleted values.
func (value QSLMedium) IsValid() bool {
	switch value {
	case CARD:
		return true
	case EQSL:
		return true
	case LOTW:
		return true
	}
	return false
}
