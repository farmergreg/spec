package qsouploadstatus

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
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

// ADIF enums are case-insensitive.
func (q QSOUploadStatus) Compare(other QSOUploadStatus) int {
	return strings.Compare(string(q), string(other))
}
