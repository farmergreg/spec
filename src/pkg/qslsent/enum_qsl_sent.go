package qslsent

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/spec"
)

var (
	// EnumQSLSentMap contains ALL records, including un-released and import-only records
	EnumQSLSentMap EnumQSLSentItemMap

	// EnumQSLSentListAll contains ALL records, including un-released and import-only records
	EnumQSLSentListAll EnumQSLSentItemList

	// EnumQSLSentList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumQSLSentList EnumQSLSentItemList
)

func init() {
	var err error
	EnumQSLSentListAll, err = spec.ParseADISpecTSV[*EnumQSLSentItem](spec.EnumQSLSentTSV)
	if err != nil {
		panic(err)
	}

	EnumQSLSentList = make([]*EnumQSLSentItem, 0, len(EnumQSLSentListAll))
	for _, item := range EnumQSLSentListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumQSLSentList = append(EnumQSLSentList, item)
		}
	}
	EnumQSLSentList = slices.Clip(EnumQSLSentList)

	EnumQSLSentMap = make(EnumQSLSentItemMap, len(EnumQSLSentList))
	for _, item := range EnumQSLSentList {
		EnumQSLSentMap[item.ID] = item
	}
}

// QSLSent represents the QSL sent status
type QSLSent string

// EnumQSLSentItemList represents a collection of QSL sent status items
type EnumQSLSentItemList []*EnumQSLSentItem

// EnumQSLSentItemMap maps QSLSent to its definition
type EnumQSLSentItemMap map[QSLSent]*EnumQSLSentItem

// EnumQSLSentItem represents a QSL sent status item
type EnumQSLSentItem struct {
	shared.ImportRecordRoot
	ID          QSLSent `csv:"Status"` // The value that is stored in the QSL_SENT, DCL_QSL_SENT, EQSL_QSL_SENT, and LOTW_QSL_SENT ADIF fields.
	Description string  `csv:"Meaning"`
}
