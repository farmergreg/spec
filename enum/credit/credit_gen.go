// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package credit provides code and constants as defined in ADIF 3.1.6 (Proposed)
package credit

import "sync"

const (
	CQDX                Credit = "CQDX"                // CQDX                 = CQ Magazine     DX                                            Mixed
	CQDXFIELD           Credit = "CQDXFIELD"           // CQDXFIELD            = CQ Magazine     DX Field                                      Mixed
	CQDXFIELD_BAND      Credit = "CQDXFIELD_BAND"      // CQDXFIELD_BAND       = CQ Magazine     DX Field                                      Band
	CQDXFIELD_MOBILE    Credit = "CQDXFIELD_MOBILE"    // CQDXFIELD_MOBILE     = CQ Magazine     DX Field                                      Mobile
	CQDXFIELD_MODE      Credit = "CQDXFIELD_MODE"      // CQDXFIELD_MODE       = CQ Magazine     DX Field                                      Mode
	CQDXFIELD_QRP       Credit = "CQDXFIELD_QRP"       // CQDXFIELD_QRP        = CQ Magazine     DX Field                                      QRP
	CQDXFIELD_SATELLITE Credit = "CQDXFIELD_SATELLITE" // CQDXFIELD_SATELLITE  = CQ Magazine     DX Field                                      Satellite
	CQDX_BAND           Credit = "CQDX_BAND"           // CQDX_BAND            = CQ Magazine     DX                                            Band
	CQDX_MOBILE         Credit = "CQDX_MOBILE"         // CQDX_MOBILE          = CQ Magazine     DX                                            Mobile
	CQDX_MODE           Credit = "CQDX_MODE"           // CQDX_MODE            = CQ Magazine     DX                                            Mode
	CQDX_QRP            Credit = "CQDX_QRP"            // CQDX_QRP             = CQ Magazine     DX                                            QRP
	CQDX_SATELLITE      Credit = "CQDX_SATELLITE"      // CQDX_SATELLITE       = CQ Magazine     DX                                            Satellite
	CQWAZ_BAND          Credit = "CQWAZ_BAND"          // CQWAZ_BAND           = CQ Magazine     Worked All Zones (WAZ)                        Band
	CQWAZ_EME           Credit = "CQWAZ_EME"           // CQWAZ_EME            = CQ Magazine     Worked All Zones (WAZ)                        EME
	CQWAZ_MIXED         Credit = "CQWAZ_MIXED"         // CQWAZ_MIXED          = CQ Magazine     Worked All Zones (WAZ)                        Mixed
	CQWAZ_MOBILE        Credit = "CQWAZ_MOBILE"        // CQWAZ_MOBILE         = CQ Magazine     Worked All Zones (WAZ)                        Mobile
	CQWAZ_MODE          Credit = "CQWAZ_MODE"          // CQWAZ_MODE           = CQ Magazine     Worked All Zones (WAZ)                        Mode
	CQWAZ_QRP           Credit = "CQWAZ_QRP"           // CQWAZ_QRP            = CQ Magazine     Worked All Zones (WAZ)                        QRP
	CQWAZ_SATELLITE     Credit = "CQWAZ_SATELLITE"     // CQWAZ_SATELLITE      = CQ Magazine     Worked All Zones (WAZ)                        Satellite
	CQWPX               Credit = "CQWPX"               // CQWPX                = CQ Magazine     WPX                                           Mixed
	CQWPX_BAND          Credit = "CQWPX_BAND"          // CQWPX_BAND           = CQ Magazine     WPX                                           Band
	CQWPX_MODE          Credit = "CQWPX_MODE"          // CQWPX_MODE           = CQ Magazine     WPX                                           Mode
	DXCC                Credit = "DXCC"                // DXCC                 = ARRL            DX Century Club (DXCC)                        Mixed
	DXCC_BAND           Credit = "DXCC_BAND"           // DXCC_BAND            = ARRL            DX Century Club (DXCC)                        Band
	DXCC_MODE           Credit = "DXCC_MODE"           // DXCC_MODE            = ARRL            DX Century Club (DXCC)                        Mode
	DXCC_SATELLITE      Credit = "DXCC_SATELLITE"      // DXCC_SATELLITE       = ARRL            DX Century Club (DXCC)                        Satellite
	EAUSTRALIA          Credit = "EAUSTRALIA"          // EAUSTRALIA           = eQSL            eAustralia                                    Mixed
	ECANADA             Credit = "ECANADA"             // ECANADA              = eQSL            eCanada                                       Mixed
	ECOUNTY_STATE       Credit = "ECOUNTY_STATE"       // ECOUNTY_STATE        = eQSL            eCounty                                       State
	EDX                 Credit = "EDX"                 // EDX                  = eQSL            eDX                                           Mixed
	EDX100              Credit = "EDX100"              // EDX100               = eQSL            eDX100                                        Mixed
	EDX100_BAND         Credit = "EDX100_BAND"         // EDX100_BAND          = eQSL            eDX100                                        Band
	EDX100_MODE         Credit = "EDX100_MODE"         // EDX100_MODE          = eQSL            eDX100                                        Mode
	EECHOLINK50         Credit = "EECHOLINK50"         // EECHOLINK50          = eQSL            eEcholink50                                   Echolink
	EGRID_BAND          Credit = "EGRID_BAND"          // EGRID_BAND           = eQSL            eGrid                                         Band
	EGRID_SATELLITE     Credit = "EGRID_SATELLITE"     // EGRID_SATELLITE      = eQSL            eGrid                                         Satellite
	EPFX300             Credit = "EPFX300"             // EPFX300              = eQSL            ePfx300                                       Mixed
	EPFX300_MODE        Credit = "EPFX300_MODE"        // EPFX300_MODE         = eQSL            ePfx300                                       Mode
	EWAS                Credit = "EWAS"                // EWAS                 = eQSL            eWAS                                          Mixed
	EWAS_BAND           Credit = "EWAS_BAND"           // EWAS_BAND            = eQSL            eWAS                                          Band
	EWAS_MODE           Credit = "EWAS_MODE"           // EWAS_MODE            = eQSL            eWAS                                          Mode
	EWAS_SATELLITE      Credit = "EWAS_SATELLITE"      // EWAS_SATELLITE       = eQSL            eWAS                                          Satellite
	EZ40                Credit = "EZ40"                // EZ40                 = eQSL            eZ40                                          Mixed
	EZ40_MODE           Credit = "EZ40_MODE"           // EZ40_MODE            = eQSL            eZ40                                          Mode
	FFMA                Credit = "FFMA"                // FFMA                 = ARRL            Fred Fish Memorial Award (FFMA)               Mixed
	IOTA                Credit = "IOTA"                // IOTA                 = RSGB            Islands on the Air (IOTA)                     Mixed
	IOTA_BASIC          Credit = "IOTA_BASIC"          // IOTA_BASIC           = RSGB            Islands on the Air (IOTA)                     Mixed
	IOTA_CONT           Credit = "IOTA_CONT"           // IOTA_CONT            = RSGB            Islands on the Air (IOTA)                     Continent
	IOTA_GROUP          Credit = "IOTA_GROUP"          // IOTA_GROUP           = RSGB            Islands on the Air (IOTA)                     Group
	RDA                 Credit = "RDA"                 // RDA                  = TAG             Russian Districts Award (RDA)                 Mixed
	USACA               Credit = "USACA"               // USACA                = CQ Magazine     United States of America Counties (USA-CA)    Mixed
	VUCC_BAND           Credit = "VUCC_BAND"           // VUCC_BAND            = ARRL            VHF/UHF Century Club Program (VUCC)           Band
	VUCC_SATELLITE      Credit = "VUCC_SATELLITE"      // VUCC_SATELLITE       = ARRL            VHF/UHF Century Club Program (VUCC)           Satellite
	WAB                 Credit = "WAB"                 // WAB                  = WAB AG          Worked All Britain (WAB)                      Mixed
	WAC                 Credit = "WAC"                 // WAC                  = IARU            Worked All Continents (WAC)                   Mixed
	WAC_BAND            Credit = "WAC_BAND"            // WAC_BAND             = IARU            Worked All Continents (WAC)                   Band
	WAE                 Credit = "WAE"                 // WAE                  = DARC            Worked All Europe (WAE)                       Mixed
	WAE_BAND            Credit = "WAE_BAND"            // WAE_BAND             = DARC            Worked All Europe (WAE)                       Band
	WAE_MODE            Credit = "WAE_MODE"            // WAE_MODE             = DARC            Worked All Europe (WAE)                       Mode
	WAIP                Credit = "WAIP"                // WAIP                 = ARI             Worked All Italian Provinces (WAIP)           Mixed
	WAIP_BAND           Credit = "WAIP_BAND"           // WAIP_BAND            = ARI             Worked All Italian Provinces (WAIP)           Band
	WAIP_MODE           Credit = "WAIP_MODE"           // WAIP_MODE            = ARI             Worked All Italian Provinces (WAIP)           Mode
	WAS                 Credit = "WAS"                 // WAS                  = ARRL            Worked All States (WAS)                       Mixed
	WAS_BAND            Credit = "WAS_BAND"            // WAS_BAND             = ARRL            Worked All States (WAS)                       Band
	WAS_EME             Credit = "WAS_EME"             // WAS_EME              = ARRL            Worked All States (WAS)                       EME
	WAS_MODE            Credit = "WAS_MODE"            // WAS_MODE             = ARRL            Worked All States (WAS)                       Mode
	WAS_NOVICE          Credit = "WAS_NOVICE"          // WAS_NOVICE           = ARRL            Worked All States (WAS)                       Novice
	WAS_QRP             Credit = "WAS_QRP"             // WAS_QRP              = ARRL            Worked All States (WAS)                       QRP
	WAS_SATELLITE       Credit = "WAS_SATELLITE"       // WAS_SATELLITE        = ARRL            Worked All States (WAS)                       Satellite
	WITUZ               Credit = "WITUZ"               // WITUZ                = RSGB            Worked ITU Zones (WITUZ)                      Mixed
	WITUZ_BAND          Credit = "WITUZ_BAND"          // WITUZ_BAND           = RSGB            Worked ITU Zones (WITUZ)                      Band
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Credit specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "CQDX", Sponsor: "CQ Magazine", Award: "DX", Facet: "Mixed"},
	{IsImportOnly: false, Key: "CQDXFIELD", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Mixed"},
	{IsImportOnly: false, Key: "CQDXFIELD_BAND", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Band"},
	{IsImportOnly: false, Key: "CQDXFIELD_MOBILE", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Mobile"},
	{IsImportOnly: false, Key: "CQDXFIELD_MODE", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Mode"},
	{IsImportOnly: false, Key: "CQDXFIELD_QRP", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "QRP"},
	{IsImportOnly: false, Key: "CQDXFIELD_SATELLITE", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Satellite"},
	{IsImportOnly: false, Key: "CQDX_BAND", Sponsor: "CQ Magazine", Award: "DX", Facet: "Band"},
	{IsImportOnly: false, Key: "CQDX_MOBILE", Sponsor: "CQ Magazine", Award: "DX", Facet: "Mobile"},
	{IsImportOnly: false, Key: "CQDX_MODE", Sponsor: "CQ Magazine", Award: "DX", Facet: "Mode"},
	{IsImportOnly: false, Key: "CQDX_QRP", Sponsor: "CQ Magazine", Award: "DX", Facet: "QRP"},
	{IsImportOnly: false, Key: "CQDX_SATELLITE", Sponsor: "CQ Magazine", Award: "DX", Facet: "Satellite"},
	{IsImportOnly: false, Key: "CQWAZ_BAND", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Band"},
	{IsImportOnly: false, Key: "CQWAZ_EME", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "EME"},
	{IsImportOnly: false, Key: "CQWAZ_MIXED", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "CQWAZ_MOBILE", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Mobile"},
	{IsImportOnly: false, Key: "CQWAZ_MODE", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Mode"},
	{IsImportOnly: false, Key: "CQWAZ_QRP", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "QRP"},
	{IsImportOnly: false, Key: "CQWAZ_SATELLITE", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Satellite"},
	{IsImportOnly: false, Key: "CQWPX", Sponsor: "CQ Magazine", Award: "WPX", Facet: "Mixed"},
	{IsImportOnly: false, Key: "CQWPX_BAND", Sponsor: "CQ Magazine", Award: "WPX", Facet: "Band"},
	{IsImportOnly: false, Key: "CQWPX_MODE", Sponsor: "CQ Magazine", Award: "WPX", Facet: "Mode"},
	{IsImportOnly: false, Key: "DXCC", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "DXCC_BAND", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Band"},
	{IsImportOnly: false, Key: "DXCC_MODE", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Mode"},
	{IsImportOnly: false, Key: "DXCC_SATELLITE", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Satellite"},
	{IsImportOnly: false, Key: "EAUSTRALIA", Sponsor: "eQSL", Award: "eAustralia", Facet: "Mixed"},
	{IsImportOnly: false, Key: "ECANADA", Sponsor: "eQSL", Award: "eCanada", Facet: "Mixed"},
	{IsImportOnly: false, Key: "ECOUNTY_STATE", Sponsor: "eQSL", Award: "eCounty", Facet: "State"},
	{IsImportOnly: false, Key: "EDX", Sponsor: "eQSL", Award: "eDX", Facet: "Mixed"},
	{IsImportOnly: false, Key: "EDX100", Sponsor: "eQSL", Award: "eDX100", Facet: "Mixed"},
	{IsImportOnly: false, Key: "EDX100_BAND", Sponsor: "eQSL", Award: "eDX100", Facet: "Band"},
	{IsImportOnly: false, Key: "EDX100_MODE", Sponsor: "eQSL", Award: "eDX100", Facet: "Mode"},
	{IsImportOnly: false, Key: "EECHOLINK50", Sponsor: "eQSL", Award: "eEcholink50", Facet: "Echolink"},
	{IsImportOnly: false, Key: "EGRID_BAND", Sponsor: "eQSL", Award: "eGrid", Facet: "Band"},
	{IsImportOnly: false, Key: "EGRID_SATELLITE", Sponsor: "eQSL", Award: "eGrid", Facet: "Satellite"},
	{IsImportOnly: false, Key: "EPFX300", Sponsor: "eQSL", Award: "ePfx300", Facet: "Mixed"},
	{IsImportOnly: false, Key: "EPFX300_MODE", Sponsor: "eQSL", Award: "ePfx300", Facet: "Mode"},
	{IsImportOnly: false, Key: "EWAS", Sponsor: "eQSL", Award: "eWAS", Facet: "Mixed"},
	{IsImportOnly: false, Key: "EWAS_BAND", Sponsor: "eQSL", Award: "eWAS", Facet: "Band"},
	{IsImportOnly: false, Key: "EWAS_MODE", Sponsor: "eQSL", Award: "eWAS", Facet: "Mode"},
	{IsImportOnly: false, Key: "EWAS_SATELLITE", Sponsor: "eQSL", Award: "eWAS", Facet: "Satellite"},
	{IsImportOnly: false, Key: "EZ40", Sponsor: "eQSL", Award: "eZ40", Facet: "Mixed"},
	{IsImportOnly: false, Key: "EZ40_MODE", Sponsor: "eQSL", Award: "eZ40", Facet: "Mode"},
	{IsImportOnly: false, Key: "FFMA", Sponsor: "ARRL", Award: "Fred Fish Memorial Award (FFMA)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "IOTA", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "IOTA_BASIC", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "IOTA_CONT", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Continent"},
	{IsImportOnly: false, Key: "IOTA_GROUP", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Group"},
	{IsImportOnly: false, Key: "RDA", Sponsor: "TAG", Award: "Russian Districts Award (RDA)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "USACA", Sponsor: "CQ Magazine", Award: "United States of America Counties (USA-CA)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "VUCC_BAND", Sponsor: "ARRL", Award: "VHF/UHF Century Club Program (VUCC)", Facet: "Band"},
	{IsImportOnly: false, Key: "VUCC_SATELLITE", Sponsor: "ARRL", Award: "VHF/UHF Century Club Program (VUCC)", Facet: "Satellite"},
	{IsImportOnly: false, Key: "WAB", Sponsor: "WAB AG", Award: "Worked All Britain (WAB)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "WAC", Sponsor: "IARU", Award: "Worked All Continents (WAC)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "WAC_BAND", Sponsor: "IARU", Award: "Worked All Continents (WAC)", Facet: "Band"},
	{IsImportOnly: false, Key: "WAE", Sponsor: "DARC", Award: "Worked All Europe (WAE)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "WAE_BAND", Sponsor: "DARC", Award: "Worked All Europe (WAE)", Facet: "Band"},
	{IsImportOnly: false, Key: "WAE_MODE", Sponsor: "DARC", Award: "Worked All Europe (WAE)", Facet: "Mode"},
	{IsImportOnly: false, Key: "WAIP", Sponsor: "ARI", Award: "Worked All Italian Provinces (WAIP)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "WAIP_BAND", Sponsor: "ARI", Award: "Worked All Italian Provinces (WAIP)", Facet: "Band"},
	{IsImportOnly: false, Key: "WAIP_MODE", Sponsor: "ARI", Award: "Worked All Italian Provinces (WAIP)", Facet: "Mode"},
	{IsImportOnly: false, Key: "WAS", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "WAS_BAND", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Band"},
	{IsImportOnly: false, Key: "WAS_EME", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "EME"},
	{IsImportOnly: false, Key: "WAS_MODE", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Mode"},
	{IsImportOnly: false, Key: "WAS_NOVICE", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Novice"},
	{IsImportOnly: false, Key: "WAS_QRP", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "QRP"},
	{IsImportOnly: false, Key: "WAS_SATELLITE", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Satellite"},
	{IsImportOnly: false, Key: "WITUZ", Sponsor: "RSGB", Award: "Worked ITU Zones (WITUZ)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "WITUZ_BAND", Sponsor: "RSGB", Award: "Worked ITU Zones (WITUZ)", Facet: "Band"},
}

// lookupMap contains all known Credit specifications.
var lookupMap = map[Credit]*Spec{
	CQDX:                &lookupList[0],
	CQDXFIELD:           &lookupList[1],
	CQDXFIELD_BAND:      &lookupList[2],
	CQDXFIELD_MOBILE:    &lookupList[3],
	CQDXFIELD_MODE:      &lookupList[4],
	CQDXFIELD_QRP:       &lookupList[5],
	CQDXFIELD_SATELLITE: &lookupList[6],
	CQDX_BAND:           &lookupList[7],
	CQDX_MOBILE:         &lookupList[8],
	CQDX_MODE:           &lookupList[9],
	CQDX_QRP:            &lookupList[10],
	CQDX_SATELLITE:      &lookupList[11],
	CQWAZ_BAND:          &lookupList[12],
	CQWAZ_EME:           &lookupList[13],
	CQWAZ_MIXED:         &lookupList[14],
	CQWAZ_MOBILE:        &lookupList[15],
	CQWAZ_MODE:          &lookupList[16],
	CQWAZ_QRP:           &lookupList[17],
	CQWAZ_SATELLITE:     &lookupList[18],
	CQWPX:               &lookupList[19],
	CQWPX_BAND:          &lookupList[20],
	CQWPX_MODE:          &lookupList[21],
	DXCC:                &lookupList[22],
	DXCC_BAND:           &lookupList[23],
	DXCC_MODE:           &lookupList[24],
	DXCC_SATELLITE:      &lookupList[25],
	EAUSTRALIA:          &lookupList[26],
	ECANADA:             &lookupList[27],
	ECOUNTY_STATE:       &lookupList[28],
	EDX:                 &lookupList[29],
	EDX100:              &lookupList[30],
	EDX100_BAND:         &lookupList[31],
	EDX100_MODE:         &lookupList[32],
	EECHOLINK50:         &lookupList[33],
	EGRID_BAND:          &lookupList[34],
	EGRID_SATELLITE:     &lookupList[35],
	EPFX300:             &lookupList[36],
	EPFX300_MODE:        &lookupList[37],
	EWAS:                &lookupList[38],
	EWAS_BAND:           &lookupList[39],
	EWAS_MODE:           &lookupList[40],
	EWAS_SATELLITE:      &lookupList[41],
	EZ40:                &lookupList[42],
	EZ40_MODE:           &lookupList[43],
	FFMA:                &lookupList[44],
	IOTA:                &lookupList[45],
	IOTA_BASIC:          &lookupList[46],
	IOTA_CONT:           &lookupList[47],
	IOTA_GROUP:          &lookupList[48],
	RDA:                 &lookupList[49],
	USACA:               &lookupList[50],
	VUCC_BAND:           &lookupList[51],
	VUCC_SATELLITE:      &lookupList[52],
	WAB:                 &lookupList[53],
	WAC:                 &lookupList[54],
	WAC_BAND:            &lookupList[55],
	WAE:                 &lookupList[56],
	WAE_BAND:            &lookupList[57],
	WAE_MODE:            &lookupList[58],
	WAIP:                &lookupList[59],
	WAIP_BAND:           &lookupList[60],
	WAIP_MODE:           &lookupList[61],
	WAS:                 &lookupList[62],
	WAS_BAND:            &lookupList[63],
	WAS_EME:             &lookupList[64],
	WAS_MODE:            &lookupList[65],
	WAS_NOVICE:          &lookupList[66],
	WAS_QRP:             &lookupList[67],
	WAS_SATELLITE:       &lookupList[68],
	WITUZ:               &lookupList[69],
	WITUZ_BAND:          &lookupList[70],
}

// Lookup returns the specification for the provided Credit.
// ADIF 3.1.6
func Lookup(c Credit) (Spec, bool) {
	spec, ok := lookupMap[c]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all Credit specifications that match the provided filter function.
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

// List returns all Credit specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns Credit specifications.
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
