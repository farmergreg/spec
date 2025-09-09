package qsodownloadstatus

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// QSODownloadStatus represents the download status of a QSO
type QSODownloadStatus string

var _ codegen.CodeGenKey = QSODownloadStatus("")

// String returns the string representation of the QSODownloadStatus.
func (q QSODownloadStatus) String() string {
	return string(q)
}
