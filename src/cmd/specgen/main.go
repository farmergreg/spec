package main

import (
	"go/format"
	"os"
	"path/filepath"

	"github.com/hamradiolog-net/adif-spec/src/pkg/adifield"
	"github.com/hamradiolog-net/adif-spec/src/pkg/aditype"
	"github.com/hamradiolog-net/adif-spec/src/pkg/antpath"
	"github.com/hamradiolog-net/adif-spec/src/pkg/arrlsection"
	"github.com/hamradiolog-net/adif-spec/src/pkg/awardsponsor"
	"github.com/hamradiolog-net/adif-spec/src/pkg/band"
	"github.com/hamradiolog-net/adif-spec/src/pkg/contest"
	"github.com/hamradiolog-net/adif-spec/src/pkg/continent"
	"github.com/hamradiolog-net/adif-spec/src/pkg/credit"
	"github.com/hamradiolog-net/adif-spec/src/pkg/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/mode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/morsekey"
	"github.com/hamradiolog-net/adif-spec/src/pkg/propagationmode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/qslmedium"
	"github.com/hamradiolog-net/adif-spec/src/pkg/qslrcvd"
	"github.com/hamradiolog-net/adif-spec/src/pkg/qslsent"
	"github.com/hamradiolog-net/adif-spec/src/pkg/qslvia"
	"github.com/hamradiolog-net/adif-spec/src/pkg/qsocomplete"
	"github.com/hamradiolog-net/adif-spec/src/pkg/qsodownloadstatus"
	"github.com/hamradiolog-net/adif-spec/src/pkg/qsouploadstatus"
	"github.com/hamradiolog-net/adif-spec/src/pkg/region"
	"github.com/hamradiolog-net/adif-spec/src/pkg/secondaryadministrativesubdivision"
	"github.com/hamradiolog-net/adif-spec/src/pkg/secondaryadministrativesubdivisionalt"
)

// Generate the Go code for the ADIF spec
// n.b. We only generate constants for currently valid values.
// As a general rule, we skip deleted, un-released, and import-only records.
func main() {
	writeToFile("aditype", "adi_data_type_gen.go", GenerateGoCodeForDataTypeDefinition(aditype.DataTypeList))
	writeToFile("adifield", "field_gen.go", GenerateGoCodeForFieldDefinition(adifield.FieldList))

	writeToFile("antpath", "enum_ant_path_gen.go", GenerateGoCodeForAntPathEnumItem(antpath.EnumAntPathList))
	writeToFile("arrlsection", "enum_arrl_section_gen.go", GenerateGoCodeForARRLSectionEnumItem(arrlsection.EnumARRLSectionList))
	writeToFile("awardsponsor", "enum_award_sponsor_gen.go", GenerateGoCodeForAwardSponsorEnumItem(awardsponsor.EnumAwardSponsorList))
	// the award enum is depreciated, and there are no constants to export, so we skip it.
	// writeToFile("award", "enum_award_gen.go", GenerateGoCodeForAwardEnumItem(award.EnumAwardList))
	writeToFile("band", "enum_band_gen.go", GenerateGoCodeForBandEnumItem(band.EnumBandList))
	writeToFile("contest", "enum_contest_gen.go", GenerateGoCodeForContestEnumItem(contest.EnumContestList))
	writeToFile("continent", "enum_continent_gen.go", GenerateGoCodeForContinentEnumItem(continent.EnumContinentList))
	writeToFile("credit", "enum_credit_gen.go", GenerateGoCodeForCreditEnumItem(credit.EnumCreditList))
	writeToFile("dxccentitycode", "enum_dxcc_entity_code_gen.go", GenerateGoCodeForDXCCEnumItem(dxccentitycode.EnumDXCCEntityCodeList))
	writeToFile("mode", "enum_mode_gen.go", GenerateGoCodeForModeEnumItem(mode.EnumModeList))
	writeToFile("morsekey", "enum_morse_key_gen.go", GenerateGoCodeForMorseKeyEnumItem(morsekey.EnumMorseKeyList))
	// These codes are not unique, so for now, we don't want to generate the list of constants...
	// writeToFile("enum_primary_administrative_subdivision_gen.go", GenerateGoCodeForPrimaryAdministrativeSubdivisionEnumItem(primaryadministrativesubdivision.EnumPrimaryAdministrativeSubdivisionList))
	writeToFile("propagationmode", "enum_propagation_mode_gen.go", GenerateGoCodeForPropagationModeEnumItem(propagationmode.EnumPropagationModeList))
	writeToFile("qslmedium", "enum_qsl_medium_gen.go", GenerateGoCodeForQSLMediumEnumItem(qslmedium.EnumQSLMediumList))
	writeToFile("qslrcvd", "enum_qsl_rcvd_gen.go", GenerateGoCodeForQSLRcvdEnumItem(qslrcvd.EnumQSLRcvdList))
	writeToFile("qslsent", "enum_qsl_sent_gen.go", GenerateGoCodeForQSLSentEnumItem(qslsent.EnumQSLSentList))
	writeToFile("qslvia", "enum_qsl_via_gen.go", GenerateGoCodeForQSLViaEnumItem(qslvia.EnumQSLViaList))
	writeToFile("qsocomplete", "enum_qso_complete_gen.go", GenerateGoCodeForQSOCompleteEnumItem(qsocomplete.EnumQSOCompleteList))
	writeToFile("qsodownloadstatus", "enum_qso_download_status_gen.go", GenerateGoCodeForQSODownloadStatusEnumItem(qsodownloadstatus.EnumQSODownloadStatusList))
	writeToFile("qsouploadstatus", "enum_qso_upload_status_gen.go", GenerateGoCodeForQSOUploadStatusEnumItem(qsouploadstatus.EnumQSOUploadStatusList))
	writeToFile("region", "enum_region_gen.go", GenerateGoCodeForRegionEnumItem(region.EnumRegionList))
	writeToFile("secondaryadministrativesubdivisionalt", "enum_secondary_administrative_subdivision_alt_gen.go", GenerateGoCodeForSecondaryAdministrativeSubdivisionAltEnumItem(secondaryadministrativesubdivisionalt.EnumSecondaryAdministrativeSubdivisionAltList))
	writeToFile("secondaryadministrativesubdivision", "enum_secondary_administrative_subdivision_gen.go", GenerateGoCodeForSecondaryAdministrativeSubdivisionEnumItem(secondaryadministrativesubdivision.EnumSecondaryAdministrativeSubdivisionList))
	writeToFile("mode", "enum_submode_gen.go", GenerateGoCodeForSubModeEnumItem(mode.EnumSubModeList))
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
