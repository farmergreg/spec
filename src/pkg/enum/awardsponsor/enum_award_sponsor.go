package awardsponsor

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumAwardSponsorMap contains ALL records, including un-released and import-only records
	EnumAwardSponsorMap EnumAwardSponsorItemMap

	// EnumAwardSponsorListAll contains ALL records, including un-released and import-only records
	EnumAwardSponsorListAll EnumAwardSponsorItemList

	// EnumAwardSponsorList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumAwardSponsorList EnumAwardSponsorItemList
)

func init() {
	var err error
	EnumAwardSponsorListAll, err = spec.ParseADISpecTSV[*EnumAwardSponsorItem](spec.EnumAwardSponsorTSV)
	if err != nil {
		panic(err)
	}

	EnumAwardSponsorList = make([]*EnumAwardSponsorItem, 0, len(EnumAwardSponsorListAll))
	for _, item := range EnumAwardSponsorListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumAwardSponsorList = append(EnumAwardSponsorList, item)
		}
	}
	EnumAwardSponsorList = slices.Clip(EnumAwardSponsorList)

	EnumAwardSponsorMap = make(EnumAwardSponsorItemMap, len(EnumAwardSponsorList))
	for _, item := range EnumAwardSponsorList {
		EnumAwardSponsorMap[item.IDPrefix] = item
	}
}

// AwardSponsorPrefix represents the unique identifier for an award sponsor
type AwardSponsorPrefix string

// EnumAwardSponsorItemList represents a collection of award sponsor definitions
type EnumAwardSponsorItemList []*EnumAwardSponsorItem

// EnumAwardSponsorItemMap maps AwardSponsorPrefix to its definition
type EnumAwardSponsorItemMap map[AwardSponsorPrefix]*EnumAwardSponsorItem

// EnumAwardSponsorItem represents an award sponsor item
type EnumAwardSponsorItem struct {
	shared.ImportRecordRoot
	IDPrefix    AwardSponsorPrefix `csv:"Sponsor"` // A prefix. The value that is stored in the AWARD_SUBMITTED and AWARD_GRANTED ADIF fields.
	Description string             `csv:"Sponsoring Organization"`
}
