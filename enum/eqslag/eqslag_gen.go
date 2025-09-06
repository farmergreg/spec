// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package eqslag provides code and constants as defined in ADIF 3.1.6 (Proposed)
package eqslag

const (
	N EQSLAG = "N" // N = the QSO is confirmed but not "Authenticity Guaranteed" by eQSL
	U EQSLAG = "U" // U = unknown
	Y EQSLAG = "Y" // Y = the QSO is confirmed and "Authenticity Guaranteed" by eQSL
)

// Lookup looks up a EQSLAG specification
func Lookup(eqslag EQSLAG) (Spec, bool) {
	spec, ok := internalEQSLAGMap[eqslag], true
	return spec, ok
}

var internalEQSLAGMap = map[EQSLAG]Spec{
	N: {IsImportOnly: false, Key: "N", Description: "the QSO is confirmed but not \"Authenticity Guaranteed\" by eQSL"},
	U: {IsImportOnly: false, Key: "U", Description: "unknown"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO is confirmed and \"Authenticity Guaranteed\" by eQSL"},
}

// All EQSLAG specifications including deprecated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var EQSLAGListAll = []Spec{
	internalEQSLAGMap[N],
	internalEQSLAGMap[U],
	internalEQSLAGMap[Y],
}

// All EQSLAG specifications that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var EQSLAGListCurrent = []Spec{
	internalEQSLAGMap[N],
	internalEQSLAGMap[U],
	internalEQSLAGMap[Y],
}
