// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package eqslag provides code and constants as defined in ADIF 3.1.6 (Proposed)
package eqslag

import "maps"

const (
	N EQSLAG = "N" // N = the QSO is confirmed but not "Authenticity Guaranteed" by eQSL
	U EQSLAG = "U" // U = unknown
	Y EQSLAG = "Y" // Y = the QSO is confirmed and "Authenticity Guaranteed" by eQSL
)

// All EQSLAG specifications in ADIF 3.1.6 (Proposed) including depreciated and import only.
func EQSLAGListAll() []Spec {
	return append([]Spec(nil), internalEQSLAGListAll...)
}

// All EQSLAG specifications values in ADIF 3.1.6 (Proposed) that are NOT marked import-only.
func EQSLAGListCurrent() []Spec {
	return append([]Spec(nil), internalEQSLAGListCurrent...)
}

// A map of all EQSLAG from ADIF 3.1.6 (Proposed).
func EQSLAGMap() map[EQSLAG]Spec {
	cp := make(map[EQSLAG]Spec, len(internalEQSLAGMap))
	maps.Copy(cp, internalEQSLAGMap)
	return cp
}

// A map of all EQSLAG specifications.
var internalEQSLAGMap = map[EQSLAG]Spec{
	N: {IsImportOnly: false, Key: "N", Description: "the QSO is confirmed but not \"Authenticity Guaranteed\" by eQSL"},
	U: {IsImportOnly: false, Key: "U", Description: "unknown"},
	Y: {IsImportOnly: false, Key: "Y", Description: "the QSO is confirmed and \"Authenticity Guaranteed\" by eQSL"},
}

var internalEQSLAGListAll = []Spec{
	internalEQSLAGMap[N],
	internalEQSLAGMap[U],
	internalEQSLAGMap[Y],
}

var internalEQSLAGListCurrent = []Spec{
	internalEQSLAGMap[N],
	internalEQSLAGMap[U],
	internalEQSLAGMap[Y],
}
