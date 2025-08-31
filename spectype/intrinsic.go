package spectype

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// Boolean is a boolean stored in the ADIF JSON specification.
type Boolean bool

func (b *Boolean) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	*b = strings.ToLower(strings.TrimSpace(val)) == "true"
	return nil
}

func (b Boolean) ToBool() bool {
	return bool(b)
}

// Integer is an integer stored in the ADIF JSON specification.
type Integer int

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

func (i Integer) ToInt() int {
	return int(i)
}

// DateTime is a date/time stored in the ADIF JSON specification.
type DateTime int64

// ToTime converts DateTime to a standard time.Time representation
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
