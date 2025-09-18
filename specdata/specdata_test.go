package specdata

import (
	"testing"

	"github.com/farmergreg/spec/v6/spec"
)

func TestLoadCompleteADIFSpec(t *testing.T) {
	spec := LoadADIFSpec()
	verify(t, &spec.AdifSpec)
}

func TestLoadADIFSpecWithExtras(t *testing.T) {
	spec := LoadADIFSpecWithExtras()
	verify(t, spec)
}

func verify(t *testing.T, spec *spec.AdifSpec) {
	if len(spec.Fields.Records) < 186 {
		t.Error("fields list is too short")
	}
	if len(spec.DataTypes.Records) < 28 {
		t.Error("datatype list is too short")
	}
	if len(spec.Enum.ARRL_Section.Records) < 90 {
		t.Error("arrl section list is too short")
	}
	if len(spec.Enum.Region.Records) < 10 {
		t.Error("arrl section list is too short")
	}
}
