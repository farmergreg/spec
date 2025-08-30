package mode

// Mode is the ADIF mode of a radio communication.
type Mode string

func (m Mode) String() string {
	return string(m)
}
