package main

import (
	_ "embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

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
	"github.com/hamradiolog-net/adif-spec/v6/spec"
	"github.com/hamradiolog-net/adif-spec/v6/specdata"
)

//go:generate go install golang.org/x/tools/cmd/goimports@latest
//go:generate go run .
//go:generate goimports -w ../../../

type ViewBag struct {
	Spec        spec.AdifSpec
	DataType    string
	PackageName string
	Records     any
	ConstPrefix string
	Comments    func(a any) string
}

func main() {
	adifSpec := specdata.LoadADIFSpecWithExtras()

	fmt.Printf("ADIF Version: %s\n", adifSpec.Version)
	fmt.Printf("Status: %s\n", adifSpec.Status)

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "",
		PackageName: "spec",
		Records:     nil,
		ConstPrefix: "",
		Comments:    func(a any) string { return "" },
	}, false, "spec.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "ADIField",
		PackageName: "adifield",
		Records:     adifSpec.Fields.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(adifield.Spec); ok {
				if spec.IsHeaderField {
					return fmt.Sprintf("Header: %s", spec.Description)
				} else {
					return fmt.Sprintf("Record: %s", spec.Description)
				}
			}
			return ""
		},
	}, false, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "ADIType",
		PackageName: "aditype",
		Records:     adifSpec.DataTypes.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(aditype.Spec); ok {
				return spec.Description
			}
			return ""
		},
	}, false, "standard.tmpl")

	/*
	 * Enumerations
	 */

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "AntPath",
		PackageName: "antpath",
		Records:     adifSpec.Enum.Ant_Path.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(antpath.Spec); ok {
				return fmt.Sprintf("%s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "ARRLSection",
		PackageName: "arrlsection",
		Records:     adifSpec.Enum.ARRL_Section.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(arrlsection.Spec); ok {
				return fmt.Sprintf("%-4s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Award",
		PackageName: "award",
		Records:     adifSpec.Enum.Award.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(award.Spec); ok {
				return string(spec.Key)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "AwardSponsorPrefix",
		PackageName: "awardsponsor",
		Records:     adifSpec.Enum.Award_Sponsor.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(awardsponsor.Spec); ok {
				return fmt.Sprintf("%-6s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Band",
		PackageName: "band",
		Records:     adifSpec.Enum.Band.Records,
		ConstPrefix: "Band",
		Comments: func(a any) string {
			if spec, ok := a.(band.Spec); ok {
				return fmt.Sprintf("%-6s = %12.4f MHz to %12.4f MHz", spec.Key, spec.LowerFreqMHz, spec.UpperFreqMHz)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Contest",
		PackageName: "contest",
		Records:     adifSpec.Enum.Contest_ID.Records,
		ConstPrefix: "Contest_",
		Comments: func(a any) string {
			if spec, ok := a.(contest.Spec); ok {
				return fmt.Sprintf("%-20s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Continent",
		PackageName: "continent",
		Records:     adifSpec.Enum.Continent.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(continent.Spec); ok {
				return fmt.Sprintf("%s = %s", spec.Key, spec.Continent)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Credit",
		PackageName: "credit",
		Records:     adifSpec.Enum.Credit.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(credit.Spec); ok {
				return fmt.Sprintf("%-20s = %-15s %-45s %-15s", spec.Key, spec.Sponsor, spec.Award, spec.Facet)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "DXCCEntityCode",
		PackageName: "dxccentitycode",
		Records:     adifSpec.Enum.DXCC_Entity_Code.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(dxccentitycode.Spec); ok {
				deleted := ""
				if spec.IsDeleted {
					deleted = " (DELETED) "
				}
				return fmt.Sprintf("%s = %s%s", spec.Key, spec.EntityName, deleted)
			}
			return ""
		},
	}, true, "dxcc.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "EQSLAG",
		PackageName: "eqslag",
		Records:     adifSpec.Enum.EQSL_AG.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(eqslag.Spec); ok {
				return fmt.Sprintf("%s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Mode",
		PackageName: "mode",
		Records:     adifSpec.Enum.Mode.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(mode.Spec); ok {
				if spec.IsImportOnly {
					return fmt.Sprintf("%-10s = %s", spec.Key, spec.Submodes)
				} else {
					return fmt.Sprintf("%-22s = %s", spec.Key, spec.Submodes)
				}
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "MorseKeyType",
		PackageName: "morsekeytype",
		Records:     adifSpec.Enum.Morse_Key_Type.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(morsekeytype.Spec); ok {
				return fmt.Sprintf("%-4s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "PropagationMode",
		PackageName: "propagationmode",
		Records:     adifSpec.Enum.Propagation_Mode.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(propagationmode.Spec); ok {
				return fmt.Sprintf("%-10s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSLMedium",
		PackageName: "qslmedium",
		Records:     adifSpec.Enum.QSL_Medium.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(qslmedium.Spec); ok {
				return fmt.Sprintf("%s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSLRcvd",
		PackageName: "qslrcvd",
		Records:     adifSpec.Enum.QSL_Rcvd.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(qslrcvd.Spec); ok {
				return fmt.Sprintf("%s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSLSent",
		PackageName: "qslsent",
		Records:     adifSpec.Enum.QSL_Sent.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(qslsent.Spec); ok {
				return fmt.Sprintf("%s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSLVia",
		PackageName: "qslvia",
		Records:     adifSpec.Enum.QSL_Via.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(qslvia.Spec); ok {
				return fmt.Sprintf("%s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSOComplete",
		PackageName: "qsocomplete",
		Records:     adifSpec.Enum.QSO_Complete.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(qsocomplete.Spec); ok {
				return fmt.Sprintf("%-4s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSODownloadStatus",
		PackageName: "qsodownloadstatus",
		Records:     adifSpec.Enum.QSO_Download_Status.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(qsodownloadstatus.Spec); ok {
				return fmt.Sprintf("%s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSOUploadStatus",
		PackageName: "qsouploadstatus",
		Records:     adifSpec.Enum.QSO_Upload_Status.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(qsouploadstatus.Spec); ok {
				return fmt.Sprintf("%s = %s", spec.Key, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "SecondaryAdministrativeSubdivision",
		PackageName: "secondaryadministrativesubdivision",
		Records:     adifSpec.Enum.Secondary_Administrative_Subdivision.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(secondaryadministrativesubdivision.Spec); ok {
				return fmt.Sprintf("%-35s = DXCC %s: %s", spec.Key, spec.DXCCEntityCode, spec.SecondaryAdminSub)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "SecondaryAdministrativeSubdivisionAlt",
		PackageName: "secondaryadministrativesubdivisionalt",
		Records:     adifSpec.Enum.Secondary_Administrative_Subdivision_Alt.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(secondaryadministrativesubdivisionalt.Spec); ok {
				return fmt.Sprintf("%-50s = DXCC: %s %s", spec.Key, spec.DXCCEntityCode, spec.Region)
			}
			return ""
		},
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "SubMode",
		PackageName: "submode",
		Records:     adifSpec.Enum.Submode.Records,
		ConstPrefix: "SubMode",
		Comments: func(a any) string {
			if spec, ok := a.(submode.Spec); ok {
				return fmt.Sprintf("%-15s = %-15s %s", spec.Key, spec.Mode, spec.Description)
			}
			return ""
		},
	}, true, "standard.tmpl")

	// PrimaryAdministrativeSubdivision and Region have composite keys and are quite different from the rest of the enums.
	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "PrimaryAdministrativeSubdivisionCompositeKey",
		PackageName: "primaryadministrativesubdivision",
		Records:     adifSpec.Enum.Primary_Administrative_Subdivision.Records,
		ConstPrefix: "PrimaryAdministrativeSubdivision",
		Comments: func(a any) string {
			if spec, ok := a.(primaryadministrativesubdivision.Spec); ok {
				return fmt.Sprintf("%5s.%-5s = %-5s ( %-5s ); IMPORTANT: This is NOT the Primary Administrative Subdivision Code. It is a lookup key for use with PrimaryAdministrativeSubdivisionCompositeKeyMap", spec.Code, spec.DXCCEntityCode, spec.Code, spec.PrimaryAdminSub)
			}
			return ""
		},
	}, true, "standard-composite-index.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "RegionCompositeKey",
		PackageName: "region",
		Records:     adifSpec.Enum.Region.Records,
		ConstPrefix: "",
		Comments: func(a any) string {
			if spec, ok := a.(region.Spec); ok {
				return fmt.Sprintf("%4s.%-3s = %-5s %-15s; IMPORTANT: This is NOT the Region Code. It is a lookup key for use with RegionCompositeKeyMap", spec.Code, spec.DXCCEntityCode, spec.Code, spec.Region)
			}
			return ""
		},
	}, true, "standard-composite-index.tmpl")
}

func generate(viewBag ViewBag, isEnum bool, tmplName string) {
	tmpl := template.New("").Funcs(tmplFuncs)
	tmpl, err := tmpl.ParseGlob("template/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	var buf strings.Builder
	err = tmpl.ExecuteTemplate(&buf, tmplName, viewBag)
	if err != nil {
		log.Fatal(err)
	}

	dir := viewBag.PackageName
	if isEnum {
		dir = path.Join("enum", viewBag.PackageName)
	}
	writeToFile(dir, viewBag.PackageName+"_gen.go", buf.String())
}

func writeToFile(dir, filename, content string) {
	fullPath := filepath.Join("../../../", dir, filename)
	os.MkdirAll(filepath.Dir(fullPath), 0755)

	formatted, fmtErr := format.Source([]byte(content))
	if fmtErr != nil {
		formatted = []byte(content)
	}

	err := os.WriteFile(fullPath, formatted, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if fmtErr != nil {
		log.Fatal(fmtErr)
	}
}
