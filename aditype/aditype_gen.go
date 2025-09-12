// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package aditype provides code and constants as defined in ADIF 3.1.6 (Proposed)
package aditype

import "sync"

const (
	AwardList                                 ADIType = "AwardList"                                 // Deprecated: a comma-delimited list of members of the Award enumeration
	Boolean                                   ADIType = "Boolean"                                   // if True, the single ASCII character Y or y if False, the single ASCII character N or n
	Character                                 ADIType = "Character"                                 // an ASCII character whose code lies in the range of 32 through 126, inclusive
	CreditList                                ADIType = "CreditList"                                // a comma-delimited list where each list item is either: A member of the Credit enumeration. A member of the Credit enumeration followed by a colon and an ampersand-delimited list of members of the QSL_Medium enumeration. For example IOTA,WAS:LOTW&CARD,DXCC:CARD
	Date                                      ADIType = "Date"                                      // 8 Digits representing a UTC date in YYYYMMDD format, where YYYY is a 4-Digit year specifier, where 1930 <= YYYY MM is a 2-Digit month specifier, where 1 <= MM <= 12 [use leading zeroes] DD is a 2-Digit day specifier, where 1 <= DD <= DaysInMonth(MM) [use leading zeroes]
	Digit                                     ADIType = "Digit"                                     // an ASCII character whose code lies in the range of 48 through 57, inclusive
	Enumeration                               ADIType = "Enumeration"                               // an explicit list of legal case-insensitive values represented in ASCII set forth in set notation, e.g. {A, B, C, D}, or defined in a table, from which a single value may be selected.
	GridSquare                                ADIType = "GridSquare"                                // a case-insensitive 2-character, 4-character, 6-character, or 8-character Maidenhead locator. Specific fields impose additional restrictions on the number of characters; see the field descriptions for the allowed numbers of characters.
	GridSquareExt                             ADIType = "GridSquareExt"                             // For a 10-character Maidenhead locator, contains characters 9 and 10. For a 12-character Maidenhead locator, contains characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0-9.
	GridSquareList                            ADIType = "GridSquareList"                            // a comma-delimited list of GridSquare items
	IOTARefNo                                 ADIType = "IOTARefNo"                                 // IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]
	Integer                                   ADIType = "Integer"                                   // a sequence of one or more Digits representing a decimal integer, optionally preceded by a minus sign (ASCII code 45). Leading zeroes are allowed.
	IntlCharacter                             ADIType = "IntlCharacter"                             // a Unicode character (encoded with UTF-8) excluding line break CR (code 13) and LF (code 10) characters
	IntlMultilineString                       ADIType = "IntlMultilineString"                       // a sequence of International Characters and line breaks. Fields of type IntlMultilineString must only be used in ADX files
	IntlString                                ADIType = "IntlString"                                // a sequence of International Characters. Fields of type IntlString must only be used in ADX files
	Location                                  ADIType = "Location"                                  // a sequence of 11 characters representing a latitude or longitude in XDDD MM.MMM format, where X is a directional Character from the set {E, W, N, S} DDD is a 3-Digit degrees specifier, where 0 <= DDD <= 180 [use leading zeroes] There is a single space character in between DDD and MM.MMM MM.MMM is an unsigned Number minutes specifier with its decimal point in the third position, where 00.000 <= MM.MMM <= 59.999 [use leading and trailing zeroes]
	MultilineString                           ADIType = "MultilineString"                           // a sequence of Characters and line-breaks, where a line break is an ASCII CR (code 13) followed immediately by an ASCII LF (code 10)
	Number                                    ADIType = "Number"                                    // a sequence of one or more Digits representing a decimal number, optionally preceded by a minus sign (ASCII code 45) and optionally including a single decimal point (ASCII code 46)
	POTARef                                   ADIType = "POTARef"                                   // a sequence of case-insensitive Characters representing a Parks on the Air park reference in the form xxxx-nnnnn[@yyyyyy] comprising 6 to 17 characters where: xxxx is the POTA national program and is 1 to 4 characters in length, typically the default callsign prefix of the national program (rather than the DX entity) nnnnn represents the unique number within the national program and is either 4 or 5 characters in length (use the exact format listed on the POTA website) yyyyyy **Optional** is the 4 to 6 character ISO 3166-2 code to differentiate which state/province/prefecture/primary administration location the contact represents, in the case that the park reference spans more than one location (such as a trail). Examples of the POTARef Data Type: ReferenceLocation K-5033Golden Hill State Forest K-10000 5-digit park numbers are reserved for future use VE-5082@CA-ABThe Great Trail of Canada (the Canadian Trailway) National Scenic Trail, within Alberta, Canada 8P-0012Chancery Lane Swamp National Park VK-0556Pieman River State Reserve K-4562@US-CAPacific Crest Trail, within California, USA Additional Notes on POTARef: A browsable and searchable list of all park references is available. A complete CSV file is available (generated nightly). For more information, visit the Parks on the Air documentation website.
	POTARefList                               ADIType = "POTARefList"                               // a comma-delimited list of one or more POTARef items.
	PositiveInteger                           ADIType = "PositiveInteger"                           // an unsigned sequence of one or more Digits representing a decimal integer that has a value greater than 0. Leading zeroes are allowed.
	SOTARef                                   ADIType = "SOTARef"                                   // a sequence of Characters representing an International SOTA Reference. The sequence comprises: an ITU prefix if applicable, a SOTA subdivision a / Character a SOTA Reference Number Examples: W2/WE-003 G/LD-003
	SecondaryAdministrativeSubdivisionListAlt ADIType = "SecondaryAdministrativeSubdivisionListAlt" // a semicolon (;) delimited, unordered list of one or more members of a Secondary_Administrative_Subdivision_Alt enumeration in the form: enumeration-name:enumeration-code Where there is more than one locality represented by the enumeration-code, they are separated by slash (/) characters. Only one of each enumeration-name valid for the DXCC entity concerned can appear in the list. Examples: <CNTY_ALT:28>NZ_Regions:Hawkes Bay/Wairoa <MY_CNTY_ALT:52>NZ_Islands:North Island;NZ_Regions:Hawkes Bay/Wairoa The first example shows the enumeration-name NZ_Regions with the region Hawkes Bay and the district Wairoa. For the purposes of illustration, the second example includes a non-existent subdivision with two available enumeration-codes, NZ_Islands:North Island and NZ_Islands:South Island. The example shows: the enumeration-name NZ_Islands with the island North Island the enumeration-name NZ_Regions with the region Hawkes Bay and the district Wairoa
	SecondarySubdivisionList                  ADIType = "SecondarySubdivisionList"                  // a colon-delimited list of two or more members of the Secondary_Administrative_Subdivision enumeration. E.g.: MA,Franklin:MA,Hampshire
	SponsoredAwardList                        ADIType = "SponsoredAwardList"                        // a comma-delimited list of members of the Sponsored_Award enumeration
	String                                    ADIType = "String"                                    // a sequence of Characters
	Time                                      ADIType = "Time"                                      // 6 Digits representing a UTC time in HHMMSS format or 4 Digits representing a time in HHMM format, where HH is a 2-Digit hour specifier, where 0 <= HH <= 23 [use leading zeroes] MM is a 2-Digit minute specifier, where 0 <= MM <= 59 [use leading zeroes] SS is a 2-Digit second specifier, where 0 <= SS <= 59 [use leading zeroes]
	WWFFRef                                   ADIType = "WWFFRef"                                   // a sequence of case-insensitive Characters representing an International WWFF (World Wildlife Flora & Fauna) reference in the form xxFF-nnnn comprising 8 to 11 characters where: xx is the WWFF national program and is 1 to 4 characters in length. FF- is two F characters followed by a dash character. nnnn represents the unique number within the national program and is 4 characters in length with leading zeros. Examples: KFF-4655 3DAFF-0002
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known ADIType specifications.
var lookupList = []Spec{
	{Key: "AwardList", DataTypeIndicator: "", Description: "a comma-delimited list of members of the Award enumeration", MinimumValue: 0, MaximumValue: 0, IsImportOnly: true, Comments: ""},
	{Key: "Boolean", DataTypeIndicator: "B", Description: "if True, the single ASCII character Y or y if False, the single ASCII character N or n", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "Character", DataTypeIndicator: "", Description: "an ASCII character whose code lies in the range of 32 through 126, inclusive", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CreditList", DataTypeIndicator: "", Description: "a comma-delimited list where each list item is either: A member of the Credit enumeration. A member of the Credit enumeration followed by a colon and an ampersand-delimited list of members of the QSL_Medium enumeration. For example IOTA,WAS:LOTW&CARD,DXCC:CARD", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "Date", DataTypeIndicator: "D", Description: "8 Digits representing a UTC date in YYYYMMDD format, where YYYY is a 4-Digit year specifier, where 1930 <= YYYY MM is a 2-Digit month specifier, where 1 <= MM <= 12 [use leading zeroes] DD is a 2-Digit day specifier, where 1 <= DD <= DaysInMonth(MM) [use leading zeroes]", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "Digit", DataTypeIndicator: "", Description: "an ASCII character whose code lies in the range of 48 through 57, inclusive", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "Enumeration", DataTypeIndicator: "E", Description: "an explicit list of legal case-insensitive values represented in ASCII set forth in set notation, e.g. {A, B, C, D}, or defined in a table, from which a single value may be selected.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "GridSquare", DataTypeIndicator: "", Description: "a case-insensitive 2-character, 4-character, 6-character, or 8-character Maidenhead locator. Specific fields impose additional restrictions on the number of characters; see the field descriptions for the allowed numbers of characters.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "GridSquareExt", DataTypeIndicator: "", Description: "For a 10-character Maidenhead locator, contains characters 9 and 10. For a 12-character Maidenhead locator, contains characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0-9.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "GridSquareList", DataTypeIndicator: "", Description: "a comma-delimited list of GridSquare items", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "IOTARefNo", DataTypeIndicator: "", Description: "IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "Integer", DataTypeIndicator: "", Description: "a sequence of one or more Digits representing a decimal integer, optionally preceded by a minus sign (ASCII code 45). Leading zeroes are allowed.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "IntlCharacter", DataTypeIndicator: "", Description: "a Unicode character (encoded with UTF-8) excluding line break CR (code 13) and LF (code 10) characters", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "IntlMultilineString", DataTypeIndicator: "G", Description: "a sequence of International Characters and line breaks. Fields of type IntlMultilineString must only be used in ADX files", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "IntlString", DataTypeIndicator: "I", Description: "a sequence of International Characters. Fields of type IntlString must only be used in ADX files", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "Location", DataTypeIndicator: "L", Description: "a sequence of 11 characters representing a latitude or longitude in XDDD MM.MMM format, where X is a directional Character from the set {E, W, N, S} DDD is a 3-Digit degrees specifier, where 0 <= DDD <= 180 [use leading zeroes] There is a single space character in between DDD and MM.MMM MM.MMM is an unsigned Number minutes specifier with its decimal point in the third position, where 00.000 <= MM.MMM <= 59.999 [use leading and trailing zeroes]", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MultilineString", DataTypeIndicator: "M", Description: "a sequence of Characters and line-breaks, where a line break is an ASCII CR (code 13) followed immediately by an ASCII LF (code 10)", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "Number", DataTypeIndicator: "N", Description: "a sequence of one or more Digits representing a decimal number, optionally preceded by a minus sign (ASCII code 45) and optionally including a single decimal point (ASCII code 46)", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "POTARef", DataTypeIndicator: "", Description: "a sequence of case-insensitive Characters representing a Parks on the Air park reference in the form xxxx-nnnnn[@yyyyyy] comprising 6 to 17 characters where: xxxx is the POTA national program and is 1 to 4 characters in length, typically the default callsign prefix of the national program (rather than the DX entity) nnnnn represents the unique number within the national program and is either 4 or 5 characters in length (use the exact format listed on the POTA website) yyyyyy **Optional** is the 4 to 6 character ISO 3166-2 code to differentiate which state/province/prefecture/primary administration location the contact represents, in the case that the park reference spans more than one location (such as a trail). Examples of the POTARef Data Type: ReferenceLocation K-5033Golden Hill State Forest K-10000 5-digit park numbers are reserved for future use VE-5082@CA-ABThe Great Trail of Canada (the Canadian Trailway) National Scenic Trail, within Alberta, Canada 8P-0012Chancery Lane Swamp National Park VK-0556Pieman River State Reserve K-4562@US-CAPacific Crest Trail, within California, USA Additional Notes on POTARef: A browsable and searchable list of all park references is available. A complete CSV file is available (generated nightly). For more information, visit the Parks on the Air documentation website.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "POTARefList", DataTypeIndicator: "", Description: "a comma-delimited list of one or more POTARef items.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "PositiveInteger", DataTypeIndicator: "", Description: "an unsigned sequence of one or more Digits representing a decimal integer that has a value greater than 0. Leading zeroes are allowed.", MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SOTARef", DataTypeIndicator: "", Description: "a sequence of Characters representing an International SOTA Reference. The sequence comprises: an ITU prefix if applicable, a SOTA subdivision a / Character a SOTA Reference Number Examples: W2/WE-003 G/LD-003", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SecondaryAdministrativeSubdivisionListAlt", DataTypeIndicator: "", Description: "a semicolon (;) delimited, unordered list of one or more members of a Secondary_Administrative_Subdivision_Alt enumeration in the form: enumeration-name:enumeration-code Where there is more than one locality represented by the enumeration-code, they are separated by slash (/) characters. Only one of each enumeration-name valid for the DXCC entity concerned can appear in the list. Examples: <CNTY_ALT:28>NZ_Regions:Hawkes Bay/Wairoa <MY_CNTY_ALT:52>NZ_Islands:North Island;NZ_Regions:Hawkes Bay/Wairoa The first example shows the enumeration-name NZ_Regions with the region Hawkes Bay and the district Wairoa. For the purposes of illustration, the second example includes a non-existent subdivision with two available enumeration-codes, NZ_Islands:North Island and NZ_Islands:South Island. The example shows: the enumeration-name NZ_Islands with the island North Island the enumeration-name NZ_Regions with the region Hawkes Bay and the district Wairoa", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SecondarySubdivisionList", DataTypeIndicator: "", Description: "a colon-delimited list of two or more members of the Secondary_Administrative_Subdivision enumeration. E.g.: MA,Franklin:MA,Hampshire", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SponsoredAwardList", DataTypeIndicator: "", Description: "a comma-delimited list of members of the Sponsored_Award enumeration", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "String", DataTypeIndicator: "S", Description: "a sequence of Characters", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "Time", DataTypeIndicator: "T", Description: "6 Digits representing a UTC time in HHMMSS format or 4 Digits representing a time in HHMM format, where HH is a 2-Digit hour specifier, where 0 <= HH <= 23 [use leading zeroes] MM is a 2-Digit minute specifier, where 0 <= MM <= 59 [use leading zeroes] SS is a 2-Digit second specifier, where 0 <= SS <= 59 [use leading zeroes]", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "WWFFRef", DataTypeIndicator: "", Description: "a sequence of case-insensitive Characters representing an International WWFF (World Wildlife Flora & Fauna) reference in the form xxFF-nnnn comprising 8 to 11 characters where: xx is the WWFF national program and is 1 to 4 characters in length. FF- is two F characters followed by a dash character. nnnn represents the unique number within the national program and is 4 characters in length with leading zeros. Examples: KFF-4655 3DAFF-0002", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
}

// lookupMap contains all known ADIType specifications.
var lookupMap = map[ADIType]*Spec{
	AwardList:           &lookupList[0],
	Boolean:             &lookupList[1],
	Character:           &lookupList[2],
	CreditList:          &lookupList[3],
	Date:                &lookupList[4],
	Digit:               &lookupList[5],
	Enumeration:         &lookupList[6],
	GridSquare:          &lookupList[7],
	GridSquareExt:       &lookupList[8],
	GridSquareList:      &lookupList[9],
	IOTARefNo:           &lookupList[10],
	Integer:             &lookupList[11],
	IntlCharacter:       &lookupList[12],
	IntlMultilineString: &lookupList[13],
	IntlString:          &lookupList[14],
	Location:            &lookupList[15],
	MultilineString:     &lookupList[16],
	Number:              &lookupList[17],
	POTARef:             &lookupList[18],
	POTARefList:         &lookupList[19],
	PositiveInteger:     &lookupList[20],
	SOTARef:             &lookupList[21],
	SecondaryAdministrativeSubdivisionListAlt: &lookupList[22],
	SecondarySubdivisionList:                  &lookupList[23],
	SponsoredAwardList:                        &lookupList[24],
	String:                                    &lookupList[25],
	Time:                                      &lookupList[26],
	WWFFRef:                                   &lookupList[27],
}

// Lookup returns the specification for the provided ADIType.
// ADIF 3.1.6
func Lookup(aditype ADIType) (Spec, bool) {
	spec, ok := lookupMap[aditype]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all ADIType specifications that match the provided filter function.
// ADIF 3.1.6
func LookupByFilter(filter func(Spec) bool) []Spec {
	result := make([]Spec, 0, len(lookupList))
	for _, v := range lookupList {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// List returns all ADIType specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns ADIType specifications.
// This list excludes those marked as import-only.
// ADIF 3.1.6
func ListActive() []Spec {
	listActiveOnce.Do(func() {
		listActive = LookupByFilter(func(spec Spec) bool { return !bool(spec.IsImportOnly) })
	})

	result := make([]Spec, len(listActive))
	copy(result, listActive)
	return result
}
