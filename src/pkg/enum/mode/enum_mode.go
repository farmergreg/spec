package mode

import (
	"slices"
	"strings"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumModeMap contains ALL records, including un-released and import-only records
	EnumModeMap map[Mode]*EnumModeItem

	// EnumModeListAll contains ALL records, including un-released and import-only records
	EnumModeListAll []*EnumModeItem

	// EnumModeList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumModeList []*EnumModeItem
)

func init() {
	var err error
	EnumModeListAll, err = spec.ParseADISpecTSV[*EnumModeItem](spec.EnumModeTSV)
	if err != nil {
		panic(err)
	}

	EnumModeList = make([]*EnumModeItem, 0, len(EnumModeListAll))
	for _, item := range EnumModeListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumModeList = append(EnumModeList, item)
		}
	}
	EnumModeList = slices.Clip(EnumModeList)

	EnumModeMap = make(map[Mode]*EnumModeItem, len(EnumModeList))
	for _, item := range EnumModeList {
		EnumModeMap[item.ID] = item
	}
}

// Mode is the ADIF mode of a radio communication.
type Mode string

// EnumModeSubmodeList is a list of submodes as defined by the Modes in the ADIF specification export.
type EnumModeSubModeList struct {
	Submodes []SubMode
}

// EnumModeItem represents a Mode item
type EnumModeItem struct {
	shared.ImportRecordRoot
	ID       Mode                `csv:"Mode"`     // The value that is stored in the MODE ADIF field.
	Submodes EnumModeSubModeList `csv:"Submodes"` // The related value that may be stored in the SUBMODE ADIF field.
}

func (s *EnumModeSubModeList) UnmarshalCSV(csv string) error {
	sub := strings.Split(csv, ",")

	if len(sub) == 1 && sub[0] == "" {
		sub = []string{}
	}

	s.Submodes = make([]SubMode, len(sub))
	for i, v := range sub {
		s.Submodes[i] = SubMode(v)
	}
	return nil
}
