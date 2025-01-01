package propagationmode

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumPropagationModeMap contains ALL records, including un-released and import-only records
	EnumPropagationModeMap EnumPropagationModeItemMap

	// EnumPropagationModeListAll contains ALL records, including un-released and import-only records
	EnumPropagationModeListAll EnumPropagationModeItemList

	// EnumPropagationModeList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumPropagationModeList EnumPropagationModeItemList
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

	EnumPropagationModeMap = make(EnumPropagationModeItemMap, len(EnumPropagationModeList))
	for _, item := range EnumPropagationModeList {
		EnumPropagationModeMap[item.ID] = item
	}
}

// PropagationMode represents the propagation mode string type
type PropagationMode string

// EnumPropagationModeItemList represents a collection of propagation mode items
type EnumPropagationModeItemList []*EnumPropagationModeItem

// EnumPropagationModeItemMap maps PropagationMode to its definition
type EnumPropagationModeItemMap map[PropagationMode]*EnumPropagationModeItem

// EnumPropagationModeItem represents a propagation mode item
type EnumPropagationModeItem struct {
	shared.ImportRecordRoot
	ID          PropagationMode `csv:"Enumeration"` // The value that is stored in the PROPAGATION_MODE ADIF field.
	Description string          `csv:"Description"`
}
