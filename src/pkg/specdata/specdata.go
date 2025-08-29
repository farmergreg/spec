package specdata

import (
	_ "embed"
	"encoding/json"
	"maps"
	"sync"

	"github.com/hamradiolog-net/adif-spec/src/pkg/field"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

// https://adif.org.uk/proposed/316/ADIF_316_resources_2025_08_27.zip
//
//go:embed all.json
var specData []byte

//go:embed fields_additional.json
var additionalFieldData []byte

var (
	adifSpec *spec.AdifSpec // adifSpec is a singleton that holds the results of parsing the ADIF Workgroup's ADIF Specification JSON
	once     sync.Once      // Thread safe singleton loading
)

// Returns a copy of the ADIF Workgroup Specification as defined in their all.json export.
func GetADIFSpec() spec.AdifSpec {
	once.Do(func() {

		// Step 1: Parse ADIF Workgroup Specification
		var container spec.AdifSpecContainer
		if err := json.Unmarshal(specData, &container); err != nil {
			panic(err)
		}

		// Step 2: Remove USERDEFn because we will add USERDEF1 - USERDEF9 in the next step.
		delete(container.AdifSpec.Fields.Records, "USERDEFn")

		// Step 3: Add Additional Fields
		var additionalFields map[field.Field]field.Spec
		if err := json.Unmarshal(additionalFieldData, &additionalFields); err != nil {
			panic(err)
		}
		maps.Copy(container.AdifSpec.Fields.Records, additionalFields)

		adifSpec = &container.AdifSpec
	})
	return *adifSpec
}
