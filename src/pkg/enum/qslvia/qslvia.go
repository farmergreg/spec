package qslvia

// QSLVia represents the method used to exchange QSL cards
type QSLVia string

func (q QSLVia) String() string {
	return string(q)
}
