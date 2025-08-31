# ADIF 3.1.6 Specification Library for Go

[![Tests](https://github.com/hamradiolog-net/adif-spec/actions/workflows/test.yml/badge.svg)](https://github.com/hamradiolog-net/adif-spec/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hamradiolog-net/adif-spec)](https://goreportcard.com/report/github.com/hamradiolog-net/adif-spec)
[![Go Reference](https://pkg.go.dev/badge/github.com/hamradiolog-net/adif-spec.svg)](https://pkg.go.dev/github.com/hamradiolog-net/adif-spec)
[![Go Version](https://img.shields.io/github/go-mod/go-version/hamradiolog-net/adif-spec)](https://github.com/hamradiolog-net/adif-spec/blob/main/go.mod)
[![License](https://img.shields.io/github/license/hamradiolog-net/adif-spec)](https://github.com/hamradiolog-net/adif-spec/blob/main/LICENSE)

## Overview

This repository contains the ADIF specification for Go.
It is generated from the [export](https://adif.org.uk/315/ADIF_315_resources_2024_11_28.zip) published by the [ADIF Workgroup](https://www.adif.org/)

## Using The Library

- Run `go get github.com/hamradiolog-net/adif-spec@latest`
- Import any of the packages in the [`src/pkg`](src/pkg) directory that you wish to use.
- Each package has constants related to the ADIF specification with the following suffixes:
  - *.`Map` Suffix - A map of the ADIF specification for quick lookups.
  - *.`List` Suffix - A list of values that are current.
  - *.`ListAll` Suffix - A list of all values, including deprecated ones.

For example, to lookup information about a band, we can use the `band` package ([Run in Go Playground](https://go.dev/play/p/6RkeK-RGayh)):

```go
package main

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/enum/band"
)

func main() {

	forty := band.BandMap[band.Band40m]
	fmt.Printf("The 40m band is between %f and %f MHz\n", forty.LowerFreqMHz, forty.UpperFreqMHz)

	fmt.Println("Current Bands")
	for _, band := range band.BandListCurrent {
		fmt.Printf("%s: %f - %f\n", band.Key, band.LowerFreqMHz, band.UpperFreqMHz)
	}

	fmt.Println("All Bands Including Import-Only and Unreleased (usually this is the same as BandListCurrent)")
	for _, band := range band.BandListAll {
		fmt.Printf("%s: %f - %f\n", band.Key, band.LowerFreqMHz, band.UpperFreqMHz)
	}
}
```

## Maintenance

The following steps are required to update the specification to the latest version.

1. Download the latest ADIF all.json file export from the ADIF Workgroup. This file must be placed into the `src/pkg/specdata/` directory of this repository.

2. Add code for any new enumerations to the src/pkg folder being careful to follow the existing style.

3. Update `src/cmd/specgen` if necessary.

4. Run `go test ./...` to run the tests.

5. Run `go generate ./...` to re-generate the Go code.

## Related Projects

If you found this library useful, you may also be interested in the following projects:

- [World's Fastest ADIF Parser](https://github.com/hamradiolog-net/adif)
- [Go ADIF Parser Benchmarks](https://github.com/hamradiolog-net/adif-benchmark)
- [g3zod/CreateADIFTestFiles](https://github.com/g3zod/CreateADIFTestFiles) ADI Test Files
- [g3zod/CreateADIFExportFiles](https://github.com/g3zod/CreateADIFExportFiles) ADIF Workgroup Specification Export Tool

## 📝 License

This project is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.
