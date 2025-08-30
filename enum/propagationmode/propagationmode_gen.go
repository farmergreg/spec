// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package propagationmode

var (
	AS       PropagationMode = "AS"       // AS         = Aircraft Scatter
	AUE      PropagationMode = "AUE"      // AUE        = Aurora-E
	AUR      PropagationMode = "AUR"      // AUR        = Aurora
	BS       PropagationMode = "BS"       // BS         = Back scatter
	ECH      PropagationMode = "ECH"      // ECH        = EchoLink
	EME      PropagationMode = "EME"      // EME        = Earth-Moon-Earth
	ES       PropagationMode = "ES"       // ES         = Sporadic E
	F2       PropagationMode = "F2"       // F2         = F2 Reflection
	FAI      PropagationMode = "FAI"      // FAI        = Field Aligned Irregularities
	GWAVE    PropagationMode = "GWAVE"    // GWAVE      = Ground Wave
	INTERNET PropagationMode = "INTERNET" // INTERNET   = Internet-assisted
	ION      PropagationMode = "ION"      // ION        = Ionoscatter
	IRL      PropagationMode = "IRL"      // IRL        = IRLP
	LOS      PropagationMode = "LOS"      // LOS        = Line of Sight (includes transmission through obstacles such as walls)
	MS       PropagationMode = "MS"       // MS         = Meteor scatter
	RPT      PropagationMode = "RPT"      // RPT        = Terrestrial or atmospheric repeater or transponder
	RS       PropagationMode = "RS"       // RS         = Rain scatter
	SAT      PropagationMode = "SAT"      // SAT        = Satellite
	TEP      PropagationMode = "TEP"      // TEP        = Trans-equatorial
	TR       PropagationMode = "TR"       // TR         = Tropospheric ducting
)

// A map of all PropagationMode specifications.
var PropagationModeMap = map[PropagationMode]Spec{
	AS:       {IsImportOnly: false, Key: "AS", Description: "Aircraft Scatter"},
	AUE:      {IsImportOnly: false, Key: "AUE", Description: "Aurora-E"},
	AUR:      {IsImportOnly: false, Key: "AUR", Description: "Aurora"},
	BS:       {IsImportOnly: false, Key: "BS", Description: "Back scatter"},
	ECH:      {IsImportOnly: false, Key: "ECH", Description: "EchoLink"},
	EME:      {IsImportOnly: false, Key: "EME", Description: "Earth-Moon-Earth"},
	ES:       {IsImportOnly: false, Key: "ES", Description: "Sporadic E"},
	F2:       {IsImportOnly: false, Key: "F2", Description: "F2 Reflection"},
	FAI:      {IsImportOnly: false, Key: "FAI", Description: "Field Aligned Irregularities"},
	GWAVE:    {IsImportOnly: false, Key: "GWAVE", Description: "Ground Wave"},
	INTERNET: {IsImportOnly: false, Key: "INTERNET", Description: "Internet-assisted"},
	ION:      {IsImportOnly: false, Key: "ION", Description: "Ionoscatter"},
	IRL:      {IsImportOnly: false, Key: "IRL", Description: "IRLP"},
	LOS:      {IsImportOnly: false, Key: "LOS", Description: "Line of Sight (includes transmission through obstacles such as walls)"},
	MS:       {IsImportOnly: false, Key: "MS", Description: "Meteor scatter"},
	RPT:      {IsImportOnly: false, Key: "RPT", Description: "Terrestrial or atmospheric repeater or transponder"},
	RS:       {IsImportOnly: false, Key: "RS", Description: "Rain scatter"},
	SAT:      {IsImportOnly: false, Key: "SAT", Description: "Satellite"},
	TEP:      {IsImportOnly: false, Key: "TEP", Description: "Trans-equatorial"},
	TR:       {IsImportOnly: false, Key: "TR", Description: "Tropospheric ducting"},
}

// All PropagationMode specifications including depreciated and import only.
var PropagationModeListAll = []Spec{
	PropagationModeMap[AS],
	PropagationModeMap[AUE],
	PropagationModeMap[AUR],
	PropagationModeMap[BS],
	PropagationModeMap[ECH],
	PropagationModeMap[EME],
	PropagationModeMap[ES],
	PropagationModeMap[F2],
	PropagationModeMap[FAI],
	PropagationModeMap[GWAVE],
	PropagationModeMap[INTERNET],
	PropagationModeMap[ION],
	PropagationModeMap[IRL],
	PropagationModeMap[LOS],
	PropagationModeMap[MS],
	PropagationModeMap[RPT],
	PropagationModeMap[RS],
	PropagationModeMap[SAT],
	PropagationModeMap[TEP],
	PropagationModeMap[TR],
}

// All PropagationMode specifications values that are NOT marked import-only.
var PropagationModeListCurrent = []Spec{
	PropagationModeMap[AS],
	PropagationModeMap[AUE],
	PropagationModeMap[AUR],
	PropagationModeMap[BS],
	PropagationModeMap[ECH],
	PropagationModeMap[EME],
	PropagationModeMap[ES],
	PropagationModeMap[F2],
	PropagationModeMap[FAI],
	PropagationModeMap[GWAVE],
	PropagationModeMap[INTERNET],
	PropagationModeMap[ION],
	PropagationModeMap[IRL],
	PropagationModeMap[LOS],
	PropagationModeMap[MS],
	PropagationModeMap[RPT],
	PropagationModeMap[RS],
	PropagationModeMap[SAT],
	PropagationModeMap[TEP],
	PropagationModeMap[TR],
}
