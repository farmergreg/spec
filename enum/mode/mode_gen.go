// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package mode provides code and constants as defined in ADIF 3.1.6 (Proposed)
package mode

import (
	"maps"

	"github.com/hamradiolog-net/adif-spec/v6/enum/submode"
)

const (
	AM           Mode = "AM"           // AM                     = []
	AMTORFEC     Mode = "AMTORFEC"     // Deprecated: AMTORFEC   = []
	ARDOP        Mode = "ARDOP"        // ARDOP                  = []
	ASCI         Mode = "ASCI"         // Deprecated: ASCI       = []
	ATV          Mode = "ATV"          // ATV                    = []
	C4FM         Mode = "C4FM"         // Deprecated: C4FM       = []
	CHIP         Mode = "CHIP"         // CHIP                   = [CHIP64 CHIP128]
	CHIP128      Mode = "CHIP128"      // Deprecated: CHIP128    = []
	CHIP64       Mode = "CHIP64"       // Deprecated: CHIP64     = []
	CLO          Mode = "CLO"          // CLO                    = []
	CONTESTI     Mode = "CONTESTI"     // CONTESTI               = []
	CW           Mode = "CW"           // CW                     = [PCW]
	DIGITALVOICE Mode = "DIGITALVOICE" // DIGITALVOICE           = [C4FM DMR DSTAR FREEDV M17]
	DOMINO       Mode = "DOMINO"       // DOMINO                 = [DOM-M DOM4 DOM5 DOM8 DOM11 DOM16 DOM22 DOM44 DOM88 DOMINOEX DOMINOF]
	DOMINOF      Mode = "DOMINOF"      // Deprecated: DOMINOF    = []
	DSTAR        Mode = "DSTAR"        // Deprecated: DSTAR      = []
	DYNAMIC      Mode = "DYNAMIC"      // DYNAMIC                = [VARA HF VARA SATELLITE VARA FM 1200 VARA FM 9600]
	FAX          Mode = "FAX"          // FAX                    = []
	FM           Mode = "FM"           // FM                     = []
	FMHELL       Mode = "FMHELL"       // Deprecated: FMHELL     = []
	FSK          Mode = "FSK"          // FSK                    = [SCAMP_FAST SCAMP_SLOW SCAMP_VSLOW]
	FSK31        Mode = "FSK31"        // Deprecated: FSK31      = []
	FSK441       Mode = "FSK441"       // FSK441                 = []
	FT8          Mode = "FT8"          // FT8                    = []
	GTOR         Mode = "GTOR"         // Deprecated: GTOR       = []
	HELL         Mode = "HELL"         // HELL                   = [FMHELL FSKH105 FSKH245 FSKHELL HELL80 HELLX5 HELLX9 HFSK PSKHELL SLOWHELL]
	HELL80       Mode = "HELL80"       // Deprecated: HELL80     = []
	HFSK         Mode = "HFSK"         // Deprecated: HFSK       = []
	ISCAT        Mode = "ISCAT"        // ISCAT                  = [ISCAT-A ISCAT-B]
	JT4          Mode = "JT4"          // JT4                    = [JT4A JT4B JT4C JT4D JT4E JT4F JT4G]
	JT44         Mode = "JT44"         // JT44                   = []
	JT4A         Mode = "JT4A"         // Deprecated: JT4A       = []
	JT4B         Mode = "JT4B"         // Deprecated: JT4B       = []
	JT4C         Mode = "JT4C"         // Deprecated: JT4C       = []
	JT4D         Mode = "JT4D"         // Deprecated: JT4D       = []
	JT4E         Mode = "JT4E"         // Deprecated: JT4E       = []
	JT4F         Mode = "JT4F"         // Deprecated: JT4F       = []
	JT4G         Mode = "JT4G"         // Deprecated: JT4G       = []
	JT65         Mode = "JT65"         // JT65                   = [JT65A JT65B JT65B2 JT65C JT65C2]
	JT65A        Mode = "JT65A"        // Deprecated: JT65A      = []
	JT65B        Mode = "JT65B"        // Deprecated: JT65B      = []
	JT65C        Mode = "JT65C"        // Deprecated: JT65C      = []
	JT6M         Mode = "JT6M"         // JT6M                   = []
	JT9          Mode = "JT9"          // JT9                    = [JT9-1 JT9-2 JT9-5 JT9-10 JT9-30 JT9A JT9B JT9C JT9D JT9E JT9E FAST JT9F JT9F FAST JT9G JT9G FAST JT9H JT9H FAST]
	MFSK         Mode = "MFSK"         // MFSK                   = [FSQCALL FST4 FST4W FT4 JS8 JTMS MFSK4 MFSK8 MFSK11 MFSK16 MFSK22 MFSK31 MFSK32 MFSK64 MFSK64L MFSK128 MFSK128L Q65]
	MFSK16       Mode = "MFSK16"       // Deprecated: MFSK16     = []
	MFSK8        Mode = "MFSK8"        // Deprecated: MFSK8      = []
	MSK144       Mode = "MSK144"       // MSK144                 = []
	MT63         Mode = "MT63"         // MT63                   = []
	MTONE        Mode = "MTONE"        // MTONE                  = [SCAMP_OO SCAMP_OO_SLW]
	OLIVIA       Mode = "OLIVIA"       // OLIVIA                 = [OLIVIA 4/125 OLIVIA 4/250 OLIVIA 8/250 OLIVIA 8/500 OLIVIA 16/500 OLIVIA 16/1000 OLIVIA 32/1000]
	OPERA        Mode = "OPERA"        // OPERA                  = [OPERA-BEACON OPERA-QSO]
	PAC          Mode = "PAC"          // PAC                    = [PAC2 PAC3 PAC4]
	PAC2         Mode = "PAC2"         // Deprecated: PAC2       = []
	PAC3         Mode = "PAC3"         // Deprecated: PAC3       = []
	PAX          Mode = "PAX"          // PAX                    = [PAX2]
	PAX2         Mode = "PAX2"         // Deprecated: PAX2       = []
	PCW          Mode = "PCW"          // Deprecated: PCW        = []
	PKT          Mode = "PKT"          // PKT                    = []
	PSK          Mode = "PSK"          // PSK                    = [8PSK125 8PSK125F 8PSK125FL 8PSK250 8PSK250F 8PSK250FL 8PSK500 8PSK500F 8PSK1000 8PSK1000F 8PSK1200F FSK31 PSK10 PSK31 PSK63 PSK63F PSK63RC4 PSK63RC5 PSK63RC10 PSK63RC20 PSK63RC32 PSK125 PSK125C12 PSK125R PSK125RC10 PSK125RC12 PSK125RC16 PSK125RC4 PSK125RC5 PSK250 PSK250C6 PSK250R PSK250RC2 PSK250RC3 PSK250RC5 PSK250RC6 PSK250RC7 PSK500 PSK500C2 PSK500C4 PSK500R PSK500RC2 PSK500RC3 PSK500RC4 PSK800C2 PSK800RC2 PSK1000 PSK1000C2 PSK1000R PSK1000RC2 PSKAM10 PSKAM31 PSKAM50 PSKFEC31 QPSK31 QPSK63 QPSK125 QPSK250 QPSK500 SIM31]
	PSK10        Mode = "PSK10"        // Deprecated: PSK10      = []
	PSK125       Mode = "PSK125"       // Deprecated: PSK125     = []
	PSK2K        Mode = "PSK2K"        // PSK2K                  = []
	PSK31        Mode = "PSK31"        // Deprecated: PSK31      = []
	PSK63        Mode = "PSK63"        // Deprecated: PSK63      = []
	PSK63F       Mode = "PSK63F"       // Deprecated: PSK63F     = []
	PSKAM10      Mode = "PSKAM10"      // Deprecated: PSKAM10    = []
	PSKAM31      Mode = "PSKAM31"      // Deprecated: PSKAM31    = []
	PSKAM50      Mode = "PSKAM50"      // Deprecated: PSKAM50    = []
	PSKFEC31     Mode = "PSKFEC31"     // Deprecated: PSKFEC31   = []
	PSKHELL      Mode = "PSKHELL"      // Deprecated: PSKHELL    = []
	Q15          Mode = "Q15"          // Q15                    = []
	QPSK125      Mode = "QPSK125"      // Deprecated: QPSK125    = []
	QPSK31       Mode = "QPSK31"       // Deprecated: QPSK31     = []
	QPSK63       Mode = "QPSK63"       // Deprecated: QPSK63     = []
	QRA64        Mode = "QRA64"        // QRA64                  = [QRA64A QRA64B QRA64C QRA64D QRA64E]
	ROS          Mode = "ROS"          // ROS                    = [ROS-EME ROS-HF ROS-MF]
	RTTY         Mode = "RTTY"         // RTTY                   = [ASCI]
	RTTYM        Mode = "RTTYM"        // RTTYM                  = []
	SSB          Mode = "SSB"          // SSB                    = [LSB USB]
	SSTV         Mode = "SSTV"         // SSTV                   = []
	T10          Mode = "T10"          // T10                    = []
	THOR         Mode = "THOR"         // THOR                   = [THOR-M THOR4 THOR5 THOR8 THOR11 THOR16 THOR22 THOR25X4 THOR50X1 THOR50X2 THOR100]
	THRB         Mode = "THRB"         // THRB                   = [THRBX THRBX1 THRBX2 THRBX4 THROB1 THROB2 THROB4]
	THRBX        Mode = "THRBX"        // Deprecated: THRBX      = []
	TOR          Mode = "TOR"          // TOR                    = [AMTORFEC GTOR NAVTEX SITORB]
	V4           Mode = "V4"           // V4                     = []
	VOI          Mode = "VOI"          // VOI                    = []
	WINMOR       Mode = "WINMOR"       // WINMOR                 = []
	WSPR         Mode = "WSPR"         // WSPR                   = []
)

