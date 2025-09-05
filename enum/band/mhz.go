package band

import (
	"encoding/json"
	"strconv"
)

// MHz is Megahertz
type MHz float64

func (a MHz) ToFloat64() float64 {
	return float64(a)
}

func (a MHz) String() string {
	return strconv.FormatFloat(float64(a), 'f', -1, 64)
}

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
