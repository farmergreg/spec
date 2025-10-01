# ADIF 3.1.6 Specification Library for Go

[![Tests](https://github.com/farmergreg/spec/actions/workflows/test.yml/badge.svg)](https://github.com/farmergreg/spec/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/farmergreg/spec/v6)](https://goreportcard.com/report/github.com/farmergreg/spec/v6)
[![Go Reference](https://pkg.go.dev/badge/github.com/farmergreg/spec/v6.svg)](https://pkg.go.dev/github.com/farmergreg/spec/v6)
[![Go Version](https://img.shields.io/github/go-mod/go-version/farmergreg/spec)](https://github.com/farmergreg/spec/blob/main/go.mod)
[![License](https://img.shields.io/github/license/farmergreg/spec)](https://github.com/farmergreg/spec/blob/main/LICENSE)

## Overview

This repository contains the ADIF specification for Go.
It is generated from the [resources](https://adif.org.uk/316/ADIF_316_resources_2025_09_15.zip) published by the [ADIF Workgroup](https://www.adif.org/)

## Installing The Library

`go get github.com/farmergreg/spec@v6`

## Examples

Please see the [example code here](example_test.go).

## Library Maintenance

If you wish to update this library, the following steps are required:

1. Download the latest ADIF all.json file export from the ADIF Workgroup. This file must be placed into the `specdata/` directory of this repository.

2. Add code for any new enumerations to the src/pkg folder being careful to follow the existing style.

3. Update `internal/cmd/specgen` if necessary.

4. Run `go test ./...` to run the tests.

5. Run `go generate ./...` to re-generate the Go code.

## Related Projects

If you found this library useful, you may also be interested in the following projects:

- [World's Fastest ADIF Parser](https://github.com/farmergreg/adif)
- [Go ADIF Parser Benchmarks](https://github.com/farmergreg/adif-benchmark)
- [g3zod/CreateADIFTestFiles](https://github.com/g3zod/CreateADIFTestFiles) ADI Test Files
- [g3zod/CreateADIFExportFiles](https://github.com/g3zod/CreateADIFExportFiles) ADIF Workgroup Specification Export Tool

## 📝 License

This project is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.
