package continent

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumContinentMap contains ALL records, including un-released and import-only records
	EnumContinentMap EnumContinentItemMap

	// EnumContinentListAll contains ALL records, including un-released and import-only records
	EnumContinentListAll EnumContinentItemList

	// EnumContinentList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumContinentList EnumContinentItemList
)

func init() {
	var err error
	EnumContinentListAll, err = spec.ParseADISpecTSV[*EnumContinentItem](spec.EnumContinentTSV)
	if err != nil {
		panic(err)
	}

	EnumContinentList = make([]*EnumContinentItem, 0, len(EnumContinentListAll))
	for _, item := range EnumContinentListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumContinentList = append(EnumContinentList, item)
		}
	}
	EnumContinentList = slices.Clip(EnumContinentList)

	EnumContinentMap = make(EnumContinentItemMap, len(EnumContinentList))
	for _, item := range EnumContinentList {
		EnumContinentMap[item.ID] = item
	}
}

type Continent string

// EnumContinentItemList represents a collection of Continent items
type EnumContinentItemList []*EnumContinentItem

// EnumContinentItemMap maps Continent to its definition
type EnumContinentItemMap map[Continent]*EnumContinentItem

// EnumContinentItem represents a continent item
type EnumContinentItem struct {
	shared.ImportRecordRoot
	ID          Continent `csv:"Abbreviation"`
	Description string    `csv:"Continent"`
}
