package spec

import (
	"fmt"
	"strings"

	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/band"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/mode"
)

// This is the root object.
// It can be used to import and process the ADIF Workgroup's all.json specification export.
type AdifSpec struct {
	Adif struct {
		Version      string           `json:"Version"`
		Status       string           `json:"Status"`
		Date         AdifSpecDateTime `json:"Date"`
		Created      AdifSpecDateTime `json:"Created"`
		DataTypes    DataTypes        `json:"DataTypes"`
		Enumerations Enumerations     `json:"Enumerations"`
		Fields       Fields           `json:"Fields"`
	} `json:"Adif"`
}

type DataType struct {
	Id                string          `json:"Data Type Name"` // Data Type Name
	DataTypeIndicator string          `json:"Data Type Indicator,omitempty"`
	Description       string          `json:"Description"`
	MinimumValue      AdifSpecInteger `json:"Minimum Value,omitempty"`
	MaximumValue      AdifSpecInteger `json:"Maximum Value,omitempty"`
	IsImportOnly      AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments          string          `json:"Comments,omitempty"`
}

type DataTypes struct {
	Header  []string            `json:"Header"`
	Records map[string]DataType `json:"Records"`
}

type Field struct {
	Id            string          `json:"Field Name"` // Field Name
	DataType      string          `json:"Data Type"`
	Enumeration   string          `json:"Enumeration,omitempty"`
	Description   string          `json:"Description"`
	IsHeaderField string          `json:"Header Field,omitempty"`
	MinimumValue  AdifSpecInteger `json:"Minimum Value,omitempty"`
	MaximumValue  AdifSpecInteger `json:"Maximum Value,omitempty"`
	IsImportOnly  AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments      string          `json:"Comments,omitempty"`
}

type Fields struct {
	Header  []string         `json:"Header"`
	Records map[string]Field `json:"Records"`
}

// Command fields used in all enumerations
type BaseEnumerationRecord struct {
	Id           string          `json:"Enumeration Name"` // Enumeration Name
	IsImportOnly AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments     string          `json:"Comments,omitempty"`
}

type AntPathRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Abbreviation"` // Abbreviation
	Description string `json:"Meaning"`
}

type AntPathEnumeration struct {
	Header  []string                 `json:"Header"`
	Records map[string]AntPathRecord `json:"Records"`
}

type ARRLSectionRecord struct {
	BaseEnumerationRecord
	Id             string                            `json:"Abbreviation"` // Abbreviation
	Description    string                            `json:"Section Name"` // Section Name
	DXCCEntityCode dxccentitycode.DXCCEntityCodeList `json:"DXCC Entity Code"`
	FromDate       AdifSpecDateTime                  `json:"From Date,omitempty"`
	DeletedDate    AdifSpecDateTime                  `json:"Deleted Date,omitempty"`
}

type ARRLSectionEnumeration struct {
	Header  []string                     `json:"Header"`
	Records map[string]ARRLSectionRecord `json:"Records"`
}

type AwardRecord struct {
	BaseEnumerationRecord
	Id string `json:"Award"` // Award
}

func (a AwardRecord) Description() string {
	return a.Id
}

type AwardEnumeration struct {
	Header  []string               `json:"Header"`
	Records map[string]AwardRecord `json:"Records"`
}

type AwardSponsorRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Sponsor"`                 // Sponsor Prefix
	Description string `json:"Sponsoring Organization"` // Sponsoring Organization
}

type AwardSponsorEnumeration struct {
	Header  []string                      `json:"Header"`
	Records map[string]AwardSponsorRecord `json:"Records"`
}

type BandRecord struct {
	BaseEnumerationRecord
	Id           string   `json:"Band"` // Band
	LowerFreqMHz band.MHz `json:"Lower Freq (MHz)"`
	UpperFreqMHz band.MHz `json:"Upper Freq (MHz)"`
}

