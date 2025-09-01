// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package arrlsection provides code and constants as defined in ADIF 3.1.6 (Proposed)
package arrlsection

import "github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"

const (
	AB  ARRLSection = "AB"  // AB   = Alberta
	AK  ARRLSection = "AK"  // AK   = Alaska
	AL  ARRLSection = "AL"  // AL   = Alabama
	AR  ARRLSection = "AR"  // AR   = Arkansas
	AZ  ARRLSection = "AZ"  // AZ   = Arizona
	BC  ARRLSection = "BC"  // BC   = British Columbia
	CO  ARRLSection = "CO"  // CO   = Colorado
	CT  ARRLSection = "CT"  // CT   = Connecticut
	DE  ARRLSection = "DE"  // DE   = Delaware
	EB  ARRLSection = "EB"  // EB   = East Bay
	EMA ARRLSection = "EMA" // EMA  = Eastern Massachusetts
	ENY ARRLSection = "ENY" // ENY  = Eastern New York
	EPA ARRLSection = "EPA" // EPA  = Eastern Pennsylvania
	EWA ARRLSection = "EWA" // EWA  = Eastern Washington
	GA  ARRLSection = "GA"  // GA   = Georgia
	GH  ARRLSection = "GH"  // GH   = Golden Horseshoe
	GTA ARRLSection = "GTA" // GTA  = Greater Toronto Area
	IA  ARRLSection = "IA"  // IA   = Iowa
	ID  ARRLSection = "ID"  // ID   = Idaho
	IL  ARRLSection = "IL"  // IL   = Illinois
	IN  ARRLSection = "IN"  // IN   = Indiana
	KS  ARRLSection = "KS"  // KS   = Kansas
	KY  ARRLSection = "KY"  // KY   = Kentucky
	LA  ARRLSection = "LA"  // LA   = Louisiana
	LAX ARRLSection = "LAX" // LAX  = Los Angeles
	MAR ARRLSection = "MAR" // MAR  = Maritime
	MB  ARRLSection = "MB"  // MB   = Manitoba
	MDC ARRLSection = "MDC" // MDC  = Maryland-DC
	ME  ARRLSection = "ME"  // ME   = Maine
	MI  ARRLSection = "MI"  // MI   = Michigan
	MN  ARRLSection = "MN"  // MN   = Minnesota
	MO  ARRLSection = "MO"  // MO   = Missouri
	MS  ARRLSection = "MS"  // MS   = Mississippi
	MT  ARRLSection = "MT"  // MT   = Montana
	NB  ARRLSection = "NB"  // NB   = New Brunswick
	NC  ARRLSection = "NC"  // NC   = North Carolina
	ND  ARRLSection = "ND"  // ND   = North Dakota
	NE  ARRLSection = "NE"  // NE   = Nebraska
	NFL ARRLSection = "NFL" // NFL  = Northern Florida
	NH  ARRLSection = "NH"  // NH   = New Hampshire
	NL  ARRLSection = "NL"  // NL   = Newfoundland/Labrador
	NLI ARRLSection = "NLI" // NLI  = New York City-Long Island
	NM  ARRLSection = "NM"  // NM   = New Mexico
	NNJ ARRLSection = "NNJ" // NNJ  = Northern New Jersey
	NNY ARRLSection = "NNY" // NNY  = Northern New York
	NS  ARRLSection = "NS"  // NS   = Nova Scotia
	NT  ARRLSection = "NT"  // NT   = Northwest Territories/Yukon/Nunavut
	NTX ARRLSection = "NTX" // NTX  = North Texas
	NV  ARRLSection = "NV"  // NV   = Nevada
	NWT ARRLSection = "NWT" // NWT  = Northwest Territories/Yukon/Nunavut
	OH  ARRLSection = "OH"  // OH   = Ohio
	OK  ARRLSection = "OK"  // OK   = Oklahoma
	ON  ARRLSection = "ON"  // ON   = Ontario
	ONE ARRLSection = "ONE" // ONE  = Ontario East
	ONN ARRLSection = "ONN" // ONN  = Ontario North
	ONS ARRLSection = "ONS" // ONS  = Ontario South
	OR  ARRLSection = "OR"  // OR   = Oregon
	ORG ARRLSection = "ORG" // ORG  = Orange
	PAC ARRLSection = "PAC" // PAC  = Pacific
	PE  ARRLSection = "PE"  // PE   = Prince Edward Island
	PR  ARRLSection = "PR"  // PR   = Puerto Rico
	QC  ARRLSection = "QC"  // QC   = Quebec
	RI  ARRLSection = "RI"  // RI   = Rhode Island
	SB  ARRLSection = "SB"  // SB   = Santa Barbara
	SC  ARRLSection = "SC"  // SC   = South Carolina
	SCV ARRLSection = "SCV" // SCV  = Santa Clara Valley
	SD  ARRLSection = "SD"  // SD   = South Dakota
	SDG ARRLSection = "SDG" // SDG  = San Diego
	SF  ARRLSection = "SF"  // SF   = San Francisco
	SFL ARRLSection = "SFL" // SFL  = Southern Florida
	SJV ARRLSection = "SJV" // SJV  = San Joaquin Valley
	SK  ARRLSection = "SK"  // SK   = Saskatchewan
	SNJ ARRLSection = "SNJ" // SNJ  = Southern New Jersey
	STX ARRLSection = "STX" // STX  = South Texas
	SV  ARRLSection = "SV"  // SV   = Sacramento Valley
	TER ARRLSection = "TER" // TER  = Territories
	TN  ARRLSection = "TN"  // TN   = Tennessee
	UT  ARRLSection = "UT"  // UT   = Utah
	VA  ARRLSection = "VA"  // VA   = Virginia
	VI  ARRLSection = "VI"  // VI   = US Virgin Islands
	VT  ARRLSection = "VT"  // VT   = Vermont
	WCF ARRLSection = "WCF" // WCF  = West Central Florida
	WI  ARRLSection = "WI"  // WI   = Wisconsin
	WMA ARRLSection = "WMA" // WMA  = Western Massachusetts
	WNY ARRLSection = "WNY" // WNY  = Western New York
	WPA ARRLSection = "WPA" // WPA  = Western Pennsylvania
	WTX ARRLSection = "WTX" // WTX  = West Texas
	WV  ARRLSection = "WV"  // WV   = West Virginia
	WWA ARRLSection = "WWA" // WWA  = Western Washington
	WY  ARRLSection = "WY"  // WY   = Wyoming
)

