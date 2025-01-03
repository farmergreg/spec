package qslmedium

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumQSLMediumMap contains ALL records, including un-released and import-only records
	EnumQSLMediumMap map[QSLMedium]*EnumQSLMediumItem

	// EnumQSLMediumListAll contains ALL records, including un-released and import-only records
	EnumQSLMediumListAll []*EnumQSLMediumItem

	// EnumQSLMediumList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumQSLMediumList []*EnumQSLMediumItem
)

func init() {
	var err error
	EnumQSLMediumListAll, err = spec.ParseADISpecTSV[*EnumQSLMediumItem](spec.EnumQSLMediumTSV)
	if err != nil {
		panic(err)
	}

	EnumQSLMediumList = make([]*EnumQSLMediumItem, 0, len(EnumQSLMediumListAll))
	for _, item := range EnumQSLMediumListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumQSLMediumList = append(EnumQSLMediumList, item)
		}
	}
	EnumQSLMediumList = slices.Clip(EnumQSLMediumList)

	EnumQSLMediumMap = make(map[QSLMedium]*EnumQSLMediumItem, len(EnumQSLMediumList))
	for _, item := range EnumQSLMediumList {
		EnumQSLMediumMap[item.ID] = item
	}
}

// QSLMedium represents the medium used for QSL exchange
type QSLMedium string

// EnumQSLMediumItem represents a QSL medium item
type EnumQSLMediumItem struct {
	shared.ImportRecordRoot
	ID          QSLMedium `csv:"Medium"`
	Description string    `csv:"Description"`
}
