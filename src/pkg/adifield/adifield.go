package adifield

import (
	"fmt"
	"slices"
	"strings"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// FieldMap contains ALL records, including un-released and import-only records
	FieldMap map[Field]*FieldDefinition

	// FieldListAll contains ALL records, including un-released and import-only records
	FieldListAll []*FieldDefinition

	// FieldList
	// is a filtered list.
	// It excludes un-released and import-only records.
	FieldList []*FieldDefinition
)

func init() {
	var err error
	FieldListAll, err = spec.ParseADISpecTSV[*FieldDefinition](spec.FieldsTSV)
	if err != nil {
		panic(err)
	}

	fieldListExtra, err := spec.ParseADISpecTSV[*FieldDefinition](spec.FieldsExtraTSV)
	if err != nil {
		panic(err)
	}
	FieldListAll = append(FieldListAll, fieldListExtra...)

	// special case, rename USERDEFn in FieldListAll to USERDEF1
	var userDefDescription string
	for _, item := range FieldListAll {
		if item.ID == USERDEF+"n" {
			item.ID = USERDEF + "1"
			userDefDescription = item.Description
			break
		}
	}

	// uppercase the field names just to be safe...
	for _, f := range FieldListAll {
		f.ID = Field(strings.ToUpper(string(f.ID)))
	}

	for i := 2; i < 10; i++ {
		FieldListAll = append(FieldListAll, &FieldDefinition{
			ID:            Field(fmt.Sprintf("USERDEF%d", i)),
			IsHeaderField: true,
			Description:   userDefDescription,
			ImportRecordRoot: shared.ImportRecordRoot{
				IsReleased:   shared.Released(true),
				IsImportOnly: shared.ImportOnly(false),
			},
		})
	}

	slices.SortStableFunc(FieldListAll, func(a, b *FieldDefinition) int {
		// Sort headers last
		if bool(a.IsHeaderField) != bool(b.IsHeaderField) {
			if bool(a.IsHeaderField) {
				return -1
			}
			return 1
		}

		return strings.Compare(string(a.ID), string(b.ID))
	})

	FieldListAll = slices.Clip(FieldListAll)

	FieldList = make([]*FieldDefinition, 0, len(FieldListAll))
	for _, item := range FieldListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			FieldList = append(FieldList, item)
		}
	}
	FieldList = slices.Clip(FieldList)

	FieldMap = make(map[Field]*FieldDefinition, len(FieldList))
	for _, item := range FieldList {
		FieldMap[item.ID] = item
	}
}

const (
	// USERDEF is the prefix for user defined fields.
	// These fields must always have a number after them, e.g. USERDEF1, USERDEF2, etc.
	USERDEF = "USERDEF"

	// APP_ is the prefix for application defined fields.
	// These field follow the pattern: APP_PROGRAMID_FIELDNAME
	// For example, APP_K9CTS_TEST_FIELD
	APP_ = "APP_"
)

// Field is the ADIF field name in and ADI file.
// By convention in this package, field name constants are always UPPERCASE.
// This is a departure from the ADIF specification, which allows field names to include lowercase letters.
type Field string

// FieldDefinition represents an ADIF field definition
type FieldDefinition struct {
	shared.ImportRecordRoot
	IsHeaderField shared.HeaderField `csv:"Header Field"`
	ID            Field              `csv:"Field Name"`
	DataType      string             `csv:"Data Type"`
	Enumeration   string             `csv:"Enumeration"`
	MinimumValue  int                `csv:"Minimum Value"`
	MaximumValue  int                `csv:"Maximum Value"`
	Description   string             `csv:"Description"`
}
