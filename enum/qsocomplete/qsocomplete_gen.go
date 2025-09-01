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

// A map of all QSOComplete specifications.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSOCompleteMap = map[QSOComplete]Spec{
	Uncertain: {IsImportOnly: false, Key: "?", Description: "uncertain"},
	N:         {IsImportOnly: false, Key: "N", Description: "no"},
	NIL:       {IsImportOnly: false, Key: "NIL", Description: "not heard"},
	Y:         {IsImportOnly: false, Key: "Y", Description: "yes"},
}

// All QSOComplete specifications including depreciated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSOCompleteListAll = []Spec{
	QSOCompleteMap[Uncertain],
	QSOCompleteMap[N],
	QSOCompleteMap[NIL],
	QSOCompleteMap[Y],
}

// All QSOComplete specifications values that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSOCompleteListCurrent = []Spec{
	QSOCompleteMap[Uncertain],
	QSOCompleteMap[N],
	QSOCompleteMap[NIL],
	QSOCompleteMap[Y],
}
