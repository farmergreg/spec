package contest

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/spec"
)

var (
	// EnumContestMap contains ALL records, including un-released and import-only records
	EnumContestMap EnumContestItemMap

	// EnumContestListAll contains ALL records, including un-released and import-only records
	EnumContestListAll EnumContestItemList

	// EnumContestList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumContestList EnumContestItemList
)

func init() {
	var err error
	EnumContestListAll, err = spec.ParseADISpecTSV[*EnumContestItem](spec.EnumContestTSV)
	if err != nil {
		panic(err)
	}

	EnumContestList = make([]*EnumContestItem, 0, len(EnumContestListAll))
	for _, item := range EnumContestListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumContestList = append(EnumContestList, item)
		}
	}
	EnumContestList = slices.Clip(EnumContestList)

	EnumContestMap = make(EnumContestItemMap, len(EnumContestList))
	for _, item := range EnumContestList {
		EnumContestMap[item.ID] = item
	}
}

// Contest represents the contest identifier
type Contest string

// EnumContestItemList represents a collection of contest definitions
type EnumContestItemList []*EnumContestItem

// EnumContestItemMap maps Contest to its definition
type EnumContestItemMap map[Contest]*EnumContestItem

// EnumContestItem represents a contest item
type EnumContestItem struct {
	shared.ImportRecordRoot
	ID          Contest `csv:"Contest-ID"`  // The value that is stored in the CONTEST_ID ADIF field.
	Description string  `csv:"Description"` // The description of the contest.
}
