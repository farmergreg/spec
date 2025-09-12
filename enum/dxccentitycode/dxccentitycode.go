package dxccentitycode

import (
	"encoding/json"
	"strconv"

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
)

// DXCCEntityCode is an ARRL DX Century Club code.
// See Also: https://www.arrl.org/dxcc
type DXCCEntityCode int

var _ codegen.CodeGenKey = DXCCEntityCode(0)

// Int returns the integer representation of the DXCCEntityCode.
func (c DXCCEntityCode) Int() int {
	return int(c)
}

func (t DXCCEntityCode) Compare(other DXCCEntityCode) int {
	return int(t) - int(other)
}

// String returns the string representation of the DXCCEntityCode.
func (c DXCCEntityCode) String() string {
	return strconv.Itoa(int(c))
}

func (c *DXCCEntityCode) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	parsedCode, err := strconv.Atoi(val)
	if err != nil {
		return err
	}

	*c = DXCCEntityCode(parsedCode)
	return nil
}
