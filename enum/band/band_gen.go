// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package band

var (
	Band1_25cm Band = "1.25cm" // 1.25cm =   24000.0000 MHz to   24250.0000 MHz
	Band1_25m  Band = "1.25m"  // 1.25m  =     222.0000 MHz to     225.0000 MHz
	Band10m    Band = "10m"    // 10m    =      28.0000 MHz to      29.7000 MHz
	Band12m    Band = "12m"    // 12m    =      24.8900 MHz to      24.9900 MHz
	Band13cm   Band = "13cm"   // 13cm   =    2300.0000 MHz to    2450.0000 MHz
	Band15m    Band = "15m"    // 15m    =      21.0000 MHz to      21.4500 MHz
	Band160m   Band = "160m"   // 160m   =       1.8000 MHz to       2.0000 MHz
	Band17m    Band = "17m"    // 17m    =      18.0680 MHz to      18.1680 MHz
	Band1mm    Band = "1mm"    // 1mm    =  241000.0000 MHz to  250000.0000 MHz
	Band2_5mm  Band = "2.5mm"  // 2.5mm  =  119980.0000 MHz to  123000.0000 MHz
	Band20m    Band = "20m"    // 20m    =      14.0000 MHz to      14.3500 MHz
	Band2190m  Band = "2190m"  // 2190m  =       0.1357 MHz to       0.1378 MHz
	Band23cm   Band = "23cm"   // 23cm   =    1240.0000 MHz to    1300.0000 MHz
	Band2m     Band = "2m"     // 2m     =     144.0000 MHz to     148.0000 MHz
	Band2mm    Band = "2mm"    // 2mm    =  134000.0000 MHz to  149000.0000 MHz
	Band30m    Band = "30m"    // 30m    =      10.1000 MHz to      10.1500 MHz
	Band33cm   Band = "33cm"   // 33cm   =     902.0000 MHz to     928.0000 MHz
	Band3cm    Band = "3cm"    // 3cm    =   10000.0000 MHz to   10500.0000 MHz
	Band40m    Band = "40m"    // 40m    =       7.0000 MHz to       7.3000 MHz
	Band4m     Band = "4m"     // 4m     =      70.0000 MHz to      71.0000 MHz
	Band4mm    Band = "4mm"    // 4mm    =   75500.0000 MHz to   81000.0000 MHz
	Band560m   Band = "560m"   // 560m   =       0.5010 MHz to       0.5040 MHz
	Band5m     Band = "5m"     // 5m     =      54.0000 MHz to      69.9000 MHz
	Band60m    Band = "60m"    // 60m    =       5.0600 MHz to       5.4500 MHz
	Band630m   Band = "630m"   // 630m   =       0.4720 MHz to       0.4790 MHz
	Band6cm    Band = "6cm"    // 6cm    =    5650.0000 MHz to    5925.0000 MHz
	Band6m     Band = "6m"     // 6m     =      50.0000 MHz to      54.0000 MHz
	Band6mm    Band = "6mm"    // 6mm    =   47000.0000 MHz to   47200.0000 MHz
	Band70cm   Band = "70cm"   // 70cm   =     420.0000 MHz to     450.0000 MHz
	Band80m    Band = "80m"    // 80m    =       3.5000 MHz to       4.0000 MHz
	Band8m     Band = "8m"     // 8m     =      40.0000 MHz to      45.0000 MHz
	Band9cm    Band = "9cm"    // 9cm    =    3300.0000 MHz to    3500.0000 MHz
	Bandsubmm  Band = "submm"  // submm  =  300000.0000 MHz to 7500000.0000 MHz
)

