// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslmedium provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslmedium

const (
	CARD QSLMedium = "CARD" // CARD = QSO confirmation via paper QSL card
	EQSL QSLMedium = "EQSL" // EQSL = QSO confirmation via eQSL.cc
	LOTW QSLMedium = "LOTW" // LOTW = QSO confirmation via ARRL Logbook of the World
)

// A map of all QSLMedium specifications.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSLMediumMap = map[QSLMedium]Spec{
	CARD: {IsImportOnly: false, Key: "CARD", Description: "QSO confirmation via paper QSL card"},
	EQSL: {IsImportOnly: false, Key: "EQSL", Description: "QSO confirmation via eQSL.cc"},
	LOTW: {IsImportOnly: false, Key: "LOTW", Description: "QSO confirmation via ARRL Logbook of the World"},
}

// All QSLMedium specifications including deprecated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSLMediumListAll = []Spec{
	QSLMediumMap[CARD],
	QSLMediumMap[EQSL],
	QSLMediumMap[LOTW],
}

// All QSLMedium specifications that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var QSLMediumListCurrent = []Spec{
	QSLMediumMap[CARD],
	QSLMediumMap[EQSL],
	QSLMediumMap[LOTW],
}
