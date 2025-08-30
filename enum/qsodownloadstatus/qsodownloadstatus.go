package qsodownloadstatus

// QSODownloadStatus represents the download status of a QSO
type QSODownloadStatus string

func (q QSODownloadStatus) String() string {
	return string(q)
}
