package spectype

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// Boolean is a boolean stored in the ADIF JSON specification.
type Boolean bool

func (a *Boolean) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	*a = strings.ToLower(strings.TrimSpace(val)) == "true"
	return nil
}

// Integer is an integer stored in the ADIF JSON specification.
type Integer int

func (a *Integer) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	result, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*a = Integer(result)

	return nil
}

// DateTime is a date/time stored in the ADIF JSON specification.
type DateTime time.Time

func (a *DateTime) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	parsedDate, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return err
	}
	*a = DateTime(parsedDate)

	return nil
}

// DateOnly is a date stored in the ADIF JSON specification.
type DateOnly time.Time

func (a *DateOnly) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	parsedDate, err := time.Parse("2006-01-02", val)
	if err != nil {
		return err
	}
	*a = DateOnly(parsedDate)

	return nil
}
