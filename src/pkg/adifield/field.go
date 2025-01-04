package adifield

import "strings"

// IsValid returns true if the field NAME is valid per the ADIF specification
// n.b. it expects the field name to be in UPPERCASE (match the constants in this package)
func (f Field) IsValid() bool {
	return FieldMap[f] != nil || strings.HasPrefix(string(f), APP_) || strings.HasPrefix(string(f), USERDEF)
}
