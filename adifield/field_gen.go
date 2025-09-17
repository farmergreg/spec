// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package adifield provides code and constants as defined in ADIF 3.1.6 (Released)
package adifield

import "sync"

const (
	ADDRESS                        Field = "address"                        // Record: the contacted station's complete mailing address: full name, street address, city, postal code, and country
	ADDRESS_INTL                   Field = "address_intl"                   // Record: the contacted station's complete mailing address: full name, street address, city, postal code, and country
	ADIF_VER                       Field = "adif_ver"                       // Header: identifies the version of ADIF used in this file in the format X.Y.Z where X is an integer designating the ADIF epoch Y is an integer between 0 and 9 designating the major version Z is an integer between 0 and 9 designating the minor version
	AGE                            Field = "age"                            // Record: the contacted station's operator's age in years in the range 0 to 120 (inclusive)
	ALTITUDE                       Field = "altitude"                       // Record: the height of the contacted station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <ALTITUDE:4>1500 and 10.5 m is <ALTITUDE:4>10.5
	ANT_AZ                         Field = "ant_az"                         // Record: the logging station's antenna azimuth, in degrees with a value between 0 to 360 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 370 is exported as 10). True north is 0 degrees with values increasing in a clockwise direction.
	ANT_EL                         Field = "ant_el"                         // Record: the logging station's antenna elevation, in degrees with a value between -90 to 90 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 100 is exported as 80). The horizon is 0 degrees with values increasing as the angle moves in an upward direction.
	ANT_PATH                       Field = "ant_path"                       // Record: the signal path
	APP_LOTW_2XQSL                 Field = "app_lotw_2xqsl"                 // Record: ARRL LOTW: [Y, N] APP_LOTW_2xQSL ="Y" Indicates this confirmation is considered to be "two-way" in the QSO's Mode. It is a confirmation that counts towards the mode-specific awards & endorsement in the WAS program. Only present if QSL_RCVD="Y".
	APP_LOTW_CQZ_INFERRED          Field = "app_lotw_cqz_inferred"          // Record: ARRL LOTW
	APP_LOTW_CQZ_INVALID           Field = "app_lotw_cqz_invalid"           // Record: ARRL LOTW
	APP_LOTW_CQZ_USERINVALID       Field = "app_lotw_cqz_userinvalid"       // Record: ARRL LOTW: Present only if the QSLing station included an invalid CQ zone value in its station location uploaded to LOTW
	APP_LOTW_CREDIT_GRANTED        Field = "app_lotw_credit_granted"        // Record: ARRL LOTW: Award credits granted based upon this QSL. Awards per LOTW award enumeration.
	APP_LOTW_CREDIT_SUBMITTED      Field = "app_lotw_credit_submitted"      // Record: ARRL LOTW: Award credits for which this QSL has been submitted on an award application to sponsor. Awards per LOTW award enumeration.
	APP_LOTW_DXCC_ENTITY_STATUS    Field = "app_lotw_dxcc_entity_status"    // Record: ARRL LOTW: [Current, Deleted] Indicates whether the ARRL DXCC entity is "Current" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or "Deleted"
	APP_LOTW_EOF                   Field = "app_lotw_eof"                   // Record: ARRL LOTW: End of File marker. No other fields expected in this record.
	APP_LOTW_GRIDSQUARE_INVALID    Field = "app_lotw_gridsquare_invalid"    // Record: ARRL LOTW
	APP_LOTW_ITUZ_INFERRED         Field = "app_lotw_ituz_inferred"         // Record: ARRL LOTW
	APP_LOTW_ITUZ_INVALID          Field = "app_lotw_ituz_invalid"          // Record: ARRL LOTW
	APP_LOTW_ITUZ_USERINVALID      Field = "app_lotw_ituz_userinvalid"      // Record: ARRL LOTW: Present only if the QSLing station included an invalid ITU zone value in its station location uploaded to LOTW
	APP_LOTW_LASTQSL               Field = "app_lotw_lastqsl"               // Header: ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl="yes". Date and time at which the most recent QSL record in this report was received. This will be the QSL date/time of the first QSO record in the file. This value only applies to the records in this file.
	APP_LOTW_LASTQSORX             Field = "app_lotw_lastqsorx"             // Header: ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl="no". Date and time at which the most recent QSO record in this report was received. This will be the QSO received (uploaded) date/time of the first QSO record in the file. This value only applies to the records in this file.
	APP_LOTW_MODE                  Field = "app_lotw_mode"                  // Record: ARRL LOTW: Field is present when the mode recorded in Logbook cannot be unambiguously represented as an ADIF-compliant value
	APP_LOTW_MODEGROUP             Field = "app_lotw_modegroup"             // Record: ARRL LOTW: [CW, PHONE, DATA] The mode group indicates whether the QSO's mode counts towards the CW, Phone or Digital awards respectively in the DXCC program. A mode group of CW indicates the QSO's mode counts towards the CW endorsement on WAS awards. A mode group of PHONE indicates the QSO's mode counts towards the WAS Phone award. A mode group of DATA indicates the QSO's mode counts towards the CQ WPX Digital award.
	APP_LOTW_MY_DXCC_ENTITY_STATUS Field = "app_lotw_my_dxcc_entity_status" // Record: ARRL LOTW: [Current, Deleted] Indicates whether the Logging station's ARRL DXCC entity is "Current" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or "Deleted"
	APP_LOTW_NPSUNIT               Field = "app_lotw_npsunit"               // Record: ARRL LOTW: Present only if the QSLing station included a valid NPOTA Park code or a comma-delimited list of valid NPOTA Park codes
	APP_LOTW_NUMREC                Field = "app_lotw_numrec"                // Header: ARRL LOTW: Number of QSO records in this download
	APP_LOTW_OWNCALL               Field = "app_lotw_owncall"               // Record: ARRL LOTW: "own" call sign of the station making the contact. Present only if qso_withown="yes". DEPRECATED -- In the near future, this field will no longer be present in the report output. Programs that import LOTW report files should not rely on the presence of this field. Use STATION_CALLSIGN instead.
	APP_LOTW_QSLMODE               Field = "app_lotw_qslmode"               // Record: ARRL LOTW: The Mode that your QSO partner indicated on their side of this QSO. Only present if QSL_RCVD="Y" and APP_LOTW_2xQSL="N".
	APP_LOTW_QSO_TIMESTAMP         Field = "app_lotw_qso_timestamp"         // Record: ARRL LOTW: Combined QSO Date & Time in ISO-8601 format
	APP_LOTW_RXQSL                 Field = "app_lotw_rxqsl"                 // Record: ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which the confirmation match (QSL) was made/updated at LOTW. Present only if QSL_RCVD="Y"
	APP_LOTW_RXQSO                 Field = "app_lotw_rxqso"                 // Record: ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which QSO record was inserted/updated at LOTW
	APP_QRZLOG_LOGID               Field = "app_qrzlog_logid"               // Record: QRZ LOG: References the internal ID of a log record in the QRZ system.
	APP_QRZLOG_QSLDATE             Field = "app_qrzlog_qsldate"             // Record: QRZ LOG: QSL Date (e.g. 20210126) This field contains the first date on which we saw a QRZ accepted confirmation for this contact. This could have been a confirmation as a result of matching contacts in the QRZ system or a credit imported from Logbook of The World.
	APP_QRZLOG_STATUS              Field = "app_qrzlog_status"              // Record: QRZ LOG: Status [C: Confirmed, A: Reserved for future use, N: Not Confirmed, 2: Confirmation Requested, S: Confirmation Requested, Request Seen, R: Confirmation Requested, Request Rejected]
	APP_SKCCLOGGER_KEYTYPE         Field = "app_skcclogger_keytype"         // Record: Straight Key Century Club (SKCC): [BUG, SK, SS] CW KEY TYPE (SS=Side Swiper / Cootie, SK=Straight, B=Bug)
	ARRL_SECT                      Field = "arrl_sect"                      // Record: the contacted station's ARRL section
	AWARD_GRANTED                  Field = "award_granted"                  // Record: the list of awards granted by a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's "Award Account" between an award sponsor and the user. For example, in response to a request "send me a list of the awards granted to AA6YQ", this might be received: <CALL:5>AA6YQ <AWARD_GRANTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m
	AWARD_SUBMITTED                Field = "award_submitted"                // Record: the list of awards submitted to a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's "Award Account" between an award sponsor and the user. For example, AA6YQ might submit a request for awards by sending the following: <CALL:5>AA6YQ <AWARD_SUBMITTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m
	A_INDEX                        Field = "a_index"                        // Record: the geomagnetic A index at the time of the QSO in the range 0 to 400 (inclusive)
	BAND                           Field = "band"                           // Record: QSO Band
	BAND_RX                        Field = "band_rx"                        // Record: in a split frequency QSO, the logging station's receiving band
	CALL                           Field = "call"                           // Record: the contacted station's callsign
	CHECK                          Field = "check"                          // Record: contest check (e.g. for ARRL Sweepstakes)
	CLASS                          Field = "class"                          // Record: contest class (e.g. for ARRL Field Day)
	CLUBLOG_QSO_UPLOAD_DATE        Field = "clublog_qso_upload_date"        // Record: the date the QSO was last uploaded to the Club Log online service
	CLUBLOG_QSO_UPLOAD_STATUS      Field = "clublog_qso_upload_status"      // Record: the upload status of the QSO on the Club Log online service
	CNTY                           Field = "cnty"                           // Record: the contacted station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format
	CNTY_ALT                       Field = "cnty_alt"                       // Record: a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the contacted station See the Data Type for details.
	COMMENT                        Field = "comment"                        // Record: comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG field recommended use: information of interest to the contacted station's operator
	COMMENT_INTL                   Field = "comment_intl"                   // Record: comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG_INTL field recommended use: information of interest to the contacted station's operator
	CONT                           Field = "cont"                           // Record: the contacted station's Continent
	CONTACTED_OP                   Field = "contacted_op"                   // Record: the callsign of the individual operating the contacted station
	CONTEST_ID                     Field = "contest_id"                     // Record: QSO Contest Identifier use enumeration values for interoperability
	COUNTRY                        Field = "country"                        // Record: the contacted station's DXCC entity name
	COUNTRY_INTL                   Field = "country_intl"                   // Record: the contacted station's DXCC entity name
	CQZ                            Field = "cqz"                            // Record: the contacted station's CQ Zone in the range 1 to 40 (inclusive)
	CREATED_TIMESTAMP              Field = "created_timestamp"              // Header: identifies the UTC date and time that the file was created in the format of 15 characters YYYYMMDD HHMMSS where YYYYMMDD is a Date data type HHMMSS is a 6 character Time data type
	CREDIT_GRANTED                 Field = "credit_granted"                 // Record: the list of credits granted to this QSO Use of data type AwardList and enumeration Award are import-only
	CREDIT_SUBMITTED               Field = "credit_submitted"               // Record: the list of credits sought for this QSO Use of data type AwardList and enumeration Award are import-only
	DARC_DOK                       Field = "darc_dok"                       // Record: the contacted station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <DARC_DOK:3>A01
	DCL_QSLRDATE                   Field = "dcl_qslrdate"                   // Record: date QSL received from DCL (only valid if DCL_QSL_RCVD is Y, I, or V)(V import-only)
	DCL_QSLSDATE                   Field = "dcl_qslsdate"                   // Record: date QSL sent to DCL (only valid if DCL_QSL_SENT is Y, Q, or I)
	DCL_QSL_RCVD                   Field = "dcl_qsl_rcvd"                   // Record: DCL QSL received status Default Value: N
	DCL_QSL_SENT                   Field = "dcl_qsl_sent"                   // Record: DCL QSL sent status Default Value: N
	DISTANCE                       Field = "distance"                       // Record: the distance between the logging station and the contacted station in kilometers via the specified signal path with a value greater than or equal to 0
	DXCC                           Field = "dxcc"                           // Record: the contacted station's DXCC Entity Code <DXCC:1>0 means that the contacted station is known not to be within a DXCC entity.
	EMAIL                          Field = "email"                          // Record: the contacted station's email address
	EQSL_AG                        Field = "eqsl_ag"                        // Record: indicates whether the QSO is known to be "Authenticity Guaranteed" by eQSL
	EQSL_QSLRDATE                  Field = "eqsl_qslrdate"                  // Record: date QSL received from eQSL.cc (only valid if EQSL_QSL_RCVD is Y, I, or V)(V import-only)
	EQSL_QSLSDATE                  Field = "eqsl_qslsdate"                  // Record: date QSL sent to eQSL.cc (only valid if EQSL_QSL_SENT is Y, Q, or I)
	EQSL_QSL_RCVD                  Field = "eqsl_qsl_rcvd"                  // Record: eQSL.cc QSL received status instead of V (import-only) use <CREDIT_GRANTED:42>CQWAZ:eqsl,CQWAZ_BAND:eqsl,CQWAZ_MODE:eqsl Default Value: N
	EQSL_QSL_SENT                  Field = "eqsl_qsl_sent"                  // Record: eQSL.cc QSL sent status Default Value: N
	EQ_CALL                        Field = "eq_call"                        // Record: the contacted station's owner's callsign
	FISTS                          Field = "fists"                          // Record: the contacted station's FISTS CW Club member number with a value greater than 0.
	FISTS_CC                       Field = "fists_cc"                       // Record: the contacted station's FISTS CW Club Century Certificate (CC) number with a value greater than 0.
	FORCE_INIT                     Field = "force_init"                     // Record: new EME "initial"
	FREQ                           Field = "freq"                           // Record: QSO frequency in Megahertz
	FREQ_RX                        Field = "freq_rx"                        // Record: in a split frequency QSO, the logging station's receiving frequency in Megahertz
	GRIDSQUARE                     Field = "gridsquare"                     // Record: the contacted station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in GRIDSQUARE and the additional 2 or 4 characters in the GRIDSQUARE_EXT field
	GRIDSQUARE_EXT                 Field = "gridsquare_ext"                 // Record: for a contacted station's 10-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9 and 10. For a contacted station's 12-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <GRIDSQUARE:8>FN01MH42 <GRIDSQUARE_EXT:2>BQ
	GUEST_OP                       Field = "guest_op"                       // Deprecated: Record: import-only: use OPERATOR instead
	HAMLOGEU_QSO_UPLOAD_DATE       Field = "hamlogeu_qso_upload_date"       // Record: the date the QSO was last uploaded to the HAMLOG.EU online service
	HAMLOGEU_QSO_UPLOAD_STATUS     Field = "hamlogeu_qso_upload_status"     // Record: the upload status of the QSO on the HAMLOG.EU online service
	HAMQTH_QSO_UPLOAD_DATE         Field = "hamqth_qso_upload_date"         // Record: the date the QSO was last uploaded to the HamQTH.com online service
	HAMQTH_QSO_UPLOAD_STATUS       Field = "hamqth_qso_upload_status"       // Record: the upload status of the QSO on the HamQTH.com online service
	HRDLOG_QSO_UPLOAD_DATE         Field = "hrdlog_qso_upload_date"         // Record: the date the QSO was last uploaded to the HRDLog.net online service
	HRDLOG_QSO_UPLOAD_STATUS       Field = "hrdlog_qso_upload_status"       // Record: the upload status of the QSO on the HRDLog.net online service
	IOTA                           Field = "iota"                           // Record: the contacted station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]
	IOTA_ISLAND_ID                 Field = "iota_island_id"                 // Record: the contacted station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]
	ITUZ                           Field = "ituz"                           // Record: the contacted station's ITU zone in the range 1 to 90 (inclusive)
	K_INDEX                        Field = "k_index"                        // Record: the geomagnetic K index at the time of the QSO in the range 0 to 9 (inclusive)
	LAT                            Field = "lat"                            // Record: the contacted station's latitude
	LON                            Field = "lon"                            // Record: the contacted station's longitude
	LOTW_QSLRDATE                  Field = "lotw_qslrdate"                  // Record: date QSL received from ARRL Logbook of the World (only valid if LOTW_QSL_RCVD is Y, I, or V)(V import-only)
	LOTW_QSLSDATE                  Field = "lotw_qslsdate"                  // Record: date QSL sent to ARRL Logbook of the World (only valid if LOTW_QSL_SENT is Y, Q, or I)
	LOTW_QSL_RCVD                  Field = "lotw_qsl_rcvd"                  // Record: ARRL Logbook of the World QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw Default Value: N
	LOTW_QSL_SENT                  Field = "lotw_qsl_sent"                  // Record: ARRL Logbook of the World QSL sent status Default Value: N
	MAX_BURSTS                     Field = "max_bursts"                     // Record: maximum length of meteor scatter bursts heard by the logging station, in seconds with a value greater than or equal to 0
	MODE                           Field = "mode"                           // Record: QSO Mode
	MORSE_KEY_INFO                 Field = "morse_key_info"                 // Record: details of the contacted station's Morse key (e.g. make, model, etc). Example: <MORSE_KEY_INFO:16>Begali Sculpture
	MORSE_KEY_TYPE                 Field = "morse_key_type"                 // Record: the contacted station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP
	MS_SHOWER                      Field = "ms_shower"                      // Record: For Meteor Scatter QSOs, the name of the meteor shower in progress
	MY_ALTITUDE                    Field = "my_altitude"                    // Record: the height of the logging station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <MY_ALTITUDE:4>1500 and 10.5 m is <MY_ALTITUDE:4>10.5
	MY_ANTENNA                     Field = "my_antenna"                     // Record: the logging station's antenna
	MY_ANTENNA_INTL                Field = "my_antenna_intl"                // Record: the logging station's antenna
	MY_ARRL_SECT                   Field = "my_arrl_sect"                   // Record: the logging station's ARRL section
	MY_CITY                        Field = "my_city"                        // Record: the logging station's city
	MY_CITY_INTL                   Field = "my_city_intl"                   // Record: the logging station's city
	MY_CNTY                        Field = "my_cnty"                        // Record: the logging station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format
	MY_CNTY_ALT                    Field = "my_cnty_alt"                    // Record: a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the logging station See the Data Type for details.
	MY_COUNTRY                     Field = "my_country"                     // Record: the logging station's DXCC entity name
	MY_COUNTRY_INTL                Field = "my_country_intl"                // Record: the logging station's DXCC entity name
	MY_CQ_ZONE                     Field = "my_cq_zone"                     // Record: the logging station's CQ Zone in the range 1 to 40 (inclusive)
	MY_DARC_DOK                    Field = "my_darc_dok"                    // Record: the logging station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <MY_DARC_DOK:3>A01
	MY_DXCC                        Field = "my_dxcc"                        // Record: the logging station's DXCC Entity Code <MY_DXCC:1>0 means that the logging station is known not to be within a DXCC entity.
	MY_FISTS                       Field = "my_fists"                       // Record: the logging station's FISTS CW Club member number with a value greater than 0.
	MY_GRIDSQUARE                  Field = "my_gridsquare"                  // Record: the logging station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in MY_GRIDSQUARE and the additional 2 or 4 characters in the MY_GRIDSQUARE_EXT field
	MY_GRIDSQUARE_EXT              Field = "my_gridsquare_ext"              // Record: for a logging station's 10-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9 and 10. For a logging station's 12-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <MY_GRIDSQUARE:8>FN01MH42 <MY_GRIDSQUARE_EXT:2>BQ
	MY_IOTA                        Field = "my_iota"                        // Record: the logging station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]
	MY_IOTA_ISLAND_ID              Field = "my_iota_island_id"              // Record: the logging station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]
	MY_ITU_ZONE                    Field = "my_itu_zone"                    // Record: the logging station's ITU zone in the range 1 to 90 (inclusive)
	MY_LAT                         Field = "my_lat"                         // Record: the logging station's latitude
	MY_LON                         Field = "my_lon"                         // Record: the logging station's longitude
	MY_MORSE_KEY_INFO              Field = "my_morse_key_info"              // Record: details of the logging station's Morse key (e.g. make, model, etc). Example: <MY_MORSE_KEY_INFO:16>Begali Sculpture
	MY_MORSE_KEY_TYPE              Field = "my_morse_key_type"              // Record: the logging station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP
	MY_NAME                        Field = "my_name"                        // Record: the logging operator's name
	MY_NAME_INTL                   Field = "my_name_intl"                   // Record: the logging operator's name
	MY_POSTAL_CODE                 Field = "my_postal_code"                 // Record: the logging station's postal code
	MY_POSTAL_CODE_INTL            Field = "my_postal_code_intl"            // Record: the logging station's postal code
	MY_POTA_REF                    Field = "my_pota_ref"                    // Record: a comma-delimited list of one or more of the logging station's POTA (Parks on the Air) reference(s). Examples: <MY_POTA_REF:6>K-0059 <MY_POTA_REF:7>K-10000 <MY_POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY
	MY_RIG                         Field = "my_rig"                         // Record: description of the logging station's equipment
	MY_RIG_INTL                    Field = "my_rig_intl"                    // Record: description of the logging station's equipment
	MY_SIG                         Field = "my_sig"                         // Record: special interest activity or event
	MY_SIG_INFO                    Field = "my_sig_info"                    // Record: special interest activity or event information
	MY_SIG_INFO_INTL               Field = "my_sig_info_intl"               // Record: special interest activity or event information
	MY_SIG_INTL                    Field = "my_sig_intl"                    // Record: special interest activity or event
	MY_SOTA_REF                    Field = "my_sota_ref"                    // Record: the logging station's International SOTA Reference.
	MY_STATE                       Field = "my_state"                       // Record: the code for the logging station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)
	MY_STREET                      Field = "my_street"                      // Record: the logging station's street
	MY_STREET_INTL                 Field = "my_street_intl"                 // Record: the logging station's street
	MY_USACA_COUNTIES              Field = "my_usaca_counties"              // Record: two US counties in the case where the logging station is located on a border between two counties, representing counties that the contacted station may claim for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire
	MY_VUCC_GRIDS                  Field = "my_vucc_grids"                  // Record: two or four adjacent Maidenhead grid locators, each four or six characters long, representing the logging station's grid squares that the contacted station may claim for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07
	MY_WWFF_REF                    Field = "my_wwff_ref"                    // Record: the logging station's WWFF (World Wildlife Flora & Fauna) reference
	NAME                           Field = "name"                           // Record: the contacted station's operator's name
	NAME_INTL                      Field = "name_intl"                      // Record: the contacted station's operator's name
	NOTES                          Field = "notes"                          // Record: QSO notes recommended use: information of interest to the logging station's operator
	NOTES_INTL                     Field = "notes_intl"                     // Record: QSO notes recommended use: information of interest to the logging station's operator
	NR_BURSTS                      Field = "nr_bursts"                      // Record: the number of meteor scatter bursts heard by the logging station with a value greater than or equal to 0
	NR_PINGS                       Field = "nr_pings"                       // Record: the number of meteor scatter pings heard by the logging station with a value greater than or equal to 0
	OPERATOR                       Field = "operator"                       // Record: the logging operator's callsign if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign
	OWNER_CALLSIGN                 Field = "owner_callsign"                 // Record: the callsign of the owner of the station used to log the contact (the callsign of the OPERATOR's host) if OWNER_CALLSIGN is absent, STATION_CALLSIGN shall be treated as both the logging station's callsign and the callsign of the owner of the station
	PFX                            Field = "pfx"                            // Record: the contacted station's WPX prefix
	POTA_REF                       Field = "pota_ref"                       // Record: a comma-delimited list of one or more of the contacted station's POTA (Parks on the Air) reference(s). Examples: <POTA_REF:6>K-5033 <POTA_REF:13>VE-5082@CA-AB <POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY
	PRECEDENCE                     Field = "precedence"                     // Record: contest precedence (e.g. for ARRL Sweepstakes)
	PROGRAMID                      Field = "programid"                      // Header: identifies the name of the logger, converter, or utility that created or processed this ADIF file To help avoid name clashes, the ADIF PROGRAMID Register provides a voluntary list of PROGRAMID values.
	PROGRAMVERSION                 Field = "programversion"                 // Header: identifies the version of the logger, converter, or utility that created or processed this ADIF file
	PROP_MODE                      Field = "prop_mode"                      // Record: QSO propagation mode
	PUBLIC_KEY                     Field = "public_key"                     // Record: public encryption key
	QRZCOM_QSO_DOWNLOAD_DATE       Field = "qrzcom_qso_download_date"       // Record: date QSO downloaded from QRZ.COM logbook
	QRZCOM_QSO_DOWNLOAD_STATUS     Field = "qrzcom_qso_download_status"     // Record: QRZ.COM logbook QSO download status
	QRZCOM_QSO_UPLOAD_DATE         Field = "qrzcom_qso_upload_date"         // Record: the date the QSO was last uploaded to the QRZ.COM online service
	QRZCOM_QSO_UPLOAD_STATUS       Field = "qrzcom_qso_upload_status"       // Record: the upload status of the QSO on the QRZ.COM online service
	QSLMSG                         Field = "qslmsg"                         // Record: a message for the contacted station's operator to be incorporated in a paper or electronic QSL
	QSLMSG_INTL                    Field = "qslmsg_intl"                    // Record: a message for the contacted station's operator to be incorporated in a paper or electronic QSL
	QSLMSG_RCVD                    Field = "qslmsg_rcvd"                    // Record: a message addressed to the logging station's operator incorporated in a paper or electronic QSL
	QSLRDATE                       Field = "qslrdate"                       // Record: QSL received date (only valid if QSL_RCVD is Y, I, or V)(V import-only)
	QSLSDATE                       Field = "qslsdate"                       // Record: QSL sent date (only valid if QSL_SENT is Y, Q, or I)
	QSL_RCVD                       Field = "qsl_rcvd"                       // Record: QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card Default Value: N
	QSL_RCVD_VIA                   Field = "qsl_rcvd_via"                   // Record: if QSL_RCVD is set to 'Y' or 'V', the means by which the QSL was received by the logging station; otherwise, the means by which the logging station requested or intends to request that the QSL be conveyed. (Note: 'V' is import-only) use of M (manager) is import-only
	QSL_SENT                       Field = "qsl_sent"                       // Record: QSL sent status Default Value: N
	QSL_SENT_VIA                   Field = "qsl_sent_via"                   // Record: if QSL_SENT is set to 'Y', the means by which the QSL was sent by the logging station; otherwise, the means by which the logging station intends to convey the QSL use of M (manager) is import-only
	QSL_VIA                        Field = "qsl_via"                        // Record: the contacted station's QSL route
	QSO_COMPLETE                   Field = "qso_complete"                   // Record: indicates whether the QSO was complete from the perspective of the logging station Y - yes N - no NIL - not heard ? - uncertain
	QSO_DATE                       Field = "qso_date"                       // Record: date on which the QSO started
	QSO_DATE_OFF                   Field = "qso_date_off"                   // Record: date on which the QSO ended
	QSO_RANDOM                     Field = "qso_random"                     // Record: indicates whether the QSO was random or scheduled
	QTH                            Field = "qth"                            // Record: the contacted station's city
	QTH_INTL                       Field = "qth_intl"                       // Record: the contacted station's city
	REGION                         Field = "region"                         // Record: the contacted station's WAE or CQ entity contained within a DXCC entity. the value None indicates that the WAE or CQ entity is the DXCC entity in the DXCC field. nothing can be inferred from the absence of the REGION field
	RIG                            Field = "rig"                            // Record: description of the contacted station's equipment
	RIG_INTL                       Field = "rig_intl"                       // Record: description of the contacted station's equipment
	RST_RCVD                       Field = "rst_rcvd"                       // Record: signal report from the contacted station
	RST_SENT                       Field = "rst_sent"                       // Record: signal report sent to the contacted station
	RX_PWR                         Field = "rx_pwr"                         // Record: the contacted station's transmitter power in Watts with a value greater than or equal to 0
	SAT_MODE                       Field = "sat_mode"                       // Record: satellite mode - a code representing the satellite's uplink band and downlink band
	SAT_NAME                       Field = "sat_name"                       // Record: name of satellite
	SFI                            Field = "sfi"                            // Record: the solar flux at the time of the QSO in the range 0 to 300 (inclusive).
	SIG                            Field = "sig"                            // Record: the name of the contacted station's special activity or interest group
	SIG_INFO                       Field = "sig_info"                       // Record: information associated with the contacted station's activity or interest group
	SIG_INFO_INTL                  Field = "sig_info_intl"                  // Record: information associated with the contacted station's activity or interest group
	SIG_INTL                       Field = "sig_intl"                       // Record: the name of the contacted station's special activity or interest group
	SILENT_KEY                     Field = "silent_key"                     // Record: 'Y' indicates that the contacted station's operator is now a Silent Key.
	SKCC                           Field = "skcc"                           // Record: the contacted station's Straight Key Century Club (SKCC) member information
	SOTA_REF                       Field = "sota_ref"                       // Record: the contacted station's International SOTA Reference.
	SRX                            Field = "srx"                            // Record: contest QSO received serial number with a value greater than or equal to 0
	SRX_STRING                     Field = "srx_string"                     // Record: contest QSO received information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail
	STATE                          Field = "state"                          // Record: the code for the contacted station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)
	STATION_CALLSIGN               Field = "station_callsign"               // Record: the logging station's callsign (the callsign used over the air) if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign
	STX                            Field = "stx"                            // Record: contest QSO transmitted serial number with a value greater than or equal to 0
	STX_STRING                     Field = "stx_string"                     // Record: contest QSO transmitted information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail
	SUBMODE                        Field = "submode"                        // Record: QSO Submode use enumeration values for interoperability
	SWL                            Field = "swl"                            // Record: indicates that the QSO information pertains to an SWL report
	TEN_TEN                        Field = "ten_ten"                        // Record: Ten-Ten number with a value greater than 0
	TIME_OFF                       Field = "time_off"                       // Record: HHMM or HHMMSS in UTC in the absence of <QSO_DATE_OFF>, the QSO duration is less than 24 hours. For example, the following is a QSO starting at 14 July 2020 23:55 and finishing at 15 July 2020 01:00: <QSO_DATE:8>20200714 <TIME_ON:4>2355 <TIME_OFF:4>0100
	TIME_ON                        Field = "time_on"                        // Record: HHMM or HHMMSS in UTC
	TX_PWR                         Field = "tx_pwr"                         // Record: the logging station's power in Watts with a value greater than or equal to 0
	UKSMG                          Field = "uksmg"                          // Record: the contacted station's UKSMG member number with a value greater than 0
	USACA_COUNTIES                 Field = "usaca_counties"                 // Record: two US counties in the case where the contacted station is located on a border between two counties, representing counties credited to the QSO for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire
	USERDEF1                       Field = "userdef1"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF2                       Field = "userdef2"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF3                       Field = "userdef3"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF4                       Field = "userdef4"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF5                       Field = "userdef5"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF6                       Field = "userdef6"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF7                       Field = "userdef7"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF8                       Field = "userdef8"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	USERDEF9                       Field = "userdef9"                       // Header: specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character
	VE_PROV                        Field = "ve_prov"                        // Deprecated: Record: import-only: use STATE instead
	VUCC_GRIDS                     Field = "vucc_grids"                     // Record: two or four adjacent Maidenhead grid locators, each four or six characters long, representing the contacted station's grid squares credited to the QSO for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07
	WEB                            Field = "web"                            // Record: the contacted station's URL
	WWFF_REF                       Field = "wwff_ref"                       // Record: the contacted station's WWFF (World Wildlife Flora & Fauna) reference
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Field specifications.
var lookupList = []Spec{
	{Key: "address", DataType: "multilinestring", Description: "the contacted station's complete mailing address: full name, street address, city, postal code, and country", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "address_intl", DataType: "intlmultilinestring", Description: "the contacted station's complete mailing address: full name, street address, city, postal code, and country", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "adif_ver", DataType: "string", Description: "identifies the version of ADIF used in this file in the format X.Y.Z where X is an integer designating the ADIF epoch Y is an integer between 0 and 9 designating the major version Z is an integer between 0 and 9 designating the minor version", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "age", DataType: "number", Description: "the contacted station's operator's age in years in the range 0 to 120 (inclusive)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 120, IsImportOnly: false, Comments: ""},
	{Key: "altitude", DataType: "number", Description: "the height of the contacted station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <ALTITUDE:4>1500 and 10.5 m is <ALTITUDE:4>10.5", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "ant_az", DataType: "number", Description: "the logging station's antenna azimuth, in degrees with a value between 0 to 360 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 370 is exported as 10). True north is 0 degrees with values increasing in a clockwise direction.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 360, IsImportOnly: false, Comments: ""},
	{Key: "ant_el", DataType: "number", Description: "the logging station's antenna elevation, in degrees with a value between -90 to 90 (inclusive). Values outside this range are import-only and must be normalized for export (e.g. 100 is exported as 80). The horizon is 0 degrees with values increasing as the angle moves in an upward direction.", IsHeaderField: false, MinimumValue: -90, MaximumValue: 90, IsImportOnly: false, Comments: ""},
	{Key: "ant_path", DataType: "enumeration", Description: "the signal path", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_2xqsl", DataType: "enumeration", Description: "ARRL LOTW: [Y, N] APP_LOTW_2xQSL =\"Y\" Indicates this confirmation is considered to be \"two-way\" in the QSO's Mode. It is a confirmation that counts towards the mode-specific awards & endorsement in the WAS program. Only present if QSL_RCVD=\"Y\".", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_cqz_inferred", DataType: "string", Description: "ARRL LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_cqz_invalid", DataType: "string", Description: "ARRL LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_cqz_userinvalid", DataType: "string", Description: "ARRL LOTW: Present only if the QSLing station included an invalid CQ zone value in its station location uploaded to LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_credit_granted", DataType: "string", Description: "ARRL LOTW: Award credits granted based upon this QSL. Awards per LOTW award enumeration.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_credit_submitted", DataType: "string", Description: "ARRL LOTW: Award credits for which this QSL has been submitted on an award application to sponsor. Awards per LOTW award enumeration.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_dxcc_entity_status", DataType: "enumeration", Description: "ARRL LOTW: [Current, Deleted] Indicates whether the ARRL DXCC entity is \"Current\" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or \"Deleted\"", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_eof", DataType: "string", Description: "ARRL LOTW: End of File marker. No other fields expected in this record.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_gridsquare_invalid", DataType: "string", Description: "ARRL LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_ituz_inferred", DataType: "string", Description: "ARRL LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_ituz_invalid", DataType: "string", Description: "ARRL LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_ituz_userinvalid", DataType: "string", Description: "ARRL LOTW: Present only if the QSLing station included an invalid ITU zone value in its station location uploaded to LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_lastqsl", DataType: "string", Description: "ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl=\"yes\". Date and time at which the most recent QSL record in this report was received. This will be the QSL date/time of the first QSO record in the file. This value only applies to the records in this file.", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_lastqsorx", DataType: "string", Description: "ARRL LOTW: date/time (YYYY-MM-DD HH:MM:SS) Present only if qso_qsl=\"no\". Date and time at which the most recent QSO record in this report was received. This will be the QSO received (uploaded) date/time of the first QSO record in the file. This value only applies to the records in this file.", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_mode", DataType: "enumeration", Description: "ARRL LOTW: Field is present when the mode recorded in Logbook cannot be unambiguously represented as an ADIF-compliant value", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_modegroup", DataType: "enumeration", Description: "ARRL LOTW: [CW, PHONE, DATA] The mode group indicates whether the QSO's mode counts towards the CW, Phone or Digital awards respectively in the DXCC program. A mode group of CW indicates the QSO's mode counts towards the CW endorsement on WAS awards. A mode group of PHONE indicates the QSO's mode counts towards the WAS Phone award. A mode group of DATA indicates the QSO's mode counts towards the CQ WPX Digital award.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_my_dxcc_entity_status", DataType: "enumeration", Description: "ARRL LOTW: [Current, Deleted] Indicates whether the Logging station's ARRL DXCC entity is \"Current\" -- counts toward Honor Roll, 5 Band DXCC and DXCC Challenege awrads -- or \"Deleted\"", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_npsunit", DataType: "string", Description: "ARRL LOTW: Present only if the QSLing station included a valid NPOTA Park code or a comma-delimited list of valid NPOTA Park codes", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_numrec", DataType: "positiveinteger", Description: "ARRL LOTW: Number of QSO records in this download", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_owncall", DataType: "string", Description: "ARRL LOTW: \"own\" call sign of the station making the contact. Present only if qso_withown=\"yes\". DEPRECATED -- In the near future, this field will no longer be present in the report output. Programs that import LOTW report files should not rely on the presence of this field. Use STATION_CALLSIGN instead.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_qslmode", DataType: "enumeration", Description: "ARRL LOTW: The Mode that your QSO partner indicated on their side of this QSO. Only present if QSL_RCVD=\"Y\" and APP_LOTW_2xQSL=\"N\".", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_qso_timestamp", DataType: "string", Description: "ARRL LOTW: Combined QSO Date & Time in ISO-8601 format", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_rxqsl", DataType: "string", Description: "ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which the confirmation match (QSL) was made/updated at LOTW. Present only if QSL_RCVD=\"Y\"", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_lotw_rxqso", DataType: "string", Description: "ARRL LOTW: Timestamp (YYYY-MM-DD HH:MM:SS) at which QSO record was inserted/updated at LOTW", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_qrzlog_logid", DataType: "positiveinteger", Description: "QRZ LOG: References the internal ID of a log record in the QRZ system.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_qrzlog_qsldate", DataType: "date", Description: "QRZ LOG: QSL Date (e.g. 20210126) This field contains the first date on which we saw a QRZ accepted confirmation for this contact. This could have been a confirmation as a result of matching contacts in the QRZ system or a credit imported from Logbook of The World.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_qrzlog_status", DataType: "enumeration", Description: "QRZ LOG: Status [C: Confirmed, A: Reserved for future use, N: Not Confirmed, 2: Confirmation Requested, S: Confirmation Requested, Request Seen, R: Confirmation Requested, Request Rejected]", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "app_skcclogger_keytype", DataType: "enumeration", Description: "Straight Key Century Club (SKCC): [BUG, SK, SS] CW KEY TYPE (SS=Side Swiper / Cootie, SK=Straight, B=Bug)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "arrl_sect", DataType: "enumeration", Description: "the contacted station's ARRL section", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "award_granted", DataType: "sponsoredawardlist", Description: "the list of awards granted by a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's \"Award Account\" between an award sponsor and the user. For example, in response to a request \"send me a list of the awards granted to AA6YQ\", this might be received: <CALL:5>AA6YQ <AWARD_GRANTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "award_submitted", DataType: "sponsoredawardlist", Description: "the list of awards submitted to a sponsor. note that this field might not be used in a QSO record. It might be used to convey information about a user's \"Award Account\" between an award sponsor and the user. For example, AA6YQ might submit a request for awards by sending the following: <CALL:5>AA6YQ <AWARD_SUBMITTED:64>ADIF_CENTURY_BASIC,ADIF_CENTURY_SILVER,ADIF_SPECTRUM_100-160m", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "a_index", DataType: "number", Description: "the geomagnetic A index at the time of the QSO in the range 0 to 400 (inclusive)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 400, IsImportOnly: false, Comments: ""},
	{Key: "band", DataType: "enumeration", Description: "QSO Band", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "band_rx", DataType: "enumeration", Description: "in a split frequency QSO, the logging station's receiving band", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "call", DataType: "string", Description: "the contacted station's callsign", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "check", DataType: "string", Description: "contest check (e.g. for ARRL Sweepstakes)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "class", DataType: "string", Description: "contest class (e.g. for ARRL Field Day)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "clublog_qso_upload_date", DataType: "date", Description: "the date the QSO was last uploaded to the Club Log online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "clublog_qso_upload_status", DataType: "enumeration", Description: "the upload status of the QSO on the Club Log online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "cnty", DataType: "enumeration", Description: "the contacted station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "cnty_alt", DataType: "secondaryadministrativesubdivisionlistalt", Description: "a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the contacted station See the Data Type for details.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "comment", DataType: "string", Description: "comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG field recommended use: information of interest to the contacted station's operator", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "comment_intl", DataType: "intlstring", Description: "comment field for QSO for a message to be incorporated in a paper or electronic QSL for the contacted station's operator, use the QSLMSG_INTL field recommended use: information of interest to the contacted station's operator", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "cont", DataType: "enumeration", Description: "the contacted station's Continent", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "contacted_op", DataType: "string", Description: "the callsign of the individual operating the contacted station", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "contest_id", DataType: "string", Description: "QSO Contest Identifier use enumeration values for interoperability", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "country", DataType: "string", Description: "the contacted station's DXCC entity name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "country_intl", DataType: "intlstring", Description: "the contacted station's DXCC entity name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "cqz", DataType: "positiveinteger", Description: "the contacted station's CQ Zone in the range 1 to 40 (inclusive)", IsHeaderField: false, MinimumValue: 1, MaximumValue: 40, IsImportOnly: false, Comments: ""},
	{Key: "created_timestamp", DataType: "string", Description: "identifies the UTC date and time that the file was created in the format of 15 characters YYYYMMDD HHMMSS where YYYYMMDD is a Date data type HHMMSS is a 6 character Time data type", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "credit_granted", DataType: "creditlist,awardlist", Description: "the list of credits granted to this QSO Use of data type AwardList and enumeration Award are import-only", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "credit_submitted", DataType: "creditlist,awardlist", Description: "the list of credits sought for this QSO Use of data type AwardList and enumeration Award are import-only", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "darc_dok", DataType: "enumeration", Description: "the contacted station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <DARC_DOK:3>A01", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "dcl_qslrdate", DataType: "date", Description: "date QSL received from DCL (only valid if DCL_QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "dcl_qslsdate", DataType: "date", Description: "date QSL sent to DCL (only valid if DCL_QSL_SENT is Y, Q, or I)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "dcl_qsl_rcvd", DataType: "enumeration", Description: "DCL QSL received status Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "dcl_qsl_sent", DataType: "enumeration", Description: "DCL QSL sent status Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "distance", DataType: "number", Description: "the distance between the logging station and the contacted station in kilometers via the specified signal path with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "dxcc", DataType: "enumeration", Description: "the contacted station's DXCC Entity Code <DXCC:1>0 means that the contacted station is known not to be within a DXCC entity.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "email", DataType: "string", Description: "the contacted station's email address", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "eqsl_ag", DataType: "enumeration", Description: "indicates whether the QSO is known to be \"Authenticity Guaranteed\" by eQSL", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "eqsl_qslrdate", DataType: "date", Description: "date QSL received from eQSL.cc (only valid if EQSL_QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "eqsl_qslsdate", DataType: "date", Description: "date QSL sent to eQSL.cc (only valid if EQSL_QSL_SENT is Y, Q, or I)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "eqsl_qsl_rcvd", DataType: "enumeration", Description: "eQSL.cc QSL received status instead of V (import-only) use <CREDIT_GRANTED:42>CQWAZ:eqsl,CQWAZ_BAND:eqsl,CQWAZ_MODE:eqsl Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "eqsl_qsl_sent", DataType: "enumeration", Description: "eQSL.cc QSL sent status Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "eq_call", DataType: "string", Description: "the contacted station's owner's callsign", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "fists", DataType: "positiveinteger", Description: "the contacted station's FISTS CW Club member number with a value greater than 0.", IsHeaderField: false, MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "fists_cc", DataType: "positiveinteger", Description: "the contacted station's FISTS CW Club Century Certificate (CC) number with a value greater than 0.", IsHeaderField: false, MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "force_init", DataType: "boolean", Description: "new EME \"initial\"", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "freq", DataType: "number", Description: "QSO frequency in Megahertz", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "freq_rx", DataType: "number", Description: "in a split frequency QSO, the logging station's receiving frequency in Megahertz", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "gridsquare", DataType: "gridsquare", Description: "the contacted station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in GRIDSQUARE and the additional 2 or 4 characters in the GRIDSQUARE_EXT field", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "gridsquare_ext", DataType: "gridsquareext", Description: "for a contacted station's 10-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9 and 10. For a contacted station's 12-character Maidenhead locator, supplements the GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <GRIDSQUARE:8>FN01MH42 <GRIDSQUARE_EXT:2>BQ", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "guest_op", DataType: "string", Description: "import-only: use OPERATOR instead", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: true, Comments: ""},
	{Key: "hamlogeu_qso_upload_date", DataType: "date", Description: "the date the QSO was last uploaded to the HAMLOG.EU online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "hamlogeu_qso_upload_status", DataType: "enumeration", Description: "the upload status of the QSO on the HAMLOG.EU online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "hamqth_qso_upload_date", DataType: "date", Description: "the date the QSO was last uploaded to the HamQTH.com online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "hamqth_qso_upload_status", DataType: "enumeration", Description: "the upload status of the QSO on the HamQTH.com online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "hrdlog_qso_upload_date", DataType: "date", Description: "the date the QSO was last uploaded to the HRDLog.net online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "hrdlog_qso_upload_status", DataType: "enumeration", Description: "the upload status of the QSO on the HRDLog.net online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "iota", DataType: "iotarefno", Description: "the contacted station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "iota_island_id", DataType: "positiveinteger", Description: "the contacted station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]", IsHeaderField: false, MinimumValue: 1, MaximumValue: 99999999, IsImportOnly: false, Comments: ""},
	{Key: "ituz", DataType: "positiveinteger", Description: "the contacted station's ITU zone in the range 1 to 90 (inclusive)", IsHeaderField: false, MinimumValue: 1, MaximumValue: 90, IsImportOnly: false, Comments: ""},
	{Key: "k_index", DataType: "integer", Description: "the geomagnetic K index at the time of the QSO in the range 0 to 9 (inclusive)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 9, IsImportOnly: false, Comments: ""},
	{Key: "lat", DataType: "location", Description: "the contacted station's latitude", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "lon", DataType: "location", Description: "the contacted station's longitude", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "lotw_qslrdate", DataType: "date", Description: "date QSL received from ARRL Logbook of the World (only valid if LOTW_QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "lotw_qslsdate", DataType: "date", Description: "date QSL sent to ARRL Logbook of the World (only valid if LOTW_QSL_SENT is Y, Q, or I)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "lotw_qsl_rcvd", DataType: "enumeration", Description: "ARRL Logbook of the World QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "lotw_qsl_sent", DataType: "enumeration", Description: "ARRL Logbook of the World QSL sent status Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "max_bursts", DataType: "number", Description: "maximum length of meteor scatter bursts heard by the logging station, in seconds with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "mode", DataType: "enumeration", Description: "QSO Mode", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "morse_key_info", DataType: "string", Description: "details of the contacted station's Morse key (e.g. make, model, etc). Example: <MORSE_KEY_INFO:16>Begali Sculpture", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "morse_key_type", DataType: "enumeration", Description: "the contacted station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "ms_shower", DataType: "string", Description: "For Meteor Scatter QSOs, the name of the meteor shower in progress", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_altitude", DataType: "number", Description: "the height of the logging station in meters relative to Mean Sea Level (MSL). For example 1.5 km is <MY_ALTITUDE:4>1500 and 10.5 m is <MY_ALTITUDE:4>10.5", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_antenna", DataType: "string", Description: "the logging station's antenna", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_antenna_intl", DataType: "intlstring", Description: "the logging station's antenna", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_arrl_sect", DataType: "enumeration", Description: "the logging station's ARRL section", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_city", DataType: "string", Description: "the logging station's city", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_city_intl", DataType: "intlstring", Description: "the logging station's city", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_cnty", DataType: "enumeration", Description: "the logging station's Secondary Administrative Subdivision (e.g. US county, JA Gun), in the specified format", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_cnty_alt", DataType: "secondaryadministrativesubdivisionlistalt", Description: "a semicolon (;) delimited, unordered list of Secondary Administrative Subdivision Alt codes for the logging station See the Data Type for details.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_country", DataType: "string", Description: "the logging station's DXCC entity name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_country_intl", DataType: "intlstring", Description: "the logging station's DXCC entity name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_cq_zone", DataType: "positiveinteger", Description: "the logging station's CQ Zone in the range 1 to 40 (inclusive)", IsHeaderField: false, MinimumValue: 1, MaximumValue: 40, IsImportOnly: false, Comments: ""},
	{Key: "my_darc_dok", DataType: "enumeration", Description: "the logging station's DARC DOK (District Location Code) A DOK comprises letters and numbers, e.g. <MY_DARC_DOK:3>A01", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_dxcc", DataType: "enumeration", Description: "the logging station's DXCC Entity Code <MY_DXCC:1>0 means that the logging station is known not to be within a DXCC entity.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_fists", DataType: "positiveinteger", Description: "the logging station's FISTS CW Club member number with a value greater than 0.", IsHeaderField: false, MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_gridsquare", DataType: "gridsquare", Description: "the logging station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square For 10 or 12 character locators, store the first 8 characters in MY_GRIDSQUARE and the additional 2 or 4 characters in the MY_GRIDSQUARE_EXT field", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_gridsquare_ext", DataType: "gridsquareext", Description: "for a logging station's 10-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9 and 10. For a logging station's 12-character Maidenhead locator, supplements the MY_GRIDSQUARE field by containing characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0 to 9. On export, the field length must be 2 or 4. On import, if the field length is greater than 4, the additional characters must be ignored. Example of exporting the 10-character locator FN01MH42BQ: <MY_GRIDSQUARE:8>FN01MH42 <MY_GRIDSQUARE_EXT:2>BQ", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_iota", DataType: "iotarefno", Description: "the logging station's IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_iota_island_id", DataType: "positiveinteger", Description: "the logging station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999 [leading zeroes optional]", IsHeaderField: false, MinimumValue: 1, MaximumValue: 99999999, IsImportOnly: false, Comments: ""},
	{Key: "my_itu_zone", DataType: "positiveinteger", Description: "the logging station's ITU zone in the range 1 to 90 (inclusive)", IsHeaderField: false, MinimumValue: 1, MaximumValue: 90, IsImportOnly: false, Comments: ""},
	{Key: "my_lat", DataType: "location", Description: "the logging station's latitude", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_lon", DataType: "location", Description: "the logging station's longitude", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_morse_key_info", DataType: "string", Description: "details of the logging station's Morse key (e.g. make, model, etc). Example: <MY_MORSE_KEY_INFO:16>Begali Sculpture", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_morse_key_type", DataType: "enumeration", Description: "the logging station's Morse key type (e.g. straight key, bug, etc). Example for a dual-lever paddle: <MORSE_KEY_TYPE:2>DP", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_name", DataType: "string", Description: "the logging operator's name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_name_intl", DataType: "intlstring", Description: "the logging operator's name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_postal_code", DataType: "string", Description: "the logging station's postal code", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_postal_code_intl", DataType: "intlstring", Description: "the logging station's postal code", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_pota_ref", DataType: "potareflist", Description: "a comma-delimited list of one or more of the logging station's POTA (Parks on the Air) reference(s). Examples: <MY_POTA_REF:6>K-0059 <MY_POTA_REF:7>K-10000 <MY_POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_rig", DataType: "string", Description: "description of the logging station's equipment", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_rig_intl", DataType: "intlstring", Description: "description of the logging station's equipment", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_sig", DataType: "string", Description: "special interest activity or event", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_sig_info", DataType: "string", Description: "special interest activity or event information", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_sig_info_intl", DataType: "intlstring", Description: "special interest activity or event information", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_sig_intl", DataType: "intlstring", Description: "special interest activity or event", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_sota_ref", DataType: "sotaref", Description: "the logging station's International SOTA Reference.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_state", DataType: "enumeration", Description: "the code for the logging station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_street", DataType: "string", Description: "the logging station's street", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_street_intl", DataType: "intlstring", Description: "the logging station's street", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_usaca_counties", DataType: "secondarysubdivisionlist", Description: "two US counties in the case where the logging station is located on a border between two counties, representing counties that the contacted station may claim for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_vucc_grids", DataType: "gridsquarelist", Description: "two or four adjacent Maidenhead grid locators, each four or six characters long, representing the logging station's grid squares that the contacted station may claim for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "my_wwff_ref", DataType: "wwffref", Description: "the logging station's WWFF (World Wildlife Flora & Fauna) reference", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "name", DataType: "string", Description: "the contacted station's operator's name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "name_intl", DataType: "intlstring", Description: "the contacted station's operator's name", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "notes", DataType: "multilinestring", Description: "QSO notes recommended use: information of interest to the logging station's operator", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "notes_intl", DataType: "intlmultilinestring", Description: "QSO notes recommended use: information of interest to the logging station's operator", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "nr_bursts", DataType: "integer", Description: "the number of meteor scatter bursts heard by the logging station with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "nr_pings", DataType: "integer", Description: "the number of meteor scatter pings heard by the logging station with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "operator", DataType: "string", Description: "the logging operator's callsign if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "owner_callsign", DataType: "string", Description: "the callsign of the owner of the station used to log the contact (the callsign of the OPERATOR's host) if OWNER_CALLSIGN is absent, STATION_CALLSIGN shall be treated as both the logging station's callsign and the callsign of the owner of the station", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "pfx", DataType: "string", Description: "the contacted station's WPX prefix", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "pota_ref", DataType: "potareflist", Description: "a comma-delimited list of one or more of the contacted station's POTA (Parks on the Air) reference(s). Examples: <POTA_REF:6>K-5033 <POTA_REF:13>VE-5082@CA-AB <POTA_REF:40>K-0817,K-4566,K-4576,K-4573,K-4578@US-WY", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "precedence", DataType: "string", Description: "contest precedence (e.g. for ARRL Sweepstakes)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "programid", DataType: "string", Description: "identifies the name of the logger, converter, or utility that created or processed this ADIF file To help avoid name clashes, the ADIF PROGRAMID Register provides a voluntary list of PROGRAMID values.", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "programversion", DataType: "string", Description: "identifies the version of the logger, converter, or utility that created or processed this ADIF file", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "prop_mode", DataType: "enumeration", Description: "QSO propagation mode", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "public_key", DataType: "string", Description: "public encryption key", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qrzcom_qso_download_date", DataType: "date", Description: "date QSO downloaded from QRZ.COM logbook", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qrzcom_qso_download_status", DataType: "enumeration", Description: "QRZ.COM logbook QSO download status", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qrzcom_qso_upload_date", DataType: "date", Description: "the date the QSO was last uploaded to the QRZ.COM online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qrzcom_qso_upload_status", DataType: "enumeration", Description: "the upload status of the QSO on the QRZ.COM online service", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qslmsg", DataType: "multilinestring", Description: "a message for the contacted station's operator to be incorporated in a paper or electronic QSL", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qslmsg_intl", DataType: "intlmultilinestring", Description: "a message for the contacted station's operator to be incorporated in a paper or electronic QSL", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qslmsg_rcvd", DataType: "multilinestring", Description: "a message addressed to the logging station's operator incorporated in a paper or electronic QSL", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qslrdate", DataType: "date", Description: "QSL received date (only valid if QSL_RCVD is Y, I, or V)(V import-only)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qslsdate", DataType: "date", Description: "QSL sent date (only valid if QSL_SENT is Y, Q, or I)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qsl_rcvd", DataType: "enumeration", Description: "QSL received status instead of V (import-only) use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qsl_rcvd_via", DataType: "enumeration", Description: "if QSL_RCVD is set to 'Y' or 'V', the means by which the QSL was received by the logging station; otherwise, the means by which the logging station requested or intends to request that the QSL be conveyed. (Note: 'V' is import-only) use of M (manager) is import-only", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qsl_sent", DataType: "enumeration", Description: "QSL sent status Default Value: N", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qsl_sent_via", DataType: "enumeration", Description: "if QSL_SENT is set to 'Y', the means by which the QSL was sent by the logging station; otherwise, the means by which the logging station intends to convey the QSL use of M (manager) is import-only", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qsl_via", DataType: "string", Description: "the contacted station's QSL route", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qso_complete", DataType: "enumeration", Description: "indicates whether the QSO was complete from the perspective of the logging station Y - yes N - no NIL - not heard ? - uncertain", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qso_date", DataType: "date", Description: "date on which the QSO started", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qso_date_off", DataType: "date", Description: "date on which the QSO ended", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qso_random", DataType: "boolean", Description: "indicates whether the QSO was random or scheduled", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qth", DataType: "string", Description: "the contacted station's city", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "qth_intl", DataType: "intlstring", Description: "the contacted station's city", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "region", DataType: "enumeration", Description: "the contacted station's WAE or CQ entity contained within a DXCC entity. the value None indicates that the WAE or CQ entity is the DXCC entity in the DXCC field. nothing can be inferred from the absence of the REGION field", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "rig", DataType: "multilinestring", Description: "description of the contacted station's equipment", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "rig_intl", DataType: "intlmultilinestring", Description: "description of the contacted station's equipment", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "rst_rcvd", DataType: "string", Description: "signal report from the contacted station", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "rst_sent", DataType: "string", Description: "signal report sent to the contacted station", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "rx_pwr", DataType: "number", Description: "the contacted station's transmitter power in Watts with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "sat_mode", DataType: "string", Description: "satellite mode - a code representing the satellite's uplink band and downlink band", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "sat_name", DataType: "string", Description: "name of satellite", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "sfi", DataType: "integer", Description: "the solar flux at the time of the QSO in the range 0 to 300 (inclusive).", IsHeaderField: false, MinimumValue: 0, MaximumValue: 300, IsImportOnly: false, Comments: ""},
	{Key: "sig", DataType: "string", Description: "the name of the contacted station's special activity or interest group", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "sig_info", DataType: "string", Description: "information associated with the contacted station's activity or interest group", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "sig_info_intl", DataType: "intlstring", Description: "information associated with the contacted station's activity or interest group", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "sig_intl", DataType: "intlstring", Description: "the name of the contacted station's special activity or interest group", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "silent_key", DataType: "boolean", Description: "'Y' indicates that the contacted station's operator is now a Silent Key.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "skcc", DataType: "string", Description: "the contacted station's Straight Key Century Club (SKCC) member information", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "sota_ref", DataType: "sotaref", Description: "the contacted station's International SOTA Reference.", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "srx", DataType: "integer", Description: "contest QSO received serial number with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "srx_string", DataType: "string", Description: "contest QSO received information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "state", DataType: "enumeration", Description: "the code for the contacted station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "station_callsign", DataType: "string", Description: "the logging station's callsign (the callsign used over the air) if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "stx", DataType: "integer", Description: "contest QSO transmitted serial number with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "stx_string", DataType: "string", Description: "contest QSO transmitted information use Cabrillo format to convey contest information for which ADIF fields are not specified in the event of a conflict between information in a dedicated contest field and this field, information in the dedicated contest field shall prevail", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "submode", DataType: "string", Description: "QSO Submode use enumeration values for interoperability", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "swl", DataType: "boolean", Description: "indicates that the QSO information pertains to an SWL report", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "ten_ten", DataType: "positiveinteger", Description: "Ten-Ten number with a value greater than 0", IsHeaderField: false, MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "time_off", DataType: "time", Description: "HHMM or HHMMSS in UTC in the absence of <QSO_DATE_OFF>, the QSO duration is less than 24 hours. For example, the following is a QSO starting at 14 July 2020 23:55 and finishing at 15 July 2020 01:00: <QSO_DATE:8>20200714 <TIME_ON:4>2355 <TIME_OFF:4>0100", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "time_on", DataType: "time", Description: "HHMM or HHMMSS in UTC", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "tx_pwr", DataType: "number", Description: "the logging station's power in Watts with a value greater than or equal to 0", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "uksmg", DataType: "positiveinteger", Description: "the contacted station's UKSMG member number with a value greater than 0", IsHeaderField: false, MinimumValue: 1, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "usaca_counties", DataType: "secondarysubdivisionlist", Description: "two US counties in the case where the contacted station is located on a border between two counties, representing counties credited to the QSO for the CQ Magazine USA-CA award program. E.g. MA,Franklin:MA,Hampshire", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "userdef1", DataType: "string", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "userdef2", DataType: "string", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "userdef3", DataType: "string", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "userdef4", DataType: "string", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "userdef5", DataType: "string", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "userdef6", DataType: "string", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "userdef7", DataType: "string", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "userdef8", DataType: "string", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "userdef9", DataType: "string", Description: "specifies the name and optional enumeration or range of the nth user-defined field, where n is a positive integer The name of a user-defined field may not be an ADIF Field name contain a comma a colon an open-angle-bracket or close-angle-bracket character an open-curly-bracket or close-curly-bracket character begin or end with a space character", IsHeaderField: true, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "ve_prov", DataType: "string", Description: "import-only: use STATE instead", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: true, Comments: ""},
	{Key: "vucc_grids", DataType: "gridsquarelist", Description: "two or four adjacent Maidenhead grid locators, each four or six characters long, representing the contacted station's grid squares credited to the QSO for the ARRL VUCC award program. E.g. EM98,FM08,EM97,FM07", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "web", DataType: "string", Description: "the contacted station's URL", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
	{Key: "wwff_ref", DataType: "wwffref", Description: "the contacted station's WWFF (World Wildlife Flora & Fauna) reference", IsHeaderField: false, MinimumValue: 0, MaximumValue: 0, IsImportOnly: false, Comments: ""},
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
