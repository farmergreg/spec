package submode

import (
	"encoding/json"
	"strings"
)

// EnumModeSubmodeList is a list of submodes as defined by the Modes in the ADIF specification export.
type SubModeList []SubMode

func (d SubModeList) String() string {
	var strList []string
	for _, v := range d {
		strList = append(strList, v.String())
	}
	return strings.Join(strList, ", ")
}

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
