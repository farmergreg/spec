// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package band provides code and constants as defined in ADIF 3.1.6 (Proposed)
package band

import "sync"

const (
	BAND_1_25CM Band = "1.25CM" // 1.25CM =   24000.0000 MHz to   24250.0000 MHz
	BAND_1_25M  Band = "1.25M"  // 1.25M  =     222.0000 MHz to     225.0000 MHz
	BAND_10M    Band = "10M"    // 10M    =      28.0000 MHz to      29.7000 MHz
	BAND_12M    Band = "12M"    // 12M    =      24.8900 MHz to      24.9900 MHz
	BAND_13CM   Band = "13CM"   // 13CM   =    2300.0000 MHz to    2450.0000 MHz
	BAND_15M    Band = "15M"    // 15M    =      21.0000 MHz to      21.4500 MHz
	BAND_160M   Band = "160M"   // 160M   =       1.8000 MHz to       2.0000 MHz
	BAND_17M    Band = "17M"    // 17M    =      18.0680 MHz to      18.1680 MHz
	BAND_1MM    Band = "1MM"    // 1MM    =  241000.0000 MHz to  250000.0000 MHz
	BAND_2_5MM  Band = "2.5MM"  // 2.5MM  =  119980.0000 MHz to  123000.0000 MHz
	BAND_20M    Band = "20M"    // 20M    =      14.0000 MHz to      14.3500 MHz
	BAND_2190M  Band = "2190M"  // 2190M  =       0.1357 MHz to       0.1378 MHz
	BAND_23CM   Band = "23CM"   // 23CM   =    1240.0000 MHz to    1300.0000 MHz
	BAND_2M     Band = "2M"     // 2M     =     144.0000 MHz to     148.0000 MHz
	BAND_2MM    Band = "2MM"    // 2MM    =  134000.0000 MHz to  149000.0000 MHz
	BAND_30M    Band = "30M"    // 30M    =      10.1000 MHz to      10.1500 MHz
	BAND_33CM   Band = "33CM"   // 33CM   =     902.0000 MHz to     928.0000 MHz
	BAND_3CM    Band = "3CM"    // 3CM    =   10000.0000 MHz to   10500.0000 MHz
	BAND_40M    Band = "40M"    // 40M    =       7.0000 MHz to       7.3000 MHz
	BAND_4M     Band = "4M"     // 4M     =      70.0000 MHz to      71.0000 MHz
	BAND_4MM    Band = "4MM"    // 4MM    =   75500.0000 MHz to   81000.0000 MHz
	BAND_560M   Band = "560M"   // 560M   =       0.5010 MHz to       0.5040 MHz
	BAND_5M     Band = "5M"     // 5M     =      54.0000 MHz to      69.9000 MHz
	BAND_60M    Band = "60M"    // 60M    =       5.0600 MHz to       5.4500 MHz
	BAND_630M   Band = "630M"   // 630M   =       0.4720 MHz to       0.4790 MHz
	BAND_6CM    Band = "6CM"    // 6CM    =    5650.0000 MHz to    5925.0000 MHz
	BAND_6M     Band = "6M"     // 6M     =      50.0000 MHz to      54.0000 MHz
	BAND_6MM    Band = "6MM"    // 6MM    =   47000.0000 MHz to   47200.0000 MHz
	BAND_70CM   Band = "70CM"   // 70CM   =     420.0000 MHz to     450.0000 MHz
	BAND_80M    Band = "80M"    // 80M    =       3.5000 MHz to       4.0000 MHz
	BAND_8M     Band = "8M"     // 8M     =      40.0000 MHz to      45.0000 MHz
	BAND_9CM    Band = "9CM"    // 9CM    =    3300.0000 MHz to    3500.0000 MHz
	BAND_SUBMM  Band = "SUBMM"  // SUBMM  =  300000.0000 MHz to 7500000.0000 MHz
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Band specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "1.25CM", LowerFreqMHz: 24000, UpperFreqMHz: 24250},
	{IsImportOnly: false, Key: "1.25M", LowerFreqMHz: 222, UpperFreqMHz: 225},
	{IsImportOnly: false, Key: "10M", LowerFreqMHz: 28, UpperFreqMHz: 29.7},
	{IsImportOnly: false, Key: "12M", LowerFreqMHz: 24.89, UpperFreqMHz: 24.99},
	{IsImportOnly: false, Key: "13CM", LowerFreqMHz: 2300, UpperFreqMHz: 2450},
	{IsImportOnly: false, Key: "15M", LowerFreqMHz: 21, UpperFreqMHz: 21.45},
	{IsImportOnly: false, Key: "160M", LowerFreqMHz: 1.8, UpperFreqMHz: 2},
	{IsImportOnly: false, Key: "17M", LowerFreqMHz: 18.068, UpperFreqMHz: 18.168},
	{IsImportOnly: false, Key: "1MM", LowerFreqMHz: 241000, UpperFreqMHz: 250000},
	{IsImportOnly: false, Key: "2.5MM", LowerFreqMHz: 119980, UpperFreqMHz: 123000},
	{IsImportOnly: false, Key: "20M", LowerFreqMHz: 14, UpperFreqMHz: 14.35},
	{IsImportOnly: false, Key: "2190M", LowerFreqMHz: 0.1357, UpperFreqMHz: 0.1378},
	{IsImportOnly: false, Key: "23CM", LowerFreqMHz: 1240, UpperFreqMHz: 1300},
	{IsImportOnly: false, Key: "2M", LowerFreqMHz: 144, UpperFreqMHz: 148},
	{IsImportOnly: false, Key: "2MM", LowerFreqMHz: 134000, UpperFreqMHz: 149000},
	{IsImportOnly: false, Key: "30M", LowerFreqMHz: 10.1, UpperFreqMHz: 10.15},
	{IsImportOnly: false, Key: "33CM", LowerFreqMHz: 902, UpperFreqMHz: 928},
	{IsImportOnly: false, Key: "3CM", LowerFreqMHz: 10000, UpperFreqMHz: 10500},
	{IsImportOnly: false, Key: "40M", LowerFreqMHz: 7, UpperFreqMHz: 7.3},
	{IsImportOnly: false, Key: "4M", LowerFreqMHz: 70, UpperFreqMHz: 71},
	{IsImportOnly: false, Key: "4MM", LowerFreqMHz: 75500, UpperFreqMHz: 81000},
	{IsImportOnly: false, Key: "560M", LowerFreqMHz: 0.501, UpperFreqMHz: 0.504},
	{IsImportOnly: false, Key: "5M", LowerFreqMHz: 54.000001, UpperFreqMHz: 69.9},
	{IsImportOnly: false, Key: "60M", LowerFreqMHz: 5.06, UpperFreqMHz: 5.45},
	{IsImportOnly: false, Key: "630M", LowerFreqMHz: 0.472, UpperFreqMHz: 0.479},
	{IsImportOnly: false, Key: "6CM", LowerFreqMHz: 5650, UpperFreqMHz: 5925},
	{IsImportOnly: false, Key: "6M", LowerFreqMHz: 50, UpperFreqMHz: 54},
	{IsImportOnly: false, Key: "6MM", LowerFreqMHz: 47000, UpperFreqMHz: 47200},
	{IsImportOnly: false, Key: "70CM", LowerFreqMHz: 420, UpperFreqMHz: 450},
	{IsImportOnly: false, Key: "80M", LowerFreqMHz: 3.5, UpperFreqMHz: 4},
	{IsImportOnly: false, Key: "8M", LowerFreqMHz: 40, UpperFreqMHz: 45},
	{IsImportOnly: false, Key: "9CM", LowerFreqMHz: 3300, UpperFreqMHz: 3500},
	{IsImportOnly: false, Key: "SUBMM", LowerFreqMHz: 300000, UpperFreqMHz: 7.5e+06},
}

