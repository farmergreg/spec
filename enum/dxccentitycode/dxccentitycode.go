package dxccentitycode

import (
	"encoding/json"
	"strconv"
	"strings"
)

// DXCCEntityCode is an ARRL DX Century Club code.
// See Also: https://www.arrl.org/dxcc
type DXCCEntityCode int

func (d DXCCEntityCode) ToInt() int {
	return int(d)
}

func (d *DXCCEntityCode) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	parsedCode, err := strconv.Atoi(val)
	if err != nil {
		return err
	}

	*d = DXCCEntityCode(parsedCode)
	return nil
}

// DXCCEntityCodeList represents a list of DXCC entity codes
type DXCCEntityCodeList []DXCCEntityCode

func (d *DXCCEntityCodeList) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	codes := strings.SplitSeq(val, ",")
	for c := range codes {
		parsedCode, err := strconv.Atoi(c)
		if err != nil {
			return err
		}
		*d = append(*d, DXCCEntityCode(parsedCode))
	}

	return nil
}

func (d DXCCEntityCode) String() string {
	return strconv.Itoa(int(d))
}
