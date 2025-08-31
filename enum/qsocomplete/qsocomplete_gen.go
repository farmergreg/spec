// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qsocomplete provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qsocomplete

import "maps"

const (
	Uncertain QSOComplete = "?"   // ?    = uncertain
	N         QSOComplete = "N"   // N    = no
	NIL       QSOComplete = "NIL" // NIL  = not heard
	Y         QSOComplete = "Y"   // Y    = yes
)

// All QSOComplete specifications in ADIF 3.1.6 (Proposed) including depreciated and import only.
func QSOCompleteListAll() []Spec {
	return append([]Spec(nil), internalQSOCompleteListAll...)
}

// All QSOComplete specifications values in ADIF 3.1.6 (Proposed) that are NOT marked import-only.
func QSOCompleteListCurrent() []Spec {
	return append([]Spec(nil), internalQSOCompleteListCurrent...)
}

// A map of all QSOComplete from ADIF 3.1.6 (Proposed).
func QSOCompleteMap() map[QSOComplete]Spec {
	cp := make(map[QSOComplete]Spec, len(internalQSOCompleteMap))
	maps.Copy(cp, internalQSOCompleteMap)
	return cp
}

// A map of all QSOComplete specifications.
var internalQSOCompleteMap = map[QSOComplete]Spec{
	Uncertain: {IsImportOnly: false, Key: "?", Description: "uncertain"},
	N:         {IsImportOnly: false, Key: "N", Description: "no"},
	NIL:       {IsImportOnly: false, Key: "NIL", Description: "not heard"},
	Y:         {IsImportOnly: false, Key: "Y", Description: "yes"},
}

var internalQSOCompleteListAll = []Spec{
	internalQSOCompleteMap[Uncertain],
	internalQSOCompleteMap[N],
	internalQSOCompleteMap[NIL],
	internalQSOCompleteMap[Y],
}

var internalQSOCompleteListCurrent = []Spec{
	internalQSOCompleteMap[Uncertain],
	internalQSOCompleteMap[N],
	internalQSOCompleteMap[NIL],
	internalQSOCompleteMap[Y],
}
