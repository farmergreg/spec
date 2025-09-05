package qsodownloadstatus

import "github.com/hamradiolog-net/adif-spec/v6/codegen"

// QSODownloadStatus represents the download status of a QSO
type QSODownloadStatus string

var _ codegen.CodeGeneratorEnumValue = QSODownloadStatus("")

func (q QSODownloadStatus) String() string {
	return string(q)
}
