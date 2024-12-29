package main

import (
	"fmt"

	"github.com/hamradiolog-net/spec"
)

func main() {
	validateDxcc()
}

func validateDxcc() {
	compareARRLSectionDXCCtoDXCCList()
}

func compareARRLSectionDXCCtoDXCCList() {
	for _, section := range spec.EnumARRLSectionList {
		for _, code := range section.DXCCEntityCodes.Code {

			if _, ok := spec.EnumDXCCEntityCodeMap[code]; !ok {
				fmt.Printf("ARRL Section \"%s\" has DXCC code %d which is not in the DXCC list.\n", section.ID, code)
			}

		}
	}
}
