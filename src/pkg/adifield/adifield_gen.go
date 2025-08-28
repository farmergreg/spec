// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package adifield

var (
	ADIF_VER          ADIField = "ADIF_VER"          // Header Field: identifies the version of ADIF used in this file in the format X.Y.Z where X is an integer designating the ADIF epoch Y is an integer between 0 and 9 designating the major version Z is an integer between 0 and 9 designating the minor version
	CREATED_TIMESTAMP ADIField = "CREATED_TIMESTAMP" // Header Field: identifies the UTC date and time that the file was created in the format of 15 characters YYYYMMDD HHMMSS where YYYYMMDD is a Date data type HHMMSS is a 6 character Time data type
	PROGRAMID         ADIField = "PROGRAMID"         // Header Field: identifies the name of the logger, converter, or utility that created or processed this ADIF file To help avoid name clashes, the ADIF PROGRAMID Register provides a voluntary list of PROGRAMID values.
	PROGRAMVERSION    ADIField = "PROGRAMVERSION"    // Header Field: identifies the version of the logger, converter, or utility that created or processed this ADIF file
)

var (
	ADDRESS                    ADIField = "ADDRESS"                    // Record Field: the contacted station's complete mailing address: full name, street address, city, postal code, and country
	ADDRESS_INTL               ADIField = "ADDRESS_INTL"               // Record Field: the contacted station's complete mailing address: full name, street address, city, postal code, and country
	AGE                        ADIField = "AGE"                        // Record Field: the contacted station's operator's age in years in the range 0 to 120 (inclusive)
	ALTITUDE                   ADIField = "ALTITUDE"                   // Record Field: the height of the contacted station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <ALTITUDE:4>1500 and 10.5 m is <ALTITUDE:4>10.5
	ANT_AZ                     ADIField = "ANT_AZ"                     // Record Field: the logging station's antenna azimuth, in degrees with a value between 0 to 360 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 370 is exported as 10). True north is 0 degrees with values increasing in a clockwise direction.
	ANT_EL                     ADIField = "ANT_EL"                     // Record Field: the logging station's antenna elevation, in degrees with a value between -90 to 90 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 100 is exported as 80). The horizon is 0 degrees with values increasing as the angle moves in an upward direction.
	ANT_PATH                   ADIField = "ANT_PATH"                   // Record Field: the signal path
	ARRL_SECT                  ADIField = "ARRL_SECT"                  // Record Field: the contacted station's ARRL section
	AWARD_GRANTED              ADIField = "AWARD_GRANTED"              // Record Field: the list of awards granted by a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's "Award Account" between an award sponsor and the user. For example, in response to a request "send me a list of the awards granted to AA6YQ", this might be received: <CALL:5>AA6YQ <AWARD_GRANTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m
	AWARD_SUBMITTED            ADIField = "AWARD_SUBMITTED"            // Record Field: the list of awards submitted to a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's "Award Account" between an award sponsor and the user. For example, AA6YQ might submit a request for awards by sending the following: <CALL:5>AA6YQ <AWARD_SUBMITTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m
	A_INDEX                    ADIField = "A_INDEX"                    // Record Field: the geomagnetic A index at the time of the QSO in the range 0 to 400 (inclusive)
	BAND                       ADIField = "BAND"                       // Record Field: QSO Band
	BAND_RX                    ADIField = "BAND_RX"                    // Record Field: in a split frequency QSO, the logging station's receiving band
	CALL                       ADIField = "CALL"                       // Record Field: the contacted station's callsign
	CHECK                      ADIField = "CHECK"                      // Record Field: contest check (e.g. for ARRL Sweepstakes)
	CLASS                      ADIField = "CLASS"                      // Record Field: contest class (e.g. for ARRL Field Day)
	CLUBLOG_QSO_UPLOAD_DATE    ADIField = "CLUBLOG_QSO_UPLOAD_DATE"    // Record Field: the date the QSO was last uploaded to the Club Log online service
	CLUBLOG_QSO_UPLOAD_STATUS  ADIField = "CLUBLOG_QSO_UPLOAD_STATUS"  // Record Field: the upload status of the QSO on the Club Log online service
	CNTY                       ADIField = "CNTY"                       // Record Field: the contacted station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format
	CNTY_ALT                   ADIField = "CNTY_ALT"                   // Record Field: a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the contacted station See the Data Type for details.
	COMMENT                    ADIField = "COMMENT"                    // Record Field: comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG field recommended use: information of interest to the contacted station's operator
	COMMENT_INTL               ADIField = "COMMENT_INTL"               // Record Field: comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG_INTL field recommended use: information of interest to the contacted station's operator
	CONT                       ADIField = "CONT"                       // Record Field: the contacted station's Continent
	CONTACTED_OP               ADIField = "CONTACTED_OP"               // Record Field: the callsign of the individual operating the contacted station
	CONTEST_ID                 ADIField = "CONTEST_ID"                 // Record Field: QSO Contest Identifier use enumeration values for interoperability
	COUNTRY                    ADIField = "COUNTRY"                    // Record Field: the contacted station's DXCC entity name
	COUNTRY_INTL               ADIField = "COUNTRY_INTL"               // Record Field: the contacted station's DXCC entity name
	CQZ                        ADIField = "CQZ"                        // Record Field: the contacted station's CQ Zone in the range 1 to 40 (inclusive)
	CREDIT_GRANTED             ADIField = "CREDIT_GRANTED"             // Record Field: the list of credits granted to this QSO Use of data type AwardList and enumeration Award are import-only
	CREDIT_SUBMITTED           ADIField = "CREDIT_SUBMITTED"           // Record Field: the list of credits sought for this QSO Use of data type AwardList and enumeration Award are import-only
	DARC_DOK                   ADIField = "DARC_DOK"                   // Record Field: the contacted station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <DARC_DOK:3>A01
	DCL_QSLRDATE               ADIField = "DCL_QSLRDATE"               // Record Field: date QSL received from DCL (only valid if DCL_QSL_RCVD is Y, I, or V)(V import-only)
	DCL_QSLSDATE               ADIField = "DCL_QSLSDATE"               // Record Field: date QSL sent to DCL (only valid if DCL_QSL_SENT is Y, Q, or I)
	DCL_QSL_RCVD               ADIField = "DCL_QSL_RCVD"               // Record Field: DCL QSL received status Default Value: N
	DCL_QSL_SENT               ADIField = "DCL_QSL_SENT"               // Record Field: DCL QSL sent status Default Value: N
	DISTANCE                   ADIField = "DISTANCE"                   // Record Field: the distance between the logging station and the contacted station in kilometers via the specified signal path with a value greater than or equal to 0
	DXCC                       ADIField = "DXCC"                       // Record Field: the contacted station's DXCC Entity Code <DXCC:1>0 means that the contacted station is known not to be within a DXCC entity.
	EMAIL                      ADIField = "EMAIL"                      // Record Field: the contacted station's email address
	EQSL_AG                    ADIField = "EQSL_AG"                    // Record Field: indicates whether the QSO is known to be "Authenticity Guaranteed" by eQSL
	EQSL_QSLRDATE              ADIField = "EQSL_QSLRDATE"              // Record Field: date QSL received from eQSL.cc (only valid if EQSL_QSL_RCVD is Y, I, or V)(V import-only)
	EQSL_QSLSDATE              ADIField = "EQSL_QSLSDATE"              // Record Field: date QSL sent to eQSL.cc (only valid if EQSL_QSL_SENT is Y, Q, or I)
	EQSL_QSL_RCVD              ADIField = "EQSL_QSL_RCVD"              // Record Field: eQSL.cc QSL received status instead of V (import-only) use <CREDIT_GRANTED:42>CQWAZ:eqsl,CQWAZ_BAND:eqsl,CQWAZ_MODE:eqsl Default Value: N
	EQSL_QSL_SENT              ADIField = "EQSL_QSL_SENT"              // Record Field: eQSL.cc QSL sent status Default Value: N
	EQ_CALL                    ADIField = "EQ_CALL"                    // Record Field: the contacted station's owner's callsign
	FISTS                      ADIField = "FISTS"                      // Record Field: the contacted station's FISTS CW Club member number with a value greater than 0.
	FISTS_CC                   ADIField = "FISTS_CC"                   // Record Field: the contacted station's FISTS CW Club Century Certificate (CC) number with a value greater than 0.
	FORCE_INIT                 ADIField = "FORCE_INIT"                 // Record Field: new EME "initial"
	FREQ                       ADIField = "FREQ"                       // Record Field: QSO frequency in Megahertz
	FREQ_RX                    ADIField = "FREQ_RX"                    // Record Field: in a split frequency QSO, the logging station's receiving frequency in Megahertz
	GRIDSQUARE                 ADIField = "GRIDSQUARE"                 // Record Field: the contacted station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in GRIDSQUARE and the additional 2 or 4 characters in the GRIDSQUARE_EXT field
	GRIDSQUARE_EXT             ADIField = "GRIDSQUARE_EXT"             // Record Field: for a contacted station's 10-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9 and 10. For a contacted station's 12-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <GRIDSQUARE:8>FN01MH42 <GRIDSQUARE_EXT:2>BQ
	GUEST_OP                   ADIField = "GUEST_OP"                   // Deprecated: Record Field: import-only: use OPERATOR instead
	HAMLOGEU_QSO_UPLOAD_DATE   ADIField = "HAMLOGEU_QSO_UPLOAD_DATE"   // Record Field: the date the QSO was last uploaded to the HAMLOG.EU online service
	HAMLOGEU_QSO_UPLOAD_STATUS ADIField = "HAMLOGEU_QSO_UPLOAD_STATUS" // Record Field: the upload status of the QSO on the HAMLOG.EU online service
	HAMQTH_QSO_UPLOAD_DATE     ADIField = "HAMQTH_QSO_UPLOAD_DATE"     // Record Field: the date the QSO was last uploaded to the HamQTH.com online service
	HAMQTH_QSO_UPLOAD_STATUS   ADIField = "HAMQTH_QSO_UPLOAD_STATUS"   // Record Field: the upload status of the QSO on the HamQTH.com online service
	HRDLOG_QSO_UPLOAD_DATE     ADIField = "HRDLOG_QSO_UPLOAD_DATE"     // Record Field: the date the QSO was last uploaded to the HRDLog.net online service
	HRDLOG_QSO_UPLOAD_STATUS   ADIField = "HRDLOG_QSO_UPLOAD_STATUS"   // Record Field: the upload status of the QSO on the HRDLog.net online service
	IOTA                       ADIField = "IOTA"                       // Record Field: the contacted station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]
	IOTA_ISLAND_ID             ADIField = "IOTA_ISLAND_ID"             // Record Field: the contacted station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]
	ITUZ                       ADIField = "ITUZ"                       // Record Field: the contacted station's ITU zone in the range 1 to 90 (inclusive)
	K_INDEX                    ADIField = "K_INDEX"                    // Record Field: the geomagnetic K index at the time of the QSO in the range 0 to 9 (inclusive)
	LAT                        ADIField = "LAT"                        // Record Field: the contacted station's latitude
	LON                        ADIField = "LON"                        // Record Field: the contacted station's longitude
	LOTW_QSLRDATE              ADIField = "LOTW_QSLRDATE"              // Record Field: date QSL received from ARRL Logbook of the World (only valid if LOTW_QSL_RCVD is Y, I, or V)(V import-only)
	LOTW_QSLSDATE              ADIField = "LOTW_QSLSDATE"              // Record Field: date QSL sent to ARRL Logbook of the World (only valid if LOTW_QSL_SENT is Y, Q, or I)
	LOTW_QSL_RCVD              ADIField = "LOTW_QSL_RCVD"              // Record Field: ARRL Logbook of the World QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw Default Value: N
	LOTW_QSL_SENT              ADIField = "LOTW_QSL_SENT"              // Record Field: ARRL Logbook of the World QSL sent status Default Value: N
	MAX_BURSTS                 ADIField = "MAX_BURSTS"                 // Record Field: maximum length of meteor scatter bursts heard by the logging station, in seconds with a value greater than or equal to 0
	MODE                       ADIField = "MODE"                       // Record Field: QSO Mode
	MORSE_KEY_INFO             ADIField = "MORSE_KEY_INFO"             // Record Field: details of the contacted station's Morse key (e.g. make, model, etc). Example: <MORSE_KEY_INFO:16>Begali Sculpture
	MORSE_KEY_TYPE             ADIField = "MORSE_KEY_TYPE"             // Record Field: the contacted station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP
	MS_SHOWER                  ADIField = "MS_SHOWER"                  // Record Field: For Meteor Scatter QSOs, the name of the meteor shower in progress
	MY_ALTITUDE                ADIField = "MY_ALTITUDE"                // Record Field: the height of the logging station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <MY_ALTITUDE:4>1500 and 10.5 m is <MY_ALTITUDE:4>10.5
	MY_ANTENNA                 ADIField = "MY_ANTENNA"                 // Record Field: the logging station's antenna
	MY_ANTENNA_INTL            ADIField = "MY_ANTENNA_INTL"            // Record Field: the logging station's antenna
	MY_ARRL_SECT               ADIField = "MY_ARRL_SECT"               // Record Field: the logging station's ARRL section
	MY_CITY                    ADIField = "MY_CITY"                    // Record Field: the logging station's city
	MY_CITY_INTL               ADIField = "MY_CITY_INTL"               // Record Field: the logging station's city
	MY_CNTY                    ADIField = "MY_CNTY"                    // Record Field: the logging station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format
	MY_CNTY_ALT                ADIField = "MY_CNTY_ALT"                // Record Field: a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the logging station See the Data Type for details.
	MY_COUNTRY                 ADIField = "MY_COUNTRY"                 // Record Field: the logging station's DXCC entity name
	MY_COUNTRY_INTL            ADIField = "MY_COUNTRY_INTL"            // Record Field: the logging station's DXCC entity name
	MY_CQ_ZONE                 ADIField = "MY_CQ_ZONE"                 // Record Field: the logging station's CQ Zone in the range 1 to 40 (inclusive)
	MY_DARC_DOK                ADIField = "MY_DARC_DOK"                // Record Field: the logging station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <MY_DARC_DOK:3>A01
	MY_DXCC                    ADIField = "MY_DXCC"                    // Record Field: the logging station's DXCC Entity Code <MY_DXCC:1>0 means that the logging station is known not to be within a DXCC entity.
	MY_FISTS                   ADIField = "MY_FISTS"                   // Record Field: the logging station's FISTS CW Club member number with a value greater than 0.
	MY_GRIDSQUARE              ADIField = "MY_GRIDSQUARE"              // Record Field: the logging station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in MY_GRIDSQUARE and the additional 2 or 4 characters in the MY_GRIDSQUARE_EXT field
	MY_GRIDSQUARE_EXT          ADIField = "MY_GRIDSQUARE_EXT"          // Record Field: for a logging station's 10-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9 and 10. For a logging station's 12-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <MY_GRIDSQUARE:8>FN01MH42 <MY_GRIDSQUARE_EXT:2>BQ
	MY_IOTA                    ADIField = "MY_IOTA"                    // Record Field: the logging station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]
	MY_IOTA_ISLAND_ID          ADIField = "MY_IOTA_ISLAND_ID"          // Record Field: the logging station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]
	MY_ITU_ZONE                ADIField = "MY_ITU_ZONE"                // Record Field: the logging station's ITU zone in the range 1 to 90 (inclusive)
	MY_LAT                     ADIField = "MY_LAT"                     // Record Field: the logging station's latitude
	MY_LON                     ADIField = "MY_LON"                     // Record Field: the logging station's longitude
	MY_MORSE_KEY_INFO          ADIField = "MY_MORSE_KEY_INFO"          // Record Field: details of the logging station's Morse key (e.g. make, model, etc). Example: <MY_MORSE_KEY_INFO:16>Begali Sculpture
	MY_MORSE_KEY_TYPE          ADIField = "MY_MORSE_KEY_TYPE"          // Record Field: the logging station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP
	MY_NAME                    ADIField = "MY_NAME"                    // Record Field: the logging operator's name
	MY_NAME_INTL               ADIField = "MY_NAME_INTL"               // Record Field: the logging operator's name
	MY_POSTAL_CODE             ADIField = "MY_POSTAL_CODE"             // Record Field: the logging station's postal code
	MY_POSTAL_CODE_INTL        ADIField = "MY_POSTAL_CODE_INTL"        // Record Field: the logging station's postal code
	MY_POTA_REF                ADIField = "MY_POTA_REF"                // Record Field: a comma-delimited list of one or more of the logging station's POTA (Parks on the Air) reference(s). Examples: <MY_POTA_REF:6>K-0059 <MY_POTA_REF:7>K-10000 <MY_POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY
	MY_RIG                     ADIField = "MY_RIG"                     // Record Field: description of the logging station's equipment
	MY_RIG_INTL                ADIField = "MY_RIG_INTL"                // Record Field: description of the logging station's equipment
	MY_SIG                     ADIField = "MY_SIG"                     // Record Field: special interest activity or event
	MY_SIG_INFO                ADIField = "MY_SIG_INFO"                // Record Field: special interest activity or event information
	MY_SIG_INFO_INTL           ADIField = "MY_SIG_INFO_INTL"           // Record Field: special interest activity or event information
	MY_SIG_INTL                ADIField = "MY_SIG_INTL"                // Record Field: special interest activity or event
	MY_SOTA_REF                ADIField = "MY_SOTA_REF"                // Record Field: the logging station's International SOTA Reference.
	MY_STATE                   ADIField = "MY_STATE"                   // Record Field: the code for the logging station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)
	MY_STREET                  ADIField = "MY_STREET"                  // Record Field: the logging station's street
	MY_STREET_INTL             ADIField = "MY_STREET_INTL"             // Record Field: the logging station's street
	MY_USACA_COUNTIES          ADIField = "MY_USACA_COUNTIES"          // Record Field: two US counties in the case where the logging station is located on a border between two counties, representing counties that the contacted station may claim for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire
	MY_VUCC_GRIDS              ADIField = "MY_VUCC_GRIDS"              // Record Field: two or four adjacent Maidenhead grid locators, each four or six characters long, representing the logging station's grid squares that the contacted station may claim for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07
	MY_WWFF_REF                ADIField = "MY_WWFF_REF"                // Record Field: the logging station's WWFF (World Wildlife Flora & Fauna) reference
	NAME                       ADIField = "NAME"                       // Record Field: the contacted station's operator's name
	NAME_INTL                  ADIField = "NAME_INTL"                  // Record Field: the contacted station's operator's name
	NOTES                      ADIField = "NOTES"                      // Record Field: QSO notes recommended use: information of interest to the logging station's operator
	NOTES_INTL                 ADIField = "NOTES_INTL"                 // Record Field: QSO notes recommended use: information of interest to the logging station's operator
	NR_BURSTS                  ADIField = "NR_BURSTS"                  // Record Field: the number of meteor scatter bursts heard by the logging station with a value greater than or equal to 0
	NR_PINGS                   ADIField = "NR_PINGS"                   // Record Field: the number of meteor scatter pings heard by the logging station with a value greater than or equal to 0
	OPERATOR                   ADIField = "OPERATOR"                   // Record Field: the logging operator's callsign if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign
	OWNER_CALLSIGN             ADIField = "OWNER_CALLSIGN"             // Record Field: the callsign of the owner of the station used to log the contact (the callsign of the OPERATOR's host) if OWNER_CALLSIGN is absent, STATION_CALLSIGN shall be treated as both the logging station's callsign and the callsign of the owner of the station
	PFX                        ADIField = "PFX"                        // Record Field: the contacted station's WPX prefix
	POTA_REF                   ADIField = "POTA_REF"                   // Record Field: a comma-delimited list of one or more of the contacted station's POTA (Parks on the Air) reference(s). Examples: <POTA_REF:6>K-5033 <POTA_REF:13>VE-5082@CA-AB <POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY
	PRECEDENCE                 ADIField = "PRECEDENCE"                 // Record Field: contest precedence (e.g. for ARRL Sweepstakes)
	PROP_MODE                  ADIField = "PROP_MODE"                  // Record Field: QSO propagation mode
	PUBLIC_KEY                 ADIField = "PUBLIC_KEY"                 // Record Field: public encryption key
	QRZCOM_QSO_DOWNLOAD_DATE   ADIField = "QRZCOM_QSO_DOWNLOAD_DATE"   // Record Field: date QSO downloaded from QRZ.COM logbook
	QRZCOM_QSO_DOWNLOAD_STATUS ADIField = "QRZCOM_QSO_DOWNLOAD_STATUS" // Record Field: QRZ.COM logbook QSO download status
	QRZCOM_QSO_UPLOAD_DATE     ADIField = "QRZCOM_QSO_UPLOAD_DATE"     // Record Field: the date the QSO was last uploaded to the QRZ.COM online service
	QRZCOM_QSO_UPLOAD_STATUS   ADIField = "QRZCOM_QSO_UPLOAD_STATUS"   // Record Field: the upload status of the QSO on the QRZ.COM online service
	QSLMSG                     ADIField = "QSLMSG"                     // Record Field: a message for the contacted station's operator to be incorporated in a paper or electronic QSL
	QSLMSG_INTL                ADIField = "QSLMSG_INTL"                // Record Field: a message for the contacted station's operator to be incorporated in a paper or electronic QSL
	QSLMSG_RCVD                ADIField = "QSLMSG_RCVD"                // Record Field: a message addressed to the logging station's operator incorporated in a paper or electronic QSL
	QSLRDATE                   ADIField = "QSLRDATE"                   // Record Field: QSL received date (only valid if QSL_RCVD is Y, I, or V)(V import-only)
	QSLSDATE                   ADIField = "QSLSDATE"                   // Record Field: QSL sent date (only valid if QSL_SENT is Y, Q, or I)
	QSL_RCVD                   ADIField = "QSL_RCVD"                   // Record Field: QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card Default Value: N
	QSL_RCVD_VIA               ADIField = "QSL_RCVD_VIA"               // Record Field: if QSL_RCVD is set to 'Y' or 'V', the means by which the QSL was received by the logging station; otherwise, the means by which the logging station requested or intends to request that the QSL be conveyed. (Note: 'V' is import-only) use of M (manager) is import-only
	QSL_SENT                   ADIField = "QSL_SENT"                   // Record Field: QSL sent status Default Value: N
	QSL_SENT_VIA               ADIField = "QSL_SENT_VIA"               // Record Field: if QSL_SENT is set to 'Y', the means by which the QSL was sent by the logging station; otherwise, the means by which the logging station intends to convey the QSL use of M (manager) is import-only
	QSL_VIA                    ADIField = "QSL_VIA"                    // Record Field: the contacted station's QSL route
	QSO_COMPLETE               ADIField = "QSO_COMPLETE"               // Record Field: indicates whether the QSO was complete from the perspective of the logging station Y - yes N - no NIL - not heard ? - uncertain
	QSO_DATE                   ADIField = "QSO_DATE"                   // Record Field: date on which the QSO started
	QSO_DATE_OFF               ADIField = "QSO_DATE_OFF"               // Record Field: date on which the QSO ended
	QSO_RANDOM                 ADIField = "QSO_RANDOM"                 // Record Field: indicates whether the QSO was random or scheduled
	QTH                        ADIField = "QTH"                        // Record Field: the contacted station's city
	QTH_INTL                   ADIField = "QTH_INTL"                   // Record Field: the contacted station's city
	REGION                     ADIField = "REGION"                     // Record Field: the contacted station's WAE or CQ entity contained within a DXCC entity. the value None indicates that the WAE or CQ entity is the DXCC entity in the DXCC field. nothing can be inferred from the absence of the REGION field
	RIG                        ADIField = "RIG"                        // Record Field: description of the contacted station's equipment
	RIG_INTL                   ADIField = "RIG_INTL"                   // Record Field: description of the contacted station's equipment
	RST_RCVD                   ADIField = "RST_RCVD"                   // Record Field: signal report from the contacted station
	RST_SENT                   ADIField = "RST_SENT"                   // Record Field: signal report sent to the contacted station
	RX_PWR                     ADIField = "RX_PWR"                     // Record Field: the contacted station's transmitter power in Watts with a value greater than or equal to 0
	SAT_MODE                   ADIField = "SAT_MODE"                   // Record Field: satellite mode - a code representing the satellite's uplink band and downlink band
	SAT_NAME                   ADIField = "SAT_NAME"                   // Record Field: name of satellite
	SFI                        ADIField = "SFI"                        // Record Field: the solar flux at the time of the QSO in the range 0 to 300 (inclusive).
	SIG                        ADIField = "SIG"                        // Record Field: the name of the contacted station's special activity or interest group
	SIG_INFO                   ADIField = "SIG_INFO"                   // Record Field: information associated with the contacted station's activity or interest group
	SIG_INFO_INTL              ADIField = "SIG_INFO_INTL"              // Record Field: information associated with the contacted station's activity or interest group
	SIG_INTL                   ADIField = "SIG_INTL"                   // Record Field: the name of the contacted station's special activity or interest group
	SILENT_KEY                 ADIField = "SILENT_KEY"                 // Record Field: 'Y' indicates that the contacted station's operator is now a Silent Key.
	SKCC                       ADIField = "SKCC"                       // Record Field: the contacted station's Straight Key Century Club (SKCC) member information
	SOTA_REF                   ADIField = "SOTA_REF"                   // Record Field: the contacted station's International SOTA Reference.
	SRX                        ADIField = "SRX"                        // Record Field: contest QSO received serial number with a value greater than or equal to 0
	SRX_STRING                 ADIField = "SRX_STRING"                 // Record Field: contest QSO received information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail
	STATE                      ADIField = "STATE"                      // Record Field: the code for the contacted station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)
	STATION_CALLSIGN           ADIField = "STATION_CALLSIGN"           // Record Field: the logging station's callsign (the callsign used over the air) if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign
	STX                        ADIField = "STX"                        // Record Field: contest QSO transmitted serial number with a value greater than or equal to 0
	STX_STRING                 ADIField = "STX_STRING"                 // Record Field: contest QSO transmitted information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail
	SUBMODE                    ADIField = "SUBMODE"                    // Record Field: QSO Submode use enumeration values for interoperability
	SWL                        ADIField = "SWL"                        // Record Field: indicates that the QSO information pertains to an SWL report
	TEN_TEN                    ADIField = "TEN_TEN"                    // Record Field: Ten-Ten number with a value greater than 0
	TIME_OFF                   ADIField = "TIME_OFF"                   // Record Field: HHMM or HHMMSS in UTC in the absence of <QSO_DATE_OFF>, the QSO duration is less than 24 hours. For example, the following is a QSO starting at 14 July 2020 23:55 and finishing at 15 July 2020 01:00: <QSO_DATE:8>20200714 <TIME_ON:4>2355 <TIME_OFF:4>0100
	TIME_ON                    ADIField = "TIME_ON"                    // Record Field: HHMM or HHMMSS in UTC
	TX_PWR                     ADIField = "TX_PWR"                     // Record Field: the logging station's power in Watts with a value greater than or equal to 0
	UKSMG                      ADIField = "UKSMG"                      // Record Field: the contacted station's UKSMG member number with a value greater than 0
	USACA_COUNTIES             ADIField = "USACA_COUNTIES"             // Record Field: two US counties in the case where the contacted station is located on a border between two counties, representing counties credited to the QSO for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire
	VE_PROV                    ADIField = "VE_PROV"                    // Deprecated: Record Field: import-only: use STATE instead
	VUCC_GRIDS                 ADIField = "VUCC_GRIDS"                 // Record Field: two or four adjacent Maidenhead grid locators, each four or six characters long, representing the contacted station's grid squares credited to the QSO for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07
	WEB                        ADIField = "WEB"                        // Record Field: the contacted station's URL
	WWFF_REF                   ADIField = "WWFF_REF"                   // Record Field: the contacted station's WWFF (World Wildlife Flora & Fauna) reference
)

