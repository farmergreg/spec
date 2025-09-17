// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package aditype provides code and constants as defined in ADIF 3.1.6 (Released)
package aditype

import "sync"

const (
	AWARDLIST                                 Type = "awardlist"                                 // Deprecated: a comma-delimited list of members of the Award enumeration
	BOOLEAN                                   Type = "boolean"                                   // if True, the single ASCII character Y or y if False, the single ASCII character N or n
	CHARACTER                                 Type = "character"                                 // an ASCII character whose code lies in the range of 32 through 126, inclusive
	CREDITLIST                                Type = "creditlist"                                // a comma-delimited list where each list item is either: A member of the Credit enumeration. A member of the Credit enumeration followed by a colon and an ampersand-delimited list of members of the QSL_Medium enumeration. For example IOTA,WAS:LOTW&CARD,DXCC:CARD
	DATE                                      Type = "date"                                      // 8 Digits representing a UTC date in YYYYMMDD format, where YYYY is a 4-Digit year specifier, where 1930 <= YYYY MM is a 2-Digit month specifier, where 1 <= MM <= 12 [use leading zeroes] DD is a 2-Digit day specifier, where 1 <= DD <= DaysInMonth(MM) [use leading zeroes]
	DIGIT                                     Type = "digit"                                     // an ASCII character whose code lies in the range of 48 through 57, inclusive
	ENUMERATION                               Type = "enumeration"                               // an explicit list of legal case-insensitive values represented in ASCII set forth in set notation, e.g. {A, B, C, D}, or defined in a table, from which a single value may be selected.
	GRIDSQUARE                                Type = "gridsquare"                                // a case-insensitive 2-character, 4-character, 6-character, or 8-character Maidenhead locator. Specific fields impose additional restrictions on the number of characters; see the field descriptions for the allowed numbers of characters.
	GRIDSQUAREEXT                             Type = "gridsquareext"                             // For a 10-character Maidenhead locator, contains characters 9 and 10. For a 12-character Maidenhead locator, contains characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0-9.
	GRIDSQUARELIST                            Type = "gridsquarelist"                            // a comma-delimited list of GridSquare items
	IOTAREFNO                                 Type = "iotarefno"                                 // IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]
	INTEGER                                   Type = "integer"                                   // a sequence of one or more Digits representing a decimal integer, optionally preceded by a minus sign (ASCII code 45). Leading zeroes are allowed.
	INTLCHARACTER                             Type = "intlcharacter"                             // a Unicode character (encoded with UTF-8) excluding line break CR (code 13) and LF (code 10) characters
	INTLMULTILINESTRING                       Type = "intlmultilinestring"                       // a sequence of International Characters and line breaks. Fields of type IntlMultilineString must only be used in ADX files
	INTLSTRING                                Type = "intlstring"                                // a sequence of International Characters. Fields of type IntlString must only be used in ADX files
	LOCATION                                  Type = "location"                                  // a sequence of 11 characters representing a latitude or longitude in XDDD MM.MMM format, where X is a directional Character from the set {E, W, N, S} DDD is a 3-Digit degrees specifier, where 0 <= DDD <= 180 [use leading zeroes] There is a single space character in between DDD and MM.MMM MM.MMM is an unsigned Number minutes specifier with its decimal point in the third position, where 00.000 <= MM.MMM <= 59.999 [use leading and trailing zeroes]
	MULTILINESTRING                           Type = "multilinestring"                           // a sequence of Characters and line-breaks, where a line break is an ASCII CR (code 13) followed immediately by an ASCII LF (code 10)
	NUMBER                                    Type = "number"                                    // a sequence of one or more Digits representing a decimal number, optionally preceded by a minus sign (ASCII code 45) and optionally including a single decimal point (ASCII code 46)
	POTAREF                                   Type = "potaref"                                   // a sequence of case-insensitive Characters representing a Parks on the Air park reference in the form xxxx-nnnnn[@yyyyyy] comprising 6 to 17 characters where: xxxx is the POTA national program and is 1 to 4 characters in length, typically the default callsign prefix of the national program (rather than the DX entity) nnnnn represents the unique number within the national program and is either 4 or 5 characters in length (use the exact format listed on the POTA website) yyyyyy **Optional** is the 4 to 6 character ISO 3166-2 code to differentiate which state/province/prefecture/primary administration location the contact represents, in the case that the park reference spans more than one location (such as a trail). Examples of the POTARef Data Type: ReferenceLocation K-5033Golden Hill State Forest K-10000 5-digit park numbers are reserved for future use VE-5082@CA-ABThe Great Trail of Canada (the Canadian Trailway) National Scenic Trail, within Alberta, Canada 8P-0012Chancery Lane Swamp National Park VK-0556Pieman River State Reserve K-4562@US-CAPacific Crest Trail, within California, USA Additional Notes on POTARef: A browsable and searchable list of all park references is available. A complete CSV file is available (generated nightly). For more information, visit the Parks on the Air documentation website.
	POTAREFLIST                               Type = "potareflist"                               // a comma-delimited list of one or more POTARef items.
	POSITIVEINTEGER                           Type = "positiveinteger"                           // an unsigned sequence of one or more Digits representing a decimal integer that has a value greater than 0. Leading zeroes are allowed.
	SOTAREF                                   Type = "sotaref"                                   // a sequence of Characters representing an International SOTA Reference. The sequence comprises: an ITU prefix if applicable, a SOTA subdivision a / Character a SOTA Reference Number Examples: W2/WE-003 G/LD-003
	SECONDARYADMINISTRATIVESUBDIVISIONLISTALT Type = "secondaryadministrativesubdivisionlistalt" // a semicolon (;) delimited, unordered list of one or more members of a Secondary_Administrative_Subdivision_Alt enumeration in the form: enumeration-name:enumeration-code Where there is more than one locality represented by the enumeration-code, they are separated by slash (/) characters. Only one of each enumeration-name valid for the DXCC entity concerned can appear in the list. Examples: <CNTY_ALT:28>NZ_Regions:Hawkes Bay/Wairoa <MY_CNTY_ALT:52>NZ_Islands:North Island;NZ_Regions:Hawkes Bay/Wairoa The first example shows the enumeration-name NZ_Regions with the region Hawkes Bay and the district Wairoa. For the purposes of illustration, the second example includes a non-existent subdivision with two available enumeration-codes, NZ_Islands:North Island and NZ_Islands:South Island. The example shows: the enumeration-name NZ_Islands with the island North Island the enumeration-name NZ_Regions with the region Hawkes Bay and the district Wairoa
	SECONDARYSUBDIVISIONLIST                  Type = "secondarysubdivisionlist"                  // a colon-delimited list of two or more members of the Secondary_Administrative_Subdivision enumeration. E.g.: MA,Franklin:MA,Hampshire
	SPONSOREDAWARDLIST                        Type = "sponsoredawardlist"                        // a comma-delimited list of members of the Sponsored_Award enumeration
	STRING                                    Type = "string"                                    // a sequence of Characters
	TIME                                      Type = "time"                                      // 6 Digits representing a UTC time in HHMMSS format or 4 Digits representing a time in HHMM format, where HH is a 2-Digit hour specifier, where 0 <= HH <= 23 [use leading zeroes] MM is a 2-Digit minute specifier, where 0 <= MM <= 59 [use leading zeroes] SS is a 2-Digit second specifier, where 0 <= SS <= 59 [use leading zeroes]
	WWFFREF                                   Type = "wwffref"                                   // a sequence of case-insensitive Characters representing an International WWFF (World Wildlife Flora & Fauna) reference in the form xxFF-nnnn comprising 8 to 11 characters where: xx is the WWFF national program and is 1 to 4 characters in length. FF- is two F characters followed by a dash character. nnnn represents the unique number within the national program and is 4 characters in length with leading zeros. Examples: KFF-4655 3DAFF-0002
)

