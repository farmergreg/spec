package band

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumBandMap contains ALL records, including un-released and import-only records
	EnumBandMap EnumBandItemMap

	// EnumBandListAll contains ALL records, including un-released and import-only records
	EnumBandListAll EnumBandItemList

	// EnumBandList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumBandList EnumBandItemList
)

func init() {
	var err error
	EnumBandListAll, err = spec.ParseADISpecTSV[*EnumBandItem](spec.EnumBandTSV)
	if err != nil {
		panic(err)
	}

	EnumBandList = make([]*EnumBandItem, 0, len(EnumBandListAll))
	for _, item := range EnumBandListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumBandList = append(EnumBandList, item)
		}
	}
	EnumBandList = slices.Clip(EnumBandList)

	EnumBandMap = make(EnumBandItemMap, len(EnumBandList))
	for _, item := range EnumBandList {
		EnumBandMap[item.ID] = item
	}
}

// Band represents a ham radio band
type Band string

// MHz represents megahertz, a type alias for float32
type MHz = float32

// EnumBandItemList represents a collection of Band definitions
type EnumBandItemList []*EnumBandItem

// EnumBandItemMap maps Band names to their definitions
type EnumBandItemMap map[Band]*EnumBandItem

// EnumBandItem represents a band item
type EnumBandItem struct {
	shared.ImportRecordRoot
	ID           Band `csv:"Band"`             // The value that is stored in the BAND ADIF field.
	LowerFreqMHz MHz  `csv:"Lower Freq (MHz)"` // The lower frequency limit for the band.
	UpperFreqMHz MHz  `csv:"Upper Freq (MHz)"` // The upper frequency limit for the band.
}

func (b *EnumBandItem) IsInBand(mhz MHz) bool {
	return mhz >= b.LowerFreqMHz && mhz <= b.UpperFreqMHz
}

// FindBandByMHz returns the band that contains the given MHz value, if any.
// It only searches the current band list.
func FindBandByMHz(mhz MHz) (*EnumBandItem, bool) {
	for _, item := range EnumBandList {
		if item.IsInBand(mhz) {
			return item, true
		}
	}

	// Not found.
	return nil, false
}

// Bandwidth returns the width of the frequency range in MHz
func (b *EnumBandItem) Bandwidth() MHz {
	return b.UpperFreqMHz - b.LowerFreqMHz
}
