// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package credit provides code and constants as defined in ADIF 3.1.6 (Released)
package credit

import "sync"

const (
	CQDX                Credit = "cqdx"                // cqdx                 = CQ Magazine     DX                                            Mixed
	CQDXFIELD           Credit = "cqdxfield"           // cqdxfield            = CQ Magazine     DX Field                                      Mixed
	CQDXFIELD_BAND      Credit = "cqdxfield_band"      // cqdxfield_band       = CQ Magazine     DX Field                                      Band
	CQDXFIELD_MOBILE    Credit = "cqdxfield_mobile"    // cqdxfield_mobile     = CQ Magazine     DX Field                                      Mobile
	CQDXFIELD_MODE      Credit = "cqdxfield_mode"      // cqdxfield_mode       = CQ Magazine     DX Field                                      Mode
	CQDXFIELD_QRP       Credit = "cqdxfield_qrp"       // cqdxfield_qrp        = CQ Magazine     DX Field                                      QRP
	CQDXFIELD_SATELLITE Credit = "cqdxfield_satellite" // cqdxfield_satellite  = CQ Magazine     DX Field                                      Satellite
	CQDX_BAND           Credit = "cqdx_band"           // cqdx_band            = CQ Magazine     DX                                            Band
	CQDX_MOBILE         Credit = "cqdx_mobile"         // cqdx_mobile          = CQ Magazine     DX                                            Mobile
	CQDX_MODE           Credit = "cqdx_mode"           // cqdx_mode            = CQ Magazine     DX                                            Mode
	CQDX_QRP            Credit = "cqdx_qrp"            // cqdx_qrp             = CQ Magazine     DX                                            QRP
	CQDX_SATELLITE      Credit = "cqdx_satellite"      // cqdx_satellite       = CQ Magazine     DX                                            Satellite
	CQWAZ_BAND          Credit = "cqwaz_band"          // cqwaz_band           = CQ Magazine     Worked All Zones (WAZ)                        Band
	CQWAZ_EME           Credit = "cqwaz_eme"           // cqwaz_eme            = CQ Magazine     Worked All Zones (WAZ)                        EME
	CQWAZ_MIXED         Credit = "cqwaz_mixed"         // cqwaz_mixed          = CQ Magazine     Worked All Zones (WAZ)                        Mixed
	CQWAZ_MOBILE        Credit = "cqwaz_mobile"        // cqwaz_mobile         = CQ Magazine     Worked All Zones (WAZ)                        Mobile
	CQWAZ_MODE          Credit = "cqwaz_mode"          // cqwaz_mode           = CQ Magazine     Worked All Zones (WAZ)                        Mode
	CQWAZ_QRP           Credit = "cqwaz_qrp"           // cqwaz_qrp            = CQ Magazine     Worked All Zones (WAZ)                        QRP
	CQWAZ_SATELLITE     Credit = "cqwaz_satellite"     // cqwaz_satellite      = CQ Magazine     Worked All Zones (WAZ)                        Satellite
	CQWPX               Credit = "cqwpx"               // cqwpx                = CQ Magazine     WPX                                           Mixed
	CQWPX_BAND          Credit = "cqwpx_band"          // cqwpx_band           = CQ Magazine     WPX                                           Band
	CQWPX_MODE          Credit = "cqwpx_mode"          // cqwpx_mode           = CQ Magazine     WPX                                           Mode
	DXCC                Credit = "dxcc"                // dxcc                 = ARRL            DX Century Club (DXCC)                        Mixed
	DXCC_BAND           Credit = "dxcc_band"           // dxcc_band            = ARRL            DX Century Club (DXCC)                        Band
	DXCC_MODE           Credit = "dxcc_mode"           // dxcc_mode            = ARRL            DX Century Club (DXCC)                        Mode
	DXCC_SATELLITE      Credit = "dxcc_satellite"      // dxcc_satellite       = ARRL            DX Century Club (DXCC)                        Satellite
	EAUSTRALIA          Credit = "eaustralia"          // eaustralia           = eQSL            eAustralia                                    Mixed
	ECANADA             Credit = "ecanada"             // ecanada              = eQSL            eCanada                                       Mixed
	ECOUNTY_STATE       Credit = "ecounty_state"       // ecounty_state        = eQSL            eCounty                                       State
	EDX                 Credit = "edx"                 // edx                  = eQSL            eDX                                           Mixed
	EDX100              Credit = "edx100"              // edx100               = eQSL            eDX100                                        Mixed
	EDX100_BAND         Credit = "edx100_band"         // edx100_band          = eQSL            eDX100                                        Band
	EDX100_MODE         Credit = "edx100_mode"         // edx100_mode          = eQSL            eDX100                                        Mode
	EECHOLINK50         Credit = "eecholink50"         // eecholink50          = eQSL            eEcholink50                                   Echolink
	EGRID_BAND          Credit = "egrid_band"          // egrid_band           = eQSL            eGrid                                         Band
	EGRID_SATELLITE     Credit = "egrid_satellite"     // egrid_satellite      = eQSL            eGrid                                         Satellite
	EPFX300             Credit = "epfx300"             // epfx300              = eQSL            ePfx300                                       Mixed
	EPFX300_MODE        Credit = "epfx300_mode"        // epfx300_mode         = eQSL            ePfx300                                       Mode
	EWAS                Credit = "ewas"                // ewas                 = eQSL            eWAS                                          Mixed
	EWAS_BAND           Credit = "ewas_band"           // ewas_band            = eQSL            eWAS                                          Band
	EWAS_MODE           Credit = "ewas_mode"           // ewas_mode            = eQSL            eWAS                                          Mode
	EWAS_SATELLITE      Credit = "ewas_satellite"      // ewas_satellite       = eQSL            eWAS                                          Satellite
	EZ40                Credit = "ez40"                // ez40                 = eQSL            eZ40                                          Mixed
	EZ40_MODE           Credit = "ez40_mode"           // ez40_mode            = eQSL            eZ40                                          Mode
	FFMA                Credit = "ffma"                // ffma                 = ARRL            Fred Fish Memorial Award (FFMA)               Mixed
	IOTA                Credit = "iota"                // iota                 = RSGB            Islands on the Air (IOTA)                     Mixed
	IOTA_BASIC          Credit = "iota_basic"          // iota_basic           = RSGB            Islands on the Air (IOTA)                     Mixed
	IOTA_CONT           Credit = "iota_cont"           // iota_cont            = RSGB            Islands on the Air (IOTA)                     Continent
	IOTA_GROUP          Credit = "iota_group"          // iota_group           = RSGB            Islands on the Air (IOTA)                     Group
	RDA                 Credit = "rda"                 // rda                  = TAG             Russian Districts Award (RDA)                 Mixed
	USACA               Credit = "usaca"               // usaca                = CQ Magazine     United States of America Counties (USA-CA)    Mixed
	VUCC_BAND           Credit = "vucc_band"           // vucc_band            = ARRL            VHF/UHF Century Club Program (VUCC)           Band
	VUCC_SATELLITE      Credit = "vucc_satellite"      // vucc_satellite       = ARRL            VHF/UHF Century Club Program (VUCC)           Satellite
	WAB                 Credit = "wab"                 // wab                  = WAB AG          Worked All Britain (WAB)                      Mixed
	WAC                 Credit = "wac"                 // wac                  = IARU            Worked All Continents (WAC)                   Mixed
	WAC_BAND            Credit = "wac_band"            // wac_band             = IARU            Worked All Continents (WAC)                   Band
	WAE                 Credit = "wae"                 // wae                  = DARC            Worked All Europe (WAE)                       Mixed
	WAE_BAND            Credit = "wae_band"            // wae_band             = DARC            Worked All Europe (WAE)                       Band
	WAE_MODE            Credit = "wae_mode"            // wae_mode             = DARC            Worked All Europe (WAE)                       Mode
	WAIP                Credit = "waip"                // waip                 = ARI             Worked All Italian Provinces (WAIP)           Mixed
	WAIP_BAND           Credit = "waip_band"           // waip_band            = ARI             Worked All Italian Provinces (WAIP)           Band
	WAIP_MODE           Credit = "waip_mode"           // waip_mode            = ARI             Worked All Italian Provinces (WAIP)           Mode
	WAS                 Credit = "was"                 // was                  = ARRL            Worked All States (WAS)                       Mixed
	WAS_BAND            Credit = "was_band"            // was_band             = ARRL            Worked All States (WAS)                       Band
	WAS_EME             Credit = "was_eme"             // was_eme              = ARRL            Worked All States (WAS)                       EME
	WAS_MODE            Credit = "was_mode"            // was_mode             = ARRL            Worked All States (WAS)                       Mode
	WAS_NOVICE          Credit = "was_novice"          // was_novice           = ARRL            Worked All States (WAS)                       Novice
	WAS_QRP             Credit = "was_qrp"             // was_qrp              = ARRL            Worked All States (WAS)                       QRP
	WAS_SATELLITE       Credit = "was_satellite"       // was_satellite        = ARRL            Worked All States (WAS)                       Satellite
	WITUZ               Credit = "wituz"               // wituz                = RSGB            Worked ITU Zones (WITUZ)                      Mixed
	WITUZ_BAND          Credit = "wituz_band"          // wituz_band           = RSGB            Worked ITU Zones (WITUZ)                      Band
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Credit specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "cqdx", Sponsor: "CQ Magazine", Award: "DX", Facet: "Mixed"},
	{IsImportOnly: false, Key: "cqdxfield", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Mixed"},
	{IsImportOnly: false, Key: "cqdxfield_band", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Band"},
	{IsImportOnly: false, Key: "cqdxfield_mobile", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Mobile"},
	{IsImportOnly: false, Key: "cqdxfield_mode", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Mode"},
	{IsImportOnly: false, Key: "cqdxfield_qrp", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "QRP"},
	{IsImportOnly: false, Key: "cqdxfield_satellite", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Satellite"},
	{IsImportOnly: false, Key: "cqdx_band", Sponsor: "CQ Magazine", Award: "DX", Facet: "Band"},
	{IsImportOnly: false, Key: "cqdx_mobile", Sponsor: "CQ Magazine", Award: "DX", Facet: "Mobile"},
	{IsImportOnly: false, Key: "cqdx_mode", Sponsor: "CQ Magazine", Award: "DX", Facet: "Mode"},
	{IsImportOnly: false, Key: "cqdx_qrp", Sponsor: "CQ Magazine", Award: "DX", Facet: "QRP"},
	{IsImportOnly: false, Key: "cqdx_satellite", Sponsor: "CQ Magazine", Award: "DX", Facet: "Satellite"},
	{IsImportOnly: false, Key: "cqwaz_band", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Band"},
	{IsImportOnly: false, Key: "cqwaz_eme", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "EME"},
	{IsImportOnly: false, Key: "cqwaz_mixed", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "cqwaz_mobile", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Mobile"},
	{IsImportOnly: false, Key: "cqwaz_mode", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Mode"},
	{IsImportOnly: false, Key: "cqwaz_qrp", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "QRP"},
	{IsImportOnly: false, Key: "cqwaz_satellite", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Satellite"},
	{IsImportOnly: false, Key: "cqwpx", Sponsor: "CQ Magazine", Award: "WPX", Facet: "Mixed"},
	{IsImportOnly: false, Key: "cqwpx_band", Sponsor: "CQ Magazine", Award: "WPX", Facet: "Band"},
	{IsImportOnly: false, Key: "cqwpx_mode", Sponsor: "CQ Magazine", Award: "WPX", Facet: "Mode"},
	{IsImportOnly: false, Key: "dxcc", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "dxcc_band", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Band"},
	{IsImportOnly: false, Key: "dxcc_mode", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Mode"},
	{IsImportOnly: false, Key: "dxcc_satellite", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Satellite"},
	{IsImportOnly: false, Key: "eaustralia", Sponsor: "eQSL", Award: "eAustralia", Facet: "Mixed"},
	{IsImportOnly: false, Key: "ecanada", Sponsor: "eQSL", Award: "eCanada", Facet: "Mixed"},
	{IsImportOnly: false, Key: "ecounty_state", Sponsor: "eQSL", Award: "eCounty", Facet: "State"},
	{IsImportOnly: false, Key: "edx", Sponsor: "eQSL", Award: "eDX", Facet: "Mixed"},
	{IsImportOnly: false, Key: "edx100", Sponsor: "eQSL", Award: "eDX100", Facet: "Mixed"},
	{IsImportOnly: false, Key: "edx100_band", Sponsor: "eQSL", Award: "eDX100", Facet: "Band"},
	{IsImportOnly: false, Key: "edx100_mode", Sponsor: "eQSL", Award: "eDX100", Facet: "Mode"},
	{IsImportOnly: false, Key: "eecholink50", Sponsor: "eQSL", Award: "eEcholink50", Facet: "Echolink"},
	{IsImportOnly: false, Key: "egrid_band", Sponsor: "eQSL", Award: "eGrid", Facet: "Band"},
	{IsImportOnly: false, Key: "egrid_satellite", Sponsor: "eQSL", Award: "eGrid", Facet: "Satellite"},
	{IsImportOnly: false, Key: "epfx300", Sponsor: "eQSL", Award: "ePfx300", Facet: "Mixed"},
	{IsImportOnly: false, Key: "epfx300_mode", Sponsor: "eQSL", Award: "ePfx300", Facet: "Mode"},
	{IsImportOnly: false, Key: "ewas", Sponsor: "eQSL", Award: "eWAS", Facet: "Mixed"},
	{IsImportOnly: false, Key: "ewas_band", Sponsor: "eQSL", Award: "eWAS", Facet: "Band"},
	{IsImportOnly: false, Key: "ewas_mode", Sponsor: "eQSL", Award: "eWAS", Facet: "Mode"},
	{IsImportOnly: false, Key: "ewas_satellite", Sponsor: "eQSL", Award: "eWAS", Facet: "Satellite"},
	{IsImportOnly: false, Key: "ez40", Sponsor: "eQSL", Award: "eZ40", Facet: "Mixed"},
	{IsImportOnly: false, Key: "ez40_mode", Sponsor: "eQSL", Award: "eZ40", Facet: "Mode"},
	{IsImportOnly: false, Key: "ffma", Sponsor: "ARRL", Award: "Fred Fish Memorial Award (FFMA)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "iota", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "iota_basic", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "iota_cont", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Continent"},
	{IsImportOnly: false, Key: "iota_group", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Group"},
	{IsImportOnly: false, Key: "rda", Sponsor: "TAG", Award: "Russian Districts Award (RDA)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "usaca", Sponsor: "CQ Magazine", Award: "United States of America Counties (USA-CA)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "vucc_band", Sponsor: "ARRL", Award: "VHF/UHF Century Club Program (VUCC)", Facet: "Band"},
	{IsImportOnly: false, Key: "vucc_satellite", Sponsor: "ARRL", Award: "VHF/UHF Century Club Program (VUCC)", Facet: "Satellite"},
	{IsImportOnly: false, Key: "wab", Sponsor: "WAB AG", Award: "Worked All Britain (WAB)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "wac", Sponsor: "IARU", Award: "Worked All Continents (WAC)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "wac_band", Sponsor: "IARU", Award: "Worked All Continents (WAC)", Facet: "Band"},
	{IsImportOnly: false, Key: "wae", Sponsor: "DARC", Award: "Worked All Europe (WAE)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "wae_band", Sponsor: "DARC", Award: "Worked All Europe (WAE)", Facet: "Band"},
	{IsImportOnly: false, Key: "wae_mode", Sponsor: "DARC", Award: "Worked All Europe (WAE)", Facet: "Mode"},
	{IsImportOnly: false, Key: "waip", Sponsor: "ARI", Award: "Worked All Italian Provinces (WAIP)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "waip_band", Sponsor: "ARI", Award: "Worked All Italian Provinces (WAIP)", Facet: "Band"},
	{IsImportOnly: false, Key: "waip_mode", Sponsor: "ARI", Award: "Worked All Italian Provinces (WAIP)", Facet: "Mode"},
	{IsImportOnly: false, Key: "was", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "was_band", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Band"},
	{IsImportOnly: false, Key: "was_eme", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "EME"},
	{IsImportOnly: false, Key: "was_mode", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Mode"},
	{IsImportOnly: false, Key: "was_novice", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Novice"},
	{IsImportOnly: false, Key: "was_qrp", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "QRP"},
	{IsImportOnly: false, Key: "was_satellite", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Satellite"},
	{IsImportOnly: false, Key: "wituz", Sponsor: "RSGB", Award: "Worked ITU Zones (WITUZ)", Facet: "Mixed"},
	{IsImportOnly: false, Key: "wituz_band", Sponsor: "RSGB", Award: "Worked ITU Zones (WITUZ)", Facet: "Band"},
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
