package specdata

import (
	_ "embed"
	"encoding/json"
	"log"
	"maps"
	"strconv"

	"github.com/hamradiolog-net/spec/v6/adifield"
	"github.com/hamradiolog-net/spec/v6/spec"
)

// https://adif.org.uk/proposed/316/ADIF_316_resources_2025_08_27.zip
//
//go:embed all.json
var specData []byte

//go:embed fields_extra.json
var extraFieldData []byte

// LoadADIFJson loads the ADIF Workgroup Specification from a JSON byte slice.
func LoadADIFJson(data []byte) (*spec.AdifSpecContainer, error) {
	var container spec.AdifSpecContainer
	if err := json.Unmarshal(data, &container); err != nil {
		return nil, err
	}
	return &container, nil
}

// Returns the ADIF Workgroup Specification EXACTLY as defined in the all.json export.
// This data is re-created every time this function is called.
func LoadADIFSpec() *spec.AdifSpecContainer {
	container, err := LoadADIFJson(specData)
	if err != nil {
		log.Fatal(err)
	}
	return container
}

// Returns the ADIF Workgroup Specification, defined in the all.json export, but with some modifications:
// 1. USERDEFn is replaced with USERDEF1, USERDEF2, ..., USERDEF9.
// 2. Extra fields from fields_extra.json are added to the specification.
// 3. All enum key values are converted to UPPERCASE to make using strings.ToUpper faster when comparing ADIF key values (enums are case insensitive).
// This data is re-created every time this function is called.
func LoadADIFSpecWithExtras() *spec.AdifSpec {
	// Step 1: Load ADIF Workgroup Specification
	container := LoadADIFSpec()

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
	return &container.AdifSpec
}
