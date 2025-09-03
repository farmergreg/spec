package dxccentitycode

import (
	"encoding/json"
	"strconv"
	"strings"
)

// DXCCEntityCodeList represents a list of DXCC entity codes
type DXCCEntityCodeList []DXCCEntityCode

func (d DXCCEntityCodeList) String() string {
	var strCodes []string
	for _, code := range d {
		strCodes = append(strCodes, code.String())
	}
	return strings.Join(strCodes, ", ")
}

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
