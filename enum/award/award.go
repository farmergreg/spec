package award

// Deprecated: Award represents an ADIF award type
type Award string

func (a Award) String() string {
	return string(a)
}
