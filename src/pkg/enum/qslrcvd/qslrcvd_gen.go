// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package qslrcvd

var (
	I QSLRcvd = "I" //
	N QSLRcvd = "N" // an incoming QSL card has not been received the QSO has not been confirmed by the online service
	R QSLRcvd = "R" // the logging station has requested a QSL card the logging station has requested the QSO be uploaded to the online service
	V QSLRcvd = "V" // Deprecated: DXCC award credit granted for the QSL card - instead use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card DXCC credit granted for the LoTW confirmation - instead use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw
	Y QSLRcvd = "Y" // an incoming QSL card has been received the QSO has been confirmed by the online service
)

/*
var QSLRcvdMap = map[qslrcvd.QSLRcvd]spec.QSLRcvdSpec{"I":spec.QSLRcvdSpec{EnumerationName:"QSL_Rcvd", IsImportOnly:false, Comments:"", Id:"I", Meaning:"ignore or invalid", Description:""}, "N":spec.QSLRcvdSpec{EnumerationName:"QSL_Rcvd", IsImportOnly:false, Comments:"", Id:"N", Meaning:"no", Description:"an incoming QSL card has not been received the QSO has not been confirmed by the online service"}, "R":spec.QSLRcvdSpec{EnumerationName:"QSL_Rcvd", IsImportOnly:false, Comments:"", Id:"R", Meaning:"requested", Description:"the logging station has requested a QSL card the logging station has requested the QSO be uploaded to the online service"}, "V":spec.QSLRcvdSpec{EnumerationName:"QSL_Rcvd", IsImportOnly:true, Comments:"", Id:"V", Meaning:"verified", Description:"DXCC award credit granted for the QSL card - instead use <CREDIT_GRANTED:39>DXCC:card,DXCC_BAND:card,DXCC_MODE:card DXCC credit granted for the LoTW confirmation - instead use <CREDIT_GRANTED:39>DXCC:lotw,DXCC_BAND:lotw,DXCC_MODE:lotw"}, "Y":spec.QSLRcvdSpec{EnumerationName:"QSL_Rcvd", IsImportOnly:false, Comments:"", Id:"Y", Meaning:"yes (confirmed)", Description:"an incoming QSL card has been received the QSO has been confirmed by the online service"}}
*/
