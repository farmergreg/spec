// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package antpath provides code and constants as defined in ADIF 3.1.6 (Proposed)
package antpath

const (
	G AntPath = "G" // G = grayline
	L AntPath = "L" // L = long path
	O AntPath = "O" // O = other
	S AntPath = "S" // S = short path
)

// Lookup looks up a AntPath specification
func Lookup(antpath AntPath) (Spec, bool) {
	spec, ok := internalAntPathMap[antpath], true
	return spec, ok
}

var internalAntPathMap = map[AntPath]Spec{
	G: {IsImportOnly: false, Key: "G", Description: "grayline"},
	L: {IsImportOnly: false, Key: "L", Description: "long path"},
	O: {IsImportOnly: false, Key: "O", Description: "other"},
	S: {IsImportOnly: false, Key: "S", Description: "short path"},
}

// All AntPath specifications including deprecated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var AntPathListAll = []Spec{
	internalAntPathMap[G],
	internalAntPathMap[L],
	internalAntPathMap[O],
	internalAntPathMap[S],
}

// All AntPath specifications that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var AntPathListCurrent = []Spec{
	internalAntPathMap[G],
	internalAntPathMap[L],
	internalAntPathMap[O],
	internalAntPathMap[S],
}
