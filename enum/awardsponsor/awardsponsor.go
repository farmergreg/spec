package awardsponsor

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// AwardSponsorPrefix represents the unique identifier for an award sponsor
type AwardSponsorPrefix string

var _ codegen.CodeGenKey = AwardSponsorPrefix("")

// String returns the string representation of the AwardSponsorPrefix.
func (a AwardSponsorPrefix) String() string {
	return string(a)
}

// Compare implements the Comparable interface.
// ADIF enums are case-insensitive.
func (t AwardSponsorPrefix) Compare(other AwardSponsorPrefix) int {
	return strings.Compare(strings.ToUpper(string(t)), strings.ToUpper(other.String()))
}