// IsValid returns true if the ADIField exists in the ADIF specification. This includes Import Only and Deleted values.
func IsValid(value ADIField) bool {
	switch value {
	case ADDRESS:
		return true
	case ADDRESS_INTL:
		return true
	case ADIF_VER:
		return true
	case AGE:
		return true
	case ALTITUDE:
		return true
	case ANT_AZ:
		return true
	case ANT_EL:
		return true
	case ANT_PATH:
		return true
	case ARRL_SECT:
		return true
	case AWARD_GRANTED:
		return true
	case AWARD_SUBMITTED:
		return true
	case A_INDEX:
		return true
	case BAND:
		return true
	case BAND_RX:
		return true
	case CALL:
		return true
	case CHECK:
		return true
	case CLASS:
		return true
	case CLUBLOG_QSO_UPLOAD_DATE:
		return true
	case CLUBLOG_QSO_UPLOAD_STATUS:
		return true
	case CNTY:
		return true
	case CNTY_ALT:
		return true
	case COMMENT:
		return true
	case COMMENT_INTL:
		return true
	case CONT:
		return true
	case CONTACTED_OP:
		return true
	case CONTEST_ID:
		return true
	case COUNTRY:
		return true
	case COUNTRY_INTL:
		return true
	case CQZ:
		return true
	case CREATED_TIMESTAMP:
		return true
	case CREDIT_GRANTED:
		return true
	case CREDIT_SUBMITTED:
		return true
	case DARC_DOK:
		return true
	case DCL_QSLRDATE:
		return true
	case DCL_QSLSDATE:
		return true
	case DCL_QSL_RCVD:
		return true
	case DCL_QSL_SENT:
		return true
	case DISTANCE:
		return true
	case DXCC:
		return true
	case EMAIL:
		return true
	case EQSL_AG:
		return true
	case EQSL_QSLRDATE:
		return true
	case EQSL_QSLSDATE:
		return true
	case EQSL_QSL_RCVD:
		return true
	case EQSL_QSL_SENT:
		return true
	case EQ_CALL:
		return true
	case FISTS:
		return true
	case FISTS_CC:
		return true
	case FORCE_INIT:
		return true
	case FREQ:
		return true
	case FREQ_RX:
		return true
	case GRIDSQUARE:
		return true
	case GRIDSQUARE_EXT:
		return true
	case GUEST_OP:
		return true
	case HAMLOGEU_QSO_UPLOAD_DATE:
		return true
	case HAMLOGEU_QSO_UPLOAD_STATUS:
		return true
	case HAMQTH_QSO_UPLOAD_DATE:
		return true
	case HAMQTH_QSO_UPLOAD_STATUS:
		return true
	case HRDLOG_QSO_UPLOAD_DATE:
		return true
	case HRDLOG_QSO_UPLOAD_STATUS:
		return true
	case IOTA:
		return true
	case IOTA_ISLAND_ID:
		return true
	case ITUZ:
		return true
	case K_INDEX:
		return true
	case LAT:
		return true
	case LON:
		return true
	case LOTW_QSLRDATE:
		return true
	case LOTW_QSLSDATE:
		return true
	case LOTW_QSL_RCVD:
		return true
	case LOTW_QSL_SENT:
		return true
	case MAX_BURSTS:
		return true
	case MODE:
		return true
	case MORSE_KEY_INFO:
		return true
	case MORSE_KEY_TYPE:
		return true
	case MS_SHOWER:
		return true
	case MY_ALTITUDE:
		return true
	case MY_ANTENNA:
		return true
	case MY_ANTENNA_INTL:
		return true
	case MY_ARRL_SECT:
		return true
	case MY_CITY:
		return true
	case MY_CITY_INTL:
		return true
	case MY_CNTY:
		return true
	case MY_CNTY_ALT:
		return true
	case MY_COUNTRY:
		return true
	case MY_COUNTRY_INTL:
		return true
	case MY_CQ_ZONE:
		return true
	case MY_DARC_DOK:
		return true
	case MY_DXCC:
		return true
	case MY_FISTS:
		return true
	case MY_GRIDSQUARE:
		return true
	case MY_GRIDSQUARE_EXT:
		return true
	case MY_IOTA:
		return true
	case MY_IOTA_ISLAND_ID:
		return true
	case MY_ITU_ZONE:
		return true
	case MY_LAT:
		return true
	case MY_LON:
		return true
	case MY_MORSE_KEY_INFO:
		return true
	case MY_MORSE_KEY_TYPE:
		return true
	case MY_NAME:
		return true
	case MY_NAME_INTL:
		return true
	case MY_POSTAL_CODE:
		return true
	case MY_POSTAL_CODE_INTL:
		return true
	case MY_POTA_REF:
		return true
	case MY_RIG:
		return true
	case MY_RIG_INTL:
		return true
	case MY_SIG:
		return true
	case MY_SIG_INFO:
		return true
	case MY_SIG_INFO_INTL:
		return true
	case MY_SIG_INTL:
		return true
	case MY_SOTA_REF:
		return true
	case MY_STATE:
		return true
	case MY_STREET:
		return true
	case MY_STREET_INTL:
		return true
	case MY_USACA_COUNTIES:
		return true
	case MY_VUCC_GRIDS:
		return true
	case MY_WWFF_REF:
		return true
	case NAME:
		return true
	case NAME_INTL:
		return true
	case NOTES:
		return true
	case NOTES_INTL:
		return true
	case NR_BURSTS:
		return true
	case NR_PINGS:
		return true
	case OPERATOR:
		return true
	case OWNER_CALLSIGN:
		return true
	case PFX:
		return true
	case POTA_REF:
		return true
	case PRECEDENCE:
		return true
	case PROGRAMID:
		return true
	case PROGRAMVERSION:
		return true
	case PROP_MODE:
		return true
	case PUBLIC_KEY:
		return true
	case QRZCOM_QSO_DOWNLOAD_DATE:
		return true
	case QRZCOM_QSO_DOWNLOAD_STATUS:
		return true
	case QRZCOM_QSO_UPLOAD_DATE:
		return true
	case QRZCOM_QSO_UPLOAD_STATUS:
		return true
	case QSLMSG:
		return true
	case QSLMSG_INTL:
		return true
	case QSLMSG_RCVD:
		return true
	case QSLRDATE:
		return true
	case QSLSDATE:
		return true
	case QSL_RCVD:
		return true
	case QSL_RCVD_VIA:
		return true
	case QSL_SENT:
		return true
	case QSL_SENT_VIA:
		return true
	case QSL_VIA:
		return true
	case QSO_COMPLETE:
		return true
	case QSO_DATE:
		return true
	case QSO_DATE_OFF:
		return true
	case QSO_RANDOM:
		return true
	case QTH:
		return true
	case QTH_INTL:
		return true
	case REGION:
		return true
	case RIG:
		return true
	case RIG_INTL:
		return true
	case RST_RCVD:
		return true
	case RST_SENT:
		return true
	case RX_PWR:
		return true
	case SAT_MODE:
		return true
	case SAT_NAME:
		return true
	case SFI:
		return true
	case SIG:
		return true
	case SIG_INFO:
		return true
	case SIG_INFO_INTL:
		return true
	case SIG_INTL:
		return true
	case SILENT_KEY:
		return true
	case SKCC:
		return true
	case SOTA_REF:
		return true
	case SRX:
		return true
	case SRX_STRING:
		return true
	case STATE:
		return true
	case STATION_CALLSIGN:
		return true
	case STX:
		return true
	case STX_STRING:
		return true
	case SUBMODE:
		return true
	case SWL:
		return true
	case TEN_TEN:
		return true
	case TIME_OFF:
		return true
	case TIME_ON:
		return true
	case TX_PWR:
		return true
	case UKSMG:
		return true
	case USACA_COUNTIES:
		return true
	case VE_PROV:
		return true
	case VUCC_GRIDS:
		return true
	case WEB:
		return true
	case WWFF_REF:
		return true
	}
	return false
}
