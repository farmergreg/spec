// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package qsouploadstatus

var (
	M QSOUploadStatus = "M" // the QSO has been modified since being uploaded to the online service
	N QSOUploadStatus = "N" // do not upload the QSO to the online service
	Y QSOUploadStatus = "Y" // the QSO has been uploaded to, and accepted by, the online service
)

// IsValid returns true if the QSOUploadStatus exists in the ADIF specification. This includes Import Only and Deleted values.
func (value QSOUploadStatus) IsValid() bool {
	switch value {
	case M:
		return true
	case N:
		return true
	case Y:
		return true
	}
	return false
}
