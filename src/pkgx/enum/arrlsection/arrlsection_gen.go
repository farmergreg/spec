// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package arrlsection

var (
	AB  ARRLSection = "AB"  // Alberta
	AK  ARRLSection = "AK"  // Alaska
	AL  ARRLSection = "AL"  // Alabama
	AR  ARRLSection = "AR"  // Arkansas
	AZ  ARRLSection = "AZ"  // Arizona
	BC  ARRLSection = "BC"  // British Columbia
	CO  ARRLSection = "CO"  // Colorado
	CT  ARRLSection = "CT"  // Connecticut
	DE  ARRLSection = "DE"  // Delaware
	EB  ARRLSection = "EB"  // East Bay
	EMA ARRLSection = "EMA" // Eastern Massachusetts
	ENY ARRLSection = "ENY" // Eastern New York
	EPA ARRLSection = "EPA" // Eastern Pennsylvania
	EWA ARRLSection = "EWA" // Eastern Washington
	GA  ARRLSection = "GA"  // Georgia
	GH  ARRLSection = "GH"  // Golden Horseshoe
	GTA ARRLSection = "GTA" // Greater Toronto Area
	IA  ARRLSection = "IA"  // Iowa
	ID  ARRLSection = "ID"  // Idaho
	IL  ARRLSection = "IL"  // Illinois
	IN  ARRLSection = "IN"  // Indiana
	KS  ARRLSection = "KS"  // Kansas
	KY  ARRLSection = "KY"  // Kentucky
	LA  ARRLSection = "LA"  // Louisiana
	LAX ARRLSection = "LAX" // Los Angeles
	MAR ARRLSection = "MAR" // Maritime
	MB  ARRLSection = "MB"  // Manitoba
	MDC ARRLSection = "MDC" // Maryland-DC
	ME  ARRLSection = "ME"  // Maine
	MI  ARRLSection = "MI"  // Michigan
	MN  ARRLSection = "MN"  // Minnesota
	MO  ARRLSection = "MO"  // Missouri
	MS  ARRLSection = "MS"  // Mississippi
	MT  ARRLSection = "MT"  // Montana
	NB  ARRLSection = "NB"  // New Brunswick
	NC  ARRLSection = "NC"  // North Carolina
	ND  ARRLSection = "ND"  // North Dakota
	NE  ARRLSection = "NE"  // Nebraska
	NFL ARRLSection = "NFL" // Northern Florida
	NH  ARRLSection = "NH"  // New Hampshire
	NL  ARRLSection = "NL"  // Newfoundland/Labrador
	NLI ARRLSection = "NLI" // New York City-Long Island
	NM  ARRLSection = "NM"  // New Mexico
	NNJ ARRLSection = "NNJ" // Northern New Jersey
	NNY ARRLSection = "NNY" // Northern New York
	NS  ARRLSection = "NS"  // Nova Scotia
	NT  ARRLSection = "NT"  // Northwest Territories/Yukon/Nunavut
	NTX ARRLSection = "NTX" // North Texas
	NV  ARRLSection = "NV"  // Nevada
	NWT ARRLSection = "NWT" // Northwest Territories/Yukon/Nunavut
	OH  ARRLSection = "OH"  // Ohio
	OK  ARRLSection = "OK"  // Oklahoma
	ON  ARRLSection = "ON"  // Ontario
	ONE ARRLSection = "ONE" // Ontario East
	ONN ARRLSection = "ONN" // Ontario North
	ONS ARRLSection = "ONS" // Ontario South
	OR  ARRLSection = "OR"  // Oregon
	ORG ARRLSection = "ORG" // Orange
	PAC ARRLSection = "PAC" // Pacific
	PE  ARRLSection = "PE"  // Prince Edward Island
	PR  ARRLSection = "PR"  // Puerto Rico
	QC  ARRLSection = "QC"  // Quebec
	RI  ARRLSection = "RI"  // Rhode Island
	SB  ARRLSection = "SB"  // Santa Barbara
	SC  ARRLSection = "SC"  // South Carolina
	SCV ARRLSection = "SCV" // Santa Clara Valley
	SD  ARRLSection = "SD"  // South Dakota
	SDG ARRLSection = "SDG" // San Diego
	SF  ARRLSection = "SF"  // San Francisco
	SFL ARRLSection = "SFL" // Southern Florida
	SJV ARRLSection = "SJV" // San Joaquin Valley
	SK  ARRLSection = "SK"  // Saskatchewan
	SNJ ARRLSection = "SNJ" // Southern New Jersey
	STX ARRLSection = "STX" // South Texas
	SV  ARRLSection = "SV"  // Sacramento Valley
	TER ARRLSection = "TER" // Territories
	TN  ARRLSection = "TN"  // Tennessee
	UT  ARRLSection = "UT"  // Utah
	VA  ARRLSection = "VA"  // Virginia
	VI  ARRLSection = "VI"  // US Virgin Islands
	VT  ARRLSection = "VT"  // Vermont
	WCF ARRLSection = "WCF" // West Central Florida
	WI  ARRLSection = "WI"  // Wisconsin
	WMA ARRLSection = "WMA" // Western Massachusetts
	WNY ARRLSection = "WNY" // Western New York
	WPA ARRLSection = "WPA" // Western Pennsylvania
	WTX ARRLSection = "WTX" // West Texas
	WV  ARRLSection = "WV"  // West Virginia
	WWA ARRLSection = "WWA" // Western Washington
	WY  ARRLSection = "WY"  // Wyoming
)

