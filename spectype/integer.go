package spectype

import (
	"encoding/json"
	"strconv"
)

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
