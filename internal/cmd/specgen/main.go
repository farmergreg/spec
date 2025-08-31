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

	"github.com/hamradiolog-net/adif-spec/v8/spec"
	"github.com/hamradiolog-net/adif-spec/v8/specdata"
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
}

func main() {
	adifSpec := specdata.GetADIFSpec()

	fmt.Printf("ADIF Version: %s\n", adifSpec.Version)
	fmt.Printf("Status: %s\n", adifSpec.Status)

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "",
		PackageName: "spec",
		Records:     nil,
		ConstPrefix: "",
	}, false, "spec.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "ADIField",
		PackageName: "adifield",
		Records:     adifSpec.Fields.Records,
		ConstPrefix: "",
	}, false, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "ADIType",
		PackageName: "aditype",
		Records:     adifSpec.DataTypes.Records,
		ConstPrefix: "",
	}, false, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "AntPath",
		PackageName: "antpath",
		Records:     adifSpec.Enum.Ant_Path.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "ARRLSection",
		PackageName: "arrlsection",
		Records:     adifSpec.Enum.ARRL_Section.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Award",
		PackageName: "award",
		Records:     adifSpec.Enum.Award.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "AwardSponsorPrefix",
		PackageName: "awardsponsor",
		Records:     adifSpec.Enum.Award_Sponsor.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Band",
		PackageName: "band",
		Records:     adifSpec.Enum.Band.Records,
		ConstPrefix: "Band",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Contest",
		PackageName: "contest",
		Records:     adifSpec.Enum.Contest_ID.Records,
		ConstPrefix: "Contest",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Continent",
		PackageName: "continent",
		Records:     adifSpec.Enum.Continent.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Credit",
		PackageName: "credit",
		Records:     adifSpec.Enum.Credit.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "DXCCEntityCode",
		PackageName: "dxccentitycode",
		Records:     adifSpec.Enum.DXCC_Entity_Code.Records,
		ConstPrefix: "",
	}, true, "dxcc.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "EQSLAG",
		PackageName: "eqslag",
		Records:     adifSpec.Enum.EQSL_AG.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "Mode",
		PackageName: "mode",
		Records:     adifSpec.Enum.Mode.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "MorseKeyType",
		PackageName: "morsekeytype",
		Records:     adifSpec.Enum.Morse_Key_Type.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "PropagationMode",
		PackageName: "propagationmode",
		Records:     adifSpec.Enum.Propagation_Mode.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSLMedium",
		PackageName: "qslmedium",
		Records:     adifSpec.Enum.QSL_Medium.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSLRcvd",
		PackageName: "qslrcvd",
		Records:     adifSpec.Enum.QSL_Rcvd.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSLSent",
		PackageName: "qslsent",
		Records:     adifSpec.Enum.QSL_Sent.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSLVia",
		PackageName: "qslvia",
		Records:     adifSpec.Enum.QSL_Via.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSOComplete",
		PackageName: "qsocomplete",
		Records:     adifSpec.Enum.QSO_Complete.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSODownloadStatus",
		PackageName: "qsodownloadstatus",
		Records:     adifSpec.Enum.QSO_Download_Status.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "QSOUploadStatus",
		PackageName: "qsouploadstatus",
		Records:     adifSpec.Enum.QSO_Upload_Status.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "SecondaryAdministrativeSubdivision",
		PackageName: "secondaryadministrativesubdivision",
		Records:     adifSpec.Enum.Secondary_Administrative_Subdivision.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "SecondaryAdministrativeSubdivisionAlt",
		PackageName: "secondaryadministrativesubdivisionalt",
		Records:     adifSpec.Enum.Secondary_Administrative_Subdivision_Alt.Records,
		ConstPrefix: "",
	}, true, "standard.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "SubMode",
		PackageName: "submode",
		Records:     adifSpec.Enum.Submode.Records,
		ConstPrefix: "SubMode",
	}, true, "standard.tmpl")

	// Region and PrimaryAdministrativeSubdivision have composite keys and are quite different from the rest.
	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "PrimaryAdministrativeSubdivisionCompositeKey",
		PackageName: "primaryadministrativesubdivision",
		Records:     adifSpec.Enum.Primary_Administrative_Subdivision.Records,
		ConstPrefix: "PrimaryAdministrativeSubdivision",
	}, true, "standard-composite-index.tmpl")

	generate(ViewBag{
		Spec:        adifSpec,
		DataType:    "RegionCompositeKey",
		PackageName: "region",
		Records:     adifSpec.Enum.Region.Records,
		ConstPrefix: "",
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
