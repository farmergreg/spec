// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package region provides code and constants as defined in ADIF 3.1.6 (Released)
package region

import (
	"sync"

	"github.com/farmergreg/spec/v6/spectype"
)

/*
 * This enumeration has a composite key.
 * It behaves differently than most of the other enumerations.
 */

const (
	AI   RegionCode = "ai"   // ai
	BI   RegionCode = "bi"   // bi
	ET   RegionCode = "et"   // et
	IV   RegionCode = "iv"   // iv
	KO   RegionCode = "ko"   // ko
	NONE RegionCode = "none" // none
	SI   RegionCode = "si"   // si
	SY   RegionCode = "sy"   // sy
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known RegionCompositeKey specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Code: "ai", DXCCEntityCode: 248, Region: "African Italy", Prefix: "IG9", Applicability: spectype.StringSlice{"CQ"}, StartDate: 0, EndDate: 0},
	{IsImportOnly: false, Code: "bi", DXCCEntityCode: 259, Region: "Bear Island", Prefix: "JW/B", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
	{IsImportOnly: false, Code: "et", DXCCEntityCode: 390, Region: "European Turkey", Prefix: "TA1", Applicability: spectype.StringSlice{"CQ"}, StartDate: 0, EndDate: 0},
	{IsImportOnly: false, Code: "iv", DXCCEntityCode: 206, Region: "ITU Vienna", Prefix: "4U1V", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
	{IsImportOnly: false, Code: "ko", DXCCEntityCode: 0, Region: "Kosovo", Prefix: "Z6", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 1347408000, EndDate: 1516406400},
	{IsImportOnly: false, Code: "ko", DXCCEntityCode: 296, Region: "Kosovo", Prefix: "YU8", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 1347321600},
	{IsImportOnly: false, Code: "ko", DXCCEntityCode: 522, Region: "Kosovo", Prefix: "Z6", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 1516492800, EndDate: 0},
	{IsImportOnly: false, Code: "none", DXCCEntityCode: 0, Region: "Not within a WAE or CQ region that is within a DXCC entity", Prefix: "", Applicability: spectype.StringSlice(nil), StartDate: 0, EndDate: 0},
	{IsImportOnly: false, Code: "si", DXCCEntityCode: 279, Region: "Shetland Islands", Prefix: "GM/S", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
	{IsImportOnly: false, Code: "sy", DXCCEntityCode: 248, Region: "Sicily", Prefix: "IT9", Applicability: spectype.StringSlice{"CQ", "WAE"}, StartDate: 0, EndDate: 0},
}

// lookupMap contains all known RegionCompositeKey specifications.
var lookupMap = map[RegionCompositeKey]*Spec{
	"ai.248": &lookupList[0],
	"bi.259": &lookupList[1],
	"et.390": &lookupList[2],
	"iv.206": &lookupList[3],
	"ko.0":   &lookupList[4],
	"ko.296": &lookupList[5],
	"ko.522": &lookupList[6],
	"none":   &lookupList[7],
	"si.279": &lookupList[8],
	"sy.248": &lookupList[9],
}

// Lookup returns the specification for the provided RegionCompositeKey.
// ADIF 3.1.6
func Lookup(r RegionCompositeKey) (Spec, bool) {
	spec, ok := lookupMap[r]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all RegionCompositeKey specifications that match the provided filter function.
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

// List returns all RegionCompositeKey specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns RegionCompositeKey specifications.
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
