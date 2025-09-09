package qsouploadstatus

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// QSOUploadStatus represents the download status of a QSO
type QSOUploadStatus string

var _ codegen.CodeGenKey = QSOUploadStatus("")

// String returns the string representation of the QSOUploadStatus.
func (q QSOUploadStatus) String() string {
	return string(q)
}
