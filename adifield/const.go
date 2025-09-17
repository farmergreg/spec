package adifield

const (
	// Header: End Of Header Marker. No data, no length.
	EOH Field = "eoh"

	// Record: End of Record Marker. No data, no length.
	EOR Field = "eor"

	// USERDEF is the prefix for user defined fields.
	// These fields must always have a number after them, e.g. USERDEF1, USERDEF2, etc.
	USERDEF = "userdef"

	// app_ is the prefix for application defined fields.
	// These field follow the pattern: app_programid_fieldname
	// For example, app_k9cts_test_field
	APP_ = "app_"
)