const DATATYPEINDICATOR_NONE DataTypeIndicator = 0 // Represents an empty/unset DataTypeIndicator (this is NOT part of the ADIF spec)

const (
	DATATYPEINDICATOR_BOOLEAN             DataTypeIndicator = 'b' // if True, the single ASCII character Y or y if False, the single ASCII character N or n
	DATATYPEINDICATOR_DATE                DataTypeIndicator = 'd' // 8 Digits representing a UTC date in YYYYMMDD format, where YYYY is a 4-Digit year specifier, where 1930 <= YYYY MM is a 2-Digit month specifier, where 1 <= MM <= 12 [use leading zeroes] DD is a 2-Digit day specifier, where 1 <= DD <= DaysInMonth(MM) [use leading zeroes]
	DATATYPEINDICATOR_ENUMERATION         DataTypeIndicator = 'e' // an explicit list of legal case-insensitive values represented in ASCII set forth in set notation, e.g. {A, B, C, D}, or defined in a table, from which a single value may be selected.
	DATATYPEINDICATOR_INTLMULTILINESTRING DataTypeIndicator = 'g' // a sequence of International Characters and line breaks. Fields of type IntlMultilineString must only be used in ADX files
	DATATYPEINDICATOR_INTLSTRING          DataTypeIndicator = 'i' // a sequence of International Characters. Fields of type IntlString must only be used in ADX files
	DATATYPEINDICATOR_LOCATION            DataTypeIndicator = 'l' // a sequence of 11 characters representing a latitude or longitude in XDDD MM.MMM format, where X is a directional Character from the set {E, W, N, S} DDD is a 3-Digit degrees specifier, where 0 <= DDD <= 180 [use leading zeroes] There is a single space character in between DDD and MM.MMM MM.MMM is an unsigned Number minutes specifier with its decimal point in the third position, where 00.000 <= MM.MMM <= 59.999 [use leading and trailing zeroes]
	DATATYPEINDICATOR_MULTILINESTRING     DataTypeIndicator = 'm' // a sequence of Characters and line-breaks, where a line break is an ASCII CR (code 13) followed immediately by an ASCII LF (code 10)
	DATATYPEINDICATOR_NUMBER              DataTypeIndicator = 'n' // a sequence of one or more Digits representing a decimal number, optionally preceded by a minus sign (ASCII code 45) and optionally including a single decimal point (ASCII code 46)
	DATATYPEINDICATOR_STRING              DataTypeIndicator = 's' // a sequence of Characters
	DATATYPEINDICATOR_TIME                DataTypeIndicator = 't' // 6 Digits representing a UTC time in HHMMSS format or 4 Digits representing a time in HHMM format, where HH is a 2-Digit hour specifier, where 0 <= HH <= 23 [use leading zeroes] MM is a 2-Digit minute specifier, where 0 <= MM <= 59 [use leading zeroes] SS is a 2-Digit second specifier, where 0 <= SS <= 59 [use leading zeroes]
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Type specifications.
var lookupList = []Spec{
	{Key: "awardlist", DataTypeIndicator: 0, Description: "a comma-delimited list of members of the Award enumeration", MinimumValue: 0, MaximumValue: 0, IsImportOnly: true, Comments: ""},
	{Key: "boolean", DataTypeIndicator: 98, Description: "if True, the single ASCII character Y or y if False, the single ASCII character N or n", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "character", DataTypeIndicator: 0, Description: "an ASCII character whose code lies in the range of 32 through 126, inclusive", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "creditlist", DataTypeIndicator: 0, Description: "a comma-delimited list where each list item is either: A member of the Credit enumeration. A member of the Credit enumeration followed by a colon and an ampersand-delimited list of members of the QSL_Medium enumeration. For example IOTA,WAS:LOTW&CARD,DXCC:CARD", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "date", DataTypeIndicator: 100, Description: "8 Digits representing a UTC date in YYYYMMDD format, where YYYY is a 4-Digit year specifier, where 1930 <= YYYY MM is a 2-Digit month specifier, where 1 <= MM <= 12 [use leading zeroes] DD is a 2-Digit day specifier, where 1 <= DD <= DaysInMonth(MM) [use leading zeroes]", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "digit", DataTypeIndicator: 0, Description: "an ASCII character whose code lies in the range of 48 through 57, inclusive", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "enumeration", DataTypeIndicator: 101, Description: "an explicit list of legal case-insensitive values represented in ASCII set forth in set notation, e.g. {A, B, C, D}, or defined in a table, from which a single value may be selected.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "gridsquare", DataTypeIndicator: 0, Description: "a case-insensitive 2-character, 4-character, 6-character, or 8-character Maidenhead locator. Specific fields impose additional restrictions on the number of characters; see the field descriptions for the allowed numbers of characters.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "gridsquareext", DataTypeIndicator: 0, Description: "For a 10-character Maidenhead locator, contains characters 9 and 10. For a 12-character Maidenhead locator, contains characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0-9.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "gridsquarelist", DataTypeIndicator: 0, Description: "a comma-delimited list of GridSquare items", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "iotarefno", DataTypeIndicator: 0, Description: "IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "integer", DataTypeIndicator: 0, Description: "a sequence of one or more Digits representing a decimal integer, optionally preceded by a minus sign (ASCII code 45). Leading zeroes are allowed.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "intlcharacter", DataTypeIndicator: 0, Description: "a Unicode character (encoded with UTF-8) excluding line break CR (code 13) and LF (code 10) characters", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "intlmultilinestring", DataTypeIndicator: 103, Description: "a sequence of International Characters and line breaks. Fields of type IntlMultilineString must only be used in ADX files", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "intlstring", DataTypeIndicator: 105, Description: "a sequence of International Characters. Fields of type IntlString must only be used in ADX files", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "location", DataTypeIndicator: 108, Description: "a sequence of 11 characters representing a latitude or longitude in XDDD MM.MMM format, where X is a directional Character from the set {E, W, N, S} DDD is a 3-Digit degrees specifier, where 0 <= DDD <= 180 [use leading zeroes] There is a single space character in between DDD and MM.MMM MM.MMM is an unsigned Number minutes specifier with its decimal point in the third position, where 00.000 <= MM.MMM <= 59.999 [use leading and trailing zeroes]", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "multilinestring", DataTypeIndicator: 109, Description: "a sequence of Characters and line-breaks, where a line break is an ASCII CR (code 13) followed immediately by an ASCII LF (code 10)", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "number", DataTypeIndicator: 110, Description: "a sequence of one or more Digits representing a decimal number, optionally preceded by a minus sign (ASCII code 45) and optionally including a single decimal point (ASCII code 46)", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "potaref", DataTypeIndicator: 0, Description: "a sequence of case-insensitive Characters representing a Parks on the Air park reference in the form xxxx-nnnnn[@yyyyyy] comprising 6 to 17 characters where: xxxx is the POTA national program and is 1 to 4 characters in length, typically the default callsign prefix of the national program (rather than the DX entity) nnnnn represents the unique number within the national program and is either 4 or 5 characters in length (use the exact format listed on the POTA website) yyyyyy **Optional** is the 4 to 6 character ISO 3166-2 code to differentiate which state/province/prefecture/primary administration location the contact represents, in the case that the park reference spans more than one location (such as a trail). Examples of the POTARef Data Type: ReferenceLocation K-5033Golden Hill State Forest K-10000 5-digit park numbers are reserved for future use VE-5082@CA-ABThe Great Trail of Canada (the Canadian Trailway) National Scenic Trail, within Alberta, Canada 8P-0012Chancery Lane Swamp National Park VK-0556Pieman River State Reserve K-4562@US-CAPacific Crest Trail, within California, USA Additional Notes on POTARef: A browsable and searchable list of all park references is available. A complete CSV file is available (generated nightly). For more information, visit the Parks on the Air documentation website.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "potareflist", DataTypeIndicator: 0, Description: "a comma-delimited list of one or more POTARef items.", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "positiveinteger", DataTypeIndicator: 0, Description: "an unsigned sequence of one or more Digits representing a decimal integer that has a value greater than 0. Leading zeroes are allowed.", MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "sotaref", DataTypeIndicator: 0, Description: "a sequence of Characters representing an International SOTA Reference. The sequence comprises: an ITU prefix if applicable, a SOTA subdivision a / Character a SOTA Reference Number Examples: W2/WE-003 G/LD-003", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "secondaryadministrativesubdivisionlistalt", DataTypeIndicator: 0, Description: "a semicolon (;) delimited, unordered list of one or more members of a Secondary_Administrative_Subdivision_Alt enumeration in the form: enumeration-name:enumeration-code Where there is more than one locality represented by the enumeration-code, they are separated by slash (/) characters. Only one of each enumeration-name valid for the DXCC entity concerned can appear in the list. Examples: <CNTY_ALT:28>NZ_Regions:Hawkes Bay/Wairoa <MY_CNTY_ALT:52>NZ_Islands:North Island;NZ_Regions:Hawkes Bay/Wairoa The first example shows the enumeration-name NZ_Regions with the region Hawkes Bay and the district Wairoa. For the purposes of illustration, the second example includes a non-existent subdivision with two available enumeration-codes, NZ_Islands:North Island and NZ_Islands:South Island. The example shows: the enumeration-name NZ_Islands with the island North Island the enumeration-name NZ_Regions with the region Hawkes Bay and the district Wairoa", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "secondarysubdivisionlist", DataTypeIndicator: 0, Description: "a colon-delimited list of two or more members of the Secondary_Administrative_Subdivision enumeration. E.g.: MA,Franklin:MA,Hampshire", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "sponsoredawardlist", DataTypeIndicator: 0, Description: "a comma-delimited list of members of the Sponsored_Award enumeration", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "string", DataTypeIndicator: 115, Description: "a sequence of Characters", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "time", DataTypeIndicator: 116, Description: "6 Digits representing a UTC time in HHMMSS format or 4 Digits representing a time in HHMM format, where HH is a 2-Digit hour specifier, where 0 <= HH <= 23 [use leading zeroes] MM is a 2-Digit minute specifier, where 0 <= MM <= 59 [use leading zeroes] SS is a 2-Digit second specifier, where 0 <= SS <= 59 [use leading zeroes]", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "wwffref", DataTypeIndicator: 0, Description: "a sequence of case-insensitive Characters representing an International WWFF (World Wildlife Flora & Fauna) reference in the form xxFF-nnnn comprising 8 to 11 characters where: xx is the WWFF national program and is 1 to 4 characters in length. FF- is two F characters followed by a dash character. nnnn represents the unique number within the national program and is 4 characters in length with leading zeros. Examples: KFF-4655 3DAFF-0002", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
}

// lookupMap contains all known Type specifications.
var lookupMap = map[Type]*Spec{
	AWARDLIST:           &lookupList[0],
	BOOLEAN:             &lookupList[1],
	CHARACTER:           &lookupList[2],
	CREDITLIST:          &lookupList[3],
	DATE:                &lookupList[4],
	DIGIT:               &lookupList[5],
	ENUMERATION:         &lookupList[6],
	GRIDSQUARE:          &lookupList[7],
	GRIDSQUAREEXT:       &lookupList[8],
	GRIDSQUARELIST:      &lookupList[9],
	IOTAREFNO:           &lookupList[10],
	INTEGER:             &lookupList[11],
	INTLCHARACTER:       &lookupList[12],
	INTLMULTILINESTRING: &lookupList[13],
	INTLSTRING:          &lookupList[14],
	LOCATION:            &lookupList[15],
	MULTILINESTRING:     &lookupList[16],
	NUMBER:              &lookupList[17],
	POTAREF:             &lookupList[18],
	POTAREFLIST:         &lookupList[19],
	POSITIVEINTEGER:     &lookupList[20],
	SOTAREF:             &lookupList[21],
	SECONDARYADMINISTRATIVESUBDIVISIONLISTALT: &lookupList[22],
	SECONDARYSUBDIVISIONLIST:                  &lookupList[23],
	SPONSOREDAWARDLIST:                        &lookupList[24],
	STRING:                                    &lookupList[25],
	TIME:                                      &lookupList[26],
	WWFFREF:                                   &lookupList[27],
}

// Lookup returns the specification for the provided Type.
// ADIF 3.1.6
func Lookup(t Type) (Spec, bool) {
	spec, ok := lookupMap[t]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all Type specifications that match the provided filter function.
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

// List returns all Type specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns Type specifications.
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
