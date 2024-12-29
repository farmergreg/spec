package dxccentitycode

import (
	"slices"
	"strconv"
	"strings"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/spec"
)

var (
	// EnumDXCCMap contains ALL records, including deleted, un-released and import-only records
	EnumDXCCMap EnumDXCCItemMap

	// EnumDXCCListAll contains ALL records, including deleted, un-released and import-only records
	EnumDXCCListAll EnumDXCCItemList

	// EnumDXCCList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumDXCCList EnumDXCCItemList
)

func init() {
	var err error
	EnumDXCCListAll, err = spec.ParseADISpecTSV[*EnumDXCCItem](spec.EnumDXCCTSV)
	if err != nil {
		panic(err)
	}

	EnumDXCCList = make([]*EnumDXCCItem, 0, len(EnumDXCCListAll))
	for _, item := range EnumDXCCListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) && !bool(item.IsDeleted) {
			EnumDXCCList = append(EnumDXCCList, item)
		}
	}
	EnumDXCCList = slices.Clip(EnumDXCCList)

	EnumDXCCMap = make(EnumDXCCItemMap, len(EnumDXCCList))
	for _, item := range EnumDXCCList {
		EnumDXCCMap[item.ID] = item
	}
}

// DXCC is the code of a DXCC entity.
type DXCC int

// EnumDXCCItemList represents a collection of DXCC entity code items
type EnumDXCCItemList []*EnumDXCCItem

// EnumDXCCItemMap maps DXCC to its definition
type EnumDXCCItemMap map[DXCC]*EnumDXCCItem

// EnumDXCCItem represents a DXCC entity code item
type EnumDXCCItem struct {
	shared.ImportRecordRoot
	ID          DXCC           `csv:"Entity Code"` // The value that is stored in the DXCC and MY_DXCC ADIF fields.
	Description string         `csv:"Entity Name"` // The value that is stored in the COUNTRY and MY_COUNTRY ADIF fields.
	IsDeleted   shared.Deleted `csv:"Deleted"`
}

// DXCCEntityList represents a list of DXCC entity codes
type DXCCEntityList struct {
	Code []DXCC
}

func (s *DXCCEntityList) UnmarshalCSV(csv string) error {
	codes := strings.Split(csv, ",")

	if len(codes) == 1 && codes[0] == "" {
		codes = []string{}
	}

	for _, c := range codes {
		parsedCode, err := strconv.Atoi(c)
		if err != nil {
			return err
		}
		s.Code = append(s.Code, DXCC(parsedCode))
	}

	return nil
}
