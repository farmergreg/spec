package dxccentitycode

import (
	"encoding/json"
	"strconv"
	"strings"
)

// DXCCEntityCode is an ARRL DX Century Club code.
// See Also: https://www.arrl.org/dxcc
type DXCCEntityCode int

// DXCCEntityCodeList represents a list of DXCC entity codes
type DXCCEntityCodeList []DXCCEntityCode

func (a *DXCCEntityCodeList) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	codes := strings.Split(val, ",")
	if len(codes) == 1 && codes[0] == "" {
		codes = []string{}
	}

	for _, c := range codes {
		parsedCode, err := strconv.Atoi(c)
		if err != nil {
			return err
		}
		*a = append(*a, DXCCEntityCode(parsedCode))
	}

	return nil
}
