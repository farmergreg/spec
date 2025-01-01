package arrlsection

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/spec"
)

var (
	// EnumARRLSectionMap contains ALL records, including un-released and import-only records
	EnumARRLSectionMap EnumARRLSectionItemMap

	// EnumARRLSectionListAll contains ALL records, including un-released and import-only records
	EnumARRLSectionListAll EnumARRLSectionItemList

	// EnumARRLSectionList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumARRLSectionList EnumARRLSectionItemList
)

func init() {
	var err error
	EnumARRLSectionListAll, err = spec.ParseADISpecTSV[*EnumARRLSectionItem](spec.EnumARRLSectionTSV)
	if err != nil {
		panic(err)
	}

	EnumARRLSectionList = make([]*EnumARRLSectionItem, 0, len(EnumARRLSectionListAll))
	for _, item := range EnumARRLSectionListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) && item.DeletedDate.Time.IsZero() {
			EnumARRLSectionList = append(EnumARRLSectionList, item)
		}
	}
	EnumARRLSectionList = slices.Clip(EnumARRLSectionList)

	EnumARRLSectionMap = make(EnumARRLSectionItemMap, len(EnumARRLSectionList))
	for _, item := range EnumARRLSectionList {
		EnumARRLSectionMap[item.ID] = item
	}
}

// ARRLSection represents an ARRL section identifier
type ARRLSection string

// EnumARRLSectionItemList represents a collection of ARRL section definitions
type EnumARRLSectionItemList []*EnumARRLSectionItem

// EnumARRLSectionItemMap maps ARRLSection to its definition
type EnumARRLSectionItemMap map[ARRLSection]*EnumARRLSectionItem

// EnumARRLSectionItem represents an ARRL section item
type EnumARRLSectionItem struct {
	shared.ImportRecordRoot
	ID              ARRLSection                       `csv:"Abbreviation"` // The value that is stored in the ARRL_SECTION and MY_ARRL_SECTION ADIF fields.
	Description     string                            `csv:"Section Name"`
	DXCCEntityCodes dxccentitycode.DXCCEntityCodeList `csv:"DXCC Entity Code"`
	FromDate        shared.ADIDate                    `csv:"From Date"`
	DeletedDate     shared.ADIDate                    `csv:"Deleted Date"`
}
