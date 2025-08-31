// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package region

import "github.com/hamradiolog-net/adif-spec/v8/spectype"

// TODO : add primary key... but with de-duplication

var (
	AI_248 RegionCompositeKey = "AI.248" //   AI.248 = AI    African Italy  ; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap
	BI_259 RegionCompositeKey = "BI.259" //   BI.259 = BI    Bear Island    ; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap
	ET_390 RegionCompositeKey = "ET.390" //   ET.390 = ET    European Turkey; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap
	IV_206 RegionCompositeKey = "IV.206" //   IV.206 = IV    ITU Vienna     ; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap
	KO_0   RegionCompositeKey = "KO.0"   //   KO.0   = KO    Kosovo         ; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap
	KO_296 RegionCompositeKey = "KO.296" //   KO.296 = KO    Kosovo         ; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap
	KO_522 RegionCompositeKey = "KO.522" //   KO.522 = KO    Kosovo         ; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap
	NONE   RegionCompositeKey = "NONE"   // NONE.0   = NONE  Not within a WAE or CQ region that is within a DXCC entity; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap
	SI_279 RegionCompositeKey = "SI.279" //   SI.279 = SI    Shetland Islands; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap
	SY_248 RegionCompositeKey = "SY.248" //   SY.248 = SY    Sicily         ; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap
)

// A map of all RegionCompositeKey specifications.
var RegionCompositeKeyMap = map[RegionCompositeKey]Spec{
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

// All RegionCompositeKey specifications including depreciated and import only.
var RegionCompositeKeyListAll = []Spec{
	RegionCompositeKeyMap[AI_248],
	RegionCompositeKeyMap[BI_259],
	RegionCompositeKeyMap[ET_390],
	RegionCompositeKeyMap[IV_206],
	RegionCompositeKeyMap[KO_0],
	RegionCompositeKeyMap[KO_296],
	RegionCompositeKeyMap[KO_522],
	RegionCompositeKeyMap[NONE],
	RegionCompositeKeyMap[SI_279],
	RegionCompositeKeyMap[SY_248],
}

// All RegionCompositeKey specifications values that are NOT marked import-only.
var RegionCompositeKeyListCurrent = []Spec{
	RegionCompositeKeyMap[AI_248],
	RegionCompositeKeyMap[BI_259],
	RegionCompositeKeyMap[ET_390],
	RegionCompositeKeyMap[IV_206],
	RegionCompositeKeyMap[KO_0],
	RegionCompositeKeyMap[KO_296],
	RegionCompositeKeyMap[KO_522],
	RegionCompositeKeyMap[NONE],
	RegionCompositeKeyMap[SI_279],
	RegionCompositeKeyMap[SY_248],
}
