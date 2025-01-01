package antpath

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumAntPathMap contains ALL records, including un-released and import-only records
	EnumAntPathMap EnumAntPathItemMap

	// EnumAntPathListAll contains ALL records, including un-released and import-only records
	EnumAntPathListAll EnumAntPathItemList

	// EnumAntPathList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumAntPathList EnumAntPathItemList
)

func init() {
	var err error
	EnumAntPathListAll, err = spec.ParseADISpecTSV[*EnumAntPathItem](spec.EnumAntPathTSV)
	if err != nil {
		panic(err)
	}

	EnumAntPathList = make([]*EnumAntPathItem, 0, len(EnumAntPathListAll))
	for _, item := range EnumAntPathListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumAntPathList = append(EnumAntPathList, item)
		}
	}
	EnumAntPathList = slices.Clip(EnumAntPathList)

	EnumAntPathMap = make(EnumAntPathItemMap, len(EnumAntPathList))
	for _, item := range EnumAntPathList {
		EnumAntPathMap[item.ID] = item
	}
}

// AntPath represents an antenna path abbreviation
type AntPath string

// EnumAntPathItemList represents a collection of antenna path definitions
type EnumAntPathItemList []*EnumAntPathItem

// EnumAntPathItemMap maps AntPath to its definition
type EnumAntPathItemMap map[AntPath]*EnumAntPathItem

// EnumAntPathItem represents an antenna path item
type EnumAntPathItem struct {
	shared.ImportRecordRoot
	ID          AntPath `csv:"Abbreviation"` // The value that is stored in the ANT_PATH ADIF field.
	Description string  `csv:"Meaning"`
}
