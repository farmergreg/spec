package region

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/spec"
)

var (
	// EnumRegionMap contains ALL records, including un-released and import-only records
	EnumRegionMap EnumRegionItemMap

	// EnumRegionListAll contains ALL records, including un-released and import-only records
	EnumRegionListAll EnumRegionItemList

	// EnumRegionList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumRegionList EnumRegionItemList
)

func init() {
	var err error
	EnumRegionListAll, err = spec.ParseADISpecTSV[*EnumRegionItem](spec.EnumRegionTSV)
	if err != nil {
		panic(err)
	}

	EnumRegionList = make([]*EnumRegionItem, 0, len(EnumRegionListAll))
	for _, item := range EnumRegionListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumRegionList = append(EnumRegionList, item)
		}
	}
	EnumRegionList = slices.Clip(EnumRegionList)

	EnumRegionMap = make(EnumRegionItemMap, len(EnumRegionList))
	for _, item := range EnumRegionList {
		EnumRegionMap[item.ID] = item
	}
}

// Region represents a region entity code
type Region string

// EnumRegionItemList represents a collection of region items
type EnumRegionItemList []*EnumRegionItem

// EnumRegionItemMap maps Region to its definition
type EnumRegionItemMap map[Region]*EnumRegionItem

// EnumRegionItem represents a region item
type EnumRegionItem struct {
	shared.ImportRecordRoot
	ID            Region              `csv:"Region Entity Code"` // The value that is stored in the REGION ADIF field.
	Description   string              `csv:"Region"`
	DXCCCode      dxccentitycode.DXCC `csv:"DXCC Entity Code"`
	Prefix        string              `csv:"Prefix"`
	Applicability string              `csv:"Applicability"`
	Comments      string              `csv:"Comments"`
}
