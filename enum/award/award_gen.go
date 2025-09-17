// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package award provides code and constants as defined in ADIF 3.1.6 (Released)
package award

import "sync"

const (
	AJA         Award = "aja"         // Deprecated: aja
	CQDX        Award = "cqdx"        // Deprecated: cqdx
	CQDXFIELD   Award = "cqdxfield"   // Deprecated: cqdxfield
	CQWAZ_160M  Award = "cqwaz_160m"  // Deprecated: cqwaz_160m
	CQWAZ_CW    Award = "cqwaz_cw"    // Deprecated: cqwaz_cw
	CQWAZ_MIXED Award = "cqwaz_mixed" // Deprecated: cqwaz_mixed
	CQWAZ_PHONE Award = "cqwaz_phone" // Deprecated: cqwaz_phone
	CQWAZ_RTTY  Award = "cqwaz_rtty"  // Deprecated: cqwaz_rtty
	CQWPX       Award = "cqwpx"       // Deprecated: cqwpx
	DARC_DOK    Award = "darc_dok"    // Deprecated: darc_dok
	DXCC        Award = "dxcc"        // Deprecated: dxcc
	DXCC_CW     Award = "dxcc_cw"     // Deprecated: dxcc_cw
	DXCC_MIXED  Award = "dxcc_mixed"  // Deprecated: dxcc_mixed
	DXCC_PHONE  Award = "dxcc_phone"  // Deprecated: dxcc_phone
	DXCC_RTTY   Award = "dxcc_rtty"   // Deprecated: dxcc_rtty
	IOTA        Award = "iota"        // Deprecated: iota
	JCC         Award = "jcc"         // Deprecated: jcc
	JCG         Award = "jcg"         // Deprecated: jcg
	MARATHON    Award = "marathon"    // Deprecated: marathon
	RDA         Award = "rda"         // Deprecated: rda
	USACA       Award = "usaca"       // Deprecated: usaca
	VUCC        Award = "vucc"        // Deprecated: vucc
	WAB         Award = "wab"         // Deprecated: wab
	WAC         Award = "wac"         // Deprecated: wac
	WAE         Award = "wae"         // Deprecated: wae
	WAIP        Award = "waip"        // Deprecated: waip
	WAJA        Award = "waja"        // Deprecated: waja
	WAS         Award = "was"         // Deprecated: was
	WAZ         Award = "waz"         // Deprecated: waz
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Award specifications.
var lookupList = []Spec{
	{IsImportOnly: true, Key: "aja"},
	{IsImportOnly: true, Key: "cqdx"},
	{IsImportOnly: true, Key: "cqdxfield"},
	{IsImportOnly: true, Key: "cqwaz_160m"},
	{IsImportOnly: true, Key: "cqwaz_cw"},
	{IsImportOnly: true, Key: "cqwaz_mixed"},
	{IsImportOnly: true, Key: "cqwaz_phone"},
	{IsImportOnly: true, Key: "cqwaz_rtty"},
	{IsImportOnly: true, Key: "cqwpx"},
	{IsImportOnly: true, Key: "darc_dok"},
	{IsImportOnly: true, Key: "dxcc"},
	{IsImportOnly: true, Key: "dxcc_cw"},
	{IsImportOnly: true, Key: "dxcc_mixed"},
	{IsImportOnly: true, Key: "dxcc_phone"},
	{IsImportOnly: true, Key: "dxcc_rtty"},
	{IsImportOnly: true, Key: "iota"},
	{IsImportOnly: true, Key: "jcc"},
	{IsImportOnly: true, Key: "jcg"},
	{IsImportOnly: true, Key: "marathon"},
	{IsImportOnly: true, Key: "rda"},
	{IsImportOnly: true, Key: "usaca"},
	{IsImportOnly: true, Key: "vucc"},
	{IsImportOnly: true, Key: "wab"},
	{IsImportOnly: true, Key: "wac"},
	{IsImportOnly: true, Key: "wae"},
	{IsImportOnly: true, Key: "waip"},
	{IsImportOnly: true, Key: "waja"},
	{IsImportOnly: true, Key: "was"},
	{IsImportOnly: true, Key: "waz"},
}

// lookupMap contains all known Award specifications.
var lookupMap = map[Award]*Spec{
	AJA:         &lookupList[0],
	CQDX:        &lookupList[1],
	CQDXFIELD:   &lookupList[2],
	CQWAZ_160M:  &lookupList[3],
	CQWAZ_CW:    &lookupList[4],
	CQWAZ_MIXED: &lookupList[5],
	CQWAZ_PHONE: &lookupList[6],
	CQWAZ_RTTY:  &lookupList[7],
	CQWPX:       &lookupList[8],
	DARC_DOK:    &lookupList[9],
	DXCC:        &lookupList[10],
	DXCC_CW:     &lookupList[11],
	DXCC_MIXED:  &lookupList[12],
	DXCC_PHONE:  &lookupList[13],
	DXCC_RTTY:   &lookupList[14],
	IOTA:        &lookupList[15],
	JCC:         &lookupList[16],
	JCG:         &lookupList[17],
	MARATHON:    &lookupList[18],
	RDA:         &lookupList[19],
	USACA:       &lookupList[20],
	VUCC:        &lookupList[21],
	WAB:         &lookupList[22],
	WAC:         &lookupList[23],
	WAE:         &lookupList[24],
	WAIP:        &lookupList[25],
	WAJA:        &lookupList[26],
	WAS:         &lookupList[27],
	WAZ:         &lookupList[28],
}

// Lookup returns the specification for the provided Award.
// ADIF 3.1.6
func Lookup(a Award) (Spec, bool) {
	spec, ok := lookupMap[a]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all Award specifications that match the provided filter function.
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

// List returns all Award specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns Award specifications.
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