// lookupMap contains all known Band specifications.
var lookupMap = map[Band]*Spec{
	BAND_1_25CM: &lookupList[0],
	BAND_1_25M:  &lookupList[1],
	BAND_10M:    &lookupList[2],
	BAND_12M:    &lookupList[3],
	BAND_13CM:   &lookupList[4],
	BAND_15M:    &lookupList[5],
	BAND_160M:   &lookupList[6],
	BAND_17M:    &lookupList[7],
	BAND_1MM:    &lookupList[8],
	BAND_2_5MM:  &lookupList[9],
	BAND_20M:    &lookupList[10],
	BAND_2190M:  &lookupList[11],
	BAND_23CM:   &lookupList[12],
	BAND_2M:     &lookupList[13],
	BAND_2MM:    &lookupList[14],
	BAND_30M:    &lookupList[15],
	BAND_33CM:   &lookupList[16],
	BAND_3CM:    &lookupList[17],
	BAND_40M:    &lookupList[18],
	BAND_4M:     &lookupList[19],
	BAND_4MM:    &lookupList[20],
	BAND_560M:   &lookupList[21],
	BAND_5M:     &lookupList[22],
	BAND_60M:    &lookupList[23],
	BAND_630M:   &lookupList[24],
	BAND_6CM:    &lookupList[25],
	BAND_6M:     &lookupList[26],
	BAND_6MM:    &lookupList[27],
	BAND_70CM:   &lookupList[28],
	BAND_80M:    &lookupList[29],
	BAND_8M:     &lookupList[30],
	BAND_9CM:    &lookupList[31],
	BAND_SUBMM:  &lookupList[32],
}

// Lookup returns the specification for the provided Band.
// ADIF 3.1.6
func Lookup(band Band) (Spec, bool) {
	spec, ok := lookupMap[band]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all Band specifications that match the provided filter function.
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

// List returns all Band specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns Band specifications.
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
