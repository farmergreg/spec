package mode

import (
	"encoding/json"
	"strings"

	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/submode"
)

// Mode is the ADIF mode of a radio communication.
type Mode string

// EnumModeSubmodeList is a list of submodes as defined by the Modes in the ADIF specification export.
type EnumModeSubModeList struct {
	Submodes []submode.SubMode
}

func (s *EnumModeSubModeList) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	sub := strings.Split(val, ",")

	if len(sub) == 1 && sub[0] == "" {
		sub = []string{}
	}

	s.Submodes = make([]submode.SubMode, len(sub))
	for i, v := range sub {
		s.Submodes[i] = submode.SubMode(v)
	}
	return nil
}