// IsValid returns true if the ARRLSection exists in the ADIF specification. This includes Import Only and Deleted values.
func (value ARRLSection) IsValid() bool {
	switch value {
	case AB:
		return true
	case AK:
		return true
	case AL:
		return true
	case AR:
		return true
	case AZ:
		return true
	case BC:
		return true
	case CO:
		return true
	case CT:
		return true
	case DE:
		return true
	case EB:
		return true
	case EMA:
		return true
	case ENY:
		return true
	case EPA:
		return true
	case EWA:
		return true
	case GA:
		return true
	case GH:
		return true
	case GTA:
		return true
	case IA:
		return true
	case ID:
		return true
	case IL:
		return true
	case IN:
		return true
	case KS:
		return true
	case KY:
		return true
	case LA:
		return true
	case LAX:
		return true
	case MAR:
		return true
	case MB:
		return true
	case MDC:
		return true
	case ME:
		return true
	case MI:
		return true
	case MN:
		return true
	case MO:
		return true
	case MS:
		return true
	case MT:
		return true
	case NB:
		return true
	case NC:
		return true
	case ND:
		return true
	case NE:
		return true
	case NFL:
		return true
	case NH:
		return true
	case NL:
		return true
	case NLI:
		return true
	case NM:
		return true
	case NNJ:
		return true
	case NNY:
		return true
	case NS:
		return true
	case NT:
		return true
	case NTX:
		return true
	case NV:
		return true
	case NWT:
		return true
	case OH:
		return true
	case OK:
		return true
	case ON:
		return true
	case ONE:
		return true
	case ONN:
		return true
	case ONS:
		return true
	case OR:
		return true
	case ORG:
		return true
	case PAC:
		return true
	case PE:
		return true
	case PR:
		return true
	case QC:
		return true
	case RI:
		return true
	case SB:
		return true
	case SC:
		return true
	case SCV:
		return true
	case SD:
		return true
	case SDG:
		return true
	case SF:
		return true
	case SFL:
		return true
	case SJV:
		return true
	case SK:
		return true
	case SNJ:
		return true
	case STX:
		return true
	case SV:
		return true
	case TER:
		return true
	case TN:
		return true
	case UT:
		return true
	case VA:
		return true
	case VI:
		return true
	case VT:
		return true
	case WCF:
		return true
	case WI:
		return true
	case WMA:
		return true
	case WNY:
		return true
	case WPA:
		return true
	case WTX:
		return true
	case WV:
		return true
	case WWA:
		return true
	case WY:
		return true
	}
	return false
}
