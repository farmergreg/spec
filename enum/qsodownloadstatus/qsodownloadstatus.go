package qsodownloadstatus

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// QSODownloadStatus represents the download status of a QSO
type QSODownloadStatus string

var _ codegen.CodeGenKey = QSODownloadStatus("")

func (q QSODownloadStatus) String() string {
	return string(q)
}
