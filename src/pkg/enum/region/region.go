package region

// Region represents a region entity code
type Region string

func (r Region) String() string {
	return string(r)
}
