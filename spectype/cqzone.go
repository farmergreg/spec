package spectype

import (
	"encoding/json"
	"strconv"
	"strings"
)

type CQZone int
type CQZoneList []CQZone

func (d CQZone) ToInt() int {
	return int(d)
}

func (d *CQZoneList) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	// TODO Hack... does the ADIF Spec need a fix?
	if val == "S=16 T=17" {
		val = "16,17"
	}

	codes := strings.SplitSeq(val, ",")
	for c := range codes {
		parsedCode, err := strconv.Atoi(c)
		if err != nil {
			return err
		}
		*d = append(*d, CQZone(parsedCode))
	}

	return nil
}
