// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package propagationmode

var (
	AS       PropagationMode = "AS"       // Aircraft Scatter
	AUE      PropagationMode = "AUE"      // Aurora-E
	AUR      PropagationMode = "AUR"      // Aurora
	BS       PropagationMode = "BS"       // Back scatter
	ECH      PropagationMode = "ECH"      // EchoLink
	EME      PropagationMode = "EME"      // Earth-Moon-Earth
	ES       PropagationMode = "ES"       // Sporadic E
	F2       PropagationMode = "F2"       // F2 Reflection
	FAI      PropagationMode = "FAI"      // Field Aligned Irregularities
	GWAVE    PropagationMode = "GWAVE"    // Ground Wave
	INTERNET PropagationMode = "INTERNET" // Internet-assisted
	ION      PropagationMode = "ION"      // Ionoscatter
	IRL      PropagationMode = "IRL"      // IRLP
	LOS      PropagationMode = "LOS"      // Line of Sight (includes transmission through obstacles such as walls)
	MS       PropagationMode = "MS"       // Meteor scatter
	RPT      PropagationMode = "RPT"      // Terrestrial or atmospheric repeater or transponder
	RS       PropagationMode = "RS"       // Rain scatter
	SAT      PropagationMode = "SAT"      // Satellite
	TEP      PropagationMode = "TEP"      // Trans-equatorial
	TR       PropagationMode = "TR"       // Tropospheric ducting
)

// IsValid returns true if the PropagationMode exists in the ADIF specification. This includes Import Only and Deleted values.
func (value PropagationMode) IsValid() bool {
	switch value {
	case AS:
		return true
	case AUE:
		return true
	case AUR:
		return true
	case BS:
		return true
	case ECH:
		return true
	case EME:
		return true
	case ES:
		return true
	case F2:
		return true
	case FAI:
		return true
	case GWAVE:
		return true
	case INTERNET:
		return true
	case ION:
		return true
	case IRL:
		return true
	case LOS:
		return true
	case MS:
		return true
	case RPT:
		return true
	case RS:
		return true
	case SAT:
		return true
	case TEP:
		return true
	case TR:
		return true
	}
	return false
}
