package shared

import "time"

// ImportOnly is a boolean that indicates if the record is an import-only record.
type ImportOnly bool

func (a *ImportOnly) UnmarshalCSV(csv string) error {
	*a = csv == "Import-only"
	return nil
}

// HeaderField is a boolean that indicates if the field is a header field in an ADI file.
type HeaderField bool

func (a *HeaderField) UnmarshalCSV(csv string) error {
	*a = csv == "Y"
	return nil
}

// Released is a boolean that indicates if the record is released per the ADIF specification.
type Released bool

func (a *Released) UnmarshalCSV(csv string) error {
	*a = csv == "Released" // || csv == "Proposed" // Uncomment when testing with un-released ADIF versions.
	return nil
}

// Deleted is a boolean that indicates if the record is deleted.
type Deleted bool

func (a *Deleted) UnmarshalCSV(csv string) error {
	*a = csv == "Deleted"
	return nil
}

// ImportRecordRoot is the root of an import record. It contains fields that are common to all imported records.
type ImportRecordRoot struct {
	IsImportOnly ImportOnly `csv:"Import-only"`
	IsReleased   Released   `csv:"ADIF Status"`
}

// ADIDate is a date in the ADIF format. It is stored as a string and is formatted as: YYYY/MM/DD
type ADIDate struct {
	Time time.Time
}

func (d *ADIDate) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}

	parsedDate, err := time.Parse("2006/01/02", csv)
	if err != nil {
		return err
	}
	d.Time = parsedDate
	return nil
}
