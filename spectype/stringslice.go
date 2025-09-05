package spectype

import (
	"encoding/json"
	"strings"
)

// StringSlice is a slice of strings stored in the ADIF JSON specification as a comma-separated string.
type StringSlice []string

func (d StringSlice) String() string {
	return strings.Join(d, ", ")
}

func (d *StringSlice) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	codes := strings.Split(val, ",")
	for i := range codes {
		codes[i] = strings.TrimSpace(codes[i])
	}
	*d = append(*d, codes...)

	return nil
}
