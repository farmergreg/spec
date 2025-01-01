package award

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumAwardMap contains ALL records, including un-released and import-only records
	EnumAwardMap EnumAwardItemMap

	// EnumAwardListAll contains ALL records, including un-released and import-only records
	EnumAwardListAll EnumAwardItemList

	// EnumAwardList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumAwardList EnumAwardItemList
)

func init() {
	var err error
	EnumAwardListAll, err = spec.ParseADISpecTSV[*EnumAwardItem](spec.EnumAwardTSV)
	if err != nil {
		panic(err)
	}

	EnumAwardList = make([]*EnumAwardItem, 0, len(EnumAwardListAll))
	for _, item := range EnumAwardListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumAwardList = append(EnumAwardList, item)
		}
	}
	EnumAwardList = slices.Clip(EnumAwardList)

	EnumAwardMap = make(EnumAwardItemMap, len(EnumAwardList))
	for _, item := range EnumAwardList {
		EnumAwardMap[item.ID] = item
	}
}

/***************************************************
 * Award is DIFFERENT
 * All of it's values are IMPORT-ONLY.
 ***************************************************/

// Award represents an ADIF award type
type Award string

// EnumAwardItemList represents a collection of Award definitions
type EnumAwardItemList []*EnumAwardItem

// EnumAwardItemMap maps Award to its definition
type EnumAwardItemMap map[Award]*EnumAwardItem

// EnumAwardItem represents an award item
type EnumAwardItem struct {
	shared.ImportRecordRoot
	ID Award `csv:"Award"` // Depreciated. The value that is stored in the CREDIT_GRANTED ADIF field. May also appear next to a Credit
}
