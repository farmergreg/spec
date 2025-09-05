package spec

import (
	"github.com/hamradiolog-net/adif-spec/v6/adifield"
	"github.com/hamradiolog-net/adif-spec/v6/aditype"
	"github.com/hamradiolog-net/adif-spec/v6/enum/antpath"
	"github.com/hamradiolog-net/adif-spec/v6/enum/arrlsection"
	"github.com/hamradiolog-net/adif-spec/v6/enum/award"
	"github.com/hamradiolog-net/adif-spec/v6/enum/awardsponsor"
	"github.com/hamradiolog-net/adif-spec/v6/enum/band"
	"github.com/hamradiolog-net/adif-spec/v6/enum/contest"
	"github.com/hamradiolog-net/adif-spec/v6/enum/continent"
	"github.com/hamradiolog-net/adif-spec/v6/enum/credit"
	"github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/v6/enum/eqslag"
	"github.com/hamradiolog-net/adif-spec/v6/enum/mode"
	"github.com/hamradiolog-net/adif-spec/v6/enum/morsekeytype"
	"github.com/hamradiolog-net/adif-spec/v6/enum/primaryadministrativesubdivision"
	"github.com/hamradiolog-net/adif-spec/v6/enum/propagationmode"
	"github.com/hamradiolog-net/adif-spec/v6/enum/qslmedium"
	"github.com/hamradiolog-net/adif-spec/v6/enum/qslrcvd"
	"github.com/hamradiolog-net/adif-spec/v6/enum/qslsent"
	"github.com/hamradiolog-net/adif-spec/v6/enum/qslvia"
	"github.com/hamradiolog-net/adif-spec/v6/enum/qsocomplete"
	"github.com/hamradiolog-net/adif-spec/v6/enum/qsodownloadstatus"
	"github.com/hamradiolog-net/adif-spec/v6/enum/qsouploadstatus"
	"github.com/hamradiolog-net/adif-spec/v6/enum/region"
	"github.com/hamradiolog-net/adif-spec/v6/enum/secondaryadministrativesubdivision"
	"github.com/hamradiolog-net/adif-spec/v6/enum/secondaryadministrativesubdivisionalt"
	"github.com/hamradiolog-net/adif-spec/v6/enum/submode"
	"github.com/hamradiolog-net/adif-spec/v6/internal/codegen"
	"github.com/hamradiolog-net/adif-spec/v6/spectype"
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

var _ codegen.CodeGenContainer = AdifSpec{}

func (a AdifSpec) CodeGeneratorMetadata() codegen.CodeGeneratorMetadataForContainer {
	return codegen.CodeGeneratorMetadataForContainer{
		PackageName: "spec",
		DataType:    "",
	}
}

func (a AdifSpec) CodeGeneratorRecords() map[codegen.CodeGeneratorEnumValue]codegen.CodeGenSpec {
	return nil
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
