package spec

import "github.com/hamradiolog-net/adif-spec/src/pkg/enum/qslrcvd"

// QSLRcvdSpec represents the specification for a single QSLRcvd
type QSLRcvdSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
	Id              string          `json:"Status"` // Status
	Meaning         string          `json:"Meaning"`
	Description     string          `json:"Description"`
}

// QSLRcvdSpecMap contains all QSLRcvdSpec specifications.
type QSLRcvdSpecMap struct {
	Header  []string                        `json:"Header"`
	Records map[qslrcvd.QSLRcvd]QSLRcvdSpec `json:"Records"`
}
