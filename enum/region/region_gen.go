// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package region

import "github.com/hamradiolog-net/adif-spec/v3/spectype"

var (
	AI_248 Region = "AI"   //   AI.248 = AI    African Italy
	BI_259 Region = "BI"   //   BI.259 = BI    Bear Island
	ET_390 Region = "ET"   //   ET.390 = ET    European Turkey
	IV_206 Region = "IV"   //   IV.206 = IV    ITU Vienna
	KO_0   Region = "KO"   //   KO.0   = KO    Kosovo
	KO_296 Region = "KO"   //   KO.296 = KO    Kosovo
	KO_522 Region = "KO"   //   KO.522 = KO    Kosovo
	NONE   Region = "NONE" // NONE.0   = NONE  Not within a WAE or CQ region that is within a DXCC entity
	SI_279 Region = "SI"   //   SI.279 = SI    Shetland Islands
	SY_248 Region = "SY"   //   SY.248 = SY    Sicily
)

// A map of all Region specifications.
var RegionMap = map[Region]Spec{
	AI_248: {IsImportOnly: false, Key: "AI", DXCCEntityCode: 248, Region: "African Italy", Prefix: "IG9", Applicability: spectype.StringSlice{"CQ"}, StartDate: 0, EndDate: 0},
	BI_259: {IsImportOnly: false, Key: "BI", DXCCEntityCode: 259, Region: "Bear Island", Prefix: "JW/B", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
	ET_390: {IsImportOnly: false, Key: "ET", DXCCEntityCode: 390, Region: "European Turkey", Prefix: "TA1", Applicability: spectype.StringSlice{"CQ"}, StartDate: 0, EndDate: 0},
	IV_206: {IsImportOnly: false, Key: "IV", DXCCEntityCode: 206, Region: "ITU Vienna", Prefix: "4U1V", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
	KO_0:   {IsImportOnly: false, Key: "KO", DXCCEntityCode: 0, Region: "Kosovo", Prefix: "Z6", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 1347408000, EndDate: 1516406400},
	KO_296: {IsImportOnly: false, Key: "KO", DXCCEntityCode: 296, Region: "Kosovo", Prefix: "YU8", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 1347321600},
	KO_522: {IsImportOnly: false, Key: "KO", DXCCEntityCode: 522, Region: "Kosovo", Prefix: "Z6", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 1516492800, EndDate: 0},
	NONE:   {IsImportOnly: false, Key: "NONE", DXCCEntityCode: 0, Region: "Not within a WAE or CQ region that is within a DXCC entity", Prefix: "", Applicability: spectype.StringSlice(nil), StartDate: 0, EndDate: 0},
	SI_279: {IsImportOnly: false, Key: "SI", DXCCEntityCode: 279, Region: "Shetland Islands", Prefix: "GM/S", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
	SY_248: {IsImportOnly: false, Key: "SY", DXCCEntityCode: 248, Region: "Sicily", Prefix: "IT9", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
}

// All Region specifications including depreciated and import only.
var RegionListAll = []Spec{
	RegionMap[AI_248],
	RegionMap[BI_259],
	RegionMap[ET_390],
	RegionMap[IV_206],
	RegionMap[KO_0],
	RegionMap[KO_296],
	RegionMap[KO_522],
	RegionMap[NONE],
	RegionMap[SI_279],
	RegionMap[SY_248],
}

// All Region specifications values that are NOT marked import-only.
var RegionListCurrent = []Spec{
	RegionMap[AI_248],
	RegionMap[BI_259],
	RegionMap[ET_390],
	RegionMap[IV_206],
	RegionMap[KO_0],
	RegionMap[KO_296],
	RegionMap[KO_522],
	RegionMap[NONE],
	RegionMap[SI_279],
	RegionMap[SY_248],
}
