package spectype

import (
	"encoding/json"
	"strconv"
)

// Oblast is an integer stored in the ADIF JSON specification.
type Oblast int

func (i *Oblast) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	result, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*i = Oblast(result)

	return nil
}

func (i Oblast) ToInt() int {
	return int(i)
}