func (b BandRecord) Description() string {
	return fmt.Sprintf("%-6s %12.4f MHz - %12.4f MHz", b.Id, b.LowerFreqMHz, b.UpperFreqMHz)
}

type BandEnumeration struct {
	Header  []string              `json:"Header"`
	Records map[string]BandRecord `json:"Records"`
}

type ContestIDRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Contest-ID"` // Contest ID
	Description string `json:"Description"`
}

type ContestIDEnumeration struct {
	Header  []string                   `json:"Header"`
	Records map[string]ContestIDRecord `json:"Records"`
}

type ContinentRecord struct {
	BaseEnumerationRecord
	Id        string `json:"Abbreviation"` // Abbreviation
	Continent string `json:"Continent"`
}

func (c ContinentRecord) Description() string {
	return c.Continent
}

type ContinentEnumeration struct {
	Header  []string                   `json:"Header"`
	Records map[string]ContinentRecord `json:"Records"`
}

type CreditRecord struct {
	BaseEnumerationRecord
	Id      string `json:"Credit For"` // Credit For
	Sponsor string `json:"Sponsor"`
	Award   string `json:"Award"`
	Facet   string `json:"Facet"`
}

func (c CreditRecord) Description() string {
	return fmt.Sprintf("%s %s %s", c.Sponsor, c.Award, c.Facet)
}

type CreditEnumeration struct {
	Header  []string                `json:"Header"`
	Records map[string]CreditRecord `json:"Records"`
}

type DXCCEntityCodeRecord struct {
	BaseEnumerationRecord
	Id         string          `json:"Entity Code"` // Entity Code
	EntityName string          `json:"Entity Name"`
	IsDeleted  AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (d DXCCEntityCodeRecord) Identifier() string {
	name := d.Description()

	if strings.HasPrefix(name, "None ") {
		name = "NONE"
	}

	if bool(d.IsDeleted) {
		name += "_DELETED"
	}

	return name
}

func (d DXCCEntityCodeRecord) Description() string {
	return d.EntityName
}

type DXCCEntityCodeEnumeration struct {
	Header  []string                        `json:"Header"`
	Records map[string]DXCCEntityCodeRecord `json:"Records"`
}

type EQSLAGRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Status"` // Status
	Description string `json:"Description"`
}

type EQSLAGEnumeration struct {
	Header  []string                `json:"Header"`
	Records map[string]EQSLAGRecord `json:"Records"`
}

type ModeRecord struct {
	BaseEnumerationRecord
	Id          string                   `json:"Mode"` // Mode
	Submodes    mode.EnumModeSubModeList `json:"Submodes,omitempty"`
	Description string                   `json:"Description,omitempty"`
}

type ModeEnumeration struct {
	Header  []string              `json:"Header"`
	Records map[string]ModeRecord `json:"Records"`
}

type MorseKeyTypeRecord struct {
	BaseEnumerationRecord
	Id               string `json:"Abbreviation"` // Abbreviation
	Description      string `json:"Meaning"`
	Characteristics  string `json:"Characteristics,omitempty"`
	MorseComposition string `json:"Morse Composition,omitempty"`
	Examples         string `json:"Examples,omitempty"`
}

type MorseKeyTypeEnumeration struct {
	Header  []string                      `json:"Header"`
	Records map[string]MorseKeyTypeRecord `json:"Records"`
}

type PrimaryAdministrativeSubdivisionRecord struct {
	BaseEnumerationRecord
	Code            string          `json:"Code"` // Code
	PrimaryAdminSub string          `json:"Primary Administrative Subdivision"`
	DXCCEntityCode  string          `json:"DXCC Entity Code"`
	ContainedWithin string          `json:"Contained Within,omitempty"`
	Oblast          string          `json:"Oblast #,omitempty"`
	CQZone          string          `json:"CQ Zone,omitempty"`
	ITUZone         string          `json:"ITU Zone,omitempty"`
	Prefix          string          `json:"Prefix,omitempty"`
	IsDeleted       AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (p PrimaryAdministrativeSubdivisionRecord) Description() string {
	return p.PrimaryAdminSub
}

type PrimaryAdministrativeSubdivisionEnumeration struct {
	Header  []string                                          `json:"Header"`
	Records map[string]PrimaryAdministrativeSubdivisionRecord `json:"Records"`
}

type PropagationModeRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Enumeration"` // Enumeration
	Description string `json:"Description"`
}

