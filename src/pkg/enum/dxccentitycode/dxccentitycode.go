package dxccentitycode

import (
	"slices"
	"strconv"
	"strings"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumDXCCEntityCodeMap contains ALL records, including deleted, un-released and import-only records
	EnumDXCCEntityCodeMap map[DXCCEntityCode]*EnumDXCCEntityCodeItem

	// EnumDXCCEntityCodeListAll contains ALL records, including deleted, un-released and import-only records
	EnumDXCCEntityCodeListAll []*EnumDXCCEntityCodeItem

	// EnumDXCCEntityCodeList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumDXCCEntityCodeList []*EnumDXCCEntityCodeItem
)

func init() {
	var err error
	EnumDXCCEntityCodeListAll, err = spec.ParseADISpecTSV[*EnumDXCCEntityCodeItem](spec.EnumDXCCTSV)
	if err != nil {
		panic(err)
	}

	EnumDXCCEntityCodeList = make([]*EnumDXCCEntityCodeItem, 0, len(EnumDXCCEntityCodeListAll))
	for _, item := range EnumDXCCEntityCodeListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) && !bool(item.IsDeleted) {
			EnumDXCCEntityCodeList = append(EnumDXCCEntityCodeList, item)
		}
	}
	EnumDXCCEntityCodeList = slices.Clip(EnumDXCCEntityCodeList)

	EnumDXCCEntityCodeMap = make(map[DXCCEntityCode]*EnumDXCCEntityCodeItem, len(EnumDXCCEntityCodeList))
	for _, item := range EnumDXCCEntityCodeList {
		EnumDXCCEntityCodeMap[item.ID] = item
	}
}

// DXCCEntityCode is the code of a DXCCEntityCode entity.
type DXCCEntityCode int

// EnumDXCCEntityCodeItem represents a DXCC entity code item
type EnumDXCCEntityCodeItem struct {
	shared.ImportRecordRoot
	ID          DXCCEntityCode `csv:"Entity Code"` // The value that is stored in the DXCC and MY_DXCC ADIF fields.
	Description string         `csv:"Entity Name"` // The value that is stored in the COUNTRY and MY_COUNTRY ADIF fields.
	IsDeleted   shared.Deleted `csv:"Deleted"`
}

// DXCCEntityCodeList represents a list of DXCC entity codes
type DXCCEntityCodeList struct {
	Code []DXCCEntityCode
}

func (s *DXCCEntityCodeList) UnmarshalCSV(csv string) error {
	codes := strings.Split(csv, ",")

	if len(codes) == 1 && codes[0] == "" {
		codes = []string{}
	}

	for _, c := range codes {
		parsedCode, err := strconv.Atoi(c)
		if err != nil {
			return err
		}
		s.Code = append(s.Code, DXCCEntityCode(parsedCode))
	}

	return nil
}
