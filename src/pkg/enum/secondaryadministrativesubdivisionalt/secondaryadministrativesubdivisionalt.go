package secondaryadministrativesubdivisionalt

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumSecondaryAdministrativeSubdivisionAltMap contains ALL records, including un-released and import-only records
	EnumSecondaryAdministrativeSubdivisionAltMap map[SecondaryAdministrativeSubdivisionAlt]*EnumSecondaryAdministrativeSubdivisionAltItem

	// EnumSecondaryAdministrativeSubdivisionAltListAll contains ALL records, including un-released and import-only records
	EnumSecondaryAdministrativeSubdivisionAltListAll []*EnumSecondaryAdministrativeSubdivisionAltItem

	// EnumSecondaryAdministrativeSubdivisionAltList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumSecondaryAdministrativeSubdivisionAltList []*EnumSecondaryAdministrativeSubdivisionAltItem
)

func init() {
	var err error
	EnumSecondaryAdministrativeSubdivisionAltListAll, err = spec.ParseADISpecTSV[*EnumSecondaryAdministrativeSubdivisionAltItem](spec.EnumSecondaryAdministrativeSubdivisionAltTSV)
	if err != nil {
		panic(err)
	}

	EnumSecondaryAdministrativeSubdivisionAltList = make([]*EnumSecondaryAdministrativeSubdivisionAltItem, 0, len(EnumSecondaryAdministrativeSubdivisionAltListAll))
	for _, item := range EnumSecondaryAdministrativeSubdivisionAltListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumSecondaryAdministrativeSubdivisionAltList = append(EnumSecondaryAdministrativeSubdivisionAltList, item)
		}
	}
	EnumSecondaryAdministrativeSubdivisionAltList = slices.Clip(EnumSecondaryAdministrativeSubdivisionAltList)

	EnumSecondaryAdministrativeSubdivisionAltMap = make(map[SecondaryAdministrativeSubdivisionAlt]*EnumSecondaryAdministrativeSubdivisionAltItem, len(EnumSecondaryAdministrativeSubdivisionAltList))
	for _, item := range EnumSecondaryAdministrativeSubdivisionAltList {
		EnumSecondaryAdministrativeSubdivisionAltMap[item.ID] = item
	}
}

// SecondaryAdministrativeSubdivisionAlt represents an antenna path abbreviation
type SecondaryAdministrativeSubdivisionAlt string

// EnumSecondaryAdministrativeSubdivisionAltItem represents an antenna path item
type EnumSecondaryAdministrativeSubdivisionAltItem struct {
	shared.ImportRecordRoot
	ID             SecondaryAdministrativeSubdivisionAlt `csv:"Code"` // The value that is stored in the ANT_PATH ADIF field.
	DXCCEntityCode dxccentitycode.DXCCEntityCode         `csv:"DXCC Entity Code"`
	Region         string                                `csv:"Region"`
	District       string                                `csv:"District"`
	Deleted        shared.Deleted                        `csv:"Deleted"`
}
