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

	"github.com/farmergreg/spec/v6/internal/codegen"
	"github.com/farmergreg/spec/v6/spec"
	"github.com/farmergreg/spec/v6/specdata"
)

//go:generate go install golang.org/x/tools/cmd/goimports@latest
//go:generate go run .
//go:generate goimports -w ../../../

type ViewBag struct {
	Spec    spec.AdifSpec
	CodeGen codegen.CodeGenContainer
}

func main() {
	adifSpec := specdata.LoadADIFSpecWithExtras()

	fmt.Printf("ADIF Version: %s\n", adifSpec.Version)
	fmt.Printf("Status: %s\n", adifSpec.Status)

	/*
	 * ADIF Overview
	 */
	generate("spec.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec})

	/*
	 * Fields and Data Types
	 */
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Fields})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.DataTypes})

	/*
	 * Enumerations
	 */
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Ant_Path})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.ARRL_Section})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Award})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Award_Sponsor})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Band})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Contest_ID})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Continent})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Credit})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.DXCC_Entity_Code})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.EQSL_AG})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Mode})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Morse_Key_Type})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Propagation_Mode})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSL_Medium})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSL_Rcvd})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSL_Sent})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSL_Via})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSO_Complete})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSO_Download_Status})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSO_Upload_Status})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Secondary_Administrative_Subdivision})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Secondary_Administrative_Subdivision_Alt})
	generate("standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Submode})

	/*
	 * Enumerations: Composite Key
	 */
	generate("standard_composite_index.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Primary_Administrative_Subdivision})
	generate("standard_composite_index.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Region})
}

func generate(tmplName string, viewBag ViewBag) {
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

	dir := viewBag.CodeGen.CodeGenMetadata().PackageName
	switch viewBag.CodeGen.CodeGenMetadata().PackageName {
	case "adifield":
		writeToFile(dir, "field_gen.go", buf.String())
	case "aditype":
		writeToFile(dir, "type_gen.go", buf.String())
	default:
		// It is an enum!
		dir = path.Join("enum", viewBag.CodeGen.CodeGenMetadata().PackageName)
		writeToFile(dir, viewBag.CodeGen.CodeGenMetadata().PackageName+"_gen.go", buf.String())
	}
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
