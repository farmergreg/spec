// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package submode provides code and constants as defined in ADIF 3.1.6 (Proposed)
package submode

import "sync"

const (
	SUBMODE_8PSK1000       SubMode = "8PSK1000"       // 8PSK1000        = PSK
	SUBMODE_8PSK1000F      SubMode = "8PSK1000F"      // 8PSK1000F       = PSK
	SUBMODE_8PSK1200F      SubMode = "8PSK1200F"      // 8PSK1200F       = PSK
	SUBMODE_8PSK125        SubMode = "8PSK125"        // 8PSK125         = PSK
	SUBMODE_8PSK125F       SubMode = "8PSK125F"       // 8PSK125F        = PSK
	SUBMODE_8PSK125FL      SubMode = "8PSK125FL"      // 8PSK125FL       = PSK
	SUBMODE_8PSK250        SubMode = "8PSK250"        // 8PSK250         = PSK
	SUBMODE_8PSK250F       SubMode = "8PSK250F"       // 8PSK250F        = PSK
	SUBMODE_8PSK250FL      SubMode = "8PSK250FL"      // 8PSK250FL       = PSK
	SUBMODE_8PSK500        SubMode = "8PSK500"        // 8PSK500         = PSK
	SUBMODE_8PSK500F       SubMode = "8PSK500F"       // 8PSK500F        = PSK
	SUBMODE_AMTORFEC       SubMode = "AMTORFEC"       // AMTORFEC        = TOR
	SUBMODE_ASCI           SubMode = "ASCI"           // ASCI            = RTTY
	SUBMODE_C4FM           SubMode = "C4FM"           // C4FM            = DIGITALVOICE    C4FM 4-level FSK See the Propagation_Mode enumeration section for examples of representing C4FM voice transmissions.
	SUBMODE_CHIP128        SubMode = "CHIP128"        // CHIP128         = CHIP
	SUBMODE_CHIP64         SubMode = "CHIP64"         // CHIP64          = CHIP
	SUBMODE_DMR            SubMode = "DMR"            // DMR             = DIGITALVOICE    Digital Mobile Radio See the Propagation_Mode enumeration section for examples of representing DMR voice transmissions.
	SUBMODE_DOM_M          SubMode = "DOM-M"          // DOM-M           = DOMINO
	SUBMODE_DOM11          SubMode = "DOM11"          // DOM11           = DOMINO
	SUBMODE_DOM16          SubMode = "DOM16"          // DOM16           = DOMINO
	SUBMODE_DOM22          SubMode = "DOM22"          // DOM22           = DOMINO
	SUBMODE_DOM4           SubMode = "DOM4"           // DOM4            = DOMINO
	SUBMODE_DOM44          SubMode = "DOM44"          // DOM44           = DOMINO
	SUBMODE_DOM5           SubMode = "DOM5"           // DOM5            = DOMINO
	SUBMODE_DOM8           SubMode = "DOM8"           // DOM8            = DOMINO
	SUBMODE_DOM88          SubMode = "DOM88"          // DOM88           = DOMINO
	SUBMODE_DOMINOEX       SubMode = "DOMINOEX"       // DOMINOEX        = DOMINO
	SUBMODE_DOMINOF        SubMode = "DOMINOF"        // DOMINOF         = DOMINO
	SUBMODE_DSTAR          SubMode = "DSTAR"          // DSTAR           = DIGITALVOICE    Digital Smart Technologies for Amateur Radio See the Propagation_Mode enumeration section for examples of representing DSTAR voice transmissions.
	SUBMODE_FMHELL         SubMode = "FMHELL"         // FMHELL          = HELL
	SUBMODE_FREEDV         SubMode = "FREEDV"         // FREEDV          = DIGITALVOICE    Digital voice mode for HF radio implemented with open source
	SUBMODE_FSK31          SubMode = "FSK31"          // FSK31           = PSK
	SUBMODE_FSKH105        SubMode = "FSKH105"        // FSKH105         = HELL
	SUBMODE_FSKH245        SubMode = "FSKH245"        // FSKH245         = HELL
	SUBMODE_FSKHELL        SubMode = "FSKHELL"        // FSKHELL         = HELL
	SUBMODE_FSQCALL        SubMode = "FSQCALL"        // FSQCALL         = MFSK            FSQCall protocol used with FSQ (Fast Simple QSO) transmission mode
	SUBMODE_FST4           SubMode = "FST4"           // FST4            = MFSK            This is a digital mode supported by the WSJT-X software
	SUBMODE_FST4W          SubMode = "FST4W"          // FST4W           = MFSK            This is a digital mode supported by the WSJT-X software that is for quasi-beacon transmissions of WSPR-style messages
	SUBMODE_FT4            SubMode = "FT4"            // FT4             = MFSK            FT4 is a digital mode designed specifically for radio contesting
	SUBMODE_GTOR           SubMode = "GTOR"           // GTOR            = TOR
	SUBMODE_HELL80         SubMode = "HELL80"         // HELL80          = HELL
	SUBMODE_HELLX5         SubMode = "HELLX5"         // HELLX5          = HELL
	SUBMODE_HELLX9         SubMode = "HELLX9"         // HELLX9          = HELL
	SUBMODE_HFSK           SubMode = "HFSK"           // HFSK            = HELL
	SUBMODE_ISCAT_A        SubMode = "ISCAT-A"        // ISCAT-A         = ISCAT
	SUBMODE_ISCAT_B        SubMode = "ISCAT-B"        // ISCAT-B         = ISCAT
	SUBMODE_JS8            SubMode = "JS8"            // JS8             = MFSK            Jordan Sherer designed 8-FSK modulation
	SUBMODE_JT4A           SubMode = "JT4A"           // JT4A            = JT4
	SUBMODE_JT4B           SubMode = "JT4B"           // JT4B            = JT4
	SUBMODE_JT4C           SubMode = "JT4C"           // JT4C            = JT4
	SUBMODE_JT4D           SubMode = "JT4D"           // JT4D            = JT4
	SUBMODE_JT4E           SubMode = "JT4E"           // JT4E            = JT4
	SUBMODE_JT4F           SubMode = "JT4F"           // JT4F            = JT4
	SUBMODE_JT4G           SubMode = "JT4G"           // JT4G            = JT4
	SUBMODE_JT65A          SubMode = "JT65A"          // JT65A           = JT65
	SUBMODE_JT65B          SubMode = "JT65B"          // JT65B           = JT65
	SUBMODE_JT65B2         SubMode = "JT65B2"         // JT65B2          = JT65
	SUBMODE_JT65C          SubMode = "JT65C"          // JT65C           = JT65
	SUBMODE_JT65C2         SubMode = "JT65C2"         // JT65C2          = JT65
	SUBMODE_JT9_1          SubMode = "JT9-1"          // JT9-1           = JT9
	SUBMODE_JT9_10         SubMode = "JT9-10"         // JT9-10          = JT9
	SUBMODE_JT9_2          SubMode = "JT9-2"          // JT9-2           = JT9
	SUBMODE_JT9_30         SubMode = "JT9-30"         // JT9-30          = JT9
	SUBMODE_JT9_5          SubMode = "JT9-5"          // JT9-5           = JT9
	SUBMODE_JT9A           SubMode = "JT9A"           // JT9A            = JT9
	SUBMODE_JT9B           SubMode = "JT9B"           // JT9B            = JT9
	SUBMODE_JT9C           SubMode = "JT9C"           // JT9C            = JT9
	SUBMODE_JT9D           SubMode = "JT9D"           // JT9D            = JT9
	SUBMODE_JT9E           SubMode = "JT9E"           // JT9E            = JT9
	SUBMODE_JT9E_FAST      SubMode = "JT9E FAST"      // JT9E FAST       = JT9
	SUBMODE_JT9F           SubMode = "JT9F"           // JT9F            = JT9
	SUBMODE_JT9F_FAST      SubMode = "JT9F FAST"      // JT9F FAST       = JT9
	SUBMODE_JT9G           SubMode = "JT9G"           // JT9G            = JT9
	SUBMODE_JT9G_FAST      SubMode = "JT9G FAST"      // JT9G FAST       = JT9
	SUBMODE_JT9H           SubMode = "JT9H"           // JT9H            = JT9
	SUBMODE_JT9H_FAST      SubMode = "JT9H FAST"      // JT9H FAST       = JT9
	SUBMODE_JTMS           SubMode = "JTMS"           // JTMS            = MFSK
	SUBMODE_LSB            SubMode = "LSB"            // LSB             = SSB             Amplitude modulated voice telephony, lower-sideband, suppressed-carrier
	SUBMODE_M17            SubMode = "M17"            // M17             = DIGITALVOICE    Digital radio protocol using the Codec 2 voice encoder
	SUBMODE_MFSK11         SubMode = "MFSK11"         // MFSK11          = MFSK
	SUBMODE_MFSK128        SubMode = "MFSK128"        // MFSK128         = MFSK
	SUBMODE_MFSK128L       SubMode = "MFSK128L"       // MFSK128L        = MFSK
	SUBMODE_MFSK16         SubMode = "MFSK16"         // MFSK16          = MFSK
	SUBMODE_MFSK22         SubMode = "MFSK22"         // MFSK22          = MFSK
	SUBMODE_MFSK31         SubMode = "MFSK31"         // MFSK31          = MFSK
	SUBMODE_MFSK32         SubMode = "MFSK32"         // MFSK32          = MFSK
	SUBMODE_MFSK4          SubMode = "MFSK4"          // MFSK4           = MFSK
	SUBMODE_MFSK64         SubMode = "MFSK64"         // MFSK64          = MFSK
	SUBMODE_MFSK64L        SubMode = "MFSK64L"        // MFSK64L         = MFSK
	SUBMODE_MFSK8          SubMode = "MFSK8"          // MFSK8           = MFSK
	SUBMODE_NAVTEX         SubMode = "NAVTEX"         // NAVTEX          = TOR
	SUBMODE_OLIVIA_16_1000 SubMode = "OLIVIA 16/1000" // OLIVIA 16/1000  = OLIVIA
	SUBMODE_OLIVIA_16_500  SubMode = "OLIVIA 16/500"  // OLIVIA 16/500   = OLIVIA
	SUBMODE_OLIVIA_32_1000 SubMode = "OLIVIA 32/1000" // OLIVIA 32/1000  = OLIVIA
	SUBMODE_OLIVIA_4_125   SubMode = "OLIVIA 4/125"   // OLIVIA 4/125    = OLIVIA
	SUBMODE_OLIVIA_4_250   SubMode = "OLIVIA 4/250"   // OLIVIA 4/250    = OLIVIA
	SUBMODE_OLIVIA_8_250   SubMode = "OLIVIA 8/250"   // OLIVIA 8/250    = OLIVIA
	SUBMODE_OLIVIA_8_500   SubMode = "OLIVIA 8/500"   // OLIVIA 8/500    = OLIVIA
	SUBMODE_OPERA_BEACON   SubMode = "OPERA-BEACON"   // OPERA-BEACON    = OPERA
	SUBMODE_OPERA_QSO      SubMode = "OPERA-QSO"      // OPERA-QSO       = OPERA
	SUBMODE_PAC2           SubMode = "PAC2"           // PAC2            = PAC
	SUBMODE_PAC3           SubMode = "PAC3"           // PAC3            = PAC
	SUBMODE_PAC4           SubMode = "PAC4"           // PAC4            = PAC
	SUBMODE_PAX2           SubMode = "PAX2"           // PAX2            = PAX
	SUBMODE_PCW            SubMode = "PCW"            // PCW             = CW              Coherent CW
	SUBMODE_PSK10          SubMode = "PSK10"          // PSK10           = PSK
	SUBMODE_PSK1000        SubMode = "PSK1000"        // PSK1000         = PSK
	SUBMODE_PSK1000RC2     SubMode = "PSK1000RC2"     // PSK1000RC2      = PSK
	SUBMODE_PSK125         SubMode = "PSK125"         // PSK125          = PSK
	SUBMODE_PSK125RC10     SubMode = "PSK125RC10"     // PSK125RC10      = PSK
	SUBMODE_PSK125RC12     SubMode = "PSK125RC12"     // PSK125RC12      = PSK
	SUBMODE_PSK125RC16     SubMode = "PSK125RC16"     // PSK125RC16      = PSK
	SUBMODE_PSK125RC4      SubMode = "PSK125RC4"      // PSK125RC4       = PSK
	SUBMODE_PSK125RC5      SubMode = "PSK125RC5"      // PSK125RC5       = PSK
	SUBMODE_PSK250         SubMode = "PSK250"         // PSK250          = PSK
	SUBMODE_PSK250RC2      SubMode = "PSK250RC2"      // PSK250RC2       = PSK
	SUBMODE_PSK250RC3      SubMode = "PSK250RC3"      // PSK250RC3       = PSK
	SUBMODE_PSK250RC5      SubMode = "PSK250RC5"      // PSK250RC5       = PSK
	SUBMODE_PSK250RC6      SubMode = "PSK250RC6"      // PSK250RC6       = PSK
	SUBMODE_PSK250RC7      SubMode = "PSK250RC7"      // PSK250RC7       = PSK
	SUBMODE_PSK31          SubMode = "PSK31"          // PSK31           = PSK
	SUBMODE_PSK500         SubMode = "PSK500"         // PSK500          = PSK
	SUBMODE_PSK500RC2      SubMode = "PSK500RC2"      // PSK500RC2       = PSK
	SUBMODE_PSK500RC3      SubMode = "PSK500RC3"      // PSK500RC3       = PSK
	SUBMODE_PSK500RC4      SubMode = "PSK500RC4"      // PSK500RC4       = PSK
	SUBMODE_PSK63          SubMode = "PSK63"          // PSK63           = PSK
	SUBMODE_PSK63F         SubMode = "PSK63F"         // PSK63F          = PSK
	SUBMODE_PSK63RC10      SubMode = "PSK63RC10"      // PSK63RC10       = PSK
	SUBMODE_PSK63RC20      SubMode = "PSK63RC20"      // PSK63RC20       = PSK
	SUBMODE_PSK63RC32      SubMode = "PSK63RC32"      // PSK63RC32       = PSK
	SUBMODE_PSK63RC4       SubMode = "PSK63RC4"       // PSK63RC4        = PSK
	SUBMODE_PSK63RC5       SubMode = "PSK63RC5"       // PSK63RC5        = PSK
	SUBMODE_PSK800RC2      SubMode = "PSK800RC2"      // PSK800RC2       = PSK
	SUBMODE_PSKAM10        SubMode = "PSKAM10"        // PSKAM10         = PSK
	SUBMODE_PSKAM31        SubMode = "PSKAM31"        // PSKAM31         = PSK
	SUBMODE_PSKAM50        SubMode = "PSKAM50"        // PSKAM50         = PSK
	SUBMODE_PSKFEC31       SubMode = "PSKFEC31"       // PSKFEC31        = PSK
	SUBMODE_PSKHELL        SubMode = "PSKHELL"        // PSKHELL         = HELL
	SUBMODE_Q65            SubMode = "Q65"            // Q65             = MFSK
	SUBMODE_QPSK125        SubMode = "QPSK125"        // QPSK125         = PSK
	SUBMODE_QPSK250        SubMode = "QPSK250"        // QPSK250         = PSK
	SUBMODE_QPSK31         SubMode = "QPSK31"         // QPSK31          = PSK
	SUBMODE_QPSK500        SubMode = "QPSK500"        // QPSK500         = PSK
	SUBMODE_QPSK63         SubMode = "QPSK63"         // QPSK63          = PSK
	SUBMODE_QRA64A         SubMode = "QRA64A"         // QRA64A          = QRA64
	SUBMODE_QRA64B         SubMode = "QRA64B"         // QRA64B          = QRA64
	SUBMODE_QRA64C         SubMode = "QRA64C"         // QRA64C          = QRA64
	SUBMODE_QRA64D         SubMode = "QRA64D"         // QRA64D          = QRA64
	SUBMODE_QRA64E         SubMode = "QRA64E"         // QRA64E          = QRA64
	SUBMODE_ROS_EME        SubMode = "ROS-EME"        // ROS-EME         = ROS
	SUBMODE_ROS_HF         SubMode = "ROS-HF"         // ROS-HF          = ROS
	SUBMODE_ROS_MF         SubMode = "ROS-MF"         // ROS-MF          = ROS
	SUBMODE_SCAMP_FAST     SubMode = "SCAMP_FAST"     // SCAMP_FAST      = FSK             SCAMP fast FSK
	SUBMODE_SCAMP_OO       SubMode = "SCAMP_OO"       // SCAMP_OO        = MTONE           SCAMP single modulated tone on/off keying
	SUBMODE_SCAMP_OO_SLW   SubMode = "SCAMP_OO_SLW"   // SCAMP_OO_SLW    = MTONE           SCAMP single modulated tone on/off slow keying
	SUBMODE_SCAMP_SLOW     SubMode = "SCAMP_SLOW"     // SCAMP_SLOW      = FSK             SCAMP slow FSK
	SUBMODE_SCAMP_VSLOW    SubMode = "SCAMP_VSLOW"    // SCAMP_VSLOW     = FSK             SCAMP very slow FSK
	SUBMODE_SIM31          SubMode = "SIM31"          // SIM31           = PSK
	SUBMODE_SITORB         SubMode = "SITORB"         // SITORB          = TOR
	SUBMODE_SLOWHELL       SubMode = "SLOWHELL"       // SLOWHELL        = HELL
	SUBMODE_THOR_M         SubMode = "THOR-M"         // THOR-M          = THOR
	SUBMODE_THOR100        SubMode = "THOR100"        // THOR100         = THOR
	SUBMODE_THOR11         SubMode = "THOR11"         // THOR11          = THOR
	SUBMODE_THOR16         SubMode = "THOR16"         // THOR16          = THOR
	SUBMODE_THOR22         SubMode = "THOR22"         // THOR22          = THOR
	SUBMODE_THOR25X4       SubMode = "THOR25X4"       // THOR25X4        = THOR
	SUBMODE_THOR4          SubMode = "THOR4"          // THOR4           = THOR
	SUBMODE_THOR5          SubMode = "THOR5"          // THOR5           = THOR
	SUBMODE_THOR50X1       SubMode = "THOR50X1"       // THOR50X1        = THOR
	SUBMODE_THOR50X2       SubMode = "THOR50X2"       // THOR50X2        = THOR
	SUBMODE_THOR8          SubMode = "THOR8"          // THOR8           = THOR
	SUBMODE_THRBX          SubMode = "THRBX"          // THRBX           = THRB
	SUBMODE_THRBX1         SubMode = "THRBX1"         // THRBX1          = THRB
	SUBMODE_THRBX2         SubMode = "THRBX2"         // THRBX2          = THRB
	SUBMODE_THRBX4         SubMode = "THRBX4"         // THRBX4          = THRB
	SUBMODE_THROB1         SubMode = "THROB1"         // THROB1          = THRB
	SUBMODE_THROB2         SubMode = "THROB2"         // THROB2          = THRB
	SUBMODE_THROB4         SubMode = "THROB4"         // THROB4          = THRB
	SUBMODE_USB            SubMode = "USB"            // USB             = SSB             Amplitude modulated voice telephony, upper-sideband, suppressed-carrier
	SUBMODE_VARA_FM_1200   SubMode = "VARA FM 1200"   // VARA FM 1200    = DYNAMIC         Channel adaptive high-speed modem for FM transceivers
	SUBMODE_VARA_FM_9600   SubMode = "VARA FM 9600"   // VARA FM 9600    = DYNAMIC         Channel adaptive high-speed modem for FM transceivers
	SUBMODE_VARA_HF        SubMode = "VARA HF"        // VARA HF         = DYNAMIC         Channel adaptive high-speed modem for HF
	SUBMODE_VARA_SATELLITE SubMode = "VARA SATELLITE" // VARA SATELLITE  = DYNAMIC         Channel adaptive high-speed modem for satellite operations
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known SubMode specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "8PSK1000", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "8PSK1000F", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "8PSK1200F", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "8PSK125", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "8PSK125F", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "8PSK125FL", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "8PSK250", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "8PSK250F", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "8PSK250FL", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "8PSK500", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "8PSK500F", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "AMTORFEC", Mode: "TOR", Description: ""},
	{IsImportOnly: false, Key: "ASCI", Mode: "RTTY", Description: ""},
	{IsImportOnly: false, Key: "C4FM", Mode: "DIGITALVOICE", Description: "C4FM 4-level FSK See the Propagation_Mode enumeration section for examples of representing C4FM voice transmissions."},
	{IsImportOnly: false, Key: "CHIP128", Mode: "CHIP", Description: ""},
	{IsImportOnly: false, Key: "CHIP64", Mode: "CHIP", Description: ""},
	{IsImportOnly: false, Key: "DMR", Mode: "DIGITALVOICE", Description: "Digital Mobile Radio See the Propagation_Mode enumeration section for examples of representing DMR voice transmissions."},
	{IsImportOnly: false, Key: "DOM-M", Mode: "DOMINO", Description: ""},
	{IsImportOnly: false, Key: "DOM11", Mode: "DOMINO", Description: ""},
	{IsImportOnly: false, Key: "DOM16", Mode: "DOMINO", Description: ""},
	{IsImportOnly: false, Key: "DOM22", Mode: "DOMINO", Description: ""},
	{IsImportOnly: false, Key: "DOM4", Mode: "DOMINO", Description: ""},
	{IsImportOnly: false, Key: "DOM44", Mode: "DOMINO", Description: ""},
	{IsImportOnly: false, Key: "DOM5", Mode: "DOMINO", Description: ""},
	{IsImportOnly: false, Key: "DOM8", Mode: "DOMINO", Description: ""},
	{IsImportOnly: false, Key: "DOM88", Mode: "DOMINO", Description: ""},
	{IsImportOnly: false, Key: "DOMINOEX", Mode: "DOMINO", Description: ""},
	{IsImportOnly: false, Key: "DOMINOF", Mode: "DOMINO", Description: ""},
	{IsImportOnly: false, Key: "DSTAR", Mode: "DIGITALVOICE", Description: "Digital Smart Technologies for Amateur Radio See the Propagation_Mode enumeration section for examples of representing DSTAR voice transmissions."},
	{IsImportOnly: false, Key: "FMHELL", Mode: "HELL", Description: ""},
	{IsImportOnly: false, Key: "FREEDV", Mode: "DIGITALVOICE", Description: "Digital voice mode for HF radio implemented with open source"},
	{IsImportOnly: false, Key: "FSK31", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "FSKH105", Mode: "HELL", Description: ""},
	{IsImportOnly: false, Key: "FSKH245", Mode: "HELL", Description: ""},
	{IsImportOnly: false, Key: "FSKHELL", Mode: "HELL", Description: ""},
	{IsImportOnly: false, Key: "FSQCALL", Mode: "MFSK", Description: "FSQCall protocol used with FSQ (Fast Simple QSO) transmission mode"},
	{IsImportOnly: false, Key: "FST4", Mode: "MFSK", Description: "This is a digital mode supported by the WSJT-X software"},
	{IsImportOnly: false, Key: "FST4W", Mode: "MFSK", Description: "This is a digital mode supported by the WSJT-X software that is for quasi-beacon transmissions of WSPR-style messages"},
	{IsImportOnly: false, Key: "FT4", Mode: "MFSK", Description: "FT4 is a digital mode designed specifically for radio contesting"},
	{IsImportOnly: false, Key: "GTOR", Mode: "TOR", Description: ""},
	{IsImportOnly: false, Key: "HELL80", Mode: "HELL", Description: ""},
	{IsImportOnly: false, Key: "HELLX5", Mode: "HELL", Description: ""},
	{IsImportOnly: false, Key: "HELLX9", Mode: "HELL", Description: ""},
	{IsImportOnly: false, Key: "HFSK", Mode: "HELL", Description: ""},
	{IsImportOnly: false, Key: "ISCAT-A", Mode: "ISCAT", Description: ""},
	{IsImportOnly: false, Key: "ISCAT-B", Mode: "ISCAT", Description: ""},
	{IsImportOnly: false, Key: "JS8", Mode: "MFSK", Description: "Jordan Sherer designed 8-FSK modulation"},
	{IsImportOnly: false, Key: "JT4A", Mode: "JT4", Description: ""},
	{IsImportOnly: false, Key: "JT4B", Mode: "JT4", Description: ""},
	{IsImportOnly: false, Key: "JT4C", Mode: "JT4", Description: ""},
	{IsImportOnly: false, Key: "JT4D", Mode: "JT4", Description: ""},
	{IsImportOnly: false, Key: "JT4E", Mode: "JT4", Description: ""},
	{IsImportOnly: false, Key: "JT4F", Mode: "JT4", Description: ""},
	{IsImportOnly: false, Key: "JT4G", Mode: "JT4", Description: ""},
	{IsImportOnly: false, Key: "JT65A", Mode: "JT65", Description: ""},
	{IsImportOnly: false, Key: "JT65B", Mode: "JT65", Description: ""},
	{IsImportOnly: false, Key: "JT65B2", Mode: "JT65", Description: ""},
	{IsImportOnly: false, Key: "JT65C", Mode: "JT65", Description: ""},
	{IsImportOnly: false, Key: "JT65C2", Mode: "JT65", Description: ""},
	{IsImportOnly: false, Key: "JT9-1", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9-10", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9-2", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9-30", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9-5", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9A", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9B", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9C", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9D", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9E", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9E FAST", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9F", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9F FAST", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9G", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9G FAST", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9H", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JT9H FAST", Mode: "JT9", Description: ""},
	{IsImportOnly: false, Key: "JTMS", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "LSB", Mode: "SSB", Description: "Amplitude modulated voice telephony, lower-sideband, suppressed-carrier"},
	{IsImportOnly: false, Key: "M17", Mode: "DIGITALVOICE", Description: "Digital radio protocol using the Codec 2 voice encoder"},
	{IsImportOnly: false, Key: "MFSK11", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "MFSK128", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "MFSK128L", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "MFSK16", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "MFSK22", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "MFSK31", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "MFSK32", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "MFSK4", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "MFSK64", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "MFSK64L", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "MFSK8", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "NAVTEX", Mode: "TOR", Description: ""},
	{IsImportOnly: false, Key: "OLIVIA 16/1000", Mode: "OLIVIA", Description: ""},
	{IsImportOnly: false, Key: "OLIVIA 16/500", Mode: "OLIVIA", Description: ""},
	{IsImportOnly: false, Key: "OLIVIA 32/1000", Mode: "OLIVIA", Description: ""},
	{IsImportOnly: false, Key: "OLIVIA 4/125", Mode: "OLIVIA", Description: ""},
	{IsImportOnly: false, Key: "OLIVIA 4/250", Mode: "OLIVIA", Description: ""},
	{IsImportOnly: false, Key: "OLIVIA 8/250", Mode: "OLIVIA", Description: ""},
	{IsImportOnly: false, Key: "OLIVIA 8/500", Mode: "OLIVIA", Description: ""},
	{IsImportOnly: false, Key: "OPERA-BEACON", Mode: "OPERA", Description: ""},
	{IsImportOnly: false, Key: "OPERA-QSO", Mode: "OPERA", Description: ""},
	{IsImportOnly: false, Key: "PAC2", Mode: "PAC", Description: ""},
	{IsImportOnly: false, Key: "PAC3", Mode: "PAC", Description: ""},
	{IsImportOnly: false, Key: "PAC4", Mode: "PAC", Description: ""},
	{IsImportOnly: false, Key: "PAX2", Mode: "PAX", Description: ""},
	{IsImportOnly: false, Key: "PCW", Mode: "CW", Description: "Coherent CW"},
	{IsImportOnly: false, Key: "PSK10", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK1000", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK1000RC2", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK125", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK125RC10", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK125RC12", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK125RC16", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK125RC4", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK125RC5", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK250", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK250RC2", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK250RC3", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK250RC5", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK250RC6", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK250RC7", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK31", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK500", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK500RC2", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK500RC3", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK500RC4", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK63", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK63F", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK63RC10", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK63RC20", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK63RC32", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK63RC4", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK63RC5", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSK800RC2", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSKAM10", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSKAM31", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSKAM50", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSKFEC31", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "PSKHELL", Mode: "HELL", Description: ""},
	{IsImportOnly: false, Key: "Q65", Mode: "MFSK", Description: ""},
	{IsImportOnly: false, Key: "QPSK125", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "QPSK250", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "QPSK31", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "QPSK500", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "QPSK63", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "QRA64A", Mode: "QRA64", Description: ""},
	{IsImportOnly: false, Key: "QRA64B", Mode: "QRA64", Description: ""},
	{IsImportOnly: false, Key: "QRA64C", Mode: "QRA64", Description: ""},
	{IsImportOnly: false, Key: "QRA64D", Mode: "QRA64", Description: ""},
	{IsImportOnly: false, Key: "QRA64E", Mode: "QRA64", Description: ""},
	{IsImportOnly: false, Key: "ROS-EME", Mode: "ROS", Description: ""},
	{IsImportOnly: false, Key: "ROS-HF", Mode: "ROS", Description: ""},
	{IsImportOnly: false, Key: "ROS-MF", Mode: "ROS", Description: ""},
	{IsImportOnly: false, Key: "SCAMP_FAST", Mode: "FSK", Description: "SCAMP fast FSK"},
	{IsImportOnly: false, Key: "SCAMP_OO", Mode: "MTONE", Description: "SCAMP single modulated tone on/off keying"},
	{IsImportOnly: false, Key: "SCAMP_OO_SLW", Mode: "MTONE", Description: "SCAMP single modulated tone on/off slow keying"},
	{IsImportOnly: false, Key: "SCAMP_SLOW", Mode: "FSK", Description: "SCAMP slow FSK"},
	{IsImportOnly: false, Key: "SCAMP_VSLOW", Mode: "FSK", Description: "SCAMP very slow FSK"},
	{IsImportOnly: false, Key: "SIM31", Mode: "PSK", Description: ""},
	{IsImportOnly: false, Key: "SITORB", Mode: "TOR", Description: ""},
	{IsImportOnly: false, Key: "SLOWHELL", Mode: "HELL", Description: ""},
	{IsImportOnly: false, Key: "THOR-M", Mode: "THOR", Description: ""},
	{IsImportOnly: false, Key: "THOR100", Mode: "THOR", Description: ""},
	{IsImportOnly: false, Key: "THOR11", Mode: "THOR", Description: ""},
	{IsImportOnly: false, Key: "THOR16", Mode: "THOR", Description: ""},
	{IsImportOnly: false, Key: "THOR22", Mode: "THOR", Description: ""},
	{IsImportOnly: false, Key: "THOR25X4", Mode: "THOR", Description: ""},
	{IsImportOnly: false, Key: "THOR4", Mode: "THOR", Description: ""},
	{IsImportOnly: false, Key: "THOR5", Mode: "THOR", Description: ""},
	{IsImportOnly: false, Key: "THOR50X1", Mode: "THOR", Description: ""},
	{IsImportOnly: false, Key: "THOR50X2", Mode: "THOR", Description: ""},
	{IsImportOnly: false, Key: "THOR8", Mode: "THOR", Description: ""},
	{IsImportOnly: false, Key: "THRBX", Mode: "THRB", Description: ""},
	{IsImportOnly: false, Key: "THRBX1", Mode: "THRB", Description: ""},
	{IsImportOnly: false, Key: "THRBX2", Mode: "THRB", Description: ""},
	{IsImportOnly: false, Key: "THRBX4", Mode: "THRB", Description: ""},
	{IsImportOnly: false, Key: "THROB1", Mode: "THRB", Description: ""},
	{IsImportOnly: false, Key: "THROB2", Mode: "THRB", Description: ""},
	{IsImportOnly: false, Key: "THROB4", Mode: "THRB", Description: ""},
	{IsImportOnly: false, Key: "USB", Mode: "SSB", Description: "Amplitude modulated voice telephony, upper-sideband, suppressed-carrier"},
	{IsImportOnly: false, Key: "VARA FM 1200", Mode: "DYNAMIC", Description: "Channel adaptive high-speed modem for FM transceivers"},
	{IsImportOnly: false, Key: "VARA FM 9600", Mode: "DYNAMIC", Description: "Channel adaptive high-speed modem for FM transceivers"},
	{IsImportOnly: false, Key: "VARA HF", Mode: "DYNAMIC", Description: "Channel adaptive high-speed modem for HF"},
	{IsImportOnly: false, Key: "VARA SATELLITE", Mode: "DYNAMIC", Description: "Channel adaptive high-speed modem for satellite operations"},
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
