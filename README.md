# ADIF 3.1.5 Specification Library for Go

[![Tests](https://github.com/hamradiolog-net/adif-spec/actions/workflows/test.yml/badge.svg)](https://github.com/hamradiolog-net/adif-spec/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hamradiolog-net/adif-spec)](https://goreportcard.com/report/github.com/hamradiolog-net/adif-spec)
[![Go Reference](https://pkg.go.dev/badge/github.com/hamradiolog-net/adif-spec.svg)](https://pkg.go.dev/github.com/hamradiolog-net/adif-spec)
[![Go Version](https://img.shields.io/github/go-mod/go-version/hamradiolog-net/adif-spec)](https://github.com/hamradiolog-net/adif-spec/blob/main/go.mod)
[![License](https://img.shields.io/github/license/hamradiolog-net/adif-spec)](https://github.com/hamradiolog-net/adif-spec/blob/main/LICENSE)

## Overview

This repository contains the ADIF specification for Go.
It is generated from the [export](https://adif.org.uk/315/ADIF_315_resources_2024_11_28.zip) published by the [ADIF Workgroup](https://www.adif.org/)

## Using The Library

- Import any of the packages in the [`src/pkg`](src/pkg) directory that you wish to use.
- Each package has constants related to the ADIF specification.
- They use init() to create three package level variables with the following suffixes:
  - *.`Map` Suffix - A map of the ADIF specification for quick lookups.
  - *.`List` Suffix - A list of values that are current.
  - *.`ListAll` Suffix - A list of all values, including deprecated ones.

For example, to lookup information about a band, we can use the `band` package ([Run in Go Playground](https://go.dev/play/p/5tPOX3wSCpK)):

```go
package main

import (
 "fmt"

 "github.com/hamradiolog-net/adif-spec/src/pkg/enum/band"
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

1. Download the latest ADIF TSV file exports from the ADIF Workgroup.  The current ADIF 3.1.5 spec is kept in the `src/spec/315` directory of this repository. You should rename the folder to match the new ADIF version number.

2. Update [`cmd/specgen/main.go`](src/cmd/specgen/main.go) and related code to use the new TSV folder.

3. Add code for any new enumerations to the src/pkg folder being careful to follow the existing style.

4. Run `go generate ./...` to re-generate the Go code.

```sh
go generate ./...
```

### Testing

```sh
go test ./...
```

## Related Projects

If you found this library useful, you may also be interested in the following projects:

- [hamradiolog-net/adif](https://github.com/hamradiolog-net/adif): High Performance ADI Library For Go
- [g3zod/CreateADIFTestFiles](https://github.com/g3zod/CreateADIFTestFiles) ADI Test Files
- [g3zod/CreateADIFExportFiles](https://github.com/g3zod/CreateADIFExportFiles) ADIF Workgroup Specification Export Tool
