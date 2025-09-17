// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package submode provides code and constants as defined in ADIF 3.1.6 (Released)
package submode

import "sync"

const (
	SUBMODE_8PSK1000       SubMode = "8psk1000"       // 8psk1000        = psk
	SUBMODE_8PSK1000F      SubMode = "8psk1000f"      // 8psk1000f       = psk
	SUBMODE_8PSK1200F      SubMode = "8psk1200f"      // 8psk1200f       = psk
	SUBMODE_8PSK125        SubMode = "8psk125"        // 8psk125         = psk
	SUBMODE_8PSK125F       SubMode = "8psk125f"       // 8psk125f        = psk
	SUBMODE_8PSK125FL      SubMode = "8psk125fl"      // 8psk125fl       = psk
	SUBMODE_8PSK250        SubMode = "8psk250"        // 8psk250         = psk
	SUBMODE_8PSK250F       SubMode = "8psk250f"       // 8psk250f        = psk
	SUBMODE_8PSK250FL      SubMode = "8psk250fl"      // 8psk250fl       = psk
	SUBMODE_8PSK500        SubMode = "8psk500"        // 8psk500         = psk
	SUBMODE_8PSK500F       SubMode = "8psk500f"       // 8psk500f        = psk
	SUBMODE_AMTORFEC       SubMode = "amtorfec"       // amtorfec        = tor
	SUBMODE_ASCI           SubMode = "asci"           // asci            = rtty
	SUBMODE_C4FM           SubMode = "c4fm"           // c4fm            = digitalvoice    C4FM 4-level FSK See the Propagation_Mode enumeration section for examples of representing C4FM voice transmissions.
	SUBMODE_CHIP128        SubMode = "chip128"        // chip128         = chip
	SUBMODE_CHIP64         SubMode = "chip64"         // chip64          = chip
	SUBMODE_DMR            SubMode = "dmr"            // dmr             = digitalvoice    Digital Mobile Radio See the Propagation_Mode enumeration section for examples of representing DMR voice transmissions.
	SUBMODE_DOM_M          SubMode = "dom-m"          // dom-m           = domino
	SUBMODE_DOM11          SubMode = "dom11"          // dom11           = domino
	SUBMODE_DOM16          SubMode = "dom16"          // dom16           = domino
	SUBMODE_DOM22          SubMode = "dom22"          // dom22           = domino
	SUBMODE_DOM4           SubMode = "dom4"           // dom4            = domino
	SUBMODE_DOM44          SubMode = "dom44"          // dom44           = domino
	SUBMODE_DOM5           SubMode = "dom5"           // dom5            = domino
	SUBMODE_DOM8           SubMode = "dom8"           // dom8            = domino
	SUBMODE_DOM88          SubMode = "dom88"          // dom88           = domino
	SUBMODE_DOMINOEX       SubMode = "dominoex"       // dominoex        = domino
	SUBMODE_DOMINOF        SubMode = "dominof"        // dominof         = domino
	SUBMODE_DSTAR          SubMode = "dstar"          // dstar           = digitalvoice    Digital Smart Technologies for Amateur Radio See the Propagation_Mode enumeration section for examples of representing DSTAR voice transmissions.
	SUBMODE_FMHELL         SubMode = "fmhell"         // fmhell          = hell
	SUBMODE_FREEDV         SubMode = "freedv"         // freedv          = digitalvoice    Digital voice mode for HF radio implemented with open source
	SUBMODE_FSK31          SubMode = "fsk31"          // fsk31           = psk
	SUBMODE_FSKH105        SubMode = "fskh105"        // fskh105         = hell
	SUBMODE_FSKH245        SubMode = "fskh245"        // fskh245         = hell
	SUBMODE_FSKHELL        SubMode = "fskhell"        // fskhell         = hell
	SUBMODE_FSQCALL        SubMode = "fsqcall"        // fsqcall         = mfsk            FSQCall protocol used with FSQ (Fast Simple QSO) transmission mode
	SUBMODE_FST4           SubMode = "fst4"           // fst4            = mfsk            This is a digital mode supported by the WSJT-X software
	SUBMODE_FST4W          SubMode = "fst4w"          // fst4w           = mfsk            This is a digital mode supported by the WSJT-X software that is for quasi-beacon transmissions of WSPR-style messages
	SUBMODE_FT4            SubMode = "ft4"            // ft4             = mfsk            FT4 is a digital mode designed specifically for radio contesting
	SUBMODE_GTOR           SubMode = "gtor"           // gtor            = tor
	SUBMODE_HELL80         SubMode = "hell80"         // hell80          = hell
	SUBMODE_HELLX5         SubMode = "hellx5"         // hellx5          = hell
	SUBMODE_HELLX9         SubMode = "hellx9"         // hellx9          = hell
	SUBMODE_HFSK           SubMode = "hfsk"           // hfsk            = hell
	SUBMODE_ISCAT_A        SubMode = "iscat-a"        // iscat-a         = iscat
	SUBMODE_ISCAT_B        SubMode = "iscat-b"        // iscat-b         = iscat
	SUBMODE_JS8            SubMode = "js8"            // js8             = mfsk            Jordan Sherer designed 8-FSK modulation
	SUBMODE_JT4A           SubMode = "jt4a"           // jt4a            = jt4
	SUBMODE_JT4B           SubMode = "jt4b"           // jt4b            = jt4
	SUBMODE_JT4C           SubMode = "jt4c"           // jt4c            = jt4
	SUBMODE_JT4D           SubMode = "jt4d"           // jt4d            = jt4
	SUBMODE_JT4E           SubMode = "jt4e"           // jt4e            = jt4
	SUBMODE_JT4F           SubMode = "jt4f"           // jt4f            = jt4
	SUBMODE_JT4G           SubMode = "jt4g"           // jt4g            = jt4
	SUBMODE_JT65A          SubMode = "jt65a"          // jt65a           = jt65
	SUBMODE_JT65B          SubMode = "jt65b"          // jt65b           = jt65
	SUBMODE_JT65B2         SubMode = "jt65b2"         // jt65b2          = jt65
	SUBMODE_JT65C          SubMode = "jt65c"          // jt65c           = jt65
	SUBMODE_JT65C2         SubMode = "jt65c2"         // jt65c2          = jt65
	SUBMODE_JT9_1          SubMode = "jt9-1"          // jt9-1           = jt9
	SUBMODE_JT9_10         SubMode = "jt9-10"         // jt9-10          = jt9
	SUBMODE_JT9_2          SubMode = "jt9-2"          // jt9-2           = jt9
	SUBMODE_JT9_30         SubMode = "jt9-30"         // jt9-30          = jt9
	SUBMODE_JT9_5          SubMode = "jt9-5"          // jt9-5           = jt9
	SUBMODE_JT9A           SubMode = "jt9a"           // jt9a            = jt9
	SUBMODE_JT9B           SubMode = "jt9b"           // jt9b            = jt9
	SUBMODE_JT9C           SubMode = "jt9c"           // jt9c            = jt9
	SUBMODE_JT9D           SubMode = "jt9d"           // jt9d            = jt9
	SUBMODE_JT9E           SubMode = "jt9e"           // jt9e            = jt9
	SUBMODE_JT9E_FAST      SubMode = "jt9e fast"      // jt9e fast       = jt9
	SUBMODE_JT9F           SubMode = "jt9f"           // jt9f            = jt9
	SUBMODE_JT9F_FAST      SubMode = "jt9f fast"      // jt9f fast       = jt9
	SUBMODE_JT9G           SubMode = "jt9g"           // jt9g            = jt9
	SUBMODE_JT9G_FAST      SubMode = "jt9g fast"      // jt9g fast       = jt9
	SUBMODE_JT9H           SubMode = "jt9h"           // jt9h            = jt9
	SUBMODE_JT9H_FAST      SubMode = "jt9h fast"      // jt9h fast       = jt9
	SUBMODE_JTMS           SubMode = "jtms"           // jtms            = mfsk
	SUBMODE_LSB            SubMode = "lsb"            // lsb             = ssb             Amplitude modulated voice telephony, lower-sideband, suppressed-carrier
	SUBMODE_M17            SubMode = "m17"            // m17             = digitalvoice    Digital radio protocol using the Codec 2 voice encoder
	SUBMODE_MFSK11         SubMode = "mfsk11"         // mfsk11          = mfsk
	SUBMODE_MFSK128        SubMode = "mfsk128"        // mfsk128         = mfsk
	SUBMODE_MFSK128L       SubMode = "mfsk128l"       // mfsk128l        = mfsk
	SUBMODE_MFSK16         SubMode = "mfsk16"         // mfsk16          = mfsk
	SUBMODE_MFSK22         SubMode = "mfsk22"         // mfsk22          = mfsk
	SUBMODE_MFSK31         SubMode = "mfsk31"         // mfsk31          = mfsk
	SUBMODE_MFSK32         SubMode = "mfsk32"         // mfsk32          = mfsk
	SUBMODE_MFSK4          SubMode = "mfsk4"          // mfsk4           = mfsk
	SUBMODE_MFSK64         SubMode = "mfsk64"         // mfsk64          = mfsk
	SUBMODE_MFSK64L        SubMode = "mfsk64l"        // mfsk64l         = mfsk
	SUBMODE_MFSK8          SubMode = "mfsk8"          // mfsk8           = mfsk
	SUBMODE_NAVTEX         SubMode = "navtex"         // navtex          = tor
	SUBMODE_OLIVIA_16_1000 SubMode = "olivia 16/1000" // olivia 16/1000  = olivia
	SUBMODE_OLIVIA_16_500  SubMode = "olivia 16/500"  // olivia 16/500   = olivia
	SUBMODE_OLIVIA_32_1000 SubMode = "olivia 32/1000" // olivia 32/1000  = olivia
	SUBMODE_OLIVIA_4_125   SubMode = "olivia 4/125"   // olivia 4/125    = olivia
	SUBMODE_OLIVIA_4_250   SubMode = "olivia 4/250"   // olivia 4/250    = olivia
	SUBMODE_OLIVIA_8_250   SubMode = "olivia 8/250"   // olivia 8/250    = olivia
	SUBMODE_OLIVIA_8_500   SubMode = "olivia 8/500"   // olivia 8/500    = olivia
	SUBMODE_OPERA_BEACON   SubMode = "opera-beacon"   // opera-beacon    = opera
	SUBMODE_OPERA_QSO      SubMode = "opera-qso"      // opera-qso       = opera
	SUBMODE_PAC2           SubMode = "pac2"           // pac2            = pac
	SUBMODE_PAC3           SubMode = "pac3"           // pac3            = pac
	SUBMODE_PAC4           SubMode = "pac4"           // pac4            = pac
	SUBMODE_PAX2           SubMode = "pax2"           // pax2            = pax
	SUBMODE_PCW            SubMode = "pcw"            // pcw             = cw              Coherent CW
	SUBMODE_PSK10          SubMode = "psk10"          // psk10           = psk
	SUBMODE_PSK1000        SubMode = "psk1000"        // psk1000         = psk
	SUBMODE_PSK1000RC2     SubMode = "psk1000rc2"     // psk1000rc2      = psk
	SUBMODE_PSK125         SubMode = "psk125"         // psk125          = psk
	SUBMODE_PSK125RC10     SubMode = "psk125rc10"     // psk125rc10      = psk
	SUBMODE_PSK125RC12     SubMode = "psk125rc12"     // psk125rc12      = psk
	SUBMODE_PSK125RC16     SubMode = "psk125rc16"     // psk125rc16      = psk
	SUBMODE_PSK125RC4      SubMode = "psk125rc4"      // psk125rc4       = psk
	SUBMODE_PSK125RC5      SubMode = "psk125rc5"      // psk125rc5       = psk
	SUBMODE_PSK250         SubMode = "psk250"         // psk250          = psk
	SUBMODE_PSK250RC2      SubMode = "psk250rc2"      // psk250rc2       = psk
	SUBMODE_PSK250RC3      SubMode = "psk250rc3"      // psk250rc3       = psk
	SUBMODE_PSK250RC5      SubMode = "psk250rc5"      // psk250rc5       = psk
	SUBMODE_PSK250RC6      SubMode = "psk250rc6"      // psk250rc6       = psk
	SUBMODE_PSK250RC7      SubMode = "psk250rc7"      // psk250rc7       = psk
	SUBMODE_PSK31          SubMode = "psk31"          // psk31           = psk
	SUBMODE_PSK500         SubMode = "psk500"         // psk500          = psk
	SUBMODE_PSK500RC2      SubMode = "psk500rc2"      // psk500rc2       = psk
	SUBMODE_PSK500RC3      SubMode = "psk500rc3"      // psk500rc3       = psk
	SUBMODE_PSK500RC4      SubMode = "psk500rc4"      // psk500rc4       = psk
	SUBMODE_PSK63          SubMode = "psk63"          // psk63           = psk
	SUBMODE_PSK63F         SubMode = "psk63f"         // psk63f          = psk
	SUBMODE_PSK63RC10      SubMode = "psk63rc10"      // psk63rc10       = psk
	SUBMODE_PSK63RC20      SubMode = "psk63rc20"      // psk63rc20       = psk
	SUBMODE_PSK63RC32      SubMode = "psk63rc32"      // psk63rc32       = psk
	SUBMODE_PSK63RC4       SubMode = "psk63rc4"       // psk63rc4        = psk
	SUBMODE_PSK63RC5       SubMode = "psk63rc5"       // psk63rc5        = psk
	SUBMODE_PSK800RC2      SubMode = "psk800rc2"      // psk800rc2       = psk
	SUBMODE_PSKAM10        SubMode = "pskam10"        // pskam10         = psk
	SUBMODE_PSKAM31        SubMode = "pskam31"        // pskam31         = psk
	SUBMODE_PSKAM50        SubMode = "pskam50"        // pskam50         = psk
	SUBMODE_PSKFEC31       SubMode = "pskfec31"       // pskfec31        = psk
	SUBMODE_PSKHELL        SubMode = "pskhell"        // pskhell         = hell
	SUBMODE_Q65            SubMode = "q65"            // q65             = mfsk
	SUBMODE_QPSK125        SubMode = "qpsk125"        // qpsk125         = psk
	SUBMODE_QPSK250        SubMode = "qpsk250"        // qpsk250         = psk
	SUBMODE_QPSK31         SubMode = "qpsk31"         // qpsk31          = psk
	SUBMODE_QPSK500        SubMode = "qpsk500"        // qpsk500         = psk
	SUBMODE_QPSK63         SubMode = "qpsk63"         // qpsk63          = psk
	SUBMODE_QRA64A         SubMode = "qra64a"         // qra64a          = qra64
	SUBMODE_QRA64B         SubMode = "qra64b"         // qra64b          = qra64
	SUBMODE_QRA64C         SubMode = "qra64c"         // qra64c          = qra64
	SUBMODE_QRA64D         SubMode = "qra64d"         // qra64d          = qra64
	SUBMODE_QRA64E         SubMode = "qra64e"         // qra64e          = qra64
	SUBMODE_ROS_EME        SubMode = "ros-eme"        // ros-eme         = ros
	SUBMODE_ROS_HF         SubMode = "ros-hf"         // ros-hf          = ros
	SUBMODE_ROS_MF         SubMode = "ros-mf"         // ros-mf          = ros
	SUBMODE_SCAMP_FAST     SubMode = "scamp_fast"     // scamp_fast      = fsk             SCAMP fast FSK
	SUBMODE_SCAMP_OO       SubMode = "scamp_oo"       // scamp_oo        = mtone           SCAMP single modulated tone on/off keying
	SUBMODE_SCAMP_OO_SLW   SubMode = "scamp_oo_slw"   // scamp_oo_slw    = mtone           SCAMP single modulated tone on/off slow keying
	SUBMODE_SCAMP_SLOW     SubMode = "scamp_slow"     // scamp_slow      = fsk             SCAMP slow FSK
	SUBMODE_SCAMP_VSLOW    SubMode = "scamp_vslow"    // scamp_vslow     = fsk             SCAMP very slow FSK
	SUBMODE_SIM31          SubMode = "sim31"          // sim31           = psk
	SUBMODE_SITORB         SubMode = "sitorb"         // sitorb          = tor
	SUBMODE_SLOWHELL       SubMode = "slowhell"       // slowhell        = hell
	SUBMODE_THOR_M         SubMode = "thor-m"         // thor-m          = thor
	SUBMODE_THOR100        SubMode = "thor100"        // thor100         = thor
	SUBMODE_THOR11         SubMode = "thor11"         // thor11          = thor
	SUBMODE_THOR16         SubMode = "thor16"         // thor16          = thor
	SUBMODE_THOR22         SubMode = "thor22"         // thor22          = thor
	SUBMODE_THOR25X4       SubMode = "thor25x4"       // thor25x4        = thor
	SUBMODE_THOR4          SubMode = "thor4"          // thor4           = thor
	SUBMODE_THOR5          SubMode = "thor5"          // thor5           = thor
	SUBMODE_THOR50X1       SubMode = "thor50x1"       // thor50x1        = thor
	SUBMODE_THOR50X2       SubMode = "thor50x2"       // thor50x2        = thor
	SUBMODE_THOR8          SubMode = "thor8"          // thor8           = thor
	SUBMODE_THRBX          SubMode = "thrbx"          // thrbx           = thrb
	SUBMODE_THRBX1         SubMode = "thrbx1"         // thrbx1          = thrb
	SUBMODE_THRBX2         SubMode = "thrbx2"         // thrbx2          = thrb
	SUBMODE_THRBX4         SubMode = "thrbx4"         // thrbx4          = thrb
	SUBMODE_THROB1         SubMode = "throb1"         // throb1          = thrb
	SUBMODE_THROB2         SubMode = "throb2"         // throb2          = thrb
	SUBMODE_THROB4         SubMode = "throb4"         // throb4          = thrb
	SUBMODE_USB            SubMode = "usb"            // usb             = ssb             Amplitude modulated voice telephony, upper-sideband, suppressed-carrier
	SUBMODE_VARA_FM_1200   SubMode = "vara fm 1200"   // vara fm 1200    = dynamic         Channel adaptive high-speed modem for FM transceivers
	SUBMODE_VARA_FM_9600   SubMode = "vara fm 9600"   // vara fm 9600    = dynamic         Channel adaptive high-speed modem for FM transceivers
	SUBMODE_VARA_HF        SubMode = "vara hf"        // vara hf         = dynamic         Channel adaptive high-speed modem for HF
	SUBMODE_VARA_SATELLITE SubMode = "vara satellite" // vara satellite  = dynamic         Channel adaptive high-speed modem for satellite operations
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known SubMode specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "8psk1000", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "8psk1000f", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "8psk1200f", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "8psk125", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "8psk125f", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "8psk125fl", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "8psk250", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "8psk250f", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "8psk250fl", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "8psk500", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "8psk500f", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "amtorfec", Mode: "tor", Description: ""},
	{IsImportOnly: false, Key: "asci", Mode: "rtty", Description: ""},
	{IsImportOnly: false, Key: "c4fm", Mode: "digitalvoice", Description: "C4FM 4-level FSK See the Propagation_Mode enumeration section for examples of representing C4FM voice transmissions."},
	{IsImportOnly: false, Key: "chip128", Mode: "chip", Description: ""},
	{IsImportOnly: false, Key: "chip64", Mode: "chip", Description: ""},
	{IsImportOnly: false, Key: "dmr", Mode: "digitalvoice", Description: "Digital Mobile Radio See the Propagation_Mode enumeration section for examples of representing DMR voice transmissions."},
	{IsImportOnly: false, Key: "dom-m", Mode: "domino", Description: ""},
	{IsImportOnly: false, Key: "dom11", Mode: "domino", Description: ""},
	{IsImportOnly: false, Key: "dom16", Mode: "domino", Description: ""},
	{IsImportOnly: false, Key: "dom22", Mode: "domino", Description: ""},
	{IsImportOnly: false, Key: "dom4", Mode: "domino", Description: ""},
	{IsImportOnly: false, Key: "dom44", Mode: "domino", Description: ""},
	{IsImportOnly: false, Key: "dom5", Mode: "domino", Description: ""},
	{IsImportOnly: false, Key: "dom8", Mode: "domino", Description: ""},
	{IsImportOnly: false, Key: "dom88", Mode: "domino", Description: ""},
	{IsImportOnly: false, Key: "dominoex", Mode: "domino", Description: ""},
	{IsImportOnly: false, Key: "dominof", Mode: "domino", Description: ""},
	{IsImportOnly: false, Key: "dstar", Mode: "digitalvoice", Description: "Digital Smart Technologies for Amateur Radio See the Propagation_Mode enumeration section for examples of representing DSTAR voice transmissions."},
	{IsImportOnly: false, Key: "fmhell", Mode: "hell", Description: ""},
	{IsImportOnly: false, Key: "freedv", Mode: "digitalvoice", Description: "Digital voice mode for HF radio implemented with open source"},
	{IsImportOnly: false, Key: "fsk31", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "fskh105", Mode: "hell", Description: ""},
	{IsImportOnly: false, Key: "fskh245", Mode: "hell", Description: ""},
	{IsImportOnly: false, Key: "fskhell", Mode: "hell", Description: ""},
	{IsImportOnly: false, Key: "fsqcall", Mode: "mfsk", Description: "FSQCall protocol used with FSQ (Fast Simple QSO) transmission mode"},
	{IsImportOnly: false, Key: "fst4", Mode: "mfsk", Description: "This is a digital mode supported by the WSJT-X software"},
	{IsImportOnly: false, Key: "fst4w", Mode: "mfsk", Description: "This is a digital mode supported by the WSJT-X software that is for quasi-beacon transmissions of WSPR-style messages"},
	{IsImportOnly: false, Key: "ft4", Mode: "mfsk", Description: "FT4 is a digital mode designed specifically for radio contesting"},
	{IsImportOnly: false, Key: "gtor", Mode: "tor", Description: ""},
	{IsImportOnly: false, Key: "hell80", Mode: "hell", Description: ""},
	{IsImportOnly: false, Key: "hellx5", Mode: "hell", Description: ""},
	{IsImportOnly: false, Key: "hellx9", Mode: "hell", Description: ""},
	{IsImportOnly: false, Key: "hfsk", Mode: "hell", Description: ""},
	{IsImportOnly: false, Key: "iscat-a", Mode: "iscat", Description: ""},
	{IsImportOnly: false, Key: "iscat-b", Mode: "iscat", Description: ""},
	{IsImportOnly: false, Key: "js8", Mode: "mfsk", Description: "Jordan Sherer designed 8-FSK modulation"},
	{IsImportOnly: false, Key: "jt4a", Mode: "jt4", Description: ""},
	{IsImportOnly: false, Key: "jt4b", Mode: "jt4", Description: ""},
	{IsImportOnly: false, Key: "jt4c", Mode: "jt4", Description: ""},
	{IsImportOnly: false, Key: "jt4d", Mode: "jt4", Description: ""},
	{IsImportOnly: false, Key: "jt4e", Mode: "jt4", Description: ""},
	{IsImportOnly: false, Key: "jt4f", Mode: "jt4", Description: ""},
	{IsImportOnly: false, Key: "jt4g", Mode: "jt4", Description: ""},
	{IsImportOnly: false, Key: "jt65a", Mode: "jt65", Description: ""},
	{IsImportOnly: false, Key: "jt65b", Mode: "jt65", Description: ""},
	{IsImportOnly: false, Key: "jt65b2", Mode: "jt65", Description: ""},
	{IsImportOnly: false, Key: "jt65c", Mode: "jt65", Description: ""},
	{IsImportOnly: false, Key: "jt65c2", Mode: "jt65", Description: ""},
	{IsImportOnly: false, Key: "jt9-1", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9-10", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9-2", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9-30", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9-5", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9a", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9b", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9c", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9d", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9e", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9e fast", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9f", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9f fast", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9g", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9g fast", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9h", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jt9h fast", Mode: "jt9", Description: ""},
	{IsImportOnly: false, Key: "jtms", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "lsb", Mode: "ssb", Description: "Amplitude modulated voice telephony, lower-sideband, suppressed-carrier"},
	{IsImportOnly: false, Key: "m17", Mode: "digitalvoice", Description: "Digital radio protocol using the Codec 2 voice encoder"},
	{IsImportOnly: false, Key: "mfsk11", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "mfsk128", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "mfsk128l", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "mfsk16", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "mfsk22", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "mfsk31", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "mfsk32", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "mfsk4", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "mfsk64", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "mfsk64l", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "mfsk8", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "navtex", Mode: "tor", Description: ""},
	{IsImportOnly: false, Key: "olivia 16/1000", Mode: "olivia", Description: ""},
	{IsImportOnly: false, Key: "olivia 16/500", Mode: "olivia", Description: ""},
	{IsImportOnly: false, Key: "olivia 32/1000", Mode: "olivia", Description: ""},
	{IsImportOnly: false, Key: "olivia 4/125", Mode: "olivia", Description: ""},
	{IsImportOnly: false, Key: "olivia 4/250", Mode: "olivia", Description: ""},
	{IsImportOnly: false, Key: "olivia 8/250", Mode: "olivia", Description: ""},
	{IsImportOnly: false, Key: "olivia 8/500", Mode: "olivia", Description: ""},
	{IsImportOnly: false, Key: "opera-beacon", Mode: "opera", Description: ""},
	{IsImportOnly: false, Key: "opera-qso", Mode: "opera", Description: ""},
	{IsImportOnly: false, Key: "pac2", Mode: "pac", Description: ""},
	{IsImportOnly: false, Key: "pac3", Mode: "pac", Description: ""},
	{IsImportOnly: false, Key: "pac4", Mode: "pac", Description: ""},
	{IsImportOnly: false, Key: "pax2", Mode: "pax", Description: ""},
	{IsImportOnly: false, Key: "pcw", Mode: "cw", Description: "Coherent CW"},
	{IsImportOnly: false, Key: "psk10", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk1000", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk1000rc2", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk125", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk125rc10", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk125rc12", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk125rc16", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk125rc4", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk125rc5", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk250", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk250rc2", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk250rc3", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk250rc5", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk250rc6", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk250rc7", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk31", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk500", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk500rc2", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk500rc3", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk500rc4", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk63", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk63f", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk63rc10", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk63rc20", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk63rc32", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk63rc4", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk63rc5", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "psk800rc2", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "pskam10", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "pskam31", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "pskam50", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "pskfec31", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "pskhell", Mode: "hell", Description: ""},
	{IsImportOnly: false, Key: "q65", Mode: "mfsk", Description: ""},
	{IsImportOnly: false, Key: "qpsk125", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "qpsk250", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "qpsk31", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "qpsk500", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "qpsk63", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "qra64a", Mode: "qra64", Description: ""},
	{IsImportOnly: false, Key: "qra64b", Mode: "qra64", Description: ""},
	{IsImportOnly: false, Key: "qra64c", Mode: "qra64", Description: ""},
	{IsImportOnly: false, Key: "qra64d", Mode: "qra64", Description: ""},
	{IsImportOnly: false, Key: "qra64e", Mode: "qra64", Description: ""},
	{IsImportOnly: false, Key: "ros-eme", Mode: "ros", Description: ""},
	{IsImportOnly: false, Key: "ros-hf", Mode: "ros", Description: ""},
	{IsImportOnly: false, Key: "ros-mf", Mode: "ros", Description: ""},
	{IsImportOnly: false, Key: "scamp_fast", Mode: "fsk", Description: "SCAMP fast FSK"},
	{IsImportOnly: false, Key: "scamp_oo", Mode: "mtone", Description: "SCAMP single modulated tone on/off keying"},
	{IsImportOnly: false, Key: "scamp_oo_slw", Mode: "mtone", Description: "SCAMP single modulated tone on/off slow keying"},
	{IsImportOnly: false, Key: "scamp_slow", Mode: "fsk", Description: "SCAMP slow FSK"},
	{IsImportOnly: false, Key: "scamp_vslow", Mode: "fsk", Description: "SCAMP very slow FSK"},
	{IsImportOnly: false, Key: "sim31", Mode: "psk", Description: ""},
	{IsImportOnly: false, Key: "sitorb", Mode: "tor", Description: ""},
	{IsImportOnly: false, Key: "slowhell", Mode: "hell", Description: ""},
	{IsImportOnly: false, Key: "thor-m", Mode: "thor", Description: ""},
	{IsImportOnly: false, Key: "thor100", Mode: "thor", Description: ""},
	{IsImportOnly: false, Key: "thor11", Mode: "thor", Description: ""},
	{IsImportOnly: false, Key: "thor16", Mode: "thor", Description: ""},
	{IsImportOnly: false, Key: "thor22", Mode: "thor", Description: ""},
	{IsImportOnly: false, Key: "thor25x4", Mode: "thor", Description: ""},
	{IsImportOnly: false, Key: "thor4", Mode: "thor", Description: ""},
	{IsImportOnly: false, Key: "thor5", Mode: "thor", Description: ""},
	{IsImportOnly: false, Key: "thor50x1", Mode: "thor", Description: ""},
	{IsImportOnly: false, Key: "thor50x2", Mode: "thor", Description: ""},
	{IsImportOnly: false, Key: "thor8", Mode: "thor", Description: ""},
	{IsImportOnly: false, Key: "thrbx", Mode: "thrb", Description: ""},
	{IsImportOnly: false, Key: "thrbx1", Mode: "thrb", Description: ""},
	{IsImportOnly: false, Key: "thrbx2", Mode: "thrb", Description: ""},
	{IsImportOnly: false, Key: "thrbx4", Mode: "thrb", Description: ""},
	{IsImportOnly: false, Key: "throb1", Mode: "thrb", Description: ""},
	{IsImportOnly: false, Key: "throb2", Mode: "thrb", Description: ""},
	{IsImportOnly: false, Key: "throb4", Mode: "thrb", Description: ""},
	{IsImportOnly: false, Key: "usb", Mode: "ssb", Description: "Amplitude modulated voice telephony, upper-sideband, suppressed-carrier"},
	{IsImportOnly: false, Key: "vara fm 1200", Mode: "dynamic", Description: "Channel adaptive high-speed modem for FM transceivers"},
	{IsImportOnly: false, Key: "vara fm 9600", Mode: "dynamic", Description: "Channel adaptive high-speed modem for FM transceivers"},
	{IsImportOnly: false, Key: "vara hf", Mode: "dynamic", Description: "Channel adaptive high-speed modem for HF"},
	{IsImportOnly: false, Key: "vara satellite", Mode: "dynamic", Description: "Channel adaptive high-speed modem for satellite operations"},
}

