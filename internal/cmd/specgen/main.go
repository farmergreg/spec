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

	"github.com/hamradiolog-net/adif-spec/v6/spec"
	"github.com/hamradiolog-net/adif-spec/v6/specdata"
)

//go:generate go install golang.org/x/tools/cmd/goimports@latest
//go:generate go run .
//go:generate goimports -w ../../../

type ViewBag struct {
	PackageName string
	DataType    string
	ConstPrefix string
	Spec        spec.AdifSpec
	Records     any
}

func main() {
	adifSpec := specdata.GetADIFSpec()

	fmt.Printf("ADIF Version: %s\n", adifSpec.Version)
	fmt.Printf("Status: %s\n", adifSpec.Status)

	generate(ViewBag{
		PackageName: "spec",
		DataType:    "",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     nil,
	}, false, "spec.tmpl")

	generate(ViewBag{
		PackageName: "adifield",
		DataType:    "ADIField",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Fields.Records,
	}, false, "standard.tmpl")

	generate(ViewBag{
		PackageName: "aditype",
		DataType:    "ADIType",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.DataTypes.Records,
	}, false, "standard.tmpl")

	generate(ViewBag{
		PackageName: "antpath",
		DataType:    "AntPath",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Ant_Path.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "arrlsection",
		DataType:    "ARRLSection",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.ARRL_Section.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "award",
		DataType:    "Award",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Award.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "awardsponsor",
		DataType:    "AwardSponsorPrefix",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Award_Sponsor.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "band",
		DataType:    "Band",
		ConstPrefix: "Band",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Band.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "contest",
		DataType:    "Contest",
		ConstPrefix: "Contest",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Contest_ID.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "continent",
		DataType:    "Continent",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Continent.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "credit",
		DataType:    "Credit",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Credit.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "dxccentitycode",
		DataType:    "DXCCEntityCode",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.DXCC_Entity_Code.Records,
	}, true, "dxcc.tmpl")

	generate(ViewBag{
		PackageName: "eqslag",
		DataType:    "EQSLAG",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.EQSL_AG.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "mode",
		DataType:    "Mode",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Mode.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "morsekeytype",
		DataType:    "MorseKeyType",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Morse_Key_Type.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "propagationmode",
		DataType:    "PropagationMode",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Propagation_Mode.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "qslmedium",
		DataType:    "QSLMedium",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.QSL_Medium.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "qslrcvd",
		DataType:    "QSLRcvd",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.QSL_Rcvd.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "qslsent",
		DataType:    "QSLSent",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.QSL_Sent.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "qslvia",
		DataType:    "QSLVia",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.QSL_Via.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "qsocomplete",
		DataType:    "QSOComplete",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.QSO_Complete.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "qsodownloadstatus",
		DataType:    "QSODownloadStatus",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.QSO_Download_Status.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "qsouploadstatus",
		DataType:    "QSOUploadStatus",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.QSO_Upload_Status.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "secondaryadministrativesubdivision",
		DataType:    "SecondaryAdministrativeSubdivision",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Secondary_Administrative_Subdivision.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "secondaryadministrativesubdivisionalt",
		DataType:    "SecondaryAdministrativeSubdivisionAlt",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Secondary_Administrative_Subdivision_Alt.Records,
	}, true, "standard.tmpl")

	generate(ViewBag{
		PackageName: "submode",
		DataType:    "SubMode",
		ConstPrefix: "SubMode",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Submode.Records,
	}, true, "standard.tmpl")

	// Region and PrimaryAdministrativeSubdivision have composite keys and are quite different from the rest of the enums.
	generate(ViewBag{
		PackageName: "primaryadministrativesubdivision",
		DataType:    "PrimaryAdministrativeSubdivisionCompositeKey",
		ConstPrefix: "PrimaryAdministrativeSubdivision",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Primary_Administrative_Subdivision.Records,
	}, true, "standard-composite-index.tmpl")

	generate(ViewBag{
		PackageName: "region",
		DataType:    "RegionCompositeKey",
		ConstPrefix: "",
		Spec:        adifSpec,
		Records:     adifSpec.Enum.Region.Records,
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
