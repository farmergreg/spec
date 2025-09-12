// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package award provides code and constants as defined in ADIF 3.1.6 (Proposed)
package award

import "sync"

const (
	AJA         Award = "AJA"         // Deprecated: AJA
	CQDX        Award = "CQDX"        // Deprecated: CQDX
	CQDXFIELD   Award = "CQDXFIELD"   // Deprecated: CQDXFIELD
	CQWAZ_160M  Award = "CQWAZ_160M"  // Deprecated: CQWAZ_160M
	CQWAZ_CW    Award = "CQWAZ_CW"    // Deprecated: CQWAZ_CW
	CQWAZ_MIXED Award = "CQWAZ_MIXED" // Deprecated: CQWAZ_MIXED
	CQWAZ_PHONE Award = "CQWAZ_PHONE" // Deprecated: CQWAZ_PHONE
	CQWAZ_RTTY  Award = "CQWAZ_RTTY"  // Deprecated: CQWAZ_RTTY
	CQWPX       Award = "CQWPX"       // Deprecated: CQWPX
	DARC_DOK    Award = "DARC_DOK"    // Deprecated: DARC_DOK
	DXCC        Award = "DXCC"        // Deprecated: DXCC
	DXCC_CW     Award = "DXCC_CW"     // Deprecated: DXCC_CW
	DXCC_MIXED  Award = "DXCC_MIXED"  // Deprecated: DXCC_MIXED
	DXCC_PHONE  Award = "DXCC_PHONE"  // Deprecated: DXCC_PHONE
	DXCC_RTTY   Award = "DXCC_RTTY"   // Deprecated: DXCC_RTTY
	IOTA        Award = "IOTA"        // Deprecated: IOTA
	JCC         Award = "JCC"         // Deprecated: JCC
	JCG         Award = "JCG"         // Deprecated: JCG
	MARATHON    Award = "MARATHON"    // Deprecated: MARATHON
	RDA         Award = "RDA"         // Deprecated: RDA
	USACA       Award = "USACA"       // Deprecated: USACA
	VUCC        Award = "VUCC"        // Deprecated: VUCC
	WAB         Award = "WAB"         // Deprecated: WAB
	WAC         Award = "WAC"         // Deprecated: WAC
	WAE         Award = "WAE"         // Deprecated: WAE
	WAIP        Award = "WAIP"        // Deprecated: WAIP
	WAJA        Award = "WAJA"        // Deprecated: WAJA
	WAS         Award = "WAS"         // Deprecated: WAS
	WAZ         Award = "WAZ"         // Deprecated: WAZ
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Award specifications.
var lookupList = []Spec{
	{IsImportOnly: true, Key: "AJA"},
	{IsImportOnly: true, Key: "CQDX"},
	{IsImportOnly: true, Key: "CQDXFIELD"},
	{IsImportOnly: true, Key: "CQWAZ_160M"},
	{IsImportOnly: true, Key: "CQWAZ_CW"},
	{IsImportOnly: true, Key: "CQWAZ_MIXED"},
	{IsImportOnly: true, Key: "CQWAZ_PHONE"},
	{IsImportOnly: true, Key: "CQWAZ_RTTY"},
	{IsImportOnly: true, Key: "CQWPX"},
	{IsImportOnly: true, Key: "DARC_DOK"},
	{IsImportOnly: true, Key: "DXCC"},
	{IsImportOnly: true, Key: "DXCC_CW"},
	{IsImportOnly: true, Key: "DXCC_MIXED"},
	{IsImportOnly: true, Key: "DXCC_PHONE"},
	{IsImportOnly: true, Key: "DXCC_RTTY"},
	{IsImportOnly: true, Key: "IOTA"},
	{IsImportOnly: true, Key: "JCC"},
	{IsImportOnly: true, Key: "JCG"},
	{IsImportOnly: true, Key: "MARATHON"},
	{IsImportOnly: true, Key: "RDA"},
	{IsImportOnly: true, Key: "USACA"},
	{IsImportOnly: true, Key: "VUCC"},
	{IsImportOnly: true, Key: "WAB"},
	{IsImportOnly: true, Key: "WAC"},
	{IsImportOnly: true, Key: "WAE"},
	{IsImportOnly: true, Key: "WAIP"},
	{IsImportOnly: true, Key: "WAJA"},
	{IsImportOnly: true, Key: "WAS"},
	{IsImportOnly: true, Key: "WAZ"},
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
func Lookup(award Award) (Spec, bool) {
	spec, ok := lookupMap[award]
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
