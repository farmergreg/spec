package spec

// This is the root object.
// It can be used to import and process the ADIF Workgroup's all.json specification export.
type AdifSpec struct {
	Adif struct {
		Version      string           `json:"Version"`
		Status       string           `json:"Status"`
		Date         AdifSpecDateTime `json:"Date"`
		Created      AdifSpecDateTime `json:"Created"`
		DataTypes    DataTypeSpecMap  `json:"DataTypes"`
		Enumerations Enumerations     `json:"Enumerations"`
		Fields       FieldSpecMap     `json:"Fields"`
	} `json:"Adif"`
}

type Enumerations struct {
	Ant_Path                                 AntPathEnumerationSpecMap                    `json:"Ant_Path,omitempty"`
	ARRL_Section                             ARRLSectionSpecMap                           `json:"ARRL_Section,omitempty"`
	Award                                    AwardSpecMap                                 `json:"Award,omitempty"`
	Award_Sponsor                            AwardSponsorSpecMap                          `json:"Award_Sponsor,omitempty"`
	Band                                     BandSpecMap                                  `json:"Band,omitempty"`
	Contest_ID                               ContestSpecMap                               `json:"Contest_ID,omitempty"`
	Continent                                ContinentSpecMap                             `json:"Continent,omitempty"`
	Credit                                   CreditSpecMap                                `json:"Credit,omitempty"`
	DXCC_Entity_Code                         DXCCEntityCodeSpecMap                        `json:"DXCC_Entity_Code,omitempty"`
	EQSL_AG                                  EQSLAGSpecMap                                `json:"EQSL_AG,omitempty"`
	Mode                                     ModeSpecMap                                  `json:"Mode,omitempty"`
	Morse_Key_Type                           MorseKeyTypeSpecMap                          `json:"Morse_Key_Type,omitempty"`
	Primary_Administrative_Subdivision       PrimaryAdministrativeSubdivisionSpecMap      `json:"Primary_Administrative_Subdivision,omitempty"`
	Propagation_Mode                         PropagationModeSpecMap                       `json:"Propagation_Mode,omitempty"`
	QSL_Medium                               QSLMediumSpecMap                             `json:"QSL_Medium,omitempty"`
	QSL_Rcvd                                 QSLRcvdSpecMap                               `json:"QSL_Rcvd,omitempty"`
	QSL_Sent                                 QSLSentSpecMap                               `json:"QSL_Sent,omitempty"`
	QSL_Via                                  QSLViaSpecMap                                `json:"QSL_Via,omitempty"`
	QSO_Complete                             QSOCompleteSpecMap                           `json:"QSO_Complete,omitempty"`
	QSO_Download_Status                      QSODownloadStatusSpecMap                     `json:"QSO_Download_Status,omitempty"`
	QSO_Upload_Status                        QSOUploadStatusSpecMap                       `json:"QSO_Upload_Status,omitempty"`
	Region                                   RegionSpecMap                                `json:"Region,omitempty"`
	Secondary_Administrative_Subdivision     SecondaryAdministrativeSubdivisionSpecMap    `json:"Secondary_Administrative_Subdivision,omitempty"`
	Secondary_Administrative_Subdivision_Alt SecondaryAdministrativeSubdivisionAltSpecMap `json:"Secondary_Administrative_Subdivision_Alt,omitempty"`
	Submode                                  SubmodeSpecMap                               `json:"Submode,omitempty"`
}
