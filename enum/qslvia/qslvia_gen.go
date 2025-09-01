// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslvia provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslvia

const (
	B QSLVia = "B" // B = bureau
	D QSLVia = "D" // D = direct
	E QSLVia = "E" // E = electronic
	M QSLVia = "M" // Deprecated: M = manager
)

// A map of all QSLVia specifications.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSLViaMap = map[QSLVia]Spec{
	B: {IsImportOnly: false, Key: "B", Description: "bureau"},
	D: {IsImportOnly: false, Key: "D", Description: "direct"},
	E: {IsImportOnly: false, Key: "E", Description: "electronic"},
	M: {IsImportOnly: true, Key: "M", Description: "manager"},
}

// All QSLVia specifications including depreciated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSLViaListAll = []Spec{
	QSLViaMap[B],
	QSLViaMap[D],
	QSLViaMap[E],
	QSLViaMap[M],
}

// All QSLVia specifications values that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSLViaListCurrent = []Spec{
	QSLViaMap[B],
	QSLViaMap[D],
	QSLViaMap[E],
}
