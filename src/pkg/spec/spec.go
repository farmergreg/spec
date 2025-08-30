package spec

import (
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/adifield"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/aditype"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/antpath"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/arrlsection"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/award"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/awardsponsor"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/band"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/contest"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/continent"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/credit"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/eqslag"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/mode"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/morsekeytype"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/primaryadministrativesubdivision"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/propagationmode"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/qslmedium"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/qslrcvd"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/qslsent"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/qslvia"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/qsocomplete"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/qsodownloadstatus"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/qsouploadstatus"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/region"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/secondaryadministrativesubdivision"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/secondaryadministrativesubdivisionalt"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/submode"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/spectype"
)

// This is the root object for JSON unmarshaling.
// It can be used to import and process the ADIF Workgroup's all.json specification export.
type AdifSpecContainer struct {
	AdifSpec AdifSpec `json:"Adif"`
}

// This is the main ADIF specification
type AdifSpec struct {
	Version   string                    `json:"Version"`
	Status    string                    `json:"Status"`
	Date      spectype.DateTime         `json:"Date"`
	Created   spectype.DateTime         `json:"Created"`
	DataTypes aditype.SpecMapContainer  `json:"DataTypes"`
	Fields    adifield.SpecMapContainer `json:"Fields"`
	Enum      Enum                      `json:"Enumerations"`
}

// Enumerations defined in the ADIF specification
type Enum struct {
	Ant_Path                                 antpath.SpecMapContainer                               `json:"Ant_Path,omitempty"`
	ARRL_Section                             arrlsection.SpecMapContainer                           `json:"ARRL_Section,omitempty"`
	Award                                    award.SpecMapContainer                                 `json:"Award,omitempty"`
	Award_Sponsor                            awardsponsor.SpecMapContainer                          `json:"Award_Sponsor,omitempty"`
	Band                                     band.SpecMapContainer                                  `json:"Band,omitempty"`
	Contest_ID                               contest.SpecMapContainer                               `json:"Contest_ID,omitempty"`
	Continent                                continent.SpecMapContainer                             `json:"Continent,omitempty"`
	Credit                                   credit.SpecMapContainer                                `json:"Credit,omitempty"`
	DXCC_Entity_Code                         dxccentitycode.SpecMapContainer                        `json:"DXCC_Entity_Code,omitempty"`
	EQSL_AG                                  eqslag.SpecMapContainer                                `json:"EQSL_AG,omitempty"`
	Mode                                     mode.SpecMapContainer                                  `json:"Mode,omitempty"`
	Morse_Key_Type                           morsekeytype.SpecMapContainer                          `json:"Morse_Key_Type,omitempty"`
	Primary_Administrative_Subdivision       primaryadministrativesubdivision.SpecMapContainer      `json:"Primary_Administrative_Subdivision,omitempty"`
	Propagation_Mode                         propagationmode.SpecMapContainer                       `json:"Propagation_Mode,omitempty"`
	QSL_Medium                               qslmedium.SpecMapContainer                             `json:"QSL_Medium,omitempty"`
	QSL_Rcvd                                 qslrcvd.SpecMapContainer                               `json:"QSL_Rcvd,omitempty"`
	QSL_Sent                                 qslsent.SpecMapContainer                               `json:"QSL_Sent,omitempty"`
	QSL_Via                                  qslvia.SpecMapContainer                                `json:"QSL_Via,omitempty"`
	QSO_Complete                             qsocomplete.SpecMapContainer                           `json:"QSO_Complete,omitempty"`
	QSO_Download_Status                      qsodownloadstatus.SpecMapContainer                     `json:"QSO_Download_Status,omitempty"`
	QSO_Upload_Status                        qsouploadstatus.SpecMapContainer                       `json:"QSO_Upload_Status,omitempty"`
	Region                                   region.SpecMapContainer                                `json:"Region,omitempty"`
	Secondary_Administrative_Subdivision     secondaryadministrativesubdivision.SpecMapContainer    `json:"Secondary_Administrative_Subdivision,omitempty"`
	Secondary_Administrative_Subdivision_Alt secondaryadministrativesubdivisionalt.SpecMapContainer `json:"Secondary_Administrative_Subdivision_Alt,omitempty"`
	Submode                                  submode.SpecMapContainer                               `json:"Submode,omitempty"`
}
