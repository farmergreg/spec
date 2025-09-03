package band

// Band represents a range of radio frequencies
type Band string

func (b Band) String() string {
	return string(b)
}