// A map of all ARRLSection specifications.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var ARRLSectionMap = map[ARRLSection]Spec{
	AB:  {IsImportOnly: false, Comments: "", Key: "AB", Description: "Alberta", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	AK:  {IsImportOnly: false, Comments: "", Key: "AK", Description: "Alaska", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{6}, FromDate: 0, DeletedDate: 0},
	AL:  {IsImportOnly: false, Comments: "", Key: "AL", Description: "Alabama", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	AR:  {IsImportOnly: false, Comments: "", Key: "AR", Description: "Arkansas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	AZ:  {IsImportOnly: false, Comments: "", Key: "AZ", Description: "Arizona", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	BC:  {IsImportOnly: false, Comments: "", Key: "BC", Description: "British Columbia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	CO:  {IsImportOnly: false, Comments: "", Key: "CO", Description: "Colorado", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	CT:  {IsImportOnly: false, Comments: "", Key: "CT", Description: "Connecticut", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	DE:  {IsImportOnly: false, Comments: "", Key: "DE", Description: "Delaware", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	EB:  {IsImportOnly: false, Comments: "", Key: "EB", Description: "East Bay", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	EMA: {IsImportOnly: false, Comments: "", Key: "EMA", Description: "Eastern Massachusetts", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	ENY: {IsImportOnly: false, Comments: "", Key: "ENY", Description: "Eastern New York", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	EPA: {IsImportOnly: false, Comments: "", Key: "EPA", Description: "Eastern Pennsylvania", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	EWA: {IsImportOnly: false, Comments: "", Key: "EWA", Description: "Eastern Washington", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	GA:  {IsImportOnly: false, Comments: "", Key: "GA", Description: "Georgia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	GH:  {IsImportOnly: false, Comments: "", Key: "GH", Description: "Golden Horseshoe", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	GTA: {IsImportOnly: false, Comments: "replaced by GH", Key: "GTA", Description: "Greater Toronto Area", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 1672531200},
	IA:  {IsImportOnly: false, Comments: "", Key: "IA", Description: "Iowa", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	ID:  {IsImportOnly: false, Comments: "", Key: "ID", Description: "Idaho", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	IL:  {IsImportOnly: false, Comments: "", Key: "IL", Description: "Illinois", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	IN:  {IsImportOnly: false, Comments: "", Key: "IN", Description: "Indiana", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	KS:  {IsImportOnly: false, Comments: "", Key: "KS", Description: "Kansas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	KY:  {IsImportOnly: false, Comments: "", Key: "KY", Description: "Kentucky", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	LA:  {IsImportOnly: false, Comments: "", Key: "LA", Description: "Louisiana", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	LAX: {IsImportOnly: false, Comments: "", Key: "LAX", Description: "Los Angeles", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	MAR: {IsImportOnly: false, Comments: "replaced by NB and NS", Key: "MAR", Description: "Maritime", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 1672531200},
	MB:  {IsImportOnly: false, Comments: "", Key: "MB", Description: "Manitoba", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	MDC: {IsImportOnly: false, Comments: "", Key: "MDC", Description: "Maryland-DC", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	ME:  {IsImportOnly: false, Comments: "", Key: "ME", Description: "Maine", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	MI:  {IsImportOnly: false, Comments: "", Key: "MI", Description: "Michigan", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	MN:  {IsImportOnly: false, Comments: "", Key: "MN", Description: "Minnesota", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	MO:  {IsImportOnly: false, Comments: "", Key: "MO", Description: "Missouri", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	MS:  {IsImportOnly: false, Comments: "", Key: "MS", Description: "Mississippi", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	MT:  {IsImportOnly: false, Comments: "", Key: "MT", Description: "Montana", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	NB:  {IsImportOnly: false, Comments: "", Key: "NB", Description: "New Brunswick", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	NC:  {IsImportOnly: false, Comments: "", Key: "NC", Description: "North Carolina", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	ND:  {IsImportOnly: false, Comments: "", Key: "ND", Description: "North Dakota", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	NE:  {IsImportOnly: false, Comments: "", Key: "NE", Description: "Nebraska", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	NFL: {IsImportOnly: false, Comments: "", Key: "NFL", Description: "Northern Florida", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	NH:  {IsImportOnly: false, Comments: "", Key: "NH", Description: "New Hampshire", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	NL:  {IsImportOnly: false, Comments: "", Key: "NL", Description: "Newfoundland/Labrador", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	NLI: {IsImportOnly: false, Comments: "", Key: "NLI", Description: "New York City-Long Island", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	NM:  {IsImportOnly: false, Comments: "", Key: "NM", Description: "New Mexico", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	NNJ: {IsImportOnly: false, Comments: "", Key: "NNJ", Description: "Northern New Jersey", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	NNY: {IsImportOnly: false, Comments: "", Key: "NNY", Description: "Northern New York", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	NS:  {IsImportOnly: false, Comments: "", Key: "NS", Description: "Nova Scotia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	NT:  {IsImportOnly: false, Comments: "replaced by TER", Key: "NT", Description: "Northwest Territories/Yukon/Nunavut", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1067644800, DeletedDate: 1672531200},
	NTX: {IsImportOnly: false, Comments: "", Key: "NTX", Description: "North Texas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	NV:  {IsImportOnly: false, Comments: "", Key: "NV", Description: "Nevada", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	NWT: {IsImportOnly: false, Comments: "replaced by NT", Key: "NWT", Description: "Northwest Territories/Yukon/Nunavut", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 1067644800},
	OH:  {IsImportOnly: false, Comments: "", Key: "OH", Description: "Ohio", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	OK:  {IsImportOnly: false, Comments: "", Key: "OK", Description: "Oklahoma", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	ON:  {IsImportOnly: false, Comments: "replaced by GTA, ONE, ONN, and ONS", Key: "ON", Description: "Ontario", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 1346457600},
	ONE: {IsImportOnly: false, Comments: "", Key: "ONE", Description: "Ontario East", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 0},
	ONN: {IsImportOnly: false, Comments: "", Key: "ONN", Description: "Ontario North", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 0},
	ONS: {IsImportOnly: false, Comments: "", Key: "ONS", Description: "Ontario South", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 0},
	OR:  {IsImportOnly: false, Comments: "", Key: "OR", Description: "Oregon", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	ORG: {IsImportOnly: false, Comments: "", Key: "ORG", Description: "Orange", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	PAC: {IsImportOnly: false, Comments: "", Key: "PAC", Description: "Pacific", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{9, 20, 103, 110, 123, 138, 166, 174, 197, 297, 515}, FromDate: 0, DeletedDate: 0},
	PE:  {IsImportOnly: false, Comments: "", Key: "PE", Description: "Prince Edward Island", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1585699200, DeletedDate: 0},
	PR:  {IsImportOnly: false, Comments: "", Key: "PR", Description: "Puerto Rico", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{43, 202}, FromDate: 0, DeletedDate: 0},
	QC:  {IsImportOnly: false, Comments: "", Key: "QC", Description: "Quebec", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	RI:  {IsImportOnly: false, Comments: "", Key: "RI", Description: "Rhode Island", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	SB:  {IsImportOnly: false, Comments: "", Key: "SB", Description: "Santa Barbara", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	SC:  {IsImportOnly: false, Comments: "", Key: "SC", Description: "South Carolina", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	SCV: {IsImportOnly: false, Comments: "", Key: "SCV", Description: "Santa Clara Valley", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	SD:  {IsImportOnly: false, Comments: "", Key: "SD", Description: "South Dakota", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	SDG: {IsImportOnly: false, Comments: "", Key: "SDG", Description: "San Diego", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	SF:  {IsImportOnly: false, Comments: "", Key: "SF", Description: "San Francisco", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	SFL: {IsImportOnly: false, Comments: "", Key: "SFL", Description: "Southern Florida", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	SJV: {IsImportOnly: false, Comments: "", Key: "SJV", Description: "San Joaquin Valley", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	SK:  {IsImportOnly: false, Comments: "", Key: "SK", Description: "Saskatchewan", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	SNJ: {IsImportOnly: false, Comments: "", Key: "SNJ", Description: "Southern New Jersey", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	STX: {IsImportOnly: false, Comments: "", Key: "STX", Description: "South Texas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	SV:  {IsImportOnly: false, Comments: "", Key: "SV", Description: "Sacramento Valley", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	TER: {IsImportOnly: false, Comments: "", Key: "TER", Description: "Territories", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	TN:  {IsImportOnly: false, Comments: "", Key: "TN", Description: "Tennessee", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	UT:  {IsImportOnly: false, Comments: "", Key: "UT", Description: "Utah", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	VA:  {IsImportOnly: false, Comments: "", Key: "VA", Description: "Virginia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	VI:  {IsImportOnly: false, Comments: "", Key: "VI", Description: "US Virgin Islands", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{105, 182, 285}, FromDate: 0, DeletedDate: 0},
	VT:  {IsImportOnly: false, Comments: "", Key: "VT", Description: "Vermont", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	WCF: {IsImportOnly: false, Comments: "", Key: "WCF", Description: "West Central Florida", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	WI:  {IsImportOnly: false, Comments: "", Key: "WI", Description: "Wisconsin", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	WMA: {IsImportOnly: false, Comments: "", Key: "WMA", Description: "Western Massachusetts", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	WNY: {IsImportOnly: false, Comments: "", Key: "WNY", Description: "Western New York", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	WPA: {IsImportOnly: false, Comments: "", Key: "WPA", Description: "Western Pennsylvania", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	WTX: {IsImportOnly: false, Comments: "", Key: "WTX", Description: "West Texas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	WV:  {IsImportOnly: false, Comments: "", Key: "WV", Description: "West Virginia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	WWA: {IsImportOnly: false, Comments: "", Key: "WWA", Description: "Western Washington", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	WY:  {IsImportOnly: false, Comments: "", Key: "WY", Description: "Wyoming", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
}

// All ARRLSection specifications including depreciated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var ARRLSectionListAll = []Spec{
	ARRLSectionMap[AB],
	ARRLSectionMap[AK],
	ARRLSectionMap[AL],
	ARRLSectionMap[AR],
	ARRLSectionMap[AZ],
	ARRLSectionMap[BC],
	ARRLSectionMap[CO],
	ARRLSectionMap[CT],
	ARRLSectionMap[DE],
	ARRLSectionMap[EB],
	ARRLSectionMap[EMA],
	ARRLSectionMap[ENY],
	ARRLSectionMap[EPA],
	ARRLSectionMap[EWA],
	ARRLSectionMap[GA],
	ARRLSectionMap[GH],
	ARRLSectionMap[GTA],
	ARRLSectionMap[IA],
	ARRLSectionMap[ID],
	ARRLSectionMap[IL],
	ARRLSectionMap[IN],
	ARRLSectionMap[KS],
	ARRLSectionMap[KY],
	ARRLSectionMap[LA],
	ARRLSectionMap[LAX],
	ARRLSectionMap[MAR],
	ARRLSectionMap[MB],
	ARRLSectionMap[MDC],
	ARRLSectionMap[ME],
	ARRLSectionMap[MI],
	ARRLSectionMap[MN],
	ARRLSectionMap[MO],
	ARRLSectionMap[MS],
	ARRLSectionMap[MT],
	ARRLSectionMap[NB],
	ARRLSectionMap[NC],
	ARRLSectionMap[ND],
	ARRLSectionMap[NE],
	ARRLSectionMap[NFL],
	ARRLSectionMap[NH],
	ARRLSectionMap[NL],
	ARRLSectionMap[NLI],
	ARRLSectionMap[NM],
	ARRLSectionMap[NNJ],
	ARRLSectionMap[NNY],
	ARRLSectionMap[NS],
	ARRLSectionMap[NT],
	ARRLSectionMap[NTX],
	ARRLSectionMap[NV],
	ARRLSectionMap[NWT],
	ARRLSectionMap[OH],
	ARRLSectionMap[OK],
	ARRLSectionMap[ON],
	ARRLSectionMap[ONE],
	ARRLSectionMap[ONN],
	ARRLSectionMap[ONS],
	ARRLSectionMap[OR],
	ARRLSectionMap[ORG],
	ARRLSectionMap[PAC],
	ARRLSectionMap[PE],
	ARRLSectionMap[PR],
	ARRLSectionMap[QC],
	ARRLSectionMap[RI],
	ARRLSectionMap[SB],
	ARRLSectionMap[SC],
	ARRLSectionMap[SCV],
	ARRLSectionMap[SD],
	ARRLSectionMap[SDG],
	ARRLSectionMap[SF],
	ARRLSectionMap[SFL],
	ARRLSectionMap[SJV],
	ARRLSectionMap[SK],
	ARRLSectionMap[SNJ],
	ARRLSectionMap[STX],
	ARRLSectionMap[SV],
	ARRLSectionMap[TER],
	ARRLSectionMap[TN],
	ARRLSectionMap[UT],
	ARRLSectionMap[VA],
	ARRLSectionMap[VI],
	ARRLSectionMap[VT],
	ARRLSectionMap[WCF],
	ARRLSectionMap[WI],
	ARRLSectionMap[WMA],
	ARRLSectionMap[WNY],
	ARRLSectionMap[WPA],
	ARRLSectionMap[WTX],
	ARRLSectionMap[WV],
	ARRLSectionMap[WWA],
	ARRLSectionMap[WY],
}

// All ARRLSection specifications values that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var ARRLSectionListCurrent = []Spec{
	ARRLSectionMap[AB],
	ARRLSectionMap[AK],
	ARRLSectionMap[AL],
	ARRLSectionMap[AR],
	ARRLSectionMap[AZ],
	ARRLSectionMap[BC],
	ARRLSectionMap[CO],
	ARRLSectionMap[CT],
	ARRLSectionMap[DE],
	ARRLSectionMap[EB],
	ARRLSectionMap[EMA],
	ARRLSectionMap[ENY],
	ARRLSectionMap[EPA],
	ARRLSectionMap[EWA],
	ARRLSectionMap[GA],
	ARRLSectionMap[GH],
	ARRLSectionMap[GTA],
	ARRLSectionMap[IA],
	ARRLSectionMap[ID],
	ARRLSectionMap[IL],
	ARRLSectionMap[IN],
	ARRLSectionMap[KS],
	ARRLSectionMap[KY],
	ARRLSectionMap[LA],
	ARRLSectionMap[LAX],
	ARRLSectionMap[MAR],
	ARRLSectionMap[MB],
	ARRLSectionMap[MDC],
	ARRLSectionMap[ME],
	ARRLSectionMap[MI],
	ARRLSectionMap[MN],
	ARRLSectionMap[MO],
	ARRLSectionMap[MS],
	ARRLSectionMap[MT],
	ARRLSectionMap[NB],
	ARRLSectionMap[NC],
	ARRLSectionMap[ND],
	ARRLSectionMap[NE],
	ARRLSectionMap[NFL],
	ARRLSectionMap[NH],
	ARRLSectionMap[NL],
	ARRLSectionMap[NLI],
	ARRLSectionMap[NM],
	ARRLSectionMap[NNJ],
	ARRLSectionMap[NNY],
	ARRLSectionMap[NS],
	ARRLSectionMap[NT],
	ARRLSectionMap[NTX],
	ARRLSectionMap[NV],
	ARRLSectionMap[NWT],
	ARRLSectionMap[OH],
	ARRLSectionMap[OK],
	ARRLSectionMap[ON],
	ARRLSectionMap[ONE],
	ARRLSectionMap[ONN],
	ARRLSectionMap[ONS],
	ARRLSectionMap[OR],
	ARRLSectionMap[ORG],
	ARRLSectionMap[PAC],
	ARRLSectionMap[PE],
	ARRLSectionMap[PR],
	ARRLSectionMap[QC],
	ARRLSectionMap[RI],
	ARRLSectionMap[SB],
	ARRLSectionMap[SC],
	ARRLSectionMap[SCV],
	ARRLSectionMap[SD],
	ARRLSectionMap[SDG],
	ARRLSectionMap[SF],
	ARRLSectionMap[SFL],
	ARRLSectionMap[SJV],
	ARRLSectionMap[SK],
	ARRLSectionMap[SNJ],
	ARRLSectionMap[STX],
	ARRLSectionMap[SV],
	ARRLSectionMap[TER],
	ARRLSectionMap[TN],
	ARRLSectionMap[UT],
	ARRLSectionMap[VA],
	ARRLSectionMap[VI],
	ARRLSectionMap[VT],
	ARRLSectionMap[WCF],
	ARRLSectionMap[WI],
	ARRLSectionMap[WMA],
	ARRLSectionMap[WNY],
	ARRLSectionMap[WPA],
	ARRLSectionMap[WTX],
	ARRLSectionMap[WV],
	ARRLSectionMap[WWA],
	ARRLSectionMap[WY],
}
