package awardsponsor

import "github.com/hamradiolog-net/spec/v6/internal/codegen"

// AwardSponsorPrefix represents the unique identifier for an award sponsor
type AwardSponsorPrefix string

var _ codegen.CodeGenKey = AwardSponsorPrefix("")

// String returns the string representation of the AwardSponsorPrefix.
func (a AwardSponsorPrefix) String() string {
	return string(a)
}
