package app

import (
	"fmt"

	"github.com/hamradiolog-net/spec/v6/enum/band"
	"github.com/hamradiolog-net/spec/v6/spec"
	"github.com/hamradiolog-net/spec/v6/specdata"
)

func Example() {

	/*
	 * You can use convenience constants and variables in packages that represent each ADIF enum type.
	 * For example, the band package represents ADIF bands:
	 */
	forty, _ := band.Lookup(band.BAND_40M)
	fmt.Printf("The 40m band is between %f and %f MHz\n", forty.LowerFreqMHz, forty.UpperFreqMHz)

	fmt.Println("Current Bands")
	for _, band := range band.ListActive() {
		fmt.Printf("%s: %f - %f\n", band.Key, band.LowerFreqMHz, band.UpperFreqMHz)
	}

	// If for some reason, you have an ADIF enum value as a string, you can convert it to the corresponding type:
	twenty := band.New("20m")
	twenty.Compare(band.BAND_20M) // true

	/*
	 * You can load an immutable copy of the specification using the spec data package.
	 */
	completeSpec := specdata.LoadADIFSpecWithExtras()
	if completeSpec.Version == spec.ADIF_VER { // e.g. "3.1.6"
		fmt.Println("The complete spec data matches the ADIF version constant")
	}

	// Output:
	// The 40m band is between 7.000000 and 7.300000 MHz
	// Current Bands
	// 1.25CM: 24000.000000 - 24250.000000
	// 1.25M: 222.000000 - 225.000000
	// 10M: 28.000000 - 29.700000
	// 12M: 24.890000 - 24.990000
	// 13CM: 2300.000000 - 2450.000000
	// 15M: 21.000000 - 21.450000
	// 160M: 1.800000 - 2.000000
	// 17M: 18.068000 - 18.168000
	// 1MM: 241000.000000 - 250000.000000
	// 2.5MM: 119980.000000 - 123000.000000
	// 20M: 14.000000 - 14.350000
	// 2190M: 0.135700 - 0.137800
	// 23CM: 1240.000000 - 1300.000000
	// 2M: 144.000000 - 148.000000
	// 2MM: 134000.000000 - 149000.000000
	// 30M: 10.100000 - 10.150000
	// 33CM: 902.000000 - 928.000000
	// 3CM: 10000.000000 - 10500.000000
	// 40M: 7.000000 - 7.300000
	// 4M: 70.000000 - 71.000000
	// 4MM: 75500.000000 - 81000.000000
	// 560M: 0.501000 - 0.504000
	// 5M: 54.000001 - 69.900000
	// 60M: 5.060000 - 5.450000
	// 630M: 0.472000 - 0.479000
	// 6CM: 5650.000000 - 5925.000000
	// 6M: 50.000000 - 54.000000
	// 6MM: 47000.000000 - 47200.000000
	// 70CM: 420.000000 - 450.000000
	// 80M: 3.500000 - 4.000000
	// 8M: 40.000000 - 45.000000
	// 9CM: 3300.000000 - 3500.000000
	// SUBMM: 300000.000000 - 7500000.000000
	// The complete spec data matches the ADIF version constant
}
