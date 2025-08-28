package continent

// Continent represents a continent such as EU for Europe
type Continent string

func (c Continent) String() string {
	return string(c)
}
