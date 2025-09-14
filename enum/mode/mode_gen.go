// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package mode provides code and constants as defined in ADIF 3.1.6 (Proposed)
package mode

import (
	"sync"

	"github.com/hamradiolog-net/spec/v6/enum/submode"
)

const (
	AM           Mode = "AM"           // AM                     =
	AMTORFEC     Mode = "AMTORFEC"     // Deprecated: AMTORFEC   =
	ARDOP        Mode = "ARDOP"        // ARDOP                  =
	ASCI         Mode = "ASCI"         // Deprecated: ASCI       =
	ATV          Mode = "ATV"          // ATV                    =
	C4FM         Mode = "C4FM"         // Deprecated: C4FM       =
	CHIP         Mode = "CHIP"         // CHIP                   = CHIP64, CHIP128
	CHIP128      Mode = "CHIP128"      // Deprecated: CHIP128    =
	CHIP64       Mode = "CHIP64"       // Deprecated: CHIP64     =
	CLO          Mode = "CLO"          // CLO                    =
	CONTESTI     Mode = "CONTESTI"     // CONTESTI               =
	CW           Mode = "CW"           // CW                     = PCW
	DIGITALVOICE Mode = "DIGITALVOICE" // DIGITALVOICE           = C4FM, DMR, DSTAR, FREEDV, M17
	DOMINO       Mode = "DOMINO"       // DOMINO                 = DOM-M, DOM4, DOM5, DOM8, DOM11, DOM16, DOM22, DOM44, DOM88, DOMINOEX, DOMINOF
	DOMINOF      Mode = "DOMINOF"      // Deprecated: DOMINOF    =
	DSTAR        Mode = "DSTAR"        // Deprecated: DSTAR      =
	DYNAMIC      Mode = "DYNAMIC"      // DYNAMIC                = VARA HF, VARA SATELLITE, VARA FM 1200, VARA FM 9600
	FAX          Mode = "FAX"          // FAX                    =
	FM           Mode = "FM"           // FM                     =
	FMHELL       Mode = "FMHELL"       // Deprecated: FMHELL     =
	FSK          Mode = "FSK"          // FSK                    = SCAMP_FAST, SCAMP_SLOW, SCAMP_VSLOW
	FSK31        Mode = "FSK31"        // Deprecated: FSK31      =
	FSK441       Mode = "FSK441"       // FSK441                 =
	FT8          Mode = "FT8"          // FT8                    =
	GTOR         Mode = "GTOR"         // Deprecated: GTOR       =
	HELL         Mode = "HELL"         // HELL                   = FMHELL, FSKH105, FSKH245, FSKHELL, HELL80, HELLX5, HELLX9, HFSK, PSKHELL, SLOWHELL
	HELL80       Mode = "HELL80"       // Deprecated: HELL80     =
	HFSK         Mode = "HFSK"         // Deprecated: HFSK       =
	ISCAT        Mode = "ISCAT"        // ISCAT                  = ISCAT-A, ISCAT-B
	JT4          Mode = "JT4"          // JT4                    = JT4A, JT4B, JT4C, JT4D, JT4E, JT4F, JT4G
	JT44         Mode = "JT44"         // JT44                   =
	JT4A         Mode = "JT4A"         // Deprecated: JT4A       =
	JT4B         Mode = "JT4B"         // Deprecated: JT4B       =
	JT4C         Mode = "JT4C"         // Deprecated: JT4C       =
	JT4D         Mode = "JT4D"         // Deprecated: JT4D       =
	JT4E         Mode = "JT4E"         // Deprecated: JT4E       =
	JT4F         Mode = "JT4F"         // Deprecated: JT4F       =
	JT4G         Mode = "JT4G"         // Deprecated: JT4G       =
	JT65         Mode = "JT65"         // JT65                   = JT65A, JT65B, JT65B2, JT65C, JT65C2
	JT65A        Mode = "JT65A"        // Deprecated: JT65A      =
	JT65B        Mode = "JT65B"        // Deprecated: JT65B      =
	JT65C        Mode = "JT65C"        // Deprecated: JT65C      =
	JT6M         Mode = "JT6M"         // JT6M                   =
	JT9          Mode = "JT9"          // JT9                    = JT9-1, JT9-2, JT9-5, JT9-10, JT9-30, JT9A, JT9B, JT9C, JT9D, JT9E, JT9E FAST, JT9F, JT9F FAST, JT9G, JT9G FAST, JT9H, JT9H FAST
	MFSK         Mode = "MFSK"         // MFSK                   = FSQCALL, FST4, FST4W, FT4, JS8, JTMS, MFSK4, MFSK8, MFSK11, MFSK16, MFSK22, MFSK31, MFSK32, MFSK64, MFSK64L, MFSK128 MFSK128L, Q65
	MFSK16       Mode = "MFSK16"       // Deprecated: MFSK16     =
	MFSK8        Mode = "MFSK8"        // Deprecated: MFSK8      =
	MSK144       Mode = "MSK144"       // MSK144                 =
	MT63         Mode = "MT63"         // MT63                   =
	MTONE        Mode = "MTONE"        // MTONE                  = SCAMP_OO, SCAMP_OO_SLW
	OLIVIA       Mode = "OLIVIA"       // OLIVIA                 = OLIVIA 4/125, OLIVIA 4/250, OLIVIA 8/250, OLIVIA 8/500, OLIVIA 16/500, OLIVIA 16/1000, OLIVIA 32/1000
	OPERA        Mode = "OPERA"        // OPERA                  = OPERA-BEACON, OPERA-QSO
	PAC          Mode = "PAC"          // PAC                    = PAC2, PAC3, PAC4
	PAC2         Mode = "PAC2"         // Deprecated: PAC2       =
	PAC3         Mode = "PAC3"         // Deprecated: PAC3       =
	PAX          Mode = "PAX"          // PAX                    = PAX2
	PAX2         Mode = "PAX2"         // Deprecated: PAX2       =
	PCW          Mode = "PCW"          // Deprecated: PCW        =
	PKT          Mode = "PKT"          // PKT                    =
	PSK          Mode = "PSK"          // PSK                    = 8PSK125, 8PSK125F, 8PSK125FL, 8PSK250, 8PSK250F, 8PSK250FL, 8PSK500, 8PSK500F, 8PSK1000, 8PSK1000F, 8PSK1200F, FSK31, PSK10, PSK31, PSK63, PSK63F, PSK63RC4, PSK63RC5, PSK63RC10, PSK63RC20, PSK63RC32, PSK125, PSK125C12, PSK125R, PSK125RC10, PSK125RC12, PSK125RC16, PSK125RC4, PSK125RC5, PSK250, PSK250C6, PSK250R, PSK250RC2, PSK250RC3, PSK250RC5, PSK250RC6, PSK250RC7, PSK500, PSK500C2, PSK500C4, PSK500R, PSK500RC2, PSK500RC3, PSK500RC4, PSK800C2, PSK800RC2, PSK1000, PSK1000C2, PSK1000R, PSK1000RC2, PSKAM10, PSKAM31, PSKAM50, PSKFEC31, QPSK31, QPSK63, QPSK125, QPSK250, QPSK500, SIM31
	PSK10        Mode = "PSK10"        // Deprecated: PSK10      =
	PSK125       Mode = "PSK125"       // Deprecated: PSK125     =
	PSK2K        Mode = "PSK2K"        // PSK2K                  =
	PSK31        Mode = "PSK31"        // Deprecated: PSK31      =
	PSK63        Mode = "PSK63"        // Deprecated: PSK63      =
	PSK63F       Mode = "PSK63F"       // Deprecated: PSK63F     =
	PSKAM10      Mode = "PSKAM10"      // Deprecated: PSKAM10    =
	PSKAM31      Mode = "PSKAM31"      // Deprecated: PSKAM31    =
	PSKAM50      Mode = "PSKAM50"      // Deprecated: PSKAM50    =
	PSKFEC31     Mode = "PSKFEC31"     // Deprecated: PSKFEC31   =
	PSKHELL      Mode = "PSKHELL"      // Deprecated: PSKHELL    =
	Q15          Mode = "Q15"          // Q15                    =
	QPSK125      Mode = "QPSK125"      // Deprecated: QPSK125    =
	QPSK31       Mode = "QPSK31"       // Deprecated: QPSK31     =
	QPSK63       Mode = "QPSK63"       // Deprecated: QPSK63     =
	QRA64        Mode = "QRA64"        // QRA64                  = QRA64A, QRA64B, QRA64C, QRA64D, QRA64E
	ROS          Mode = "ROS"          // ROS                    = ROS-EME, ROS-HF, ROS-MF
	RTTY         Mode = "RTTY"         // RTTY                   = ASCI
	RTTYM        Mode = "RTTYM"        // RTTYM                  =
	SSB          Mode = "SSB"          // SSB                    = LSB, USB
	SSTV         Mode = "SSTV"         // SSTV                   =
	T10          Mode = "T10"          // T10                    =
	THOR         Mode = "THOR"         // THOR                   = THOR-M, THOR4, THOR5, THOR8, THOR11, THOR16, THOR22, THOR25X4, THOR50X1, THOR50X2, THOR100
	THRB         Mode = "THRB"         // THRB                   = THRBX, THRBX1, THRBX2, THRBX4, THROB1, THROB2, THROB4
	THRBX        Mode = "THRBX"        // Deprecated: THRBX      =
	TOR          Mode = "TOR"          // TOR                    = AMTORFEC, GTOR, NAVTEX, SITORB
	V4           Mode = "V4"           // V4                     =
	VOI          Mode = "VOI"          // VOI                    =
	WINMOR       Mode = "WINMOR"       // WINMOR                 =
	WSPR         Mode = "WSPR"         // WSPR                   =
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Mode specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "AM", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "AMTORFEC", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "ARDOP", Submodes: submode.SubModeList(nil), Description: "Amateur Radio Digital Open Protocol"},
	{IsImportOnly: true, Key: "ASCI", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "ATV", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "C4FM", Submodes: submode.SubModeList(nil), Description: "C4FM 4-level FSK Technology Imported QSOs with <MODE:4>C4FM> must be exported as: <MODE:12>DIGITALVOICE <SUBMODE:4>C4FM"},
	{IsImportOnly: false, Key: "CHIP", Submodes: submode.SubModeList{"CHIP64", "CHIP128"}, Description: ""},
	{IsImportOnly: true, Key: "CHIP128", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "CHIP64", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "CLO", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "CONTESTI", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "CW", Submodes: submode.SubModeList{"PCW"}, Description: ""},
	{IsImportOnly: false, Key: "DIGITALVOICE", Submodes: submode.SubModeList{"C4FM", "DMR", "DSTAR", "FREEDV", "M17"}, Description: ""},
	{IsImportOnly: false, Key: "DOMINO", Submodes: submode.SubModeList{"DOM-M", "DOM4", "DOM5", "DOM8", "DOM11", "DOM16", "DOM22", "DOM44", "DOM88", "DOMINOEX", "DOMINOF"}, Description: ""},
	{IsImportOnly: true, Key: "DOMINOF", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "DSTAR", Submodes: submode.SubModeList(nil), Description: "Digital Smart Technologies for Amateur Radio Imported QSOs with <MODE:5>DSTAR must be exported as: <MODE:12>DIGITALVOICE <SUBMODE:5>DSTAR"},
	{IsImportOnly: false, Key: "DYNAMIC", Submodes: submode.SubModeList{"VARA HF", "VARA SATELLITE", "VARA FM 1200", "VARA FM 9600"}, Description: ""},
	{IsImportOnly: false, Key: "FAX", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "FM", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "FMHELL", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "FSK", Submodes: submode.SubModeList{"SCAMP_FAST", "SCAMP_SLOW", "SCAMP_VSLOW"}, Description: "Frequency shift keying"},
	{IsImportOnly: true, Key: "FSK31", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "FSK441", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "FT8", Submodes: submode.SubModeList(nil), Description: "Franke-Taylor design, 8-FSK modulation"},
	{IsImportOnly: true, Key: "GTOR", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "HELL", Submodes: submode.SubModeList{"FMHELL", "FSKH105", "FSKH245", "FSKHELL", "HELL80", "HELLX5", "HELLX9", "HFSK", "PSKHELL", "SLOWHELL"}, Description: ""},
	{IsImportOnly: true, Key: "HELL80", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "HFSK", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "ISCAT", Submodes: submode.SubModeList{"ISCAT-A", "ISCAT-B"}, Description: ""},
	{IsImportOnly: false, Key: "JT4", Submodes: submode.SubModeList{"JT4A", "JT4B", "JT4C", "JT4D", "JT4E", "JT4F", "JT4G"}, Description: ""},
	{IsImportOnly: false, Key: "JT44", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "JT4A", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "JT4B", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "JT4C", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "JT4D", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "JT4E", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "JT4F", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "JT4G", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "JT65", Submodes: submode.SubModeList{"JT65A", "JT65B", "JT65B2", "JT65C", "JT65C2"}, Description: ""},
	{IsImportOnly: true, Key: "JT65A", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "JT65B", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "JT65C", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "JT6M", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "JT9", Submodes: submode.SubModeList{"JT9-1", "JT9-2", "JT9-5", "JT9-10", "JT9-30", "JT9A", "JT9B", "JT9C", "JT9D", "JT9E", "JT9E FAST", "JT9F", "JT9F FAST", "JT9G", "JT9G FAST", "JT9H", "JT9H FAST"}, Description: ""},
	{IsImportOnly: false, Key: "MFSK", Submodes: submode.SubModeList{"FSQCALL", "FST4", "FST4W", "FT4", "JS8", "JTMS", "MFSK4", "MFSK8", "MFSK11", "MFSK16", "MFSK22", "MFSK31", "MFSK32", "MFSK64", "MFSK64L", "MFSK128 MFSK128L", "Q65"}, Description: ""},
	{IsImportOnly: true, Key: "MFSK16", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "MFSK8", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "MSK144", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "MT63", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "MTONE", Submodes: submode.SubModeList{"SCAMP_OO", "SCAMP_OO_SLW"}, Description: "Single modulated tone"},
	{IsImportOnly: false, Key: "OLIVIA", Submodes: submode.SubModeList{"OLIVIA 4/125", "OLIVIA 4/250", "OLIVIA 8/250", "OLIVIA 8/500", "OLIVIA 16/500", "OLIVIA 16/1000", "OLIVIA 32/1000"}, Description: ""},
	{IsImportOnly: false, Key: "OPERA", Submodes: submode.SubModeList{"OPERA-BEACON", "OPERA-QSO"}, Description: ""},
	{IsImportOnly: false, Key: "PAC", Submodes: submode.SubModeList{"PAC2", "PAC3", "PAC4"}, Description: ""},
	{IsImportOnly: true, Key: "PAC2", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "PAC3", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "PAX", Submodes: submode.SubModeList{"PAX2"}, Description: ""},
	{IsImportOnly: true, Key: "PAX2", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "PCW", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "PKT", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "PSK", Submodes: submode.SubModeList{"8PSK125", "8PSK125F", "8PSK125FL", "8PSK250", "8PSK250F", "8PSK250FL", "8PSK500", "8PSK500F", "8PSK1000", "8PSK1000F", "8PSK1200F", "FSK31", "PSK10", "PSK31", "PSK63", "PSK63F", "PSK63RC4", "PSK63RC5", "PSK63RC10", "PSK63RC20", "PSK63RC32", "PSK125", "PSK125C12", "PSK125R", "PSK125RC10", "PSK125RC12", "PSK125RC16", "PSK125RC4", "PSK125RC5", "PSK250", "PSK250C6", "PSK250R", "PSK250RC2", "PSK250RC3", "PSK250RC5", "PSK250RC6", "PSK250RC7", "PSK500", "PSK500C2", "PSK500C4", "PSK500R", "PSK500RC2", "PSK500RC3", "PSK500RC4", "PSK800C2", "PSK800RC2", "PSK1000", "PSK1000C2", "PSK1000R", "PSK1000RC2", "PSKAM10", "PSKAM31", "PSKAM50", "PSKFEC31", "QPSK31", "QPSK63", "QPSK125", "QPSK250", "QPSK500", "SIM31"}, Description: ""},
	{IsImportOnly: true, Key: "PSK10", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "PSK125", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "PSK2K", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "PSK31", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "PSK63", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "PSK63F", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "PSKAM10", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "PSKAM31", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "PSKAM50", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "PSKFEC31", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "PSKHELL", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "Q15", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "QPSK125", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "QPSK31", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "QPSK63", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "QRA64", Submodes: submode.SubModeList{"QRA64A", "QRA64B", "QRA64C", "QRA64D", "QRA64E"}, Description: ""},
	{IsImportOnly: false, Key: "ROS", Submodes: submode.SubModeList{"ROS-EME", "ROS-HF", "ROS-MF"}, Description: ""},
	{IsImportOnly: false, Key: "RTTY", Submodes: submode.SubModeList{"ASCI"}, Description: ""},
	{IsImportOnly: false, Key: "RTTYM", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "SSB", Submodes: submode.SubModeList{"LSB", "USB"}, Description: ""},
	{IsImportOnly: false, Key: "SSTV", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "T10", Submodes: submode.SubModeList(nil), Description: "Tonal 10 digital mode with focus on sensitivity, band capacity and resistance to the HF Doppler frequency spread"},
	{IsImportOnly: false, Key: "THOR", Submodes: submode.SubModeList{"THOR-M", "THOR4", "THOR5", "THOR8", "THOR11", "THOR16", "THOR22", "THOR25X4", "THOR50X1", "THOR50X2", "THOR100"}, Description: ""},
	{IsImportOnly: false, Key: "THRB", Submodes: submode.SubModeList{"THRBX", "THRBX1", "THRBX2", "THRBX4", "THROB1", "THROB2", "THROB4"}, Description: ""},
	{IsImportOnly: true, Key: "THRBX", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "TOR", Submodes: submode.SubModeList{"AMTORFEC", "GTOR", "NAVTEX", "SITORB"}, Description: ""},
	{IsImportOnly: false, Key: "V4", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "VOI", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "WINMOR", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "WSPR", Submodes: submode.SubModeList(nil), Description: ""},
}

