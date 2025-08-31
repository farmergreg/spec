// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package region provides code and constants as defined in ADIF 3.1.6 (Proposed)
package region

import (
	"maps"

	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

/*
 * This enumeration has a composite key.
 * It works differently from most of the other enumerations.
 */

// All RegionCompositeKey specifications including depreciated and import only.
func RegionCompositeKeyListAll() []Spec {
	return append([]Spec(nil), internalRegionCompositeKeyListAll...)
}

// All RegionCompositeKey specifications values that are NOT marked import-only.
func RegionCompositeKeyListCurrent() []Spec {
	return append([]Spec(nil), internalRegionCompositeKeyListCurrent...)
}

// A map of all RegionCompositeKey specifications.
func RegionCompositeKeyMap() map[RegionCompositeKey]Spec {
	cp := make(map[RegionCompositeKey]Spec, len(internalRegionCompositeKeyMap))
	maps.Copy(cp, internalRegionCompositeKeyMap)
	return cp
}

// A map of all RegionCompositeKey specifications.
var internalRegionCompositeKeyMap = map[RegionCompositeKey]Spec{
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

// All RegionCompositeKey specifications including depreciated and import only.
var internalRegionCompositeKeyListAll = []Spec{
	internalRegionCompositeKeyMap["AI.248"],
	internalRegionCompositeKeyMap["BI.259"],
	internalRegionCompositeKeyMap["ET.390"],
	internalRegionCompositeKeyMap["IV.206"],
	internalRegionCompositeKeyMap["KO.0"],
	internalRegionCompositeKeyMap["KO.296"],
	internalRegionCompositeKeyMap["KO.522"],
	internalRegionCompositeKeyMap["NONE"],
	internalRegionCompositeKeyMap["SI.279"],
	internalRegionCompositeKeyMap["SY.248"],
}

// All RegionCompositeKey specifications values that are NOT marked import-only.
var internalRegionCompositeKeyListCurrent = []Spec{
	internalRegionCompositeKeyMap["AI.248"],
	internalRegionCompositeKeyMap["BI.259"],
	internalRegionCompositeKeyMap["ET.390"],
	internalRegionCompositeKeyMap["IV.206"],
	internalRegionCompositeKeyMap["KO.0"],
	internalRegionCompositeKeyMap["KO.296"],
	internalRegionCompositeKeyMap["KO.522"],
	internalRegionCompositeKeyMap["NONE"],
	internalRegionCompositeKeyMap["SI.279"],
	internalRegionCompositeKeyMap["SY.248"],
}
