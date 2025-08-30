// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package qsocomplete

var (
	Uncertain QSOComplete = "?"   // ?    = uncertain
	N         QSOComplete = "N"   // N    = no
	NIL       QSOComplete = "NIL" // NIL  = not heard
	Y         QSOComplete = "Y"   // Y    = yes
)

// A map of all QSOComplete specifications.
var QSOCompleteMap = map[QSOComplete]Spec{
	Uncertain: {IsImportOnly: false, Key: "?", Description: "uncertain"},
	N:         {IsImportOnly: false, Key: "N", Description: "no"},
	NIL:       {IsImportOnly: false, Key: "NIL", Description: "not heard"},
	Y:         {IsImportOnly: false, Key: "Y", Description: "yes"},
}

// All QSOComplete specifications including depreciated and import only.
var QSOCompleteListAll = []Spec{
	QSOCompleteMap[Uncertain],
	QSOCompleteMap[N],
	QSOCompleteMap[NIL],
	QSOCompleteMap[Y],
}

// All QSOComplete specifications values that are NOT marked import-only.
var QSOCompleteListCurrent = []Spec{
	QSOCompleteMap[Uncertain],
	QSOCompleteMap[N],
	QSOCompleteMap[NIL],
	QSOCompleteMap[Y],
}
