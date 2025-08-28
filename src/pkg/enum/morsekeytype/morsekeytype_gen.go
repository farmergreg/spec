// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package morsekeytype

var (
	BUG MorseKeyType = "BUG" // Mechanical semi-automatic keyer or Bug
	CPU MorseKeyType = "CPU" // Computer Driven
	DP  MorseKeyType = "DP"  // Dual Paddle
	FAB MorseKeyType = "FAB" // Mechanical fully-automatic keyer or Bug
	SK  MorseKeyType = "SK"  // Straight Key
	SP  MorseKeyType = "SP"  // Single Paddle
	SS  MorseKeyType = "SS"  // Sideswiper
)

// IsValid returns true if the MorseKeyType exists in the ADIF specification. This includes Import Only and Deleted values.
func (value MorseKeyType) IsValid() bool {
	switch value {
	case BUG:
		return true
	case CPU:
		return true
	case DP:
		return true
	case FAB:
		return true
	case SK:
		return true
	case SP:
		return true
	case SS:
		return true
	}
	return false
}
