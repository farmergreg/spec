// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package award provides code and constants as defined in ADIF 3.1.6 (Proposed)
package award

import "maps"

const (
	AJA         Award = "AJA"         // Deprecated: AJA
	CQDX        Award = "CQDX"        // Deprecated: CQDX
	CQDXFIELD   Award = "CQDXFIELD"   // Deprecated: CQDXFIELD
	CQWAZ_160m  Award = "CQWAZ_160m"  // Deprecated: CQWAZ_160m
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

// All Award specifications including depreciated and import only.
func AwardListAll() []Spec {
	return append([]Spec(nil), internalAwardListAll...)
}

// All Award specifications values that are NOT marked import-only.
func AwardListCurrent() []Spec {
	return append([]Spec(nil), internalAwardListCurrent...)
}

// A map of all Award specifications.
func AwardMap() map[Award]Spec {
	cp := make(map[Award]Spec, len(internalAwardMap))
	maps.Copy(cp, internalAwardMap)
	return cp
}

// A map of all Award specifications.
var internalAwardMap = map[Award]Spec{
	AJA:         {IsImportOnly: true, Key: "AJA"},
	CQDX:        {IsImportOnly: true, Key: "CQDX"},
	CQDXFIELD:   {IsImportOnly: true, Key: "CQDXFIELD"},
	CQWAZ_160m:  {IsImportOnly: true, Key: "CQWAZ_160m"},
	CQWAZ_CW:    {IsImportOnly: true, Key: "CQWAZ_CW"},
	CQWAZ_MIXED: {IsImportOnly: true, Key: "CQWAZ_MIXED"},
	CQWAZ_PHONE: {IsImportOnly: true, Key: "CQWAZ_PHONE"},
	CQWAZ_RTTY:  {IsImportOnly: true, Key: "CQWAZ_RTTY"},
	CQWPX:       {IsImportOnly: true, Key: "CQWPX"},
	DARC_DOK:    {IsImportOnly: true, Key: "DARC_DOK"},
	DXCC:        {IsImportOnly: true, Key: "DXCC"},
	DXCC_CW:     {IsImportOnly: true, Key: "DXCC_CW"},
	DXCC_MIXED:  {IsImportOnly: true, Key: "DXCC_MIXED"},
	DXCC_PHONE:  {IsImportOnly: true, Key: "DXCC_PHONE"},
	DXCC_RTTY:   {IsImportOnly: true, Key: "DXCC_RTTY"},
	IOTA:        {IsImportOnly: true, Key: "IOTA"},
	JCC:         {IsImportOnly: true, Key: "JCC"},
	JCG:         {IsImportOnly: true, Key: "JCG"},
	MARATHON:    {IsImportOnly: true, Key: "MARATHON"},
	RDA:         {IsImportOnly: true, Key: "RDA"},
	USACA:       {IsImportOnly: true, Key: "USACA"},
	VUCC:        {IsImportOnly: true, Key: "VUCC"},
	WAB:         {IsImportOnly: true, Key: "WAB"},
	WAC:         {IsImportOnly: true, Key: "WAC"},
	WAE:         {IsImportOnly: true, Key: "WAE"},
	WAIP:        {IsImportOnly: true, Key: "WAIP"},
	WAJA:        {IsImportOnly: true, Key: "WAJA"},
	WAS:         {IsImportOnly: true, Key: "WAS"},
	WAZ:         {IsImportOnly: true, Key: "WAZ"},
}

var internalAwardListAll = []Spec{
	internalAwardMap[AJA],
	internalAwardMap[CQDX],
	internalAwardMap[CQDXFIELD],
	internalAwardMap[CQWAZ_160m],
	internalAwardMap[CQWAZ_CW],
	internalAwardMap[CQWAZ_MIXED],
	internalAwardMap[CQWAZ_PHONE],
	internalAwardMap[CQWAZ_RTTY],
	internalAwardMap[CQWPX],
	internalAwardMap[DARC_DOK],
	internalAwardMap[DXCC],
	internalAwardMap[DXCC_CW],
	internalAwardMap[DXCC_MIXED],
	internalAwardMap[DXCC_PHONE],
	internalAwardMap[DXCC_RTTY],
	internalAwardMap[IOTA],
	internalAwardMap[JCC],
	internalAwardMap[JCG],
	internalAwardMap[MARATHON],
	internalAwardMap[RDA],
	internalAwardMap[USACA],
	internalAwardMap[VUCC],
	internalAwardMap[WAB],
	internalAwardMap[WAC],
	internalAwardMap[WAE],
	internalAwardMap[WAIP],
	internalAwardMap[WAJA],
	internalAwardMap[WAS],
	internalAwardMap[WAZ],
}

var internalAwardListCurrent = []Spec{}
