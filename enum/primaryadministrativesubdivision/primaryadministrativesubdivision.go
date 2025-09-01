package primaryadministrativesubdivision

import "github.com/hamradiolog-net/adif-spec/v6/enum/dxccentitycode"

type PrimaryAdministrativeSubdivisionCompositeKey string

func (p PrimaryAdministrativeSubdivisionCompositeKey) String() string {
	return string(p)
}

// LookupRegion looks up a Region specification by its composite key (RegionCompositeKey + DXCCEntityCode).
func LookupPrimaryAdministrativeSubdivision(code PrimaryAdministrativeSubdivisionCode, dxccEntityCode dxccentitycode.DXCCEntityCode) (Spec, bool) {
	spec, ok := PrimaryAdministrativeSubdivisionCompositeKeyMap[PrimaryAdministrativeSubdivisionCompositeKey(string(code)+"."+dxccEntityCode.String())]
	return spec, ok
}
