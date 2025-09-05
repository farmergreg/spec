package spectype

import (
	"encoding/json"
	"strings"
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
