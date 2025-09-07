// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package qslrcvd provides code and constants as defined in ADIF 3.1.6 (Proposed)
package qslrcvd

import "sync"

const (
	I QSLRcvd = "I" // I =
	N QSLRcvd = "N" // N = an incoming QSL card has not been received the QSO has not been confirmed by the online service
	R QSLRcvd = "R" // R = the logging station has requested a QSL card the logging station has requested the QSO be uploaded to the online service
	V QSLRcvd = "V" // Deprecated: V = DXCC award credit granted for the QSL card - instead use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card DXCC credit granted for the LoTW confirmation - instead use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw
	Y QSLRcvd = "Y" // Y = an incoming QSL card has been received the QSO has been confirmed by the online service
)

var (
	listActive     []Spec
	listActiveOnce sync.Once
)

// lookupList contains all known QSLRcvd specifications
var lookupList = []Spec{
	{IsImportOnly: false, Key: "I", Meaning: "ignore or invalid", Description: ""},
	{IsImportOnly: false, Key: "N", Meaning: "no", Description: "an incoming QSL card has not been received the QSO has not been confirmed by the online service"},
	{IsImportOnly: false, Key: "R", Meaning: "requested", Description: "the logging station has requested a QSL card the logging station has requested the QSO be uploaded to the online service"},
	{IsImportOnly: true, Key: "V", Meaning: "verified", Description: "DXCC award credit granted for the QSL card - instead use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card DXCC credit granted for the LoTW confirmation - instead use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw"},
	{IsImportOnly: false, Key: "Y", Meaning: "yes (confirmed)", Description: "an incoming QSL card has been received the QSO has been confirmed by the online service"},
}

// lookupMap contains all known QSLRcvd specifications
var lookupMap = map[QSLRcvd]*Spec{
	I: &lookupList[0],
	N: &lookupList[1],
	R: &lookupList[2],
	V: &lookupList[3],
	Y: &lookupList[4],
}

// Lookup locates the ADIF 3.1.6 specification for the provided QSLRcvd
func Lookup(qslrcvd QSLRcvd) (Spec, bool) {
	spec, ok := lookupMap[qslrcvd]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all ADIF 3.1.6 QSLRcvd specifications that match the provided filter function.
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(lookupList))
	for _, v := range lookupList {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// ListActive returns a slice of ADIF 3.1.6 QSLRcvd specifications, but excludes those marked as import-only.
func ListActive() []Spec {
	listActiveOnce.Do(func() {
		listActive = LookupByFilter(func(spec Spec) bool { return !bool(spec.IsImportOnly) })
	})
	return listActive
}

// List returns a slice of all ADIF 3.1.6 QSLRcvd specifications.
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}