// All Mode specifications in ADIF 3.1.6 (Proposed) including depreciated and import only.
func ModeListAll() []Spec {
	return append([]Spec(nil), internalModeListAll...)
}

// All Mode specifications values in ADIF 3.1.6 (Proposed) that are NOT marked import-only.
func ModeListCurrent() []Spec {
	return append([]Spec(nil), internalModeListCurrent...)
}

// A map of all Mode from ADIF 3.1.6 (Proposed).
func ModeMap() map[Mode]Spec {
	cp := make(map[Mode]Spec, len(internalModeMap))
	maps.Copy(cp, internalModeMap)
	return cp
}

// A map of all Mode specifications.
var internalModeMap = map[Mode]Spec{
	AM:           {IsImportOnly: false, Key: "AM", Submodes: submode.SubModeList(nil), Description: ""},
	AMTORFEC:     {IsImportOnly: true, Key: "AMTORFEC", Submodes: submode.SubModeList(nil), Description: ""},
	ARDOP:        {IsImportOnly: false, Key: "ARDOP", Submodes: submode.SubModeList(nil), Description: "Amateur Radio Digital Open Protocol"},
	ASCI:         {IsImportOnly: true, Key: "ASCI", Submodes: submode.SubModeList(nil), Description: ""},
	ATV:          {IsImportOnly: false, Key: "ATV", Submodes: submode.SubModeList(nil), Description: ""},
	C4FM:         {IsImportOnly: true, Key: "C4FM", Submodes: submode.SubModeList(nil), Description: "C4FM 4-level FSK Technology Imported QSOs with <MODE:4>C4FM> must be exported as: <MODE:12>DIGITALVOICE <SUBMODE:4>C4FM"},
	CHIP:         {IsImportOnly: false, Key: "CHIP", Submodes: submode.SubModeList{"CHIP64", "CHIP128"}, Description: ""},
	CHIP128:      {IsImportOnly: true, Key: "CHIP128", Submodes: submode.SubModeList(nil), Description: ""},
	CHIP64:       {IsImportOnly: true, Key: "CHIP64", Submodes: submode.SubModeList(nil), Description: ""},
	CLO:          {IsImportOnly: false, Key: "CLO", Submodes: submode.SubModeList(nil), Description: ""},
	CONTESTI:     {IsImportOnly: false, Key: "CONTESTI", Submodes: submode.SubModeList(nil), Description: ""},
	CW:           {IsImportOnly: false, Key: "CW", Submodes: submode.SubModeList{"PCW"}, Description: ""},
	DIGITALVOICE: {IsImportOnly: false, Key: "DIGITALVOICE", Submodes: submode.SubModeList{"C4FM", "DMR", "DSTAR", "FREEDV", "M17"}, Description: ""},
	DOMINO:       {IsImportOnly: false, Key: "DOMINO", Submodes: submode.SubModeList{"DOM-M", "DOM4", "DOM5", "DOM8", "DOM11", "DOM16", "DOM22", "DOM44", "DOM88", "DOMINOEX", "DOMINOF"}, Description: ""},
	DOMINOF:      {IsImportOnly: true, Key: "DOMINOF", Submodes: submode.SubModeList(nil), Description: ""},
	DSTAR:        {IsImportOnly: true, Key: "DSTAR", Submodes: submode.SubModeList(nil), Description: "Digital Smart Technologies for Amateur Radio Imported QSOs with <MODE:5>DSTAR must be exported as: <MODE:12>DIGITALVOICE <SUBMODE:5>DSTAR"},
	DYNAMIC:      {IsImportOnly: false, Key: "DYNAMIC", Submodes: submode.SubModeList{"VARA HF", "VARA SATELLITE", "VARA FM 1200", "VARA FM 9600"}, Description: ""},
	FAX:          {IsImportOnly: false, Key: "FAX", Submodes: submode.SubModeList(nil), Description: ""},
	FM:           {IsImportOnly: false, Key: "FM", Submodes: submode.SubModeList(nil), Description: ""},
	FMHELL:       {IsImportOnly: true, Key: "FMHELL", Submodes: submode.SubModeList(nil), Description: ""},
	FSK:          {IsImportOnly: false, Key: "FSK", Submodes: submode.SubModeList{"SCAMP_FAST", "SCAMP_SLOW", "SCAMP_VSLOW"}, Description: "Frequency shift keying"},
	FSK31:        {IsImportOnly: true, Key: "FSK31", Submodes: submode.SubModeList(nil), Description: ""},
	FSK441:       {IsImportOnly: false, Key: "FSK441", Submodes: submode.SubModeList(nil), Description: ""},
	FT8:          {IsImportOnly: false, Key: "FT8", Submodes: submode.SubModeList(nil), Description: "Franke-Taylor design, 8-FSK modulation"},
	GTOR:         {IsImportOnly: true, Key: "GTOR", Submodes: submode.SubModeList(nil), Description: ""},
	HELL:         {IsImportOnly: false, Key: "HELL", Submodes: submode.SubModeList{"FMHELL", "FSKH105", "FSKH245", "FSKHELL", "HELL80", "HELLX5", "HELLX9", "HFSK", "PSKHELL", "SLOWHELL"}, Description: ""},
	HELL80:       {IsImportOnly: true, Key: "HELL80", Submodes: submode.SubModeList(nil), Description: ""},
	HFSK:         {IsImportOnly: true, Key: "HFSK", Submodes: submode.SubModeList(nil), Description: ""},
	ISCAT:        {IsImportOnly: false, Key: "ISCAT", Submodes: submode.SubModeList{"ISCAT-A", "ISCAT-B"}, Description: ""},
	JT4:          {IsImportOnly: false, Key: "JT4", Submodes: submode.SubModeList{"JT4A", "JT4B", "JT4C", "JT4D", "JT4E", "JT4F", "JT4G"}, Description: ""},
	JT44:         {IsImportOnly: false, Key: "JT44", Submodes: submode.SubModeList(nil), Description: ""},
	JT4A:         {IsImportOnly: true, Key: "JT4A", Submodes: submode.SubModeList(nil), Description: ""},
	JT4B:         {IsImportOnly: true, Key: "JT4B", Submodes: submode.SubModeList(nil), Description: ""},
	JT4C:         {IsImportOnly: true, Key: "JT4C", Submodes: submode.SubModeList(nil), Description: ""},
	JT4D:         {IsImportOnly: true, Key: "JT4D", Submodes: submode.SubModeList(nil), Description: ""},
	JT4E:         {IsImportOnly: true, Key: "JT4E", Submodes: submode.SubModeList(nil), Description: ""},
	JT4F:         {IsImportOnly: true, Key: "JT4F", Submodes: submode.SubModeList(nil), Description: ""},
	JT4G:         {IsImportOnly: true, Key: "JT4G", Submodes: submode.SubModeList(nil), Description: ""},
	JT65:         {IsImportOnly: false, Key: "JT65", Submodes: submode.SubModeList{"JT65A", "JT65B", "JT65B2", "JT65C", "JT65C2"}, Description: ""},
	JT65A:        {IsImportOnly: true, Key: "JT65A", Submodes: submode.SubModeList(nil), Description: ""},
	JT65B:        {IsImportOnly: true, Key: "JT65B", Submodes: submode.SubModeList(nil), Description: ""},
	JT65C:        {IsImportOnly: true, Key: "JT65C", Submodes: submode.SubModeList(nil), Description: ""},
	JT6M:         {IsImportOnly: false, Key: "JT6M", Submodes: submode.SubModeList(nil), Description: ""},
	JT9:          {IsImportOnly: false, Key: "JT9", Submodes: submode.SubModeList{"JT9-1", "JT9-2", "JT9-5", "JT9-10", "JT9-30", "JT9A", "JT9B", "JT9C", "JT9D", "JT9E", "JT9E FAST", "JT9F", "JT9F FAST", "JT9G", "JT9G FAST", "JT9H", "JT9H FAST"}, Description: ""},
	MFSK:         {IsImportOnly: false, Key: "MFSK", Submodes: submode.SubModeList{"FSQCALL", "FST4", "FST4W", "FT4", "JS8", "JTMS", "MFSK4", "MFSK8", "MFSK11", "MFSK16", "MFSK22", "MFSK31", "MFSK32", "MFSK64", "MFSK64L", "MFSK128 MFSK128L", "Q65"}, Description: ""},
	MFSK16:       {IsImportOnly: true, Key: "MFSK16", Submodes: submode.SubModeList(nil), Description: ""},
	MFSK8:        {IsImportOnly: true, Key: "MFSK8", Submodes: submode.SubModeList(nil), Description: ""},
	MSK144:       {IsImportOnly: false, Key: "MSK144", Submodes: submode.SubModeList(nil), Description: ""},
	MT63:         {IsImportOnly: false, Key: "MT63", Submodes: submode.SubModeList(nil), Description: ""},
	MTONE:        {IsImportOnly: false, Key: "MTONE", Submodes: submode.SubModeList{"SCAMP_OO", "SCAMP_OO_SLW"}, Description: "Single modulated tone"},
	OLIVIA:       {IsImportOnly: false, Key: "OLIVIA", Submodes: submode.SubModeList{"OLIVIA 4/125", "OLIVIA 4/250", "OLIVIA 8/250", "OLIVIA 8/500", "OLIVIA 16/500", "OLIVIA 16/1000", "OLIVIA 32/1000"}, Description: ""},
	OPERA:        {IsImportOnly: false, Key: "OPERA", Submodes: submode.SubModeList{"OPERA-BEACON", "OPERA-QSO"}, Description: ""},
	PAC:          {IsImportOnly: false, Key: "PAC", Submodes: submode.SubModeList{"PAC2", "PAC3", "PAC4"}, Description: ""},
	PAC2:         {IsImportOnly: true, Key: "PAC2", Submodes: submode.SubModeList(nil), Description: ""},
	PAC3:         {IsImportOnly: true, Key: "PAC3", Submodes: submode.SubModeList(nil), Description: ""},
	PAX:          {IsImportOnly: false, Key: "PAX", Submodes: submode.SubModeList{"PAX2"}, Description: ""},
	PAX2:         {IsImportOnly: true, Key: "PAX2", Submodes: submode.SubModeList(nil), Description: ""},
	PCW:          {IsImportOnly: true, Key: "PCW", Submodes: submode.SubModeList(nil), Description: ""},
	PKT:          {IsImportOnly: false, Key: "PKT", Submodes: submode.SubModeList(nil), Description: ""},
	PSK:          {IsImportOnly: false, Key: "PSK", Submodes: submode.SubModeList{"8PSK125", "8PSK125F", "8PSK125FL", "8PSK250", "8PSK250F", "8PSK250FL", "8PSK500", "8PSK500F", "8PSK1000", "8PSK1000F", "8PSK1200F", "FSK31", "PSK10", "PSK31", "PSK63", "PSK63F", "PSK63RC4", "PSK63RC5", "PSK63RC10", "PSK63RC20", "PSK63RC32", "PSK125", "PSK125C12", "PSK125R", "PSK125RC10", "PSK125RC12", "PSK125RC16", "PSK125RC4", "PSK125RC5", "PSK250", "PSK250C6", "PSK250R", "PSK250RC2", "PSK250RC3", "PSK250RC5", "PSK250RC6", "PSK250RC7", "PSK500", "PSK500C2", "PSK500C4", "PSK500R", "PSK500RC2", "PSK500RC3", "PSK500RC4", "PSK800C2", "PSK800RC2", "PSK1000", "PSK1000C2", "PSK1000R", "PSK1000RC2", "PSKAM10", "PSKAM31", "PSKAM50", "PSKFEC31", "QPSK31", "QPSK63", "QPSK125", "QPSK250", "QPSK500", "SIM31"}, Description: ""},
	PSK10:        {IsImportOnly: true, Key: "PSK10", Submodes: submode.SubModeList(nil), Description: ""},
	PSK125:       {IsImportOnly: true, Key: "PSK125", Submodes: submode.SubModeList(nil), Description: ""},
	PSK2K:        {IsImportOnly: false, Key: "PSK2K", Submodes: submode.SubModeList(nil), Description: ""},
	PSK31:        {IsImportOnly: true, Key: "PSK31", Submodes: submode.SubModeList(nil), Description: ""},
	PSK63:        {IsImportOnly: true, Key: "PSK63", Submodes: submode.SubModeList(nil), Description: ""},
	PSK63F:       {IsImportOnly: true, Key: "PSK63F", Submodes: submode.SubModeList(nil), Description: ""},
	PSKAM10:      {IsImportOnly: true, Key: "PSKAM10", Submodes: submode.SubModeList(nil), Description: ""},
	PSKAM31:      {IsImportOnly: true, Key: "PSKAM31", Submodes: submode.SubModeList(nil), Description: ""},
	PSKAM50:      {IsImportOnly: true, Key: "PSKAM50", Submodes: submode.SubModeList(nil), Description: ""},
	PSKFEC31:     {IsImportOnly: true, Key: "PSKFEC31", Submodes: submode.SubModeList(nil), Description: ""},
	PSKHELL:      {IsImportOnly: true, Key: "PSKHELL", Submodes: submode.SubModeList(nil), Description: ""},
	Q15:          {IsImportOnly: false, Key: "Q15", Submodes: submode.SubModeList(nil), Description: ""},
	QPSK125:      {IsImportOnly: true, Key: "QPSK125", Submodes: submode.SubModeList(nil), Description: ""},
	QPSK31:       {IsImportOnly: true, Key: "QPSK31", Submodes: submode.SubModeList(nil), Description: ""},
	QPSK63:       {IsImportOnly: true, Key: "QPSK63", Submodes: submode.SubModeList(nil), Description: ""},
	QRA64:        {IsImportOnly: false, Key: "QRA64", Submodes: submode.SubModeList{"QRA64A", "QRA64B", "QRA64C", "QRA64D", "QRA64E"}, Description: ""},
	ROS:          {IsImportOnly: false, Key: "ROS", Submodes: submode.SubModeList{"ROS-EME", "ROS-HF", "ROS-MF"}, Description: ""},
	RTTY:         {IsImportOnly: false, Key: "RTTY", Submodes: submode.SubModeList{"ASCI"}, Description: ""},
	RTTYM:        {IsImportOnly: false, Key: "RTTYM", Submodes: submode.SubModeList(nil), Description: ""},
	SSB:          {IsImportOnly: false, Key: "SSB", Submodes: submode.SubModeList{"LSB", "USB"}, Description: ""},
	SSTV:         {IsImportOnly: false, Key: "SSTV", Submodes: submode.SubModeList(nil), Description: ""},
	T10:          {IsImportOnly: false, Key: "T10", Submodes: submode.SubModeList(nil), Description: "Tonal 10 digital mode with focus on sensitivity, band capacity and resistance to the HF Doppler frequency spread"},
	THOR:         {IsImportOnly: false, Key: "THOR", Submodes: submode.SubModeList{"THOR-M", "THOR4", "THOR5", "THOR8", "THOR11", "THOR16", "THOR22", "THOR25X4", "THOR50X1", "THOR50X2", "THOR100"}, Description: ""},
	THRB:         {IsImportOnly: false, Key: "THRB", Submodes: submode.SubModeList{"THRBX", "THRBX1", "THRBX2", "THRBX4", "THROB1", "THROB2", "THROB4"}, Description: ""},
	THRBX:        {IsImportOnly: true, Key: "THRBX", Submodes: submode.SubModeList(nil), Description: ""},
	TOR:          {IsImportOnly: false, Key: "TOR", Submodes: submode.SubModeList{"AMTORFEC", "GTOR", "NAVTEX", "SITORB"}, Description: ""},
	V4:           {IsImportOnly: false, Key: "V4", Submodes: submode.SubModeList(nil), Description: ""},
	VOI:          {IsImportOnly: false, Key: "VOI", Submodes: submode.SubModeList(nil), Description: ""},
	WINMOR:       {IsImportOnly: false, Key: "WINMOR", Submodes: submode.SubModeList(nil), Description: ""},
	WSPR:         {IsImportOnly: false, Key: "WSPR", Submodes: submode.SubModeList(nil), Description: ""},
}

