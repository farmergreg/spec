// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package awardsponsor provides code and constants as defined in ADIF 3.1.6 (Proposed)
package awardsponsor

const (
	ADIF  AwardSponsorPrefix = "ADIF_"  // ADIF_  = ADIF Development Group
	ARI   AwardSponsorPrefix = "ARI_"   // ARI_   = ARI - l'Associazione Radioamatori Italiani
	ARRL  AwardSponsorPrefix = "ARRL_"  // ARRL_  = ARRL - American Radio Relay League
	CQ    AwardSponsorPrefix = "CQ_"    // CQ_    = CQ Magazine
	DARC  AwardSponsorPrefix = "DARC_"  // DARC_  = DARC - Deutscher Amateur-Radio-Club e.V.
	EQSL  AwardSponsorPrefix = "EQSL_"  // EQSL_  = eQSL
	IARU  AwardSponsorPrefix = "IARU_"  // IARU_  = IARU - International Amateur Radio Union
	JARL  AwardSponsorPrefix = "JARL_"  // JARL_  = JARL - Japan Amateur Radio League
	RSGB  AwardSponsorPrefix = "RSGB_"  // RSGB_  = RSGB - Radio Society of Great Britain
	TAG   AwardSponsorPrefix = "TAG_"   // TAG_   = TAG - Tambov award group
	WABAG AwardSponsorPrefix = "WABAG_" // WABAG_ = WAB - Worked all Britain
)

// Lookup look up a specification for AwardSponsorPrefix
func Lookup(awardsponsorprefix AwardSponsorPrefix) (Spec, bool) {
	spec, ok := internalMap[awardsponsorprefix]
	return spec, ok
}

// IsValid returns true if the specification for AwardSponsorPrefix exists and is not import only.
func IsValid(awardsponsorprefix AwardSponsorPrefix) bool {
	spec, ok := internalMap[awardsponsorprefix]
	if ok && bool(spec.IsImportOnly) {
		return false
	}
	return ok
}

// All AwardSponsorPrefix specifications INCLUDING ones marked import only.
func AllAwardSponsorPrefix() []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		result = append(result, v)
	}
	return result
}

// AllActiveAwardSponsorPrefix specifications EXCLUDING ones marked import only.
func AllActiveAwardSponsorPrefix() []Spec {
	return LookupByFilter(func(s Spec) bool {
		return !bool(s.IsImportOnly)
	})
}

// LookupByFilter returns all specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

var internalMap = map[AwardSponsorPrefix]Spec{
	ADIF:  {IsImportOnly: false, Key: "ADIF_", Description: "ADIF Development Group"},
	ARI:   {IsImportOnly: false, Key: "ARI_", Description: "ARI - l'Associazione Radioamatori Italiani"},
	ARRL:  {IsImportOnly: false, Key: "ARRL_", Description: "ARRL - American Radio Relay League"},
	CQ:    {IsImportOnly: false, Key: "CQ_", Description: "CQ Magazine"},
	DARC:  {IsImportOnly: false, Key: "DARC_", Description: "DARC - Deutscher Amateur-Radio-Club e.V."},
	EQSL:  {IsImportOnly: false, Key: "EQSL_", Description: "eQSL"},
	IARU:  {IsImportOnly: false, Key: "IARU_", Description: "IARU - International Amateur Radio Union"},
	JARL:  {IsImportOnly: false, Key: "JARL_", Description: "JARL - Japan Amateur Radio League"},
	RSGB:  {IsImportOnly: false, Key: "RSGB_", Description: "RSGB - Radio Society of Great Britain"},
	TAG:   {IsImportOnly: false, Key: "TAG_", Description: "TAG - Tambov award group"},
	WABAG: {IsImportOnly: false, Key: "WABAG_", Description: "WAB - Worked all Britain"},
}
