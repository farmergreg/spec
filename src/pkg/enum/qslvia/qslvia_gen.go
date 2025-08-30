// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package qslvia

var (
	B QSLVia = "B" // B = bureau
	D QSLVia = "D" // D = direct
	E QSLVia = "E" // E = electronic
	M QSLVia = "M" // Deprecated: M = manager
)

// A map of all QSLVia specifications.
var QSLViaMap = map[QSLVia]Spec{
	B: {IsImportOnly: false, Comments: "", Key: "B", Description: "bureau"},
	D: {IsImportOnly: false, Comments: "", Key: "D", Description: "direct"},
	E: {IsImportOnly: false, Comments: "", Key: "E", Description: "electronic"},
	M: {IsImportOnly: true, Comments: "", Key: "M", Description: "manager"},
}

// All QSLVia specifications including depreciated and import only.
var QSLViaListAll = []Spec{
	QSLViaMap[B],
	QSLViaMap[D],
	QSLViaMap[E],
	QSLViaMap[M],
}

// All QSLVia specifications values that are NOT marked import-only.
var QSLViaListCurrent = []Spec{
	QSLViaMap[B],
	QSLViaMap[D],
	QSLViaMap[E],
}
