// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package region provides code and constants as defined in ADIF 3.1.6 (Proposed)
package region

import "github.com/hamradiolog-net/adif-spec/v6/spectype"

/*
 * This enumeration has a composite key.
 * It behaves differently than most of the other enumerations.
 */

const (
	AI   RegionCode = "AI"   // AI
	BI   RegionCode = "BI"   // BI
	ET   RegionCode = "ET"   // ET
	IV   RegionCode = "IV"   // IV
	KO   RegionCode = "KO"   // KO
	NONE RegionCode = "NONE" // NONE
	SI   RegionCode = "SI"   // SI
	SY   RegionCode = "SY"   // SY
)

// Lookup look up a specification for RegionCompositeKey
func Lookup(regioncompositekey RegionCompositeKey) (Spec, bool) {
	spec, ok := internalMap[regioncompositekey], true
	return spec, ok
}

// All RegionCompositeKey specifications INCLUDING ones marked import only.
func AllRegionCompositeKey() []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		result = append(result, v)
	}
	return result
}

// AllActiveRegionCompositeKey specifications EXCLUDING ones marked import only.
func AllActiveRegionCompositeKey() []Spec {
	return LookupByFilter(func(s Spec) bool {
		return !bool(s.IsImportOnly)
	})
}

// LookupByFilter returns all specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(internalMap))
	for _, v := range internalMap {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

var internalMap = map[RegionCompositeKey]Spec{
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
