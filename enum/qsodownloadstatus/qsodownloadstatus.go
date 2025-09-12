package qsodownloadstatus

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// QSODownloadStatus represents the download status of a QSO
type QSODownloadStatus string

var _ codegen.CodeGenKey = QSODownloadStatus("")

// String returns the string representation of the QSODownloadStatus.
func (q QSODownloadStatus) String() string {
	return string(q)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t QSODownloadStatus) Compare(other QSODownloadStatus) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
