// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package qsocomplete

var (
	Uncertain QSOComplete = "?"   // uncertain
	N         QSOComplete = "N"   // no
	NIL       QSOComplete = "NIL" // not heard
	Y         QSOComplete = "Y"   // yes
)

// IsValid returns true if the QSOComplete exists in the ADIF specification. This includes Import Only and Deleted values.
func (value QSOComplete) IsValid() bool {
	switch value {
	case Uncertain:
		return true
	case N:
		return true
	case NIL:
		return true
	case Y:
		return true
	}
	return false
}
