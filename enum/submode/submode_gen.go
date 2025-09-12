// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package submode provides code and constants as defined in ADIF 3.1.6 (Proposed)
package submode

import "sync"

const (
	SubMode8PSK1000       SubMode = "8PSK1000"       // 8PSK1000        = PSK
	SubMode8PSK1000F      SubMode = "8PSK1000F"      // 8PSK1000F       = PSK
	SubMode8PSK1200F      SubMode = "8PSK1200F"      // 8PSK1200F       = PSK
	SubMode8PSK125        SubMode = "8PSK125"        // 8PSK125         = PSK
	SubMode8PSK125F       SubMode = "8PSK125F"       // 8PSK125F        = PSK
	SubMode8PSK125FL      SubMode = "8PSK125FL"      // 8PSK125FL       = PSK
	SubMode8PSK250        SubMode = "8PSK250"        // 8PSK250         = PSK
	SubMode8PSK250F       SubMode = "8PSK250F"       // 8PSK250F        = PSK
	SubMode8PSK250FL      SubMode = "8PSK250FL"      // 8PSK250FL       = PSK
	SubMode8PSK500        SubMode = "8PSK500"        // 8PSK500         = PSK
	SubMode8PSK500F       SubMode = "8PSK500F"       // 8PSK500F        = PSK
	SubModeAMTORFEC       SubMode = "AMTORFEC"       // AMTORFEC        = TOR
	SubModeASCI           SubMode = "ASCI"           // ASCI            = RTTY
	SubModeC4FM           SubMode = "C4FM"           // C4FM            = DIGITALVOICE    C4FM 4-level FSK See the Propagation_Mode enumeration section for examples of representing C4FM voice transmissions.
	SubModeCHIP128        SubMode = "CHIP128"        // CHIP128         = CHIP
	SubModeCHIP64         SubMode = "CHIP64"         // CHIP64          = CHIP
	SubModeDMR            SubMode = "DMR"            // DMR             = DIGITALVOICE    Digital Mobile Radio See the Propagation_Mode enumeration section for examples of representing DMR voice transmissions.
	SubModeDOM_M          SubMode = "DOM-M"          // DOM-M           = DOMINO
	SubModeDOM11          SubMode = "DOM11"          // DOM11           = DOMINO
	SubModeDOM16          SubMode = "DOM16"          // DOM16           = DOMINO
	SubModeDOM22          SubMode = "DOM22"          // DOM22           = DOMINO
	SubModeDOM4           SubMode = "DOM4"           // DOM4            = DOMINO
	SubModeDOM44          SubMode = "DOM44"          // DOM44           = DOMINO
	SubModeDOM5           SubMode = "DOM5"           // DOM5            = DOMINO
	SubModeDOM8           SubMode = "DOM8"           // DOM8            = DOMINO
	SubModeDOM88          SubMode = "DOM88"          // DOM88           = DOMINO
	SubModeDOMINOEX       SubMode = "DOMINOEX"       // DOMINOEX        = DOMINO
	SubModeDOMINOF        SubMode = "DOMINOF"        // DOMINOF         = DOMINO
	SubModeDSTAR          SubMode = "DSTAR"          // DSTAR           = DIGITALVOICE    Digital Smart Technologies for Amateur Radio See the Propagation_Mode enumeration section for examples of representing DSTAR voice transmissions.
	SubModeFMHELL         SubMode = "FMHELL"         // FMHELL          = HELL
	SubModeFREEDV         SubMode = "FREEDV"         // FREEDV          = DIGITALVOICE    Digital voice mode for HF radio implemented with open source
	SubModeFSK31          SubMode = "FSK31"          // FSK31           = PSK
	SubModeFSKH105        SubMode = "FSKH105"        // FSKH105         = HELL
	SubModeFSKH245        SubMode = "FSKH245"        // FSKH245         = HELL
	SubModeFSKHELL        SubMode = "FSKHELL"        // FSKHELL         = HELL
	SubModeFSQCALL        SubMode = "FSQCALL"        // FSQCALL         = MFSK            FSQCall protocol used with FSQ (Fast Simple QSO) transmission mode
	SubModeFST4           SubMode = "FST4"           // FST4            = MFSK            This is a digital mode supported by the WSJT-X software
	SubModeFST4W          SubMode = "FST4W"          // FST4W           = MFSK            This is a digital mode supported by the WSJT-X software that is for quasi-beacon transmissions of WSPR-style messages
	SubModeFT4            SubMode = "FT4"            // FT4             = MFSK            FT4 is a digital mode designed specifically for radio contesting
	SubModeGTOR           SubMode = "GTOR"           // GTOR            = TOR
	SubModeHELL80         SubMode = "HELL80"         // HELL80          = HELL
	SubModeHELLX5         SubMode = "HELLX5"         // HELLX5          = HELL
	SubModeHELLX9         SubMode = "HELLX9"         // HELLX9          = HELL
	SubModeHFSK           SubMode = "HFSK"           // HFSK            = HELL
	SubModeISCAT_A        SubMode = "ISCAT-A"        // ISCAT-A         = ISCAT
	SubModeISCAT_B        SubMode = "ISCAT-B"        // ISCAT-B         = ISCAT
	SubModeJS8            SubMode = "JS8"            // JS8             = MFSK            Jordan Sherer designed 8-FSK modulation
	SubModeJT4A           SubMode = "JT4A"           // JT4A            = JT4
	SubModeJT4B           SubMode = "JT4B"           // JT4B            = JT4
	SubModeJT4C           SubMode = "JT4C"           // JT4C            = JT4
	SubModeJT4D           SubMode = "JT4D"           // JT4D            = JT4
	SubModeJT4E           SubMode = "JT4E"           // JT4E            = JT4
	SubModeJT4F           SubMode = "JT4F"           // JT4F            = JT4
	SubModeJT4G           SubMode = "JT4G"           // JT4G            = JT4
	SubModeJT65A          SubMode = "JT65A"          // JT65A           = JT65
	SubModeJT65B          SubMode = "JT65B"          // JT65B           = JT65
	SubModeJT65B2         SubMode = "JT65B2"         // JT65B2          = JT65
	SubModeJT65C          SubMode = "JT65C"          // JT65C           = JT65
	SubModeJT65C2         SubMode = "JT65C2"         // JT65C2          = JT65
	SubModeJT9_1          SubMode = "JT9-1"          // JT9-1           = JT9
	SubModeJT9_10         SubMode = "JT9-10"         // JT9-10          = JT9
	SubModeJT9_2          SubMode = "JT9-2"          // JT9-2           = JT9
	SubModeJT9_30         SubMode = "JT9-30"         // JT9-30          = JT9
	SubModeJT9_5          SubMode = "JT9-5"          // JT9-5           = JT9
	SubModeJT9A           SubMode = "JT9A"           // JT9A            = JT9
	SubModeJT9B           SubMode = "JT9B"           // JT9B            = JT9
	SubModeJT9C           SubMode = "JT9C"           // JT9C            = JT9
	SubModeJT9D           SubMode = "JT9D"           // JT9D            = JT9
	SubModeJT9E           SubMode = "JT9E"           // JT9E            = JT9
	SubModeJT9E_FAST      SubMode = "JT9E FAST"      // JT9E FAST       = JT9
	SubModeJT9F           SubMode = "JT9F"           // JT9F            = JT9
	SubModeJT9F_FAST      SubMode = "JT9F FAST"      // JT9F FAST       = JT9
	SubModeJT9G           SubMode = "JT9G"           // JT9G            = JT9
	SubModeJT9G_FAST      SubMode = "JT9G FAST"      // JT9G FAST       = JT9
	SubModeJT9H           SubMode = "JT9H"           // JT9H            = JT9
	SubModeJT9H_FAST      SubMode = "JT9H FAST"      // JT9H FAST       = JT9
	SubModeJTMS           SubMode = "JTMS"           // JTMS            = MFSK
	SubModeLSB            SubMode = "LSB"            // LSB             = SSB             Amplitude modulated voice telephony, lower-sideband, suppressed-carrier
	SubModeM17            SubMode = "M17"            // M17             = DIGITALVOICE    Digital radio protocol using the Codec 2 voice encoder
	SubModeMFSK11         SubMode = "MFSK11"         // MFSK11          = MFSK
	SubModeMFSK128        SubMode = "MFSK128"        // MFSK128         = MFSK
	SubModeMFSK128L       SubMode = "MFSK128L"       // MFSK128L        = MFSK
	SubModeMFSK16         SubMode = "MFSK16"         // MFSK16          = MFSK
	SubModeMFSK22         SubMode = "MFSK22"         // MFSK22          = MFSK
	SubModeMFSK31         SubMode = "MFSK31"         // MFSK31          = MFSK
	SubModeMFSK32         SubMode = "MFSK32"         // MFSK32          = MFSK
	SubModeMFSK4          SubMode = "MFSK4"          // MFSK4           = MFSK
	SubModeMFSK64         SubMode = "MFSK64"         // MFSK64          = MFSK
	SubModeMFSK64L        SubMode = "MFSK64L"        // MFSK64L         = MFSK
	SubModeMFSK8          SubMode = "MFSK8"          // MFSK8           = MFSK
	SubModeNAVTEX         SubMode = "NAVTEX"         // NAVTEX          = TOR
	SubModeOLIVIA_16_1000 SubMode = "OLIVIA 16/1000" // OLIVIA 16/1000  = OLIVIA
	SubModeOLIVIA_16_500  SubMode = "OLIVIA 16/500"  // OLIVIA 16/500   = OLIVIA
	SubModeOLIVIA_32_1000 SubMode = "OLIVIA 32/1000" // OLIVIA 32/1000  = OLIVIA
	SubModeOLIVIA_4_125   SubMode = "OLIVIA 4/125"   // OLIVIA 4/125    = OLIVIA
	SubModeOLIVIA_4_250   SubMode = "OLIVIA 4/250"   // OLIVIA 4/250    = OLIVIA
	SubModeOLIVIA_8_250   SubMode = "OLIVIA 8/250"   // OLIVIA 8/250    = OLIVIA
	SubModeOLIVIA_8_500   SubMode = "OLIVIA 8/500"   // OLIVIA 8/500    = OLIVIA
	SubModeOPERA_BEACON   SubMode = "OPERA-BEACON"   // OPERA-BEACON    = OPERA
	SubModeOPERA_QSO      SubMode = "OPERA-QSO"      // OPERA-QSO       = OPERA
	SubModePAC2           SubMode = "PAC2"           // PAC2            = PAC
	SubModePAC3           SubMode = "PAC3"           // PAC3            = PAC
	SubModePAC4           SubMode = "PAC4"           // PAC4            = PAC
	SubModePAX2           SubMode = "PAX2"           // PAX2            = PAX
	SubModePCW            SubMode = "PCW"            // PCW             = CW              Coherent CW
	SubModePSK10          SubMode = "PSK10"          // PSK10           = PSK
	SubModePSK1000        SubMode = "PSK1000"        // PSK1000         = PSK
	SubModePSK1000RC2     SubMode = "PSK1000RC2"     // PSK1000RC2      = PSK
	SubModePSK125         SubMode = "PSK125"         // PSK125          = PSK
	SubModePSK125RC10     SubMode = "PSK125RC10"     // PSK125RC10      = PSK
	SubModePSK125RC12     SubMode = "PSK125RC12"     // PSK125RC12      = PSK
	SubModePSK125RC16     SubMode = "PSK125RC16"     // PSK125RC16      = PSK
	SubModePSK125RC4      SubMode = "PSK125RC4"      // PSK125RC4       = PSK
	SubModePSK125RC5      SubMode = "PSK125RC5"      // PSK125RC5       = PSK
	SubModePSK250         SubMode = "PSK250"         // PSK250          = PSK
	SubModePSK250RC2      SubMode = "PSK250RC2"      // PSK250RC2       = PSK
	SubModePSK250RC3      SubMode = "PSK250RC3"      // PSK250RC3       = PSK
	SubModePSK250RC5      SubMode = "PSK250RC5"      // PSK250RC5       = PSK
	SubModePSK250RC6      SubMode = "PSK250RC6"      // PSK250RC6       = PSK
	SubModePSK250RC7      SubMode = "PSK250RC7"      // PSK250RC7       = PSK
	SubModePSK31          SubMode = "PSK31"          // PSK31           = PSK
	SubModePSK500         SubMode = "PSK500"         // PSK500          = PSK
	SubModePSK500RC2      SubMode = "PSK500RC2"      // PSK500RC2       = PSK
	SubModePSK500RC3      SubMode = "PSK500RC3"      // PSK500RC3       = PSK
	SubModePSK500RC4      SubMode = "PSK500RC4"      // PSK500RC4       = PSK
	SubModePSK63          SubMode = "PSK63"          // PSK63           = PSK
	SubModePSK63F         SubMode = "PSK63F"         // PSK63F          = PSK
	SubModePSK63RC10      SubMode = "PSK63RC10"      // PSK63RC10       = PSK
	SubModePSK63RC20      SubMode = "PSK63RC20"      // PSK63RC20       = PSK
	SubModePSK63RC32      SubMode = "PSK63RC32"      // PSK63RC32       = PSK
	SubModePSK63RC4       SubMode = "PSK63RC4"       // PSK63RC4        = PSK
	SubModePSK63RC5       SubMode = "PSK63RC5"       // PSK63RC5        = PSK
	SubModePSK800RC2      SubMode = "PSK800RC2"      // PSK800RC2       = PSK
	SubModePSKAM10        SubMode = "PSKAM10"        // PSKAM10         = PSK
	SubModePSKAM31        SubMode = "PSKAM31"        // PSKAM31         = PSK
	SubModePSKAM50        SubMode = "PSKAM50"        // PSKAM50         = PSK
	SubModePSKFEC31       SubMode = "PSKFEC31"       // PSKFEC31        = PSK
	SubModePSKHELL        SubMode = "PSKHELL"        // PSKHELL         = HELL
	SubModeQ65            SubMode = "Q65"            // Q65             = MFSK
	SubModeQPSK125        SubMode = "QPSK125"        // QPSK125         = PSK
	SubModeQPSK250        SubMode = "QPSK250"        // QPSK250         = PSK
	SubModeQPSK31         SubMode = "QPSK31"         // QPSK31          = PSK
	SubModeQPSK500        SubMode = "QPSK500"        // QPSK500         = PSK
	SubModeQPSK63         SubMode = "QPSK63"         // QPSK63          = PSK
	SubModeQRA64A         SubMode = "QRA64A"         // QRA64A          = QRA64
	SubModeQRA64B         SubMode = "QRA64B"         // QRA64B          = QRA64
	SubModeQRA64C         SubMode = "QRA64C"         // QRA64C          = QRA64
	SubModeQRA64D         SubMode = "QRA64D"         // QRA64D          = QRA64
	SubModeQRA64E         SubMode = "QRA64E"         // QRA64E          = QRA64
	SubModeROS_EME        SubMode = "ROS-EME"        // ROS-EME         = ROS
	SubModeROS_HF         SubMode = "ROS-HF"         // ROS-HF          = ROS
	SubModeROS_MF         SubMode = "ROS-MF"         // ROS-MF          = ROS
	SubModeSCAMP_FAST     SubMode = "SCAMP_FAST"     // SCAMP_FAST      = FSK             SCAMP fast FSK
	SubModeSCAMP_OO       SubMode = "SCAMP_OO"       // SCAMP_OO        = MTONE           SCAMP single modulated tone on/off keying
	SubModeSCAMP_OO_SLW   SubMode = "SCAMP_OO_SLW"   // SCAMP_OO_SLW    = MTONE           SCAMP single modulated tone on/off slow keying
	SubModeSCAMP_SLOW     SubMode = "SCAMP_SLOW"     // SCAMP_SLOW      = FSK             SCAMP slow FSK
	SubModeSCAMP_VSLOW    SubMode = "SCAMP_VSLOW"    // SCAMP_VSLOW     = FSK             SCAMP very slow FSK
	SubModeSIM31          SubMode = "SIM31"          // SIM31           = PSK
	SubModeSITORB         SubMode = "SITORB"         // SITORB          = TOR
	SubModeSLOWHELL       SubMode = "SLOWHELL"       // SLOWHELL        = HELL
	SubModeTHOR_M         SubMode = "THOR-M"         // THOR-M          = THOR
	SubModeTHOR100        SubMode = "THOR100"        // THOR100         = THOR
	SubModeTHOR11         SubMode = "THOR11"         // THOR11          = THOR
	SubModeTHOR16         SubMode = "THOR16"         // THOR16          = THOR
	SubModeTHOR22         SubMode = "THOR22"         // THOR22          = THOR
	SubModeTHOR25X4       SubMode = "THOR25X4"       // THOR25X4        = THOR
	SubModeTHOR4          SubMode = "THOR4"          // THOR4           = THOR
	SubModeTHOR5          SubMode = "THOR5"          // THOR5           = THOR
	SubModeTHOR50X1       SubMode = "THOR50X1"       // THOR50X1        = THOR
	SubModeTHOR50X2       SubMode = "THOR50X2"       // THOR50X2        = THOR
	SubModeTHOR8          SubMode = "THOR8"          // THOR8           = THOR
	SubModeTHRBX          SubMode = "THRBX"          // THRBX           = THRB
	SubModeTHRBX1         SubMode = "THRBX1"         // THRBX1          = THRB
	SubModeTHRBX2         SubMode = "THRBX2"         // THRBX2          = THRB
	SubModeTHRBX4         SubMode = "THRBX4"         // THRBX4          = THRB
	SubModeTHROB1         SubMode = "THROB1"         // THROB1          = THRB
	SubModeTHROB2         SubMode = "THROB2"         // THROB2          = THRB
	SubModeTHROB4         SubMode = "THROB4"         // THROB4          = THRB
	SubModeUSB            SubMode = "USB"            // USB             = SSB             Amplitude modulated voice telephony, upper-sideband, suppressed-carrier
	SubModeVARA_FM_1200   SubMode = "VARA FM 1200"   // VARA FM 1200    = DYNAMIC         Channel adaptive high-speed modem for FM transceivers
	SubModeVARA_FM_9600   SubMode = "VARA FM 9600"   // VARA FM 9600    = DYNAMIC         Channel adaptive high-speed modem for FM transceivers
	SubModeVARA_HF        SubMode = "VARA HF"        // VARA HF         = DYNAMIC         Channel adaptive high-speed modem for HF
	SubModeVARA_SATELLITE SubMode = "VARA SATELLITE" // VARA SATELLITE  = DYNAMIC         Channel adaptive high-speed modem for satellite operations
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
	SubMode8PSK1000:       &lookupList[0],
	SubMode8PSK1000F:      &lookupList[1],
	SubMode8PSK1200F:      &lookupList[2],
	SubMode8PSK125:        &lookupList[3],
	SubMode8PSK125F:       &lookupList[4],
	SubMode8PSK125FL:      &lookupList[5],
	SubMode8PSK250:        &lookupList[6],
	SubMode8PSK250F:       &lookupList[7],
	SubMode8PSK250FL:      &lookupList[8],
	SubMode8PSK500:        &lookupList[9],
	SubMode8PSK500F:       &lookupList[10],
	SubModeAMTORFEC:       &lookupList[11],
	SubModeASCI:           &lookupList[12],
	SubModeC4FM:           &lookupList[13],
	SubModeCHIP128:        &lookupList[14],
	SubModeCHIP64:         &lookupList[15],
	SubModeDMR:            &lookupList[16],
	SubModeDOM_M:          &lookupList[17],
	SubModeDOM11:          &lookupList[18],
	SubModeDOM16:          &lookupList[19],
	SubModeDOM22:          &lookupList[20],
	SubModeDOM4:           &lookupList[21],
	SubModeDOM44:          &lookupList[22],
	SubModeDOM5:           &lookupList[23],
	SubModeDOM8:           &lookupList[24],
	SubModeDOM88:          &lookupList[25],
	SubModeDOMINOEX:       &lookupList[26],
	SubModeDOMINOF:        &lookupList[27],
	SubModeDSTAR:          &lookupList[28],
	SubModeFMHELL:         &lookupList[29],
	SubModeFREEDV:         &lookupList[30],
	SubModeFSK31:          &lookupList[31],
	SubModeFSKH105:        &lookupList[32],
	SubModeFSKH245:        &lookupList[33],
	SubModeFSKHELL:        &lookupList[34],
	SubModeFSQCALL:        &lookupList[35],
	SubModeFST4:           &lookupList[36],
	SubModeFST4W:          &lookupList[37],
	SubModeFT4:            &lookupList[38],
	SubModeGTOR:           &lookupList[39],
	SubModeHELL80:         &lookupList[40],
	SubModeHELLX5:         &lookupList[41],
	SubModeHELLX9:         &lookupList[42],
	SubModeHFSK:           &lookupList[43],
	SubModeISCAT_A:        &lookupList[44],
	SubModeISCAT_B:        &lookupList[45],
	SubModeJS8:            &lookupList[46],
	SubModeJT4A:           &lookupList[47],
	SubModeJT4B:           &lookupList[48],
	SubModeJT4C:           &lookupList[49],
	SubModeJT4D:           &lookupList[50],
	SubModeJT4E:           &lookupList[51],
	SubModeJT4F:           &lookupList[52],
	SubModeJT4G:           &lookupList[53],
	SubModeJT65A:          &lookupList[54],
	SubModeJT65B:          &lookupList[55],
	SubModeJT65B2:         &lookupList[56],
	SubModeJT65C:          &lookupList[57],
	SubModeJT65C2:         &lookupList[58],
	SubModeJT9_1:          &lookupList[59],
	SubModeJT9_10:         &lookupList[60],
	SubModeJT9_2:          &lookupList[61],
	SubModeJT9_30:         &lookupList[62],
	SubModeJT9_5:          &lookupList[63],
	SubModeJT9A:           &lookupList[64],
	SubModeJT9B:           &lookupList[65],
	SubModeJT9C:           &lookupList[66],
	SubModeJT9D:           &lookupList[67],
	SubModeJT9E:           &lookupList[68],
	SubModeJT9E_FAST:      &lookupList[69],
	SubModeJT9F:           &lookupList[70],
	SubModeJT9F_FAST:      &lookupList[71],
	SubModeJT9G:           &lookupList[72],
	SubModeJT9G_FAST:      &lookupList[73],
	SubModeJT9H:           &lookupList[74],
	SubModeJT9H_FAST:      &lookupList[75],
	SubModeJTMS:           &lookupList[76],
	SubModeLSB:            &lookupList[77],
	SubModeM17:            &lookupList[78],
	SubModeMFSK11:         &lookupList[79],
	SubModeMFSK128:        &lookupList[80],
	SubModeMFSK128L:       &lookupList[81],
	SubModeMFSK16:         &lookupList[82],
	SubModeMFSK22:         &lookupList[83],
	SubModeMFSK31:         &lookupList[84],
	SubModeMFSK32:         &lookupList[85],
	SubModeMFSK4:          &lookupList[86],
	SubModeMFSK64:         &lookupList[87],
	SubModeMFSK64L:        &lookupList[88],
	SubModeMFSK8:          &lookupList[89],
	SubModeNAVTEX:         &lookupList[90],
	SubModeOLIVIA_16_1000: &lookupList[91],
	SubModeOLIVIA_16_500:  &lookupList[92],
	SubModeOLIVIA_32_1000: &lookupList[93],
	SubModeOLIVIA_4_125:   &lookupList[94],
	SubModeOLIVIA_4_250:   &lookupList[95],
	SubModeOLIVIA_8_250:   &lookupList[96],
	SubModeOLIVIA_8_500:   &lookupList[97],
	SubModeOPERA_BEACON:   &lookupList[98],
	SubModeOPERA_QSO:      &lookupList[99],
	SubModePAC2:           &lookupList[100],
	SubModePAC3:           &lookupList[101],
	SubModePAC4:           &lookupList[102],
	SubModePAX2:           &lookupList[103],
	SubModePCW:            &lookupList[104],
	SubModePSK10:          &lookupList[105],
	SubModePSK1000:        &lookupList[106],
	SubModePSK1000RC2:     &lookupList[107],
	SubModePSK125:         &lookupList[108],
	SubModePSK125RC10:     &lookupList[109],
	SubModePSK125RC12:     &lookupList[110],
	SubModePSK125RC16:     &lookupList[111],
	SubModePSK125RC4:      &lookupList[112],
	SubModePSK125RC5:      &lookupList[113],
	SubModePSK250:         &lookupList[114],
	SubModePSK250RC2:      &lookupList[115],
	SubModePSK250RC3:      &lookupList[116],
	SubModePSK250RC5:      &lookupList[117],
	SubModePSK250RC6:      &lookupList[118],
	SubModePSK250RC7:      &lookupList[119],
	SubModePSK31:          &lookupList[120],
	SubModePSK500:         &lookupList[121],
	SubModePSK500RC2:      &lookupList[122],
	SubModePSK500RC3:      &lookupList[123],
	SubModePSK500RC4:      &lookupList[124],
	SubModePSK63:          &lookupList[125],
	SubModePSK63F:         &lookupList[126],
	SubModePSK63RC10:      &lookupList[127],
	SubModePSK63RC20:      &lookupList[128],
	SubModePSK63RC32:      &lookupList[129],
	SubModePSK63RC4:       &lookupList[130],
	SubModePSK63RC5:       &lookupList[131],
	SubModePSK800RC2:      &lookupList[132],
	SubModePSKAM10:        &lookupList[133],
	SubModePSKAM31:        &lookupList[134],
	SubModePSKAM50:        &lookupList[135],
	SubModePSKFEC31:       &lookupList[136],
	SubModePSKHELL:        &lookupList[137],
	SubModeQ65:            &lookupList[138],
	SubModeQPSK125:        &lookupList[139],
	SubModeQPSK250:        &lookupList[140],
	SubModeQPSK31:         &lookupList[141],
	SubModeQPSK500:        &lookupList[142],
	SubModeQPSK63:         &lookupList[143],
	SubModeQRA64A:         &lookupList[144],
	SubModeQRA64B:         &lookupList[145],
	SubModeQRA64C:         &lookupList[146],
	SubModeQRA64D:         &lookupList[147],
	SubModeQRA64E:         &lookupList[148],
	SubModeROS_EME:        &lookupList[149],
	SubModeROS_HF:         &lookupList[150],
	SubModeROS_MF:         &lookupList[151],
	SubModeSCAMP_FAST:     &lookupList[152],
	SubModeSCAMP_OO:       &lookupList[153],
	SubModeSCAMP_OO_SLW:   &lookupList[154],
	SubModeSCAMP_SLOW:     &lookupList[155],
	SubModeSCAMP_VSLOW:    &lookupList[156],
	SubModeSIM31:          &lookupList[157],
	SubModeSITORB:         &lookupList[158],
	SubModeSLOWHELL:       &lookupList[159],
	SubModeTHOR_M:         &lookupList[160],
	SubModeTHOR100:        &lookupList[161],
	SubModeTHOR11:         &lookupList[162],
	SubModeTHOR16:         &lookupList[163],
	SubModeTHOR22:         &lookupList[164],
	SubModeTHOR25X4:       &lookupList[165],
	SubModeTHOR4:          &lookupList[166],
	SubModeTHOR5:          &lookupList[167],
	SubModeTHOR50X1:       &lookupList[168],
	SubModeTHOR50X2:       &lookupList[169],
	SubModeTHOR8:          &lookupList[170],
	SubModeTHRBX:          &lookupList[171],
	SubModeTHRBX1:         &lookupList[172],
	SubModeTHRBX2:         &lookupList[173],
	SubModeTHRBX4:         &lookupList[174],
	SubModeTHROB1:         &lookupList[175],
	SubModeTHROB2:         &lookupList[176],
	SubModeTHROB4:         &lookupList[177],
	SubModeUSB:            &lookupList[178],
	SubModeVARA_FM_1200:   &lookupList[179],
	SubModeVARA_FM_9600:   &lookupList[180],
	SubModeVARA_HF:        &lookupList[181],
	SubModeVARA_SATELLITE: &lookupList[182],
}

// Lookup returns the specification for the provided SubMode.
// ADIF 3.1.6
func Lookup(submode SubMode) (Spec, bool) {
	spec, ok := lookupMap[submode]
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