var internalModeListAll = []Spec{
	internalModeMap[AM],
	internalModeMap[AMTORFEC],
	internalModeMap[ARDOP],
	internalModeMap[ASCI],
	internalModeMap[ATV],
	internalModeMap[C4FM],
	internalModeMap[CHIP],
	internalModeMap[CHIP128],
	internalModeMap[CHIP64],
	internalModeMap[CLO],
	internalModeMap[CONTESTI],
	internalModeMap[CW],
	internalModeMap[DIGITALVOICE],
	internalModeMap[DOMINO],
	internalModeMap[DOMINOF],
	internalModeMap[DSTAR],
	internalModeMap[DYNAMIC],
	internalModeMap[FAX],
	internalModeMap[FM],
	internalModeMap[FMHELL],
	internalModeMap[FSK],
	internalModeMap[FSK31],
	internalModeMap[FSK441],
	internalModeMap[FT8],
	internalModeMap[GTOR],
	internalModeMap[HELL],
	internalModeMap[HELL80],
	internalModeMap[HFSK],
	internalModeMap[ISCAT],
	internalModeMap[JT4],
	internalModeMap[JT44],
	internalModeMap[JT4A],
	internalModeMap[JT4B],
	internalModeMap[JT4C],
	internalModeMap[JT4D],
	internalModeMap[JT4E],
	internalModeMap[JT4F],
	internalModeMap[JT4G],
	internalModeMap[JT65],
	internalModeMap[JT65A],
	internalModeMap[JT65B],
	internalModeMap[JT65C],
	internalModeMap[JT6M],
	internalModeMap[JT9],
	internalModeMap[MFSK],
	internalModeMap[MFSK16],
	internalModeMap[MFSK8],
	internalModeMap[MSK144],
	internalModeMap[MT63],
	internalModeMap[MTONE],
	internalModeMap[OLIVIA],
	internalModeMap[OPERA],
	internalModeMap[PAC],
	internalModeMap[PAC2],
	internalModeMap[PAC3],
	internalModeMap[PAX],
	internalModeMap[PAX2],
	internalModeMap[PCW],
	internalModeMap[PKT],
	internalModeMap[PSK],
	internalModeMap[PSK10],
	internalModeMap[PSK125],
	internalModeMap[PSK2K],
	internalModeMap[PSK31],
	internalModeMap[PSK63],
	internalModeMap[PSK63F],
	internalModeMap[PSKAM10],
	internalModeMap[PSKAM31],
	internalModeMap[PSKAM50],
	internalModeMap[PSKFEC31],
	internalModeMap[PSKHELL],
	internalModeMap[Q15],
	internalModeMap[QPSK125],
	internalModeMap[QPSK31],
	internalModeMap[QPSK63],
	internalModeMap[QRA64],
	internalModeMap[ROS],
	internalModeMap[RTTY],
	internalModeMap[RTTYM],
	internalModeMap[SSB],
	internalModeMap[SSTV],
	internalModeMap[T10],
	internalModeMap[THOR],
	internalModeMap[THRB],
	internalModeMap[THRBX],
	internalModeMap[TOR],
	internalModeMap[V4],
	internalModeMap[VOI],
	internalModeMap[WINMOR],
	internalModeMap[WSPR],
}

