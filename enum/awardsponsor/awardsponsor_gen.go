// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package awardsponsor provides code and constants as defined in ADIF 3.1.6 (Released)
package awardsponsor

import "sync"

const (
	ADIF  AwardSponsorPrefix = "adif_"  // adif_  = ADIF Development Group
	ARI   AwardSponsorPrefix = "ari_"   // ari_   = ARI - l'Associazione Radioamatori Italiani
	ARRL  AwardSponsorPrefix = "arrl_"  // arrl_  = ARRL - American Radio Relay League
	CQ    AwardSponsorPrefix = "cq_"    // cq_    = CQ Magazine
	DARC  AwardSponsorPrefix = "darc_"  // darc_  = DARC - Deutscher Amateur-Radio-Club e.V.
	EQSL  AwardSponsorPrefix = "eqsl_"  // eqsl_  = eQSL
	IARU  AwardSponsorPrefix = "iaru_"  // iaru_  = IARU - International Amateur Radio Union
	JARL  AwardSponsorPrefix = "jarl_"  // jarl_  = JARL - Japan Amateur Radio League
	RSGB  AwardSponsorPrefix = "rsgb_"  // rsgb_  = RSGB - Radio Society of Great Britain
	TAG   AwardSponsorPrefix = "tag_"   // tag_   = TAG - Tambov award group
	WABAG AwardSponsorPrefix = "wabag_" // wabag_ = WAB - Worked all Britain
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known AwardSponsorPrefix specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "adif_", Description: "ADIF Development Group"},
	{IsImportOnly: false, Key: "ari_", Description: "ARI - l'Associazione Radioamatori Italiani"},
	{IsImportOnly: false, Key: "arrl_", Description: "ARRL - American Radio Relay League"},
	{IsImportOnly: false, Key: "cq_", Description: "CQ Magazine"},
	{IsImportOnly: false, Key: "darc_", Description: "DARC - Deutscher Amateur-Radio-Club e.V."},
	{IsImportOnly: false, Key: "eqsl_", Description: "eQSL"},
	{IsImportOnly: false, Key: "iaru_", Description: "IARU - International Amateur Radio Union"},
	{IsImportOnly: false, Key: "jarl_", Description: "JARL - Japan Amateur Radio League"},
	{IsImportOnly: false, Key: "rsgb_", Description: "RSGB - Radio Society of Great Britain"},
	{IsImportOnly: false, Key: "tag_", Description: "TAG - Tambov award group"},
	{IsImportOnly: false, Key: "wabag_", Description: "WAB - Worked all Britain"},
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
