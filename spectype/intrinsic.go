package spectype

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// Boolean is a boolean stored in the ADIF JSON specification.
type Boolean bool

// ToBool converts Boolean to bool
func (b Boolean) ToBool() bool {
	return bool(b)
}

func (b *Boolean) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	*b = strings.ToLower(strings.TrimSpace(val)) == "true"
	return nil
}

// Integer is an integer stored in the ADIF JSON specification.
type Integer int

// ToInt converts Integer to int
func (i Integer) ToInt() int {
	return int(i)
}

func (i *Integer) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	result, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*i = Integer(result)

	return nil
}

// DateTime is a date/time stored in the ADIF JSON specification.
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
