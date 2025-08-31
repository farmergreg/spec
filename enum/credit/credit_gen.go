// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package credit provides code and constants as defined in ADIF 3.1.6 (Proposed)
package credit

import "maps"

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

// All Credit specifications in ADIF 3.1.6 (Proposed) including depreciated and import only.
func CreditListAll() []Spec {
	return append([]Spec(nil), internalCreditListAll...)
}

// All Credit specifications values in ADIF 3.1.6 (Proposed) that are NOT marked import-only.
func CreditListCurrent() []Spec {
	return append([]Spec(nil), internalCreditListCurrent...)
}

// A map of all Credit from ADIF 3.1.6 (Proposed).
func CreditMap() map[Credit]Spec {
	cp := make(map[Credit]Spec, len(internalCreditMap))
	maps.Copy(cp, internalCreditMap)
	return cp
}

// A map of all Credit specifications.
var internalCreditMap = map[Credit]Spec{
	CQDX:                {IsImportOnly: false, Key: "CQDX", Sponsor: "CQ Magazine", Award: "DX", Facet: "Mixed"},
	CQDXFIELD:           {IsImportOnly: false, Key: "CQDXFIELD", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Mixed"},
	CQDXFIELD_BAND:      {IsImportOnly: false, Key: "CQDXFIELD_BAND", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Band"},
	CQDXFIELD_MOBILE:    {IsImportOnly: false, Key: "CQDXFIELD_MOBILE", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Mobile"},
	CQDXFIELD_MODE:      {IsImportOnly: false, Key: "CQDXFIELD_MODE", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Mode"},
	CQDXFIELD_QRP:       {IsImportOnly: false, Key: "CQDXFIELD_QRP", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "QRP"},
	CQDXFIELD_SATELLITE: {IsImportOnly: false, Key: "CQDXFIELD_SATELLITE", Sponsor: "CQ Magazine", Award: "DX Field", Facet: "Satellite"},
	CQDX_BAND:           {IsImportOnly: false, Key: "CQDX_BAND", Sponsor: "CQ Magazine", Award: "DX", Facet: "Band"},
	CQDX_MOBILE:         {IsImportOnly: false, Key: "CQDX_MOBILE", Sponsor: "CQ Magazine", Award: "DX", Facet: "Mobile"},
	CQDX_MODE:           {IsImportOnly: false, Key: "CQDX_MODE", Sponsor: "CQ Magazine", Award: "DX", Facet: "Mode"},
	CQDX_QRP:            {IsImportOnly: false, Key: "CQDX_QRP", Sponsor: "CQ Magazine", Award: "DX", Facet: "QRP"},
	CQDX_SATELLITE:      {IsImportOnly: false, Key: "CQDX_SATELLITE", Sponsor: "CQ Magazine", Award: "DX", Facet: "Satellite"},
	CQWAZ_BAND:          {IsImportOnly: false, Key: "CQWAZ_BAND", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Band"},
	CQWAZ_EME:           {IsImportOnly: false, Key: "CQWAZ_EME", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "EME"},
	CQWAZ_MIXED:         {IsImportOnly: false, Key: "CQWAZ_MIXED", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Mixed"},
	CQWAZ_MOBILE:        {IsImportOnly: false, Key: "CQWAZ_MOBILE", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Mobile"},
	CQWAZ_MODE:          {IsImportOnly: false, Key: "CQWAZ_MODE", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Mode"},
	CQWAZ_QRP:           {IsImportOnly: false, Key: "CQWAZ_QRP", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "QRP"},
	CQWAZ_SATELLITE:     {IsImportOnly: false, Key: "CQWAZ_SATELLITE", Sponsor: "CQ Magazine", Award: "Worked All Zones (WAZ)", Facet: "Satellite"},
	CQWPX:               {IsImportOnly: false, Key: "CQWPX", Sponsor: "CQ Magazine", Award: "WPX", Facet: "Mixed"},
	CQWPX_BAND:          {IsImportOnly: false, Key: "CQWPX_BAND", Sponsor: "CQ Magazine", Award: "WPX", Facet: "Band"},
	CQWPX_MODE:          {IsImportOnly: false, Key: "CQWPX_MODE", Sponsor: "CQ Magazine", Award: "WPX", Facet: "Mode"},
	DXCC:                {IsImportOnly: false, Key: "DXCC", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Mixed"},
	DXCC_BAND:           {IsImportOnly: false, Key: "DXCC_BAND", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Band"},
	DXCC_MODE:           {IsImportOnly: false, Key: "DXCC_MODE", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Mode"},
	DXCC_SATELLITE:      {IsImportOnly: false, Key: "DXCC_SATELLITE", Sponsor: "ARRL", Award: "DX Century Club (DXCC)", Facet: "Satellite"},
	EAUSTRALIA:          {IsImportOnly: false, Key: "EAUSTRALIA", Sponsor: "eQSL", Award: "eAustralia", Facet: "Mixed"},
	ECANADA:             {IsImportOnly: false, Key: "ECANADA", Sponsor: "eQSL", Award: "eCanada", Facet: "Mixed"},
	ECOUNTY_STATE:       {IsImportOnly: false, Key: "ECOUNTY_STATE", Sponsor: "eQSL", Award: "eCounty", Facet: "State"},
	EDX:                 {IsImportOnly: false, Key: "EDX", Sponsor: "eQSL", Award: "eDX", Facet: "Mixed"},
	EDX100:              {IsImportOnly: false, Key: "EDX100", Sponsor: "eQSL", Award: "eDX100", Facet: "Mixed"},
	EDX100_BAND:         {IsImportOnly: false, Key: "EDX100_BAND", Sponsor: "eQSL", Award: "eDX100", Facet: "Band"},
	EDX100_MODE:         {IsImportOnly: false, Key: "EDX100_MODE", Sponsor: "eQSL", Award: "eDX100", Facet: "Mode"},
	EECHOLINK50:         {IsImportOnly: false, Key: "EECHOLINK50", Sponsor: "eQSL", Award: "eEcholink50", Facet: "Echolink"},
	EGRID_BAND:          {IsImportOnly: false, Key: "EGRID_BAND", Sponsor: "eQSL", Award: "eGrid", Facet: "Band"},
	EGRID_SATELLITE:     {IsImportOnly: false, Key: "EGRID_SATELLITE", Sponsor: "eQSL", Award: "eGrid", Facet: "Satellite"},
	EPFX300:             {IsImportOnly: false, Key: "EPFX300", Sponsor: "eQSL", Award: "ePfx300", Facet: "Mixed"},
	EPFX300_MODE:        {IsImportOnly: false, Key: "EPFX300_MODE", Sponsor: "eQSL", Award: "ePfx300", Facet: "Mode"},
	EWAS:                {IsImportOnly: false, Key: "EWAS", Sponsor: "eQSL", Award: "eWAS", Facet: "Mixed"},
	EWAS_BAND:           {IsImportOnly: false, Key: "EWAS_BAND", Sponsor: "eQSL", Award: "eWAS", Facet: "Band"},
	EWAS_MODE:           {IsImportOnly: false, Key: "EWAS_MODE", Sponsor: "eQSL", Award: "eWAS", Facet: "Mode"},
	EWAS_SATELLITE:      {IsImportOnly: false, Key: "EWAS_SATELLITE", Sponsor: "eQSL", Award: "eWAS", Facet: "Satellite"},
	EZ40:                {IsImportOnly: false, Key: "EZ40", Sponsor: "eQSL", Award: "eZ40", Facet: "Mixed"},
	EZ40_MODE:           {IsImportOnly: false, Key: "EZ40_MODE", Sponsor: "eQSL", Award: "eZ40", Facet: "Mode"},
	FFMA:                {IsImportOnly: false, Key: "FFMA", Sponsor: "ARRL", Award: "Fred Fish Memorial Award (FFMA)", Facet: "Mixed"},
	IOTA:                {IsImportOnly: false, Key: "IOTA", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Mixed"},
	IOTA_BASIC:          {IsImportOnly: false, Key: "IOTA_BASIC", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Mixed"},
	IOTA_CONT:           {IsImportOnly: false, Key: "IOTA_CONT", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Continent"},
	IOTA_GROUP:          {IsImportOnly: false, Key: "IOTA_GROUP", Sponsor: "RSGB", Award: "Islands on the Air (IOTA)", Facet: "Group"},
	RDA:                 {IsImportOnly: false, Key: "RDA", Sponsor: "TAG", Award: "Russian Districts Award (RDA)", Facet: "Mixed"},
	USACA:               {IsImportOnly: false, Key: "USACA", Sponsor: "CQ Magazine", Award: "United States of America Counties (USA-CA)", Facet: "Mixed"},
	VUCC_BAND:           {IsImportOnly: false, Key: "VUCC_BAND", Sponsor: "ARRL", Award: "VHF/UHF Century Club Program (VUCC)", Facet: "Band"},
	VUCC_SATELLITE:      {IsImportOnly: false, Key: "VUCC_SATELLITE", Sponsor: "ARRL", Award: "VHF/UHF Century Club Program (VUCC)", Facet: "Satellite"},
	WAB:                 {IsImportOnly: false, Key: "WAB", Sponsor: "WAB AG", Award: "Worked All Britain (WAB)", Facet: "Mixed"},
	WAC:                 {IsImportOnly: false, Key: "WAC", Sponsor: "IARU", Award: "Worked All Continents (WAC)", Facet: "Mixed"},
	WAC_BAND:            {IsImportOnly: false, Key: "WAC_BAND", Sponsor: "IARU", Award: "Worked All Continents (WAC)", Facet: "Band"},
	WAE:                 {IsImportOnly: false, Key: "WAE", Sponsor: "DARC", Award: "Worked All Europe (WAE)", Facet: "Mixed"},
	WAE_BAND:            {IsImportOnly: false, Key: "WAE_BAND", Sponsor: "DARC", Award: "Worked All Europe (WAE)", Facet: "Band"},
	WAE_MODE:            {IsImportOnly: false, Key: "WAE_MODE", Sponsor: "DARC", Award: "Worked All Europe (WAE)", Facet: "Mode"},
	WAIP:                {IsImportOnly: false, Key: "WAIP", Sponsor: "ARI", Award: "Worked All Italian Provinces (WAIP)", Facet: "Mixed"},
	WAIP_BAND:           {IsImportOnly: false, Key: "WAIP_BAND", Sponsor: "ARI", Award: "Worked All Italian Provinces (WAIP)", Facet: "Band"},
	WAIP_MODE:           {IsImportOnly: false, Key: "WAIP_MODE", Sponsor: "ARI", Award: "Worked All Italian Provinces (WAIP)", Facet: "Mode"},
	WAS:                 {IsImportOnly: false, Key: "WAS", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Mixed"},
	WAS_BAND:            {IsImportOnly: false, Key: "WAS_BAND", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Band"},
	WAS_EME:             {IsImportOnly: false, Key: "WAS_EME", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "EME"},
	WAS_MODE:            {IsImportOnly: false, Key: "WAS_MODE", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Mode"},
	WAS_NOVICE:          {IsImportOnly: false, Key: "WAS_NOVICE", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Novice"},
	WAS_QRP:             {IsImportOnly: false, Key: "WAS_QRP", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "QRP"},
	WAS_SATELLITE:       {IsImportOnly: false, Key: "WAS_SATELLITE", Sponsor: "ARRL", Award: "Worked All States (WAS)", Facet: "Satellite"},
	WITUZ:               {IsImportOnly: false, Key: "WITUZ", Sponsor: "RSGB", Award: "Worked ITU Zones (WITUZ)", Facet: "Mixed"},
	WITUZ_BAND:          {IsImportOnly: false, Key: "WITUZ_BAND", Sponsor: "RSGB", Award: "Worked ITU Zones (WITUZ)", Facet: "Band"},
}

var internalCreditListAll = []Spec{
	internalCreditMap[CQDX],
	internalCreditMap[CQDXFIELD],
	internalCreditMap[CQDXFIELD_BAND],
	internalCreditMap[CQDXFIELD_MOBILE],
	internalCreditMap[CQDXFIELD_MODE],
	internalCreditMap[CQDXFIELD_QRP],
	internalCreditMap[CQDXFIELD_SATELLITE],
	internalCreditMap[CQDX_BAND],
	internalCreditMap[CQDX_MOBILE],
	internalCreditMap[CQDX_MODE],
	internalCreditMap[CQDX_QRP],
	internalCreditMap[CQDX_SATELLITE],
	internalCreditMap[CQWAZ_BAND],
	internalCreditMap[CQWAZ_EME],
	internalCreditMap[CQWAZ_MIXED],
	internalCreditMap[CQWAZ_MOBILE],
	internalCreditMap[CQWAZ_MODE],
	internalCreditMap[CQWAZ_QRP],
	internalCreditMap[CQWAZ_SATELLITE],
	internalCreditMap[CQWPX],
	internalCreditMap[CQWPX_BAND],
	internalCreditMap[CQWPX_MODE],
	internalCreditMap[DXCC],
	internalCreditMap[DXCC_BAND],
	internalCreditMap[DXCC_MODE],
	internalCreditMap[DXCC_SATELLITE],
	internalCreditMap[EAUSTRALIA],
	internalCreditMap[ECANADA],
	internalCreditMap[ECOUNTY_STATE],
	internalCreditMap[EDX],
	internalCreditMap[EDX100],
	internalCreditMap[EDX100_BAND],
	internalCreditMap[EDX100_MODE],
	internalCreditMap[EECHOLINK50],
	internalCreditMap[EGRID_BAND],
	internalCreditMap[EGRID_SATELLITE],
	internalCreditMap[EPFX300],
	internalCreditMap[EPFX300_MODE],
	internalCreditMap[EWAS],
	internalCreditMap[EWAS_BAND],
	internalCreditMap[EWAS_MODE],
	internalCreditMap[EWAS_SATELLITE],
	internalCreditMap[EZ40],
	internalCreditMap[EZ40_MODE],
	internalCreditMap[FFMA],
	internalCreditMap[IOTA],
	internalCreditMap[IOTA_BASIC],
	internalCreditMap[IOTA_CONT],
	internalCreditMap[IOTA_GROUP],
	internalCreditMap[RDA],
	internalCreditMap[USACA],
	internalCreditMap[VUCC_BAND],
	internalCreditMap[VUCC_SATELLITE],
	internalCreditMap[WAB],
	internalCreditMap[WAC],
	internalCreditMap[WAC_BAND],
	internalCreditMap[WAE],
	internalCreditMap[WAE_BAND],
	internalCreditMap[WAE_MODE],
	internalCreditMap[WAIP],
	internalCreditMap[WAIP_BAND],
	internalCreditMap[WAIP_MODE],
	internalCreditMap[WAS],
	internalCreditMap[WAS_BAND],
	internalCreditMap[WAS_EME],
	internalCreditMap[WAS_MODE],
	internalCreditMap[WAS_NOVICE],
	internalCreditMap[WAS_QRP],
	internalCreditMap[WAS_SATELLITE],
	internalCreditMap[WITUZ],
	internalCreditMap[WITUZ_BAND],
}

var internalCreditListCurrent = []Spec{
	internalCreditMap[CQDX],
	internalCreditMap[CQDXFIELD],
	internalCreditMap[CQDXFIELD_BAND],
	internalCreditMap[CQDXFIELD_MOBILE],
	internalCreditMap[CQDXFIELD_MODE],
	internalCreditMap[CQDXFIELD_QRP],
	internalCreditMap[CQDXFIELD_SATELLITE],
	internalCreditMap[CQDX_BAND],
	internalCreditMap[CQDX_MOBILE],
	internalCreditMap[CQDX_MODE],
	internalCreditMap[CQDX_QRP],
	internalCreditMap[CQDX_SATELLITE],
	internalCreditMap[CQWAZ_BAND],
	internalCreditMap[CQWAZ_EME],
	internalCreditMap[CQWAZ_MIXED],
	internalCreditMap[CQWAZ_MOBILE],
	internalCreditMap[CQWAZ_MODE],
	internalCreditMap[CQWAZ_QRP],
	internalCreditMap[CQWAZ_SATELLITE],
	internalCreditMap[CQWPX],
	internalCreditMap[CQWPX_BAND],
	internalCreditMap[CQWPX_MODE],
	internalCreditMap[DXCC],
	internalCreditMap[DXCC_BAND],
	internalCreditMap[DXCC_MODE],
	internalCreditMap[DXCC_SATELLITE],
	internalCreditMap[EAUSTRALIA],
	internalCreditMap[ECANADA],
	internalCreditMap[ECOUNTY_STATE],
	internalCreditMap[EDX],
	internalCreditMap[EDX100],
	internalCreditMap[EDX100_BAND],
	internalCreditMap[EDX100_MODE],
	internalCreditMap[EECHOLINK50],
	internalCreditMap[EGRID_BAND],
	internalCreditMap[EGRID_SATELLITE],
	internalCreditMap[EPFX300],
	internalCreditMap[EPFX300_MODE],
	internalCreditMap[EWAS],
	internalCreditMap[EWAS_BAND],
	internalCreditMap[EWAS_MODE],
	internalCreditMap[EWAS_SATELLITE],
	internalCreditMap[EZ40],
	internalCreditMap[EZ40_MODE],
	internalCreditMap[FFMA],
	internalCreditMap[IOTA],
	internalCreditMap[IOTA_BASIC],
	internalCreditMap[IOTA_CONT],
	internalCreditMap[IOTA_GROUP],
	internalCreditMap[RDA],
	internalCreditMap[USACA],
	internalCreditMap[VUCC_BAND],
	internalCreditMap[VUCC_SATELLITE],
	internalCreditMap[WAB],
	internalCreditMap[WAC],
	internalCreditMap[WAC_BAND],
	internalCreditMap[WAE],
	internalCreditMap[WAE_BAND],
	internalCreditMap[WAE_MODE],
	internalCreditMap[WAIP],
	internalCreditMap[WAIP_BAND],
	internalCreditMap[WAIP_MODE],
	internalCreditMap[WAS],
	internalCreditMap[WAS_BAND],
	internalCreditMap[WAS_EME],
	internalCreditMap[WAS_MODE],
	internalCreditMap[WAS_NOVICE],
	internalCreditMap[WAS_QRP],
	internalCreditMap[WAS_SATELLITE],
	internalCreditMap[WITUZ],
	internalCreditMap[WITUZ_BAND],
}