var internalModeListCurrent = []Spec{
	internalModeMap[AM],
	internalModeMap[ARDOP],
	internalModeMap[ATV],
	internalModeMap[CHIP],
	internalModeMap[CLO],
	internalModeMap[CONTESTI],
	internalModeMap[CW],
	internalModeMap[DIGITALVOICE],
	internalModeMap[DOMINO],
	internalModeMap[DYNAMIC],
	internalModeMap[FAX],
	internalModeMap[FM],
	internalModeMap[FSK],
	internalModeMap[FSK441],
	internalModeMap[FT8],
	internalModeMap[HELL],
	internalModeMap[ISCAT],
	internalModeMap[JT4],
	internalModeMap[JT44],
	internalModeMap[JT65],
	internalModeMap[JT6M],
	internalModeMap[JT9],
	internalModeMap[MFSK],
	internalModeMap[MSK144],
	internalModeMap[MT63],
	internalModeMap[MTONE],
	internalModeMap[OLIVIA],
	internalModeMap[OPERA],
	internalModeMap[PAC],
	internalModeMap[PAX],
	internalModeMap[PKT],
	internalModeMap[PSK],
	internalModeMap[PSK2K],
	internalModeMap[Q15],
	internalModeMap[QRA64],
	internalModeMap[ROS],
	internalModeMap[RTTY],
	internalModeMap[RTTYM],
	internalModeMap[SSB],
	internalModeMap[SSTV],
	internalModeMap[T10],
	internalModeMap[THOR],
	internalModeMap[THRB],
	internalModeMap[TOR],
	internalModeMap[V4],
	internalModeMap[VOI],
	internalModeMap[WINMOR],
	internalModeMap[WSPR],
}
