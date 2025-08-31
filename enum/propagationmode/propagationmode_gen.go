// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package propagationmode provides code and constants as defined in ADIF 3.1.6 (Proposed)
package propagationmode

import "maps"

const (
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

// All PropagationMode specifications in ADIF 3.1.6 (Proposed) including depreciated and import only.
func PropagationModeListAll() []Spec {
	return append([]Spec(nil), internalPropagationModeListAll...)
}

// All PropagationMode specifications values in ADIF 3.1.6 (Proposed) that are NOT marked import-only.
func PropagationModeListCurrent() []Spec {
	return append([]Spec(nil), internalPropagationModeListCurrent...)
}

// A map of all PropagationMode from ADIF 3.1.6 (Proposed).
func PropagationModeMap() map[PropagationMode]Spec {
	cp := make(map[PropagationMode]Spec, len(internalPropagationModeMap))
	maps.Copy(cp, internalPropagationModeMap)
	return cp
}

// A map of all PropagationMode specifications.
var internalPropagationModeMap = map[PropagationMode]Spec{
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

var internalPropagationModeListAll = []Spec{
	internalPropagationModeMap[AS],
	internalPropagationModeMap[AUE],
	internalPropagationModeMap[AUR],
	internalPropagationModeMap[BS],
	internalPropagationModeMap[ECH],
	internalPropagationModeMap[EME],
	internalPropagationModeMap[ES],
	internalPropagationModeMap[F2],
	internalPropagationModeMap[FAI],
	internalPropagationModeMap[GWAVE],
	internalPropagationModeMap[INTERNET],
	internalPropagationModeMap[ION],
	internalPropagationModeMap[IRL],
	internalPropagationModeMap[LOS],
	internalPropagationModeMap[MS],
	internalPropagationModeMap[RPT],
	internalPropagationModeMap[RS],
	internalPropagationModeMap[SAT],
	internalPropagationModeMap[TEP],
	internalPropagationModeMap[TR],
}

var internalPropagationModeListCurrent = []Spec{
	internalPropagationModeMap[AS],
	internalPropagationModeMap[AUE],
	internalPropagationModeMap[AUR],
	internalPropagationModeMap[BS],
	internalPropagationModeMap[ECH],
	internalPropagationModeMap[EME],
	internalPropagationModeMap[ES],
	internalPropagationModeMap[F2],
	internalPropagationModeMap[FAI],
	internalPropagationModeMap[GWAVE],
	internalPropagationModeMap[INTERNET],
	internalPropagationModeMap[ION],
	internalPropagationModeMap[IRL],
	internalPropagationModeMap[LOS],
	internalPropagationModeMap[MS],
	internalPropagationModeMap[RPT],
	internalPropagationModeMap[RS],
	internalPropagationModeMap[SAT],
	internalPropagationModeMap[TEP],
	internalPropagationModeMap[TR],
}
