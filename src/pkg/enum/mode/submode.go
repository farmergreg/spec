package mode

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumSubModeMap contains ALL records, including un-released and import-only records
	EnumSubModeMap map[SubMode]*EnumSubModeItem

	// EnumSubModeListAll contains ALL records, including un-released and import-only records
	EnumSubModeListAll []*EnumSubModeItem

	// EnumSubModeList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumSubModeList []*EnumSubModeItem
)

func init() {
	var err error
	EnumSubModeListAll, err = spec.ParseADISpecTSV[*EnumSubModeItem](spec.EnumSubModeTSV)
	if err != nil {
		panic(err)
	}

	EnumSubModeList = make([]*EnumSubModeItem, 0, len(EnumSubModeListAll))
	for _, item := range EnumSubModeListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumSubModeList = append(EnumSubModeList, item)
		}
	}
	EnumSubModeList = slices.Clip(EnumSubModeList)

	EnumSubModeMap = make(map[SubMode]*EnumSubModeItem, len(EnumSubModeList))
	for _, item := range EnumSubModeList {
		EnumSubModeMap[item.ID] = item
	}
}

// SubMode represents the submode of an ADIF record
type SubMode string

// EnumSubModeItem represents a submode item
type EnumSubModeItem struct {
	shared.ImportRecordRoot
	ID   SubMode `csv:"Submode"` // The value that is stored in the SUBMODE ADIF field.
	Mode Mode    `csv:"Mode"`    // The related value that must be stored in the MODE ADIF field.
}