// lookupMap contains all known Mode specifications.
var lookupMap = map[Mode]*Spec{
	AM:           &lookupList[0],
	AMTORFEC:     &lookupList[1],
	ARDOP:        &lookupList[2],
	ASCI:         &lookupList[3],
	ATV:          &lookupList[4],
	C4FM:         &lookupList[5],
	CHIP:         &lookupList[6],
	CHIP128:      &lookupList[7],
	CHIP64:       &lookupList[8],
	CLO:          &lookupList[9],
	CONTESTI:     &lookupList[10],
	CW:           &lookupList[11],
	DIGITALVOICE: &lookupList[12],
	DOMINO:       &lookupList[13],
	DOMINOF:      &lookupList[14],
	DSTAR:        &lookupList[15],
	DYNAMIC:      &lookupList[16],
	FAX:          &lookupList[17],
	FM:           &lookupList[18],
	FMHELL:       &lookupList[19],
	FSK:          &lookupList[20],
	FSK31:        &lookupList[21],
	FSK441:       &lookupList[22],
	FT8:          &lookupList[23],
	GTOR:         &lookupList[24],
	HELL:         &lookupList[25],
	HELL80:       &lookupList[26],
	HFSK:         &lookupList[27],
	ISCAT:        &lookupList[28],
	JT4:          &lookupList[29],
	JT44:         &lookupList[30],
	JT4A:         &lookupList[31],
	JT4B:         &lookupList[32],
	JT4C:         &lookupList[33],
	JT4D:         &lookupList[34],
	JT4E:         &lookupList[35],
	JT4F:         &lookupList[36],
	JT4G:         &lookupList[37],
	JT65:         &lookupList[38],
	JT65A:        &lookupList[39],
	JT65B:        &lookupList[40],
	JT65C:        &lookupList[41],
	JT6M:         &lookupList[42],
	JT9:          &lookupList[43],
	MFSK:         &lookupList[44],
	MFSK16:       &lookupList[45],
	MFSK8:        &lookupList[46],
	MSK144:       &lookupList[47],
	MT63:         &lookupList[48],
	MTONE:        &lookupList[49],
	OLIVIA:       &lookupList[50],
	OPERA:        &lookupList[51],
	PAC:          &lookupList[52],
	PAC2:         &lookupList[53],
	PAC3:         &lookupList[54],
	PAX:          &lookupList[55],
	PAX2:         &lookupList[56],
	PCW:          &lookupList[57],
	PKT:          &lookupList[58],
	PSK:          &lookupList[59],
	PSK10:        &lookupList[60],
	PSK125:       &lookupList[61],
	PSK2K:        &lookupList[62],
	PSK31:        &lookupList[63],
	PSK63:        &lookupList[64],
	PSK63F:       &lookupList[65],
	PSKAM10:      &lookupList[66],
	PSKAM31:      &lookupList[67],
	PSKAM50:      &lookupList[68],
	PSKFEC31:     &lookupList[69],
	PSKHELL:      &lookupList[70],
	Q15:          &lookupList[71],
	QPSK125:      &lookupList[72],
	QPSK31:       &lookupList[73],
	QPSK63:       &lookupList[74],
	QRA64:        &lookupList[75],
	ROS:          &lookupList[76],
	RTTY:         &lookupList[77],
	RTTYM:        &lookupList[78],
	SSB:          &lookupList[79],
	SSTV:         &lookupList[80],
	T10:          &lookupList[81],
	THOR:         &lookupList[82],
	THRB:         &lookupList[83],
	THRBX:        &lookupList[84],
	TOR:          &lookupList[85],
	V4:           &lookupList[86],
	VOI:          &lookupList[87],
	WINMOR:       &lookupList[88],
	WSPR:         &lookupList[89],
}

// Lookup returns the specification for the provided Mode.
// ADIF 3.1.6
func Lookup(m Mode) (Spec, bool) {
	spec, ok := lookupMap[m]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all Mode specifications that match the provided filter function.
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

// List returns all Mode specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns Mode specifications.
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
