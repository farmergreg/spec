package specdata

import (
	_ "embed"
	"encoding/json"
	"maps"
	"strconv"
	"sync"

	"github.com/hamradiolog-net/adif-spec/v8/adifield"
	"github.com/hamradiolog-net/adif-spec/v8/spec"
)

// https://adif.org.uk/proposed/316/ADIF_316_resources_2025_08_27.zip
//
//go:embed all.json
var specData []byte

//go:embed fields_extra.json
var extraFieldData []byte

var (
	adifSpec *spec.AdifSpec // adifSpec is a singleton that holds the results of parsing the ADIF Workgroup's ADIF Specification JSON
	once     sync.Once      // Thread safe singleton loading
)

// Returns a copy of the ADIF Workgroup Specification as defined in their all.json export.
// Some modifications are made to the original data:
// 1. USERDEFn is replaced with USERDEF1, USERDEF2, ..., USERDEF9
// 2. Extra fields from fields_extra.json are added to the specification
func GetADIFSpec() spec.AdifSpec {
	once.Do(func() {

		// Step 1: Parse ADIF Workgroup Specification
		var container spec.AdifSpecContainer
		if err := json.Unmarshal(specData, &container); err != nil {
			panic(err)
		}

		// Step 2: Remove USERDEFn and replace with USERDEF1, USERDEF2, etc...
		userdefn := container.AdifSpec.Fields.Records["USERDEFn"]
		for i := 1; i < 10; i++ {
			next := userdefn
			next.Key = adifield.ADIField("USERDEF" + strconv.Itoa(i))
			container.AdifSpec.Fields.Records[next.Key] = next
		}
		delete(container.AdifSpec.Fields.Records, "USERDEFn")

		// Step 3: Add Extra Fields
		var extraFields map[adifield.ADIField]adifield.Spec
		if err := json.Unmarshal(extraFieldData, &extraFields); err != nil {
			panic(err)
		}
		maps.Copy(container.AdifSpec.Fields.Records, extraFields)

		adifSpec = &container.AdifSpec
	})
	return *adifSpec
}
