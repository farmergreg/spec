package credit

// Credit represents an award credit identifier
type Credit string

func (c Credit) String() string {
	return string(c)
}
