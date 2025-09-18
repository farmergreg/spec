// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package mode provides code and constants as defined in ADIF 3.1.6 (Released)
package mode

import (
	"sync"

	"github.com/farmergreg/spec/v6/enum/submode"
)

const (
	AM           Mode = "am"           // am                     =
	AMTORFEC     Mode = "amtorfec"     // Deprecated: amtorfec   =
	ARDOP        Mode = "ardop"        // ardop                  =
	ASCI         Mode = "asci"         // Deprecated: asci       =
	ATV          Mode = "atv"          // atv                    =
	C4FM         Mode = "c4fm"         // Deprecated: c4fm       =
	CHIP         Mode = "chip"         // chip                   = chip64, chip128
	CHIP128      Mode = "chip128"      // Deprecated: chip128    =
	CHIP64       Mode = "chip64"       // Deprecated: chip64     =
	CLO          Mode = "clo"          // clo                    =
	CONTESTI     Mode = "contesti"     // contesti               =
	CW           Mode = "cw"           // cw                     = pcw
	DIGITALVOICE Mode = "digitalvoice" // digitalvoice           = c4fm, dmr, dstar, freedv, m17
	DOMINO       Mode = "domino"       // domino                 = dom-m, dom4, dom5, dom8, dom11, dom16, dom22, dom44, dom88, dominoex, dominof
	DOMINOF      Mode = "dominof"      // Deprecated: dominof    =
	DSTAR        Mode = "dstar"        // Deprecated: dstar      =
	DYNAMIC      Mode = "dynamic"      // dynamic                = vara hf, vara satellite, vara fm 1200, vara fm 9600
	FAX          Mode = "fax"          // fax                    =
	FM           Mode = "fm"           // fm                     =
	FMHELL       Mode = "fmhell"       // Deprecated: fmhell     =
	FSK          Mode = "fsk"          // fsk                    = scamp_fast, scamp_slow, scamp_vslow
	FSK31        Mode = "fsk31"        // Deprecated: fsk31      =
	FSK441       Mode = "fsk441"       // fsk441                 =
	FT8          Mode = "ft8"          // ft8                    =
	GTOR         Mode = "gtor"         // Deprecated: gtor       =
	HELL         Mode = "hell"         // hell                   = fmhell, fskh105, fskh245, fskhell, hell80, hellx5, hellx9, hfsk, pskhell, slowhell
	HELL80       Mode = "hell80"       // Deprecated: hell80     =
	HFSK         Mode = "hfsk"         // Deprecated: hfsk       =
	ISCAT        Mode = "iscat"        // iscat                  = iscat-a, iscat-b
	JT4          Mode = "jt4"          // jt4                    = jt4a, jt4b, jt4c, jt4d, jt4e, jt4f, jt4g
	JT44         Mode = "jt44"         // jt44                   =
	JT4A         Mode = "jt4a"         // Deprecated: jt4a       =
	JT4B         Mode = "jt4b"         // Deprecated: jt4b       =
	JT4C         Mode = "jt4c"         // Deprecated: jt4c       =
	JT4D         Mode = "jt4d"         // Deprecated: jt4d       =
	JT4E         Mode = "jt4e"         // Deprecated: jt4e       =
	JT4F         Mode = "jt4f"         // Deprecated: jt4f       =
	JT4G         Mode = "jt4g"         // Deprecated: jt4g       =
	JT65         Mode = "jt65"         // jt65                   = jt65a, jt65b, jt65b2, jt65c, jt65c2
	JT65A        Mode = "jt65a"        // Deprecated: jt65a      =
	JT65B        Mode = "jt65b"        // Deprecated: jt65b      =
	JT65C        Mode = "jt65c"        // Deprecated: jt65c      =
	JT6M         Mode = "jt6m"         // jt6m                   =
	JT9          Mode = "jt9"          // jt9                    = jt9-1, jt9-2, jt9-5, jt9-10, jt9-30, jt9a, jt9b, jt9c, jt9d, jt9e, jt9e fast, jt9f, jt9f fast, jt9g, jt9g fast, jt9h, jt9h fast
	MFSK         Mode = "mfsk"         // mfsk                   = fsqcall, fst4, fst4w, ft4, js8, jtms, mfsk4, mfsk8, mfsk11, mfsk16, mfsk22, mfsk31, mfsk32, mfsk64, mfsk64l, mfsk128 mfsk128l, q65
	MFSK16       Mode = "mfsk16"       // Deprecated: mfsk16     =
	MFSK8        Mode = "mfsk8"        // Deprecated: mfsk8      =
	MSK144       Mode = "msk144"       // msk144                 =
	MT63         Mode = "mt63"         // mt63                   =
	MTONE        Mode = "mtone"        // mtone                  = scamp_oo, scamp_oo_slw
	OLIVIA       Mode = "olivia"       // olivia                 = olivia 4/125, olivia 4/250, olivia 8/250, olivia 8/500, olivia 16/500, olivia 16/1000, olivia 32/1000
	OPERA        Mode = "opera"        // opera                  = opera-beacon, opera-qso
	PAC          Mode = "pac"          // pac                    = pac2, pac3, pac4
	PAC2         Mode = "pac2"         // Deprecated: pac2       =
	PAC3         Mode = "pac3"         // Deprecated: pac3       =
	PAX          Mode = "pax"          // pax                    = pax2
	PAX2         Mode = "pax2"         // Deprecated: pax2       =
	PCW          Mode = "pcw"          // Deprecated: pcw        =
	PKT          Mode = "pkt"          // pkt                    =
	PSK          Mode = "psk"          // psk                    = 8psk125, 8psk125f, 8psk125fl, 8psk250, 8psk250f, 8psk250fl, 8psk500, 8psk500f, 8psk1000, 8psk1000f, 8psk1200f, fsk31, psk10, psk31, psk63, psk63f, psk63rc4, psk63rc5, psk63rc10, psk63rc20, psk63rc32, psk125, psk125c12, psk125r, psk125rc10, psk125rc12, psk125rc16, psk125rc4, psk125rc5, psk250, psk250c6, psk250r, psk250rc2, psk250rc3, psk250rc5, psk250rc6, psk250rc7, psk500, psk500c2, psk500c4, psk500r, psk500rc2, psk500rc3, psk500rc4, psk800c2, psk800rc2, psk1000, psk1000c2, psk1000r, psk1000rc2, pskam10, pskam31, pskam50, pskfec31, qpsk31, qpsk63, qpsk125, qpsk250, qpsk500, sim31
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
	QRA64        Mode = "qra64"        // qra64                  = qra64a, qra64b, qra64c, qra64d, qra64e
	ROS          Mode = "ros"          // ros                    = ros-eme, ros-hf, ros-mf
	RTTY         Mode = "rtty"         // rtty                   = asci
	RTTYM        Mode = "rttym"        // rttym                  =
	SSB          Mode = "ssb"          // ssb                    = lsb, usb
	SSTV         Mode = "sstv"         // sstv                   =
	T10          Mode = "t10"          // t10                    =
	THOR         Mode = "thor"         // thor                   = thor-m, thor4, thor5, thor8, thor11, thor16, thor22, thor25x4, thor50x1, thor50x2, thor100
	THRB         Mode = "thrb"         // thrb                   = thrbx, thrbx1, thrbx2, thrbx4, throb1, throb2, throb4
	THRBX        Mode = "thrbx"        // Deprecated: thrbx      =
	TOR          Mode = "tor"          // tor                    = amtorfec, gtor, navtex, sitorb
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
	{IsImportOnly: false, Key: "chip", Submodes: submode.SubModeList{"chip64", "chip128"}, Description: ""},
	{IsImportOnly: true, Key: "chip128", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "chip64", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "clo", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "contesti", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "cw", Submodes: submode.SubModeList{"pcw"}, Description: ""},
	{IsImportOnly: false, Key: "digitalvoice", Submodes: submode.SubModeList{"c4fm", "dmr", "dstar", "freedv", "m17"}, Description: ""},
	{IsImportOnly: false, Key: "domino", Submodes: submode.SubModeList{"dom-m", "dom4", "dom5", "dom8", "dom11", "dom16", "dom22", "dom44", "dom88", "dominoex", "dominof"}, Description: ""},
	{IsImportOnly: true, Key: "dominof", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "dstar", Submodes: submode.SubModeList(nil), Description: "Digital Smart Technologies for Amateur Radio Imported QSOs with <MODE:5>DSTAR must be exported as: <MODE:12>DIGITALVOICE <SUBMODE:5>DSTAR"},
	{IsImportOnly: false, Key: "dynamic", Submodes: submode.SubModeList{"vara hf", "vara satellite", "vara fm 1200", "vara fm 9600"}, Description: ""},
	{IsImportOnly: false, Key: "fax", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "fm", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "fmhell", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "fsk", Submodes: submode.SubModeList{"scamp_fast", "scamp_slow", "scamp_vslow"}, Description: "Frequency shift keying"},
	{IsImportOnly: true, Key: "fsk31", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "fsk441", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "ft8", Submodes: submode.SubModeList(nil), Description: "Franke-Taylor design, 8-FSK modulation"},
	{IsImportOnly: true, Key: "gtor", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "hell", Submodes: submode.SubModeList{"fmhell", "fskh105", "fskh245", "fskhell", "hell80", "hellx5", "hellx9", "hfsk", "pskhell", "slowhell"}, Description: ""},
	{IsImportOnly: true, Key: "hell80", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "hfsk", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "iscat", Submodes: submode.SubModeList{"iscat-a", "iscat-b"}, Description: ""},
	{IsImportOnly: false, Key: "jt4", Submodes: submode.SubModeList{"jt4a", "jt4b", "jt4c", "jt4d", "jt4e", "jt4f", "jt4g"}, Description: ""},
	{IsImportOnly: false, Key: "jt44", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4a", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4b", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4c", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4d", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4e", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4f", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt4g", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "jt65", Submodes: submode.SubModeList{"jt65a", "jt65b", "jt65b2", "jt65c", "jt65c2"}, Description: ""},
	{IsImportOnly: true, Key: "jt65a", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt65b", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "jt65c", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "jt6m", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "jt9", Submodes: submode.SubModeList{"jt9-1", "jt9-2", "jt9-5", "jt9-10", "jt9-30", "jt9a", "jt9b", "jt9c", "jt9d", "jt9e", "jt9e fast", "jt9f", "jt9f fast", "jt9g", "jt9g fast", "jt9h", "jt9h fast"}, Description: ""},
	{IsImportOnly: false, Key: "mfsk", Submodes: submode.SubModeList{"fsqcall", "fst4", "fst4w", "ft4", "js8", "jtms", "mfsk4", "mfsk8", "mfsk11", "mfsk16", "mfsk22", "mfsk31", "mfsk32", "mfsk64", "mfsk64l", "mfsk128 mfsk128l", "q65"}, Description: ""},
	{IsImportOnly: true, Key: "mfsk16", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "mfsk8", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "msk144", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "mt63", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "mtone", Submodes: submode.SubModeList{"scamp_oo", "scamp_oo_slw"}, Description: "Single modulated tone"},
	{IsImportOnly: false, Key: "olivia", Submodes: submode.SubModeList{"olivia 4/125", "olivia 4/250", "olivia 8/250", "olivia 8/500", "olivia 16/500", "olivia 16/1000", "olivia 32/1000"}, Description: ""},
	{IsImportOnly: false, Key: "opera", Submodes: submode.SubModeList{"opera-beacon", "opera-qso"}, Description: ""},
	{IsImportOnly: false, Key: "pac", Submodes: submode.SubModeList{"pac2", "pac3", "pac4"}, Description: ""},
	{IsImportOnly: true, Key: "pac2", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "pac3", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "pax", Submodes: submode.SubModeList{"pax2"}, Description: ""},
	{IsImportOnly: true, Key: "pax2", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: true, Key: "pcw", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "pkt", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "psk", Submodes: submode.SubModeList{"8psk125", "8psk125f", "8psk125fl", "8psk250", "8psk250f", "8psk250fl", "8psk500", "8psk500f", "8psk1000", "8psk1000f", "8psk1200f", "fsk31", "psk10", "psk31", "psk63", "psk63f", "psk63rc4", "psk63rc5", "psk63rc10", "psk63rc20", "psk63rc32", "psk125", "psk125c12", "psk125r", "psk125rc10", "psk125rc12", "psk125rc16", "psk125rc4", "psk125rc5", "psk250", "psk250c6", "psk250r", "psk250rc2", "psk250rc3", "psk250rc5", "psk250rc6", "psk250rc7", "psk500", "psk500c2", "psk500c4", "psk500r", "psk500rc2", "psk500rc3", "psk500rc4", "psk800c2", "psk800rc2", "psk1000", "psk1000c2", "psk1000r", "psk1000rc2", "pskam10", "pskam31", "pskam50", "pskfec31", "qpsk31", "qpsk63", "qpsk125", "qpsk250", "qpsk500", "sim31"}, Description: ""},
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
	{IsImportOnly: false, Key: "qra64", Submodes: submode.SubModeList{"qra64a", "qra64b", "qra64c", "qra64d", "qra64e"}, Description: ""},
	{IsImportOnly: false, Key: "ros", Submodes: submode.SubModeList{"ros-eme", "ros-hf", "ros-mf"}, Description: ""},
	{IsImportOnly: false, Key: "rtty", Submodes: submode.SubModeList{"asci"}, Description: ""},
	{IsImportOnly: false, Key: "rttym", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "ssb", Submodes: submode.SubModeList{"lsb", "usb"}, Description: ""},
	{IsImportOnly: false, Key: "sstv", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "t10", Submodes: submode.SubModeList(nil), Description: "Tonal 10 digital mode with focus on sensitivity, band capacity and resistance to the HF Doppler frequency spread"},
	{IsImportOnly: false, Key: "thor", Submodes: submode.SubModeList{"thor-m", "thor4", "thor5", "thor8", "thor11", "thor16", "thor22", "thor25x4", "thor50x1", "thor50x2", "thor100"}, Description: ""},
	{IsImportOnly: false, Key: "thrb", Submodes: submode.SubModeList{"thrbx", "thrbx1", "thrbx2", "thrbx4", "throb1", "throb2", "throb4"}, Description: ""},
	{IsImportOnly: true, Key: "thrbx", Submodes: submode.SubModeList(nil), Description: ""},
	{IsImportOnly: false, Key: "tor", Submodes: submode.SubModeList{"amtorfec", "gtor", "navtex", "sitorb"}, Description: ""},
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
