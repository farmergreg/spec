package arrlsection

// ARRLSection represents an ARRL section identifier
type ARRLSection string

func (a ARRLSection) String() string {
	return string(a)
}
