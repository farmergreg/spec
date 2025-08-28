package qsocomplete

// QSOComplete represents the completion status of a QSO
type QSOComplete string

func (q QSOComplete) String() string {
	return string(q)
}
