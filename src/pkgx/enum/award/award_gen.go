// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package award

var (
	AJA         Award = "AJA"         // Deprecated: AJA
	CQDX        Award = "CQDX"        // Deprecated: CQDX
	CQDXFIELD   Award = "CQDXFIELD"   // Deprecated: CQDXFIELD
	CQWAZ_160m  Award = "CQWAZ_160m"  // Deprecated: CQWAZ_160m
	CQWAZ_CW    Award = "CQWAZ_CW"    // Deprecated: CQWAZ_CW
	CQWAZ_MIXED Award = "CQWAZ_MIXED" // Deprecated: CQWAZ_MIXED
	CQWAZ_PHONE Award = "CQWAZ_PHONE" // Deprecated: CQWAZ_PHONE
	CQWAZ_RTTY  Award = "CQWAZ_RTTY"  // Deprecated: CQWAZ_RTTY
	CQWPX       Award = "CQWPX"       // Deprecated: CQWPX
	DARC_DOK    Award = "DARC_DOK"    // Deprecated: DARC_DOK
	DXCC        Award = "DXCC"        // Deprecated: DXCC
	DXCC_CW     Award = "DXCC_CW"     // Deprecated: DXCC_CW
	DXCC_MIXED  Award = "DXCC_MIXED"  // Deprecated: DXCC_MIXED
	DXCC_PHONE  Award = "DXCC_PHONE"  // Deprecated: DXCC_PHONE
	DXCC_RTTY   Award = "DXCC_RTTY"   // Deprecated: DXCC_RTTY
	IOTA        Award = "IOTA"        // Deprecated: IOTA
	JCC         Award = "JCC"         // Deprecated: JCC
	JCG         Award = "JCG"         // Deprecated: JCG
	MARATHON    Award = "MARATHON"    // Deprecated: MARATHON
	RDA         Award = "RDA"         // Deprecated: RDA
	USACA       Award = "USACA"       // Deprecated: USACA
	VUCC        Award = "VUCC"        // Deprecated: VUCC
	WAB         Award = "WAB"         // Deprecated: WAB
	WAC         Award = "WAC"         // Deprecated: WAC
	WAE         Award = "WAE"         // Deprecated: WAE
	WAIP        Award = "WAIP"        // Deprecated: WAIP
	WAJA        Award = "WAJA"        // Deprecated: WAJA
	WAS         Award = "WAS"         // Deprecated: WAS
	WAZ         Award = "WAZ"         // Deprecated: WAZ
)

// IsValid returns true if the Award exists in the ADIF specification. This includes Import Only and Deleted values.
func (value Award) IsValid() bool {
	switch value {
	case AJA:
		return true
	case CQDX:
		return true
	case CQDXFIELD:
		return true
	case CQWAZ_160m:
		return true
	case CQWAZ_CW:
		return true
	case CQWAZ_MIXED:
		return true
	case CQWAZ_PHONE:
		return true
	case CQWAZ_RTTY:
		return true
	case CQWPX:
		return true
	case DARC_DOK:
		return true
	case DXCC:
		return true
	case DXCC_CW:
		return true
	case DXCC_MIXED:
		return true
	case DXCC_PHONE:
		return true
	case DXCC_RTTY:
		return true
	case IOTA:
		return true
	case JCC:
		return true
	case JCG:
		return true
	case MARATHON:
		return true
	case RDA:
		return true
	case USACA:
		return true
	case VUCC:
		return true
	case WAB:
		return true
	case WAC:
		return true
	case WAE:
		return true
	case WAIP:
		return true
	case WAJA:
		return true
	case WAS:
		return true
	case WAZ:
		return true
	}
	return false
}
