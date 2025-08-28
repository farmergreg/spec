package spec

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// AdifSpecBoolean is a boolean stored in the ADIF JSON specification.
type AdifSpecBoolean bool

func (a *AdifSpecBoolean) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	*a = strings.ToLower(strings.TrimSpace(val)) == "true"
	return nil
}

// AdifSpecInteger is an integer stored in the ADIF JSON specification.
type AdifSpecInteger int

func (a *AdifSpecInteger) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	result, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*a = AdifSpecInteger(result)

	return nil
}

// AdifSpecDateTime is a date/time stored in the ADIF JSON specification.
type AdifSpecDateTime time.Time

func (a *AdifSpecDateTime) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	parsedDate, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return err
	}
	*a = AdifSpecDateTime(parsedDate)

	return nil
}

// AdifSpecDateOnly is a date stored in the ADIF JSON specification.
type AdifSpecDateOnly time.Time

func (a *AdifSpecDateOnly) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	parsedDate, err := time.Parse("2006-01-02", val)
	if err != nil {
		return err
	}
	*a = AdifSpecDateOnly(parsedDate)

	return nil
}
