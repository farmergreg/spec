package qslrcvd

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// QSLRcvdSpec represents the specification for a single QSLRcvd
type QSLRcvdSpec struct {
	EnumerationName string                   `json:"Enumeration Name"`
	IsImportOnly    spectype.AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string                   `json:"Comments,omitempty"`
	Id              string                   `json:"Status"` // Status
	Meaning         string                   `json:"Meaning"`
	Description     string                   `json:"Description"`
}

// SpecMap contains all QSLRcvdSpec specifications.
type SpecMap struct {
	Header  []string                `json:"Header"`
	Records map[QSLRcvd]QSLRcvdSpec `json:"Records"`
}
