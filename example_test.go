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
	forty, _ := band.Lookup(band.Band_40m)
	fmt.Printf("The 40m band is between %f and %f MHz\n", forty.LowerFreqMHz, forty.UpperFreqMHz)

	fmt.Println("Current Bands")
	for _, band := range band.ListActive() {
		fmt.Printf("%s: %f - %f\n", band.Key, band.LowerFreqMHz, band.UpperFreqMHz)
	}

	// If for some reason, you have an ADIF enum value as a string (e.g. post from a web form), you can convert it to the corresponding type:
	twenty := band.New("20m")
	if twenty == band.Band_20m {
		fmt.Println("The twenty variable is equal to the Band_20m constant")

		// You can also compare two ADIF enum values using the Compare method:
		twenty.Compare(band.Band_20m) // return 0 because they are equal
	}

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
	// 1.25cm: 24000.000000 - 24250.000000
	// 1.25m: 222.000000 - 225.000000
	// 10m: 28.000000 - 29.700000
	// 12m: 24.890000 - 24.990000
	// 13cm: 2300.000000 - 2450.000000
	// 15m: 21.000000 - 21.450000
	// 160m: 1.800000 - 2.000000
	// 17m: 18.068000 - 18.168000
	// 1mm: 241000.000000 - 250000.000000
	// 2.5mm: 119980.000000 - 123000.000000
	// 20m: 14.000000 - 14.350000
	// 2190m: 0.135700 - 0.137800
	// 23cm: 1240.000000 - 1300.000000
	// 2m: 144.000000 - 148.000000
	// 2mm: 134000.000000 - 149000.000000
	// 30m: 10.100000 - 10.150000
	// 33cm: 902.000000 - 928.000000
	// 3cm: 10000.000000 - 10500.000000
	// 40m: 7.000000 - 7.300000
	// 4m: 70.000000 - 71.000000
	// 4mm: 75500.000000 - 81000.000000
	// 560m: 0.501000 - 0.504000
	// 5m: 54.000001 - 69.900000
	// 60m: 5.060000 - 5.450000
	// 630m: 0.472000 - 0.479000
	// 6cm: 5650.000000 - 5925.000000
	// 6m: 50.000000 - 54.000000
	// 6mm: 47000.000000 - 47200.000000
	// 70cm: 420.000000 - 450.000000
	// 80m: 3.500000 - 4.000000
	// 8m: 40.000000 - 45.000000
	// 9cm: 3300.000000 - 3500.000000
	// submm: 300000.000000 - 7500000.000000
	// The complete spec data matches the ADIF version constant
}
