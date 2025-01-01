package adifield

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// FieldMap contains ALL records, including un-released and import-only records
	FieldMap FieldDefinitionMap

	// FieldListAll contains ALL records, including un-released and import-only records
	FieldListAll FieldDefinitionList

	// FieldList
	// is a filtered list.
	// It excludes un-released and import-only records.
	FieldList FieldDefinitionList
)

func init() {
	var err error
	FieldListAll, err = spec.ParseADISpecTSV[*FieldDefinition](spec.FieldsTSV)
	if err != nil {
		panic(err)
	}

	// special case, rename USERDEFn in FieldListAll to USERDEF1
	for _, item := range FieldListAll {
		if item.ID == USERDEF+"n" {
			item.ID = USERDEF + "1"
			break
		}
	}

	FieldList = make([]*FieldDefinition, 0, len(FieldListAll))
	for _, item := range FieldListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			FieldList = append(FieldList, item)
		}
	}
	FieldList = slices.Clip(FieldList)

	FieldMap = make(FieldDefinitionMap, len(FieldList))
	for _, item := range FieldList {
		FieldMap[item.ID] = item
	}
}

const (
	EOR Field = "EOR" // Special Field: End of Record
	EOH Field = "EOH" // Special Field: End of Header

	// USERDEF is the prefix for user defined fields.
	// These fields must always have a number after them, e.g. USERDEF1, USERDEF2, etc.
	USERDEF = "USERDEF"

	// APP_ is the prefix for application defined fields.
	// These field follow the pattern: APP_PROGRAMID_FIELDNAME
	// For example, APP_K9CTS_TEST_FIELD
	APP_ = "APP_"
)

// Field is the ADIF field name in and ADI file.
type Field string

// FieldDefinitionList represents a collection of ADIF field definitions
type FieldDefinitionList []*FieldDefinition

// FieldDefinitionMap maps Field to its definition
type FieldDefinitionMap map[Field]*FieldDefinition

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
