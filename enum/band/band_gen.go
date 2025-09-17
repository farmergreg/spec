// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package band provides code and constants as defined in ADIF 3.1.6 (Released)
package band

import "sync"

const (
	BAND_1_25CM Band = "1.25cm" // 1.25cm =   24000.0000 MHz to   24250.0000 MHz
	BAND_1_25M  Band = "1.25m"  // 1.25m  =     222.0000 MHz to     225.0000 MHz
	BAND_10M    Band = "10m"    // 10m    =      28.0000 MHz to      29.7000 MHz
	BAND_12M    Band = "12m"    // 12m    =      24.8900 MHz to      24.9900 MHz
	BAND_13CM   Band = "13cm"   // 13cm   =    2300.0000 MHz to    2450.0000 MHz
	BAND_15M    Band = "15m"    // 15m    =      21.0000 MHz to      21.4500 MHz
	BAND_160M   Band = "160m"   // 160m   =       1.8000 MHz to       2.0000 MHz
	BAND_17M    Band = "17m"    // 17m    =      18.0680 MHz to      18.1680 MHz
	BAND_1MM    Band = "1mm"    // 1mm    =  241000.0000 MHz to  250000.0000 MHz
	BAND_2_5MM  Band = "2.5mm"  // 2.5mm  =  119980.0000 MHz to  123000.0000 MHz
	BAND_20M    Band = "20m"    // 20m    =      14.0000 MHz to      14.3500 MHz
	BAND_2190M  Band = "2190m"  // 2190m  =       0.1357 MHz to       0.1378 MHz
	BAND_23CM   Band = "23cm"   // 23cm   =    1240.0000 MHz to    1300.0000 MHz
	BAND_2M     Band = "2m"     // 2m     =     144.0000 MHz to     148.0000 MHz
	BAND_2MM    Band = "2mm"    // 2mm    =  134000.0000 MHz to  149000.0000 MHz
	BAND_30M    Band = "30m"    // 30m    =      10.1000 MHz to      10.1500 MHz
	BAND_33CM   Band = "33cm"   // 33cm   =     902.0000 MHz to     928.0000 MHz
	BAND_3CM    Band = "3cm"    // 3cm    =   10000.0000 MHz to   10500.0000 MHz
	BAND_40M    Band = "40m"    // 40m    =       7.0000 MHz to       7.3000 MHz
	BAND_4M     Band = "4m"     // 4m     =      70.0000 MHz to      71.0000 MHz
	BAND_4MM    Band = "4mm"    // 4mm    =   75500.0000 MHz to   81000.0000 MHz
	BAND_560M   Band = "560m"   // 560m   =       0.5010 MHz to       0.5040 MHz
	BAND_5M     Band = "5m"     // 5m     =      54.0000 MHz to      69.9000 MHz
	BAND_60M    Band = "60m"    // 60m    =       5.0600 MHz to       5.4500 MHz
	BAND_630M   Band = "630m"   // 630m   =       0.4720 MHz to       0.4790 MHz
	BAND_6CM    Band = "6cm"    // 6cm    =    5650.0000 MHz to    5925.0000 MHz
	BAND_6M     Band = "6m"     // 6m     =      50.0000 MHz to      54.0000 MHz
	BAND_6MM    Band = "6mm"    // 6mm    =   47000.0000 MHz to   47200.0000 MHz
	BAND_70CM   Band = "70cm"   // 70cm   =     420.0000 MHz to     450.0000 MHz
	BAND_80M    Band = "80m"    // 80m    =       3.5000 MHz to       4.0000 MHz
	BAND_8M     Band = "8m"     // 8m     =      40.0000 MHz to      45.0000 MHz
	BAND_9CM    Band = "9cm"    // 9cm    =    3300.0000 MHz to    3500.0000 MHz
	BAND_SUBMM  Band = "submm"  // submm  =  300000.0000 MHz to 7500000.0000 MHz
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Band specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "1.25cm", LowerFreqMHz: 24000, UpperFreqMHz: 24250},
	{IsImportOnly: false, Key: "1.25m", LowerFreqMHz: 222, UpperFreqMHz: 225},
	{IsImportOnly: false, Key: "10m", LowerFreqMHz: 28, UpperFreqMHz: 29.7},
	{IsImportOnly: false, Key: "12m", LowerFreqMHz: 24.89, UpperFreqMHz: 24.99},
	{IsImportOnly: false, Key: "13cm", LowerFreqMHz: 2300, UpperFreqMHz: 2450},
	{IsImportOnly: false, Key: "15m", LowerFreqMHz: 21, UpperFreqMHz: 21.45},
	{IsImportOnly: false, Key: "160m", LowerFreqMHz: 1.8, UpperFreqMHz: 2},
	{IsImportOnly: false, Key: "17m", LowerFreqMHz: 18.068, UpperFreqMHz: 18.168},
	{IsImportOnly: false, Key: "1mm", LowerFreqMHz: 241000, UpperFreqMHz: 250000},
	{IsImportOnly: false, Key: "2.5mm", LowerFreqMHz: 119980, UpperFreqMHz: 123000},
	{IsImportOnly: false, Key: "20m", LowerFreqMHz: 14, UpperFreqMHz: 14.35},
	{IsImportOnly: false, Key: "2190m", LowerFreqMHz: 0.1357, UpperFreqMHz: 0.1378},
	{IsImportOnly: false, Key: "23cm", LowerFreqMHz: 1240, UpperFreqMHz: 1300},
	{IsImportOnly: false, Key: "2m", LowerFreqMHz: 144, UpperFreqMHz: 148},
	{IsImportOnly: false, Key: "2mm", LowerFreqMHz: 134000, UpperFreqMHz: 149000},
	{IsImportOnly: false, Key: "30m", LowerFreqMHz: 10.1, UpperFreqMHz: 10.15},
	{IsImportOnly: false, Key: "33cm", LowerFreqMHz: 902, UpperFreqMHz: 928},
	{IsImportOnly: false, Key: "3cm", LowerFreqMHz: 10000, UpperFreqMHz: 10500},
	{IsImportOnly: false, Key: "40m", LowerFreqMHz: 7, UpperFreqMHz: 7.3},
	{IsImportOnly: false, Key: "4m", LowerFreqMHz: 70, UpperFreqMHz: 71},
	{IsImportOnly: false, Key: "4mm", LowerFreqMHz: 75500, UpperFreqMHz: 81000},
	{IsImportOnly: false, Key: "560m", LowerFreqMHz: 0.501, UpperFreqMHz: 0.504},
	{IsImportOnly: false, Key: "5m", LowerFreqMHz: 54.000001, UpperFreqMHz: 69.9},
	{IsImportOnly: false, Key: "60m", LowerFreqMHz: 5.06, UpperFreqMHz: 5.45},
	{IsImportOnly: false, Key: "630m", LowerFreqMHz: 0.472, UpperFreqMHz: 0.479},
	{IsImportOnly: false, Key: "6cm", LowerFreqMHz: 5650, UpperFreqMHz: 5925},
	{IsImportOnly: false, Key: "6m", LowerFreqMHz: 50, UpperFreqMHz: 54},
	{IsImportOnly: false, Key: "6mm", LowerFreqMHz: 47000, UpperFreqMHz: 47200},
	{IsImportOnly: false, Key: "70cm", LowerFreqMHz: 420, UpperFreqMHz: 450},
	{IsImportOnly: false, Key: "80m", LowerFreqMHz: 3.5, UpperFreqMHz: 4},
	{IsImportOnly: false, Key: "8m", LowerFreqMHz: 40, UpperFreqMHz: 45},
	{IsImportOnly: false, Key: "9cm", LowerFreqMHz: 3300, UpperFreqMHz: 3500},
	{IsImportOnly: false, Key: "submm", LowerFreqMHz: 300000, UpperFreqMHz: 7.5e+06},
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
func Lookup(b Band) (Spec, bool) {
	spec, ok := lookupMap[b]
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
