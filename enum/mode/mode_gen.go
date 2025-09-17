// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package mode provides code and constants as defined in ADIF 3.1.6 (Released)
package mode

import (
	"sync"

	"github.com/hamradiolog-net/spec/v6/enum/submode"
)

const (
	AM           Mode = "am"           // am                     =
	AMTORFEC     Mode = "amtorfec"     // Deprecated: amtorfec   =
	ARDOP        Mode = "ardop"        // ardop                  =
	ASCI         Mode = "asci"         // Deprecated: asci       =
	ATV          Mode = "atv"          // atv                    =
	C4FM         Mode = "c4fm"         // Deprecated: c4fm       =
	CHIP         Mode = "chip"         // chip                   = CHIP64, CHIP128
	CHIP128      Mode = "chip128"      // Deprecated: chip128    =
	CHIP64       Mode = "chip64"       // Deprecated: chip64     =
	CLO          Mode = "clo"          // clo                    =
	CONTESTI     Mode = "contesti"     // contesti               =
	CW           Mode = "cw"           // cw                     = PCW
	DIGITALVOICE Mode = "digitalvoice" // digitalvoice           = C4FM, DMR, DSTAR, FREEDV, M17
	DOMINO       Mode = "domino"       // domino                 = DOM-M, DOM4, DOM5, DOM8, DOM11, DOM16, DOM22, DOM44, DOM88, DOMINOEX, DOMINOF
	DOMINOF      Mode = "dominof"      // Deprecated: dominof    =
	DSTAR        Mode = "dstar"        // Deprecated: dstar      =
	DYNAMIC      Mode = "dynamic"      // dynamic                = VARA HF, VARA SATELLITE, VARA FM 1200, VARA FM 9600
	FAX          Mode = "fax"          // fax                    =
	FM           Mode = "fm"           // fm                     =
	FMHELL       Mode = "fmhell"       // Deprecated: fmhell     =
	FSK          Mode = "fsk"          // fsk                    = SCAMP_FAST, SCAMP_SLOW, SCAMP_VSLOW
	FSK31        Mode = "fsk31"        // Deprecated: fsk31      =
	FSK441       Mode = "fsk441"       // fsk441                 =
	FT8          Mode = "ft8"          // ft8                    =
	GTOR         Mode = "gtor"         // Deprecated: gtor       =
	HELL         Mode = "hell"         // hell                   = FMHELL, FSKH105, FSKH245, FSKHELL, HELL80, HELLX5, HELLX9, HFSK, PSKHELL, SLOWHELL
	HELL80       Mode = "hell80"       // Deprecated: hell80     =
	HFSK         Mode = "hfsk"         // Deprecated: hfsk       =
	ISCAT        Mode = "iscat"        // iscat                  = ISCAT-A, ISCAT-B
	JT4          Mode = "jt4"          // jt4                    = JT4A, JT4B, JT4C, JT4D, JT4E, JT4F, JT4G
	JT44         Mode = "jt44"         // jt44                   =
	JT4A         Mode = "jt4a"         // Deprecated: jt4a       =
	JT4B         Mode = "jt4b"         // Deprecated: jt4b       =
	JT4C         Mode = "jt4c"         // Deprecated: jt4c       =
	JT4D         Mode = "jt4d"         // Deprecated: jt4d       =
	JT4E         Mode = "jt4e"         // Deprecated: jt4e       =
	JT4F         Mode = "jt4f"         // Deprecated: jt4f       =
	JT4G         Mode = "jt4g"         // Deprecated: jt4g       =
	JT65         Mode = "jt65"         // jt65                   = JT65A, JT65B, JT65B2, JT65C, JT65C2
	JT65A        Mode = "jt65a"        // Deprecated: jt65a      =
	JT65B        Mode = "jt65b"        // Deprecated: jt65b      =
	JT65C        Mode = "jt65c"        // Deprecated: jt65c      =
	JT6M         Mode = "jt6m"         // jt6m                   =
	JT9          Mode = "jt9"          // jt9                    = JT9-1, JT9-2, JT9-5, JT9-10, JT9-30, JT9A, JT9B, JT9C, JT9D, JT9E, JT9E FAST, JT9F, JT9F FAST, JT9G, JT9G FAST, JT9H, JT9H FAST
	MFSK         Mode = "mfsk"         // mfsk                   = FSQCALL, FST4, FST4W, FT4, JS8, JTMS, MFSK4, MFSK8, MFSK11, MFSK16, MFSK22, MFSK31, MFSK32, MFSK64, MFSK64L, MFSK128 MFSK128L, Q65
	MFSK16       Mode = "mfsk16"       // Deprecated: mfsk16     =
	MFSK8        Mode = "mfsk8"        // Deprecated: mfsk8      =
	MSK144       Mode = "msk144"       // msk144                 =
	MT63         Mode = "mt63"         // mt63                   =
	MTONE        Mode = "mtone"        // mtone                  = SCAMP_OO, SCAMP_OO_SLW
	OLIVIA       Mode = "olivia"       // olivia                 = OLIVIA 4/125, OLIVIA 4/250, OLIVIA 8/250, OLIVIA 8/500, OLIVIA 16/500, OLIVIA 16/1000, OLIVIA 32/1000
	OPERA        Mode = "opera"        // opera                  = OPERA-BEACON, OPERA-QSO
	PAC          Mode = "pac"          // pac                    = PAC2, PAC3, PAC4
	PAC2         Mode = "pac2"         // Deprecated: pac2       =
	PAC3         Mode = "pac3"         // Deprecated: pac3       =
	PAX          Mode = "pax"          // pax                    = PAX2
	PAX2         Mode = "pax2"         // Deprecated: pax2       =
	PCW          Mode = "pcw"          // Deprecated: pcw        =
	PKT          Mode = "pkt"          // pkt                    =
	PSK          Mode = "psk"          // psk                    = 8PSK125, 8PSK125F, 8PSK125FL, 8PSK250, 8PSK250F, 8PSK250FL, 8PSK500, 8PSK500F, 8PSK1000, 8PSK1000F, 8PSK1200F, FSK31, PSK10, PSK31, PSK63, PSK63F, PSK63RC4, PSK63RC5, PSK63RC10, PSK63RC20, PSK63RC32, PSK125, PSK125C12, PSK125R, PSK125RC10, PSK125RC12, PSK125RC16, PSK125RC4, PSK125RC5, PSK250, PSK250C6, PSK250R, PSK250RC2, PSK250RC3, PSK250RC5, PSK250RC6, PSK250RC7, PSK500, PSK500C2, PSK500C4, PSK500R, PSK500RC2, PSK500RC3, PSK500RC4, PSK800C2, PSK800RC2, PSK1000, PSK1000C2, PSK1000R, PSK1000RC2, PSKAM10, PSKAM31, PSKAM50, PSKFEC31, QPSK31, QPSK63, QPSK125, QPSK250, QPSK500, SIM31
	PSK10        Mode = "psk10"        // Deprecated: psk10      =
	PSK125       Mode = "psk125"       // Deprecated: psk125     =
	PSK2K        Mode = "psk2k"        // psk2k                  =
	PSK31        Mode = "psk31"        // Deprecated: psk31      =
	PSK63        Mode = "psk63"        // Deprecated: psk63      =
	PSK63F       Mode = "psk63f"       // Deprecated: psk63f     =
	PSKAM10      Mode = "pskam10"      // Deprecated: pskam10    =
	PSKAM31      Mode = "pskam31"      // Deprecated: pskam31    =
	PSKAM50      Mode = "pskam50"      // Deprecated: pskam50    =
	PSKFEC31     Mode = "pskfec31"     // Deprecated: pskfec31   =
	PSKHELL      Mode = "pskhell"      // Deprecated: pskhell    =
	Q15          Mode = "q15"          // q15                    =
	QPSK125      Mode = "qpsk125"      // Deprecated: qpsk125    =
	QPSK31       Mode = "qpsk31"       // Deprecated: qpsk31     =
	QPSK63       Mode = "qpsk63"       // Deprecated: qpsk63     =
	QRA64        Mode = "qra64"        // qra64                  = QRA64A, QRA64B, QRA64C, QRA64D, QRA64E
	ROS          Mode = "ros"          // ros                    = ROS-EME, ROS-HF, ROS-MF
	RTTY         Mode = "rtty"         // rtty                   = ASCI
	RTTYM        Mode = "rttym"        // rttym                  =
	SSB          Mode = "ssb"          // ssb                    = LSB, USB
	SSTV         Mode = "sstv"         // sstv                   =
	T10          Mode = "t10"          // t10                    =
	THOR         Mode = "thor"         // thor                   = THOR-M, THOR4, THOR5, THOR8, THOR11, THOR16, THOR22, THOR25X4, THOR50X1, THOR50X2, THOR100
	THRB         Mode = "thrb"         // thrb                   = THRBX, THRBX1, THRBX2, THRBX4, THROB1, THROB2, THROB4
	THRBX        Mode = "thrbx"        // Deprecated: thrbx      =
	TOR          Mode = "tor"          // tor                    = AMTORFEC, GTOR, NAVTEX, SITORB
	V4           Mode = "v4"           // v4                     =
	VOI          Mode = "voi"          // voi                    =
	WINMOR       Mode = "winmor"       // winmor                 =
	WSPR         Mode = "wspr"         // wspr                   =
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Mode specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "am", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "amtorfec", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "ardop", Submodes: submode.SubModeList(nil), Description: "Amateur Radio Digital Open Protocol"},
	{IsImportOnly: true, Key: "asci", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "atv", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "c4fm", Submodes: submode.SubModeList(nil), Description: "C4FM 4-level FSK Technology Imported QSOs with <MODE:4>C4FM> must be exported as: <MODE:12>DIGITALVOICE <SUBMODE:4>C4FM"},
	{IsImportOnly: false, Key: "chip", Submodes: submode.SubModeList{"CHIP64", "CHIP128"}, Description: ""},
	{IsImportOnly: true, Key: "chip128", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "chip64", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "clo", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "contesti", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "cw", Submodes: submode.SubModeList{"PCW"}, Description: ""},
	{IsImportOnly: false, Key: "digitalvoice", Submodes: submode.SubModeList{"C4FM", "DMR", "DSTAR", "FREEDV", "M17"}, Description: ""},
	{IsImportOnly: false, Key: "domino", Submodes: submode.SubModeList{"DOM-M", "DOM4", "DOM5", "DOM8", "DOM11", "DOM16", "DOM22", "DOM44", "DOM88", "DOMINOEX", "DOMINOF"}, Description: ""},
	{IsImportOnly: true, Key: "dominof", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "dstar", Submodes: submode.SubModeList(nil), Description: "Digital Smart Technologies for Amateur Radio Imported QSOs with <MODE:5>DSTAR must be exported as: <MODE:12>DIGITALVOICE <SUBMODE:5>DSTAR"},
	{IsImportOnly: false, Key: "dynamic", Submodes: submode.SubModeList{"VARA HF", "VARA SATELLITE", "VARA FM 1200", "VARA FM 9600"}, Description: ""},
	{IsImportOnly: false, Key: "fax", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "fm", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "fmhell", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "fsk", Submodes: submode.SubModeList{"SCAMP_FAST", "SCAMP_SLOW", "SCAMP_VSLOW"}, Description: "Frequency shift keying"},
	{IsImportOnly: true, Key: "fsk31", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "fsk441", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "ft8", Submodes: submode.SubModeList(nil), Description: "Franke-Taylor design, 8-FSK modulation"},
	{IsImportOnly: true, Key: "gtor", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "hell", Submodes: submode.SubModeList{"FMHELL", "FSKH105", "FSKH245", "FSKHELL", "HELL80", "HELLX5", "HELLX9", "HFSK", "PSKHELL", "SLOWHELL"}, Description: ""},
	{IsImportOnly: true, Key: "hell80", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "hfsk", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "iscat", Submodes: submode.SubModeList{"ISCAT-A", "ISCAT-B"}, Description: ""},
	{IsImportOnly: false, Key: "jt4", Submodes: submode.SubModeList{"JT4A", "JT4B", "JT4C", "JT4D", "JT4E", "JT4F", "JT4G"}, Description: ""},
	{IsImportOnly: false, Key: "jt44", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4a", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4b", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4c", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4d", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4e", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4f", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4g", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "jt65", Submodes: submode.SubModeList{"JT65A", "JT65B", "JT65B2", "JT65C", "JT65C2"}, Description: ""},
	{IsImportOnly: true, Key: "jt65a", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt65b", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt65c", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "jt6m", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "jt9", Submodes: submode.SubModeList{"JT9-1", "JT9-2", "JT9-5", "JT9-10", "JT9-30", "JT9A", "JT9B", "JT9C", "JT9D", "JT9E", "JT9E FAST", "JT9F", "JT9F FAST", "JT9G", "JT9G FAST", "JT9H", "JT9H FAST"}, Description: ""},
	{IsImportOnly: false, Key: "mfsk", Submodes: submode.SubModeList{"FSQCALL", "FST4", "FST4W", "FT4", "JS8", "JTMS", "MFSK4", "MFSK8", "MFSK11", "MFSK16", "MFSK22", "MFSK31", "MFSK32", "MFSK64", "MFSK64L", "MFSK128 MFSK128L", "Q65"}, Description: ""},
	{IsImportOnly: true, Key: "mfsk16", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "mfsk8", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "msk144", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "mt63", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "mtone", Submodes: submode.SubModeList{"SCAMP_OO", "SCAMP_OO_SLW"}, Description: "Single modulated tone"},
	{IsImportOnly: false, Key: "olivia", Submodes: submode.SubModeList{"OLIVIA 4/125", "OLIVIA 4/250", "OLIVIA 8/250", "OLIVIA 8/500", "OLIVIA 16/500", "OLIVIA 16/1000", "OLIVIA 32/1000"}, Description: ""},
	{IsImportOnly: false, Key: "opera", Submodes: submode.SubModeList{"OPERA-BEACON", "OPERA-QSO"}, Description: ""},
	{IsImportOnly: false, Key: "pac", Submodes: submode.SubModeList{"PAC2", "PAC3", "PAC4"}, Description: ""},
	{IsImportOnly: true, Key: "pac2", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "pac3", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "pax", Submodes: submode.SubModeList{"PAX2"}, Description: ""},
	{IsImportOnly: true, Key: "pax2", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "pcw", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "pkt", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "psk", Submodes: submode.SubModeList{"8PSK125", "8PSK125F", "8PSK125FL", "8PSK250", "8PSK250F", "8PSK250FL", "8PSK500", "8PSK500F", "8PSK1000", "8PSK1000F", "8PSK1200F", "FSK31", "PSK10", "PSK31", "PSK63", "PSK63F", "PSK63RC4", "PSK63RC5", "PSK63RC10", "PSK63RC20", "PSK63RC32", "PSK125", "PSK125C12", "PSK125R", "PSK125RC10", "PSK125RC12", "PSK125RC16", "PSK125RC4", "PSK125RC5", "PSK250", "PSK250C6", "PSK250R", "PSK250RC2", "PSK250RC3", "PSK250RC5", "PSK250RC6", "PSK250RC7", "PSK500", "PSK500C2", "PSK500C4", "PSK500R", "PSK500RC2", "PSK500RC3", "PSK500RC4", "PSK800C2", "PSK800RC2", "PSK1000", "PSK1000C2", "PSK1000R", "PSK1000RC2", "PSKAM10", "PSKAM31", "PSKAM50", "PSKFEC31", "QPSK31", "QPSK63", "QPSK125", "QPSK250", "QPSK500", "SIM31"}, Description: ""},
	{IsImportOnly: true, Key: "psk10", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "psk125", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "psk2k", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "psk31", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "psk63", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "psk63f", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "pskam10", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "pskam31", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "pskam50", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "pskfec31", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "pskhell", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "q15", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "qpsk125", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "qpsk31", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "qpsk63", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "qra64", Submodes: submode.SubModeList{"QRA64A", "QRA64B", "QRA64C", "QRA64D", "QRA64E"}, Description: ""},
	{IsImportOnly: false, Key: "ros", Submodes: submode.SubModeList{"ROS-EME", "ROS-HF", "ROS-MF"}, Description: ""},
	{IsImportOnly: false, Key: "rtty", Submodes: submode.SubModeList{"ASCI"}, Description: ""},
	{IsImportOnly: false, Key: "rttym", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "ssb", Submodes: submode.SubModeList{"LSB", "USB"}, Description: ""},
	{IsImportOnly: false, Key: "sstv", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "t10", Submodes: submode.SubModeList(nil), Description: "Tonal 10 digital mode with focus on sensitivity, band capacity and resistance to the HF Doppler frequency spread"},
	{IsImportOnly: false, Key: "thor", Submodes: submode.SubModeList{"THOR-M", "THOR4", "THOR5", "THOR8", "THOR11", "THOR16", "THOR22", "THOR25X4", "THOR50X1", "THOR50X2", "THOR100"}, Description: ""},
	{IsImportOnly: false, Key: "thrb", Submodes: submode.SubModeList{"THRBX", "THRBX1", "THRBX2", "THRBX4", "THROB1", "THROB2", "THROB4"}, Description: ""},
	{IsImportOnly: true, Key: "thrbx", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "tor", Submodes: submode.SubModeList{"AMTORFEC", "GTOR", "NAVTEX", "SITORB"}, Description: ""},
	{IsImportOnly: false, Key: "v4", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "voi", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "winmor", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "wspr", Submodes: submode.SubModeList(nil), Description: ""},
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
