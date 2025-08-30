// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package region

var (
	AI_248 Region = "AI.248" //   AI.248 = African Italy
	BI_259 Region = "BI.259" //   BI.259 = Bear Island
	ET_390 Region = "ET.390" //   ET.390 = European Turkey
	IV_206 Region = "IV.206" //   IV.206 = ITU Vienna
	KO_0   Region = "KO.0"   //   KO.0   = Kosovo
	KO_296 Region = "KO.296" //   KO.296 = Kosovo
	KO_522 Region = "KO.522" //   KO.522 = Kosovo
	NONE   Region = "NONE"   // NONE.    = Not within a WAE or CQ region that is within a DXCC entity
	SI_279 Region = "SI.279" //   SI.279 = Shetland Islands
	SY_248 Region = "SY.248" //   SY.248 = Sicily
)

// A map of all Region specifications.
var RegionMap = map[Region]Spec{
	AI_248: {IsImportOnly: false, Comments: "", Key: "AI", DXCCEntityCode: "248", Region: "African Italy", Prefix: "IG9", Applicability: "CQ", StartDate: 0, EndDate: 0},
	BI_259: {IsImportOnly: false, Comments: "", Key: "BI", DXCCEntityCode: "259", Region: "Bear Island", Prefix: "JW/B", Applicability: "CQ, WAE", StartDate: 0, EndDate: 0},
	ET_390: {IsImportOnly: false, Comments: "", Key: "ET", DXCCEntityCode: "390", Region: "European Turkey", Prefix: "TA1", Applicability: "CQ", StartDate: 0, EndDate: 0},
	IV_206: {IsImportOnly: false, Comments: "", Key: "IV", DXCCEntityCode: "206", Region: "ITU Vienna", Prefix: "4U1V", Applicability: "CQ, WAE", StartDate: 0, EndDate: 0},
	KO_0:   {IsImportOnly: false, Comments: "", Key: "KO", DXCCEntityCode: "0", Region: "Kosovo", Prefix: "Z6", Applicability: "CQ, WAE", StartDate: 1347408000, EndDate: 1516406400},
	KO_296: {IsImportOnly: false, Comments: "", Key: "KO", DXCCEntityCode: "296", Region: "Kosovo", Prefix: "YU8", Applicability: "CQ, WAE", StartDate: 0, EndDate: 1347321600},
	KO_522: {IsImportOnly: false, Comments: "", Key: "KO", DXCCEntityCode: "522", Region: "Kosovo", Prefix: "Z6", Applicability: "CQ, WAE", StartDate: 1516492800, EndDate: 0},
	NONE:   {IsImportOnly: false, Comments: "", Key: "NONE", DXCCEntityCode: "", Region: "Not within a WAE or CQ region that is within a DXCC entity", Prefix: "", Applicability: "", StartDate: 0, EndDate: 0},
	SI_279: {IsImportOnly: false, Comments: "", Key: "SI", DXCCEntityCode: "279", Region: "Shetland Islands", Prefix: "GM/S", Applicability: "CQ, WAE", StartDate: 0, EndDate: 0},
	SY_248: {IsImportOnly: false, Comments: "", Key: "SY", DXCCEntityCode: "248", Region: "Sicily", Prefix: "IT9", Applicability: "CQ, WAE", StartDate: 0, EndDate: 0},
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
