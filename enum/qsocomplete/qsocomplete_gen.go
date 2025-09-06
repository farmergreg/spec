// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsocomplete provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qsocomplete

const (
	Uncertain QSOComplete = "?"   // ?    = uncertain
	N         QSOComplete = "N"   // N    = no
	NIL       QSOComplete = "NIL" // NIL  = not heard
	Y         QSOComplete = "Y"   // Y    = yes
)

// Lookup looks up a QSOComplete specification
func Lookup(qsocomplete QSOComplete) (Spec, bool) {
	spec, ok := internalQSOCompleteMap[qsocomplete], true
	return spec, ok
}

var internalQSOCompleteMap = map[QSOComplete]Spec{
	Uncertain: {IsImportOnly: false, Key: "?", Description: "uncertain"},
	N:         {IsImportOnly: false, Key: "N", Description: "no"},
	NIL:       {IsImportOnly: false, Key: "NIL", Description: "not heard"},
	Y:         {IsImportOnly: false, Key: "Y", Description: "yes"},
}

// All QSOComplete specifications including deprecated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSOCompleteListAll = []Spec{
	internalQSOCompleteMap[Uncertain],
	internalQSOCompleteMap[N],
	internalQSOCompleteMap[NIL],
	internalQSOCompleteMap[Y],
}

// All QSOComplete specifications that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSOCompleteListCurrent = []Spec{
	internalQSOCompleteMap[Uncertain],
	internalQSOCompleteMap[N],
	internalQSOCompleteMap[NIL],
	internalQSOCompleteMap[Y],
}
