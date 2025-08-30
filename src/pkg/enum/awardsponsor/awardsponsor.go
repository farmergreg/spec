package awardsponsor

// AwardSponsorPrefix represents the unique identifier for an award sponsor
type AwardSponsorPrefix string

func (a AwardSponsorPrefix) String() string {
	return string(a)
}
