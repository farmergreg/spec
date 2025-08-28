package qslmedium

// QSLMedium represents the medium used for QSL exchange
type QSLMedium string

func (q QSLMedium) String() string {
	return string(q)
}
