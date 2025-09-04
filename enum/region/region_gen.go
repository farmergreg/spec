// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package region provides code and constants as defined in ADIF 3.1.6 (Proposed)
package region

import "github.com/hamradiolog-net/adif-spec/v6/spectype"

/*
 * This enumeration has a composite key.
 * It behaves differently than most of the other enumerations.
 */

// A map of all RegionCompositeKey specifications.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var RegionCompositeKeyMap = map[RegionCompositeKey]Spec{
	"AI.248": {IsImportOnly: false, Code: "AI", DXCCEntityCode: 248, Region: "African Italy", Prefix: "IG9", Applicability: spectype.StringSlice{"CQ"}, StartDate: 0, EndDate: 0},
	"BI.259": {IsImportOnly: false, Code: "BI", DXCCEntityCode: 259, Region: "Bear Island", Prefix: "JW/B", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
	"ET.390": {IsImportOnly: false, Code: "ET", DXCCEntityCode: 390, Region: "European Turkey", Prefix: "TA1", Applicability: spectype.StringSlice{"CQ"}, StartDate: 0, EndDate: 0},
	"IV.206": {IsImportOnly: false, Code: "IV", DXCCEntityCode: 206, Region: "ITU Vienna", Prefix: "4U1V", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
	"KO.0":   {IsImportOnly: false, Code: "KO", DXCCEntityCode: 0, Region: "Kosovo", Prefix: "Z6", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 1347408000, EndDate: 1516406400},
	"KO.296": {IsImportOnly: false, Code: "KO", DXCCEntityCode: 296, Region: "Kosovo", Prefix: "YU8", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 1347321600},
	"KO.522": {IsImportOnly: false, Code: "KO", DXCCEntityCode: 522, Region: "Kosovo", Prefix: "Z6", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 1516492800, EndDate: 0},
	"NONE":   {IsImportOnly: false, Code: "NONE", DXCCEntityCode: 0, Region: "Not within a WAE or CQ region that is within a DXCC entity", Prefix: "", Applicability: spectype.StringSlice(nil), StartDate: 0, EndDate: 0},
	"SI.279": {IsImportOnly: false, Code: "SI", DXCCEntityCode: 279, Region: "Shetland Islands", Prefix: "GM/S", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
	"SY.248": {IsImportOnly: false, Code: "SY", DXCCEntityCode: 248, Region: "Sicily", Prefix: "IT9", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
}

// All RegionCompositeKey specifications including deprecated and import only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var RegionCompositeKeyListAll = []Spec{
	RegionCompositeKeyMap["AI.248"],
	RegionCompositeKeyMap["BI.259"],
	RegionCompositeKeyMap["ET.390"],
	RegionCompositeKeyMap["IV.206"],
	RegionCompositeKeyMap["KO.0"],
	RegionCompositeKeyMap["KO.296"],
	RegionCompositeKeyMap["KO.522"],
	RegionCompositeKeyMap["NONE"],
	RegionCompositeKeyMap["SI.279"],
	RegionCompositeKeyMap["SY.248"],
}

// All RegionCompositeKey specifications that are NOT marked import-only.
// For convenience, this data is mutable.
// If you require immutable data, please use the specdata package.
var RegionCompositeKeyListCurrent = []Spec{
	RegionCompositeKeyMap["AI.248"],
	RegionCompositeKeyMap["BI.259"],
	RegionCompositeKeyMap["ET.390"],
	RegionCompositeKeyMap["IV.206"],
	RegionCompositeKeyMap["KO.0"],
	RegionCompositeKeyMap["KO.296"],
	RegionCompositeKeyMap["KO.522"],
	RegionCompositeKeyMap["NONE"],
	RegionCompositeKeyMap["SI.279"],
	RegionCompositeKeyMap["SY.248"],
}
