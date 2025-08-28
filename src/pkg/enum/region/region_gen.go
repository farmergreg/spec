// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package region

var (
	AI_248 Region = "AI.248" // African Italy
	BI_259 Region = "BI.259" // Bear Island
	ET_390 Region = "ET.390" // European Turkey
	IV_206 Region = "IV.206" // ITU Vienna
	KO_0   Region = "KO.0"   // Kosovo
	KO_296 Region = "KO.296" // Kosovo
	KO_522 Region = "KO.522" // Kosovo
	NONE   Region = "NONE"   // Not within a WAE or CQ region that is within a DXCC entity
	SI_279 Region = "SI.279" // Shetland Islands
	SY_248 Region = "SY.248" // Sicily
)

// IsValid returns true if the Region exists in the ADIF specification. This includes Import Only and Deleted values.
func (value Region) IsValid() bool {
	switch value {
	case AI_248:
		return true
	case BI_259:
		return true
	case ET_390:
		return true
	case IV_206:
		return true
	case KO_0:
		return true
	case KO_296:
		return true
	case KO_522:
		return true
	case NONE:
		return true
	case SI_279:
		return true
	case SY_248:
		return true
	}
	return false
}
