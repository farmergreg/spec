package qslrcvd

// QSLRcvd represents the QSL received status
type QSLRcvd string

func (q QSLRcvd) String() string {
	return string(q)
}
