package qslrcvd

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/spec"
)

var (
	// EnumQSLRcvdMap contains ALL records, including un-released and import-only records
	EnumQSLRcvdMap EnumQSLRcvdItemMap

	// EnumQSLRcvdListAll contains ALL records, including un-released and import-only records
	EnumQSLRcvdListAll EnumQSLRcvdItemList

	// EnumQSLRcvdList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumQSLRcvdList EnumQSLRcvdItemList
)

func init() {
	var err error
	EnumQSLRcvdListAll, err = spec.ParseADISpecTSV[*EnumQSLRcvdItem](spec.EnumQSLRcvdTSV)
	if err != nil {
		panic(err)
	}

	EnumQSLRcvdList = make([]*EnumQSLRcvdItem, 0, len(EnumQSLRcvdListAll))
	for _, item := range EnumQSLRcvdListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumQSLRcvdList = append(EnumQSLRcvdList, item)
		}
	}
	EnumQSLRcvdList = slices.Clip(EnumQSLRcvdList)

	EnumQSLRcvdMap = make(EnumQSLRcvdItemMap, len(EnumQSLRcvdList))
	for _, item := range EnumQSLRcvdList {
		EnumQSLRcvdMap[item.ID] = item
	}
}

// QSLRcvd represents the QSL received status
type QSLRcvd string

// EnumQSLRcvdItemList represents a collection of QSL received status definitions
type EnumQSLRcvdItemList []*EnumQSLRcvdItem

// EnumQSLRcvdItemMap maps QSLRcvd to its definition
type EnumQSLRcvdItemMap map[QSLRcvd]*EnumQSLRcvdItem

// EnumQSLRcvdItem represents a QSL received status item
type EnumQSLRcvdItem struct {
	shared.ImportRecordRoot
	ID          QSLRcvd `csv:"Status"` // The value that is stored in the QSL_RCVD, DCL_QSL_RCVD, EQSL_QSL_RCVD, and LOTW_QSL_RCVD ADIF fields.
	Description string  `csv:"Meaning"`
}
