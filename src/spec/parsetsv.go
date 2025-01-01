package spec

import (
	"bytes"
	"encoding/csv"
	"slices"

	"github.com/gocarina/gocsv"
)

// ParseADISpecTSV builds a map from a TSV files
func ParseADISpecTSV[T any](tsvData []byte) ([]T, error) {
	var err error
	r := csv.NewReader(bytes.NewReader(tsvData))
	r.Comma = '\t'
	r.LazyQuotes = true

	var parsedTSV []T
	err = gocsv.UnmarshalCSV(r, &parsedTSV)
	parsedTSV = slices.Clip(parsedTSV)

	return parsedTSV, err
}
