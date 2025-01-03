package qslvia

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumQSLViaMap contains ALL records, including un-released and import-only records
	EnumQSLViaMap map[QSLVia]*EnumQSLViaItem

	// EnumQSLViaListAll contains all QSL Via records, including un-released and import-only records
	EnumQSLViaListAll []*EnumQSLViaItem

	// EnumQSLViaList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumQSLViaList []*EnumQSLViaItem
)

func init() {
	var err error
	EnumQSLViaListAll, err = spec.ParseADISpecTSV[*EnumQSLViaItem](spec.EnumQSLViaTSV)
	if err != nil {
		panic(err)
	}

	EnumQSLViaList = make([]*EnumQSLViaItem, 0, len(EnumQSLViaListAll))
	for _, item := range EnumQSLViaListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumQSLViaList = append(EnumQSLViaList, item)
		}
	}
	EnumQSLViaList = slices.Clip(EnumQSLViaList)

	EnumQSLViaMap = make(map[QSLVia]*EnumQSLViaItem, len(EnumQSLViaList))
	for _, item := range EnumQSLViaList {
		EnumQSLViaMap[item.ID] = item
	}
}

// QSLVia represents the method used to exchange QSL cards
type QSLVia string

// EnumQSLViaItem represents a QSL Via item
type EnumQSLViaItem struct {
	shared.ImportRecordRoot
	ID          QSLVia `csv:"Via"` // The value that is stored in the QSL_RCVD_VIA and QSL_SENT_VIA ADIF fields.
	Description string `csv:"Description"`
}
