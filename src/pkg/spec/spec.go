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

// This is the root object.
// It can be used to import and process the ADIF Workgroup's all.json specification export.
type AdifSpec struct {
	Adif struct {
		Version      string                    `json:"Version"`
		Status       string                    `json:"Status"`
		Date         spectype.AdifSpecDateTime `json:"Date"`
		Created      spectype.AdifSpecDateTime `json:"Created"`
		DataTypes    datatype.DataTypeSpecMap  `json:"DataTypes"`
		Enumerations Enumerations              `json:"Enumerations"`
		Fields       field.FieldSpecMap        `json:"Fields"`
	} `json:"Adif"`
}

type Enumerations struct {
	Ant_Path                                 antpath.AntPathEnumerationSpecMap                                                  `json:"Ant_Path,omitempty"`
	ARRL_Section                             arrlsection.ARRLSectionSpecMap                                                     `json:"ARRL_Section,omitempty"`
	Award                                    award.AwardSpecMap                                                                 `json:"Award,omitempty"`
	Award_Sponsor                            awardsponsor.AwardSponsorSpecMap                                                   `json:"Award_Sponsor,omitempty"`
	Band                                     band.BandSpecMap                                                                   `json:"Band,omitempty"`
	Contest_ID                               contest.ContestSpecMap                                                             `json:"Contest_ID,omitempty"`
	Continent                                continent.ContinentSpecMap                                                         `json:"Continent,omitempty"`
	Credit                                   credit.CreditSpecMap                                                               `json:"Credit,omitempty"`
	DXCC_Entity_Code                         dxccentitycode.DXCCEntityCodeSpecMap                                               `json:"DXCC_Entity_Code,omitempty"`
	EQSL_AG                                  eqslag.EQSLAGSpecMap                                                               `json:"EQSL_AG,omitempty"`
	Mode                                     mode.ModeSpecMap                                                                   `json:"Mode,omitempty"`
	Morse_Key_Type                           morsekeytype.MorseKeyTypeSpecMap                                                   `json:"Morse_Key_Type,omitempty"`
	Primary_Administrative_Subdivision       primaryadministrativesubdivision.PrimaryAdministrativeSubdivisionSpecMap           `json:"Primary_Administrative_Subdivision,omitempty"`
	Propagation_Mode                         propagationmode.PropagationModeSpecMap                                             `json:"Propagation_Mode,omitempty"`
	QSL_Medium                               qslmedium.QSLMediumSpecMap                                                         `json:"QSL_Medium,omitempty"`
	QSL_Rcvd                                 qslrcvd.QSLRcvdSpecMap                                                             `json:"QSL_Rcvd,omitempty"`
	QSL_Sent                                 qslsent.QSLSentSpecMap                                                             `json:"QSL_Sent,omitempty"`
	QSL_Via                                  qslvia.QSLViaSpecMap                                                               `json:"QSL_Via,omitempty"`
	QSO_Complete                             qsocomplete.QSOCompleteSpecMap                                                     `json:"QSO_Complete,omitempty"`
	QSO_Download_Status                      qsodownloadstatus.QSODownloadStatusSpecMap                                         `json:"QSO_Download_Status,omitempty"`
	QSO_Upload_Status                        qsouploadstatus.QSOUploadStatusSpecMap                                             `json:"QSO_Upload_Status,omitempty"`
	Region                                   region.RegionSpecMap                                                               `json:"Region,omitempty"`
	Secondary_Administrative_Subdivision     secondaryadministrativesubdivision.SecondaryAdministrativeSubdivisionSpecMap       `json:"Secondary_Administrative_Subdivision,omitempty"`
	Secondary_Administrative_Subdivision_Alt secondaryadministrativesubdivisionalt.SecondaryAdministrativeSubdivisionAltSpecMap `json:"Secondary_Administrative_Subdivision_Alt,omitempty"`
	Submode                                  submode.SubmodeSpecMap                                                             `json:"Submode,omitempty"`
}
