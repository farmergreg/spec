// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package adifield

var (
	ADIF_VER           ADIField = "ADIF_VER"           // Header Field: identifies the version of ADIF used in this file in the format X.Y.Z where X is an integer designating the ADIF epoch Y is an integer between 0 and 9 designating the major version Z is an integer between 0 and 9 designating the minor version
	APP_LOTW_LASTQSL   ADIField = "APP_LOTW_LASTQSL"   // Header Field: ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl="yes". Date and time at which the most recent QSL record in this report was received. This will be the QSL date/time of the first QSO record in the file. This value only applies to the records in this file.
	APP_LOTW_LASTQSORX ADIField = "APP_LOTW_LASTQSORX" // Header Field: ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl="no". Date and time at which the most recent QSO record in this report was received. This will be the QSO received (uploaded) date/time of the first QSO record in the file. This value only applies to the records in this file.
	APP_LOTW_NUMREC    ADIField = "APP_LOTW_NUMREC"    // Header Field: ARRL LOTW: Number of QSO records in this download
	CREATED_TIMESTAMP  ADIField = "CREATED_TIMESTAMP"  // Header Field: identifies the UTC date and time that the file was created in the format of 15 characters YYYYMMDD HHMMSS where YYYYMMDD is a Date data type HHMMSS is a 6 character Time data type
	PROGRAMID          ADIField = "PROGRAMID"          // Header Field: identifies the name of the logger, converter, or utility that created or processed this ADIF file To help avoid name clashes, the ADIF PROGRAMID Register provides a voluntary list of PROGRAMID values.
	PROGRAMVERSION     ADIField = "PROGRAMVERSION"     // Header Field: identifies the version of the logger, converter, or utility that created or processed this ADIF file
	USERDEF1           ADIField = "USERDEF1"           // Header Field: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF2           ADIField = "USERDEF2"           // Header Field: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF3           ADIField = "USERDEF3"           // Header Field: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF4           ADIField = "USERDEF4"           // Header Field: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF5           ADIField = "USERDEF5"           // Header Field: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF6           ADIField = "USERDEF6"           // Header Field: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF7           ADIField = "USERDEF7"           // Header Field: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF8           ADIField = "USERDEF8"           // Header Field: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF9           ADIField = "USERDEF9"           // Header Field: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
)

