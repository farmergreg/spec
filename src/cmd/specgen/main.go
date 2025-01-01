package main

import (
	"go/format"
	"os"
	"path/filepath"

	"github.com/hamradiolog-net/adif-spec/src/pkg/datatype"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/antpath"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/arrlsection"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/awardsponsor"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/band"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/contest"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/continent"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/credit"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/mode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/morsekey"
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
	"github.com/hamradiolog-net/adif-spec/src/pkg/field"
)

// Generate the Go code for the ADIF spec
// n.b. We only generate constants for currently valid values.
// As a general rule, we skip deleted, un-released, and import-only records.
func main() {
	writeToFile("datatype", "datatype_gen.go", GenerateGoCodeForDataTypeDefinition(datatype.DataTypeList))
	writeToFile("field", "field_gen.go", GenerateGoCodeForFieldDefinition(field.FieldList))

	writeToFile("enum/antpath", "enum_ant_path_gen.go", GenerateGoCodeForAntPathEnumItem(antpath.EnumAntPathList))
	writeToFile("enum/arrlsection", "enum_arrl_section_gen.go", GenerateGoCodeForARRLSectionEnumItem(arrlsection.EnumARRLSectionList))
	writeToFile("enum/awardsponsor", "enum_award_sponsor_gen.go", GenerateGoCodeForAwardSponsorEnumItem(awardsponsor.EnumAwardSponsorList))
	// the award enum is depreciated, and there are no constants to export, so we skip it.
	// writeToFile("award", "enum_award_gen.go", GenerateGoCodeForAwardEnumItem(award.EnumAwardList))
	writeToFile("enum/band", "enum_band_gen.go", GenerateGoCodeForBandEnumItem(band.EnumBandList))
	writeToFile("enum/contest", "enum_contest_gen.go", GenerateGoCodeForContestEnumItem(contest.EnumContestList))
	writeToFile("enum/continent", "enum_continent_gen.go", GenerateGoCodeForContinentEnumItem(continent.EnumContinentList))
	writeToFile("enum/credit", "enum_credit_gen.go", GenerateGoCodeForCreditEnumItem(credit.EnumCreditList))
	writeToFile("enum/dxccentitycode", "enum_dxcc_entity_code_gen.go", GenerateGoCodeForDXCCEnumItem(dxccentitycode.EnumDXCCEntityCodeList))
	writeToFile("enum/mode", "enum_mode_gen.go", GenerateGoCodeForModeEnumItem(mode.EnumModeList))
	writeToFile("enum/morsekey", "enum_morse_key_gen.go", GenerateGoCodeForMorseKeyEnumItem(morsekey.EnumMorseKeyList))
	// These codes are not unique, so for now, we don't want to generate the list of constants...
	// writeToFile("enum/primary_administrative_subdivision_gen.go", GenerateGoCodeForPrimaryAdministrativeSubdivisionEnumItem(primaryadministrativesubdivision.EnumPrimaryAdministrativeSubdivisionList))
	writeToFile("enum/propagationmode", "enum_propagation_mode_gen.go", GenerateGoCodeForPropagationModeEnumItem(propagationmode.EnumPropagationModeList))
	writeToFile("enum/qslmedium", "enum_qsl_medium_gen.go", GenerateGoCodeForQSLMediumEnumItem(qslmedium.EnumQSLMediumList))
	writeToFile("enum/qslrcvd", "enum_qsl_rcvd_gen.go", GenerateGoCodeForQSLRcvdEnumItem(qslrcvd.EnumQSLRcvdList))
	writeToFile("enum/qslsent", "enum_qsl_sent_gen.go", GenerateGoCodeForQSLSentEnumItem(qslsent.EnumQSLSentList))
	writeToFile("enum/qslvia", "enum_qsl_via_gen.go", GenerateGoCodeForQSLViaEnumItem(qslvia.EnumQSLViaList))
	writeToFile("enum/qsocomplete", "enum_qso_complete_gen.go", GenerateGoCodeForQSOCompleteEnumItem(qsocomplete.EnumQSOCompleteList))
	writeToFile("enum/qsodownloadstatus", "enum_qso_download_status_gen.go", GenerateGoCodeForQSODownloadStatusEnumItem(qsodownloadstatus.EnumQSODownloadStatusList))
	writeToFile("enum/qsouploadstatus", "enum_qso_upload_status_gen.go", GenerateGoCodeForQSOUploadStatusEnumItem(qsouploadstatus.EnumQSOUploadStatusList))
	writeToFile("enum/region", "enum_region_gen.go", GenerateGoCodeForRegionEnumItem(region.EnumRegionList))
	writeToFile("enum/secondaryadministrativesubdivisionalt", "enum_secondary_administrative_subdivision_alt_gen.go", GenerateGoCodeForSecondaryAdministrativeSubdivisionAltEnumItem(secondaryadministrativesubdivisionalt.EnumSecondaryAdministrativeSubdivisionAltList))
	writeToFile("enum/secondaryadministrativesubdivision", "enum_secondary_administrative_subdivision_gen.go", GenerateGoCodeForSecondaryAdministrativeSubdivisionEnumItem(secondaryadministrativesubdivision.EnumSecondaryAdministrativeSubdivisionList))
	writeToFile("enum/mode", "enum_submode_gen.go", GenerateGoCodeForSubModeEnumItem(mode.EnumSubModeList))
}

func writeToFile(dir, filename, content string) {
	fullPath := filepath.Join("../../pkg", dir, filename)
	os.MkdirAll(filepath.Dir(fullPath), 0755)

	content = "// DO NOT EDIT\n// Code generated by ./cmd/codegen/codegen\n\n// This file contains CURRENTLY valid constants.\n// It is not a full list of all possible historic and/or future values.\n\n" + content

	formatted, fmtErr := format.Source([]byte(content))
	if fmtErr != nil {
		panic(fmtErr)
		// formatted = []byte(content)
	}

	err := os.WriteFile(fullPath, formatted, 0644)
	if err != nil {
		panic(err)
	}
}
