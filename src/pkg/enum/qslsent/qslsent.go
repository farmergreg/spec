package qslsent

// QSLSent represents the QSL sent status
type QSLSent string

func (q QSLSent) String() string {
	return string(q)
}