// A map of all Band specifications.
var BandMap = map[Band]Spec{
	Band1_25cm: {IsImportOnly: false, Key: "1.25cm", LowerFreqMHz: 24000, UpperFreqMHz: 24250},
	Band1_25m:  {IsImportOnly: false, Key: "1.25m", LowerFreqMHz: 222, UpperFreqMHz: 225},
	Band10m:    {IsImportOnly: false, Key: "10m", LowerFreqMHz: 28, UpperFreqMHz: 29.7},
	Band12m:    {IsImportOnly: false, Key: "12m", LowerFreqMHz: 24.89, UpperFreqMHz: 24.99},
	Band13cm:   {IsImportOnly: false, Key: "13cm", LowerFreqMHz: 2300, UpperFreqMHz: 2450},
	Band15m:    {IsImportOnly: false, Key: "15m", LowerFreqMHz: 21, UpperFreqMHz: 21.45},
	Band160m:   {IsImportOnly: false, Key: "160m", LowerFreqMHz: 1.8, UpperFreqMHz: 2},
	Band17m:    {IsImportOnly: false, Key: "17m", LowerFreqMHz: 18.068, UpperFreqMHz: 18.168},
	Band1mm:    {IsImportOnly: false, Key: "1mm", LowerFreqMHz: 241000, UpperFreqMHz: 250000},
	Band2_5mm:  {IsImportOnly: false, Key: "2.5mm", LowerFreqMHz: 119980, UpperFreqMHz: 123000},
	Band20m:    {IsImportOnly: false, Key: "20m", LowerFreqMHz: 14, UpperFreqMHz: 14.35},
	Band2190m:  {IsImportOnly: false, Key: "2190m", LowerFreqMHz: 0.1357, UpperFreqMHz: 0.1378},
	Band23cm:   {IsImportOnly: false, Key: "23cm", LowerFreqMHz: 1240, UpperFreqMHz: 1300},
	Band2m:     {IsImportOnly: false, Key: "2m", LowerFreqMHz: 144, UpperFreqMHz: 148},
	Band2mm:    {IsImportOnly: false, Key: "2mm", LowerFreqMHz: 134000, UpperFreqMHz: 149000},
	Band30m:    {IsImportOnly: false, Key: "30m", LowerFreqMHz: 10.1, UpperFreqMHz: 10.15},
	Band33cm:   {IsImportOnly: false, Key: "33cm", LowerFreqMHz: 902, UpperFreqMHz: 928},
	Band3cm:    {IsImportOnly: false, Key: "3cm", LowerFreqMHz: 10000, UpperFreqMHz: 10500},
	Band40m:    {IsImportOnly: false, Key: "40m", LowerFreqMHz: 7, UpperFreqMHz: 7.3},
	Band4m:     {IsImportOnly: false, Key: "4m", LowerFreqMHz: 70, UpperFreqMHz: 71},
	Band4mm:    {IsImportOnly: false, Key: "4mm", LowerFreqMHz: 75500, UpperFreqMHz: 81000},
	Band560m:   {IsImportOnly: false, Key: "560m", LowerFreqMHz: 0.501, UpperFreqMHz: 0.504},
	Band5m:     {IsImportOnly: false, Key: "5m", LowerFreqMHz: 54.000001, UpperFreqMHz: 69.9},
	Band60m:    {IsImportOnly: false, Key: "60m", LowerFreqMHz: 5.06, UpperFreqMHz: 5.45},
	Band630m:   {IsImportOnly: false, Key: "630m", LowerFreqMHz: 0.472, UpperFreqMHz: 0.479},
	Band6cm:    {IsImportOnly: false, Key: "6cm", LowerFreqMHz: 5650, UpperFreqMHz: 5925},
	Band6m:     {IsImportOnly: false, Key: "6m", LowerFreqMHz: 50, UpperFreqMHz: 54},
	Band6mm:    {IsImportOnly: false, Key: "6mm", LowerFreqMHz: 47000, UpperFreqMHz: 47200},
	Band70cm:   {IsImportOnly: false, Key: "70cm", LowerFreqMHz: 420, UpperFreqMHz: 450},
	Band80m:    {IsImportOnly: false, Key: "80m", LowerFreqMHz: 3.5, UpperFreqMHz: 4},
	Band8m:     {IsImportOnly: false, Key: "8m", LowerFreqMHz: 40, UpperFreqMHz: 45},
	Band9cm:    {IsImportOnly: false, Key: "9cm", LowerFreqMHz: 3300, UpperFreqMHz: 3500},
	Bandsubmm:  {IsImportOnly: false, Key: "submm", LowerFreqMHz: 300000, UpperFreqMHz: 7.5e+06},
}

// All Band specifications including depreciated and import only.
var BandListAll = []Spec{
	BandMap[Band1_25cm],
	BandMap[Band1_25m],
	BandMap[Band10m],
	BandMap[Band12m],
	BandMap[Band13cm],
	BandMap[Band15m],
	BandMap[Band160m],
	BandMap[Band17m],
	BandMap[Band1mm],
	BandMap[Band2_5mm],
	BandMap[Band20m],
	BandMap[Band2190m],
	BandMap[Band23cm],
	BandMap[Band2m],
	BandMap[Band2mm],
	BandMap[Band30m],
	BandMap[Band33cm],
	BandMap[Band3cm],
	BandMap[Band40m],
	BandMap[Band4m],
	BandMap[Band4mm],
	BandMap[Band560m],
	BandMap[Band5m],
	BandMap[Band60m],
	BandMap[Band630m],
	BandMap[Band6cm],
	BandMap[Band6m],
	BandMap[Band6mm],
	BandMap[Band70cm],
	BandMap[Band80m],
	BandMap[Band8m],
	BandMap[Band9cm],
	BandMap[Bandsubmm],
}

// All Band specifications values that are NOT marked import-only.
var BandListCurrent = []Spec{
	BandMap[Band1_25cm],
	BandMap[Band1_25m],
	BandMap[Band10m],
	BandMap[Band12m],
	BandMap[Band13cm],
	BandMap[Band15m],
	BandMap[Band160m],
	BandMap[Band17m],
	BandMap[Band1mm],
	BandMap[Band2_5mm],
	BandMap[Band20m],
	BandMap[Band2190m],
	BandMap[Band23cm],
	BandMap[Band2m],
	BandMap[Band2mm],
	BandMap[Band30m],
	BandMap[Band33cm],
	BandMap[Band3cm],
	BandMap[Band40m],
	BandMap[Band4m],
	BandMap[Band4mm],
	BandMap[Band560m],
	BandMap[Band5m],
	BandMap[Band60m],
	BandMap[Band630m],
	BandMap[Band6cm],
	BandMap[Band6m],
	BandMap[Band6mm],
	BandMap[Band70cm],
	BandMap[Band80m],
	BandMap[Band8m],
	BandMap[Band9cm],
	BandMap[Bandsubmm],
}
