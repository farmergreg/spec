package qsocomplete

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumQSOCompleteMap contains ALL records, including un-released and import-only values
	EnumQSOCompleteMap map[QSOComplete]*EnumQSOCompleteItem

	// EnumQSOCompleteListAll contains ALL records, including un-released and import-only contests
	EnumQSOCompleteListAll []*EnumQSOCompleteItem

	// EnumQSOCompleteList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumQSOCompleteList []*EnumQSOCompleteItem
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

	EnumQSOCompleteMap = make(map[QSOComplete]*EnumQSOCompleteItem, len(EnumQSOCompleteList))
	for _, item := range EnumQSOCompleteList {
		EnumQSOCompleteMap[item.ID] = item
	}
}

// QSOComplete represents the completion status of a QSO
type QSOComplete string

// EnumQSOCompleteItem represents a QSO Complete item
type EnumQSOCompleteItem struct {
	shared.ImportRecordRoot
	ID          QSOComplete `csv:"Abbreviation"` // The value that is stored in the QSO_COMPLETE ADIF field.
	Description string      `csv:"Meaning"`
}
