// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package qslvia

var (
	B QSLVia = "B" // bureau
	D QSLVia = "D" // direct
	E QSLVia = "E" // electronic
	M QSLVia = "M" // Deprecated: manager
)

// IsValid returns true if the QSLVia exists in the ADIF specification. This includes Import Only and Deleted values.
func (value QSLVia) IsValid() bool {
	switch value {
	case B:
		return true
	case D:
		return true
	case E:
		return true
	case M:
		return true
	}
	return false
}
