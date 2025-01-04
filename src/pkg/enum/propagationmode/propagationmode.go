package propagationmode

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumPropagationModeMap contains ALL records, including un-released and import-only records
	EnumPropagationModeMap map[PropagationMode]*EnumPropagationModeItem

	// EnumPropagationModeListAll contains ALL records, including un-released and import-only records
	EnumPropagationModeListAll []*EnumPropagationModeItem

	// EnumPropagationModeList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumPropagationModeList []*EnumPropagationModeItem
)

func init() {
	var err error
	EnumPropagationModeListAll, err = spec.ParseADISpecTSV[*EnumPropagationModeItem](spec.EnumPropagationModeTSV)
	if err != nil {
		panic(err)
	}

	EnumPropagationModeList = make([]*EnumPropagationModeItem, 0, len(EnumPropagationModeListAll))
	for _, item := range EnumPropagationModeListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumPropagationModeList = append(EnumPropagationModeList, item)
		}
	}
	EnumPropagationModeList = slices.Clip(EnumPropagationModeList)

	EnumPropagationModeMap = make(map[PropagationMode]*EnumPropagationModeItem, len(EnumPropagationModeList))
	for _, item := range EnumPropagationModeList {
		EnumPropagationModeMap[item.ID] = item
	}
}

// PropagationMode represents the propagation mode string type
type PropagationMode string

// EnumPropagationModeItem represents a propagation mode item
type EnumPropagationModeItem struct {
	shared.ImportRecordRoot
	ID          PropagationMode `csv:"Enumeration"` // The value that is stored in the PROPAGATION_MODE ADIF field.
	Description string          `csv:"Description"`
}
