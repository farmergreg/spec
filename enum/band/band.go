package band

import (
	"encoding/json"
	"strconv"
)

// Band represents a ham radio band
type Band string

// MHz is Megahertz
type MHz float64

func (a *MHz) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	mhz, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}

	*a = MHz(mhz)
	return nil
}

func (b Band) String() string {
	return string(b)
}
