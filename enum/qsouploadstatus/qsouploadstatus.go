package qsouploadstatus

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// QSOUploadStatus represents the download status of a QSO
type QSOUploadStatus string

var _ codegen.CodeGeneratorEnumValue = QSOUploadStatus("")

func (q QSOUploadStatus) String() string {
	return string(q)
}
