package main

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/arrlsection"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
)

func main() {
	validateDxcc()
}

func validateDxcc() {
	compareARRLSectionDXCCtoDXCCList()
}

func compareARRLSectionDXCCtoDXCCList() {
	for _, section := range arrlsection.EnumARRLSectionList {
		for _, code := range section.DXCCEntityCodes.Code {

			if _, ok := dxccentitycode.EnumDXCCEntityCodeMap[code]; !ok {
				fmt.Printf("ARRL Section \"%s\" has DXCC code %d which is not in the DXCC list.\n", section.ID, code)
			}

		}
	}
}
