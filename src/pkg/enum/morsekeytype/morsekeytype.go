package morsekeytype

type MorseKeyType string

func (m MorseKeyType) String() string {
	return string(m)
}
