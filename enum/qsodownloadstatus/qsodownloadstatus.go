package qsodownloadstatus

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSODownloadStatus represents the download status of a QSO
type QSODownloadStatus string

var _ codegen.CodeGenKey = QSODownloadStatus("")

// New creates a new QSODownloadStatus from the provided string.
func New(value string) QSODownloadStatus {
	return QSODownloadStatus(strings.ToUpper(value))
}

// String returns the string representation of the QSODownloadStatus.
func (q QSODownloadStatus) String() string {
	return string(q)
}

// ADIF enums are case-insensitive.
func (q QSODownloadStatus) Compare(other QSODownloadStatus) int {
	return strings.Compare(string(q), string(other))
}
