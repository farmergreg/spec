package contest

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumContestMap contains ALL records, including un-released and import-only records
	EnumContestMap map[Contest]*EnumContestItem

	// EnumContestListAll contains ALL records, including un-released and import-only records
	EnumContestListAll []*EnumContestItem

	// EnumContestList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumContestList []*EnumContestItem
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

	EnumContestMap = make(map[Contest]*EnumContestItem, len(EnumContestList))
	for _, item := range EnumContestList {
		EnumContestMap[item.ID] = item
	}
}

// Contest represents the contest identifier
type Contest string

// EnumContestItem represents a contest item
type EnumContestItem struct {
	shared.ImportRecordRoot
	ID          Contest `csv:"Contest-ID"`  // The value that is stored in the CONTEST_ID ADIF field.
	Description string  `csv:"Description"` // The description of the contest.
}
