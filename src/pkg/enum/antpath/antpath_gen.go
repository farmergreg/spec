// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package antpath

var (
	G AntPath = "G" // G = grayline
	L AntPath = "L" // L = long path
	O AntPath = "O" // O = other
	S AntPath = "S" // S = short path
)

// A map of all AntPath specifications.
var AntPathMap = map[AntPath]Spec{
	G: {IsImportOnly: false, Key: "G", Description: "grayline"},
	L: {IsImportOnly: false, Key: "L", Description: "long path"},
	O: {IsImportOnly: false, Key: "O", Description: "other"},
	S: {IsImportOnly: false, Key: "S", Description: "short path"},
}

// All AntPath specifications including depreciated and import only.
var AntPathListAll = []Spec{
	AntPathMap[G],
	AntPathMap[L],
	AntPathMap[O],
	AntPathMap[S],
}

// All AntPath specifications values that are NOT marked import-only.
var AntPathListCurrent = []Spec{
	AntPathMap[G],
	AntPathMap[L],
	AntPathMap[O],
	AntPathMap[S],
}
