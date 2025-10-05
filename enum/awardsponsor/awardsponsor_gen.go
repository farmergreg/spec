// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package awardsponsor provides code and constants as defined in ADIF 3.1.6 (Released)
package awardsponsor

import "sync"

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

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known AwardSponsorPrefix specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "ADIF_", Description: "ADIF Development Group"},
	{IsImportOnly: false, Key: "ARI_", Description: "ARI - l'Associazione Radioamatori Italiani"},
	{IsImportOnly: false, Key: "ARRL_", Description: "ARRL - American Radio Relay League"},
	{IsImportOnly: false, Key: "CQ_", Description: "CQ Magazine"},
	{IsImportOnly: false, Key: "DARC_", Description: "DARC - Deutscher Amateur-Radio-Club e.V."},
	{IsImportOnly: false, Key: "EQSL_", Description: "eQSL"},
	{IsImportOnly: false, Key: "IARU_", Description: "IARU - International Amateur Radio Union"},
	{IsImportOnly: false, Key: "JARL_", Description: "JARL - Japan Amateur Radio League"},
	{IsImportOnly: false, Key: "RSGB_", Description: "RSGB - Radio Society of Great Britain"},
	{IsImportOnly: false, Key: "TAG_", Description: "TAG - Tambov award group"},
	{IsImportOnly: false, Key: "WABAG_", Description: "WAB - Worked all Britain"},
}

// lookupMap contains all known AwardSponsorPrefix specifications.
var lookupMap = map[AwardSponsorPrefix]*Spec{
	ADIF:  &lookupList[0],
	ARI:   &lookupList[1],
	ARRL:  &lookupList[2],
	CQ:    &lookupList[3],
	DARC:  &lookupList[4],
	EQSL:  &lookupList[5],
	IARU:  &lookupList[6],
	JARL:  &lookupList[7],
	RSGB:  &lookupList[8],
	TAG:   &lookupList[9],
	WABAG: &lookupList[10],
}

// Lookup returns the specification for the provided AwardSponsorPrefix.
// ADIF 3.1.6
func Lookup(a AwardSponsorPrefix) (Spec, bool) {
	spec, ok := lookupMap[a]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all AwardSponsorPrefix specifications that match the provided filter function.
// ADIF 3.1.6
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(lookupList))
	for _, v := range lookupList {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// List returns all AwardSponsorPrefix specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns AwardSponsorPrefix specifications.
// This list excludes those marked as import-only.
// ADIF 3.1.6
func ListActive() []Spec {
	listActiveOnce.Do(func() {
		listActive = LookupByFilter(func(spec Spec) bool { return !bool(spec.IsImportOnly) })
	})

	result := make([]Spec, len(listActive))
	copy(result, listActive)
	return result
}
