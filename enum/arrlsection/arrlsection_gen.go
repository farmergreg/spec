// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package arrlsection provides code and constants as defined in ADIF 3.1.6 (Released)
package arrlsection

import (
	"sync"

	"github.com/hamradiolog-net/spec/v6/enum/dxccentitycode"
)

const (
	AB  ARRLSection = "ab"  // ab   = Alberta
	AK  ARRLSection = "ak"  // ak   = Alaska
	AL  ARRLSection = "al"  // al   = Alabama
	AR  ARRLSection = "ar"  // ar   = Arkansas
	AZ  ARRLSection = "az"  // az   = Arizona
	BC  ARRLSection = "bc"  // bc   = British Columbia
	CO  ARRLSection = "co"  // co   = Colorado
	CT  ARRLSection = "ct"  // ct   = Connecticut
	DE  ARRLSection = "de"  // de   = Delaware
	EB  ARRLSection = "eb"  // eb   = East Bay
	EMA ARRLSection = "ema" // ema  = Eastern Massachusetts
	ENY ARRLSection = "eny" // eny  = Eastern New York
	EPA ARRLSection = "epa" // epa  = Eastern Pennsylvania
	EWA ARRLSection = "ewa" // ewa  = Eastern Washington
	GA  ARRLSection = "ga"  // ga   = Georgia
	GH  ARRLSection = "gh"  // gh   = Golden Horseshoe
	GTA ARRLSection = "gta" // gta  = Greater Toronto Area
	IA  ARRLSection = "ia"  // ia   = Iowa
	ID  ARRLSection = "id"  // id   = Idaho
	IL  ARRLSection = "il"  // il   = Illinois
	IN  ARRLSection = "in"  // in   = Indiana
	KS  ARRLSection = "ks"  // ks   = Kansas
	KY  ARRLSection = "ky"  // ky   = Kentucky
	LA  ARRLSection = "la"  // la   = Louisiana
	LAX ARRLSection = "lax" // lax  = Los Angeles
	MAR ARRLSection = "mar" // mar  = Maritime
	MB  ARRLSection = "mb"  // mb   = Manitoba
	MDC ARRLSection = "mdc" // mdc  = Maryland-DC
	ME  ARRLSection = "me"  // me   = Maine
	MI  ARRLSection = "mi"  // mi   = Michigan
	MN  ARRLSection = "mn"  // mn   = Minnesota
	MO  ARRLSection = "mo"  // mo   = Missouri
	MS  ARRLSection = "ms"  // ms   = Mississippi
	MT  ARRLSection = "mt"  // mt   = Montana
	NB  ARRLSection = "nb"  // nb   = New Brunswick
	NC  ARRLSection = "nc"  // nc   = North Carolina
	ND  ARRLSection = "nd"  // nd   = North Dakota
	NE  ARRLSection = "ne"  // ne   = Nebraska
	NFL ARRLSection = "nfl" // nfl  = Northern Florida
	NH  ARRLSection = "nh"  // nh   = New Hampshire
	NL  ARRLSection = "nl"  // nl   = Newfoundland/Labrador
	NLI ARRLSection = "nli" // nli  = New York City-Long Island
	NM  ARRLSection = "nm"  // nm   = New Mexico
	NNJ ARRLSection = "nnj" // nnj  = Northern New Jersey
	NNY ARRLSection = "nny" // nny  = Northern New York
	NS  ARRLSection = "ns"  // ns   = Nova Scotia
	NT  ARRLSection = "nt"  // nt   = Northwest Territories/Yukon/Nunavut
	NTX ARRLSection = "ntx" // ntx  = North Texas
	NV  ARRLSection = "nv"  // nv   = Nevada
	NWT ARRLSection = "nwt" // nwt  = Northwest Territories/Yukon/Nunavut
	OH  ARRLSection = "oh"  // oh   = Ohio
	OK  ARRLSection = "ok"  // ok   = Oklahoma
	ON  ARRLSection = "on"  // on   = Ontario
	ONE ARRLSection = "one" // one  = Ontario East
	ONN ARRLSection = "onn" // onn  = Ontario North
	ONS ARRLSection = "ons" // ons  = Ontario South
	OR  ARRLSection = "or"  // or   = Oregon
	ORG ARRLSection = "org" // org  = Orange
	PAC ARRLSection = "pac" // pac  = Pacific
	PE  ARRLSection = "pe"  // pe   = Prince Edward Island
	PR  ARRLSection = "pr"  // pr   = Puerto Rico
	QC  ARRLSection = "qc"  // qc   = Quebec
	RI  ARRLSection = "ri"  // ri   = Rhode Island
	SB  ARRLSection = "sb"  // sb   = Santa Barbara
	SC  ARRLSection = "sc"  // sc   = South Carolina
	SCV ARRLSection = "scv" // scv  = Santa Clara Valley
	SD  ARRLSection = "sd"  // sd   = South Dakota
	SDG ARRLSection = "sdg" // sdg  = San Diego
	SF  ARRLSection = "sf"  // sf   = San Francisco
	SFL ARRLSection = "sfl" // sfl  = Southern Florida
	SJV ARRLSection = "sjv" // sjv  = San Joaquin Valley
	SK  ARRLSection = "sk"  // sk   = Saskatchewan
	SNJ ARRLSection = "snj" // snj  = Southern New Jersey
	STX ARRLSection = "stx" // stx  = South Texas
	SV  ARRLSection = "sv"  // sv   = Sacramento Valley
	TER ARRLSection = "ter" // ter  = Territories
	TN  ARRLSection = "tn"  // tn   = Tennessee
	UT  ARRLSection = "ut"  // ut   = Utah
	VA  ARRLSection = "va"  // va   = Virginia
	VI  ARRLSection = "vi"  // vi   = US Virgin Islands
	VT  ARRLSection = "vt"  // vt   = Vermont
	WCF ARRLSection = "wcf" // wcf  = West Central Florida
	WI  ARRLSection = "wi"  // wi   = Wisconsin
	WMA ARRLSection = "wma" // wma  = Western Massachusetts
	WNY ARRLSection = "wny" // wny  = Western New York
	WPA ARRLSection = "wpa" // wpa  = Western Pennsylvania
	WTX ARRLSection = "wtx" // wtx  = West Texas
	WV  ARRLSection = "wv"  // wv   = West Virginia
	WWA ARRLSection = "wwa" // wwa  = Western Washington
	WY  ARRLSection = "wy"  // wy   = Wyoming
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known ARRLSection specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Comments: "", Key: "ab", Description: "Alberta", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ak", Description: "Alaska", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{6}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "al", Description: "Alabama", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ar", Description: "Arkansas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "az", Description: "Arizona", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "bc", Description: "British Columbia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "co", Description: "Colorado", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ct", Description: "Connecticut", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "de", Description: "Delaware", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "eb", Description: "East Bay", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ema", Description: "Eastern Massachusetts", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "eny", Description: "Eastern New York", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "epa", Description: "Eastern Pennsylvania", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ewa", Description: "Eastern Washington", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ga", Description: "Georgia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "gh", Description: "Golden Horseshoe", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	{IsImportOnly: false, Comments: "replaced by GH", Key: "gta", Description: "Greater Toronto Area", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 1672531200},
	{IsImportOnly: false, Comments: "", Key: "ia", Description: "Iowa", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "id", Description: "Idaho", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "il", Description: "Illinois", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "in", Description: "Indiana", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ks", Description: "Kansas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ky", Description: "Kentucky", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "la", Description: "Louisiana", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "lax", Description: "Los Angeles", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "replaced by NB and NS", Key: "mar", Description: "Maritime", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 1672531200},
	{IsImportOnly: false, Comments: "", Key: "mb", Description: "Manitoba", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "mdc", Description: "Maryland-DC", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "me", Description: "Maine", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "mi", Description: "Michigan", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "mn", Description: "Minnesota", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "mo", Description: "Missouri", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ms", Description: "Mississippi", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "mt", Description: "Montana", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "nb", Description: "New Brunswick", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "nc", Description: "North Carolina", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "nd", Description: "North Dakota", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ne", Description: "Nebraska", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "nfl", Description: "Northern Florida", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "nh", Description: "New Hampshire", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "nl", Description: "Newfoundland/Labrador", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "nli", Description: "New York City-Long Island", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "nm", Description: "New Mexico", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "nnj", Description: "Northern New Jersey", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "nny", Description: "Northern New York", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ns", Description: "Nova Scotia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	{IsImportOnly: false, Comments: "replaced by TER", Key: "nt", Description: "Northwest Territories/Yukon/Nunavut", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1067644800, DeletedDate: 1672531200},
	{IsImportOnly: false, Comments: "", Key: "ntx", Description: "North Texas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "nv", Description: "Nevada", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "replaced by NT", Key: "nwt", Description: "Northwest Territories/Yukon/Nunavut", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 1067644800},
	{IsImportOnly: false, Comments: "", Key: "oh", Description: "Ohio", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ok", Description: "Oklahoma", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "replaced by GTA, ONE, ONN, and ONS", Key: "on", Description: "Ontario", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 1346457600},
	{IsImportOnly: false, Comments: "", Key: "one", Description: "Ontario East", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "onn", Description: "Ontario North", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ons", Description: "Ontario South", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1346457600, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "or", Description: "Oregon", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "org", Description: "Orange", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "pac", Description: "Pacific", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{9, 20, 103, 110, 123, 138, 166, 174, 197, 297, 515}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "pe", Description: "Prince Edward Island", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1585699200, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "pr", Description: "Puerto Rico", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{43, 202}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "qc", Description: "Quebec", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ri", Description: "Rhode Island", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "sb", Description: "Santa Barbara", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "sc", Description: "South Carolina", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "scv", Description: "Santa Clara Valley", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "sd", Description: "South Dakota", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "sdg", Description: "San Diego", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "sf", Description: "San Francisco", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "sfl", Description: "Southern Florida", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "sjv", Description: "San Joaquin Valley", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "sk", Description: "Saskatchewan", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "snj", Description: "Southern New Jersey", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "stx", Description: "South Texas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "sv", Description: "Sacramento Valley", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ter", Description: "Territories", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{1}, FromDate: 1672531200, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "tn", Description: "Tennessee", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "ut", Description: "Utah", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "va", Description: "Virginia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "vi", Description: "US Virgin Islands", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{105, 182, 285}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "vt", Description: "Vermont", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "wcf", Description: "West Central Florida", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "wi", Description: "Wisconsin", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "wma", Description: "Western Massachusetts", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "wny", Description: "Western New York", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "wpa", Description: "Western Pennsylvania", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "wtx", Description: "West Texas", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "wv", Description: "West Virginia", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "wwa", Description: "Western Washington", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
	{IsImportOnly: false, Comments: "", Key: "wy", Description: "Wyoming", DXCCEntityCode: dxccentitycode.DXCCEntityCodeList{291}, FromDate: 0, DeletedDate: 0},
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
func Lookup(a ARRLSection) (Spec, bool) {
	spec, ok := lookupMap[a]
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
