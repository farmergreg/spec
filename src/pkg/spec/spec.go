// package spec exposes the ADIF specification TSV files from:
// https://adif.org/
package spec

import _ "embed"

const (
	// ADIFVersion is the version of the ADIF specification.
	// This is the version that is currently supported by this package.
	// The specification was downloaded from:
	//
	// https://adif.org.uk/315/ADIF_315_resources_2024_11_28.zip
	ADIFVersion      = "3.1.5"
	ADIFVersionMajor = 3 // ADIF Major Version
	ADIFVersionMinor = 1 // ADIF Minor Version
	ADIFVersionPatch = 5 // ADIF Patch Version
)

//go:embed 315/enumerations_ant_path.tsv
var EnumAntPathTSV []byte

//go:embed 315/enumerations_arrl_section.tsv
var EnumARRLSectionTSV []byte

//go:embed 315/enumerations_award.tsv
var EnumAwardTSV []byte

//go:embed 315/enumerations_award_sponsor.tsv
var EnumAwardSponsorTSV []byte

//go:embed 315/enumerations_band.tsv
var EnumBandTSV []byte

//go:embed 315/enumerations_contest_id.tsv
var EnumContestTSV []byte

//go:embed 315/enumerations_continent.tsv
var EnumContinentTSV []byte

//go:embed 315/enumerations_credit.tsv
var EnumCreditTSV []byte

//go:embed 315/datatypes.tsv
var AdiDataTypesTSV []byte

//go:embed 315/enumerations_dxcc_entity_code.tsv
var EnumDXCCTSV []byte

//go:embed 315/fields.tsv
var FieldsTSV []byte

// FieldsExtraTSV is the TSV file containing common APP_ fields defined by 3rd parties.
// This file is not part of the ADIF specification.
// If you add to this file, please keep it alphabetized.
//
//go:embed fields_extra.tsv
var FieldsExtraTSV []byte

//go:embed 315/enumerations_mode.tsv
var EnumModeTSV []byte

//go:embed 315/enumerations_morse_key_type.tsv
var EnumMorseKeyTSV []byte

//go:embed 315/enumerations_primary_administrative_subdivision.tsv
var EnumPrimaryAdministrativeSubdivisionTSV []byte

//go:embed 315/enumerations_propagation_mode.tsv
var EnumPropagationModeTSV []byte

//go:embed 315/enumerations_qsl_medium.tsv
var EnumQSLMediumTSV []byte

//go:embed 315/enumerations_qsl_rcvd.tsv
var EnumQSLRcvdTSV []byte

//go:embed 315/enumerations_qsl_sent.tsv
var EnumQSLSentTSV []byte

//go:embed 315/enumerations_qsl_via.tsv
var EnumQSLViaTSV []byte

//go:embed 315/enumerations_qso_complete.tsv
var EnumQSOCompleteTSV []byte

//go:embed 315/enumerations_qso_download_status.tsv
var EnumQSODownloadStatusTSV []byte

//go:embed 315/enumerations_qso_upload_status.tsv
var EnumQSOUploadStatusTSV []byte

//go:embed 315/enumerations_region.tsv
var EnumRegionTSV []byte

//go:embed 315/enumerations_secondary_administrative_subdivision.tsv
var EnumSecondaryAdministrativeSubdivisionTSV []byte

//go:embed 315/enumerations_secondary_administrative_subdivision_alt.tsv
var EnumSecondaryAdministrativeSubdivisionAltTSV []byte

//go:embed 315/enumerations_submode.tsv
var EnumSubModeTSV []byte
