package continent

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumContinentMap contains ALL records, including un-released and import-only records
	EnumContinentMap map[Continent]*EnumContinentItem

	// EnumContinentListAll contains ALL records, including un-released and import-only records
	EnumContinentListAll []*EnumContinentItem

	// EnumContinentList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumContinentList []*EnumContinentItem
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

	EnumContinentMap = make(map[Continent]*EnumContinentItem, len(EnumContinentList))
	for _, item := range EnumContinentList {
		EnumContinentMap[item.ID] = item
	}
}

type Continent string

// EnumContinentItem represents a continent item
type EnumContinentItem struct {
	shared.ImportRecordRoot
	ID          Continent `csv:"Abbreviation"`
	Description string    `csv:"Continent"`
}
