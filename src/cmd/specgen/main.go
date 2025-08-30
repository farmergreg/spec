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

	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/spec"
	"github.com/hamradiolog-net/adif-spec/v2/src/pkg/specdata"
)

//go:generate go install golang.org/x/tools/cmd/goimports@latest
//go:generate go run .
//go:generate goimports -w ../../pkg

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

	generate(adifSpec, nil, false, "spec.tmpl", "spec", "", "")

	generate(adifSpec, adifSpec.Fields.Records, false, "field.tmpl", "adifield", "ADIField", "")
	generate(adifSpec, adifSpec.DataTypes.Records, false, "standard.tmpl", "aditype", "ADIType", "")

	generate(adifSpec, adifSpec.Enum.Ant_Path.Records, true, "standard.tmpl", "antpath", "AntPath", "")
	generate(adifSpec, adifSpec.Enum.ARRL_Section.Records, true, "standard.tmpl", "arrlsection", "ARRLSection", "")
	generate(adifSpec, adifSpec.Enum.Award.Records, true, "standard.tmpl", "award", "Award", "")
	generate(adifSpec, adifSpec.Enum.Award_Sponsor.Records, true, "standard.tmpl", "awardsponsor", "AwardSponsorPrefix", "")
	generate(adifSpec, adifSpec.Enum.Band.Records, true, "standard.tmpl", "band", "Band", "Band")
	generate(adifSpec, adifSpec.Enum.Contest_ID.Records, true, "standard.tmpl", "contest", "Contest", "Contest")
	generate(adifSpec, adifSpec.Enum.Continent.Records, true, "standard.tmpl", "continent", "Continent", "")
	generate(adifSpec, adifSpec.Enum.Credit.Records, true, "standard.tmpl", "credit", "Credit", "")
	generate(adifSpec, adifSpec.Enum.DXCC_Entity_Code.Records, true, "dxcc.tmpl", "dxccentitycode", "DXCCEntityCode", "")
	generate(adifSpec, adifSpec.Enum.EQSL_AG.Records, true, "standard.tmpl", "eqslag", "EQSLAG", "")
	generate(adifSpec, adifSpec.Enum.Mode.Records, true, "standard.tmpl", "mode", "Mode", "")
	generate(adifSpec, adifSpec.Enum.Morse_Key_Type.Records, true, "standard.tmpl", "morsekeytype", "MorseKeyType", "")
	// todo research primary key / Primary Administrative Subdivision
	// generate(adifSpec, adifSpec.Enum.Primary_Administrative_Subdivision.Records, true, "standard.tmpl", "primaryadministrativesubdivision", "PrimaryAdministrativeSubdivision", "PrimaryAdministrativeSubdivision")
	generate(adifSpec, adifSpec.Enum.Propagation_Mode.Records, true, "standard.tmpl", "propagationmode", "PropagationMode", "")
	generate(adifSpec, adifSpec.Enum.QSL_Medium.Records, true, "standard.tmpl", "qslmedium", "QSLMedium", "")
	generate(adifSpec, adifSpec.Enum.QSL_Rcvd.Records, true, "standard.tmpl", "qslrcvd", "QSLRcvd", "")
	generate(adifSpec, adifSpec.Enum.QSL_Sent.Records, true, "standard.tmpl", "qslsent", "QSLSent", "")
	generate(adifSpec, adifSpec.Enum.QSL_Via.Records, true, "standard.tmpl", "qslvia", "QSLVia", "")
	generate(adifSpec, adifSpec.Enum.QSO_Complete.Records, true, "standard.tmpl", "qsocomplete", "QSOComplete", "")
	generate(adifSpec, adifSpec.Enum.QSO_Download_Status.Records, true, "standard.tmpl", "qsodownloadstatus", "QSODownloadStatus", "")
	generate(adifSpec, adifSpec.Enum.QSO_Upload_Status.Records, true, "standard.tmpl", "qsouploadstatus", "QSOUploadStatus", "")
	generate(adifSpec, adifSpec.Enum.Region.Records, true, "standard.tmpl", "region", "Region", "")
	generate(adifSpec, adifSpec.Enum.Secondary_Administrative_Subdivision.Records, true, "standard.tmpl", "secondaryadministrativesubdivision", "SecondaryAdministrativeSubdivision", "")
	generate(adifSpec, adifSpec.Enum.Secondary_Administrative_Subdivision_Alt.Records, true, "standard.tmpl", "secondaryadministrativesubdivisionalt", "SecondaryAdministrativeSubdivisionAlt", "")
	generate(adifSpec, adifSpec.Enum.Submode.Records, true, "standard.tmpl", "submode", "SubMode", "SubMode")
}

func generate(spec spec.AdifSpec, records any, isEnum bool, tmplName, packageName, dataType, constPrefix string) {
	tmpl := template.New("").Funcs(tmplFuncs)
	tmpl, err := tmpl.ParseGlob("template/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	var buf strings.Builder
	err = tmpl.ExecuteTemplate(&buf, tmplName, ViewBag{
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
	if isEnum {
		dir = path.Join("enum", packageName)
	}
	writeToFile(dir, packageName+"_gen.go", buf.String())
}

func writeToFile(dir, filename, content string) {
	fullPath := filepath.Join("../../pkg", dir, filename)
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
