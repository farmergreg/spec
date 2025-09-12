package awardsponsor

import (
	"strings"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// AwardSponsorPrefix represents the unique identifier for an award sponsor
type AwardSponsorPrefix string

var _ codegen.CodeGenKey = AwardSponsorPrefix("")

// New creates a new AwardSponsorPrefix from the provided string.
func New(value string) AwardSponsorPrefix {
	return AwardSponsorPrefix(strings.ToUpper(value))
}

// String returns the string representation of the AwardSponsorPrefix.
func (a AwardSponsorPrefix) String() string {
	return string(a)
}

// ADIF enums are case-insensitive.
func (t AwardSponsorPrefix) Compare(other AwardSponsorPrefix) int {
	return strings.Compare(string(t), string(other))
}
