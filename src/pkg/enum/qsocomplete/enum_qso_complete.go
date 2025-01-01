package qsocomplete

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumQSOCompleteMap contains ALL records, including un-released and import-only values
	EnumQSOCompleteMap EnumQSOCompleteItemMap

	// EnumQSOCompleteListAll contains ALL records, including un-released and import-only contests
	EnumQSOCompleteListAll EnumQSOCompleteItemList

	// EnumQSOCompleteList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumQSOCompleteList EnumQSOCompleteItemList
)

func init() {
	var err error
	EnumQSOCompleteListAll, err = spec.ParseADISpecTSV[*EnumQSOCompleteItem](spec.EnumQSOCompleteTSV)
	if err != nil {
		panic(err)
	}

	EnumQSOCompleteList = make([]*EnumQSOCompleteItem, 0, len(EnumQSOCompleteListAll))
	for _, item := range EnumQSOCompleteListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumQSOCompleteList = append(EnumQSOCompleteList, item)
		}
	}
	EnumQSOCompleteList = slices.Clip(EnumQSOCompleteList)

	EnumQSOCompleteMap = make(EnumQSOCompleteItemMap, len(EnumQSOCompleteList))
	for _, item := range EnumQSOCompleteList {
		EnumQSOCompleteMap[item.ID] = item
	}
}

// QSOComplete represents the completion status of a QSO
type QSOComplete string

// EnumQSOCompleteItemList represents a collection of QSO Complete definitions
type EnumQSOCompleteItemList []*EnumQSOCompleteItem

// EnumQSOCompleteItemMap maps QSOComplete to its definition
type EnumQSOCompleteItemMap map[QSOComplete]*EnumQSOCompleteItem

// EnumQSOCompleteItem represents a QSO Complete item
type EnumQSOCompleteItem struct {
	shared.ImportRecordRoot
	ID          QSOComplete `csv:"Abbreviation"` // The value that is stored in the QSO_COMPLETE ADIF field.
	Description string      `csv:"Meaning"`
}
