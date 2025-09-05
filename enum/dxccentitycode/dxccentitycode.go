package dxccentitycode

import (
	"encoding/json"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/internal/codegen"
)

// DXCCEntityCode is an ARRL DX Century Club code.
// See Also: https://www.arrl.org/dxcc
type DXCCEntityCode int

var _ codegen.CodeGeneratorEnumValue = DXCCEntityCode(0)

func (d DXCCEntityCode) ToInt() int {
	return int(d)
}

func (d DXCCEntityCode) String() string {
	return strconv.Itoa(int(d))
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
