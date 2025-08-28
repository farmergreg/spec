// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package awardsponsor

var (
	ADIF  AwardSponsorPrefix = "ADIF_"  // ADIF Development Group
	ARI   AwardSponsorPrefix = "ARI_"   // ARI - l'Associazione Radioamatori Italiani
	ARRL  AwardSponsorPrefix = "ARRL_"  // ARRL - American Radio Relay League
	CQ    AwardSponsorPrefix = "CQ_"    // CQ Magazine
	DARC  AwardSponsorPrefix = "DARC_"  // DARC - Deutscher Amateur-Radio-Club e.V.
	EQSL  AwardSponsorPrefix = "EQSL_"  // eQSL
	IARU  AwardSponsorPrefix = "IARU_"  // IARU - International Amateur Radio Union
	JARL  AwardSponsorPrefix = "JARL_"  // JARL - Japan Amateur Radio League
	RSGB  AwardSponsorPrefix = "RSGB_"  // RSGB - Radio Society of Great Britain
	TAG   AwardSponsorPrefix = "TAG_"   // TAG - Tambov award group
	WABAG AwardSponsorPrefix = "WABAG_" // WAB - Worked all Britain
)

// IsValid returns true if the AwardSponsorPrefix exists in the ADIF specification. This includes Import Only and Deleted values.
func (value AwardSponsorPrefix) IsValid() bool {
	switch value {
	case ADIF:
		return true
	case ARI:
		return true
	case ARRL:
		return true
	case CQ:
		return true
	case DARC:
		return true
	case EQSL:
		return true
	case IARU:
		return true
	case JARL:
		return true
	case RSGB:
		return true
	case TAG:
		return true
	case WABAG:
		return true
	}
	return false
}
