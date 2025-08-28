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

// IsValid returns true if the Band exists in the ADIF specification. This includes Import Only and Deleted values.
func (value Band) IsValid() bool {
	switch value {
	case Band1_25cm:
		return true
	case Band1_25m:
		return true
	case Band10m:
		return true
	case Band12m:
		return true
	case Band13cm:
		return true
	case Band15m:
		return true
	case Band160m:
		return true
	case Band17m:
		return true
	case Band1mm:
		return true
	case Band2_5mm:
		return true
	case Band20m:
		return true
	case Band2190m:
		return true
	case Band23cm:
		return true
	case Band2m:
		return true
	case Band2mm:
		return true
	case Band30m:
		return true
	case Band33cm:
		return true
	case Band3cm:
		return true
	case Band40m:
		return true
	case Band4m:
		return true
	case Band4mm:
		return true
	case Band560m:
		return true
	case Band5m:
		return true
	case Band60m:
		return true
	case Band630m:
		return true
	case Band6cm:
		return true
	case Band6m:
		return true
	case Band6mm:
		return true
	case Band70cm:
		return true
	case Band80m:
		return true
	case Band8m:
		return true
	case Band9cm:
		return true
	case Bandsubmm:
		return true
	}
	return false
}
