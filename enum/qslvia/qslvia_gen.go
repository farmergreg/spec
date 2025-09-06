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

// Lookup looks up a QSLVia specification
func Lookup(qslvia QSLVia) (Spec, bool) {
	spec, ok := internalQSLViaMap[qslvia], true
	return spec, ok
}

var internalQSLViaMap = map[QSLVia]Spec{
	B: {IsImportOnly: false, Key: "B", Description: "bureau"},
	D: {IsImportOnly: false, Key: "D", Description: "direct"},
	E: {IsImportOnly: false, Key: "E", Description: "electronic"},
	M: {IsImportOnly: true, Key: "M", Description: "manager"},
}

// All QSLVia specifications including deprecated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSLViaListAll = []Spec{
	internalQSLViaMap[B],
	internalQSLViaMap[D],
	internalQSLViaMap[E],
	internalQSLViaMap[M],
}

// All QSLVia specifications that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSLViaListCurrent = []Spec{
	internalQSLViaMap[B],
	internalQSLViaMap[D],
	internalQSLViaMap[E],
}
