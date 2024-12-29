# ADIF Specification for Go

## Overview

This repository contains the ADIF specification for Go.
It is generated from the [export](https://adif.org.uk/315/ADIF_315_resources_2024_11_28.zip) published by the [ADIF Workgroup](https://www.adif.org/)

## Using The Library

Import any of the packages in the `src/pkg` directory that you wish to use.
They all follow the same pattern.
Each pacakge has constants related to the ADIF specification.
They also proved three package level variables:

- `Map` - A map of the ADIF specification for quick lookups.
- `List` - A list of values that are current.
- `ListAll` - A list of all values, including deprecated ones.

For example, to lookup information about a band, you can use the `band` package ([Run in Go Playground](https://go.dev/play/p/HJW91fhyvdJ)):

```go
package main

import (
 "fmt"

 "github.com/hamradiolog-net/adif-spec/src/pkg/band"
)

func main() {

 forty := band.EnumBandMap[band.Band40m]
 fmt.Printf("The 40m band is between %f and %f MHz\n", forty.LowerFreqMHz, forty.UpperFreqMHz)

 fmt.Println("Current Bands")
 for _, band := range band.EnumBandList {
  fmt.Printf("%s: %f - %f\n", band.ID, band.LowerFreqMHz, band.UpperFreqMHz)
 }

 fmt.Println("All Bands Including Import-Only and Unreleased (usually this is the same as EnumBandList)")
 for _, band := range band.EnumBandListAll {
  fmt.Printf("%s: %f - %f\n", band.ID, band.LowerFreqMHz, band.UpperFreqMHz)
 }
}

```

## Maintenance

The following steps are required to update the specification to the latest version.

1. Download the latest ADIF TSV file exports from the ADIF Workgroup.  The current ADIF 3.1.5 spec is in the `src/spec/315` directory. You should rename the folder to the new ADIF version number.

2. Update `cmd/specgen/main.go` to use the new TSV folder.

3. Add code for any new enumerations to the src/pkg folder being careful to follow the existing style.

4. Run the `generate.sh` script to generate the Go code.

```sh
./generate.sh
```

### Testing

```sh
./test.sh
```
