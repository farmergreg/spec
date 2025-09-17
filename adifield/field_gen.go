// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package adifield provides code and constants as defined in ADIF 3.1.6 (Released)
package adifield

import "sync"

const (
	ADDRESS                        Field = "ADDRESS"                        // Record: the contacted station's complete mailing address: full name, street address, city, postal code, and country
	ADDRESS_INTL                   Field = "ADDRESS_INTL"                   // Record: the contacted station's complete mailing address: full name, street address, city, postal code, and country
	ADIF_VER                       Field = "ADIF_VER"                       // Header: identifies the version of ADIF used in this file in the format X.Y.Z where X is an integer designating the ADIF epoch Y is an integer between 0 and 9 designating the major version Z is an integer between 0 and 9 designating the minor version
	AGE                            Field = "AGE"                            // Record: the contacted station's operator's age in years in the range 0 to 120 (inclusive)
	ALTITUDE                       Field = "ALTITUDE"                       // Record: the height of the contacted station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <ALTITUDE:4>1500 and 10.5 m is <ALTITUDE:4>10.5
	ANT_AZ                         Field = "ANT_AZ"                         // Record: the logging station's antenna azimuth, in degrees with a value between 0 to 360 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 370 is exported as 10). True north is 0 degrees with values increasing in a clockwise direction.
	ANT_EL                         Field = "ANT_EL"                         // Record: the logging station's antenna elevation, in degrees with a value between -90 to 90 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 100 is exported as 80). The horizon is 0 degrees with values increasing as the angle moves in an upward direction.
	ANT_PATH                       Field = "ANT_PATH"                       // Record: the signal path
	APP_LOTW_2XQSL                 Field = "APP_LOTW_2XQSL"                 // Record: ARRL LOTW: [Y, N] APP_LOTW_2xQSL ="Y" Indicates this confirmation is considered to be "two-way" in the QSO's Mode. It is a confirmation that counts towards the mode-specific awards & endorsement in the WAS program. Only present if QSL_RCVD="Y".
	APP_LOTW_CQZ_INFERRED          Field = "APP_LOTW_CQZ_INFERRED"          // Record: ARRL LOTW
	APP_LOTW_CQZ_INVALID           Field = "APP_LOTW_CQZ_INVALID"           // Record: ARRL LOTW
	APP_LOTW_CQZ_USERINVALID       Field = "APP_LOTW_CQZ_USERINVALID"       // Record: ARRL LOTW: Present only if the QSLing station included an invalid CQ zone value in its station location uploaded to LOTW
	APP_LOTW_CREDIT_GRANTED        Field = "APP_LOTW_CREDIT_GRANTED"        // Record: ARRL LOTW: Award credits granted based upon this QSL. Awards per LOTW award enumeration.
	APP_LOTW_CREDIT_SUBMITTED      Field = "APP_LOTW_CREDIT_SUBMITTED"      // Record: ARRL LOTW: Award credits for which this QSL has been submitted on an award application to sponsor. Awards per LOTW award enumeration.
	APP_LOTW_DXCC_ENTITY_STATUS    Field = "APP_LOTW_DXCC_ENTITY_STATUS"    // Record: ARRL LOTW: [Current, Deleted] Indicates whether the ARRL DXCC entity is "Current" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or "Deleted"
	APP_LOTW_EOF                   Field = "APP_LOTW_EOF"                   // Record: ARRL LOTW: End of File marker. No other fields expected in this record.
	APP_LOTW_GRIDSQUARE_INVALID    Field = "APP_LOTW_GRIDSQUARE_INVALID"    // Record: ARRL LOTW
	APP_LOTW_ITUZ_INFERRED         Field = "APP_LOTW_ITUZ_INFERRED"         // Record: ARRL LOTW
	APP_LOTW_ITUZ_INVALID          Field = "APP_LOTW_ITUZ_INVALID"          // Record: ARRL LOTW
	APP_LOTW_ITUZ_USERINVALID      Field = "APP_LOTW_ITUZ_USERINVALID"      // Record: ARRL LOTW: Present only if the QSLing station included an invalid ITU zone value in its station location uploaded to LOTW
	APP_LOTW_LASTQSL               Field = "APP_LOTW_LASTQSL"               // Header: ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl="yes". Date and time at which the most recent QSL record in this report was received. This will be the QSL date/time of the first QSO record in the file. This value only applies to the records in this file.
	APP_LOTW_LASTQSORX             Field = "APP_LOTW_LASTQSORX"             // Header: ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl="no". Date and time at which the most recent QSO record in this report was received. This will be the QSO received (uploaded) date/time of the first QSO record in the file. This value only applies to the records in this file.
	APP_LOTW_MODE                  Field = "APP_LOTW_MODE"                  // Record: ARRL LOTW: Field is present when the mode recorded in Logbook cannot be unambiguously represented as an ADIF-compliant value
	APP_LOTW_MODEGROUP             Field = "APP_LOTW_MODEGROUP"             // Record: ARRL LOTW: [CW, PHONE, DATA] The mode group indicates whether the QSO's mode counts towards the CW, Phone or Digital awards respectively in the DXCC program. A mode group of CW indicates the QSO's mode counts towards the CW endorsement on WAS awards. A mode group of PHONE indicates the QSO's mode counts towards the WAS Phone award. A mode group of DATA indicates the QSO's mode counts towards the CQ WPX Digital award.
	APP_LOTW_MY_DXCC_ENTITY_STATUS Field = "APP_LOTW_MY_DXCC_ENTITY_STATUS" // Record: ARRL LOTW: [Current, Deleted] Indicates whether the Logging station's ARRL DXCC entity is "Current" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or "Deleted"
	APP_LOTW_NPSUNIT               Field = "APP_LOTW_NPSUNIT"               // Record: ARRL LOTW: Present only if the QSLing station included a valid NPOTA Park code or a comma-delimited list of valid NPOTA Park codes
	APP_LOTW_NUMREC                Field = "APP_LOTW_NUMREC"                // Header: ARRL LOTW: Number of QSO records in this download
	APP_LOTW_OWNCALL               Field = "APP_LOTW_OWNCALL"               // Record: ARRL LOTW: "own" call sign of the station making the contact. Present only if qso_withown="yes". DEPRECATED -- In the near future, this field will no longer be present in the report output. Programs that import LOTW report files should not rely on the presence of this field. Use STATION_CALLSIGN instead.
	APP_LOTW_QSLMODE               Field = "APP_LOTW_QSLMODE"               // Record: ARRL LOTW: The Mode that your QSO partner indicated on their side of this QSO. Only present if QSL_RCVD="Y" and APP_LOTW_2xQSL="N".
	APP_LOTW_QSO_TIMESTAMP         Field = "APP_LOTW_QSO_TIMESTAMP"         // Record: ARRL LOTW: Combined QSO Date & Time in ISO-8601 format
	APP_LOTW_RXQSL                 Field = "APP_LOTW_RXQSL"                 // Record: ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which the confirmation match (QSL) was made/updated at LOTW. Present only if QSL_RCVD="Y"
	APP_LOTW_RXQSO                 Field = "APP_LOTW_RXQSO"                 // Record: ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which QSO record was inserted/updated at LOTW
	APP_QRZLOG_LOGID               Field = "APP_QRZLOG_LOGID"               // Record: QRZ LOG: References the internal ID of a log record in the QRZ system.
	APP_QRZLOG_QSLDATE             Field = "APP_QRZLOG_QSLDATE"             // Record: QRZ LOG: QSL Date (e.g. 20210126) This field contains the first date on which we saw a QRZ accepted confirmation for this contact. This could have been a confirmation as a result of matching contacts in the QRZ system or a credit imported from Logbook of The World.
	APP_QRZLOG_STATUS              Field = "APP_QRZLOG_STATUS"              // Record: QRZ LOG: Status [C: Confirmed, A: Reserved for future use, N: Not Confirmed, 2: Confirmation Requested, S: Confirmation Requested, Request Seen, R: Confirmation Requested, Request Rejected]
	APP_SKCCLOGGER_KEYTYPE         Field = "APP_SKCCLOGGER_KEYTYPE"         // Record: Straight Key Century Club (SKCC): [BUG, SK, SS] CW KEY TYPE (SS=Side Swiper / Cootie, SK=Straight, B=Bug)
	ARRL_SECT                      Field = "ARRL_SECT"                      // Record: the contacted station's ARRL section
	AWARD_GRANTED                  Field = "AWARD_GRANTED"                  // Record: the list of awards granted by a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's "Award Account" between an award sponsor and the user. For example, in response to a request "send me a list of the awards granted to AA6YQ", this might be received: <CALL:5>AA6YQ <AWARD_GRANTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m
	AWARD_SUBMITTED                Field = "AWARD_SUBMITTED"                // Record: the list of awards submitted to a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's "Award Account" between an award sponsor and the user. For example, AA6YQ might submit a request for awards by sending the following: <CALL:5>AA6YQ <AWARD_SUBMITTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m
	A_INDEX                        Field = "A_INDEX"                        // Record: the geomagnetic A index at the time of the QSO in the range 0 to 400 (inclusive)
	BAND                           Field = "BAND"                           // Record: QSO Band
	BAND_RX                        Field = "BAND_RX"                        // Record: in a split frequency QSO, the logging station's receiving band
	CALL                           Field = "CALL"                           // Record: the contacted station's callsign
	CHECK                          Field = "CHECK"                          // Record: contest check (e.g. for ARRL Sweepstakes)
	CLASS                          Field = "CLASS"                          // Record: contest class (e.g. for ARRL Field Day)
	CLUBLOG_QSO_UPLOAD_DATE        Field = "CLUBLOG_QSO_UPLOAD_DATE"        // Record: the date the QSO was last uploaded to the Club Log online service
	CLUBLOG_QSO_UPLOAD_STATUS      Field = "CLUBLOG_QSO_UPLOAD_STATUS"      // Record: the upload status of the QSO on the Club Log online service
	CNTY                           Field = "CNTY"                           // Record: the contacted station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format
	CNTY_ALT                       Field = "CNTY_ALT"                       // Record: a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the contacted station See the Data Type for details.
	COMMENT                        Field = "COMMENT"                        // Record: comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG field recommended use: information of interest to the contacted station's operator
	COMMENT_INTL                   Field = "COMMENT_INTL"                   // Record: comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG_INTL field recommended use: information of interest to the contacted station's operator
	CONT                           Field = "CONT"                           // Record: the contacted station's Continent
	CONTACTED_OP                   Field = "CONTACTED_OP"                   // Record: the callsign of the individual operating the contacted station
	CONTEST_ID                     Field = "CONTEST_ID"                     // Record: QSO Contest Identifier use enumeration values for interoperability
	COUNTRY                        Field = "COUNTRY"                        // Record: the contacted station's DXCC entity name
	COUNTRY_INTL                   Field = "COUNTRY_INTL"                   // Record: the contacted station's DXCC entity name
	CQZ                            Field = "CQZ"                            // Record: the contacted station's CQ Zone in the range 1 to 40 (inclusive)
	CREATED_TIMESTAMP              Field = "CREATED_TIMESTAMP"              // Header: identifies the UTC date and time that the file was created in the format of 15 characters YYYYMMDD HHMMSS where YYYYMMDD is a Date data type HHMMSS is a 6 character Time data type
	CREDIT_GRANTED                 Field = "CREDIT_GRANTED"                 // Record: the list of credits granted to this QSO Use of data type AwardList and enumeration Award are import-only
	CREDIT_SUBMITTED               Field = "CREDIT_SUBMITTED"               // Record: the list of credits sought for this QSO Use of data type AwardList and enumeration Award are import-only
	DARC_DOK                       Field = "DARC_DOK"                       // Record: the contacted station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <DARC_DOK:3>A01
	DCL_QSLRDATE                   Field = "DCL_QSLRDATE"                   // Record: date QSL received from DCL (only valid if DCL_QSL_RCVD is Y, I, or V)(V import-only)
	DCL_QSLSDATE                   Field = "DCL_QSLSDATE"                   // Record: date QSL sent to DCL (only valid if DCL_QSL_SENT is Y, Q, or I)
	DCL_QSL_RCVD                   Field = "DCL_QSL_RCVD"                   // Record: DCL QSL received status Default Value: N
	DCL_QSL_SENT                   Field = "DCL_QSL_SENT"                   // Record: DCL QSL sent status Default Value: N
	DISTANCE                       Field = "DISTANCE"                       // Record: the distance between the logging station and the contacted station in kilometers via the specified signal path with a value greater than or equal to 0
	DXCC                           Field = "DXCC"                           // Record: the contacted station's DXCC Entity Code <DXCC:1>0 means that the contacted station is known not to be within a DXCC entity.
	EMAIL                          Field = "EMAIL"                          // Record: the contacted station's email address
	EQSL_AG                        Field = "EQSL_AG"                        // Record: indicates whether the QSO is known to be "Authenticity Guaranteed" by eQSL
	EQSL_QSLRDATE                  Field = "EQSL_QSLRDATE"                  // Record: date QSL received from eQSL.cc (only valid if EQSL_QSL_RCVD is Y, I, or V)(V import-only)
	EQSL_QSLSDATE                  Field = "EQSL_QSLSDATE"                  // Record: date QSL sent to eQSL.cc (only valid if EQSL_QSL_SENT is Y, Q, or I)
	EQSL_QSL_RCVD                  Field = "EQSL_QSL_RCVD"                  // Record: eQSL.cc QSL received status instead of V (import-only) use <CREDIT_GRANTED:42>CQWAZ:eqsl,CQWAZ_BAND:eqsl,CQWAZ_MODE:eqsl Default Value: N
	EQSL_QSL_SENT                  Field = "EQSL_QSL_SENT"                  // Record: eQSL.cc QSL sent status Default Value: N
	EQ_CALL                        Field = "EQ_CALL"                        // Record: the contacted station's owner's callsign
	FISTS                          Field = "FISTS"                          // Record: the contacted station's FISTS CW Club member number with a value greater than 0.
	FISTS_CC                       Field = "FISTS_CC"                       // Record: the contacted station's FISTS CW Club Century Certificate (CC) number with a value greater than 0.
	FORCE_INIT                     Field = "FORCE_INIT"                     // Record: new EME "initial"
	FREQ                           Field = "FREQ"                           // Record: QSO frequency in Megahertz
	FREQ_RX                        Field = "FREQ_RX"                        // Record: in a split frequency QSO, the logging station's receiving frequency in Megahertz
	GRIDSQUARE                     Field = "GRIDSQUARE"                     // Record: the contacted station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in GRIDSQUARE and the additional 2 or 4 characters in the GRIDSQUARE_EXT field
	GRIDSQUARE_EXT                 Field = "GRIDSQUARE_EXT"                 // Record: for a contacted station's 10-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9 and 10. For a contacted station's 12-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <GRIDSQUARE:8>FN01MH42 <GRIDSQUARE_EXT:2>BQ
	GUEST_OP                       Field = "GUEST_OP"                       // Deprecated: Record: import-only: use OPERATOR instead
	HAMLOGEU_QSO_UPLOAD_DATE       Field = "HAMLOGEU_QSO_UPLOAD_DATE"       // Record: the date the QSO was last uploaded to the HAMLOG.EU online service
	HAMLOGEU_QSO_UPLOAD_STATUS     Field = "HAMLOGEU_QSO_UPLOAD_STATUS"     // Record: the upload status of the QSO on the HAMLOG.EU online service
	HAMQTH_QSO_UPLOAD_DATE         Field = "HAMQTH_QSO_UPLOAD_DATE"         // Record: the date the QSO was last uploaded to the HamQTH.com online service
	HAMQTH_QSO_UPLOAD_STATUS       Field = "HAMQTH_QSO_UPLOAD_STATUS"       // Record: the upload status of the QSO on the HamQTH.com online service
	HRDLOG_QSO_UPLOAD_DATE         Field = "HRDLOG_QSO_UPLOAD_DATE"         // Record: the date the QSO was last uploaded to the HRDLog.net online service
	HRDLOG_QSO_UPLOAD_STATUS       Field = "HRDLOG_QSO_UPLOAD_STATUS"       // Record: the upload status of the QSO on the HRDLog.net online service
	IOTA                           Field = "IOTA"                           // Record: the contacted station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]
	IOTA_ISLAND_ID                 Field = "IOTA_ISLAND_ID"                 // Record: the contacted station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]
	ITUZ                           Field = "ITUZ"                           // Record: the contacted station's ITU zone in the range 1 to 90 (inclusive)
	K_INDEX                        Field = "K_INDEX"                        // Record: the geomagnetic K index at the time of the QSO in the range 0 to 9 (inclusive)
	LAT                            Field = "LAT"                            // Record: the contacted station's latitude
	LON                            Field = "LON"                            // Record: the contacted station's longitude
	LOTW_QSLRDATE                  Field = "LOTW_QSLRDATE"                  // Record: date QSL received from ARRL Logbook of the World (only valid if LOTW_QSL_RCVD is Y, I, or V)(V import-only)
	LOTW_QSLSDATE                  Field = "LOTW_QSLSDATE"                  // Record: date QSL sent to ARRL Logbook of the World (only valid if LOTW_QSL_SENT is Y, Q, or I)
	LOTW_QSL_RCVD                  Field = "LOTW_QSL_RCVD"                  // Record: ARRL Logbook of the World QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw Default Value: N
	LOTW_QSL_SENT                  Field = "LOTW_QSL_SENT"                  // Record: ARRL Logbook of the World QSL sent status Default Value: N
	MAX_BURSTS                     Field = "MAX_BURSTS"                     // Record: maximum length of meteor scatter bursts heard by the logging station, in seconds with a value greater than or equal to 0
	MODE                           Field = "MODE"                           // Record: QSO Mode
	MORSE_KEY_INFO                 Field = "MORSE_KEY_INFO"                 // Record: details of the contacted station's Morse key (e.g. make, model, etc). Example: <MORSE_KEY_INFO:16>Begali Sculpture
	MORSE_KEY_TYPE                 Field = "MORSE_KEY_TYPE"                 // Record: the contacted station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP
	MS_SHOWER                      Field = "MS_SHOWER"                      // Record: For Meteor Scatter QSOs, the name of the meteor shower in progress
	MY_ALTITUDE                    Field = "MY_ALTITUDE"                    // Record: the height of the logging station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <MY_ALTITUDE:4>1500 and 10.5 m is <MY_ALTITUDE:4>10.5
	MY_ANTENNA                     Field = "MY_ANTENNA"                     // Record: the logging station's antenna
	MY_ANTENNA_INTL                Field = "MY_ANTENNA_INTL"                // Record: the logging station's antenna
	MY_ARRL_SECT                   Field = "MY_ARRL_SECT"                   // Record: the logging station's ARRL section
	MY_CITY                        Field = "MY_CITY"                        // Record: the logging station's city
	MY_CITY_INTL                   Field = "MY_CITY_INTL"                   // Record: the logging station's city
	MY_CNTY                        Field = "MY_CNTY"                        // Record: the logging station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format
	MY_CNTY_ALT                    Field = "MY_CNTY_ALT"                    // Record: a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the logging station See the Data Type for details.
	MY_COUNTRY                     Field = "MY_COUNTRY"                     // Record: the logging station's DXCC entity name
	MY_COUNTRY_INTL                Field = "MY_COUNTRY_INTL"                // Record: the logging station's DXCC entity name
	MY_CQ_ZONE                     Field = "MY_CQ_ZONE"                     // Record: the logging station's CQ Zone in the range 1 to 40 (inclusive)
	MY_DARC_DOK                    Field = "MY_DARC_DOK"                    // Record: the logging station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <MY_DARC_DOK:3>A01
	MY_DXCC                        Field = "MY_DXCC"                        // Record: the logging station's DXCC Entity Code <MY_DXCC:1>0 means that the logging station is known not to be within a DXCC entity.
	MY_FISTS                       Field = "MY_FISTS"                       // Record: the logging station's FISTS CW Club member number with a value greater than 0.
	MY_GRIDSQUARE                  Field = "MY_GRIDSQUARE"                  // Record: the logging station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in MY_GRIDSQUARE and the additional 2 or 4 characters in the MY_GRIDSQUARE_EXT field
	MY_GRIDSQUARE_EXT              Field = "MY_GRIDSQUARE_EXT"              // Record: for a logging station's 10-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9 and 10. For a logging station's 12-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <MY_GRIDSQUARE:8>FN01MH42 <MY_GRIDSQUARE_EXT:2>BQ
	MY_IOTA                        Field = "MY_IOTA"                        // Record: the logging station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]
	MY_IOTA_ISLAND_ID              Field = "MY_IOTA_ISLAND_ID"              // Record: the logging station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]
	MY_ITU_ZONE                    Field = "MY_ITU_ZONE"                    // Record: the logging station's ITU zone in the range 1 to 90 (inclusive)
	MY_LAT                         Field = "MY_LAT"                         // Record: the logging station's latitude
	MY_LON                         Field = "MY_LON"                         // Record: the logging station's longitude
	MY_MORSE_KEY_INFO              Field = "MY_MORSE_KEY_INFO"              // Record: details of the logging station's Morse key (e.g. make, model, etc). Example: <MY_MORSE_KEY_INFO:16>Begali Sculpture
	MY_MORSE_KEY_TYPE              Field = "MY_MORSE_KEY_TYPE"              // Record: the logging station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP
	MY_NAME                        Field = "MY_NAME"                        // Record: the logging operator's name
	MY_NAME_INTL                   Field = "MY_NAME_INTL"                   // Record: the logging operator's name
	MY_POSTAL_CODE                 Field = "MY_POSTAL_CODE"                 // Record: the logging station's postal code
	MY_POSTAL_CODE_INTL            Field = "MY_POSTAL_CODE_INTL"            // Record: the logging station's postal code
	MY_POTA_REF                    Field = "MY_POTA_REF"                    // Record: a comma-delimited list of one or more of the logging station's POTA (Parks on the Air) reference(s). Examples: <MY_POTA_REF:6>K-0059 <MY_POTA_REF:7>K-10000 <MY_POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY
	MY_RIG                         Field = "MY_RIG"                         // Record: description of the logging station's equipment
	MY_RIG_INTL                    Field = "MY_RIG_INTL"                    // Record: description of the logging station's equipment
	MY_SIG                         Field = "MY_SIG"                         // Record: special interest activity or event
	MY_SIG_INFO                    Field = "MY_SIG_INFO"                    // Record: special interest activity or event information
	MY_SIG_INFO_INTL               Field = "MY_SIG_INFO_INTL"               // Record: special interest activity or event information
	MY_SIG_INTL                    Field = "MY_SIG_INTL"                    // Record: special interest activity or event
	MY_SOTA_REF                    Field = "MY_SOTA_REF"                    // Record: the logging station's International SOTA Reference.
	MY_STATE                       Field = "MY_STATE"                       // Record: the code for the logging station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)
	MY_STREET                      Field = "MY_STREET"                      // Record: the logging station's street
	MY_STREET_INTL                 Field = "MY_STREET_INTL"                 // Record: the logging station's street
	MY_USACA_COUNTIES              Field = "MY_USACA_COUNTIES"              // Record: two US counties in the case where the logging station is located on a border between two counties, representing counties that the contacted station may claim for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire
	MY_VUCC_GRIDS                  Field = "MY_VUCC_GRIDS"                  // Record: two or four adjacent Maidenhead grid locators, each four or six characters long, representing the logging station's grid squares that the contacted station may claim for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07
	MY_WWFF_REF                    Field = "MY_WWFF_REF"                    // Record: the logging station's WWFF (World Wildlife Flora & Fauna) reference
	NAME                           Field = "NAME"                           // Record: the contacted station's operator's name
	NAME_INTL                      Field = "NAME_INTL"                      // Record: the contacted station's operator's name
	NOTES                          Field = "NOTES"                          // Record: QSO notes recommended use: information of interest to the logging station's operator
	NOTES_INTL                     Field = "NOTES_INTL"                     // Record: QSO notes recommended use: information of interest to the logging station's operator
	NR_BURSTS                      Field = "NR_BURSTS"                      // Record: the number of meteor scatter bursts heard by the logging station with a value greater than or equal to 0
	NR_PINGS                       Field = "NR_PINGS"                       // Record: the number of meteor scatter pings heard by the logging station with a value greater than or equal to 0
	OPERATOR                       Field = "OPERATOR"                       // Record: the logging operator's callsign if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign
	OWNER_CALLSIGN                 Field = "OWNER_CALLSIGN"                 // Record: the callsign of the owner of the station used to log the contact (the callsign of the OPERATOR's host) if OWNER_CALLSIGN is absent, STATION_CALLSIGN shall be treated as both the logging station's callsign and the callsign of the owner of the station
	PFX                            Field = "PFX"                            // Record: the contacted station's WPX prefix
	POTA_REF                       Field = "POTA_REF"                       // Record: a comma-delimited list of one or more of the contacted station's POTA (Parks on the Air) reference(s). Examples: <POTA_REF:6>K-5033 <POTA_REF:13>VE-5082@CA-AB <POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY
	PRECEDENCE                     Field = "PRECEDENCE"                     // Record: contest precedence (e.g. for ARRL Sweepstakes)
	PROGRAMID                      Field = "PROGRAMID"                      // Header: identifies the name of the logger, converter, or utility that created or processed this ADIF file To help avoid name clashes, the ADIF PROGRAMID Register provides a voluntary list of PROGRAMID values.
	PROGRAMVERSION                 Field = "PROGRAMVERSION"                 // Header: identifies the version of the logger, converter, or utility that created or processed this ADIF file
	PROP_MODE                      Field = "PROP_MODE"                      // Record: QSO propagation mode
	PUBLIC_KEY                     Field = "PUBLIC_KEY"                     // Record: public encryption key
	QRZCOM_QSO_DOWNLOAD_DATE       Field = "QRZCOM_QSO_DOWNLOAD_DATE"       // Record: date QSO downloaded from QRZ.COM logbook
	QRZCOM_QSO_DOWNLOAD_STATUS     Field = "QRZCOM_QSO_DOWNLOAD_STATUS"     // Record: QRZ.COM logbook QSO download status
	QRZCOM_QSO_UPLOAD_DATE         Field = "QRZCOM_QSO_UPLOAD_DATE"         // Record: the date the QSO was last uploaded to the QRZ.COM online service
	QRZCOM_QSO_UPLOAD_STATUS       Field = "QRZCOM_QSO_UPLOAD_STATUS"       // Record: the upload status of the QSO on the QRZ.COM online service
	QSLMSG                         Field = "QSLMSG"                         // Record: a message for the contacted station's operator to be incorporated in a paper or electronic QSL
	QSLMSG_INTL                    Field = "QSLMSG_INTL"                    // Record: a message for the contacted station's operator to be incorporated in a paper or electronic QSL
	QSLMSG_RCVD                    Field = "QSLMSG_RCVD"                    // Record: a message addressed to the logging station's operator incorporated in a paper or electronic QSL
	QSLRDATE                       Field = "QSLRDATE"                       // Record: QSL received date (only valid if QSL_RCVD is Y, I, or V)(V import-only)
	QSLSDATE                       Field = "QSLSDATE"                       // Record: QSL sent date (only valid if QSL_SENT is Y, Q, or I)
	QSL_RCVD                       Field = "QSL_RCVD"                       // Record: QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card Default Value: N
	QSL_RCVD_VIA                   Field = "QSL_RCVD_VIA"                   // Record: if QSL_RCVD is set to 'Y' or 'V', the means by which the QSL was received by the logging station; otherwise, the means by which the logging station requested or intends to request that the QSL be conveyed. (Note: 'V' is import-only) use of M (manager) is import-only
	QSL_SENT                       Field = "QSL_SENT"                       // Record: QSL sent status Default Value: N
	QSL_SENT_VIA                   Field = "QSL_SENT_VIA"                   // Record: if QSL_SENT is set to 'Y', the means by which the QSL was sent by the logging station; otherwise, the means by which the logging station intends to convey the QSL use of M (manager) is import-only
	QSL_VIA                        Field = "QSL_VIA"                        // Record: the contacted station's QSL route
	QSO_COMPLETE                   Field = "QSO_COMPLETE"                   // Record: indicates whether the QSO was complete from the perspective of the logging station Y - yes N - no NIL - not heard ? - uncertain
	QSO_DATE                       Field = "QSO_DATE"                       // Record: date on which the QSO started
	QSO_DATE_OFF                   Field = "QSO_DATE_OFF"                   // Record: date on which the QSO ended
	QSO_RANDOM                     Field = "QSO_RANDOM"                     // Record: indicates whether the QSO was random or scheduled
	QTH                            Field = "QTH"                            // Record: the contacted station's city
	QTH_INTL                       Field = "QTH_INTL"                       // Record: the contacted station's city
	REGION                         Field = "REGION"                         // Record: the contacted station's WAE or CQ entity contained within a DXCC entity. the value None indicates that the WAE or CQ entity is the DXCC entity in the DXCC field. nothing can be inferred from the absence of the REGION field
	RIG                            Field = "RIG"                            // Record: description of the contacted station's equipment
	RIG_INTL                       Field = "RIG_INTL"                       // Record: description of the contacted station's equipment
	RST_RCVD                       Field = "RST_RCVD"                       // Record: signal report from the contacted station
	RST_SENT                       Field = "RST_SENT"                       // Record: signal report sent to the contacted station
	RX_PWR                         Field = "RX_PWR"                         // Record: the contacted station's transmitter power in Watts with a value greater than or equal to 0
	SAT_MODE                       Field = "SAT_MODE"                       // Record: satellite mode - a code representing the satellite's uplink band and downlink band
	SAT_NAME                       Field = "SAT_NAME"                       // Record: name of satellite
	SFI                            Field = "SFI"                            // Record: the solar flux at the time of the QSO in the range 0 to 300 (inclusive).
	SIG                            Field = "SIG"                            // Record: the name of the contacted station's special activity or interest group
	SIG_INFO                       Field = "SIG_INFO"                       // Record: information associated with the contacted station's activity or interest group
	SIG_INFO_INTL                  Field = "SIG_INFO_INTL"                  // Record: information associated with the contacted station's activity or interest group
	SIG_INTL                       Field = "SIG_INTL"                       // Record: the name of the contacted station's special activity or interest group
	SILENT_KEY                     Field = "SILENT_KEY"                     // Record: 'Y' indicates that the contacted station's operator is now a Silent Key.
	SKCC                           Field = "SKCC"                           // Record: the contacted station's Straight Key Century Club (SKCC) member information
	SOTA_REF                       Field = "SOTA_REF"                       // Record: the contacted station's International SOTA Reference.
	SRX                            Field = "SRX"                            // Record: contest QSO received serial number with a value greater than or equal to 0
	SRX_STRING                     Field = "SRX_STRING"                     // Record: contest QSO received information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail
	STATE                          Field = "STATE"                          // Record: the code for the contacted station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)
	STATION_CALLSIGN               Field = "STATION_CALLSIGN"               // Record: the logging station's callsign (the callsign used over the air) if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign
	STX                            Field = "STX"                            // Record: contest QSO transmitted serial number with a value greater than or equal to 0
	STX_STRING                     Field = "STX_STRING"                     // Record: contest QSO transmitted information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail
	SUBMODE                        Field = "SUBMODE"                        // Record: QSO Submode use enumeration values for interoperability
	SWL                            Field = "SWL"                            // Record: indicates that the QSO information pertains to an SWL report
	TEN_TEN                        Field = "TEN_TEN"                        // Record: Ten-Ten number with a value greater than 0
	TIME_OFF                       Field = "TIME_OFF"                       // Record: HHMM or HHMMSS in UTC in the absence of <QSO_DATE_OFF>, the QSO duration is less than 24 hours. For example, the following is a QSO starting at 14 July 2020 23:55 and finishing at 15 July 2020 01:00: <QSO_DATE:8>20200714 <TIME_ON:4>2355 <TIME_OFF:4>0100
	TIME_ON                        Field = "TIME_ON"                        // Record: HHMM or HHMMSS in UTC
	TX_PWR                         Field = "TX_PWR"                         // Record: the logging station's power in Watts with a value greater than or equal to 0
	UKSMG                          Field = "UKSMG"                          // Record: the contacted station's UKSMG member number with a value greater than 0
	USACA_COUNTIES                 Field = "USACA_COUNTIES"                 // Record: two US counties in the case where the contacted station is located on a border between two counties, representing counties credited to the QSO for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire
	USERDEF1                       Field = "USERDEF1"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF2                       Field = "USERDEF2"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF3                       Field = "USERDEF3"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF4                       Field = "USERDEF4"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF5                       Field = "USERDEF5"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF6                       Field = "USERDEF6"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF7                       Field = "USERDEF7"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF8                       Field = "USERDEF8"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF9                       Field = "USERDEF9"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	VE_PROV                        Field = "VE_PROV"                        // Deprecated: Record: import-only: use STATE instead
	VUCC_GRIDS                     Field = "VUCC_GRIDS"                     // Record: two or four adjacent Maidenhead grid locators, each four or six characters long, representing the contacted station's grid squares credited to the QSO for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07
	WEB                            Field = "WEB"                            // Record: the contacted station's URL
	WWFF_REF                       Field = "WWFF_REF"                       // Record: the contacted station's WWFF (World Wildlife Flora & Fauna) reference
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Field specifications.
var lookupList = []Spec{
	{Key: "ADDRESS", DataType: "MULTILINESTRING", Description: "the contacted station's complete mailing address: full name, street address, city, postal code, and country", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "ADDRESS_INTL", DataType: "INTLMULTILINESTRING", Description: "the contacted station's complete mailing address: full name, street address, city, postal code, and country", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "ADIF_VER", DataType: "STRING", Description: "identifies the version of ADIF used in this file in the format X.Y.Z where X is an integer designating the ADIF epoch Y is an integer between 0 and 9 designating the major version Z is an integer between 0 and 9 designating the minor version", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "AGE", DataType: "NUMBER", Description: "the contacted station's operator's age in years in the range 0 to 120 (inclusive)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 120, IsImportOnly: false, Comments: ""},
	{Key: "ALTITUDE", DataType: "NUMBER", Description: "the height of the contacted station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <ALTITUDE:4>1500 and 10.5 m is <ALTITUDE:4>10.5", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "ANT_AZ", DataType: "NUMBER", Description: "the logging station's antenna azimuth, in degrees with a value between 0 to 360 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 370 is exported as 10). True north is 0 degrees with values increasing in a clockwise direction.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 360, IsImportOnly: false, Comments: ""},
	{Key: "ANT_EL", DataType: "NUMBER", Description: "the logging station's antenna elevation, in degrees with a value between -90 to 90 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 100 is exported as 80). The horizon is 0 degrees with values increasing as the angle moves in an upward direction.", IsHeaderField: false, MinimumValue: -90, MaximumValue: 90, IsImportOnly: false, Comments: ""},
	{Key: "ANT_PATH", DataType: "ENUMERATION", Description: "the signal path", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_2XQSL", DataType: "ENUMERATION", Description: "ARRL LOTW: [Y, N] APP_LOTW_2xQSL =\"Y\" Indicates this confirmation is considered to be \"two-way\" in the QSO's Mode. It is a confirmation that counts towards the mode-specific awards & endorsement in the WAS program. Only present if QSL_RCVD=\"Y\".", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_CQZ_INFERRED", DataType: "STRING", Description: "ARRL LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_CQZ_INVALID", DataType: "STRING", Description: "ARRL LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_CQZ_USERINVALID", DataType: "STRING", Description: "ARRL LOTW: Present only if the QSLing station included an invalid CQ zone value in its station location uploaded to LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_CREDIT_GRANTED", DataType: "STRING", Description: "ARRL LOTW: Award credits granted based upon this QSL. Awards per LOTW award enumeration.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_CREDIT_SUBMITTED", DataType: "STRING", Description: "ARRL LOTW: Award credits for which this QSL has been submitted on an award application to sponsor. Awards per LOTW award enumeration.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_DXCC_ENTITY_STATUS", DataType: "ENUMERATION", Description: "ARRL LOTW: [Current, Deleted] Indicates whether the ARRL DXCC entity is \"Current\" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or \"Deleted\"", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_EOF", DataType: "STRING", Description: "ARRL LOTW: End of File marker. No other fields expected in this record.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_GRIDSQUARE_INVALID", DataType: "STRING", Description: "ARRL LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_ITUZ_INFERRED", DataType: "STRING", Description: "ARRL LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_ITUZ_INVALID", DataType: "STRING", Description: "ARRL LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_ITUZ_USERINVALID", DataType: "STRING", Description: "ARRL LOTW: Present only if the QSLing station included an invalid ITU zone value in its station location uploaded to LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_LASTQSL", DataType: "STRING", Description: "ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl=\"yes\". Date and time at which the most recent QSL record in this report was received. This will be the QSL date/time of the first QSO record in the file. This value only applies to the records in this file.", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_LASTQSORX", DataType: "STRING", Description: "ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl=\"no\". Date and time at which the most recent QSO record in this report was received. This will be the QSO received (uploaded) date/time of the first QSO record in the file. This value only applies to the records in this file.", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_MODE", DataType: "ENUMERATION", Description: "ARRL LOTW: Field is present when the mode recorded in Logbook cannot be unambiguously represented as an ADIF-compliant value", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_MODEGROUP", DataType: "ENUMERATION", Description: "ARRL LOTW: [CW, PHONE, DATA] The mode group indicates whether the QSO's mode counts towards the CW, Phone or Digital awards respectively in the DXCC program. A mode group of CW indicates the QSO's mode counts towards the CW endorsement on WAS awards. A mode group of PHONE indicates the QSO's mode counts towards the WAS Phone award. A mode group of DATA indicates the QSO's mode counts towards the CQ WPX Digital award.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_MY_DXCC_ENTITY_STATUS", DataType: "ENUMERATION", Description: "ARRL LOTW: [Current, Deleted] Indicates whether the Logging station's ARRL DXCC entity is \"Current\" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or \"Deleted\"", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_NPSUNIT", DataType: "STRING", Description: "ARRL LOTW: Present only if the QSLing station included a valid NPOTA Park code or a comma-delimited list of valid NPOTA Park codes", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_NUMREC", DataType: "POSITIVEINTEGER", Description: "ARRL LOTW: Number of QSO records in this download", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_OWNCALL", DataType: "STRING", Description: "ARRL LOTW: \"own\" call sign of the station making the contact. Present only if qso_withown=\"yes\". DEPRECATED -- In the near future, this field will no longer be present in the report output. Programs that import LOTW report files should not rely on the presence of this field. Use STATION_CALLSIGN instead.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_QSLMODE", DataType: "ENUMERATION", Description: "ARRL LOTW: The Mode that your QSO partner indicated on their side of this QSO. Only present if QSL_RCVD=\"Y\" and APP_LOTW_2xQSL=\"N\".", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_QSO_TIMESTAMP", DataType: "STRING", Description: "ARRL LOTW: Combined QSO Date & Time in ISO-8601 format", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_RXQSL", DataType: "STRING", Description: "ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which the confirmation match (QSL) was made/updated at LOTW. Present only if QSL_RCVD=\"Y\"", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_LOTW_RXQSO", DataType: "STRING", Description: "ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which QSO record was inserted/updated at LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_QRZLOG_LOGID", DataType: "POSITIVEINTEGER", Description: "QRZ LOG: References the internal ID of a log record in the QRZ system.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_QRZLOG_QSLDATE", DataType: "DATE", Description: "QRZ LOG: QSL Date (e.g. 20210126) This field contains the first date on which we saw a QRZ accepted confirmation for this contact. This could have been a confirmation as a result of matching contacts in the QRZ system or a credit imported from Logbook of The World.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_QRZLOG_STATUS", DataType: "ENUMERATION", Description: "QRZ LOG: Status [C: Confirmed, A: Reserved for future use, N: Not Confirmed, 2: Confirmation Requested, S: Confirmation Requested, Request Seen, R: Confirmation Requested, Request Rejected]", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "APP_SKCCLOGGER_KEYTYPE", DataType: "ENUMERATION", Description: "Straight Key Century Club (SKCC): [BUG, SK, SS] CW KEY TYPE (SS=Side Swiper / Cootie, SK=Straight, B=Bug)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "ARRL_SECT", DataType: "ENUMERATION", Description: "the contacted station's ARRL section", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "AWARD_GRANTED", DataType: "SPONSOREDAWARDLIST", Description: "the list of awards granted by a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's \"Award Account\" between an award sponsor and the user. For example, in response to a request \"send me a list of the awards granted to AA6YQ\", this might be received: <CALL:5>AA6YQ <AWARD_GRANTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "AWARD_SUBMITTED", DataType: "SPONSOREDAWARDLIST", Description: "the list of awards submitted to a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's \"Award Account\" between an award sponsor and the user. For example, AA6YQ might submit a request for awards by sending the following: <CALL:5>AA6YQ <AWARD_SUBMITTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "A_INDEX", DataType: "NUMBER", Description: "the geomagnetic A index at the time of the QSO in the range 0 to 400 (inclusive)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 400, IsImportOnly: false, Comments: ""},
	{Key: "BAND", DataType: "ENUMERATION", Description: "QSO Band", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "BAND_RX", DataType: "ENUMERATION", Description: "in a split frequency QSO, the logging station's receiving band", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CALL", DataType: "STRING", Description: "the contacted station's callsign", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CHECK", DataType: "STRING", Description: "contest check (e.g. for ARRL Sweepstakes)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CLASS", DataType: "STRING", Description: "contest class (e.g. for ARRL Field Day)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CLUBLOG_QSO_UPLOAD_DATE", DataType: "DATE", Description: "the date the QSO was last uploaded to the Club Log online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CLUBLOG_QSO_UPLOAD_STATUS", DataType: "ENUMERATION", Description: "the upload status of the QSO on the Club Log online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CNTY", DataType: "ENUMERATION", Description: "the contacted station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CNTY_ALT", DataType: "SECONDARYADMINISTRATIVESUBDIVISIONLISTALT", Description: "a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the contacted station See the Data Type for details.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "COMMENT", DataType: "STRING", Description: "comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG field recommended use: information of interest to the contacted station's operator", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "COMMENT_INTL", DataType: "INTLSTRING", Description: "comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG_INTL field recommended use: information of interest to the contacted station's operator", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CONT", DataType: "ENUMERATION", Description: "the contacted station's Continent", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CONTACTED_OP", DataType: "STRING", Description: "the callsign of the individual operating the contacted station", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CONTEST_ID", DataType: "STRING", Description: "QSO Contest Identifier use enumeration values for interoperability", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "COUNTRY", DataType: "STRING", Description: "the contacted station's DXCC entity name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "COUNTRY_INTL", DataType: "INTLSTRING", Description: "the contacted station's DXCC entity name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CQZ", DataType: "POSITIVEINTEGER", Description: "the contacted station's CQ Zone in the range 1 to 40 (inclusive)", IsHeaderField: false, MinimumValue: 1, MaximumValue: 40, IsImportOnly: false, Comments: ""},
	{Key: "CREATED_TIMESTAMP", DataType: "STRING", Description: "identifies the UTC date and time that the file was created in the format of 15 characters YYYYMMDD HHMMSS where YYYYMMDD is a Date data type HHMMSS is a 6 character Time data type", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CREDIT_GRANTED", DataType: "CREDITLIST,AWARDLIST", Description: "the list of credits granted to this QSO Use of data type AwardList and enumeration Award are import-only", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "CREDIT_SUBMITTED", DataType: "CREDITLIST,AWARDLIST", Description: "the list of credits sought for this QSO Use of data type AwardList and enumeration Award are import-only", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "DARC_DOK", DataType: "ENUMERATION", Description: "the contacted station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <DARC_DOK:3>A01", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "DCL_QSLRDATE", DataType: "DATE", Description: "date QSL received from DCL (only valid if DCL_QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "DCL_QSLSDATE", DataType: "DATE", Description: "date QSL sent to DCL (only valid if DCL_QSL_SENT is Y, Q, or I)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "DCL_QSL_RCVD", DataType: "ENUMERATION", Description: "DCL QSL received status Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "DCL_QSL_SENT", DataType: "ENUMERATION", Description: "DCL QSL sent status Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "DISTANCE", DataType: "NUMBER", Description: "the distance between the logging station and the contacted station in kilometers via the specified signal path with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "DXCC", DataType: "ENUMERATION", Description: "the contacted station's DXCC Entity Code <DXCC:1>0 means that the contacted station is known not to be within a DXCC entity.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "EMAIL", DataType: "STRING", Description: "the contacted station's email address", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "EQSL_AG", DataType: "ENUMERATION", Description: "indicates whether the QSO is known to be \"Authenticity Guaranteed\" by eQSL", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "EQSL_QSLRDATE", DataType: "DATE", Description: "date QSL received from eQSL.cc (only valid if EQSL_QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "EQSL_QSLSDATE", DataType: "DATE", Description: "date QSL sent to eQSL.cc (only valid if EQSL_QSL_SENT is Y, Q, or I)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "EQSL_QSL_RCVD", DataType: "ENUMERATION", Description: "eQSL.cc QSL received status instead of V (import-only) use <CREDIT_GRANTED:42>CQWAZ:eqsl,CQWAZ_BAND:eqsl,CQWAZ_MODE:eqsl Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "EQSL_QSL_SENT", DataType: "ENUMERATION", Description: "eQSL.cc QSL sent status Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "EQ_CALL", DataType: "STRING", Description: "the contacted station's owner's callsign", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "FISTS", DataType: "POSITIVEINTEGER", Description: "the contacted station's FISTS CW Club member number with a value greater than 0.", IsHeaderField: false, MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "FISTS_CC", DataType: "POSITIVEINTEGER", Description: "the contacted station's FISTS CW Club Century Certificate (CC) number with a value greater than 0.", IsHeaderField: false, MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "FORCE_INIT", DataType: "BOOLEAN", Description: "new EME \"initial\"", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "FREQ", DataType: "NUMBER", Description: "QSO frequency in Megahertz", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "FREQ_RX", DataType: "NUMBER", Description: "in a split frequency QSO, the logging station's receiving frequency in Megahertz", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "GRIDSQUARE", DataType: "GRIDSQUARE", Description: "the contacted station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in GRIDSQUARE and the additional 2 or 4 characters in the GRIDSQUARE_EXT field", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "GRIDSQUARE_EXT", DataType: "GRIDSQUAREEXT", Description: "for a contacted station's 10-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9 and 10. For a contacted station's 12-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <GRIDSQUARE:8>FN01MH42 <GRIDSQUARE_EXT:2>BQ", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "GUEST_OP", DataType: "STRING", Description: "import-only: use OPERATOR instead", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: true, Comments: ""},
	{Key: "HAMLOGEU_QSO_UPLOAD_DATE", DataType: "DATE", Description: "the date the QSO was last uploaded to the HAMLOG.EU online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "HAMLOGEU_QSO_UPLOAD_STATUS", DataType: "ENUMERATION", Description: "the upload status of the QSO on the HAMLOG.EU online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "HAMQTH_QSO_UPLOAD_DATE", DataType: "DATE", Description: "the date the QSO was last uploaded to the HamQTH.com online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "HAMQTH_QSO_UPLOAD_STATUS", DataType: "ENUMERATION", Description: "the upload status of the QSO on the HamQTH.com online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "HRDLOG_QSO_UPLOAD_DATE", DataType: "DATE", Description: "the date the QSO was last uploaded to the HRDLog.net online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "HRDLOG_QSO_UPLOAD_STATUS", DataType: "ENUMERATION", Description: "the upload status of the QSO on the HRDLog.net online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "IOTA", DataType: "IOTAREFNO", Description: "the contacted station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "IOTA_ISLAND_ID", DataType: "POSITIVEINTEGER", Description: "the contacted station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]", IsHeaderField: false, MinimumValue: 1, MaximumValue: 99999999, IsImportOnly: false, Comments: ""},
	{Key: "ITUZ", DataType: "POSITIVEINTEGER", Description: "the contacted station's ITU zone in the range 1 to 90 (inclusive)", IsHeaderField: false, MinimumValue: 1, MaximumValue: 90, IsImportOnly: false, Comments: ""},
	{Key: "K_INDEX", DataType: "INTEGER", Description: "the geomagnetic K index at the time of the QSO in the range 0 to 9 (inclusive)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 9, IsImportOnly: false, Comments: ""},
	{Key: "LAT", DataType: "LOCATION", Description: "the contacted station's latitude", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "LON", DataType: "LOCATION", Description: "the contacted station's longitude", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "LOTW_QSLRDATE", DataType: "DATE", Description: "date QSL received from ARRL Logbook of the World (only valid if LOTW_QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "LOTW_QSLSDATE", DataType: "DATE", Description: "date QSL sent to ARRL Logbook of the World (only valid if LOTW_QSL_SENT is Y, Q, or I)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "LOTW_QSL_RCVD", DataType: "ENUMERATION", Description: "ARRL Logbook of the World QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "LOTW_QSL_SENT", DataType: "ENUMERATION", Description: "ARRL Logbook of the World QSL sent status Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MAX_BURSTS", DataType: "NUMBER", Description: "maximum length of meteor scatter bursts heard by the logging station, in seconds with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MODE", DataType: "ENUMERATION", Description: "QSO Mode", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MORSE_KEY_INFO", DataType: "STRING", Description: "details of the contacted station's Morse key (e.g. make, model, etc). Example: <MORSE_KEY_INFO:16>Begali Sculpture", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MORSE_KEY_TYPE", DataType: "ENUMERATION", Description: "the contacted station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MS_SHOWER", DataType: "STRING", Description: "For Meteor Scatter QSOs, the name of the meteor shower in progress", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_ALTITUDE", DataType: "NUMBER", Description: "the height of the logging station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <MY_ALTITUDE:4>1500 and 10.5 m is <MY_ALTITUDE:4>10.5", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_ANTENNA", DataType: "STRING", Description: "the logging station's antenna", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_ANTENNA_INTL", DataType: "INTLSTRING", Description: "the logging station's antenna", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_ARRL_SECT", DataType: "ENUMERATION", Description: "the logging station's ARRL section", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_CITY", DataType: "STRING", Description: "the logging station's city", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_CITY_INTL", DataType: "INTLSTRING", Description: "the logging station's city", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_CNTY", DataType: "ENUMERATION", Description: "the logging station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_CNTY_ALT", DataType: "SECONDARYADMINISTRATIVESUBDIVISIONLISTALT", Description: "a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the logging station See the Data Type for details.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_COUNTRY", DataType: "STRING", Description: "the logging station's DXCC entity name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_COUNTRY_INTL", DataType: "INTLSTRING", Description: "the logging station's DXCC entity name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_CQ_ZONE", DataType: "POSITIVEINTEGER", Description: "the logging station's CQ Zone in the range 1 to 40 (inclusive)", IsHeaderField: false, MinimumValue: 1, MaximumValue: 40, IsImportOnly: false, Comments: ""},
	{Key: "MY_DARC_DOK", DataType: "ENUMERATION", Description: "the logging station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <MY_DARC_DOK:3>A01", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_DXCC", DataType: "ENUMERATION", Description: "the logging station's DXCC Entity Code <MY_DXCC:1>0 means that the logging station is known not to be within a DXCC entity.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_FISTS", DataType: "POSITIVEINTEGER", Description: "the logging station's FISTS CW Club member number with a value greater than 0.", IsHeaderField: false, MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_GRIDSQUARE", DataType: "GRIDSQUARE", Description: "the logging station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in MY_GRIDSQUARE and the additional 2 or 4 characters in the MY_GRIDSQUARE_EXT field", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_GRIDSQUARE_EXT", DataType: "GRIDSQUAREEXT", Description: "for a logging station's 10-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9 and 10. For a logging station's 12-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <MY_GRIDSQUARE:8>FN01MH42 <MY_GRIDSQUARE_EXT:2>BQ", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_IOTA", DataType: "IOTAREFNO", Description: "the logging station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_IOTA_ISLAND_ID", DataType: "POSITIVEINTEGER", Description: "the logging station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]", IsHeaderField: false, MinimumValue: 1, MaximumValue: 99999999, IsImportOnly: false, Comments: ""},
	{Key: "MY_ITU_ZONE", DataType: "POSITIVEINTEGER", Description: "the logging station's ITU zone in the range 1 to 90 (inclusive)", IsHeaderField: false, MinimumValue: 1, MaximumValue: 90, IsImportOnly: false, Comments: ""},
	{Key: "MY_LAT", DataType: "LOCATION", Description: "the logging station's latitude", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_LON", DataType: "LOCATION", Description: "the logging station's longitude", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_MORSE_KEY_INFO", DataType: "STRING", Description: "details of the logging station's Morse key (e.g. make, model, etc). Example: <MY_MORSE_KEY_INFO:16>Begali Sculpture", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_MORSE_KEY_TYPE", DataType: "ENUMERATION", Description: "the logging station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_NAME", DataType: "STRING", Description: "the logging operator's name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_NAME_INTL", DataType: "INTLSTRING", Description: "the logging operator's name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_POSTAL_CODE", DataType: "STRING", Description: "the logging station's postal code", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_POSTAL_CODE_INTL", DataType: "INTLSTRING", Description: "the logging station's postal code", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_POTA_REF", DataType: "POTAREFLIST", Description: "a comma-delimited list of one or more of the logging station's POTA (Parks on the Air) reference(s). Examples: <MY_POTA_REF:6>K-0059 <MY_POTA_REF:7>K-10000 <MY_POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_RIG", DataType: "STRING", Description: "description of the logging station's equipment", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_RIG_INTL", DataType: "INTLSTRING", Description: "description of the logging station's equipment", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_SIG", DataType: "STRING", Description: "special interest activity or event", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_SIG_INFO", DataType: "STRING", Description: "special interest activity or event information", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_SIG_INFO_INTL", DataType: "INTLSTRING", Description: "special interest activity or event information", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_SIG_INTL", DataType: "INTLSTRING", Description: "special interest activity or event", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_SOTA_REF", DataType: "SOTAREF", Description: "the logging station's International SOTA Reference.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_STATE", DataType: "ENUMERATION", Description: "the code for the logging station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_STREET", DataType: "STRING", Description: "the logging station's street", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_STREET_INTL", DataType: "INTLSTRING", Description: "the logging station's street", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_USACA_COUNTIES", DataType: "SECONDARYSUBDIVISIONLIST", Description: "two US counties in the case where the logging station is located on a border between two counties, representing counties that the contacted station may claim for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_VUCC_GRIDS", DataType: "GRIDSQUARELIST", Description: "two or four adjacent Maidenhead grid locators, each four or six characters long, representing the logging station's grid squares that the contacted station may claim for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "MY_WWFF_REF", DataType: "WWFFREF", Description: "the logging station's WWFF (World Wildlife Flora & Fauna) reference", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "NAME", DataType: "STRING", Description: "the contacted station's operator's name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "NAME_INTL", DataType: "INTLSTRING", Description: "the contacted station's operator's name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "NOTES", DataType: "MULTILINESTRING", Description: "QSO notes recommended use: information of interest to the logging station's operator", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "NOTES_INTL", DataType: "INTLMULTILINESTRING", Description: "QSO notes recommended use: information of interest to the logging station's operator", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "NR_BURSTS", DataType: "INTEGER", Description: "the number of meteor scatter bursts heard by the logging station with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "NR_PINGS", DataType: "INTEGER", Description: "the number of meteor scatter pings heard by the logging station with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "OPERATOR", DataType: "STRING", Description: "the logging operator's callsign if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "OWNER_CALLSIGN", DataType: "STRING", Description: "the callsign of the owner of the station used to log the contact (the callsign of the OPERATOR's host) if OWNER_CALLSIGN is absent, STATION_CALLSIGN shall be treated as both the logging station's callsign and the callsign of the owner of the station", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "PFX", DataType: "STRING", Description: "the contacted station's WPX prefix", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "POTA_REF", DataType: "POTAREFLIST", Description: "a comma-delimited list of one or more of the contacted station's POTA (Parks on the Air) reference(s). Examples: <POTA_REF:6>K-5033 <POTA_REF:13>VE-5082@CA-AB <POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "PRECEDENCE", DataType: "STRING", Description: "contest precedence (e.g. for ARRL Sweepstakes)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "PROGRAMID", DataType: "STRING", Description: "identifies the name of the logger, converter, or utility that created or processed this ADIF file To help avoid name clashes, the ADIF PROGRAMID Register provides a voluntary list of PROGRAMID values.", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "PROGRAMVERSION", DataType: "STRING", Description: "identifies the version of the logger, converter, or utility that created or processed this ADIF file", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "PROP_MODE", DataType: "ENUMERATION", Description: "QSO propagation mode", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "PUBLIC_KEY", DataType: "STRING", Description: "public encryption key", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QRZCOM_QSO_DOWNLOAD_DATE", DataType: "DATE", Description: "date QSO downloaded from QRZ.COM logbook", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QRZCOM_QSO_DOWNLOAD_STATUS", DataType: "ENUMERATION", Description: "QRZ.COM logbook QSO download status", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QRZCOM_QSO_UPLOAD_DATE", DataType: "DATE", Description: "the date the QSO was last uploaded to the QRZ.COM online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QRZCOM_QSO_UPLOAD_STATUS", DataType: "ENUMERATION", Description: "the upload status of the QSO on the QRZ.COM online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSLMSG", DataType: "MULTILINESTRING", Description: "a message for the contacted station's operator to be incorporated in a paper or electronic QSL", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSLMSG_INTL", DataType: "INTLMULTILINESTRING", Description: "a message for the contacted station's operator to be incorporated in a paper or electronic QSL", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSLMSG_RCVD", DataType: "MULTILINESTRING", Description: "a message addressed to the logging station's operator incorporated in a paper or electronic QSL", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSLRDATE", DataType: "DATE", Description: "QSL received date (only valid if QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSLSDATE", DataType: "DATE", Description: "QSL sent date (only valid if QSL_SENT is Y, Q, or I)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSL_RCVD", DataType: "ENUMERATION", Description: "QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSL_RCVD_VIA", DataType: "ENUMERATION", Description: "if QSL_RCVD is set to 'Y' or 'V', the means by which the QSL was received by the logging station; otherwise, the means by which the logging station requested or intends to request that the QSL be conveyed. (Note: 'V' is import-only) use of M (manager) is import-only", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSL_SENT", DataType: "ENUMERATION", Description: "QSL sent status Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSL_SENT_VIA", DataType: "ENUMERATION", Description: "if QSL_SENT is set to 'Y', the means by which the QSL was sent by the logging station; otherwise, the means by which the logging station intends to convey the QSL use of M (manager) is import-only", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSL_VIA", DataType: "STRING", Description: "the contacted station's QSL route", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSO_COMPLETE", DataType: "ENUMERATION", Description: "indicates whether the QSO was complete from the perspective of the logging station Y - yes N - no NIL - not heard ? - uncertain", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSO_DATE", DataType: "DATE", Description: "date on which the QSO started", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSO_DATE_OFF", DataType: "DATE", Description: "date on which the QSO ended", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QSO_RANDOM", DataType: "BOOLEAN", Description: "indicates whether the QSO was random or scheduled", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QTH", DataType: "STRING", Description: "the contacted station's city", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "QTH_INTL", DataType: "INTLSTRING", Description: "the contacted station's city", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "REGION", DataType: "ENUMERATION", Description: "the contacted station's WAE or CQ entity contained within a DXCC entity. the value None indicates that the WAE or CQ entity is the DXCC entity in the DXCC field. nothing can be inferred from the absence of the REGION field", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "RIG", DataType: "MULTILINESTRING", Description: "description of the contacted station's equipment", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "RIG_INTL", DataType: "INTLMULTILINESTRING", Description: "description of the contacted station's equipment", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "RST_RCVD", DataType: "STRING", Description: "signal report from the contacted station", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "RST_SENT", DataType: "STRING", Description: "signal report sent to the contacted station", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "RX_PWR", DataType: "NUMBER", Description: "the contacted station's transmitter power in Watts with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SAT_MODE", DataType: "STRING", Description: "satellite mode - a code representing the satellite's uplink band and downlink band", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SAT_NAME", DataType: "STRING", Description: "name of satellite", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SFI", DataType: "INTEGER", Description: "the solar flux at the time of the QSO in the range 0 to 300 (inclusive).", IsHeaderField: false, MinimumValue: 0, MaximumValue: 300, IsImportOnly: false, Comments: ""},
	{Key: "SIG", DataType: "STRING", Description: "the name of the contacted station's special activity or interest group", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SIG_INFO", DataType: "STRING", Description: "information associated with the contacted station's activity or interest group", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SIG_INFO_INTL", DataType: "INTLSTRING", Description: "information associated with the contacted station's activity or interest group", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SIG_INTL", DataType: "INTLSTRING", Description: "the name of the contacted station's special activity or interest group", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SILENT_KEY", DataType: "BOOLEAN", Description: "'Y' indicates that the contacted station's operator is now a Silent Key.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SKCC", DataType: "STRING", Description: "the contacted station's Straight Key Century Club (SKCC) member information", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SOTA_REF", DataType: "SOTAREF", Description: "the contacted station's International SOTA Reference.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SRX", DataType: "INTEGER", Description: "contest QSO received serial number with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SRX_STRING", DataType: "STRING", Description: "contest QSO received information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "STATE", DataType: "ENUMERATION", Description: "the code for the contacted station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "STATION_CALLSIGN", DataType: "STRING", Description: "the logging station's callsign (the callsign used over the air) if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "STX", DataType: "INTEGER", Description: "contest QSO transmitted serial number with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "STX_STRING", DataType: "STRING", Description: "contest QSO transmitted information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SUBMODE", DataType: "STRING", Description: "QSO Submode use enumeration values for interoperability", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "SWL", DataType: "BOOLEAN", Description: "indicates that the QSO information pertains to an SWL report", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "TEN_TEN", DataType: "POSITIVEINTEGER", Description: "Ten-Ten number with a value greater than 0", IsHeaderField: false, MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "TIME_OFF", DataType: "TIME", Description: "HHMM or HHMMSS in UTC in the absence of <QSO_DATE_OFF>, the QSO duration is less than 24 hours. For example, the following is a QSO starting at 14 July 2020 23:55 and finishing at 15 July 2020 01:00: <QSO_DATE:8>20200714 <TIME_ON:4>2355 <TIME_OFF:4>0100", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "TIME_ON", DataType: "TIME", Description: "HHMM or HHMMSS in UTC", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "TX_PWR", DataType: "NUMBER", Description: "the logging station's power in Watts with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "UKSMG", DataType: "POSITIVEINTEGER", Description: "the contacted station's UKSMG member number with a value greater than 0", IsHeaderField: false, MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "USACA_COUNTIES", DataType: "SECONDARYSUBDIVISIONLIST", Description: "two US counties in the case where the contacted station is located on a border between two counties, representing counties credited to the QSO for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "USERDEF1", DataType: "STRING", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "USERDEF2", DataType: "STRING", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "USERDEF3", DataType: "STRING", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "USERDEF4", DataType: "STRING", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "USERDEF5", DataType: "STRING", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "USERDEF6", DataType: "STRING", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "USERDEF7", DataType: "STRING", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "USERDEF8", DataType: "STRING", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "USERDEF9", DataType: "STRING", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "VE_PROV", DataType: "STRING", Description: "import-only: use STATE instead", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: true, Comments: ""},
	{Key: "VUCC_GRIDS", DataType: "GRIDSQUARELIST", Description: "two or four adjacent Maidenhead grid locators, each four or six characters long, representing the contacted station's grid squares credited to the QSO for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "WEB", DataType: "STRING", Description: "the contacted station's URL", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "WWFF_REF", DataType: "WWFFREF", Description: "the contacted station's WWFF (World Wildlife Flora & Fauna) reference", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
}

// lookupMap contains all known Field specifications.
var lookupMap = map[Field]*Spec{
	ADDRESS:                        &lookupList[0],
	ADDRESS_INTL:                   &lookupList[1],
	ADIF_VER:                       &lookupList[2],
	AGE:                            &lookupList[3],
	ALTITUDE:                       &lookupList[4],
	ANT_AZ:                         &lookupList[5],
	ANT_EL:                         &lookupList[6],
	ANT_PATH:                       &lookupList[7],
	APP_LOTW_2XQSL:                 &lookupList[8],
	APP_LOTW_CQZ_INFERRED:          &lookupList[9],
	APP_LOTW_CQZ_INVALID:           &lookupList[10],
	APP_LOTW_CQZ_USERINVALID:       &lookupList[11],
	APP_LOTW_CREDIT_GRANTED:        &lookupList[12],
	APP_LOTW_CREDIT_SUBMITTED:      &lookupList[13],
	APP_LOTW_DXCC_ENTITY_STATUS:    &lookupList[14],
	APP_LOTW_EOF:                   &lookupList[15],
	APP_LOTW_GRIDSQUARE_INVALID:    &lookupList[16],
	APP_LOTW_ITUZ_INFERRED:         &lookupList[17],
	APP_LOTW_ITUZ_INVALID:          &lookupList[18],
	APP_LOTW_ITUZ_USERINVALID:      &lookupList[19],
	APP_LOTW_LASTQSL:               &lookupList[20],
	APP_LOTW_LASTQSORX:             &lookupList[21],
	APP_LOTW_MODE:                  &lookupList[22],
	APP_LOTW_MODEGROUP:             &lookupList[23],
	APP_LOTW_MY_DXCC_ENTITY_STATUS: &lookupList[24],
	APP_LOTW_NPSUNIT:               &lookupList[25],
	APP_LOTW_NUMREC:                &lookupList[26],
	APP_LOTW_OWNCALL:               &lookupList[27],
	APP_LOTW_QSLMODE:               &lookupList[28],
	APP_LOTW_QSO_TIMESTAMP:         &lookupList[29],
	APP_LOTW_RXQSL:                 &lookupList[30],
	APP_LOTW_RXQSO:                 &lookupList[31],
	APP_QRZLOG_LOGID:               &lookupList[32],
	APP_QRZLOG_QSLDATE:             &lookupList[33],
	APP_QRZLOG_STATUS:              &lookupList[34],
	APP_SKCCLOGGER_KEYTYPE:         &lookupList[35],
	ARRL_SECT:                      &lookupList[36],
	AWARD_GRANTED:                  &lookupList[37],
	AWARD_SUBMITTED:                &lookupList[38],
	A_INDEX:                        &lookupList[39],
	BAND:                           &lookupList[40],
	BAND_RX:                        &lookupList[41],
	CALL:                           &lookupList[42],
	CHECK:                          &lookupList[43],
	CLASS:                          &lookupList[44],
	CLUBLOG_QSO_UPLOAD_DATE:        &lookupList[45],
	CLUBLOG_QSO_UPLOAD_STATUS:      &lookupList[46],
	CNTY:                           &lookupList[47],
	CNTY_ALT:                       &lookupList[48],
	COMMENT:                        &lookupList[49],
	COMMENT_INTL:                   &lookupList[50],
	CONT:                           &lookupList[51],
	CONTACTED_OP:                   &lookupList[52],
	CONTEST_ID:                     &lookupList[53],
	COUNTRY:                        &lookupList[54],
	COUNTRY_INTL:                   &lookupList[55],
	CQZ:                            &lookupList[56],
	CREATED_TIMESTAMP:              &lookupList[57],
	CREDIT_GRANTED:                 &lookupList[58],
	CREDIT_SUBMITTED:               &lookupList[59],
	DARC_DOK:                       &lookupList[60],
	DCL_QSLRDATE:                   &lookupList[61],
	DCL_QSLSDATE:                   &lookupList[62],
	DCL_QSL_RCVD:                   &lookupList[63],
	DCL_QSL_SENT:                   &lookupList[64],
	DISTANCE:                       &lookupList[65],
	DXCC:                           &lookupList[66],
	EMAIL:                          &lookupList[67],
	EQSL_AG:                        &lookupList[68],
	EQSL_QSLRDATE:                  &lookupList[69],
	EQSL_QSLSDATE:                  &lookupList[70],
	EQSL_QSL_RCVD:                  &lookupList[71],
	EQSL_QSL_SENT:                  &lookupList[72],
	EQ_CALL:                        &lookupList[73],
	FISTS:                          &lookupList[74],
	FISTS_CC:                       &lookupList[75],
	FORCE_INIT:                     &lookupList[76],
	FREQ:                           &lookupList[77],
	FREQ_RX:                        &lookupList[78],
	GRIDSQUARE:                     &lookupList[79],
	GRIDSQUARE_EXT:                 &lookupList[80],
	GUEST_OP:                       &lookupList[81],
	HAMLOGEU_QSO_UPLOAD_DATE:       &lookupList[82],
	HAMLOGEU_QSO_UPLOAD_STATUS:     &lookupList[83],
	HAMQTH_QSO_UPLOAD_DATE:         &lookupList[84],
	HAMQTH_QSO_UPLOAD_STATUS:       &lookupList[85],
	HRDLOG_QSO_UPLOAD_DATE:         &lookupList[86],
	HRDLOG_QSO_UPLOAD_STATUS:       &lookupList[87],
	IOTA:                           &lookupList[88],
	IOTA_ISLAND_ID:                 &lookupList[89],
	ITUZ:                           &lookupList[90],
	K_INDEX:                        &lookupList[91],
	LAT:                            &lookupList[92],
	LON:                            &lookupList[93],
	LOTW_QSLRDATE:                  &lookupList[94],
	LOTW_QSLSDATE:                  &lookupList[95],
	LOTW_QSL_RCVD:                  &lookupList[96],
	LOTW_QSL_SENT:                  &lookupList[97],
	MAX_BURSTS:                     &lookupList[98],
	MODE:                           &lookupList[99],
	MORSE_KEY_INFO:                 &lookupList[100],
	MORSE_KEY_TYPE:                 &lookupList[101],
	MS_SHOWER:                      &lookupList[102],
	MY_ALTITUDE:                    &lookupList[103],
	MY_ANTENNA:                     &lookupList[104],
	MY_ANTENNA_INTL:                &lookupList[105],
	MY_ARRL_SECT:                   &lookupList[106],
	MY_CITY:                        &lookupList[107],
	MY_CITY_INTL:                   &lookupList[108],
	MY_CNTY:                        &lookupList[109],
	MY_CNTY_ALT:                    &lookupList[110],
	MY_COUNTRY:                     &lookupList[111],
	MY_COUNTRY_INTL:                &lookupList[112],
	MY_CQ_ZONE:                     &lookupList[113],
	MY_DARC_DOK:                    &lookupList[114],
	MY_DXCC:                        &lookupList[115],
	MY_FISTS:                       &lookupList[116],
	MY_GRIDSQUARE:                  &lookupList[117],
	MY_GRIDSQUARE_EXT:              &lookupList[118],
	MY_IOTA:                        &lookupList[119],
	MY_IOTA_ISLAND_ID:              &lookupList[120],
	MY_ITU_ZONE:                    &lookupList[121],
	MY_LAT:                         &lookupList[122],
	MY_LON:                         &lookupList[123],
	MY_MORSE_KEY_INFO:              &lookupList[124],
	MY_MORSE_KEY_TYPE:              &lookupList[125],
	MY_NAME:                        &lookupList[126],
	MY_NAME_INTL:                   &lookupList[127],
	MY_POSTAL_CODE:                 &lookupList[128],
	MY_POSTAL_CODE_INTL:            &lookupList[129],
	MY_POTA_REF:                    &lookupList[130],
	MY_RIG:                         &lookupList[131],
	MY_RIG_INTL:                    &lookupList[132],
	MY_SIG:                         &lookupList[133],
	MY_SIG_INFO:                    &lookupList[134],
	MY_SIG_INFO_INTL:               &lookupList[135],
	MY_SIG_INTL:                    &lookupList[136],
	MY_SOTA_REF:                    &lookupList[137],
	MY_STATE:                       &lookupList[138],
	MY_STREET:                      &lookupList[139],
	MY_STREET_INTL:                 &lookupList[140],
	MY_USACA_COUNTIES:              &lookupList[141],
	MY_VUCC_GRIDS:                  &lookupList[142],
	MY_WWFF_REF:                    &lookupList[143],
	NAME:                           &lookupList[144],
	NAME_INTL:                      &lookupList[145],
	NOTES:                          &lookupList[146],
	NOTES_INTL:                     &lookupList[147],
	NR_BURSTS:                      &lookupList[148],
	NR_PINGS:                       &lookupList[149],
	OPERATOR:                       &lookupList[150],
	OWNER_CALLSIGN:                 &lookupList[151],
	PFX:                            &lookupList[152],
	POTA_REF:                       &lookupList[153],
	PRECEDENCE:                     &lookupList[154],
	PROGRAMID:                      &lookupList[155],
	PROGRAMVERSION:                 &lookupList[156],
	PROP_MODE:                      &lookupList[157],
	PUBLIC_KEY:                     &lookupList[158],
	QRZCOM_QSO_DOWNLOAD_DATE:       &lookupList[159],
	QRZCOM_QSO_DOWNLOAD_STATUS:     &lookupList[160],
	QRZCOM_QSO_UPLOAD_DATE:         &lookupList[161],
	QRZCOM_QSO_UPLOAD_STATUS:       &lookupList[162],
	QSLMSG:                         &lookupList[163],
	QSLMSG_INTL:                    &lookupList[164],
	QSLMSG_RCVD:                    &lookupList[165],
	QSLRDATE:                       &lookupList[166],
	QSLSDATE:                       &lookupList[167],
	QSL_RCVD:                       &lookupList[168],
	QSL_RCVD_VIA:                   &lookupList[169],
	QSL_SENT:                       &lookupList[170],
	QSL_SENT_VIA:                   &lookupList[171],
	QSL_VIA:                        &lookupList[172],
	QSO_COMPLETE:                   &lookupList[173],
	QSO_DATE:                       &lookupList[174],
	QSO_DATE_OFF:                   &lookupList[175],
	QSO_RANDOM:                     &lookupList[176],
	QTH:                            &lookupList[177],
	QTH_INTL:                       &lookupList[178],
	REGION:                         &lookupList[179],
	RIG:                            &lookupList[180],
	RIG_INTL:                       &lookupList[181],
	RST_RCVD:                       &lookupList[182],
	RST_SENT:                       &lookupList[183],
	RX_PWR:                         &lookupList[184],
	SAT_MODE:                       &lookupList[185],
	SAT_NAME:                       &lookupList[186],
	SFI:                            &lookupList[187],
	SIG:                            &lookupList[188],
	SIG_INFO:                       &lookupList[189],
	SIG_INFO_INTL:                  &lookupList[190],
	SIG_INTL:                       &lookupList[191],
	SILENT_KEY:                     &lookupList[192],
	SKCC:                           &lookupList[193],
	SOTA_REF:                       &lookupList[194],
	SRX:                            &lookupList[195],
	SRX_STRING:                     &lookupList[196],
	STATE:                          &lookupList[197],
	STATION_CALLSIGN:               &lookupList[198],
	STX:                            &lookupList[199],
	STX_STRING:                     &lookupList[200],
	SUBMODE:                        &lookupList[201],
	SWL:                            &lookupList[202],
	TEN_TEN:                        &lookupList[203],
	TIME_OFF:                       &lookupList[204],
	TIME_ON:                        &lookupList[205],
	TX_PWR:                         &lookupList[206],
	UKSMG:                          &lookupList[207],
	USACA_COUNTIES:                 &lookupList[208],
	USERDEF1:                       &lookupList[209],
	USERDEF2:                       &lookupList[210],
	USERDEF3:                       &lookupList[211],
	USERDEF4:                       &lookupList[212],
	USERDEF5:                       &lookupList[213],
	USERDEF6:                       &lookupList[214],
	USERDEF7:                       &lookupList[215],
	USERDEF8:                       &lookupList[216],
	USERDEF9:                       &lookupList[217],
	VE_PROV:                        &lookupList[218],
	VUCC_GRIDS:                     &lookupList[219],
	WEB:                            &lookupList[220],
	WWFF_REF:                       &lookupList[221],
}

// Lookup returns the specification for the provided Field.
// ADIF 3.1.6
func Lookup(f Field) (Spec, bool) {
	spec, ok := lookupMap[f]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all Field specifications that match the provided filter function.
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

// List returns all Field specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns Field specifications.
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
