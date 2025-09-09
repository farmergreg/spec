// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package arrlsection provides code and constants as defined in ADIF 3.1.6 (Proposed)
package arrlsection

import (
	"sync"

	"github.com/hamradiolog-net/spec/v6/enum/dxccentitycode"
)

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

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known ARRLSection specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Comments: "", Key: "AB", Description: "Alberta", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "AK", Description: "Alaska", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{6}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "AL", Description: "Alabama", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "AR", Description: "Arkansas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "AZ", Description: "Arizona", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "BC", Description: "British Columbia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "CO", Description: "Colorado", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "CT", Description: "Connecticut", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "DE", Description: "Delaware", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "EB", Description: "East Bay", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "EMA", Description: "Eastern Massachusetts", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ENY", Description: "Eastern New York", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "EPA", Description: "Eastern Pennsylvania", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "EWA", Description: "Eastern Washington", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "GA", Description: "Georgia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "GH", Description: "Golden Horseshoe", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	{IsImportOnly: false, Comments: "replaced by GH", Key: "GTA", Description: "Greater Toronto Area", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 1672531200},
	{IsImportOnly: false, Comments: "", Key: "IA", Description: "Iowa", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ID", Description: "Idaho", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "IL", Description: "Illinois", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "IN", Description: "Indiana", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "KS", Description: "Kansas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "KY", Description: "Kentucky", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "LA", Description: "Louisiana", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "LAX", Description: "Los Angeles", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "replaced by NB and NS", Key: "MAR", Description: "Maritime", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 1672531200},
	{IsImportOnly: false, Comments: "", Key: "MB", Description: "Manitoba", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "MDC", Description: "Maryland-DC", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ME", Description: "Maine", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "MI", Description: "Michigan", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "MN", Description: "Minnesota", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "MO", Description: "Missouri", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "MS", Description: "Mississippi", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "MT", Description: "Montana", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NB", Description: "New Brunswick", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NC", Description: "North Carolina", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ND", Description: "North Dakota", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NE", Description: "Nebraska", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NFL", Description: "Northern Florida", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NH", Description: "New Hampshire", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NL", Description: "Newfoundland/Labrador", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NLI", Description: "New York City-Long Island", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NM", Description: "New Mexico", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NNJ", Description: "Northern New Jersey", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NNY", Description: "Northern New York", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NS", Description: "Nova Scotia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	{IsImportOnly: false, Comments: "replaced by TER", Key: "NT", Description: "Northwest Territories/Yukon/Nunavut", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1067644800, DeletedDate: 1672531200},
	{IsImportOnly: false, Comments: "", Key: "NTX", Description: "North Texas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "NV", Description: "Nevada", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "replaced by NT", Key: "NWT", Description: "Northwest Territories/Yukon/Nunavut", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 1067644800},
	{IsImportOnly: false, Comments: "", Key: "OH", Description: "Ohio", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "OK", Description: "Oklahoma", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "replaced by GTA, ONE, ONN, and ONS", Key: "ON", Description: "Ontario", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 1346457600},
	{IsImportOnly: false, Comments: "", Key: "ONE", Description: "Ontario East", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ONN", Description: "Ontario North", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ONS", Description: "Ontario South", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "OR", Description: "Oregon", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ORG", Description: "Orange", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "PAC", Description: "Pacific", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{9, 20, 103, 110, 123, 138, 166, 174, 197, 297, 515}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "PE", Description: "Prince Edward Island", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1585699200, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "PR", Description: "Puerto Rico", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{43, 202}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "QC", Description: "Quebec", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "RI", Description: "Rhode Island", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "SB", Description: "Santa Barbara", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "SC", Description: "South Carolina", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "SCV", Description: "Santa Clara Valley", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "SD", Description: "South Dakota", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "SDG", Description: "San Diego", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "SF", Description: "San Francisco", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "SFL", Description: "Southern Florida", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "SJV", Description: "San Joaquin Valley", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "SK", Description: "Saskatchewan", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "SNJ", Description: "Southern New Jersey", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "STX", Description: "South Texas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "SV", Description: "Sacramento Valley", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "TER", Description: "Territories", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "TN", Description: "Tennessee", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "UT", Description: "Utah", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "VA", Description: "Virginia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "VI", Description: "US Virgin Islands", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{105, 182, 285}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "VT", Description: "Vermont", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "WCF", Description: "West Central Florida", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "WI", Description: "Wisconsin", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "WMA", Description: "Western Massachusetts", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "WNY", Description: "Western New York", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "WPA", Description: "Western Pennsylvania", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "WTX", Description: "West Texas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "WV", Description: "West Virginia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "WWA", Description: "Western Washington", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "WY", Description: "Wyoming", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
}

// lookupMap contains all known ARRLSection specifications.
var lookupMap = map[ARRLSection]*Spec{
	AB:  &lookupList[0],
	AK:  &lookupList[1],
	AL:  &lookupList[2],
	AR:  &lookupList[3],
	AZ:  &lookupList[4],
	BC:  &lookupList[5],
	CO:  &lookupList[6],
	CT:  &lookupList[7],
	DE:  &lookupList[8],
	EB:  &lookupList[9],
	EMA: &lookupList[10],
	ENY: &lookupList[11],
	EPA: &lookupList[12],
	EWA: &lookupList[13],
	GA:  &lookupList[14],
	GH:  &lookupList[15],
	GTA: &lookupList[16],
	IA:  &lookupList[17],
	ID:  &lookupList[18],
	IL:  &lookupList[19],
	IN:  &lookupList[20],
	KS:  &lookupList[21],
	KY:  &lookupList[22],
	LA:  &lookupList[23],
	LAX: &lookupList[24],
	MAR: &lookupList[25],
	MB:  &lookupList[26],
	MDC: &lookupList[27],
	ME:  &lookupList[28],
	MI:  &lookupList[29],
	MN:  &lookupList[30],
	MO:  &lookupList[31],
	MS:  &lookupList[32],
	MT:  &lookupList[33],
	NB:  &lookupList[34],
	NC:  &lookupList[35],
	ND:  &lookupList[36],
	NE:  &lookupList[37],
	NFL: &lookupList[38],
	NH:  &lookupList[39],
	NL:  &lookupList[40],
	NLI: &lookupList[41],
	NM:  &lookupList[42],
	NNJ: &lookupList[43],
	NNY: &lookupList[44],
	NS:  &lookupList[45],
	NT:  &lookupList[46],
	NTX: &lookupList[47],
	NV:  &lookupList[48],
	NWT: &lookupList[49],
	OH:  &lookupList[50],
	OK:  &lookupList[51],
	ON:  &lookupList[52],
	ONE: &lookupList[53],
	ONN: &lookupList[54],
	ONS: &lookupList[55],
	OR:  &lookupList[56],
	ORG: &lookupList[57],
	PAC: &lookupList[58],
	PE:  &lookupList[59],
	PR:  &lookupList[60],
	QC:  &lookupList[61],
	RI:  &lookupList[62],
	SB:  &lookupList[63],
	SC:  &lookupList[64],
	SCV: &lookupList[65],
	SD:  &lookupList[66],
	SDG: &lookupList[67],
	SF:  &lookupList[68],
	SFL: &lookupList[69],
	SJV: &lookupList[70],
	SK:  &lookupList[71],
	SNJ: &lookupList[72],
	STX: &lookupList[73],
	SV:  &lookupList[74],
	TER: &lookupList[75],
	TN:  &lookupList[76],
	UT:  &lookupList[77],
	VA:  &lookupList[78],
	VI:  &lookupList[79],
	VT:  &lookupList[80],
	WCF: &lookupList[81],
	WI:  &lookupList[82],
	WMA: &lookupList[83],
	WNY: &lookupList[84],
	WPA: &lookupList[85],
	WTX: &lookupList[86],
	WV:  &lookupList[87],
	WWA: &lookupList[88],
	WY:  &lookupList[89],
}

// Lookup returns the specification for the provided ARRLSection.
// ADIF 3.1.6
func Lookup(arrlsection ARRLSection) (Spec, bool) {
	spec, ok := lookupMap[arrlsection]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all ARRLSection specifications that match the provided filter function.
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

// List returns all ARRLSection specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns ARRLSection specifications.
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
