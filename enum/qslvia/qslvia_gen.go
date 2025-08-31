// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslvia provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslvia

import "maps"

const (
	B QSLVia = "B" // B = bureau
	D QSLVia = "D" // D = direct
	E QSLVia = "E" // E = electronic
	M QSLVia = "M" // Deprecated: M = manager
)

// All QSLVia specifications in ADIF 3.1.6 (Proposed) including depreciated and import only.
func QSLViaListAll() []Spec {
	return append([]Spec(nil), internalQSLViaListAll...)
}

// All QSLVia specifications values in ADIF 3.1.6 (Proposed) that are NOT marked import-only.
func QSLViaListCurrent() []Spec {
	return append([]Spec(nil), internalQSLViaListCurrent...)
}

// A map of all QSLVia from ADIF 3.1.6 (Proposed).
func QSLViaMap() map[QSLVia]Spec {
	cp := make(map[QSLVia]Spec, len(internalQSLViaMap))
	maps.Copy(cp, internalQSLViaMap)
	return cp
}

// A map of all QSLVia specifications.
var internalQSLViaMap = map[QSLVia]Spec{
	B: {IsImportOnly: false, Key: "B", Description: "bureau"},
	D: {IsImportOnly: false, Key: "D", Description: "direct"},
	E: {IsImportOnly: false, Key: "E", Description: "electronic"},
	M: {IsImportOnly: true, Key: "M", Description: "manager"},
}

var internalQSLViaListAll = []Spec{
	internalQSLViaMap[B],
	internalQSLViaMap[D],
	internalQSLViaMap[E],
	internalQSLViaMap[M],
}

var internalQSLViaListCurrent = []Spec{
	internalQSLViaMap[B],
	internalQSLViaMap[D],
	internalQSLViaMap[E],
}
