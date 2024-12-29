package main

import (
	"fmt"
	"strings"

	"github.com/hamradiolog-net/spec"
)

func GenerateGoCodeForDataTypeDefinition(a []*spec.DataTypeDefinition) string {
	return generateGoCode(
		a,
		"aditype",
		"",
		"DataType",
		nil,
		func(a *spec.DataTypeDefinition) string { return string(a.ID) },
		func(a *spec.DataTypeDefinition) string { return string(a.ID) },
		func(a *spec.DataTypeDefinition) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForFieldDefinition(a []*spec.FieldDefinition) string {
	return generateGoCode(
		a,
		"adifield",
		"",
		"Field",
		func(a *spec.FieldDefinition) bool {
			switch a.ID {
			case "USERDEFn":
				// USERDEFn is just an example field that appears in our source data.
				// Removing it because it is not real.
				return false
			default:
				return true
			}
		},
		func(a *spec.FieldDefinition) string { return string(a.ID) },
		func(a *spec.FieldDefinition) string { return string(a.ID) },
		func(a *spec.FieldDefinition) string {

			header := "Record"
			if a.IsHeaderField {
				header = "Header"
			}
			return fmt.Sprintf("%-30s = %s: %s", a.ID, header, a.Description)
		},
	)
}

func GenerateGoCodeForAntPathEnumItem(a []*spec.EnumAntPathItem) string {
	return generateGoCode(
		a,
		"antpath",
		"",
		"AntPath",
		nil,
		func(a *spec.EnumAntPathItem) string { return string(a.ID) },
		func(a *spec.EnumAntPathItem) string { return string(a.ID) },
		func(a *spec.EnumAntPathItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForARRLSectionEnumItem(a []*spec.EnumARRLSectionItem) string {
	return generateGoCode(
		a,
		"arrlsection",
		"",
		"ARRLSection",
		nil,
		func(a *spec.EnumARRLSectionItem) string { return string(a.ID) },
		func(a *spec.EnumARRLSectionItem) string { return string(a.ID) },
		func(a *spec.EnumARRLSectionItem) string { return fmt.Sprintf("%-5s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForAwardSponsorEnumItem(a []*spec.EnumAwardSponsorItem) string {
	return generateGoCode(
		a,
		"awardsponsorprefix",
		"",
		"AwardSponsorPrefix",
		nil,
		func(a *spec.EnumAwardSponsorItem) string { return string(a.IDPrefix) },
		func(a *spec.EnumAwardSponsorItem) string { return string(a.IDPrefix) },
		func(a *spec.EnumAwardSponsorItem) string {
			return fmt.Sprintf("%-7s = %s", a.IDPrefix, a.Description)
		},
	)
}

func GenerateGoCodeForAwardEnumItem(a []*spec.EnumAwardItem) string {
	return generateGoCode(
		a,
		"award",
		"",
		"Award",
		nil,
		func(a *spec.EnumAwardItem) string { return string(a.ID) },
		func(a *spec.EnumAwardItem) string { return string(a.ID) },
		func(a *spec.EnumAwardItem) string { return string(a.ID) },
	)
}

func GenerateGoCodeForBandEnumItem(a []*spec.EnumBandItem) string {
	return generateGoCode(
		a,
		"band",
		"Band",
		"Band",
		nil,
		func(a *spec.EnumBandItem) string { return string(a.ID) },
		func(a *spec.EnumBandItem) string { return string(a.ID) },
		func(a *spec.EnumBandItem) string {
			return fmt.Sprintf("%12.4f MHz - %12.4f MHz = %s", a.LowerFreq, a.UpperFreq, a.ID)
		},
	)
}

func GenerateGoCodeForContestEnumItem(a []*spec.EnumContestItem) string {
	return generateGoCode(
		a,
		"contest",
		"Contest_",
		"Contest",
		nil,
		func(a *spec.EnumContestItem) string { return string(a.ID) },
		func(a *spec.EnumContestItem) string { return string(a.ID) },
		func(a *spec.EnumContestItem) string { return fmt.Sprintf("%-24s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForContinentEnumItem(a []*spec.EnumContinentItem) string {
	return generateGoCode(
		a,
		"continent",
		"",
		"Continent",
		nil,
		func(a *spec.EnumContinentItem) string { return string(a.ID) },
		func(a *spec.EnumContinentItem) string { return string(a.ID) },
		func(a *spec.EnumContinentItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForCreditEnumItem(a []*spec.EnumCreditItem) string {
	return generateGoCode(
		a,
		"credit",
		"",
		"Credit",
		nil,
		func(a *spec.EnumCreditItem) string { return string(a.ID) },
		func(a *spec.EnumCreditItem) string { return string(a.ID) },
		func(a *spec.EnumCreditItem) string {
			return fmt.Sprintf("%-20s %-10s %-15s %s", a.ID, a.Facet, a.Sponsor, a.Award)
		},
	)
}

func GenerateGoCodeForDXCCEnumItem(a []*spec.EnumDXCCEntityCodeItem) string {
	a[0].Description = "" // A hack... we want "none" to be blank so that it shows up nicely in drop down lists.

	return generateGoCode(
		a,
		"dxccentitycode",
		"",
		"DXCC",
		nil,
		func(a *spec.EnumDXCCEntityCodeItem) string {
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
		func(a *spec.EnumDXCCEntityCodeItem) int { return int(a.ID) },
		func(a *spec.EnumDXCCEntityCodeItem) string { return fmt.Sprintf("%-4d = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForModeEnumItem(a []*spec.EnumModeItem) string {
	return generateGoCode(
		a,
		"mode",
		"Mode",
		"Mode",
		nil,
		func(a *spec.EnumModeItem) string { return string(a.ID) },
		func(a *spec.EnumModeItem) string { return string(a.ID) },
		func(a *spec.EnumModeItem) string {
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

func GenerateGoCodeForMorseKeyEnumItem(a []*spec.EnumMorseKeyItem) string {
	return generateGoCode(
		a,
		"morsekey",
		"",
		"MorseKey",
		nil,
		func(a *spec.EnumMorseKeyItem) string { return string(a.ID) },
		func(a *spec.EnumMorseKeyItem) string { return string(a.ID) },
		func(a *spec.EnumMorseKeyItem) string { return fmt.Sprintf("%-3s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForPrimaryAdministrativeSubdivisionEnumItem(a []*spec.EnumPrimaryAdministrativeSubdivisionItem) string {
	return generateGoCode(
		a,
		"primaryadministrativesubdivision",
		"",
		"PrimaryAdministrativeSubdivision",
		nil,
		func(a *spec.EnumPrimaryAdministrativeSubdivisionItem) string { return string(a.ID) },
		func(a *spec.EnumPrimaryAdministrativeSubdivisionItem) string { return string(a.ID) },
		func(a *spec.EnumPrimaryAdministrativeSubdivisionItem) string {
			return fmt.Sprintf("%s = %s", a.ID, a.Description)
		},
	)
}

func GenerateGoCodeForPropagationModeEnumItem(a []*spec.EnumPropagationModeItem) string {
	return generateGoCode(
		a,
		"propagationmode",
		"",
		"PropagationMode",
		nil,
		func(a *spec.EnumPropagationModeItem) string { return string(a.ID) },
		func(a *spec.EnumPropagationModeItem) string { return string(a.ID) },
		func(a *spec.EnumPropagationModeItem) string { return fmt.Sprintf("%-10s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSLMediumEnumItem(a []*spec.EnumQSLMediumItem) string {
	return generateGoCode(
		a,
		"qslmedium",
		"",
		"QSLMedium",
		nil,
		func(a *spec.EnumQSLMediumItem) string { return string(a.ID) },
		func(a *spec.EnumQSLMediumItem) string { return string(a.ID) },
		func(a *spec.EnumQSLMediumItem) string { return fmt.Sprintf("%-5s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSLRcvdEnumItem(a []*spec.EnumQSLRcvdItem) string {
	return generateGoCode(
		a,
		"qslrcvd",
		"",
		"QSLRcvd",
		nil,
		func(a *spec.EnumQSLRcvdItem) string { return string(a.ID) },
		func(a *spec.EnumQSLRcvdItem) string { return string(a.ID) },
		func(a *spec.EnumQSLRcvdItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSLSentEnumItem(a []*spec.EnumQSLSentItem) string {
	return generateGoCode(
		a,
		"qslsent",
		"",
		"QSLSent",
		nil,
		func(a *spec.EnumQSLSentItem) string { return string(a.ID) },
		func(a *spec.EnumQSLSentItem) string { return string(a.ID) },
		func(a *spec.EnumQSLSentItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSLViaEnumItem(a []*spec.EnumQSLViaItem) string {
	return generateGoCode(
		a,
		"qslvia",
		"",
		"QSLVia",
		nil,
		func(a *spec.EnumQSLViaItem) string { return string(a.ID) },
		func(a *spec.EnumQSLViaItem) string { return string(a.ID) },
		func(a *spec.EnumQSLViaItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSOCompleteEnumItem(a []*spec.EnumQSOCompleteItem) string {
	return generateGoCode(
		a,
		"qsocomplete",
		"",
		"QSOComplete",
		nil,
		func(a *spec.EnumQSOCompleteItem) string { return string(a.ID) },
		func(a *spec.EnumQSOCompleteItem) string { return string(a.ID) },
		func(a *spec.EnumQSOCompleteItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSODownloadStatusEnumItem(a []*spec.EnumQSODownloadStatusItem) string {
	return generateGoCode(
		a,
		"qsodownloadstatus",
		"",
		"QSODownloadStatus",
		nil,
		func(a *spec.EnumQSODownloadStatusItem) string { return string(a.ID) },
		func(a *spec.EnumQSODownloadStatusItem) string { return string(a.ID) },
		func(a *spec.EnumQSODownloadStatusItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForQSOUploadStatusEnumItem(a []*spec.EnumQSOUploadStatusItem) string {
	return generateGoCode(
		a,
		"qsouploadstatus",
		"",
		"QSOUploadStatus",
		nil,
		func(a *spec.EnumQSOUploadStatusItem) string { return string(a.ID) },
		func(a *spec.EnumQSOUploadStatusItem) string { return string(a.ID) },
		func(a *spec.EnumQSOUploadStatusItem) string { return fmt.Sprintf("%s = %s", a.ID, a.Description) },
	)
}

func GenerateGoCodeForRegionEnumItem(a []*spec.EnumRegionItem) string {
	return generateGoCode(
		a,
		"region",
		"",
		"Region",
		nil,
		func(a *spec.EnumRegionItem) string { return string(a.ID) },
		func(a *spec.EnumRegionItem) string { return string(a.ID) },
		func(a *spec.EnumRegionItem) string {
			return fmt.Sprintf("%-5s = DXCC: %-3d / %-5s / %s", a.ID, a.DXCCCode, a.Prefix, a.Description)
		},
	)
}

func GenerateGoCodeForSecondaryAdministrativeSubdivisionAltEnumItem(a []*spec.EnumSecondaryAdministrativeSubdivisionAltItem) string {
	return generateGoCode(
		a,
		"secondaryadministrativesubdivisionalt",
		"",
		"SecondaryAdministrativeSubdivisionAlt",
		nil,
		func(a *spec.EnumSecondaryAdministrativeSubdivisionAltItem) string { return string(a.ID) },
		func(a *spec.EnumSecondaryAdministrativeSubdivisionAltItem) string { return string(a.ID) },
		func(a *spec.EnumSecondaryAdministrativeSubdivisionAltItem) string {
			return fmt.Sprintf("%-50s = DXCC: %-3d / %-20s / %s", a.ID, a.DXCCEntityCode, a.Region, a.District)
		},
	)
}

func GenerateGoCodeForSecondaryAdministrativeSubdivisionEnumItem(a []*spec.EnumSecondaryAdministrativeSubdivisionItem) string {
	return generateGoCode(
		a,
		"secondaryadministrativesubdivision",
		"",
		"SecondaryAdministrativeSubdivision",
		nil,
		func(a *spec.EnumSecondaryAdministrativeSubdivisionItem) string { return string(a.ID) },
		func(a *spec.EnumSecondaryAdministrativeSubdivisionItem) string { return string(a.ID) },
		func(a *spec.EnumSecondaryAdministrativeSubdivisionItem) string {
			return fmt.Sprintf("%-40s = DXCC: %-3d / %-35s / %s", a.ID, a.DXCCEntityCode, a.SecondaryAdministrativeSubdivision, a.AlaskaJudicialDistrict)
		},
	)
}

func GenerateGoCodeForSubModeEnumItem(a []*spec.EnumSubModeItem) string {
	return generateGoCode(
		a,
		"mode",
		"SubMode",
		"SubMode",
		nil,
		func(a *spec.EnumSubModeItem) string { return string(a.ID) },
		func(a *spec.EnumSubModeItem) string { return string(a.ID) },
		func(a *spec.EnumSubModeItem) string { return fmt.Sprintf("%-12s / %s", a.Mode, a.ID) },
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
