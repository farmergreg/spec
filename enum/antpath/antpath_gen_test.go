package antpath

import (
	"sync"
	"testing"
)

func TestLookup_Success(t *testing.T) {
	tests := []struct {
		name string
		arg  AntPath
		want Spec
	}{
		{name: "g", arg: G, want: *lookupMap[G]},
		{name: "l", arg: L, want: *lookupMap[L]},
		{name: "o", arg: O, want: *lookupMap[O]},
		{name: "s", arg: S, want: *lookupMap[S]},
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
	got, ok := Lookup("ZZ")
	if ok {
		t.Errorf("Lookup(ZZ) = %v, want nil", got)
	}
}

func TestLookupByFilter_Success(t *testing.T) {
	got := LookupByFilter(func(spec Spec) bool { return spec.Key == G })
	if len(got) != 1 || got[0].Key != G {
		t.Errorf("LookupByFilter(%v) = %v, want %v", G, got, 1)
	}
}

func TestLookupByFilter_NoResults(t *testing.T) {
	got := LookupByFilter(func(spec Spec) bool { return spec.Key == "ZZ" })
	if len(got) != 0 {
		t.Errorf("LookupByFilter(%v) = %v, want %v", G, got, 0)
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

func TestLookup_Immutability(t *testing.T) {
	spec, ok := Lookup(G)
	if !ok {
		t.Errorf("lookupMap[G] not found")
	}

	spec.Key = "Z"
	_ = spec.Key // we want to see if this newly assigned value affects the next lookup

	spec2, ok := Lookup(G)
	if !ok {
		t.Errorf("lookupMap[G] not found")
	}
	if spec2.Key == "Z" {
		t.Errorf("lookupMap is not immutable")
	}
}

func TestLookupByFilter_Immutability(t *testing.T) {
	filterFunc := func(spec Spec) bool { return spec.Key == G }
	got := LookupByFilter(filterFunc)
	if len(got) != 1 {
		t.Errorf("LookupByFilter(%v) = %v, want %v", G, got, 1)
	}
	got[0].Key = "Z"
	got2 := LookupByFilter(filterFunc)
	if len(got2) != 1 {
		t.Errorf("LookupByFilter(%v) = %v, want %v", G, got2, 1)
	}
	if got2[0].Key == "Z" {
		t.Errorf("LookupByFilter is not immutable")
	}
}

func TestListImmutability(t *testing.T) {
	list1 := List()
	list1[0].Key = "ZZ"
	list2 := List()
	if list2[0].Key == "ZZ" {
		t.Errorf("List() is not immutable")
	}
}

func TestListActiveImmutability(t *testing.T) {
	list1 := ListActive()
	list1[0].Key = "ZZ"
	list2 := ListActive()
	if list2[0].Key == "ZZ" {
		t.Errorf("ListActive() is not immutable")
	}
}

func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	for range 10000 {
		wg.Add(3)
		go func() {
			defer wg.Done()
			List()
		}()
		go func() {
			defer wg.Done()
			ListActive()
		}()
		go func() {
			defer wg.Done()
			Lookup(G)
		}()
	}
	wg.Wait()
}
