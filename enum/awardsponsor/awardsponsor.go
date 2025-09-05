package awardsponsor

import "github.com/hamradiolog-net/adif-spec/v6/codegen"

// AwardSponsorPrefix represents the unique identifier for an award sponsor
type AwardSponsorPrefix string

var _ codegen.CodeGeneratorEnumValue = AwardSponsorPrefix("")

func (a AwardSponsorPrefix) String() string {
	return string(a)
}
