// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package band

var (
	Band1_25cm Band = "1.25cm" // 1.25cm   24000.0000 MHz -   24250.0000 MHz
	Band1_25m  Band = "1.25m"  // 1.25m      222.0000 MHz -     225.0000 MHz
	Band10m    Band = "10m"    // 10m         28.0000 MHz -      29.7000 MHz
	Band12m    Band = "12m"    // 12m         24.8900 MHz -      24.9900 MHz
	Band13cm   Band = "13cm"   // 13cm      2300.0000 MHz -    2450.0000 MHz
	Band15m    Band = "15m"    // 15m         21.0000 MHz -      21.4500 MHz
	Band160m   Band = "160m"   // 160m         1.8000 MHz -       2.0000 MHz
	Band17m    Band = "17m"    // 17m         18.0680 MHz -      18.1680 MHz
	Band1mm    Band = "1mm"    // 1mm     241000.0000 MHz -  250000.0000 MHz
	Band2_5mm  Band = "2.5mm"  // 2.5mm   119980.0000 MHz -  123000.0000 MHz
	Band20m    Band = "20m"    // 20m         14.0000 MHz -      14.3500 MHz
	Band2190m  Band = "2190m"  // 2190m        0.1357 MHz -       0.1378 MHz
	Band23cm   Band = "23cm"   // 23cm      1240.0000 MHz -    1300.0000 MHz
	Band2m     Band = "2m"     // 2m         144.0000 MHz -     148.0000 MHz
	Band2mm    Band = "2mm"    // 2mm     134000.0000 MHz -  149000.0000 MHz
	Band30m    Band = "30m"    // 30m         10.1000 MHz -      10.1500 MHz
	Band33cm   Band = "33cm"   // 33cm       902.0000 MHz -     928.0000 MHz
	Band3cm    Band = "3cm"    // 3cm      10000.0000 MHz -   10500.0000 MHz
	Band40m    Band = "40m"    // 40m          7.0000 MHz -       7.3000 MHz
	Band4m     Band = "4m"     // 4m          70.0000 MHz -      71.0000 MHz
	Band4mm    Band = "4mm"    // 4mm      75500.0000 MHz -   81000.0000 MHz
	Band560m   Band = "560m"   // 560m         0.5010 MHz -       0.5040 MHz
	Band5m     Band = "5m"     // 5m          54.0000 MHz -      69.9000 MHz
	Band60m    Band = "60m"    // 60m          5.0600 MHz -       5.4500 MHz
	Band630m   Band = "630m"   // 630m         0.4720 MHz -       0.4790 MHz
	Band6cm    Band = "6cm"    // 6cm       5650.0000 MHz -    5925.0000 MHz
	Band6m     Band = "6m"     // 6m          50.0000 MHz -      54.0000 MHz
	Band6mm    Band = "6mm"    // 6mm      47000.0000 MHz -   47200.0000 MHz
	Band70cm   Band = "70cm"   // 70cm       420.0000 MHz -     450.0000 MHz
	Band80m    Band = "80m"    // 80m          3.5000 MHz -       4.0000 MHz
	Band8m     Band = "8m"     // 8m          40.0000 MHz -      45.0000 MHz
	Band9cm    Band = "9cm"    // 9cm       3300.0000 MHz -    3500.0000 MHz
	Bandsubmm  Band = "submm"  // submm   300000.0000 MHz - 7500000.0000 MHz
)

/*
var BandMap = map[band.Band]spec.BandSpec{"1.25cm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"1.25cm", LowerFreqMHz:24000, UpperFreqMHz:24250}, "1.25m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"1.25m", LowerFreqMHz:222, UpperFreqMHz:225}, "10m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"10m", LowerFreqMHz:28, UpperFreqMHz:29.7}, "12m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"12m", LowerFreqMHz:24.89, UpperFreqMHz:24.99}, "13cm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"13cm", LowerFreqMHz:2300, UpperFreqMHz:2450}, "15m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"15m", LowerFreqMHz:21, UpperFreqMHz:21.45}, "160m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"160m", LowerFreqMHz:1.8, UpperFreqMHz:2}, "17m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"17m", LowerFreqMHz:18.068, UpperFreqMHz:18.168}, "1mm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"1mm", LowerFreqMHz:241000, UpperFreqMHz:250000}, "2.5mm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"2.5mm", LowerFreqMHz:119980, UpperFreqMHz:123000}, "20m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"20m", LowerFreqMHz:14, UpperFreqMHz:14.35}, "2190m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"2190m", LowerFreqMHz:0.1357, UpperFreqMHz:0.1378}, "23cm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"23cm", LowerFreqMHz:1240, UpperFreqMHz:1300}, "2m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"2m", LowerFreqMHz:144, UpperFreqMHz:148}, "2mm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"2mm", LowerFreqMHz:134000, UpperFreqMHz:149000}, "30m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"30m", LowerFreqMHz:10.1, UpperFreqMHz:10.15}, "33cm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"33cm", LowerFreqMHz:902, UpperFreqMHz:928}, "3cm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"3cm", LowerFreqMHz:10000, UpperFreqMHz:10500}, "40m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"40m", LowerFreqMHz:7, UpperFreqMHz:7.3}, "4m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"4m", LowerFreqMHz:70, UpperFreqMHz:71}, "4mm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"4mm", LowerFreqMHz:75500, UpperFreqMHz:81000}, "560m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"560m", LowerFreqMHz:0.501, UpperFreqMHz:0.504}, "5m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"5m", LowerFreqMHz:54.000001, UpperFreqMHz:69.9}, "60m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"60m", LowerFreqMHz:5.06, UpperFreqMHz:5.45}, "630m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"630m", LowerFreqMHz:0.472, UpperFreqMHz:0.479}, "6cm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"6cm", LowerFreqMHz:5650, UpperFreqMHz:5925}, "6m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"6m", LowerFreqMHz:50, UpperFreqMHz:54}, "6mm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"6mm", LowerFreqMHz:47000, UpperFreqMHz:47200}, "70cm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"70cm", LowerFreqMHz:420, UpperFreqMHz:450}, "80m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"80m", LowerFreqMHz:3.5, UpperFreqMHz:4}, "8m":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"8m", LowerFreqMHz:40, UpperFreqMHz:45}, "9cm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"9cm", LowerFreqMHz:3300, UpperFreqMHz:3500}, "submm":spec.BandSpec{EnumerationName:"Band", IsImportOnly:false, Comments:"", Id:"submm", LowerFreqMHz:300000, UpperFreqMHz:7.5e+06}}
*/
