package awardsponsor

import (
	"strings"

	"github.com/farmergreg/spec/v6/internal/codegen"
)

// AwardSponsorPrefix represents the unique identifier for an award sponsor
type AwardSponsorPrefix string

var _ codegen.CodeGenKey = AwardSponsorPrefix("")

// New creates a new AwardSponsorPrefix from the provided string.
func New(value string) AwardSponsorPrefix {
	return AwardSponsorPrefix(strings.ToLower(value))
}

// String returns the string representation of the AwardSponsorPrefix.
func (a AwardSponsorPrefix) String() string {
	return string(a)
}

// Compare returns an integer comparing two AwardSponsorPrefix values lexicographically.
// ADIF enums are case-insensitive.
func (t AwardSponsorPrefix) Compare(other AwardSponsorPrefix) int {
	return strings.Compare(strings.ToLower(string(t)), strings.ToLower(string(other)))
}

// Equals returns true if this AwardSponsorPrefix equals the other AwardSponsorPrefix.
// ADIF enums are case-insensitive.
func (t AwardSponsorPrefix) Equals(other AwardSponsorPrefix) bool {
	return strings.EqualFold(string(t), string(other))
}
