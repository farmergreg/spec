package submode

import (
	"encoding/json"
	"strings"
)

// SubMode represents the submode of an ADIF record
type SubMode string

// EnumModeSubmodeList is a list of submodes as defined by the Modes in the ADIF specification export.
type SubModeList []SubMode

func (d *SubModeList) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	codes := strings.SplitSeq(val, ",")
	for c := range codes {
		*d = append(*d, SubMode(c))
	}

	return nil
}

func (s SubMode) String() string {
	return string(s)
}
