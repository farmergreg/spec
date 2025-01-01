package main

import (
	"fmt"
	"strings"

	"github.com/hamradiolog-net/adif-spec/src/pkg/adifield"
	"github.com/hamradiolog-net/adif-spec/src/pkg/aditype"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/antpath"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/arrlsection"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/award"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/awardsponsor"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/band"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/contest"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/continent"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/credit"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/mode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/morsekey"
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
)

func GenerateGoCodeForDataTypeDefinition(a []*aditype.DataTypeDefinition) string {
	return generateGoCode(
		a,
		"aditype",
		"",
		"DataType",
		nil,
		func(a *aditype.DataTypeDefinition) string { return string(a.ID) },
		func(a *aditype.DataTypeDefinition) string { return string(a.ID) },
		func(a *aditype.DataTypeDefinition) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForFieldDefinition(a []*adifield.FieldDefinition) string {
	return generateGoCode(
		a,
		"adifield",
		"",
		"Field",
		func(a *adifield.FieldDefinition) bool {
			switch a.ID {
			case "USERDEFn":
				// USERDEFn is just an example field that appears in our source data.
				// Removing it because it is not real.
				return false
			default:
				return true
			}
		},
		func(a *adifield.FieldDefinition) string { return string(a.ID) },
		func(a *adifield.FieldDefinition) string { return string(a.ID) },
		func(a *adifield.FieldDefinition) string {

			header := "QSO"
			if a.IsHeaderField {
				header = "Header"
			}
			return fmt.Sprintf("%-30s = %s: %s", a.ID, header, a.Description)
		},
	)
}

func GenerateGoCodeForAntPathEnumItem(a []*antpath.EnumAntPathItem) string {
	return generateGoCode(
		a,
		"antpath",
		"",
		"AntPath",
		nil,
		func(a *antpath.EnumAntPathItem) string { return string(a.ID) },
		func(a *antpath.EnumAntPathItem) string { return string(a.ID) },
		func(a *antpath.EnumAntPathItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForARRLSectionEnumItem(a []*arrlsection.EnumARRLSectionItem) string {
	return generateGoCode(
		a,
		"arrlsection",
		"",
		"ARRLSection",
		nil,
		func(a *arrlsection.EnumARRLSectionItem) string { return string(a.ID) },
		func(a *arrlsection.EnumARRLSectionItem) string { return string(a.ID) },
		func(a *arrlsection.EnumARRLSectionItem) string { return fmt.Sprintf("%-5s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForAwardSponsorEnumItem(a []*awardsponsor.EnumAwardSponsorItem) string {
	return generateGoCode(
		a,
		"awardsponsor",
		"",
		"AwardSponsorPrefix",
		nil,
		func(a *awardsponsor.EnumAwardSponsorItem) string { return string(a.IDPrefix) },
		func(a *awardsponsor.EnumAwardSponsorItem) string { return string(a.IDPrefix) },
		func(a *awardsponsor.EnumAwardSponsorItem) string {
			return fmt.Sprintf("%-7s = %s", a.IDPrefix, a.Description)
		},
	)
}

func GenerateGoCodeForAwardEnumItem(a []*award.EnumAwardItem) string {
	return generateGoCode(
		a,
		"award",
		"",
		"Award",
		nil,
		func(a *award.EnumAwardItem) string { return string(a.ID) },
		func(a *award.EnumAwardItem) string { return string(a.ID) },
		func(a *award.EnumAwardItem) string { return string(a.ID) },
	)
}

func GenerateGoCodeForBandEnumItem(a []*band.EnumBandItem) string {
	return generateGoCode(
		a,
		"band",
		"Band",
		"Band",
		nil,
		func(a *band.EnumBandItem) string { return string(a.ID) },
		func(a *band.EnumBandItem) string { return string(a.ID) },
		func(a *band.EnumBandItem) string {
			return fmt.Sprintf("%12.4f MHz - %12.4f MHz = %s", a.LowerFreqMHz, a.UpperFreqMHz, a.ID)
		},
	)
}

func GenerateGoCodeForContestEnumItem(a []*contest.EnumContestItem) string {
	return generateGoCode(
		a,
		"contest",
		"Contest_",
		"Contest",
		nil,
		func(a *contest.EnumContestItem) string { return string(a.ID) },
		func(a *contest.EnumContestItem) string { return string(a.ID) },
		func(a *contest.EnumContestItem) string { return fmt.Sprintf("%-24s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForContinentEnumItem(a []*continent.EnumContinentItem) string {
	return generateGoCode(
		a,
		"continent",
		"",
		"Continent",
		nil,
		func(a *continent.EnumContinentItem) string { return string(a.ID) },
		func(a *continent.EnumContinentItem) string { return string(a.ID) },
		func(a *continent.EnumContinentItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForCreditEnumItem(a []*credit.EnumCreditItem) string {
	return generateGoCode(
		a,
		"credit",
		"",
		"Credit",
		nil,
		func(a *credit.EnumCreditItem) string { return string(a.ID) },
		func(a *credit.EnumCreditItem) string { return string(a.ID) },
		func(a *credit.EnumCreditItem) string {
			return fmt.Sprintf("%-20s %-10s %-15s %s", a.ID, a.Facet, a.Sponsor, a.Award)
		},
	)
}

func GenerateGoCodeForDXCCEnumItem(a []*dxccentitycode.EnumDXCCEntityCodeItem) string {
	a[0].Description = "" // A hack... we want "none" to be blank so that it shows up nicely in drop down lists.

	return generateGoCode(
		a,
		"dxccentitycode",
		"",
		"DXCCEntityCode",
		nil,
		func(a *dxccentitycode.EnumDXCCEntityCodeItem) string {
			name := a.Description

			if name == "" {
				// See also, the hack at the top of this function.
				name = "NONE"
			}

			if bool(a.IsDeleted) {
				name += "_DELETED"
			}

			return name
		},
		func(a *dxccentitycode.EnumDXCCEntityCodeItem) int { return int(a.ID) },
		func(a *dxccentitycode.EnumDXCCEntityCodeItem) string {
			return fmt.Sprintf("%-4d = %s", a.ID, a.Description)
		},
	)
}

func GenerateGoCodeForModeEnumItem(a []*mode.EnumModeItem) string {
	return generateGoCode(
		a,
		"mode",
		"Mode",
		"Mode",
		nil,
		func(a *mode.EnumModeItem) string { return string(a.ID) },
		func(a *mode.EnumModeItem) string { return string(a.ID) },
		func(a *mode.EnumModeItem) string {
			if len(a.Submodes.Submodes) > 0 {
				submodeListAsString := make([]string, len(a.Submodes.Submodes))
				for i, subMode := range a.Submodes.Submodes {
					submodeListAsString[i] = string(subMode)
				}
				submodesStr := strings.Join(submodeListAsString, ", ")
				return fmt.Sprintf("%-15s = [ %s ]", a.ID, submodesStr)
			}
			return string(a.ID)
		},
	)
}

func GenerateGoCodeForMorseKeyEnumItem(a []*morsekey.EnumMorseKeyItem) string {
	return generateGoCode(
		a,
		"morsekey",
		"",
		"MorseKey",
		nil,
		func(a *morsekey.EnumMorseKeyItem) string { return string(a.ID) },
		func(a *morsekey.EnumMorseKeyItem) string { return string(a.ID) },
		func(a *morsekey.EnumMorseKeyItem) string { return fmt.Sprintf("%-3s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForPrimaryAdministrativeSubdivisionEnumItem(a []*primaryadministrativesubdivision.EnumPrimaryAdministrativeSubdivisionItem) string {
	return generateGoCode(
		a,
		"primaryadministrativesubdivision",
		"",
		"PrimaryAdministrativeSubdivision",
		nil,
		func(a *primaryadministrativesubdivision.EnumPrimaryAdministrativeSubdivisionItem) string {
			return string(a.ID)
		},
		func(a *primaryadministrativesubdivision.EnumPrimaryAdministrativeSubdivisionItem) string {
			return string(a.ID)
		},
		func(a *primaryadministrativesubdivision.EnumPrimaryAdministrativeSubdivisionItem) string {
			return fmt.Sprintf("%s = %s", a.ID, a.Description)
		},
	)
}

func GenerateGoCodeForPropagationModeEnumItem(a []*propagationmode.EnumPropagationModeItem) string {
	return generateGoCode(
		a,
		"propagationmode",
		"",
		"PropagationMode",
		nil,
		func(a *propagationmode.EnumPropagationModeItem) string { return string(a.ID) },
		func(a *propagationmode.EnumPropagationModeItem) string { return string(a.ID) },
		func(a *propagationmode.EnumPropagationModeItem) string {
			return fmt.Sprintf("%-10s = %s", a.ID, a.Description)
		},
	)
}

func GenerateGoCodeForQSLMediumEnumItem(a []*qslmedium.EnumQSLMediumItem) string {
	return generateGoCode(
		a,
		"qslmedium",
		"",
		"QSLMedium",
		nil,
		func(a *qslmedium.EnumQSLMediumItem) string { return string(a.ID) },
		func(a *qslmedium.EnumQSLMediumItem) string { return string(a.ID) },
		func(a *qslmedium.EnumQSLMediumItem) string { return fmt.Sprintf("%-5s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSLRcvdEnumItem(a []*qslrcvd.EnumQSLRcvdItem) string {
	return generateGoCode(
		a,
		"qslrcvd",
		"",
		"QSLRcvd",
		nil,
		func(a *qslrcvd.EnumQSLRcvdItem) string { return string(a.ID) },
		func(a *qslrcvd.EnumQSLRcvdItem) string { return string(a.ID) },
		func(a *qslrcvd.EnumQSLRcvdItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSLSentEnumItem(a []*qslsent.EnumQSLSentItem) string {
	return generateGoCode(
		a,
		"qslsent",
		"",
		"QSLSent",
		nil,
		func(a *qslsent.EnumQSLSentItem) string { return string(a.ID) },
		func(a *qslsent.EnumQSLSentItem) string { return string(a.ID) },
		func(a *qslsent.EnumQSLSentItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSLViaEnumItem(a []*qslvia.EnumQSLViaItem) string {
	return generateGoCode(
		a,
		"qslvia",
		"",
		"QSLVia",
		nil,
		func(a *qslvia.EnumQSLViaItem) string { return string(a.ID) },
		func(a *qslvia.EnumQSLViaItem) string { return string(a.ID) },
		func(a *qslvia.EnumQSLViaItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSOCompleteEnumItem(a []*qsocomplete.EnumQSOCompleteItem) string {
	return generateGoCode(
		a,
		"qsocomplete",
		"",
		"QSOComplete",
		nil,
		func(a *qsocomplete.EnumQSOCompleteItem) string { return string(a.ID) },
		func(a *qsocomplete.EnumQSOCompleteItem) string { return string(a.ID) },
		func(a *qsocomplete.EnumQSOCompleteItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSODownloadStatusEnumItem(a []*qsodownloadstatus.EnumQSODownloadStatusItem) string {
	return generateGoCode(
		a,
		"qsodownloadstatus",
		"",
		"QSODownloadStatus",
		nil,
		func(a *qsodownloadstatus.EnumQSODownloadStatusItem) string { return string(a.ID) },
		func(a *qsodownloadstatus.EnumQSODownloadStatusItem) string { return string(a.ID) },
		func(a *qsodownloadstatus.EnumQSODownloadStatusItem) string {
			return fmt.Sprintf("%s = %s", a.ID, a.Description)
		},
	)
}

func GenerateGoCodeForQSOUploadStatusEnumItem(a []*qsouploadstatus.EnumQSOUploadStatusItem) string {
	return generateGoCode(
		a,
		"qsouploadstatus",
		"",
		"QSOUploadStatus",
		nil,
		func(a *qsouploadstatus.EnumQSOUploadStatusItem) string { return string(a.ID) },
		func(a *qsouploadstatus.EnumQSOUploadStatusItem) string { return string(a.ID) },
		func(a *qsouploadstatus.EnumQSOUploadStatusItem) string {
			return fmt.Sprintf("%s = %s", a.ID, a.Description)
		},
	)
}

func GenerateGoCodeForRegionEnumItem(a []*region.EnumRegionItem) string {
	return generateGoCode(
		a,
		"region",
		"",
		"Region",
		nil,
		func(a *region.EnumRegionItem) string { return string(a.ID) },
		func(a *region.EnumRegionItem) string { return string(a.ID) },
		func(a *region.EnumRegionItem) string {
			return fmt.Sprintf("%-5s = DXCC: %-3d / %-5s / %s", a.ID, a.DXCCCode, a.Prefix, a.Description)
		},
	)
}

func GenerateGoCodeForSecondaryAdministrativeSubdivisionAltEnumItem(a []*secondaryadministrativesubdivisionalt.EnumSecondaryAdministrativeSubdivisionAltItem) string {
	return generateGoCode(
		a,
		"secondaryadministrativesubdivisionalt",
		"",
		"SecondaryAdministrativeSubdivisionAlt",
		nil,
		func(a *secondaryadministrativesubdivisionalt.EnumSecondaryAdministrativeSubdivisionAltItem) string {
			return string(a.ID)
		},
		func(a *secondaryadministrativesubdivisionalt.EnumSecondaryAdministrativeSubdivisionAltItem) string {
			return string(a.ID)
		},
		func(a *secondaryadministrativesubdivisionalt.EnumSecondaryAdministrativeSubdivisionAltItem) string {
			return fmt.Sprintf("%-50s = DXCC: %-3d / %-20s / %s", a.ID, a.DXCCEntityCode, a.Region, a.District)
		},
	)
}

func GenerateGoCodeForSecondaryAdministrativeSubdivisionEnumItem(a []*secondaryadministrativesubdivision.EnumSecondaryAdministrativeSubdivisionItem) string {
	return generateGoCode(
		a,
		"secondaryadministrativesubdivision",
		"",
		"SecondaryAdministrativeSubdivision",
		nil,
		func(a *secondaryadministrativesubdivision.EnumSecondaryAdministrativeSubdivisionItem) string {
			return string(a.ID)
		},
		func(a *secondaryadministrativesubdivision.EnumSecondaryAdministrativeSubdivisionItem) string {
			return string(a.ID)
		},
		func(a *secondaryadministrativesubdivision.EnumSecondaryAdministrativeSubdivisionItem) string {
			return fmt.Sprintf("%-40s = DXCC: %-3d / %-35s / %s", a.ID, a.DXCCEntityCode, a.SecondaryAdministrativeSubdivision, a.AlaskaJudicialDistrict)
		},
	)
}

func GenerateGoCodeForSubModeEnumItem(a []*mode.EnumSubModeItem) string {
	return generateGoCode(
		a,
		"mode",
		"SubMode",
		"SubMode",
		nil,
		func(a *mode.EnumSubModeItem) string { return string(a.ID) },
		func(a *mode.EnumSubModeItem) string { return string(a.ID) },
		func(a *mode.EnumSubModeItem) string { return fmt.Sprintf("%-12s / %s", a.Mode, a.ID) },
	)
}

func generateGoCode[T, K any](
	a []K,
	packageName string,
	constPrefix string,
	dataType string,
	filterFunc func(K) bool,
	keyFieldNameFunc func(K) string,
	keyFieldValueFunc func(K) T,
	commentFunc func(K) string) string {

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("package %s\n", strings.ToLower(packageName)))

	// Generate the map lookup constants
	builder.WriteString("var (\n")
	for _, value := range a {
		if filterFunc == nil || filterFunc(value) {
			builder.WriteString(fmt.Sprintf("%s%s %s = %#v // %s\n",
				constPrefix, convertToGoIdentifier(keyFieldNameFunc(value)), dataType, keyFieldValueFunc(value), commentFunc(value)))
		}
	}
	builder.WriteString(")\n")
	return builder.String()
}

func convertToGoIdentifier(s string) string {
	var replacer = strings.NewReplacer(
		" ", "_",
		".", "_",
		"&", "AND",
		"(", "",
		")", "",
		",", "",
		"-", "_",
		"/", "_",
		"'", "",
		":", "_",
		"?", "Uncertain", // Special case for QSOComplete '?'
	)
	result := replacer.Replace(s)
	result = strings.ReplaceAll(result, "__", "_")
	result = strings.TrimSuffix(result, "_")
	return result
}
