package specdata

import (
	_ "embed"
	"encoding/json"
	"log"
	"maps"
	"strconv"

	"github.com/hamradiolog-net/adif-spec/v6/adifield"
	"github.com/hamradiolog-net/adif-spec/v6/spec"
)

// https://adif.org.uk/proposed/316/ADIF_316_resources_2025_08_27.zip
//
//go:embed all.json
var specData []byte

//go:embed fields_extra.json
var extraFieldData []byte

// Returns a copy of the ADIF Workgroup Specification as defined in their all.json export.
// Some modifications are made to the original data:
// 1. USERDEFn is replaced with USERDEF1, USERDEF2, ..., USERDEF9
// 2. Extra fields from fields_extra.json are added to the specification
func GetADIFSpec() spec.AdifSpec {
	// Step 1: Parse ADIF Workgroup Specification
	var container spec.AdifSpecContainer
	if err := json.Unmarshal(specData, &container); err != nil {
		log.Fatal(err)
	}

	// Step 2: Remove USERDEFn and replace with USERDEF1, USERDEF2, etc...
	userdefn := container.AdifSpec.Fields.Records["USERDEFn"]
	delete(container.AdifSpec.Fields.Records, "USERDEFn")
	for i := 1; i < 10; i++ {
		next := userdefn
		next.Key = adifield.ADIField("USERDEF" + strconv.Itoa(i))
		container.AdifSpec.Fields.Records[next.Key] = next
	}

	// Step 3: Add Extra Fields
	var extraFields map[adifield.ADIField]adifield.Spec
	if err := json.Unmarshal(extraFieldData, &extraFields); err != nil {
		panic(err)
	}
	maps.Copy(container.AdifSpec.Fields.Records, extraFields)

	// Step N: All done, return...
	return container.AdifSpec
}
