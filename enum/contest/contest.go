package contest

// Contest represents the contest identifier
type Contest string

func (c Contest) String() string {
	return string(c)
}
