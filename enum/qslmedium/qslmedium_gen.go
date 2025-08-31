// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslmedium provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslmedium

import "maps"

const (
	CARD QSLMedium = "CARD" // CARD = QSO confirmation via paper QSL card
	EQSL QSLMedium = "EQSL" // EQSL = QSO confirmation via eQSL.cc
	LOTW QSLMedium = "LOTW" // LOTW = QSO confirmation via ARRL Logbook of the World
)

// All QSLMedium specifications in ADIF 3.1.6 (Proposed) including depreciated and import only.
func QSLMediumListAll() []Spec {
	return append([]Spec(nil), internalQSLMediumListAll...)
}

// All QSLMedium specifications values in ADIF 3.1.6 (Proposed) that are NOT marked import-only.
func QSLMediumListCurrent() []Spec {
	return append([]Spec(nil), internalQSLMediumListCurrent...)
}

// A map of all QSLMedium from ADIF 3.1.6 (Proposed).
func QSLMediumMap() map[QSLMedium]Spec {
	cp := make(map[QSLMedium]Spec, len(internalQSLMediumMap))
	maps.Copy(cp, internalQSLMediumMap)
	return cp
}

// A map of all QSLMedium specifications.
var internalQSLMediumMap = map[QSLMedium]Spec{
	CARD: {IsImportOnly: false, Key: "CARD", Description: "QSO confirmation via paper QSL card"},
	EQSL: {IsImportOnly: false, Key: "EQSL", Description: "QSO confirmation via eQSL.cc"},
	LOTW: {IsImportOnly: false, Key: "LOTW", Description: "QSO confirmation via ARRL Logbook of the World"},
}

var internalQSLMediumListAll = []Spec{
	internalQSLMediumMap[CARD],
	internalQSLMediumMap[EQSL],
	internalQSLMediumMap[LOTW],
}

var internalQSLMediumListCurrent = []Spec{
	internalQSLMediumMap[CARD],
	internalQSLMediumMap[EQSL],
	internalQSLMediumMap[LOTW],
}
