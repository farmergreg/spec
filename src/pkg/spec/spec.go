package spec

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

// BaseEnumerationSpec contains common fields present in all enumerations
type BaseEnumerationSpec struct {
	EnumerationName string          `json:"Enumeration Name"`
	IsImportOnly    AdifSpecBoolean `json:"Import-only,omitempty"`
	Comments        string          `json:"Comments,omitempty"`
}

type Enumerations struct {
	Ant_Path                                 AntPathEnumerationSpec                           `json:"Ant_Path,omitempty"`
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
