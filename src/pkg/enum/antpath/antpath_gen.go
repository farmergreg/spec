// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package antpath

var (
	G AntPath = "G" // grayline
	L AntPath = "L" // long path
	O AntPath = "O" // other
	S AntPath = "S" // short path
)

// IsValid returns true if the AntPath exists in the ADIF specification. This includes Import Only and Deleted values.
func (value AntPath) IsValid() bool {
	switch value {
	case G:
		return true
	case L:
		return true
	case O:
		return true
	case S:
		return true
	}
	return false
}
