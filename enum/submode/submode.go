package submode

// SubMode represents the submode of an ADIF record
type SubMode string

func (s SubMode) String() string {
	return string(s)
}
