// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package qslmedium

var (
	CARD QSLMedium = "CARD" // CARD = QSO confirmation via paper QSL card
	EQSL QSLMedium = "EQSL" // EQSL = QSO confirmation via eQSL.cc
	LOTW QSLMedium = "LOTW" // LOTW = QSO confirmation via ARRL Logbook of the World
)

// A map of all QSLMedium specifications.
var QSLMediumMap = map[QSLMedium]Spec{
	CARD: {IsImportOnly: false, Comments: "", Key: "CARD", Description: "QSO confirmation via paper QSL card"},
	EQSL: {IsImportOnly: false, Comments: "", Key: "EQSL", Description: "QSO confirmation via eQSL.cc"},
	LOTW: {IsImportOnly: false, Comments: "", Key: "LOTW", Description: "QSO confirmation via ARRL Logbook of the World"},
}

// All QSLMedium specifications including depreciated and import only.
var QSLMediumListAll = []Spec{
	QSLMediumMap[CARD],
	QSLMediumMap[EQSL],
	QSLMediumMap[LOTW],
}

// All QSLMedium specifications values that are NOT marked import-only.
var QSLMediumListCurrent = []Spec{
	QSLMediumMap[CARD],
	QSLMediumMap[EQSL],
	QSLMediumMap[LOTW],
}
