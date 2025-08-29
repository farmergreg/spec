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

	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
	"github.com/hamradiolog-net/adif-spec/src/pkg/specdata"
)

//go:generate go run .

type ViewBag struct {
	Spec        spec.AdifSpec
	DataType    string
	PackageName string
	Records     any
	ConstPrefix string
}

func main() {
	adifSpec := specdata.GetADIFSpec()

	fmt.Printf("ADIF Version: %s\n", adifSpec.Adif.Version)
	fmt.Printf("Status: %s\n", adifSpec.Adif.Status)

	generate(adifSpec, nil, "spec", "spec", "", "")
	generate(adifSpec, adifSpec.Adif.Fields.Records, "field", "field", "Field", "")
	generate(adifSpec, adifSpec.Adif.DataTypes.Records, "datatype", "datatype", "DataType", "")
	generate(adifSpec, adifSpec.Adif.Enum.Ant_Path.Records, "enum", "antpath", "AntPath", "")
	generate(adifSpec, adifSpec.Adif.Enum.ARRL_Section.Records, "enum", "arrlsection", "ARRLSection", "")
	generate(adifSpec, adifSpec.Adif.Enum.Award.Records, "enum", "award", "Award", "")
	generate(adifSpec, adifSpec.Adif.Enum.Award_Sponsor.Records, "enum", "awardsponsor", "AwardSponsorPrefix", "")
	generate(adifSpec, adifSpec.Adif.Enum.Band.Records, "enum", "band", "Band", "Band")
	generate(adifSpec, adifSpec.Adif.Enum.Contest_ID.Records, "enum", "contest", "Contest", "Contest")
	generate(adifSpec, adifSpec.Adif.Enum.Continent.Records, "enum", "continent", "Continent", "")
	generate(adifSpec, adifSpec.Adif.Enum.Credit.Records, "enum", "credit", "Credit", "")
	generate(adifSpec, adifSpec.Adif.Enum.DXCC_Entity_Code.Records, "enum_dxccentitycode", "dxccentitycode", "DXCCEntityCode", "")
	generate(adifSpec, adifSpec.Adif.Enum.EQSL_AG.Records, "enum", "eqslag", "EQSLAG", "")
	generate(adifSpec, adifSpec.Adif.Enum.Mode.Records, "enum", "mode", "Mode", "")
	generate(adifSpec, adifSpec.Adif.Enum.Morse_Key_Type.Records, "enum", "morsekeytype", "MorseKeyType", "")
	generate(adifSpec, adifSpec.Adif.Enum.Primary_Administrative_Subdivision.Records, "enum", "primaryadministrativesubdivision", "PrimaryAdministrativeSubdivision", "PrimaryAdministrativeSubdivision")
	generate(adifSpec, adifSpec.Adif.Enum.Propagation_Mode.Records, "enum", "propagationmode", "PropagationMode", "")
	generate(adifSpec, adifSpec.Adif.Enum.QSL_Medium.Records, "enum", "qslmedium", "QSLMedium", "")
	generate(adifSpec, adifSpec.Adif.Enum.QSL_Rcvd.Records, "enum", "qslrcvd", "QSLRcvd", "")
	generate(adifSpec, adifSpec.Adif.Enum.QSL_Sent.Records, "enum", "qslsent", "QSLSent", "")
	generate(adifSpec, adifSpec.Adif.Enum.QSL_Via.Records, "enum", "qslvia", "QSLVia", "")
	generate(adifSpec, adifSpec.Adif.Enum.QSO_Complete.Records, "enum", "qsocomplete", "QSOComplete", "")
	generate(adifSpec, adifSpec.Adif.Enum.QSO_Download_Status.Records, "enum", "qsodownloadstatus", "QSODownloadStatus", "")
	generate(adifSpec, adifSpec.Adif.Enum.QSO_Upload_Status.Records, "enum", "qsouploadstatus", "QSOUploadStatus", "")
	generate(adifSpec, adifSpec.Adif.Enum.Region.Records, "enum", "region", "Region", "")
	generate(adifSpec, adifSpec.Adif.Enum.Secondary_Administrative_Subdivision.Records, "enum", "secondaryadministrativesubdivision", "SecondaryAdministrativeSubdivision", "")
	generate(adifSpec, adifSpec.Adif.Enum.Secondary_Administrative_Subdivision_Alt.Records, "enum", "secondaryadministrativesubdivisionalt", "SecondaryAdministrativeSubdivisionAlt", "")
	generate(adifSpec, adifSpec.Adif.Enum.Submode.Records, "enum", "submode", "SubMode", "SubMode")
}

func generate(spec spec.AdifSpec, records any, tmplName, packageName, dataType, constPrefix string) {
	fileName := filepath.Join("src/cmd/specgen/template/", tmplName+".tmpl")

	tmpl := template.New(tmplName + ".tmpl").Funcs(tmplFuncs)
	tmpl, err := tmpl.ParseFiles(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, ViewBag{
		Spec:        spec,
		DataType:    dataType,
		PackageName: packageName,
		Records:     records,
		ConstPrefix: constPrefix,
	})
	if err != nil {
		log.Fatal(err)
	}

	dir := packageName
	if strings.HasPrefix(tmplName, "enum") {
		dir = path.Join("enum", packageName)
	}
	writeToFile(dir, packageName+"_gen.go", buf.String())
}

func writeToFile(dir, filename, content string) {
	fullPath := filepath.Join("src/pkg", dir, filename)
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