type PropagationModeEnumeration struct {
	Header  []string                         `json:"Header"`
	Records map[string]PropagationModeRecord `json:"Records"`
}

type QSLMediumRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Medium"` // Medium
	Description string `json:"Description"`
}

type QSLMediumEnumeration struct {
	Header  []string                   `json:"Header"`
	Records map[string]QSLMediumRecord `json:"Records"`
}

type QSLRcvdRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Status"` // Status
	Meaning     string `json:"Meaning"`
	Description string `json:"Description"`
}

type QSLRcvdEnumeration struct {
	Header  []string                 `json:"Header"`
	Records map[string]QSLRcvdRecord `json:"Records"`
}

type QSLSentRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Status"` // Status
	Meaning     string `json:"Meaning"`
	Description string `json:"Description"`
}

type QSLSentEnumeration struct {
	Header  []string                 `json:"Header"`
	Records map[string]QSLSentRecord `json:"Records"`
}

type QSLViaRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Via"` // Via
	Description string `json:"Description"`
}

type QSLViaEnumeration struct {
	Header  []string                `json:"Header"`
	Records map[string]QSLViaRecord `json:"Records"`
}

type QSOCompleteRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Abbreviation"` // Abbreviation
	Description string `json:"Meaning"`
}

type QSOCompleteEnumeration struct {
	Header  []string                     `json:"Header"`
	Records map[string]QSOCompleteRecord `json:"Records"`
}

type QSODownloadStatusRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Status"` // Status
	Description string `json:"Description"`
}

type QSODownloadStatusEnumeration struct {
	Header  []string                           `json:"Header"`
	Records map[string]QSODownloadStatusRecord `json:"Records"`
}

type QSOUploadStatusRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Status"` // Status
	Description string `json:"Description"`
}

type QSOUploadStatusEnumeration struct {
	Header  []string                         `json:"Header"`
	Records map[string]QSOUploadStatusRecord `json:"Records"`
}

type RegionRecord struct {
	BaseEnumerationRecord
	Id             string           `json:"Region Entity Code"` // Region Entity Code
	DXCCEntityCode string           `json:"DXCC Entity Code"`
	Region         string           `json:"Region"`
	Prefix         string           `json:"Prefix,omitempty"`
	Applicability  string           `json:"Applicability,omitempty"`
	StartDate      AdifSpecDateOnly `json:"Start Date,omitempty"`
	EndDate        AdifSpecDateOnly `json:"End Date,omitempty"`
}

func (r RegionRecord) Description() string {
	return r.Region
}

type RegionEnumeration struct {
	Header  []string                `json:"Header"`
	Records map[string]RegionRecord `json:"Records"`
}

