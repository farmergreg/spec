package qsouploadstatus

import (
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
)

// QSOUploadStatus represents the download status of a QSO
type QSOUploadStatus string

var _ codegen.CodeGenKey = QSOUploadStatus("")

// New creates a new QSOUploadStatus from the provided string.
func New(value string) QSOUploadStatus {
	return QSOUploadStatus(strings.ToUpper(value))
}

// String returns the string representation of the QSOUploadStatus.
func (q QSOUploadStatus) String() string {
	return string(q)
}

// Compare returns an integer comparing two QSOUploadStatus values lexicographically.
// ADIF enums are case-insensitive.
func (q QSOUploadStatus) Compare(other QSOUploadStatus) int {
	return strings.Compare(strings.ToUpper(string(q)), strings.ToUpper(string(other)))
}

// Equals returns true if this QSOUploadStatus equals the other QSOUploadStatus.
// ADIF enums are case-insensitive.
func (q QSOUploadStatus) Equals(other QSOUploadStatus) bool {
	return strings.EqualFold(string(q), string(other))
}
