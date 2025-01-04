package credit

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumCreditMap contains ALL records, including un-released and import-only records
	EnumCreditMap map[Credit]*EnumCreditItem

	// EnumCreditListAll contains ALL records, including un-released and import-only records
	EnumCreditListAll []*EnumCreditItem

	// EnumCreditList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumCreditList []*EnumCreditItem
)

func init() {
	var err error
	EnumCreditListAll, err = spec.ParseADISpecTSV[*EnumCreditItem](spec.EnumCreditTSV)
	if err != nil {
		panic(err)
	}

	EnumCreditList = make([]*EnumCreditItem, 0, len(EnumCreditListAll))
	for _, item := range EnumCreditListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumCreditList = append(EnumCreditList, item)
		}
	}
	EnumCreditList = slices.Clip(EnumCreditList)

	EnumCreditMap = make(map[Credit]*EnumCreditItem, len(EnumCreditList))
	for _, item := range EnumCreditList {
		EnumCreditMap[item.ID] = item
	}
}

// Credit represents an award credit identifier
type Credit string

// EnumCreditItem represents an award credit item
type EnumCreditItem struct {
	shared.ImportRecordRoot
	ID      Credit `csv:"Credit For"` // The value that is stored in the CREDIT_SUBMITTED and CREDIT_GRANTED ADIF fields..
	Sponsor string `csv:"Sponsor"`
	Award   string `csv:"Award"`
	Facet   string `csv:"Facet"`
}
