package qsodownloadstatus

import (
	"fmt"

	"github.com/hamradiolog-net/adif-spec/v6/spectype"
)

// SpecMapContainer contains all QSODownloadStatus specifications as defined by the ADIF Workgroup specification exports.
type SpecMapContainer struct {
	// Header  []string         `json:"Header"`
	Records map[QSODownloadStatus]Spec `json:"Records"`
}

// Spec represents the specification for a single QSODownloadStatus as defined by the ADIF Workgroup specification exports.
type Spec struct {
	// EnumerationName string           `json:"Enumeration Name"`
	IsImportOnly spectype.Boolean `json:"Import-only,omitempty"`
	// Comments     string            `json:"Comments,omitempty"`
	Key         QSODownloadStatus `json:"Status"` // Status
	Description string            `json:"Description"`
}

// Depreciated: CodeGeneratorMetadata is not part of the stable API and may change without warning in the future even for minor version numbers.
func (s Spec) CodeGeneratorMetadata() string {
	return fmt.Sprintf("%s = %s", s.Key, s.Description)
}
