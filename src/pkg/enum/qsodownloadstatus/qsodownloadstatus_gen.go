// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package qsodownloadstatus

var (
	I QSODownloadStatus = "I" // ignore or invalid
	N QSODownloadStatus = "N" // the QSO has not been downloaded from the online service
	Y QSODownloadStatus = "Y" // the QSO has been downloaded from the online service
)

// IsValid returns true if the QSODownloadStatus exists in the ADIF specification. This includes Import Only and Deleted values.
func (value QSODownloadStatus) IsValid() bool {
	switch value {
	case I:
		return true
	case N:
		return true
	case Y:
		return true
	}
	return false
}