var (
	ADDRESS                        ADIField = "ADDRESS"                        // Record Field: the contacted station's complete mailing address: full name, street address, city, postal code, and country
	ADDRESS_INTL                   ADIField = "ADDRESS_INTL"                   // Record Field: the contacted station's complete mailing address: full name, street address, city, postal code, and country
	AGE                            ADIField = "AGE"                            // Record Field: the contacted station's operator's age in years in the range 0 to 120 (inclusive)
	ALTITUDE                       ADIField = "ALTITUDE"                       // Record Field: the height of the contacted station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <ALTITUDE:4>1500 and 10.5 m is <ALTITUDE:4>10.5
	ANT_AZ                         ADIField = "ANT_AZ"                         // Record Field: the logging station's antenna azimuth, in degrees with a value between 0 to 360 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 370 is exported as 10). True north is 0 degrees with values increasing in a clockwise direction.
	ANT_EL                         ADIField = "ANT_EL"                         // Record Field: the logging station's antenna elevation, in degrees with a value between -90 to 90 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 100 is exported as 80). The horizon is 0 degrees with values increasing as the angle moves in an upward direction.
	ANT_PATH                       ADIField = "ANT_PATH"                       // Record Field: the signal path
	APP_LOTW_2XQSL                 ADIField = "APP_LOTW_2XQSL"                 // Record Field: ARRL LOTW: [Y, N] APP_LOTW_2xQSL ="Y" Indicates this confirmation is considered to be "two-way" in the QSO's Mode. It is a confirmation that counts towards the mode-specific awards & endorsement in the WAS program. Only present if QSL_RCVD="Y".
	APP_LOTW_CQZ_INFERRED          ADIField = "APP_LOTW_CQZ_INFERRED"          // Record Field: ARRL LOTW
	APP_LOTW_CQZ_INVALID           ADIField = "APP_LOTW_CQZ_INVALID"           // Record Field: ARRL LOTW
	APP_LOTW_CQZ_USERINVALID       ADIField = "APP_LOTW_CQZ_USERINVALID"       // Record Field: ARRL LOTW: Present only if the QSLing station included an invalid CQ zone value in its station location uploaded to LOTW
	APP_LOTW_CREDIT_GRANTED        ADIField = "APP_LOTW_CREDIT_GRANTED"        // Record Field: ARRL LOTW: Award credits granted based upon this QSL. Awards per LOTW award enumeration.
	APP_LOTW_CREDIT_SUBMITTED      ADIField = "APP_LOTW_CREDIT_SUBMITTED"      // Record Field: ARRL LOTW: Award credits for which this QSL has been submitted on an award application to sponsor. Awards per LOTW award enumeration.
	APP_LOTW_DXCC_ENTITY_STATUS    ADIField = "APP_LOTW_DXCC_ENTITY_STATUS"    // Record Field: ARRL LOTW: [Current, Deleted] Indicates whether the ARRL DXCC entity is "Current" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or "Deleted"
	APP_LOTW_EOF                   ADIField = "APP_LOTW_EOF"                   // Record Field: ARRL LOTW: End of File marker. No other fields expected in this record.
	APP_LOTW_GRIDSQUARE_INVALID    ADIField = "APP_LOTW_GRIDSQUARE_INVALID"    // Record Field: ARRL LOTW
	APP_LOTW_ITUZ_INFERRED         ADIField = "APP_LOTW_ITUZ_INFERRED"         // Record Field: ARRL LOTW
	APP_LOTW_ITUZ_INVALID          ADIField = "APP_LOTW_ITUZ_INVALID"          // Record Field: ARRL LOTW
	APP_LOTW_ITUZ_USERINVALID      ADIField = "APP_LOTW_ITUZ_USERINVALID"      // Record Field: ARRL LOTW: Present only if the QSLing station included an invalid ITU zone value in its station location uploaded to LOTW
	APP_LOTW_MODE                  ADIField = "APP_LOTW_MODE"                  // Record Field: ARRL LOTW: Field is present when the mode recorded in Logbook cannot be unambiguously represented as an ADIF-compliant value
	APP_LOTW_MODEGROUP             ADIField = "APP_LOTW_MODEGROUP"             // Record Field: ARRL LOTW: [CW, PHONE, DATA] The mode group indicates whether the QSO's mode counts towards the CW, Phone or Digital awards respectively in the DXCC program. A mode group of CW indicates the QSO's mode counts towards the CW endorsement on WAS awards. A mode group of PHONE indicates the QSO's mode counts towards the WAS Phone award. A mode group of DATA indicates the QSO's mode counts towards the CQ WPX Digital award.
	APP_LOTW_MY_DXCC_ENTITY_STATUS ADIField = "APP_LOTW_MY_DXCC_ENTITY_STATUS" // Record Field: ARRL LOTW: [Current, Deleted] Indicates whether the Logging station's ARRL DXCC entity is "Current" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or "Deleted"
	APP_LOTW_NPSUNIT               ADIField = "APP_LOTW_NPSUNIT"               // Record Field: ARRL LOTW: Present only if the QSLing station included a valid NPOTA Park code or a comma-delimited list of valid NPOTA Park codes
	APP_LOTW_OWNCALL               ADIField = "APP_LOTW_OWNCALL"               // Record Field: ARRL LOTW: "own" call sign of the station making the contact. Present only if qso_withown="yes". DEPRECATED -- In the near future, this field will no longer be present in the report output. Programs that import LOTW report files should not rely on the presence of this field. Use STATION_CALLSIGN instead.
	APP_LOTW_QSLMODE               ADIField = "APP_LOTW_QSLMODE"               // Record Field: ARRL LOTW: The Mode that your QSO partner indicated on their side of this QSO. Only present if QSL_RCVD="Y" and APP_LOTW_2xQSL="N".
	APP_LOTW_QSO_TIMESTAMP         ADIField = "APP_LOTW_QSO_TIMESTAMP"         // Record Field: ARRL LOTW: Combined QSO Date & Time in ISO-8601 format
	APP_LOTW_RXQSL                 ADIField = "APP_LOTW_RXQSL"                 // Record Field: ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which the confirmation match (QSL) was made/updated at LOTW. Present only if QSL_RCVD="Y"
	APP_LOTW_RXQSO                 ADIField = "APP_LOTW_RXQSO"                 // Record Field: ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which QSO record was inserted/updated at LOTW
	APP_QRZLOG_LOGID               ADIField = "APP_QRZLOG_LOGID"               // Record Field: QRZ LOG: References the internal ID of a log record in the QRZ system.
	APP_QRZLOG_QSLDATE             ADIField = "APP_QRZLOG_QSLDATE"             // Record Field: QRZ LOG: QSL Date (e.g. 20210126) This field contains the first date on which we saw a QRZ accepted confirmation for this contact. This could have been a confirmation as a result of matching contacts in the QRZ system or a credit imported from Logbook of The World.
	APP_QRZLOG_STATUS              ADIField = "APP_QRZLOG_STATUS"              // Record Field: QRZ LOG: Status [C: Confirmed, A: Reserved for future use, N: Not Confirmed, 2: Confirmation Requested, S: Confirmation Requested, Request Seen, R: Confirmation Requested, Request Rejected]
	APP_SKCCLOGGER_KEYTYPE         ADIField = "APP_SKCCLOGGER_KEYTYPE"         // Record Field: Straight Key Century Club (SKCC): [BUG, SK, SS] CW KEY TYPE (SS=Side Swiper / Cootie, SK=Straight, B=Bug)
	ARRL_SECT                      ADIField = "ARRL_SECT"                      // Record Field: the contacted station's ARRL section
	AWARD_GRANTED                  ADIField = "AWARD_GRANTED"                  // Record Field: the list of awards granted by a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's "Award Account" between an award sponsor and the user. For example, in response to a request "send me a list of the awards granted to AA6YQ", this might be received: <CALL:5>AA6YQ <AWARD_GRANTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m
	AWARD_SUBMITTED                ADIField = "AWARD_SUBMITTED"                // Record Field: the list of awards submitted to a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's "Award Account" between an award sponsor and the user. For example, AA6YQ might submit a request for awards by sending the following: <CALL:5>AA6YQ <AWARD_SUBMITTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m
	A_INDEX                        ADIField = "A_INDEX"                        // Record Field: the geomagnetic A index at the time of the QSO in the range 0 to 400 (inclusive)
	BAND                           ADIField = "BAND"                           // Record Field: QSO Band
	BAND_RX                        ADIField = "BAND_RX"                        // Record Field: in a split frequency QSO, the logging station's receiving band
	CALL                           ADIField = "CALL"                           // Record Field: the contacted station's callsign
	CHECK                          ADIField = "CHECK"                          // Record Field: contest check (e.g. for ARRL Sweepstakes)
	CLASS                          ADIField = "CLASS"                          // Record Field: contest class (e.g. for ARRL Field Day)
	CLUBLOG_QSO_UPLOAD_DATE        ADIField = "CLUBLOG_QSO_UPLOAD_DATE"        // Record Field: the date the QSO was last uploaded to the Club Log online service
	CLUBLOG_QSO_UPLOAD_STATUS      ADIField = "CLUBLOG_QSO_UPLOAD_STATUS"      // Record Field: the upload status of the QSO on the Club Log online service
	CNTY                           ADIField = "CNTY"                           // Record Field: the contacted station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format
	CNTY_ALT                       ADIField = "CNTY_ALT"                       // Record Field: a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the contacted station See the Data Type for details.
	COMMENT                        ADIField = "COMMENT"                        // Record Field: comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG field recommended use: information of interest to the contacted station's operator
	COMMENT_INTL                   ADIField = "COMMENT_INTL"                   // Record Field: comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG_INTL field recommended use: information of interest to the contacted station's operator
	CONT                           ADIField = "CONT"                           // Record Field: the contacted station's Continent
	CONTACTED_OP                   ADIField = "CONTACTED_OP"                   // Record Field: the callsign of the individual operating the contacted station
	CONTEST_ID                     ADIField = "CONTEST_ID"                     // Record Field: QSO Contest Identifier use enumeration values for interoperability
	COUNTRY                        ADIField = "COUNTRY"                        // Record Field: the contacted station's DXCC entity name
	COUNTRY_INTL                   ADIField = "COUNTRY_INTL"                   // Record Field: the contacted station's DXCC entity name
	CQZ                            ADIField = "CQZ"                            // Record Field: the contacted station's CQ Zone in the range 1 to 40 (inclusive)
	CREDIT_GRANTED                 ADIField = "CREDIT_GRANTED"                 // Record Field: the list of credits granted to this QSO Use of data type AwardList and enumeration Award are import-only
	CREDIT_SUBMITTED               ADIField = "CREDIT_SUBMITTED"               // Record Field: the list of credits sought for this QSO Use of data type AwardList and enumeration Award are import-only
	DARC_DOK                       ADIField = "DARC_DOK"                       // Record Field: the contacted station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <DARC_DOK:3>A01
	DCL_QSLRDATE                   ADIField = "DCL_QSLRDATE"                   // Record Field: date QSL received from DCL (only valid if DCL_QSL_RCVD is Y, I, or V)(V import-only)
	DCL_QSLSDATE                   ADIField = "DCL_QSLSDATE"                   // Record Field: date QSL sent to DCL (only valid if DCL_QSL_SENT is Y, Q, or I)
	DCL_QSL_RCVD                   ADIField = "DCL_QSL_RCVD"                   // Record Field: DCL QSL received status Default Value: N
	DCL_QSL_SENT                   ADIField = "DCL_QSL_SENT"                   // Record Field: DCL QSL sent status Default Value: N
	DISTANCE                       ADIField = "DISTANCE"                       // Record Field: the distance between the logging station and the contacted station in kilometers via the specified signal path with a value greater than or equal to 0
	DXCC                           ADIField = "DXCC"                           // Record Field: the contacted station's DXCC Entity Code <DXCC:1>0 means that the contacted station is known not to be within a DXCC entity.
	EMAIL                          ADIField = "EMAIL"                          // Record Field: the contacted station's email address
	EQSL_AG                        ADIField = "EQSL_AG"                        // Record Field: indicates whether the QSO is known to be "Authenticity Guaranteed" by eQSL
	EQSL_QSLRDATE                  ADIField = "EQSL_QSLRDATE"                  // Record Field: date QSL received from eQSL.cc (only valid if EQSL_QSL_RCVD is Y, I, or V)(V import-only)
	EQSL_QSLSDATE                  ADIField = "EQSL_QSLSDATE"                  // Record Field: date QSL sent to eQSL.cc (only valid if EQSL_QSL_SENT is Y, Q, or I)
	EQSL_QSL_RCVD                  ADIField = "EQSL_QSL_RCVD"                  // Record Field: eQSL.cc QSL received status instead of V (import-only) use <CREDIT_GRANTED:42>CQWAZ:eqsl,CQWAZ_BAND:eqsl,CQWAZ_MODE:eqsl Default Value: N
	EQSL_QSL_SENT                  ADIField = "EQSL_QSL_SENT"                  // Record Field: eQSL.cc QSL sent status Default Value: N
	EQ_CALL                        ADIField = "EQ_CALL"                        // Record Field: the contacted station's owner's callsign
	FISTS                          ADIField = "FISTS"                          // Record Field: the contacted station's FISTS CW Club member number with a value greater than 0.
	FISTS_CC                       ADIField = "FISTS_CC"                       // Record Field: the contacted station's FISTS CW Club Century Certificate (CC) number with a value greater than 0.
	FORCE_INIT                     ADIField = "FORCE_INIT"                     // Record Field: new EME "initial"
	FREQ                           ADIField = "FREQ"                           // Record Field: QSO frequency in Megahertz
	FREQ_RX                        ADIField = "FREQ_RX"                        // Record Field: in a split frequency QSO, the logging station's receiving frequency in Megahertz
	GRIDSQUARE                     ADIField = "GRIDSQUARE"                     // Record Field: the contacted station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in GRIDSQUARE and the additional 2 or 4 characters in the GRIDSQUARE_EXT field
	GRIDSQUARE_EXT                 ADIField = "GRIDSQUARE_EXT"                 // Record Field: for a contacted station's 10-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9 and 10. For a contacted station's 12-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <GRIDSQUARE:8>FN01MH42 <GRIDSQUARE_EXT:2>BQ
	GUEST_OP                       ADIField = "GUEST_OP"                       // Deprecated: Record Field: import-only: use OPERATOR instead
	HAMLOGEU_QSO_UPLOAD_DATE       ADIField = "HAMLOGEU_QSO_UPLOAD_DATE"       // Record Field: the date the QSO was last uploaded to the HAMLOG.EU online service
	HAMLOGEU_QSO_UPLOAD_STATUS     ADIField = "HAMLOGEU_QSO_UPLOAD_STATUS"     // Record Field: the upload status of the QSO on the HAMLOG.EU online service
	HAMQTH_QSO_UPLOAD_DATE         ADIField = "HAMQTH_QSO_UPLOAD_DATE"         // Record Field: the date the QSO was last uploaded to the HamQTH.com online service
	HAMQTH_QSO_UPLOAD_STATUS       ADIField = "HAMQTH_QSO_UPLOAD_STATUS"       // Record Field: the upload status of the QSO on the HamQTH.com online service
	HRDLOG_QSO_UPLOAD_DATE         ADIField = "HRDLOG_QSO_UPLOAD_DATE"         // Record Field: the date the QSO was last uploaded to the HRDLog.net online service
	HRDLOG_QSO_UPLOAD_STATUS       ADIField = "HRDLOG_QSO_UPLOAD_STATUS"       // Record Field: the upload status of the QSO on the HRDLog.net online service
	IOTA                           ADIField = "IOTA"                           // Record Field: the contacted station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]
	IOTA_ISLAND_ID                 ADIField = "IOTA_ISLAND_ID"                 // Record Field: the contacted station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]
	ITUZ                           ADIField = "ITUZ"                           // Record Field: the contacted station's ITU zone in the range 1 to 90 (inclusive)
	K_INDEX                        ADIField = "K_INDEX"                        // Record Field: the geomagnetic K index at the time of the QSO in the range 0 to 9 (inclusive)
	LAT                            ADIField = "LAT"                            // Record Field: the contacted station's latitude
	LON                            ADIField = "LON"                            // Record Field: the contacted station's longitude
	LOTW_QSLRDATE                  ADIField = "LOTW_QSLRDATE"                  // Record Field: date QSL received from ARRL Logbook of the World (only valid if LOTW_QSL_RCVD is Y, I, or V)(V import-only)
	LOTW_QSLSDATE                  ADIField = "LOTW_QSLSDATE"                  // Record Field: date QSL sent to ARRL Logbook of the World (only valid if LOTW_QSL_SENT is Y, Q, or I)
	LOTW_QSL_RCVD                  ADIField = "LOTW_QSL_RCVD"                  // Record Field: ARRL Logbook of the World QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw Default Value: N
	LOTW_QSL_SENT                  ADIField = "LOTW_QSL_SENT"                  // Record Field: ARRL Logbook of the World QSL sent status Default Value: N
	MAX_BURSTS                     ADIField = "MAX_BURSTS"                     // Record Field: maximum length of meteor scatter bursts heard by the logging station, in seconds with a value greater than or equal to 0
	MODE                           ADIField = "MODE"                           // Record Field: QSO Mode
	MORSE_KEY_INFO                 ADIField = "MORSE_KEY_INFO"                 // Record Field: details of the contacted station's Morse key (e.g. make, model, etc). Example: <MORSE_KEY_INFO:16>Begali Sculpture
	MORSE_KEY_TYPE                 ADIField = "MORSE_KEY_TYPE"                 // Record Field: the contacted station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP
	MS_SHOWER                      ADIField = "MS_SHOWER"                      // Record Field: For Meteor Scatter QSOs, the name of the meteor shower in progress
	MY_ALTITUDE                    ADIField = "MY_ALTITUDE"                    // Record Field: the height of the logging station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <MY_ALTITUDE:4>1500 and 10.5 m is <MY_ALTITUDE:4>10.5
	MY_ANTENNA                     ADIField = "MY_ANTENNA"                     // Record Field: the logging station's antenna
	MY_ANTENNA_INTL                ADIField = "MY_ANTENNA_INTL"                // Record Field: the logging station's antenna
	MY_ARRL_SECT                   ADIField = "MY_ARRL_SECT"                   // Record Field: the logging station's ARRL section
	MY_CITY                        ADIField = "MY_CITY"                        // Record Field: the logging station's city
	MY_CITY_INTL                   ADIField = "MY_CITY_INTL"                   // Record Field: the logging station's city
	MY_CNTY                        ADIField = "MY_CNTY"                        // Record Field: the logging station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format
	MY_CNTY_ALT                    ADIField = "MY_CNTY_ALT"                    // Record Field: a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the logging station See the Data Type for details.
	MY_COUNTRY                     ADIField = "MY_COUNTRY"                     // Record Field: the logging station's DXCC entity name
	MY_COUNTRY_INTL                ADIField = "MY_COUNTRY_INTL"                // Record Field: the logging station's DXCC entity name
	MY_CQ_ZONE                     ADIField = "MY_CQ_ZONE"                     // Record Field: the logging station's CQ Zone in the range 1 to 40 (inclusive)
	MY_DARC_DOK                    ADIField = "MY_DARC_DOK"                    // Record Field: the logging station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <MY_DARC_DOK:3>A01
	MY_DXCC                        ADIField = "MY_DXCC"                        // Record Field: the logging station's DXCC Entity Code <MY_DXCC:1>0 means that the logging station is known not to be within a DXCC entity.
	MY_FISTS                       ADIField = "MY_FISTS"                       // Record Field: the logging station's FISTS CW Club member number with a value greater than 0.
	MY_GRIDSQUARE                  ADIField = "MY_GRIDSQUARE"                  // Record Field: the logging station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in MY_GRIDSQUARE and the additional 2 or 4 characters in the MY_GRIDSQUARE_EXT field
	MY_GRIDSQUARE_EXT              ADIField = "MY_GRIDSQUARE_EXT"              // Record Field: for a logging station's 10-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9 and 10. For a logging station's 12-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <MY_GRIDSQUARE:8>FN01MH42 <MY_GRIDSQUARE_EXT:2>BQ
	MY_IOTA                        ADIField = "MY_IOTA"                        // Record Field: the logging station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]
	MY_IOTA_ISLAND_ID              ADIField = "MY_IOTA_ISLAND_ID"              // Record Field: the logging station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]
	MY_ITU_ZONE                    ADIField = "MY_ITU_ZONE"                    // Record Field: the logging station's ITU zone in the range 1 to 90 (inclusive)
	MY_LAT                         ADIField = "MY_LAT"                         // Record Field: the logging station's latitude
	MY_LON                         ADIField = "MY_LON"                         // Record Field: the logging station's longitude
	MY_MORSE_KEY_INFO              ADIField = "MY_MORSE_KEY_INFO"              // Record Field: details of the logging station's Morse key (e.g. make, model, etc). Example: <MY_MORSE_KEY_INFO:16>Begali Sculpture
	MY_MORSE_KEY_TYPE              ADIField = "MY_MORSE_KEY_TYPE"              // Record Field: the logging station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP
	MY_NAME                        ADIField = "MY_NAME"                        // Record Field: the logging operator's name
	MY_NAME_INTL                   ADIField = "MY_NAME_INTL"                   // Record Field: the logging operator's name
	MY_POSTAL_CODE                 ADIField = "MY_POSTAL_CODE"                 // Record Field: the logging station's postal code
	MY_POSTAL_CODE_INTL            ADIField = "MY_POSTAL_CODE_INTL"            // Record Field: the logging station's postal code
	MY_POTA_REF                    ADIField = "MY_POTA_REF"                    // Record Field: a comma-delimited list of one or more of the logging station's POTA (Parks on the Air) reference(s). Examples: <MY_POTA_REF:6>K-0059 <MY_POTA_REF:7>K-10000 <MY_POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY
	MY_RIG                         ADIField = "MY_RIG"                         // Record Field: description of the logging station's equipment
	MY_RIG_INTL                    ADIField = "MY_RIG_INTL"                    // Record Field: description of the logging station's equipment
	MY_SIG                         ADIField = "MY_SIG"                         // Record Field: special interest activity or event
	MY_SIG_INFO                    ADIField = "MY_SIG_INFO"                    // Record Field: special interest activity or event information
	MY_SIG_INFO_INTL               ADIField = "MY_SIG_INFO_INTL"               // Record Field: special interest activity or event information
	MY_SIG_INTL                    ADIField = "MY_SIG_INTL"                    // Record Field: special interest activity or event
	MY_SOTA_REF                    ADIField = "MY_SOTA_REF"                    // Record Field: the logging station's International SOTA Reference.
	MY_STATE                       ADIField = "MY_STATE"                       // Record Field: the code for the logging station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)
	MY_STREET                      ADIField = "MY_STREET"                      // Record Field: the logging station's street
	MY_STREET_INTL                 ADIField = "MY_STREET_INTL"                 // Record Field: the logging station's street
	MY_USACA_COUNTIES              ADIField = "MY_USACA_COUNTIES"              // Record Field: two US counties in the case where the logging station is located on a border between two counties, representing counties that the contacted station may claim for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire
	MY_VUCC_GRIDS                  ADIField = "MY_VUCC_GRIDS"                  // Record Field: two or four adjacent Maidenhead grid locators, each four or six characters long, representing the logging station's grid squares that the contacted station may claim for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07
	MY_WWFF_REF                    ADIField = "MY_WWFF_REF"                    // Record Field: the logging station's WWFF (World Wildlife Flora & Fauna) reference
	NAME                           ADIField = "NAME"                           // Record Field: the contacted station's operator's name
	NAME_INTL                      ADIField = "NAME_INTL"                      // Record Field: the contacted station's operator's name
	NOTES                          ADIField = "NOTES"                          // Record Field: QSO notes recommended use: information of interest to the logging station's operator
	NOTES_INTL                     ADIField = "NOTES_INTL"                     // Record Field: QSO notes recommended use: information of interest to the logging station's operator
	NR_BURSTS                      ADIField = "NR_BURSTS"                      // Record Field: the number of meteor scatter bursts heard by the logging station with a value greater than or equal to 0
	NR_PINGS                       ADIField = "NR_PINGS"                       // Record Field: the number of meteor scatter pings heard by the logging station with a value greater than or equal to 0
	OPERATOR                       ADIField = "OPERATOR"                       // Record Field: the logging operator's callsign if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign
	OWNER_CALLSIGN                 ADIField = "OWNER_CALLSIGN"                 // Record Field: the callsign of the owner of the station used to log the contact (the callsign of the OPERATOR's host) if OWNER_CALLSIGN is absent, STATION_CALLSIGN shall be treated as both the logging station's callsign and the callsign of the owner of the station
	PFX                            ADIField = "PFX"                            // Record Field: the contacted station's WPX prefix
	POTA_REF                       ADIField = "POTA_REF"                       // Record Field: a comma-delimited list of one or more of the contacted station's POTA (Parks on the Air) reference(s). Examples: <POTA_REF:6>K-5033 <POTA_REF:13>VE-5082@CA-AB <POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY
	PRECEDENCE                     ADIField = "PRECEDENCE"                     // Record Field: contest precedence (e.g. for ARRL Sweepstakes)
	PROP_MODE                      ADIField = "PROP_MODE"                      // Record Field: QSO propagation mode
	PUBLIC_KEY                     ADIField = "PUBLIC_KEY"                     // Record Field: public encryption key
	QRZCOM_QSO_DOWNLOAD_DATE       ADIField = "QRZCOM_QSO_DOWNLOAD_DATE"       // Record Field: date QSO downloaded from QRZ.COM logbook
	QRZCOM_QSO_DOWNLOAD_STATUS     ADIField = "QRZCOM_QSO_DOWNLOAD_STATUS"     // Record Field: QRZ.COM logbook QSO download status
	QRZCOM_QSO_UPLOAD_DATE         ADIField = "QRZCOM_QSO_UPLOAD_DATE"         // Record Field: the date the QSO was last uploaded to the QRZ.COM online service
	QRZCOM_QSO_UPLOAD_STATUS       ADIField = "QRZCOM_QSO_UPLOAD_STATUS"       // Record Field: the upload status of the QSO on the QRZ.COM online service
	QSLMSG                         ADIField = "QSLMSG"                         // Record Field: a message for the contacted station's operator to be incorporated in a paper or electronic QSL
	QSLMSG_INTL                    ADIField = "QSLMSG_INTL"                    // Record Field: a message for the contacted station's operator to be incorporated in a paper or electronic QSL
	QSLMSG_RCVD                    ADIField = "QSLMSG_RCVD"                    // Record Field: a message addressed to the logging station's operator incorporated in a paper or electronic QSL
	QSLRDATE                       ADIField = "QSLRDATE"                       // Record Field: QSL received date (only valid if QSL_RCVD is Y, I, or V)(V import-only)
	QSLSDATE                       ADIField = "QSLSDATE"                       // Record Field: QSL sent date (only valid if QSL_SENT is Y, Q, or I)
	QSL_RCVD                       ADIField = "QSL_RCVD"                       // Record Field: QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card Default Value: N
	QSL_RCVD_VIA                   ADIField = "QSL_RCVD_VIA"                   // Record Field: if QSL_RCVD is set to 'Y' or 'V', the means by which the QSL was received by the logging station; otherwise, the means by which the logging station requested or intends to request that the QSL be conveyed. (Note: 'V' is import-only) use of M (manager) is import-only
	QSL_SENT                       ADIField = "QSL_SENT"                       // Record Field: QSL sent status Default Value: N
	QSL_SENT_VIA                   ADIField = "QSL_SENT_VIA"                   // Record Field: if QSL_SENT is set to 'Y', the means by which the QSL was sent by the logging station; otherwise, the means by which the logging station intends to convey the QSL use of M (manager) is import-only
	QSL_VIA                        ADIField = "QSL_VIA"                        // Record Field: the contacted station's QSL route
	QSO_COMPLETE                   ADIField = "QSO_COMPLETE"                   // Record Field: indicates whether the QSO was complete from the perspective of the logging station Y - yes N - no NIL - not heard ? - uncertain
	QSO_DATE                       ADIField = "QSO_DATE"                       // Record Field: date on which the QSO started
	QSO_DATE_OFF                   ADIField = "QSO_DATE_OFF"                   // Record Field: date on which the QSO ended
	QSO_RANDOM                     ADIField = "QSO_RANDOM"                     // Record Field: indicates whether the QSO was random or scheduled
	QTH                            ADIField = "QTH"                            // Record Field: the contacted station's city
	QTH_INTL                       ADIField = "QTH_INTL"                       // Record Field: the contacted station's city
	REGION                         ADIField = "REGION"                         // Record Field: the contacted station's WAE or CQ entity contained within a DXCC entity. the value None indicates that the WAE or CQ entity is the DXCC entity in the DXCC field. nothing can be inferred from the absence of the REGION field
	RIG                            ADIField = "RIG"                            // Record Field: description of the contacted station's equipment
	RIG_INTL                       ADIField = "RIG_INTL"                       // Record Field: description of the contacted station's equipment
	RST_RCVD                       ADIField = "RST_RCVD"                       // Record Field: signal report from the contacted station
	RST_SENT                       ADIField = "RST_SENT"                       // Record Field: signal report sent to the contacted station
	RX_PWR                         ADIField = "RX_PWR"                         // Record Field: the contacted station's transmitter power in Watts with a value greater than or equal to 0
	SAT_MODE                       ADIField = "SAT_MODE"                       // Record Field: satellite mode - a code representing the satellite's uplink band and downlink band
	SAT_NAME                       ADIField = "SAT_NAME"                       // Record Field: name of satellite
	SFI                            ADIField = "SFI"                            // Record Field: the solar flux at the time of the QSO in the range 0 to 300 (inclusive).
	SIG                            ADIField = "SIG"                            // Record Field: the name of the contacted station's special activity or interest group
	SIG_INFO                       ADIField = "SIG_INFO"                       // Record Field: information associated with the contacted station's activity or interest group
	SIG_INFO_INTL                  ADIField = "SIG_INFO_INTL"                  // Record Field: information associated with the contacted station's activity or interest group
	SIG_INTL                       ADIField = "SIG_INTL"                       // Record Field: the name of the contacted station's special activity or interest group
	SILENT_KEY                     ADIField = "SILENT_KEY"                     // Record Field: 'Y' indicates that the contacted station's operator is now a Silent Key.
	SKCC                           ADIField = "SKCC"                           // Record Field: the contacted station's Straight Key Century Club (SKCC) member information
	SOTA_REF                       ADIField = "SOTA_REF"                       // Record Field: the contacted station's International SOTA Reference.
	SRX                            ADIField = "SRX"                            // Record Field: contest QSO received serial number with a value greater than or equal to 0
	SRX_STRING                     ADIField = "SRX_STRING"                     // Record Field: contest QSO received information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail
	STATE                          ADIField = "STATE"                          // Record Field: the code for the contacted station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)
	STATION_CALLSIGN               ADIField = "STATION_CALLSIGN"               // Record Field: the logging station's callsign (the callsign used over the air) if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign
	STX                            ADIField = "STX"                            // Record Field: contest QSO transmitted serial number with a value greater than or equal to 0
	STX_STRING                     ADIField = "STX_STRING"                     // Record Field: contest QSO transmitted information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail
	SUBMODE                        ADIField = "SUBMODE"                        // Record Field: QSO Submode use enumeration values for interoperability
	SWL                            ADIField = "SWL"                            // Record Field: indicates that the QSO information pertains to an SWL report
	TEN_TEN                        ADIField = "TEN_TEN"                        // Record Field: Ten-Ten number with a value greater than 0
	TIME_OFF                       ADIField = "TIME_OFF"                       // Record Field: HHMM or HHMMSS in UTC in the absence of <QSO_DATE_OFF>, the QSO duration is less than 24 hours. For example, the following is a QSO starting at 14 July 2020 23:55 and finishing at 15 July 2020 01:00: <QSO_DATE:8>20200714 <TIME_ON:4>2355 <TIME_OFF:4>0100
	TIME_ON                        ADIField = "TIME_ON"                        // Record Field: HHMM or HHMMSS in UTC
	TX_PWR                         ADIField = "TX_PWR"                         // Record Field: the logging station's power in Watts with a value greater than or equal to 0
	UKSMG                          ADIField = "UKSMG"                          // Record Field: the contacted station's UKSMG member number with a value greater than 0
	USACA_COUNTIES                 ADIField = "USACA_COUNTIES"                 // Record Field: two US counties in the case where the contacted station is located on a border between two counties, representing counties credited to the QSO for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire
	VE_PROV                        ADIField = "VE_PROV"                        // Deprecated: Record Field: import-only: use STATE instead
	VUCC_GRIDS                     ADIField = "VUCC_GRIDS"                     // Record Field: two or four adjacent Maidenhead grid locators, each four or six characters long, representing the contacted station's grid squares credited to the QSO for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07
	WEB                            ADIField = "WEB"                            // Record Field: the contacted station's URL
	WWFF_REF                       ADIField = "WWFF_REF"                       // Record Field: the contacted station's WWFF (World Wildlife Flora & Fauna) reference
)