type SecondaryAdministrativeSubdivisionRecord struct {
	BaseEnumerationRecord
	Id                     string          `json:"Code"` // Code
	SecondaryAdminSub      string          `json:"Secondary Administrative Subdivision"`
	DXCCEntityCode         string          `json:"DXCC Entity Code"`
	AlaskaJudicialDistrict string          `json:"Alaska Judicial District,omitempty"`
	IsDeleted              AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (s SecondaryAdministrativeSubdivisionRecord) Description() string {
	return s.SecondaryAdminSub
}

type SecondaryAdministrativeSubdivisionEnumeration struct {
	Header  []string                                            `json:"Header"`
	Records map[string]SecondaryAdministrativeSubdivisionRecord `json:"Records"`
}

type SecondaryAdministrativeSubdivisionAltRecord struct {
	BaseEnumerationRecord
	Id             string          `json:"Code"` // Code
	DXCCEntityCode string          `json:"DXCC Entity Code"`
	Region         string          `json:"Region"`
	District       string          `json:"District"`
	IsDeleted      AdifSpecBoolean `json:"Deleted,omitempty"`
}

func (s SecondaryAdministrativeSubdivisionAltRecord) Description() string {
	return s.Region
}

type SecondaryAdministrativeSubdivisionAltEnumeration struct {
	Header  []string                                               `json:"Header"`
	Records map[string]SecondaryAdministrativeSubdivisionAltRecord `json:"Records"`
}

type SubmodeRecord struct {
	BaseEnumerationRecord
	Id          string `json:"Submode"` // Submode
	Mode        string `json:"Mode"`
	Description string `json:"Description,omitempty"`
}

type SubmodeEnumeration struct {
	Header  []string                 `json:"Header"`
	Records map[string]SubmodeRecord `json:"Records"`
}

type Enumerations struct {
	Ant_Path                                 AntPathEnumeration                               `json:"Ant_Path,omitempty"`
	ARRL_Section                             ARRLSectionEnumeration                           `json:"ARRL_Section,omitempty"`
	Award                                    AwardEnumeration                                 `json:"Award,omitempty"`
	Award_Sponsor                            AwardSponsorEnumeration                          `json:"Award_Sponsor,omitempty"`
	Band                                     BandEnumeration                                  `json:"Band,omitempty"`
	Contest_ID                               ContestIDEnumeration                             `json:"Contest_ID,omitempty"`
	Continent                                ContinentEnumeration                             `json:"Continent,omitempty"`
	Credit                                   CreditEnumeration                                `json:"Credit,omitempty"`
	DXCC_Entity_Code                         DXCCEntityCodeEnumeration                        `json:"DXCC_Entity_Code,omitempty"`
	EQSL_AG                                  EQSLAGEnumeration                                `json:"EQSL_AG,omitempty"`
	Mode                                     ModeEnumeration                                  `json:"Mode,omitempty"`
	Morse_Key_Type                           MorseKeyTypeEnumeration                          `json:"Morse_Key_Type,omitempty"`
	Primary_Administrative_Subdivision       PrimaryAdministrativeSubdivisionEnumeration      `json:"Primary_Administrative_Subdivision,omitempty"`
	Propagation_Mode                         PropagationModeEnumeration                       `json:"Propagation_Mode,omitempty"`
	QSL_Medium                               QSLMediumEnumeration                             `json:"QSL_Medium,omitempty"`
	QSL_Rcvd                                 QSLRcvdEnumeration                               `json:"QSL_Rcvd,omitempty"`
	QSL_Sent                                 QSLSentEnumeration                               `json:"QSL_Sent,omitempty"`
	QSL_Via                                  QSLViaEnumeration                                `json:"QSL_Via,omitempty"`
	QSO_Complete                             QSOCompleteEnumeration                           `json:"QSO_Complete,omitempty"`
	QSO_Download_Status                      QSODownloadStatusEnumeration                     `json:"QSO_Download_Status,omitempty"`
	QSO_Upload_Status                        QSOUploadStatusEnumeration                       `json:"QSO_Upload_Status,omitempty"`
	Region                                   RegionEnumeration                                `json:"Region,omitempty"`
	Secondary_Administrative_Subdivision     SecondaryAdministrativeSubdivisionEnumeration    `json:"Secondary_Administrative_Subdivision,omitempty"`
	Secondary_Administrative_Subdivision_Alt SecondaryAdministrativeSubdivisionAltEnumeration `json:"Secondary_Administrative_Subdivision_Alt,omitempty"`
	Submode                                  SubmodeEnumeration                               `json:"Submode,omitempty"`
}
