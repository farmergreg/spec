package spec

import (
	"github.com/hamradiolog-net/adif-spec/src/pkg/datatype"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/antpath"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/arrlsection"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/award"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/awardsponsor"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/band"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/contest"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/continent"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/credit"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/eqslag"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/mode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/morsekeytype"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/primaryadministrativesubdivision"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/propagationmode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/qslmedium"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/qslrcvd"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/qslsent"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/qslvia"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/qsocomplete"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/qsodownloadstatus"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/qsouploadstatus"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/region"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/secondaryadministrativesubdivision"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/secondaryadministrativesubdivisionalt"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/submode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/field"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spectype"
)

// This is the root object for JSON unmarshaling.
// It can be used to import and process the ADIF Workgroup's all.json specification export.
type AdifSpecContainer struct {
	AdifSpec AdifSpec `json:"Adif"`
}

// This is the main ADIF specification structure
type AdifSpec struct {
	Version   string            `json:"Version"`
	Status    string            `json:"Status"`
	Date      spectype.DateTime `json:"Date"`
	Created   spectype.DateTime `json:"Created"`
	DataTypes datatype.SpecMap  `json:"DataTypes"`
	Fields    field.SpecMap     `json:"Fields"`
	Enum      Enum              `json:"Enumerations"`
}

// Enumerations defined in the ADIF specification.
type Enum struct {
	Ant_Path                                 antpath.SpecMap                               `json:"Ant_Path,omitempty"`
	ARRL_Section                             arrlsection.SpecMap                           `json:"ARRL_Section,omitempty"`
	Award                                    award.SpecMap                                 `json:"Award,omitempty"`
	Award_Sponsor                            awardsponsor.SpecMap                          `json:"Award_Sponsor,omitempty"`
	Band                                     band.SpecMap                                  `json:"Band,omitempty"`
	Contest_ID                               contest.SpecMap                               `json:"Contest_ID,omitempty"`
	Continent                                continent.SpecMap                             `json:"Continent,omitempty"`
	Credit                                   credit.SpecMap                                `json:"Credit,omitempty"`
	DXCC_Entity_Code                         dxccentitycode.SpecMap                        `json:"DXCC_Entity_Code,omitempty"`
	EQSL_AG                                  eqslag.SpecMap                                `json:"EQSL_AG,omitempty"`
	Mode                                     mode.SpecMap                                  `json:"Mode,omitempty"`
	Morse_Key_Type                           morsekeytype.SpecMap                          `json:"Morse_Key_Type,omitempty"`
	Primary_Administrative_Subdivision       primaryadministrativesubdivision.SpecMap      `json:"Primary_Administrative_Subdivision,omitempty"`
	Propagation_Mode                         propagationmode.SpecMap                       `json:"Propagation_Mode,omitempty"`
	QSL_Medium                               qslmedium.SpecMap                             `json:"QSL_Medium,omitempty"`
	QSL_Rcvd                                 qslrcvd.SpecMap                               `json:"QSL_Rcvd,omitempty"`
	QSL_Sent                                 qslsent.SpecMap                               `json:"QSL_Sent,omitempty"`
	QSL_Via                                  qslvia.SpecMap                                `json:"QSL_Via,omitempty"`
	QSO_Complete                             qsocomplete.SpecMap                           `json:"QSO_Complete,omitempty"`
	QSO_Download_Status                      qsodownloadstatus.SpecMap                     `json:"QSO_Download_Status,omitempty"`
	QSO_Upload_Status                        qsouploadstatus.SpecMap                       `json:"QSO_Upload_Status,omitempty"`
	Region                                   region.SpecMap                                `json:"Region,omitempty"`
	Secondary_Administrative_Subdivision     secondaryadministrativesubdivision.SpecMap    `json:"Secondary_Administrative_Subdivision,omitempty"`
	Secondary_Administrative_Subdivision_Alt secondaryadministrativesubdivisionalt.SpecMap `json:"Secondary_Administrative_Subdivision_Alt,omitempty"`
	Submode                                  submode.SpecMap                               `json:"Submode,omitempty"`
}
