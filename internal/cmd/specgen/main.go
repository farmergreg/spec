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

	"github.com/hamradiolog-net/spec/v6/internal/codegen"
	"github.com/hamradiolog-net/spec/v6/spec"
	"github.com/hamradiolog-net/spec/v6/specdata"
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
	generate(false, "spec.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec})

	/*
	 * Fields and Data Types
	 */
	generate(false, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Fields})
	generate(false, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.DataTypes})

	/*
	 * Enumerations
	 */
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Ant_Path})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.ARRL_Section})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Award})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Award_Sponsor})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Band})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Contest_ID})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Continent})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Credit})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.DXCC_Entity_Code})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.EQSL_AG})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Mode})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Morse_Key_Type})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Propagation_Mode})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSL_Medium})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSL_Rcvd})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSL_Sent})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSL_Via})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSO_Complete})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSO_Download_Status})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.QSO_Upload_Status})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Secondary_Administrative_Subdivision})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Secondary_Administrative_Subdivision_Alt})
	generate(true, "standard.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Submode})

	/*
	 * Enumerations: Composite Key
	 */
	generate(true, "standard_composite_index.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Primary_Administrative_Subdivision})
	generate(true, "standard_composite_index.tmpl", ViewBag{Spec: *adifSpec, CodeGen: adifSpec.Enum.Region})
}

func generate(isEnum bool, tmplName string, viewBag ViewBag) {
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
	if isEnum {
		dir = path.Join("enum", viewBag.CodeGen.CodeGenMetadata().PackageName)
	}
	writeToFile(dir, viewBag.CodeGen.CodeGenMetadata().PackageName+"_gen.go", buf.String())
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
