package qsouploadstatus

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumQSOUploadStatusMap contains ALL records, including un-released and import-only statuses
	EnumQSOUploadStatusMap EnumQSOUploadStatusItemMap

	// EnumQSOUploadStatusListAll contains ALL records, including un-released and import-only statuses
	EnumQSOUploadStatusListAll EnumQSOUploadStatusItemList

	// EnumQSOUploadStatusList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumQSOUploadStatusList EnumQSOUploadStatusItemList
)

func init() {
	var err error
	EnumQSOUploadStatusListAll, err = spec.ParseADISpecTSV[*EnumQSOUploadStatusItem](spec.EnumQSOUploadStatusTSV)
	if err != nil {
		panic(err)
	}

	EnumQSOUploadStatusList = make([]*EnumQSOUploadStatusItem, 0, len(EnumQSOUploadStatusListAll))
	for _, item := range EnumQSOUploadStatusListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumQSOUploadStatusList = append(EnumQSOUploadStatusList, item)
		}
	}
	EnumQSOUploadStatusList = slices.Clip(EnumQSOUploadStatusList)

	EnumQSOUploadStatusMap = make(EnumQSOUploadStatusItemMap, len(EnumQSOUploadStatusList))
	for _, item := range EnumQSOUploadStatusList {
		EnumQSOUploadStatusMap[item.ID] = item
	}
}

// QSOUploadStatus represents the upload status of a QSO
type QSOUploadStatus string

// EnumQSOUploadStatusItemList represents a collection of QSO upload status items
type EnumQSOUploadStatusItemList []*EnumQSOUploadStatusItem

// EnumQSOUploadStatusItemMap maps QSOUploadStatus to its definition
type EnumQSOUploadStatusItemMap map[QSOUploadStatus]*EnumQSOUploadStatusItem

// EnumQSOUploadStatusItem represents a QSO upload status item
type EnumQSOUploadStatusItem struct {
	shared.ImportRecordRoot
	ID          QSOUploadStatus `csv:"Status"` // The value that is stored in the CLUBLOG_QSO_UPLOAD_STATUS, QRZCOM_QSO_UPLOAD_STATUS, HAMLOGEU_QSO_UPLOAD_STATUS, HAMQTH_QSO_UPLOAD_STATUS, HRDLOG_QSO_UPLOAD_STATUS ADIF field.
	Description string          `csv:"Description"`
}
