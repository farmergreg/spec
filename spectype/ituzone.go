package spectype

import (
	"encoding/json"
	"strconv"
	"strings"
)

type ITUZone int
type ITUZoneList []ITUZone

func (d ITUZone) ToInt() int {
	return int(d)
}

func (d *ITUZoneList) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	// TODO Hack... does the ADIF Spec need a fix?
	val = strings.ReplaceAll(val, "/", ",")

	codes := strings.SplitSeq(val, ",")
	for c := range codes {
		parsedCode, err := strconv.Atoi(c)
		if err != nil {
			return err
		}
		*d = append(*d, ITUZone(parsedCode))
	}

	return nil
}
