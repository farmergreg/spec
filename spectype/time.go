package spectype

import (
	"encoding/json"
	"time"
)

// DateTime is a date/time stored in the ADIF JSON specification.
// It is represented as an Unix timestamp.
type DateTime int64

// ToTime converts DateTime to time.Time
func (d DateTime) ToTime() time.Time {
	return time.Unix(int64(d), 0)
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	parsedDate, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return err
	}
	*d = DateTime(parsedDate.Unix())

	return nil
}
