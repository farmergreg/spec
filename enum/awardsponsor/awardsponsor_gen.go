// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package awardsponsor provides code and constants as defined in ADIF 3.1.6 (Proposed)
package awardsponsor

import "maps"

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

// All AwardSponsorPrefix specifications including depreciated and import only.
func AwardSponsorPrefixListAll() []Spec {
	return append([]Spec(nil), internalAwardSponsorPrefixListAll...)
}

// All AwardSponsorPrefix specifications values that are NOT marked import-only.
func AwardSponsorPrefixListCurrent() []Spec {
	return append([]Spec(nil), internalAwardSponsorPrefixListCurrent...)
}

// A map of all AwardSponsorPrefix specifications.
func AwardSponsorPrefixMap() map[AwardSponsorPrefix]Spec {
	cp := make(map[AwardSponsorPrefix]Spec, len(internalAwardSponsorPrefixMap))
	maps.Copy(cp, internalAwardSponsorPrefixMap)
	return cp
}

// A map of all AwardSponsorPrefix specifications.
var internalAwardSponsorPrefixMap = map[AwardSponsorPrefix]Spec{
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

var internalAwardSponsorPrefixListAll = []Spec{
	internalAwardSponsorPrefixMap[ADIF],
	internalAwardSponsorPrefixMap[ARI],
	internalAwardSponsorPrefixMap[ARRL],
	internalAwardSponsorPrefixMap[CQ],
	internalAwardSponsorPrefixMap[DARC],
	internalAwardSponsorPrefixMap[EQSL],
	internalAwardSponsorPrefixMap[IARU],
	internalAwardSponsorPrefixMap[JARL],
	internalAwardSponsorPrefixMap[RSGB],
	internalAwardSponsorPrefixMap[TAG],
	internalAwardSponsorPrefixMap[WABAG],
}

var internalAwardSponsorPrefixListCurrent = []Spec{
	internalAwardSponsorPrefixMap[ADIF],
	internalAwardSponsorPrefixMap[ARI],
	internalAwardSponsorPrefixMap[ARRL],
	internalAwardSponsorPrefixMap[CQ],
	internalAwardSponsorPrefixMap[DARC],
	internalAwardSponsorPrefixMap[EQSL],
	internalAwardSponsorPrefixMap[IARU],
	internalAwardSponsorPrefixMap[JARL],
	internalAwardSponsorPrefixMap[RSGB],
	internalAwardSponsorPrefixMap[TAG],
	internalAwardSponsorPrefixMap[WABAG],
}
