package antpath

import (
	"testing"
)

func TestLookup_Success(t *testing.T) {
	tests := []struct {
		name string
		arg  AntPath
		want Spec
	}{
		{name: "G", arg: G, want: *lookupMap[G]},
		{name: "L", arg: L, want: *lookupMap[L]},
		{name: "O", arg: O, want: *lookupMap[O]},
		{name: "S", arg: S, want: *lookupMap[S]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Lookup(tt.arg)
			if !ok {
				t.Errorf("Lookup(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestLookup_Failure(t *testing.T) {
	got, ok := Lookup("Z")
	if ok {
		t.Errorf("Lookup(Z) = %v, want nil", got)
	}
}

func TestLookupByFilter_Success(t *testing.T) {
	got := LookupByFilter(func(spec Spec) bool { return spec.Key == "G" })
	if len(got) != 1 || got[0].Key != "G" {
		t.Errorf("LookupByFilter(%v) = %v, want %v", "G", got, 1)
	}
}

func TestLookupByFilter_Failure(t *testing.T) {
	got := LookupByFilter(func(spec Spec) bool { return spec.Key == "A" })
	if len(got) != 0 {
		t.Errorf("LookupByFilter(%v) = %v, want %v", "A", got, 0)
	}
}

func TestListActive(t *testing.T) {
	got := ListActive()
	if len(got) != 4 {
		t.Errorf("ListActive() = %v, want %v", got, 4)
	}
}

func TestList(t *testing.T) {
	got := List()
	if len(got) != 4 {
		t.Errorf("List() = %v, want %v", got, 4)
	}
}
