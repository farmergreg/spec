package primaryadministrativesubdivision

import (
	"slices"
	"strconv"
	"strings"

	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

// n.b. There is no map like the other enumerations have. This is because the IDs are not unique.
var (
	// EnumPrimaryAdministrativeSubdivisionListAll contains ALL records, including deleted, un-released and import-only records
	EnumPrimaryAdministrativeSubdivisionListAll []*EnumPrimaryAdministrativeSubdivisionItem

	// EnumPrimaryAdministrativeSubdivisionList
	// is a filtered list.
	// It excludes deleted, un-released and import-only records.
	EnumPrimaryAdministrativeSubdivisionList []*EnumPrimaryAdministrativeSubdivisionItem
)

func init() {
	var err error
	EnumPrimaryAdministrativeSubdivisionListAll, err = spec.ParseADISpecTSV[*EnumPrimaryAdministrativeSubdivisionItem](spec.EnumPrimaryAdministrativeSubdivisionTSV)
	if err != nil {
		panic(err)
	}

	EnumPrimaryAdministrativeSubdivisionList = make([]*EnumPrimaryAdministrativeSubdivisionItem, 0, len(EnumPrimaryAdministrativeSubdivisionListAll))
	for _, item := range EnumPrimaryAdministrativeSubdivisionListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) && !bool(item.IsDeleted) {
			EnumPrimaryAdministrativeSubdivisionList = append(EnumPrimaryAdministrativeSubdivisionList, item)
		}
	}
	EnumPrimaryAdministrativeSubdivisionList = slices.Clip(EnumPrimaryAdministrativeSubdivisionList)
}

// PrimaryAdministrativeSubdivision represents a primary administrative subdivision identifier
type PrimaryAdministrativeSubdivision string

// EnumPrimaryAdministrativeSubdivisionItem represents an ARRL section item
type EnumPrimaryAdministrativeSubdivisionItem struct {
	shared.ImportRecordRoot
	IsDeleted       shared.Deleted                    `csv:"Deleted"`
	ID              PrimaryAdministrativeSubdivision  `csv:"Code"` // n.b. NOT UNIQUE.
	Description     string                            `csv:"Primary Administrative Subdivision"`
	DXCCEntityCodes dxccentitycode.DXCCEntityCodeList `csv:"DXCC Entity Code"`
	ContainedWithin string                            `csv:"Contained Within"`
	Oblast          int                               `csv:"Oblast #"`
	CQZone          CQZoneList                        `csv:"CQ Zone"`
	ITUZone         ITUZoneList                       `csv:"ITU Zone"`
	Prefix          string                            `csv:"Prefix"`
	Comments        string                            `csv:"Comments"`
}

// CQZone represents a CQ zone
type CQZone int

// CQZoneList represents a list of CQ zones
type CQZoneList struct {
	Code []CQZone
}

func (s *CQZoneList) UnmarshalCSV(csv string) error {

	// Special case for ADIF specification weirdness with OB / Orenburg ...
	if csv == "S=16 T=17" {
		s.Code = []CQZone{16, 17}
		return nil
	}

	codes := strings.Split(csv, ",")

	if len(codes) == 1 && codes[0] == "" {
		codes = []string{}
	}

	for _, c := range codes {
		parsedCode, err := strconv.Atoi(c)
		if err != nil {
			return err
		}
		s.Code = append(s.Code, CQZone(parsedCode))
	}

	return nil
}

// ITUZone represents an ITU zone
type ITUZone int

// ITUZoneList represents a list of ITU zones
type ITUZoneList struct {
	Code []ITUZone
}

func (s *ITUZoneList) UnmarshalCSV(csv string) error {
	// For whatever reason, some of the Zones are split by "," and some are split by "/" in the ADIF specification...
	var codes []string
	if strings.Contains(csv, ",") {
		codes = strings.Split(csv, ",")
	} else {
		codes = strings.Split(csv, "/")
	}

	if len(codes) == 1 && codes[0] == "" {
		codes = []string{}
	}

	for _, c := range codes {
		parsedCode, err := strconv.Atoi(c)
		if err != nil {
			return err
		}
		s.Code = append(s.Code, ITUZone(parsedCode))
	}

	return nil
}
