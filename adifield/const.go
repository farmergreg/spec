package adifield

const (
	// Header: End Of Header Marker. No data, no length.
	EOH ADIField = "EOH"

	// Record: End of Record Marker. No data, no length.
	EOR ADIField = "EOR"

	// USERDEF is the prefix for user defined fields.
	// These fields must always have a number after them, e.g. USERDEF1, USERDEF2, etc.
	USERDEF = "USERDEF"

	// APP_ is the prefix for application defined fields.
	// These field follow the pattern: APP_PROGRAMID_FIELDNAME
	// For example, APP_K9CTS_TEST_FIELD
	APP_ = "APP_"
)
