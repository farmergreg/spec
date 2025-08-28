package specdata

import (
	_ "embed"
	"encoding/json"
	"sync"

	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

// https://adif.org.uk/proposed/316/ADIF_316_resources_2025_08_27.zip
//
//go:embed all.json
var jsonData []byte

var (
	adifSpec *spec.AdifSpec // adifSpec is a singleton that holds the results of parsing the ADIF Workgroup's ADIF Specification JSON
	once     sync.Once      // Thread safe singleton loading
)

// Returns a copy of the ADIF Workgroup Specification as defined in their all.json export.
func GetADIFSpec() spec.AdifSpec {
	once.Do(func() {
		if err := json.Unmarshal(jsonData, &adifSpec); err != nil {
			panic(err)
		}
	})
	return *adifSpec
}
