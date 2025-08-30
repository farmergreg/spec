package qsouploadstatus

// QSOUploadStatus represents the download status of a QSO
type QSOUploadStatus string

func (q QSOUploadStatus) String() string {
	return string(q)
}
