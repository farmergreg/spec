package qsodownloadstatus

import (
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
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

// Compare returns an integer comparing two QSODownloadStatus values lexicographically.
// ADIF enums are case-insensitive.
func (q QSODownloadStatus) Compare(other QSODownloadStatus) int {
	return strings.Compare(strings.ToUpper(string(q)), strings.ToUpper(string(other)))
}

// Equals returns true if this QSODownloadStatus equals the other QSODownloadStatus.
// ADIF enums are case-insensitive.
func (q QSODownloadStatus) Equals(other QSODownloadStatus) bool {
	return strings.EqualFold(string(q), string(other))
}
