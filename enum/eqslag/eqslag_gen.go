// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package eqslag

const (
	N EQSLAG = "N" // N = the QSO is confirmed but not "Authenticity Guaranteed" by eQSL
	U EQSLAG = "U" // U = unknown
	Y EQSLAG = "Y" // Y = the QSO is confirmed and "Authenticity Guaranteed" by eQSL
)

// A map of all EQSLAG specifications.
var EQSLAGMap = map[EQSLAG]Spec{
	N: {IsImportOnly: false, Key: "N", Description: "the QSO is confirmed but not \"Authenticity Guaranteed\" by eQSL"},
	U: {IsImportOnly: false, Key: "U", Description: "unknown"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO is confirmed and \"Authenticity Guaranteed\" by eQSL"},
}

// All EQSLAG specifications including depreciated and import only.
var EQSLAGListAll = []Spec{
	EQSLAGMap[N],
	EQSLAGMap[U],
	EQSLAGMap[Y],
}

// All EQSLAG specifications values that are NOT marked import-only.
var EQSLAGListCurrent = []Spec{
	EQSLAGMap[N],
	EQSLAGMap[U],
	EQSLAGMap[Y],
}
