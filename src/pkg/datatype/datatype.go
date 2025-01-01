package datatype

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// DataTypeMap contains ALL records, including un-released and import-only records
	DataTypeMap DataTypeDefinitionMap

	// DataTypeListAll contains ALL records, including un-released and import-only records
	DataTypeListAll DataTypeDefinitionList

	// DataTypeList
	// is a filtered list.
	// It excludes un-released and import-only records.
	DataTypeList DataTypeDefinitionList
)

func init() {
	var err error
	DataTypeListAll, err = spec.ParseADISpecTSV[*DataTypeDefinition](spec.AdiDataTypesTSV)
	if err != nil {
		panic(err)
	}

	DataTypeList = make([]*DataTypeDefinition, 0, len(DataTypeListAll))
	for _, item := range DataTypeListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			DataTypeList = append(DataTypeList, item)
		}
	}
	DataTypeList = slices.Clip(DataTypeList)

	DataTypeMap = make(DataTypeDefinitionMap, len(DataTypeList))
	for _, item := range DataTypeList {
		DataTypeMap[item.ID] = item
	}
}

// DataTypeIndicator is a single character that represents the ADIF Data Type Indicator that precedes the data field in an ADIF record.
type DataTypeIndicator rune

// DataType represents the ADIF data type of a data field
type DataType string

// DataTypeDefinitionList represents a collection of ADIF data type definitions
type DataTypeDefinitionList []*DataTypeDefinition

// DataTypeDefinitionMap maps DataType to its definition
type DataTypeDefinitionMap map[DataType]*DataTypeDefinition

// DataTypeDefinition represents an ADIF data type definition
type DataTypeDefinition struct {
	shared.ImportRecordRoot
	ID           DataType          `csv:"Data Type Name"`
	Indicator    DataTypeIndicator `csv:"Data Type Indicator"`
	Description  string            `csv:"Description"`
	MinimumValue int               `csv:"Minimum Value"`
	MaximumValue int               `csv:"Maximum Value"`
	Comments     string            `csv:"Comments"`
}

// UnmarshalCSV implements the encoding.TextUnmarshaler interface for ADIFDataTypeIndicator
func (a *DataTypeIndicator) UnmarshalCSV(csv string) error {
	if len(csv) == 0 {
		*a = 0
		return nil
	}
	if len(csv) > 1 {
		return ErrDataTypeMustBeSingleCharacter
	}

	*a = DataTypeIndicator(csv[0])
	return nil
}
