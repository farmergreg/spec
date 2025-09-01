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

// A map of all AntPath specifications.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var AntPathMap = map[AntPath]Spec{
	G: {IsImportOnly: false, Key: "G", Description: "grayline"},
	L: {IsImportOnly: false, Key: "L", Description: "long path"},
	O: {IsImportOnly: false, Key: "O", Description: "other"},
	S: {IsImportOnly: false, Key: "S", Description: "short path"},
}

// All AntPath specifications including depreciated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var AntPathListAll = []Spec{
	AntPathMap[G],
	AntPathMap[L],
	AntPathMap[O],
	AntPathMap[S],
}

// All AntPath specifications values that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var AntPathListCurrent = []Spec{
	AntPathMap[G],
	AntPathMap[L],
	AntPathMap[O],
	AntPathMap[S],
}