// A map of all ADIField specifications.
var ADIFieldMap = map[ADIField]Spec{
	ADDRESS:                        {Key: "ADDRESS", DataType: "MultilineString", Enumeration: "", Description: "the contacted station's complete mailing address: full name, street address, city, postal code, and country", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	ADDRESS_INTL:                   {Key: "ADDRESS_INTL", DataType: "IntlMultilineString", Enumeration: "", Description: "the contacted station's complete mailing address: full name, street address, city, postal code, and country", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	ADIF_VER:                       {Key: "ADIF_VER", DataType: "String", Enumeration: "", Description: "identifies the version of ADIF used in this file in the format X.Y.Z where X is an integer designating the ADIF epoch Y is an integer between 0 and 9 designating the major version Z is an integer between 0 and 9 designating the minor version", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	AGE:                            {Key: "AGE", DataType: "Number", Enumeration: "", Description: "the contacted station's operator's age in years in the range 0 to 120 (inclusive)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 120, IsImportOnly: false, Comments: ""},
	ALTITUDE:                       {Key: "ALTITUDE", DataType: "Number", Enumeration: "", Description: "the height of the contacted station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <ALTITUDE:4>1500 and 10.5 m is <ALTITUDE:4>10.5", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	ANT_AZ:                         {Key: "ANT_AZ", DataType: "Number", Enumeration: "", Description: "the logging station's antenna azimuth, in degrees with a value between 0 to 360 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 370 is exported as 10). True north is 0 degrees with values increasing in a clockwise direction.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 360, IsImportOnly: false, Comments: ""},
	ANT_EL:                         {Key: "ANT_EL", DataType: "Number", Enumeration: "", Description: "the logging station's antenna elevation, in degrees with a value between -90 to 90 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 100 is exported as 80). The horizon is 0 degrees with values increasing as the angle moves in an upward direction.", IsHeaderField: "", MinimumValue: -90, MaximumValue: 90, IsImportOnly: false, Comments: ""},
	ANT_PATH:                       {Key: "ANT_PATH", DataType: "Enumeration", Enumeration: "Ant_Path", Description: "the signal path", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_2XQSL:                 {Key: "APP_LOTW_2XQSL", DataType: "Enumeration", Enumeration: "", Description: "ARRL LOTW: [Y, N] APP_LOTW_2xQSL =\"Y\" Indicates this confirmation is considered to be \"two-way\" in the QSO's Mode. It is a confirmation that counts towards the mode-specific awards & endorsement in the WAS program. Only present if QSL_RCVD=\"Y\".", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_CQZ_INFERRED:          {Key: "APP_LOTW_CQZ_INFERRED", DataType: "String", Enumeration: "", Description: "ARRL LOTW", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_CQZ_INVALID:           {Key: "APP_LOTW_CQZ_INVALID", DataType: "String", Enumeration: "", Description: "ARRL LOTW", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_CQZ_USERINVALID:       {Key: "APP_LOTW_CQZ_USERINVALID", DataType: "String", Enumeration: "", Description: "ARRL LOTW: Present only if the QSLing station included an invalid CQ zone value in its station location uploaded to LOTW", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_CREDIT_GRANTED:        {Key: "APP_LOTW_CREDIT_GRANTED", DataType: "String", Enumeration: "", Description: "ARRL LOTW: Award credits granted based upon this QSL. Awards per LOTW award enumeration.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_CREDIT_SUBMITTED:      {Key: "APP_LOTW_CREDIT_SUBMITTED", DataType: "String", Enumeration: "", Description: "ARRL LOTW: Award credits for which this QSL has been submitted on an award application to sponsor. Awards per LOTW award enumeration.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_DXCC_ENTITY_STATUS:    {Key: "APP_LOTW_DXCC_ENTITY_STATUS", DataType: "Enumeration", Enumeration: "", Description: "ARRL LOTW: [Current, Deleted] Indicates whether the ARRL DXCC entity is \"Current\" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or \"Deleted\"", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_EOF:                   {Key: "APP_LOTW_EOF", DataType: "String", Enumeration: "", Description: "ARRL LOTW: End of File marker. No other fields expected in this record.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_GRIDSQUARE_INVALID:    {Key: "APP_LOTW_GRIDSQUARE_INVALID", DataType: "String", Enumeration: "", Description: "ARRL LOTW", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_ITUZ_INFERRED:         {Key: "APP_LOTW_ITUZ_INFERRED", DataType: "String", Enumeration: "", Description: "ARRL LOTW", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_ITUZ_INVALID:          {Key: "APP_LOTW_ITUZ_INVALID", DataType: "String", Enumeration: "", Description: "ARRL LOTW", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_ITUZ_USERINVALID:      {Key: "APP_LOTW_ITUZ_USERINVALID", DataType: "String", Enumeration: "", Description: "ARRL LOTW: Present only if the QSLing station included an invalid ITU zone value in its station location uploaded to LOTW", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_LASTQSL:               {Key: "APP_LOTW_LASTQSL", DataType: "String", Enumeration: "", Description: "ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl=\"yes\". Date and time at which the most recent QSL record in this report was received. This will be the QSL date/time of the first QSO record in the file. This value only applies to the records in this file.", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_LASTQSORX:             {Key: "APP_LOTW_LASTQSORX", DataType: "String", Enumeration: "", Description: "ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl=\"no\". Date and time at which the most recent QSO record in this report was received. This will be the QSO received (uploaded) date/time of the first QSO record in the file. This value only applies to the records in this file.", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_MODE:                  {Key: "APP_LOTW_MODE", DataType: "Enumeration", Enumeration: "", Description: "ARRL LOTW: Field is present when the mode recorded in Logbook cannot be unambiguously represented as an ADIF-compliant value", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_MODEGROUP:             {Key: "APP_LOTW_MODEGROUP", DataType: "Enumeration", Enumeration: "", Description: "ARRL LOTW: [CW, PHONE, DATA] The mode group indicates whether the QSO's mode counts towards the CW, Phone or Digital awards respectively in the DXCC program. A mode group of CW indicates the QSO's mode counts towards the CW endorsement on WAS awards. A mode group of PHONE indicates the QSO's mode counts towards the WAS Phone award. A mode group of DATA indicates the QSO's mode counts towards the CQ WPX Digital award.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_MY_DXCC_ENTITY_STATUS: {Key: "APP_LOTW_MY_DXCC_ENTITY_STATUS", DataType: "Enumeration", Enumeration: "", Description: "ARRL LOTW: [Current, Deleted] Indicates whether the Logging station's ARRL DXCC entity is \"Current\" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or \"Deleted\"", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_NPSUNIT:               {Key: "APP_LOTW_NPSUNIT", DataType: "String", Enumeration: "", Description: "ARRL LOTW: Present only if the QSLing station included a valid NPOTA Park code or a comma-delimited list of valid NPOTA Park codes", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_NUMREC:                {Key: "APP_LOTW_NUMREC", DataType: "PositiveInteger", Enumeration: "", Description: "ARRL LOTW: Number of QSO records in this download", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_OWNCALL:               {Key: "APP_LOTW_OWNCALL", DataType: "String", Enumeration: "", Description: "ARRL LOTW: \"own\" call sign of the station making the contact. Present only if qso_withown=\"yes\". DEPRECATED -- In the near future, this field will no longer be present in the report output. Programs that import LOTW report files should not rely on the presence of this field. Use STATION_CALLSIGN instead.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_QSLMODE:               {Key: "APP_LOTW_QSLMODE", DataType: "Enumeration", Enumeration: "", Description: "ARRL LOTW: The Mode that your QSO partner indicated on their side of this QSO. Only present if QSL_RCVD=\"Y\" and APP_LOTW_2xQSL=\"N\".", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_QSO_TIMESTAMP:         {Key: "APP_LOTW_QSO_TIMESTAMP", DataType: "String", Enumeration: "", Description: "ARRL LOTW: Combined QSO Date & Time in ISO-8601 format", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_RXQSL:                 {Key: "APP_LOTW_RXQSL", DataType: "String", Enumeration: "", Description: "ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which the confirmation match (QSL) was made/updated at LOTW. Present only if QSL_RCVD=\"Y\"", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_LOTW_RXQSO:                 {Key: "APP_LOTW_RXQSO", DataType: "String", Enumeration: "", Description: "ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which QSO record was inserted/updated at LOTW", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_QRZLOG_LOGID:               {Key: "APP_QRZLOG_LOGID", DataType: "PositiveInteger", Enumeration: "", Description: "QRZ LOG: References the internal ID of a log record in the QRZ system.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_QRZLOG_QSLDATE:             {Key: "APP_QRZLOG_QSLDATE", DataType: "Date", Enumeration: "", Description: "QRZ LOG: QSL Date (e.g. 20210126) This field contains the first date on which we saw a QRZ accepted confirmation for this contact. This could have been a confirmation as a result of matching contacts in the QRZ system or a credit imported from Logbook of The World.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_QRZLOG_STATUS:              {Key: "APP_QRZLOG_STATUS", DataType: "Enumeration", Enumeration: "", Description: "QRZ LOG: Status [C: Confirmed, A: Reserved for future use, N: Not Confirmed, 2: Confirmation Requested, S: Confirmation Requested, Request Seen, R: Confirmation Requested, Request Rejected]", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	APP_SKCCLOGGER_KEYTYPE:         {Key: "APP_SKCCLOGGER_KEYTYPE", DataType: "Enumeration", Enumeration: "", Description: "Straight Key Century Club (SKCC): [BUG, SK, SS] CW KEY TYPE (SS=Side Swiper / Cootie, SK=Straight, B=Bug)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	ARRL_SECT:                      {Key: "ARRL_SECT", DataType: "Enumeration", Enumeration: "ARRL_Section", Description: "the contacted station's ARRL section", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	AWARD_GRANTED:                  {Key: "AWARD_GRANTED", DataType: "SponsoredAwardList", Enumeration: "Sponsored_Award", Description: "the list of awards granted by a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's \"Award Account\" between an award sponsor and the user. For example, in response to a request \"send me a list of the awards granted to AA6YQ\", this might be received: <CALL:5>AA6YQ <AWARD_GRANTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	AWARD_SUBMITTED:                {Key: "AWARD_SUBMITTED", DataType: "SponsoredAwardList", Enumeration: "Sponsored_Award", Description: "the list of awards submitted to a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's \"Award Account\" between an award sponsor and the user. For example, AA6YQ might submit a request for awards by sending the following: <CALL:5>AA6YQ <AWARD_SUBMITTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	A_INDEX:                        {Key: "A_INDEX", DataType: "Number", Enumeration: "", Description: "the geomagnetic A index at the time of the QSO in the range 0 to 400 (inclusive)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 400, IsImportOnly: false, Comments: ""},
	BAND:                           {Key: "BAND", DataType: "Enumeration", Enumeration: "Band", Description: "QSO Band", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	BAND_RX:                        {Key: "BAND_RX", DataType: "Enumeration", Enumeration: "Band", Description: "in a split frequency QSO, the logging station's receiving band", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CALL:                           {Key: "CALL", DataType: "String", Enumeration: "", Description: "the contacted station's callsign", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CHECK:                          {Key: "CHECK", DataType: "String", Enumeration: "", Description: "contest check (e.g. for ARRL Sweepstakes)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CLASS:                          {Key: "CLASS", DataType: "String", Enumeration: "", Description: "contest class (e.g. for ARRL Field Day)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CLUBLOG_QSO_UPLOAD_DATE:        {Key: "CLUBLOG_QSO_UPLOAD_DATE", DataType: "Date", Enumeration: "", Description: "the date the QSO was last uploaded to the Club Log online service", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CLUBLOG_QSO_UPLOAD_STATUS:      {Key: "CLUBLOG_QSO_UPLOAD_STATUS", DataType: "Enumeration", Enumeration: "QSO_Upload_Status", Description: "the upload status of the QSO on the Club Log online service", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CNTY:                           {Key: "CNTY", DataType: "Enumeration", Enumeration: "Secondary_Administrative_Subdivision[DXCC]", Description: "the contacted station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CNTY_ALT:                       {Key: "CNTY_ALT", DataType: "SecondaryAdministrativeSubdivisionListAlt", Enumeration: "", Description: "a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the contacted station See the Data Type for details.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	COMMENT:                        {Key: "COMMENT", DataType: "String", Enumeration: "", Description: "comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG field recommended use: information of interest to the contacted station's operator", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	COMMENT_INTL:                   {Key: "COMMENT_INTL", DataType: "IntlString", Enumeration: "", Description: "comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG_INTL field recommended use: information of interest to the contacted station's operator", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CONT:                           {Key: "CONT", DataType: "Enumeration", Enumeration: "Continent", Description: "the contacted station's Continent", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CONTACTED_OP:                   {Key: "CONTACTED_OP", DataType: "String", Enumeration: "", Description: "the callsign of the individual operating the contacted station", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CONTEST_ID:                     {Key: "CONTEST_ID", DataType: "String", Enumeration: "Contest_ID", Description: "QSO Contest Identifier use enumeration values for interoperability", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	COUNTRY:                        {Key: "COUNTRY", DataType: "String", Enumeration: "", Description: "the contacted station's DXCC entity name", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	COUNTRY_INTL:                   {Key: "COUNTRY_INTL", DataType: "IntlString", Enumeration: "", Description: "the contacted station's DXCC entity name", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CQZ:                            {Key: "CQZ", DataType: "PositiveInteger", Enumeration: "", Description: "the contacted station's CQ Zone in the range 1 to 40 (inclusive)", IsHeaderField: "", MinimumValue: 1, MaximumValue: 40, IsImportOnly: false, Comments: ""},
	CREATED_TIMESTAMP:              {Key: "CREATED_TIMESTAMP", DataType: "String", Enumeration: "", Description: "identifies the UTC date and time that the file was created in the format of 15 characters YYYYMMDD HHMMSS where YYYYMMDD is a Date data type HHMMSS is a 6 character Time data type", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CREDIT_GRANTED:                 {Key: "CREDIT_GRANTED", DataType: "CreditList,AwardList", Enumeration: "Credit,Award", Description: "the list of credits granted to this QSO Use of data type AwardList and enumeration Award are import-only", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	CREDIT_SUBMITTED:               {Key: "CREDIT_SUBMITTED", DataType: "CreditList,AwardList", Enumeration: "Credit,Award", Description: "the list of credits sought for this QSO Use of data type AwardList and enumeration Award are import-only", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	DARC_DOK:                       {Key: "DARC_DOK", DataType: "Enumeration", Enumeration: "", Description: "the contacted station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <DARC_DOK:3>A01", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	DCL_QSLRDATE:                   {Key: "DCL_QSLRDATE", DataType: "Date", Enumeration: "", Description: "date QSL received from DCL (only valid if DCL_QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	DCL_QSLSDATE:                   {Key: "DCL_QSLSDATE", DataType: "Date", Enumeration: "", Description: "date QSL sent to DCL (only valid if DCL_QSL_SENT is Y, Q, or I)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	DCL_QSL_RCVD:                   {Key: "DCL_QSL_RCVD", DataType: "Enumeration", Enumeration: "QSL_Rcvd", Description: "DCL QSL received status Default Value: N", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	DCL_QSL_SENT:                   {Key: "DCL_QSL_SENT", DataType: "Enumeration", Enumeration: "QSL_Sent", Description: "DCL QSL sent status Default Value: N", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	DISTANCE:                       {Key: "DISTANCE", DataType: "Number", Enumeration: "", Description: "the distance between the logging station and the contacted station in kilometers via the specified signal path with a value greater than or equal to 0", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	DXCC:                           {Key: "DXCC", DataType: "Enumeration", Enumeration: "DXCC_Entity_Code", Description: "the contacted station's DXCC Entity Code <DXCC:1>0 means that the contacted station is known not to be within a DXCC entity.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	EMAIL:                          {Key: "EMAIL", DataType: "String", Enumeration: "", Description: "the contacted station's email address", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	EQSL_AG:                        {Key: "EQSL_AG", DataType: "Enumeration", Enumeration: "EQSL_AG", Description: "indicates whether the QSO is known to be \"Authenticity Guaranteed\" by eQSL", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	EQSL_QSLRDATE:                  {Key: "EQSL_QSLRDATE", DataType: "Date", Enumeration: "", Description: "date QSL received from eQSL.cc (only valid if EQSL_QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	EQSL_QSLSDATE:                  {Key: "EQSL_QSLSDATE", DataType: "Date", Enumeration: "", Description: "date QSL sent to eQSL.cc (only valid if EQSL_QSL_SENT is Y, Q, or I)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	EQSL_QSL_RCVD:                  {Key: "EQSL_QSL_RCVD", DataType: "Enumeration", Enumeration: "QSL_Rcvd", Description: "eQSL.cc QSL received status instead of V (import-only) use <CREDIT_GRANTED:42>CQWAZ:eqsl,CQWAZ_BAND:eqsl,CQWAZ_MODE:eqsl Default Value: N", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	EQSL_QSL_SENT:                  {Key: "EQSL_QSL_SENT", DataType: "Enumeration", Enumeration: "QSL_Sent", Description: "eQSL.cc QSL sent status Default Value: N", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	EQ_CALL:                        {Key: "EQ_CALL", DataType: "String", Enumeration: "", Description: "the contacted station's owner's callsign", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	FISTS:                          {Key: "FISTS", DataType: "PositiveInteger", Enumeration: "", Description: "the contacted station's FISTS CW Club member number with a value greater than 0.", IsHeaderField: "", MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	FISTS_CC:                       {Key: "FISTS_CC", DataType: "PositiveInteger", Enumeration: "", Description: "the contacted station's FISTS CW Club Century Certificate (CC) number with a value greater than 0.", IsHeaderField: "", MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	FORCE_INIT:                     {Key: "FORCE_INIT", DataType: "Boolean", Enumeration: "", Description: "new EME \"initial\"", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	FREQ:                           {Key: "FREQ", DataType: "Number", Enumeration: "", Description: "QSO frequency in Megahertz", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	FREQ_RX:                        {Key: "FREQ_RX", DataType: "Number", Enumeration: "", Description: "in a split frequency QSO, the logging station's receiving frequency in Megahertz", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	GRIDSQUARE:                     {Key: "GRIDSQUARE", DataType: "GridSquare", Enumeration: "", Description: "the contacted station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in GRIDSQUARE and the additional 2 or 4 characters in the GRIDSQUARE_EXT field", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	GRIDSQUARE_EXT:                 {Key: "GRIDSQUARE_EXT", DataType: "GridSquareExt", Enumeration: "", Description: "for a contacted station's 10-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9 and 10. For a contacted station's 12-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <GRIDSQUARE:8>FN01MH42 <GRIDSQUARE_EXT:2>BQ", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	GUEST_OP:                       {Key: "GUEST_OP", DataType: "String", Enumeration: "", Description: "import-only: use OPERATOR instead", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: true, Comments: ""},
	HAMLOGEU_QSO_UPLOAD_DATE:       {Key: "HAMLOGEU_QSO_UPLOAD_DATE", DataType: "Date", Enumeration: "", Description: "the date the QSO was last uploaded to the HAMLOG.EU online service", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	HAMLOGEU_QSO_UPLOAD_STATUS:     {Key: "HAMLOGEU_QSO_UPLOAD_STATUS", DataType: "Enumeration", Enumeration: "QSO_Upload_Status", Description: "the upload status of the QSO on the HAMLOG.EU online service", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	HAMQTH_QSO_UPLOAD_DATE:         {Key: "HAMQTH_QSO_UPLOAD_DATE", DataType: "Date", Enumeration: "", Description: "the date the QSO was last uploaded to the HamQTH.com online service", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	HAMQTH_QSO_UPLOAD_STATUS:       {Key: "HAMQTH_QSO_UPLOAD_STATUS", DataType: "Enumeration", Enumeration: "QSO_Upload_Status", Description: "the upload status of the QSO on the HamQTH.com online service", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	HRDLOG_QSO_UPLOAD_DATE:         {Key: "HRDLOG_QSO_UPLOAD_DATE", DataType: "Date", Enumeration: "", Description: "the date the QSO was last uploaded to the HRDLog.net online service", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	HRDLOG_QSO_UPLOAD_STATUS:       {Key: "HRDLOG_QSO_UPLOAD_STATUS", DataType: "Enumeration", Enumeration: "QSO_Upload_Status", Description: "the upload status of the QSO on the HRDLog.net online service", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	IOTA:                           {Key: "IOTA", DataType: "IOTARefNo", Enumeration: "", Description: "the contacted station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	IOTA_ISLAND_ID:                 {Key: "IOTA_ISLAND_ID", DataType: "PositiveInteger", Enumeration: "", Description: "the contacted station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]", IsHeaderField: "", MinimumValue: 1, MaximumValue: 99999999, IsImportOnly: false, Comments: ""},
	ITUZ:                           {Key: "ITUZ", DataType: "PositiveInteger", Enumeration: "", Description: "the contacted station's ITU zone in the range 1 to 90 (inclusive)", IsHeaderField: "", MinimumValue: 1, MaximumValue: 90, IsImportOnly: false, Comments: ""},
	K_INDEX:                        {Key: "K_INDEX", DataType: "Integer", Enumeration: "", Description: "the geomagnetic K index at the time of the QSO in the range 0 to 9 (inclusive)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 9, IsImportOnly: false, Comments: ""},
	LAT:                            {Key: "LAT", DataType: "Location", Enumeration: "", Description: "the contacted station's latitude", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	LON:                            {Key: "LON", DataType: "Location", Enumeration: "", Description: "the contacted station's longitude", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	LOTW_QSLRDATE:                  {Key: "LOTW_QSLRDATE", DataType: "Date", Enumeration: "", Description: "date QSL received from ARRL Logbook of the World (only valid if LOTW_QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	LOTW_QSLSDATE:                  {Key: "LOTW_QSLSDATE", DataType: "Date", Enumeration: "", Description: "date QSL sent to ARRL Logbook of the World (only valid if LOTW_QSL_SENT is Y, Q, or I)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	LOTW_QSL_RCVD:                  {Key: "LOTW_QSL_RCVD", DataType: "Enumeration", Enumeration: "QSL_Rcvd", Description: "ARRL Logbook of the World QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw Default Value: N", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	LOTW_QSL_SENT:                  {Key: "LOTW_QSL_SENT", DataType: "Enumeration", Enumeration: "QSL_Sent", Description: "ARRL Logbook of the World QSL sent status Default Value: N", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MAX_BURSTS:                     {Key: "MAX_BURSTS", DataType: "Number", Enumeration: "", Description: "maximum length of meteor scatter bursts heard by the logging station, in seconds with a value greater than or equal to 0", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MODE:                           {Key: "MODE", DataType: "Enumeration", Enumeration: "Mode", Description: "QSO Mode", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MORSE_KEY_INFO:                 {Key: "MORSE_KEY_INFO", DataType: "String", Enumeration: "", Description: "details of the contacted station's Morse key (e.g. make, model, etc). Example: <MORSE_KEY_INFO:16>Begali Sculpture", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MORSE_KEY_TYPE:                 {Key: "MORSE_KEY_TYPE", DataType: "Enumeration", Enumeration: "Morse_Key_Type", Description: "the contacted station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MS_SHOWER:                      {Key: "MS_SHOWER", DataType: "String", Enumeration: "", Description: "For Meteor Scatter QSOs, the name of the meteor shower in progress", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_ALTITUDE:                    {Key: "MY_ALTITUDE", DataType: "Number", Enumeration: "", Description: "the height of the logging station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <MY_ALTITUDE:4>1500 and 10.5 m is <MY_ALTITUDE:4>10.5", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_ANTENNA:                     {Key: "MY_ANTENNA", DataType: "String", Enumeration: "", Description: "the logging station's antenna", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_ANTENNA_INTL:                {Key: "MY_ANTENNA_INTL", DataType: "IntlString", Enumeration: "", Description: "the logging station's antenna", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_ARRL_SECT:                   {Key: "MY_ARRL_SECT", DataType: "Enumeration", Enumeration: "ARRL_Section", Description: "the logging station's ARRL section", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_CITY:                        {Key: "MY_CITY", DataType: "String", Enumeration: "", Description: "the logging station's city", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_CITY_INTL:                   {Key: "MY_CITY_INTL", DataType: "IntlString", Enumeration: "", Description: "the logging station's city", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_CNTY:                        {Key: "MY_CNTY", DataType: "Enumeration", Enumeration: "Secondary_Administrative_Subdivision[MY_DXCC]", Description: "the logging station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_CNTY_ALT:                    {Key: "MY_CNTY_ALT", DataType: "SecondaryAdministrativeSubdivisionListAlt", Enumeration: "", Description: "a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the logging station See the Data Type for details.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_COUNTRY:                     {Key: "MY_COUNTRY", DataType: "String", Enumeration: "Country", Description: "the logging station's DXCC entity name", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_COUNTRY_INTL:                {Key: "MY_COUNTRY_INTL", DataType: "IntlString", Enumeration: "Country", Description: "the logging station's DXCC entity name", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_CQ_ZONE:                     {Key: "MY_CQ_ZONE", DataType: "PositiveInteger", Enumeration: "", Description: "the logging station's CQ Zone in the range 1 to 40 (inclusive)", IsHeaderField: "", MinimumValue: 1, MaximumValue: 40, IsImportOnly: false, Comments: ""},
	MY_DARC_DOK:                    {Key: "MY_DARC_DOK", DataType: "Enumeration", Enumeration: "", Description: "the logging station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <MY_DARC_DOK:3>A01", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_DXCC:                        {Key: "MY_DXCC", DataType: "Enumeration", Enumeration: "DXCC_Entity_Code", Description: "the logging station's DXCC Entity Code <MY_DXCC:1>0 means that the logging station is known not to be within a DXCC entity.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_FISTS:                       {Key: "MY_FISTS", DataType: "PositiveInteger", Enumeration: "", Description: "the logging station's FISTS CW Club member number with a value greater than 0.", IsHeaderField: "", MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_GRIDSQUARE:                  {Key: "MY_GRIDSQUARE", DataType: "GridSquare", Enumeration: "", Description: "the logging station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in MY_GRIDSQUARE and the additional 2 or 4 characters in the MY_GRIDSQUARE_EXT field", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_GRIDSQUARE_EXT:              {Key: "MY_GRIDSQUARE_EXT", DataType: "GridSquareExt", Enumeration: "", Description: "for a logging station's 10-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9 and 10. For a logging station's 12-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <MY_GRIDSQUARE:8>FN01MH42 <MY_GRIDSQUARE_EXT:2>BQ", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_IOTA:                        {Key: "MY_IOTA", DataType: "IOTARefNo", Enumeration: "", Description: "the logging station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_IOTA_ISLAND_ID:              {Key: "MY_IOTA_ISLAND_ID", DataType: "PositiveInteger", Enumeration: "", Description: "the logging station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]", IsHeaderField: "", MinimumValue: 1, MaximumValue: 99999999, IsImportOnly: false, Comments: ""},
	MY_ITU_ZONE:                    {Key: "MY_ITU_ZONE", DataType: "PositiveInteger", Enumeration: "", Description: "the logging station's ITU zone in the range 1 to 90 (inclusive)", IsHeaderField: "", MinimumValue: 1, MaximumValue: 90, IsImportOnly: false, Comments: ""},
	MY_LAT:                         {Key: "MY_LAT", DataType: "Location", Enumeration: "", Description: "the logging station's latitude", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_LON:                         {Key: "MY_LON", DataType: "Location", Enumeration: "", Description: "the logging station's longitude", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_MORSE_KEY_INFO:              {Key: "MY_MORSE_KEY_INFO", DataType: "String", Enumeration: "", Description: "details of the logging station's Morse key (e.g. make, model, etc). Example: <MY_MORSE_KEY_INFO:16>Begali Sculpture", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_MORSE_KEY_TYPE:              {Key: "MY_MORSE_KEY_TYPE", DataType: "Enumeration", Enumeration: "Morse_Key_Type", Description: "the logging station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_NAME:                        {Key: "MY_NAME", DataType: "String", Enumeration: "", Description: "the logging operator's name", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_NAME_INTL:                   {Key: "MY_NAME_INTL", DataType: "IntlString", Enumeration: "", Description: "the logging operator's name", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_POSTAL_CODE:                 {Key: "MY_POSTAL_CODE", DataType: "String", Enumeration: "", Description: "the logging station's postal code", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_POSTAL_CODE_INTL:            {Key: "MY_POSTAL_CODE_INTL", DataType: "IntlString", Enumeration: "", Description: "the logging station's postal code", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_POTA_REF:                    {Key: "MY_POTA_REF", DataType: "POTARefList", Enumeration: "", Description: "a comma-delimited list of one or more of the logging station's POTA (Parks on the Air) reference(s). Examples: <MY_POTA_REF:6>K-0059 <MY_POTA_REF:7>K-10000 <MY_POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_RIG:                         {Key: "MY_RIG", DataType: "String", Enumeration: "", Description: "description of the logging station's equipment", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_RIG_INTL:                    {Key: "MY_RIG_INTL", DataType: "IntlString", Enumeration: "", Description: "description of the logging station's equipment", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_SIG:                         {Key: "MY_SIG", DataType: "String", Enumeration: "", Description: "special interest activity or event", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_SIG_INFO:                    {Key: "MY_SIG_INFO", DataType: "String", Enumeration: "", Description: "special interest activity or event information", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_SIG_INFO_INTL:               {Key: "MY_SIG_INFO_INTL", DataType: "IntlString", Enumeration: "", Description: "special interest activity or event information", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_SIG_INTL:                    {Key: "MY_SIG_INTL", DataType: "IntlString", Enumeration: "", Description: "special interest activity or event", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_SOTA_REF:                    {Key: "MY_SOTA_REF", DataType: "SOTARef", Enumeration: "", Description: "the logging station's International SOTA Reference.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_STATE:                       {Key: "MY_STATE", DataType: "Enumeration", Enumeration: "Primary_Administrative_Subdivision[MY_DXCC]", Description: "the code for the logging station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_STREET:                      {Key: "MY_STREET", DataType: "String", Enumeration: "", Description: "the logging station's street", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_STREET_INTL:                 {Key: "MY_STREET_INTL", DataType: "IntlString", Enumeration: "", Description: "the logging station's street", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_USACA_COUNTIES:              {Key: "MY_USACA_COUNTIES", DataType: "SecondarySubdivisionList", Enumeration: "", Description: "two US counties in the case where the logging station is located on a border between two counties, representing counties that the contacted station may claim for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_VUCC_GRIDS:                  {Key: "MY_VUCC_GRIDS", DataType: "GridSquareList", Enumeration: "", Description: "two or four adjacent Maidenhead grid locators, each four or six characters long, representing the logging station's grid squares that the contacted station may claim for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	MY_WWFF_REF:                    {Key: "MY_WWFF_REF", DataType: "WWFFRef", Enumeration: "", Description: "the logging station's WWFF (World Wildlife Flora & Fauna) reference", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	NAME:                           {Key: "NAME", DataType: "String", Enumeration: "", Description: "the contacted station's operator's name", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	NAME_INTL:                      {Key: "NAME_INTL", DataType: "IntlString", Enumeration: "", Description: "the contacted station's operator's name", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	NOTES:                          {Key: "NOTES", DataType: "MultilineString", Enumeration: "", Description: "QSO notes recommended use: information of interest to the logging station's operator", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	NOTES_INTL:                     {Key: "NOTES_INTL", DataType: "IntlMultilineString", Enumeration: "", Description: "QSO notes recommended use: information of interest to the logging station's operator", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	NR_BURSTS:                      {Key: "NR_BURSTS", DataType: "Integer", Enumeration: "", Description: "the number of meteor scatter bursts heard by the logging station with a value greater than or equal to 0", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	NR_PINGS:                       {Key: "NR_PINGS", DataType: "Integer", Enumeration: "", Description: "the number of meteor scatter pings heard by the logging station with a value greater than or equal to 0", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	OPERATOR:                       {Key: "OPERATOR", DataType: "String", Enumeration: "", Description: "the logging operator's callsign if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	OWNER_CALLSIGN:                 {Key: "OWNER_CALLSIGN", DataType: "String", Enumeration: "", Description: "the callsign of the owner of the station used to log the contact (the callsign of the OPERATOR's host) if OWNER_CALLSIGN is absent, STATION_CALLSIGN shall be treated as both the logging station's callsign and the callsign of the owner of the station", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	PFX:                            {Key: "PFX", DataType: "String", Enumeration: "", Description: "the contacted station's WPX prefix", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	POTA_REF:                       {Key: "POTA_REF", DataType: "POTARefList", Enumeration: "", Description: "a comma-delimited list of one or more of the contacted station's POTA (Parks on the Air) reference(s). Examples: <POTA_REF:6>K-5033 <POTA_REF:13>VE-5082@CA-AB <POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	PRECEDENCE:                     {Key: "PRECEDENCE", DataType: "String", Enumeration: "", Description: "contest precedence (e.g. for ARRL Sweepstakes)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	PROGRAMID:                      {Key: "PROGRAMID", DataType: "String", Enumeration: "", Description: "identifies the name of the logger, converter, or utility that created or processed this ADIF file To help avoid name clashes, the ADIF PROGRAMID Register provides a voluntary list of PROGRAMID values.", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	PROGRAMVERSION:                 {Key: "PROGRAMVERSION", DataType: "String", Enumeration: "", Description: "identifies the version of the logger, converter, or utility that created or processed this ADIF file", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	PROP_MODE:                      {Key: "PROP_MODE", DataType: "Enumeration", Enumeration: "Propagation_Mode", Description: "QSO propagation mode", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	PUBLIC_KEY:                     {Key: "PUBLIC_KEY", DataType: "String", Enumeration: "", Description: "public encryption key", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QRZCOM_QSO_DOWNLOAD_DATE:       {Key: "QRZCOM_QSO_DOWNLOAD_DATE", DataType: "Date", Enumeration: "", Description: "date QSO downloaded from QRZ.COM logbook", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QRZCOM_QSO_DOWNLOAD_STATUS:     {Key: "QRZCOM_QSO_DOWNLOAD_STATUS", DataType: "Enumeration", Enumeration: "QSO_Download_Status", Description: "QRZ.COM logbook QSO download status", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QRZCOM_QSO_UPLOAD_DATE:         {Key: "QRZCOM_QSO_UPLOAD_DATE", DataType: "Date", Enumeration: "", Description: "the date the QSO was last uploaded to the QRZ.COM online service", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QRZCOM_QSO_UPLOAD_STATUS:       {Key: "QRZCOM_QSO_UPLOAD_STATUS", DataType: "Enumeration", Enumeration: "QSO_Upload_Status", Description: "the upload status of the QSO on the QRZ.COM online service", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSLMSG:                         {Key: "QSLMSG", DataType: "MultilineString", Enumeration: "", Description: "a message for the contacted station's operator to be incorporated in a paper or electronic QSL", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSLMSG_INTL:                    {Key: "QSLMSG_INTL", DataType: "IntlMultilineString", Enumeration: "", Description: "a message for the contacted station's operator to be incorporated in a paper or electronic QSL", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSLMSG_RCVD:                    {Key: "QSLMSG_RCVD", DataType: "MultilineString", Enumeration: "", Description: "a message addressed to the logging station's operator incorporated in a paper or electronic QSL", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSLRDATE:                       {Key: "QSLRDATE", DataType: "Date", Enumeration: "", Description: "QSL received date (only valid if QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSLSDATE:                       {Key: "QSLSDATE", DataType: "Date", Enumeration: "", Description: "QSL sent date (only valid if QSL_SENT is Y, Q, or I)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSL_RCVD:                       {Key: "QSL_RCVD", DataType: "Enumeration", Enumeration: "QSL_Rcvd", Description: "QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card Default Value: N", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSL_RCVD_VIA:                   {Key: "QSL_RCVD_VIA", DataType: "Enumeration", Enumeration: "QSL_Via", Description: "if QSL_RCVD is set to 'Y' or 'V', the means by which the QSL was received by the logging station; otherwise, the means by which the logging station requested or intends to request that the QSL be conveyed. (Note: 'V' is import-only) use of M (manager) is import-only", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSL_SENT:                       {Key: "QSL_SENT", DataType: "Enumeration", Enumeration: "QSL_Sent", Description: "QSL sent status Default Value: N", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSL_SENT_VIA:                   {Key: "QSL_SENT_VIA", DataType: "Enumeration", Enumeration: "QSL_Via", Description: "if QSL_SENT is set to 'Y', the means by which the QSL was sent by the logging station; otherwise, the means by which the logging station intends to convey the QSL use of M (manager) is import-only", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSL_VIA:                        {Key: "QSL_VIA", DataType: "String", Enumeration: "", Description: "the contacted station's QSL route", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSO_COMPLETE:                   {Key: "QSO_COMPLETE", DataType: "Enumeration", Enumeration: "QSO_Complete", Description: "indicates whether the QSO was complete from the perspective of the logging station Y - yes N - no NIL - not heard ? - uncertain", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSO_DATE:                       {Key: "QSO_DATE", DataType: "Date", Enumeration: "", Description: "date on which the QSO started", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSO_DATE_OFF:                   {Key: "QSO_DATE_OFF", DataType: "Date", Enumeration: "", Description: "date on which the QSO ended", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QSO_RANDOM:                     {Key: "QSO_RANDOM", DataType: "Boolean", Enumeration: "", Description: "indicates whether the QSO was random or scheduled", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QTH:                            {Key: "QTH", DataType: "String", Enumeration: "", Description: "the contacted station's city", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	QTH_INTL:                       {Key: "QTH_INTL", DataType: "IntlString", Enumeration: "", Description: "the contacted station's city", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	REGION:                         {Key: "REGION", DataType: "Enumeration", Enumeration: "Region", Description: "the contacted station's WAE or CQ entity contained within a DXCC entity. the value None indicates that the WAE or CQ entity is the DXCC entity in the DXCC field. nothing can be inferred from the absence of the REGION field", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	RIG:                            {Key: "RIG", DataType: "MultilineString", Enumeration: "", Description: "description of the contacted station's equipment", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	RIG_INTL:                       {Key: "RIG_INTL", DataType: "IntlMultilineString", Enumeration: "", Description: "description of the contacted station's equipment", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	RST_RCVD:                       {Key: "RST_RCVD", DataType: "String", Enumeration: "", Description: "signal report from the contacted station", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	RST_SENT:                       {Key: "RST_SENT", DataType: "String", Enumeration: "", Description: "signal report sent to the contacted station", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	RX_PWR:                         {Key: "RX_PWR", DataType: "Number", Enumeration: "", Description: "the contacted station's transmitter power in Watts with a value greater than or equal to 0", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SAT_MODE:                       {Key: "SAT_MODE", DataType: "String", Enumeration: "", Description: "satellite mode - a code representing the satellite's uplink band and downlink band", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SAT_NAME:                       {Key: "SAT_NAME", DataType: "String", Enumeration: "", Description: "name of satellite", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SFI:                            {Key: "SFI", DataType: "Integer", Enumeration: "", Description: "the solar flux at the time of the QSO in the range 0 to 300 (inclusive).", IsHeaderField: "", MinimumValue: 0, MaximumValue: 300, IsImportOnly: false, Comments: ""},
	SIG:                            {Key: "SIG", DataType: "String", Enumeration: "", Description: "the name of the contacted station's special activity or interest group", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SIG_INFO:                       {Key: "SIG_INFO", DataType: "String", Enumeration: "", Description: "information associated with the contacted station's activity or interest group", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SIG_INFO_INTL:                  {Key: "SIG_INFO_INTL", DataType: "IntlString", Enumeration: "", Description: "information associated with the contacted station's activity or interest group", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SIG_INTL:                       {Key: "SIG_INTL", DataType: "IntlString", Enumeration: "", Description: "the name of the contacted station's special activity or interest group", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SILENT_KEY:                     {Key: "SILENT_KEY", DataType: "Boolean", Enumeration: "", Description: "'Y' indicates that the contacted station's operator is now a Silent Key.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SKCC:                           {Key: "SKCC", DataType: "String", Enumeration: "", Description: "the contacted station's Straight Key Century Club (SKCC) member information", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SOTA_REF:                       {Key: "SOTA_REF", DataType: "SOTARef", Enumeration: "", Description: "the contacted station's International SOTA Reference.", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SRX:                            {Key: "SRX", DataType: "Integer", Enumeration: "", Description: "contest QSO received serial number with a value greater than or equal to 0", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SRX_STRING:                     {Key: "SRX_STRING", DataType: "String", Enumeration: "", Description: "contest QSO received information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	STATE:                          {Key: "STATE", DataType: "Enumeration", Enumeration: "Primary_Administrative_Subdivision[DXCC]", Description: "the code for the contacted station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	STATION_CALLSIGN:               {Key: "STATION_CALLSIGN", DataType: "String", Enumeration: "", Description: "the logging station's callsign (the callsign used over the air) if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	STX:                            {Key: "STX", DataType: "Integer", Enumeration: "", Description: "contest QSO transmitted serial number with a value greater than or equal to 0", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	STX_STRING:                     {Key: "STX_STRING", DataType: "String", Enumeration: "", Description: "contest QSO transmitted information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SUBMODE:                        {Key: "SUBMODE", DataType: "String", Enumeration: "Submode[MODE]", Description: "QSO Submode use enumeration values for interoperability", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	SWL:                            {Key: "SWL", DataType: "Boolean", Enumeration: "", Description: "indicates that the QSO information pertains to an SWL report", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	TEN_TEN:                        {Key: "TEN_TEN", DataType: "PositiveInteger", Enumeration: "", Description: "Ten-Ten number with a value greater than 0", IsHeaderField: "", MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	TIME_OFF:                       {Key: "TIME_OFF", DataType: "Time", Enumeration: "", Description: "HHMM or HHMMSS in UTC in the absence of <QSO_DATE_OFF>, the QSO duration is less than 24 hours. For example, the following is a QSO starting at 14 July 2020 23:55 and finishing at 15 July 2020 01:00: <QSO_DATE:8>20200714 <TIME_ON:4>2355 <TIME_OFF:4>0100", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	TIME_ON:                        {Key: "TIME_ON", DataType: "Time", Enumeration: "", Description: "HHMM or HHMMSS in UTC", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	TX_PWR:                         {Key: "TX_PWR", DataType: "Number", Enumeration: "", Description: "the logging station's power in Watts with a value greater than or equal to 0", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	UKSMG:                          {Key: "UKSMG", DataType: "PositiveInteger", Enumeration: "", Description: "the contacted station's UKSMG member number with a value greater than 0", IsHeaderField: "", MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	USACA_COUNTIES:                 {Key: "USACA_COUNTIES", DataType: "SecondarySubdivisionList", Enumeration: "", Description: "two US counties in the case where the contacted station is located on a border between two counties, representing counties credited to the QSO for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	USERDEF1:                       {Key: "USERDEF1", DataType: "String", Enumeration: "", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	USERDEF2:                       {Key: "USERDEF2", DataType: "String", Enumeration: "", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	USERDEF3:                       {Key: "USERDEF3", DataType: "String", Enumeration: "", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	USERDEF4:                       {Key: "USERDEF4", DataType: "String", Enumeration: "", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	USERDEF5:                       {Key: "USERDEF5", DataType: "String", Enumeration: "", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	USERDEF6:                       {Key: "USERDEF6", DataType: "String", Enumeration: "", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	USERDEF7:                       {Key: "USERDEF7", DataType: "String", Enumeration: "", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	USERDEF8:                       {Key: "USERDEF8", DataType: "String", Enumeration: "", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	USERDEF9:                       {Key: "USERDEF9", DataType: "String", Enumeration: "", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: "true", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	VE_PROV:                        {Key: "VE_PROV", DataType: "String", Enumeration: "", Description: "import-only: use STATE instead", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: true, Comments: ""},
	VUCC_GRIDS:                     {Key: "VUCC_GRIDS", DataType: "GridSquareList", Enumeration: "", Description: "two or four adjacent Maidenhead grid locators, each four or six characters long, representing the contacted station's grid squares credited to the QSO for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	WEB:                            {Key: "WEB", DataType: "String", Enumeration: "", Description: "the contacted station's URL", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	WWFF_REF:                       {Key: "WWFF_REF", DataType: "WWFFRef", Enumeration: "", Description: "the contacted station's WWFF (World Wildlife Flora & Fauna) reference", IsHeaderField: "", MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
}

// All ADIField specifications including depreciated and import only.
var ADIFieldListAll = []Spec{
	ADIFieldMap[ADDRESS],
	ADIFieldMap[ADDRESS_INTL],
	ADIFieldMap[ADIF_VER],
	ADIFieldMap[AGE],
	ADIFieldMap[ALTITUDE],
	ADIFieldMap[ANT_AZ],
	ADIFieldMap[ANT_EL],
	ADIFieldMap[ANT_PATH],
	ADIFieldMap[APP_LOTW_2XQSL],
	ADIFieldMap[APP_LOTW_CQZ_INFERRED],
	ADIFieldMap[APP_LOTW_CQZ_INVALID],
	ADIFieldMap[APP_LOTW_CQZ_USERINVALID],
	ADIFieldMap[APP_LOTW_CREDIT_GRANTED],
	ADIFieldMap[APP_LOTW_CREDIT_SUBMITTED],
	ADIFieldMap[APP_LOTW_DXCC_ENTITY_STATUS],
	ADIFieldMap[APP_LOTW_EOF],
	ADIFieldMap[APP_LOTW_GRIDSQUARE_INVALID],
	ADIFieldMap[APP_LOTW_ITUZ_INFERRED],
	ADIFieldMap[APP_LOTW_ITUZ_INVALID],
	ADIFieldMap[APP_LOTW_ITUZ_USERINVALID],
	ADIFieldMap[APP_LOTW_LASTQSL],
	ADIFieldMap[APP_LOTW_LASTQSORX],
	ADIFieldMap[APP_LOTW_MODE],
	ADIFieldMap[APP_LOTW_MODEGROUP],
	ADIFieldMap[APP_LOTW_MY_DXCC_ENTITY_STATUS],
	ADIFieldMap[APP_LOTW_NPSUNIT],
	ADIFieldMap[APP_LOTW_NUMREC],
	ADIFieldMap[APP_LOTW_OWNCALL],
	ADIFieldMap[APP_LOTW_QSLMODE],
	ADIFieldMap[APP_LOTW_QSO_TIMESTAMP],
	ADIFieldMap[APP_LOTW_RXQSL],
	ADIFieldMap[APP_LOTW_RXQSO],
	ADIFieldMap[APP_QRZLOG_LOGID],
	ADIFieldMap[APP_QRZLOG_QSLDATE],
	ADIFieldMap[APP_QRZLOG_STATUS],
	ADIFieldMap[APP_SKCCLOGGER_KEYTYPE],
	ADIFieldMap[ARRL_SECT],
	ADIFieldMap[AWARD_GRANTED],
	ADIFieldMap[AWARD_SUBMITTED],
	ADIFieldMap[A_INDEX],
	ADIFieldMap[BAND],
	ADIFieldMap[BAND_RX],
	ADIFieldMap[CALL],
	ADIFieldMap[CHECK],
	ADIFieldMap[CLASS],
	ADIFieldMap[CLUBLOG_QSO_UPLOAD_DATE],
	ADIFieldMap[CLUBLOG_QSO_UPLOAD_STATUS],
	ADIFieldMap[CNTY],
	ADIFieldMap[CNTY_ALT],
	ADIFieldMap[COMMENT],
	ADIFieldMap[COMMENT_INTL],
	ADIFieldMap[CONT],
	ADIFieldMap[CONTACTED_OP],
	ADIFieldMap[CONTEST_ID],
	ADIFieldMap[COUNTRY],
	ADIFieldMap[COUNTRY_INTL],
	ADIFieldMap[CQZ],
	ADIFieldMap[CREATED_TIMESTAMP],
	ADIFieldMap[CREDIT_GRANTED],
	ADIFieldMap[CREDIT_SUBMITTED],
	ADIFieldMap[DARC_DOK],
	ADIFieldMap[DCL_QSLRDATE],
	ADIFieldMap[DCL_QSLSDATE],
	ADIFieldMap[DCL_QSL_RCVD],
	ADIFieldMap[DCL_QSL_SENT],
	ADIFieldMap[DISTANCE],
	ADIFieldMap[DXCC],
	ADIFieldMap[EMAIL],
	ADIFieldMap[EQSL_AG],
	ADIFieldMap[EQSL_QSLRDATE],
	ADIFieldMap[EQSL_QSLSDATE],
	ADIFieldMap[EQSL_QSL_RCVD],
	ADIFieldMap[EQSL_QSL_SENT],
	ADIFieldMap[EQ_CALL],
	ADIFieldMap[FISTS],
	ADIFieldMap[FISTS_CC],
	ADIFieldMap[FORCE_INIT],
	ADIFieldMap[FREQ],
	ADIFieldMap[FREQ_RX],
	ADIFieldMap[GRIDSQUARE],
	ADIFieldMap[GRIDSQUARE_EXT],
	ADIFieldMap[GUEST_OP],
	ADIFieldMap[HAMLOGEU_QSO_UPLOAD_DATE],
	ADIFieldMap[HAMLOGEU_QSO_UPLOAD_STATUS],
	ADIFieldMap[HAMQTH_QSO_UPLOAD_DATE],
	ADIFieldMap[HAMQTH_QSO_UPLOAD_STATUS],
	ADIFieldMap[HRDLOG_QSO_UPLOAD_DATE],
	ADIFieldMap[HRDLOG_QSO_UPLOAD_STATUS],
	ADIFieldMap[IOTA],
	ADIFieldMap[IOTA_ISLAND_ID],
	ADIFieldMap[ITUZ],
	ADIFieldMap[K_INDEX],
	ADIFieldMap[LAT],
	ADIFieldMap[LON],
	ADIFieldMap[LOTW_QSLRDATE],
	ADIFieldMap[LOTW_QSLSDATE],
	ADIFieldMap[LOTW_QSL_RCVD],
	ADIFieldMap[LOTW_QSL_SENT],
	ADIFieldMap[MAX_BURSTS],
	ADIFieldMap[MODE],
	ADIFieldMap[MORSE_KEY_INFO],
	ADIFieldMap[MORSE_KEY_TYPE],
	ADIFieldMap[MS_SHOWER],
	ADIFieldMap[MY_ALTITUDE],
	ADIFieldMap[MY_ANTENNA],
	ADIFieldMap[MY_ANTENNA_INTL],
	ADIFieldMap[MY_ARRL_SECT],
	ADIFieldMap[MY_CITY],
	ADIFieldMap[MY_CITY_INTL],
	ADIFieldMap[MY_CNTY],
	ADIFieldMap[MY_CNTY_ALT],
	ADIFieldMap[MY_COUNTRY],
	ADIFieldMap[MY_COUNTRY_INTL],
	ADIFieldMap[MY_CQ_ZONE],
	ADIFieldMap[MY_DARC_DOK],
	ADIFieldMap[MY_DXCC],
	ADIFieldMap[MY_FISTS],
	ADIFieldMap[MY_GRIDSQUARE],
	ADIFieldMap[MY_GRIDSQUARE_EXT],
	ADIFieldMap[MY_IOTA],
	ADIFieldMap[MY_IOTA_ISLAND_ID],
	ADIFieldMap[MY_ITU_ZONE],
	ADIFieldMap[MY_LAT],
	ADIFieldMap[MY_LON],
	ADIFieldMap[MY_MORSE_KEY_INFO],
	ADIFieldMap[MY_MORSE_KEY_TYPE],
	ADIFieldMap[MY_NAME],
	ADIFieldMap[MY_NAME_INTL],
	ADIFieldMap[MY_POSTAL_CODE],
	ADIFieldMap[MY_POSTAL_CODE_INTL],
	ADIFieldMap[MY_POTA_REF],
	ADIFieldMap[MY_RIG],
	ADIFieldMap[MY_RIG_INTL],
	ADIFieldMap[MY_SIG],
	ADIFieldMap[MY_SIG_INFO],
	ADIFieldMap[MY_SIG_INFO_INTL],
	ADIFieldMap[MY_SIG_INTL],
	ADIFieldMap[MY_SOTA_REF],
	ADIFieldMap[MY_STATE],
	ADIFieldMap[MY_STREET],
	ADIFieldMap[MY_STREET_INTL],
	ADIFieldMap[MY_USACA_COUNTIES],
	ADIFieldMap[MY_VUCC_GRIDS],
	ADIFieldMap[MY_WWFF_REF],
	ADIFieldMap[NAME],
	ADIFieldMap[NAME_INTL],
	ADIFieldMap[NOTES],
	ADIFieldMap[NOTES_INTL],
	ADIFieldMap[NR_BURSTS],
	ADIFieldMap[NR_PINGS],
	ADIFieldMap[OPERATOR],
	ADIFieldMap[OWNER_CALLSIGN],
	ADIFieldMap[PFX],
	ADIFieldMap[POTA_REF],
	ADIFieldMap[PRECEDENCE],
	ADIFieldMap[PROGRAMID],
	ADIFieldMap[PROGRAMVERSION],
	ADIFieldMap[PROP_MODE],
	ADIFieldMap[PUBLIC_KEY],
	ADIFieldMap[QRZCOM_QSO_DOWNLOAD_DATE],
	ADIFieldMap[QRZCOM_QSO_DOWNLOAD_STATUS],
	ADIFieldMap[QRZCOM_QSO_UPLOAD_DATE],
	ADIFieldMap[QRZCOM_QSO_UPLOAD_STATUS],
	ADIFieldMap[QSLMSG],
	ADIFieldMap[QSLMSG_INTL],
	ADIFieldMap[QSLMSG_RCVD],
	ADIFieldMap[QSLRDATE],
	ADIFieldMap[QSLSDATE],
	ADIFieldMap[QSL_RCVD],
	ADIFieldMap[QSL_RCVD_VIA],
	ADIFieldMap[QSL_SENT],
	ADIFieldMap[QSL_SENT_VIA],
	ADIFieldMap[QSL_VIA],
	ADIFieldMap[QSO_COMPLETE],
	ADIFieldMap[QSO_DATE],
	ADIFieldMap[QSO_DATE_OFF],
	ADIFieldMap[QSO_RANDOM],
	ADIFieldMap[QTH],
	ADIFieldMap[QTH_INTL],
	ADIFieldMap[REGION],
	ADIFieldMap[RIG],
	ADIFieldMap[RIG_INTL],
	ADIFieldMap[RST_RCVD],
	ADIFieldMap[RST_SENT],
	ADIFieldMap[RX_PWR],
	ADIFieldMap[SAT_MODE],
	ADIFieldMap[SAT_NAME],
	ADIFieldMap[SFI],
	ADIFieldMap[SIG],
	ADIFieldMap[SIG_INFO],
	ADIFieldMap[SIG_INFO_INTL],
	ADIFieldMap[SIG_INTL],
	ADIFieldMap[SILENT_KEY],
	ADIFieldMap[SKCC],
	ADIFieldMap[SOTA_REF],
	ADIFieldMap[SRX],
	ADIFieldMap[SRX_STRING],
	ADIFieldMap[STATE],
	ADIFieldMap[STATION_CALLSIGN],
	ADIFieldMap[STX],
	ADIFieldMap[STX_STRING],
	ADIFieldMap[SUBMODE],
	ADIFieldMap[SWL],
	ADIFieldMap[TEN_TEN],
	ADIFieldMap[TIME_OFF],
	ADIFieldMap[TIME_ON],
	ADIFieldMap[TX_PWR],
	ADIFieldMap[UKSMG],
	ADIFieldMap[USACA_COUNTIES],
	ADIFieldMap[USERDEF1],
	ADIFieldMap[USERDEF2],
	ADIFieldMap[USERDEF3],
	ADIFieldMap[USERDEF4],
	ADIFieldMap[USERDEF5],
	ADIFieldMap[USERDEF6],
	ADIFieldMap[USERDEF7],
	ADIFieldMap[USERDEF8],
	ADIFieldMap[USERDEF9],
	ADIFieldMap[VE_PROV],
	ADIFieldMap[VUCC_GRIDS],
	ADIFieldMap[WEB],
	ADIFieldMap[WWFF_REF],
}

// All ADIField specifications values that are NOT marked import-only.
var ADIFieldListCurrent = []Spec{
	ADIFieldMap[ADDRESS],
	ADIFieldMap[ADDRESS_INTL],
	ADIFieldMap[ADIF_VER],
	ADIFieldMap[AGE],
	ADIFieldMap[ALTITUDE],
	ADIFieldMap[ANT_AZ],
	ADIFieldMap[ANT_EL],
	ADIFieldMap[ANT_PATH],
	ADIFieldMap[APP_LOTW_2XQSL],
	ADIFieldMap[APP_LOTW_CQZ_INFERRED],
	ADIFieldMap[APP_LOTW_CQZ_INVALID],
	ADIFieldMap[APP_LOTW_CQZ_USERINVALID],
	ADIFieldMap[APP_LOTW_CREDIT_GRANTED],
	ADIFieldMap[APP_LOTW_CREDIT_SUBMITTED],
	ADIFieldMap[APP_LOTW_DXCC_ENTITY_STATUS],
	ADIFieldMap[APP_LOTW_EOF],
	ADIFieldMap[APP_LOTW_GRIDSQUARE_INVALID],
	ADIFieldMap[APP_LOTW_ITUZ_INFERRED],
	ADIFieldMap[APP_LOTW_ITUZ_INVALID],
	ADIFieldMap[APP_LOTW_ITUZ_USERINVALID],
	ADIFieldMap[APP_LOTW_LASTQSL],
	ADIFieldMap[APP_LOTW_LASTQSORX],
	ADIFieldMap[APP_LOTW_MODE],
	ADIFieldMap[APP_LOTW_MODEGROUP],
	ADIFieldMap[APP_LOTW_MY_DXCC_ENTITY_STATUS],
	ADIFieldMap[APP_LOTW_NPSUNIT],
	ADIFieldMap[APP_LOTW_NUMREC],
	ADIFieldMap[APP_LOTW_OWNCALL],
	ADIFieldMap[APP_LOTW_QSLMODE],
	ADIFieldMap[APP_LOTW_QSO_TIMESTAMP],
	ADIFieldMap[APP_LOTW_RXQSL],
	ADIFieldMap[APP_LOTW_RXQSO],
	ADIFieldMap[APP_QRZLOG_LOGID],
	ADIFieldMap[APP_QRZLOG_QSLDATE],
	ADIFieldMap[APP_QRZLOG_STATUS],
	ADIFieldMap[APP_SKCCLOGGER_KEYTYPE],
	ADIFieldMap[ARRL_SECT],
	ADIFieldMap[AWARD_GRANTED],
	ADIFieldMap[AWARD_SUBMITTED],
	ADIFieldMap[A_INDEX],
	ADIFieldMap[BAND],
	ADIFieldMap[BAND_RX],
	ADIFieldMap[CALL],
	ADIFieldMap[CHECK],
	ADIFieldMap[CLASS],
	ADIFieldMap[CLUBLOG_QSO_UPLOAD_DATE],
	ADIFieldMap[CLUBLOG_QSO_UPLOAD_STATUS],
	ADIFieldMap[CNTY],
	ADIFieldMap[CNTY_ALT],
	ADIFieldMap[COMMENT],
	ADIFieldMap[COMMENT_INTL],
	ADIFieldMap[CONT],
	ADIFieldMap[CONTACTED_OP],
	ADIFieldMap[CONTEST_ID],
	ADIFieldMap[COUNTRY],
	ADIFieldMap[COUNTRY_INTL],
	ADIFieldMap[CQZ],
	ADIFieldMap[CREATED_TIMESTAMP],
	ADIFieldMap[CREDIT_GRANTED],
	ADIFieldMap[CREDIT_SUBMITTED],
	ADIFieldMap[DARC_DOK],
	ADIFieldMap[DCL_QSLRDATE],
	ADIFieldMap[DCL_QSLSDATE],
	ADIFieldMap[DCL_QSL_RCVD],
	ADIFieldMap[DCL_QSL_SENT],
	ADIFieldMap[DISTANCE],
	ADIFieldMap[DXCC],
	ADIFieldMap[EMAIL],
	ADIFieldMap[EQSL_AG],
	ADIFieldMap[EQSL_QSLRDATE],
	ADIFieldMap[EQSL_QSLSDATE],
	ADIFieldMap[EQSL_QSL_RCVD],
	ADIFieldMap[EQSL_QSL_SENT],
	ADIFieldMap[EQ_CALL],
	ADIFieldMap[FISTS],
	ADIFieldMap[FISTS_CC],
	ADIFieldMap[FORCE_INIT],
	ADIFieldMap[FREQ],
	ADIFieldMap[FREQ_RX],
	ADIFieldMap[GRIDSQUARE],
	ADIFieldMap[GRIDSQUARE_EXT],
	ADIFieldMap[HAMLOGEU_QSO_UPLOAD_DATE],
	ADIFieldMap[HAMLOGEU_QSO_UPLOAD_STATUS],
	ADIFieldMap[HAMQTH_QSO_UPLOAD_DATE],
	ADIFieldMap[HAMQTH_QSO_UPLOAD_STATUS],
	ADIFieldMap[HRDLOG_QSO_UPLOAD_DATE],
	ADIFieldMap[HRDLOG_QSO_UPLOAD_STATUS],
	ADIFieldMap[IOTA],
	ADIFieldMap[IOTA_ISLAND_ID],
	ADIFieldMap[ITUZ],
	ADIFieldMap[K_INDEX],
	ADIFieldMap[LAT],
	ADIFieldMap[LON],
	ADIFieldMap[LOTW_QSLRDATE],
	ADIFieldMap[LOTW_QSLSDATE],
	ADIFieldMap[LOTW_QSL_RCVD],
	ADIFieldMap[LOTW_QSL_SENT],
	ADIFieldMap[MAX_BURSTS],
	ADIFieldMap[MODE],
	ADIFieldMap[MORSE_KEY_INFO],
	ADIFieldMap[MORSE_KEY_TYPE],
	ADIFieldMap[MS_SHOWER],
	ADIFieldMap[MY_ALTITUDE],
	ADIFieldMap[MY_ANTENNA],
	ADIFieldMap[MY_ANTENNA_INTL],
	ADIFieldMap[MY_ARRL_SECT],
	ADIFieldMap[MY_CITY],
	ADIFieldMap[MY_CITY_INTL],
	ADIFieldMap[MY_CNTY],
	ADIFieldMap[MY_CNTY_ALT],
	ADIFieldMap[MY_COUNTRY],
	ADIFieldMap[MY_COUNTRY_INTL],
	ADIFieldMap[MY_CQ_ZONE],
	ADIFieldMap[MY_DARC_DOK],
	ADIFieldMap[MY_DXCC],
	ADIFieldMap[MY_FISTS],
	ADIFieldMap[MY_GRIDSQUARE],
	ADIFieldMap[MY_GRIDSQUARE_EXT],
	ADIFieldMap[MY_IOTA],
	ADIFieldMap[MY_IOTA_ISLAND_ID],
	ADIFieldMap[MY_ITU_ZONE],
	ADIFieldMap[MY_LAT],
	ADIFieldMap[MY_LON],
	ADIFieldMap[MY_MORSE_KEY_INFO],
	ADIFieldMap[MY_MORSE_KEY_TYPE],
	ADIFieldMap[MY_NAME],
	ADIFieldMap[MY_NAME_INTL],
	ADIFieldMap[MY_POSTAL_CODE],
	ADIFieldMap[MY_POSTAL_CODE_INTL],
	ADIFieldMap[MY_POTA_REF],
	ADIFieldMap[MY_RIG],
	ADIFieldMap[MY_RIG_INTL],
	ADIFieldMap[MY_SIG],
	ADIFieldMap[MY_SIG_INFO],
	ADIFieldMap[MY_SIG_INFO_INTL],
	ADIFieldMap[MY_SIG_INTL],
	ADIFieldMap[MY_SOTA_REF],
	ADIFieldMap[MY_STATE],
	ADIFieldMap[MY_STREET],
	ADIFieldMap[MY_STREET_INTL],
	ADIFieldMap[MY_USACA_COUNTIES],
	ADIFieldMap[MY_VUCC_GRIDS],
	ADIFieldMap[MY_WWFF_REF],
	ADIFieldMap[NAME],
	ADIFieldMap[NAME_INTL],
	ADIFieldMap[NOTES],
	ADIFieldMap[NOTES_INTL],
	ADIFieldMap[NR_BURSTS],
	ADIFieldMap[NR_PINGS],
	ADIFieldMap[OPERATOR],
	ADIFieldMap[OWNER_CALLSIGN],
	ADIFieldMap[PFX],
	ADIFieldMap[POTA_REF],
	ADIFieldMap[PRECEDENCE],
	ADIFieldMap[PROGRAMID],
	ADIFieldMap[PROGRAMVERSION],
	ADIFieldMap[PROP_MODE],
	ADIFieldMap[PUBLIC_KEY],
	ADIFieldMap[QRZCOM_QSO_DOWNLOAD_DATE],
	ADIFieldMap[QRZCOM_QSO_DOWNLOAD_STATUS],
	ADIFieldMap[QRZCOM_QSO_UPLOAD_DATE],
	ADIFieldMap[QRZCOM_QSO_UPLOAD_STATUS],
	ADIFieldMap[QSLMSG],
	ADIFieldMap[QSLMSG_INTL],
	ADIFieldMap[QSLMSG_RCVD],
	ADIFieldMap[QSLRDATE],
	ADIFieldMap[QSLSDATE],
	ADIFieldMap[QSL_RCVD],
	ADIFieldMap[QSL_RCVD_VIA],
	ADIFieldMap[QSL_SENT],
	ADIFieldMap[QSL_SENT_VIA],
	ADIFieldMap[QSL_VIA],
	ADIFieldMap[QSO_COMPLETE],
	ADIFieldMap[QSO_DATE],
	ADIFieldMap[QSO_DATE_OFF],
	ADIFieldMap[QSO_RANDOM],
	ADIFieldMap[QTH],
	ADIFieldMap[QTH_INTL],
	ADIFieldMap[REGION],
	ADIFieldMap[RIG],
	ADIFieldMap[RIG_INTL],
	ADIFieldMap[RST_RCVD],
	ADIFieldMap[RST_SENT],
	ADIFieldMap[RX_PWR],
	ADIFieldMap[SAT_MODE],
	ADIFieldMap[SAT_NAME],
	ADIFieldMap[SFI],
	ADIFieldMap[SIG],
	ADIFieldMap[SIG_INFO],
	ADIFieldMap[SIG_INFO_INTL],
	ADIFieldMap[SIG_INTL],
	ADIFieldMap[SILENT_KEY],
	ADIFieldMap[SKCC],
	ADIFieldMap[SOTA_REF],
	ADIFieldMap[SRX],
	ADIFieldMap[SRX_STRING],
	ADIFieldMap[STATE],
	ADIFieldMap[STATION_CALLSIGN],
	ADIFieldMap[STX],
	ADIFieldMap[STX_STRING],
	ADIFieldMap[SUBMODE],
	ADIFieldMap[SWL],
	ADIFieldMap[TEN_TEN],
	ADIFieldMap[TIME_OFF],
	ADIFieldMap[TIME_ON],
	ADIFieldMap[TX_PWR],
	ADIFieldMap[UKSMG],
	ADIFieldMap[USACA_COUNTIES],
	ADIFieldMap[USERDEF1],
	ADIFieldMap[USERDEF2],
	ADIFieldMap[USERDEF3],
	ADIFieldMap[USERDEF4],
	ADIFieldMap[USERDEF5],
	ADIFieldMap[USERDEF6],
	ADIFieldMap[USERDEF7],
	ADIFieldMap[USERDEF8],
	ADIFieldMap[USERDEF9],
	ADIFieldMap[VUCC_GRIDS],
	ADIFieldMap[WEB],
	ADIFieldMap[WWFF_REF],
}
