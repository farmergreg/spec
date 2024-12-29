package qsodownloadstatus

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/spec"
)

var (
	// EnumQSODownloadStatusMap contains ALL records, including un-released and import-only statuses
	EnumQSODownloadStatusMap EnumQSODownloadStatusItemMap

	// EnumQSODownloadStatusListAll contains ALL records, including un-released and import-only statuses
	EnumQSODownloadStatusListAll EnumQSODownloadStatusItemList

	// EnumQSODownloadStatusList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumQSODownloadStatusList EnumQSODownloadStatusItemList
)

func init() {
	var err error
	EnumQSODownloadStatusListAll, err = spec.ParseADISpecTSV[*EnumQSODownloadStatusItem](spec.EnumQSODownloadStatusTSV)
	if err != nil {
		panic(err)
	}

	EnumQSODownloadStatusList = make([]*EnumQSODownloadStatusItem, 0, len(EnumQSODownloadStatusListAll))
	for _, item := range EnumQSODownloadStatusListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumQSODownloadStatusList = append(EnumQSODownloadStatusList, item)
		}
	}
	EnumQSODownloadStatusList = slices.Clip(EnumQSODownloadStatusList)

	EnumQSODownloadStatusMap = make(EnumQSODownloadStatusItemMap, len(EnumQSODownloadStatusList))
	for _, item := range EnumQSODownloadStatusList {
		EnumQSODownloadStatusMap[item.ID] = item
	}
}

// QSODownloadStatus represents the download status of a QSO
type QSODownloadStatus string

// EnumQSODownloadStatusItemList represents a collection of QSO download status items
type EnumQSODownloadStatusItemList []*EnumQSODownloadStatusItem

// EnumQSODownloadStatusItemMap maps QSODownloadStatus to its definition
type EnumQSODownloadStatusItemMap map[QSODownloadStatus]*EnumQSODownloadStatusItem

// EnumQSODownloadStatusItem represents a QSO download status item
type EnumQSODownloadStatusItem struct {
	shared.ImportRecordRoot
	ID          QSODownloadStatus `csv:"Status"` // The value that is stored in the QRZCOM_QSO_DOWNLOAD_STATUS ADIF field.
	Description string            `csv:"Description"`
}