// lookupMap contains all known SubMode specifications.
var lookupMap = map[SubMode]*Spec{
	SUBMODE_8PSK1000:       &lookupList[0],
	SUBMODE_8PSK1000F:      &lookupList[1],
	SUBMODE_8PSK1200F:      &lookupList[2],
	SUBMODE_8PSK125:        &lookupList[3],
	SUBMODE_8PSK125F:       &lookupList[4],
	SUBMODE_8PSK125FL:      &lookupList[5],
	SUBMODE_8PSK250:        &lookupList[6],
	SUBMODE_8PSK250F:       &lookupList[7],
	SUBMODE_8PSK250FL:      &lookupList[8],
	SUBMODE_8PSK500:        &lookupList[9],
	SUBMODE_8PSK500F:       &lookupList[10],
	SUBMODE_AMTORFEC:       &lookupList[11],
	SUBMODE_ASCI:           &lookupList[12],
	SUBMODE_C4FM:           &lookupList[13],
	SUBMODE_CHIP128:        &lookupList[14],
	SUBMODE_CHIP64:         &lookupList[15],
	SUBMODE_DMR:            &lookupList[16],
	SUBMODE_DOM_M:          &lookupList[17],
	SUBMODE_DOM11:          &lookupList[18],
	SUBMODE_DOM16:          &lookupList[19],
	SUBMODE_DOM22:          &lookupList[20],
	SUBMODE_DOM4:           &lookupList[21],
	SUBMODE_DOM44:          &lookupList[22],
	SUBMODE_DOM5:           &lookupList[23],
	SUBMODE_DOM8:           &lookupList[24],
	SUBMODE_DOM88:          &lookupList[25],
	SUBMODE_DOMINOEX:       &lookupList[26],
	SUBMODE_DOMINOF:        &lookupList[27],
	SUBMODE_DSTAR:          &lookupList[28],
	SUBMODE_FMHELL:         &lookupList[29],
	SUBMODE_FREEDV:         &lookupList[30],
	SUBMODE_FSK31:          &lookupList[31],
	SUBMODE_FSKH105:        &lookupList[32],
	SUBMODE_FSKH245:        &lookupList[33],
	SUBMODE_FSKHELL:        &lookupList[34],
	SUBMODE_FSQCALL:        &lookupList[35],
	SUBMODE_FST4:           &lookupList[36],
	SUBMODE_FST4W:          &lookupList[37],
	SUBMODE_FT4:            &lookupList[38],
	SUBMODE_GTOR:           &lookupList[39],
	SUBMODE_HELL80:         &lookupList[40],
	SUBMODE_HELLX5:         &lookupList[41],
	SUBMODE_HELLX9:         &lookupList[42],
	SUBMODE_HFSK:           &lookupList[43],
	SUBMODE_ISCAT_A:        &lookupList[44],
	SUBMODE_ISCAT_B:        &lookupList[45],
	SUBMODE_JS8:            &lookupList[46],
	SUBMODE_JT4A:           &lookupList[47],
	SUBMODE_JT4B:           &lookupList[48],
	SUBMODE_JT4C:           &lookupList[49],
	SUBMODE_JT4D:           &lookupList[50],
	SUBMODE_JT4E:           &lookupList[51],
	SUBMODE_JT4F:           &lookupList[52],
	SUBMODE_JT4G:           &lookupList[53],
	SUBMODE_JT65A:          &lookupList[54],
	SUBMODE_JT65B:          &lookupList[55],
	SUBMODE_JT65B2:         &lookupList[56],
	SUBMODE_JT65C:          &lookupList[57],
	SUBMODE_JT65C2:         &lookupList[58],
	SUBMODE_JT9_1:          &lookupList[59],
	SUBMODE_JT9_10:         &lookupList[60],
	SUBMODE_JT9_2:          &lookupList[61],
	SUBMODE_JT9_30:         &lookupList[62],
	SUBMODE_JT9_5:          &lookupList[63],
	SUBMODE_JT9A:           &lookupList[64],
	SUBMODE_JT9B:           &lookupList[65],
	SUBMODE_JT9C:           &lookupList[66],
	SUBMODE_JT9D:           &lookupList[67],
	SUBMODE_JT9E:           &lookupList[68],
	SUBMODE_JT9E_FAST:      &lookupList[69],
	SUBMODE_JT9F:           &lookupList[70],
	SUBMODE_JT9F_FAST:      &lookupList[71],
	SUBMODE_JT9G:           &lookupList[72],
	SUBMODE_JT9G_FAST:      &lookupList[73],
	SUBMODE_JT9H:           &lookupList[74],
	SUBMODE_JT9H_FAST:      &lookupList[75],
	SUBMODE_JTMS:           &lookupList[76],
	SUBMODE_LSB:            &lookupList[77],
	SUBMODE_M17:            &lookupList[78],
	SUBMODE_MFSK11:         &lookupList[79],
	SUBMODE_MFSK128:        &lookupList[80],
	SUBMODE_MFSK128L:       &lookupList[81],
	SUBMODE_MFSK16:         &lookupList[82],
	SUBMODE_MFSK22:         &lookupList[83],
	SUBMODE_MFSK31:         &lookupList[84],
	SUBMODE_MFSK32:         &lookupList[85],
	SUBMODE_MFSK4:          &lookupList[86],
	SUBMODE_MFSK64:         &lookupList[87],
	SUBMODE_MFSK64L:        &lookupList[88],
	SUBMODE_MFSK8:          &lookupList[89],
	SUBMODE_NAVTEX:         &lookupList[90],
	SUBMODE_OLIVIA_16_1000: &lookupList[91],
	SUBMODE_OLIVIA_16_500:  &lookupList[92],
	SUBMODE_OLIVIA_32_1000: &lookupList[93],
	SUBMODE_OLIVIA_4_125:   &lookupList[94],
	SUBMODE_OLIVIA_4_250:   &lookupList[95],
	SUBMODE_OLIVIA_8_250:   &lookupList[96],
	SUBMODE_OLIVIA_8_500:   &lookupList[97],
	SUBMODE_OPERA_BEACON:   &lookupList[98],
	SUBMODE_OPERA_QSO:      &lookupList[99],
	SUBMODE_PAC2:           &lookupList[100],
	SUBMODE_PAC3:           &lookupList[101],
	SUBMODE_PAC4:           &lookupList[102],
	SUBMODE_PAX2:           &lookupList[103],
	SUBMODE_PCW:            &lookupList[104],
	SUBMODE_PSK10:          &lookupList[105],
	SUBMODE_PSK1000:        &lookupList[106],
	SUBMODE_PSK1000RC2:     &lookupList[107],
	SUBMODE_PSK125:         &lookupList[108],
	SUBMODE_PSK125RC10:     &lookupList[109],
	SUBMODE_PSK125RC12:     &lookupList[110],
	SUBMODE_PSK125RC16:     &lookupList[111],
	SUBMODE_PSK125RC4:      &lookupList[112],
	SUBMODE_PSK125RC5:      &lookupList[113],
	SUBMODE_PSK250:         &lookupList[114],
	SUBMODE_PSK250RC2:      &lookupList[115],
	SUBMODE_PSK250RC3:      &lookupList[116],
	SUBMODE_PSK250RC5:      &lookupList[117],
	SUBMODE_PSK250RC6:      &lookupList[118],
	SUBMODE_PSK250RC7:      &lookupList[119],
	SUBMODE_PSK31:          &lookupList[120],
	SUBMODE_PSK500:         &lookupList[121],
	SUBMODE_PSK500RC2:      &lookupList[122],
	SUBMODE_PSK500RC3:      &lookupList[123],
	SUBMODE_PSK500RC4:      &lookupList[124],
	SUBMODE_PSK63:          &lookupList[125],
	SUBMODE_PSK63F:         &lookupList[126],
	SUBMODE_PSK63RC10:      &lookupList[127],
	SUBMODE_PSK63RC20:      &lookupList[128],
	SUBMODE_PSK63RC32:      &lookupList[129],
	SUBMODE_PSK63RC4:       &lookupList[130],
	SUBMODE_PSK63RC5:       &lookupList[131],
	SUBMODE_PSK800RC2:      &lookupList[132],
	SUBMODE_PSKAM10:        &lookupList[133],
	SUBMODE_PSKAM31:        &lookupList[134],
	SUBMODE_PSKAM50:        &lookupList[135],
	SUBMODE_PSKFEC31:       &lookupList[136],
	SUBMODE_PSKHELL:        &lookupList[137],
	SUBMODE_Q65:            &lookupList[138],
	SUBMODE_QPSK125:        &lookupList[139],
	SUBMODE_QPSK250:        &lookupList[140],
	SUBMODE_QPSK31:         &lookupList[141],
	SUBMODE_QPSK500:        &lookupList[142],
	SUBMODE_QPSK63:         &lookupList[143],
	SUBMODE_QRA64A:         &lookupList[144],
	SUBMODE_QRA64B:         &lookupList[145],
	SUBMODE_QRA64C:         &lookupList[146],
	SUBMODE_QRA64D:         &lookupList[147],
	SUBMODE_QRA64E:         &lookupList[148],
	SUBMODE_ROS_EME:        &lookupList[149],
	SUBMODE_ROS_HF:         &lookupList[150],
	SUBMODE_ROS_MF:         &lookupList[151],
	SUBMODE_SCAMP_FAST:     &lookupList[152],
	SUBMODE_SCAMP_OO:       &lookupList[153],
	SUBMODE_SCAMP_OO_SLW:   &lookupList[154],
	SUBMODE_SCAMP_SLOW:     &lookupList[155],
	SUBMODE_SCAMP_VSLOW:    &lookupList[156],
	SUBMODE_SIM31:          &lookupList[157],
	SUBMODE_SITORB:         &lookupList[158],
	SUBMODE_SLOWHELL:       &lookupList[159],
	SUBMODE_THOR_M:         &lookupList[160],
	SUBMODE_THOR100:        &lookupList[161],
	SUBMODE_THOR11:         &lookupList[162],
	SUBMODE_THOR16:         &lookupList[163],
	SUBMODE_THOR22:         &lookupList[164],
	SUBMODE_THOR25X4:       &lookupList[165],
	SUBMODE_THOR4:          &lookupList[166],
	SUBMODE_THOR5:          &lookupList[167],
	SUBMODE_THOR50X1:       &lookupList[168],
	SUBMODE_THOR50X2:       &lookupList[169],
	SUBMODE_THOR8:          &lookupList[170],
	SUBMODE_THRBX:          &lookupList[171],
	SUBMODE_THRBX1:         &lookupList[172],
	SUBMODE_THRBX2:         &lookupList[173],
	SUBMODE_THRBX4:         &lookupList[174],
	SUBMODE_THROB1:         &lookupList[175],
	SUBMODE_THROB2:         &lookupList[176],
	SUBMODE_THROB4:         &lookupList[177],
	SUBMODE_USB:            &lookupList[178],
	SUBMODE_VARA_FM_1200:   &lookupList[179],
	SUBMODE_VARA_FM_9600:   &lookupList[180],
	SUBMODE_VARA_HF:        &lookupList[181],
	SUBMODE_VARA_SATELLITE: &lookupList[182],
}

// Lookup returns the specification for the provided SubMode.
// ADIF 3.1.6
func Lookup(s SubMode) (Spec, bool) {
	spec, ok := lookupMap[s]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all SubMode specifications that match the provided filter function.
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

// List returns all SubMode specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns SubMode specifications.
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
