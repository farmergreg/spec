// DO NOT EDIT; GENERATED CODE
// ADIF: 3.1.6 Proposed

package mode

import "github.com/hamradiolog-net/adif-spec/v2/src/pkg/enum/submode"

var (
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

// A map of all Mode specifications.
var ModeMap = map[Mode]Spec{
	AM:           {IsImportOnly: false, Comments: "", Key: "AM", Submodes: submode.SubModeList(nil), Description: ""},
	AMTORFEC:     {IsImportOnly: true, Comments: "", Key: "AMTORFEC", Submodes: submode.SubModeList(nil), Description: ""},
	ARDOP:        {IsImportOnly: false, Comments: "", Key: "ARDOP", Submodes: submode.SubModeList(nil), Description: "Amateur Radio Digital Open Protocol"},
	ASCI:         {IsImportOnly: true, Comments: "", Key: "ASCI", Submodes: submode.SubModeList(nil), Description: ""},
	ATV:          {IsImportOnly: false, Comments: "", Key: "ATV", Submodes: submode.SubModeList(nil), Description: ""},
	C4FM:         {IsImportOnly: true, Comments: "", Key: "C4FM", Submodes: submode.SubModeList(nil), Description: "C4FM 4-level FSK Technology Imported QSOs with <MODE:4>C4FM> must be exported as: <MODE:12>DIGITALVOICE <SUBMODE:4>C4FM"},
	CHIP:         {IsImportOnly: false, Comments: "", Key: "CHIP", Submodes: submode.SubModeList{"CHIP64", "CHIP128"}, Description: ""},
	CHIP128:      {IsImportOnly: true, Comments: "", Key: "CHIP128", Submodes: submode.SubModeList(nil), Description: ""},
	CHIP64:       {IsImportOnly: true, Comments: "", Key: "CHIP64", Submodes: submode.SubModeList(nil), Description: ""},
	CLO:          {IsImportOnly: false, Comments: "", Key: "CLO", Submodes: submode.SubModeList(nil), Description: ""},
	CONTESTI:     {IsImportOnly: false, Comments: "", Key: "CONTESTI", Submodes: submode.SubModeList(nil), Description: ""},
	CW:           {IsImportOnly: false, Comments: "", Key: "CW", Submodes: submode.SubModeList{"PCW"}, Description: ""},
	DIGITALVOICE: {IsImportOnly: false, Comments: "", Key: "DIGITALVOICE", Submodes: submode.SubModeList{"C4FM", "DMR", "DSTAR", "FREEDV", "M17"}, Description: ""},
	DOMINO:       {IsImportOnly: false, Comments: "", Key: "DOMINO", Submodes: submode.SubModeList{"DOM-M", "DOM4", "DOM5", "DOM8", "DOM11", "DOM16", "DOM22", "DOM44", "DOM88", "DOMINOEX", "DOMINOF"}, Description: ""},
	DOMINOF:      {IsImportOnly: true, Comments: "", Key: "DOMINOF", Submodes: submode.SubModeList(nil), Description: ""},
	DSTAR:        {IsImportOnly: true, Comments: "", Key: "DSTAR", Submodes: submode.SubModeList(nil), Description: "Digital Smart Technologies for Amateur Radio Imported QSOs with <MODE:5>DSTAR must be exported as: <MODE:12>DIGITALVOICE <SUBMODE:5>DSTAR"},
	DYNAMIC:      {IsImportOnly: false, Comments: "", Key: "DYNAMIC", Submodes: submode.SubModeList{"VARA HF", "VARA SATELLITE", "VARA FM 1200", "VARA FM 9600"}, Description: ""},
	FAX:          {IsImportOnly: false, Comments: "", Key: "FAX", Submodes: submode.SubModeList(nil), Description: ""},
	FM:           {IsImportOnly: false, Comments: "", Key: "FM", Submodes: submode.SubModeList(nil), Description: ""},
	FMHELL:       {IsImportOnly: true, Comments: "", Key: "FMHELL", Submodes: submode.SubModeList(nil), Description: ""},
	FSK:          {IsImportOnly: false, Comments: "", Key: "FSK", Submodes: submode.SubModeList{"SCAMP_FAST", "SCAMP_SLOW", "SCAMP_VSLOW"}, Description: "Frequency shift keying"},
	FSK31:        {IsImportOnly: true, Comments: "", Key: "FSK31", Submodes: submode.SubModeList(nil), Description: ""},
	FSK441:       {IsImportOnly: false, Comments: "", Key: "FSK441", Submodes: submode.SubModeList(nil), Description: ""},
	FT8:          {IsImportOnly: false, Comments: "", Key: "FT8", Submodes: submode.SubModeList(nil), Description: "Franke-Taylor design, 8-FSK modulation"},
	GTOR:         {IsImportOnly: true, Comments: "", Key: "GTOR", Submodes: submode.SubModeList(nil), Description: ""},
	HELL:         {IsImportOnly: false, Comments: "", Key: "HELL", Submodes: submode.SubModeList{"FMHELL", "FSKH105", "FSKH245", "FSKHELL", "HELL80", "HELLX5", "HELLX9", "HFSK", "PSKHELL", "SLOWHELL"}, Description: ""},
	HELL80:       {IsImportOnly: true, Comments: "", Key: "HELL80", Submodes: submode.SubModeList(nil), Description: ""},
	HFSK:         {IsImportOnly: true, Comments: "", Key: "HFSK", Submodes: submode.SubModeList(nil), Description: ""},
	ISCAT:        {IsImportOnly: false, Comments: "", Key: "ISCAT", Submodes: submode.SubModeList{"ISCAT-A", "ISCAT-B"}, Description: ""},
	JT4:          {IsImportOnly: false, Comments: "", Key: "JT4", Submodes: submode.SubModeList{"JT4A", "JT4B", "JT4C", "JT4D", "JT4E", "JT4F", "JT4G"}, Description: ""},
	JT44:         {IsImportOnly: false, Comments: "", Key: "JT44", Submodes: submode.SubModeList(nil), Description: ""},
	JT4A:         {IsImportOnly: true, Comments: "", Key: "JT4A", Submodes: submode.SubModeList(nil), Description: ""},
	JT4B:         {IsImportOnly: true, Comments: "", Key: "JT4B", Submodes: submode.SubModeList(nil), Description: ""},
	JT4C:         {IsImportOnly: true, Comments: "", Key: "JT4C", Submodes: submode.SubModeList(nil), Description: ""},
	JT4D:         {IsImportOnly: true, Comments: "", Key: "JT4D", Submodes: submode.SubModeList(nil), Description: ""},
	JT4E:         {IsImportOnly: true, Comments: "", Key: "JT4E", Submodes: submode.SubModeList(nil), Description: ""},
	JT4F:         {IsImportOnly: true, Comments: "", Key: "JT4F", Submodes: submode.SubModeList(nil), Description: ""},
	JT4G:         {IsImportOnly: true, Comments: "", Key: "JT4G", Submodes: submode.SubModeList(nil), Description: ""},
	JT65:         {IsImportOnly: false, Comments: "", Key: "JT65", Submodes: submode.SubModeList{"JT65A", "JT65B", "JT65B2", "JT65C", "JT65C2"}, Description: ""},
	JT65A:        {IsImportOnly: true, Comments: "", Key: "JT65A", Submodes: submode.SubModeList(nil), Description: ""},
	JT65B:        {IsImportOnly: true, Comments: "", Key: "JT65B", Submodes: submode.SubModeList(nil), Description: ""},
	JT65C:        {IsImportOnly: true, Comments: "", Key: "JT65C", Submodes: submode.SubModeList(nil), Description: ""},
	JT6M:         {IsImportOnly: false, Comments: "", Key: "JT6M", Submodes: submode.SubModeList(nil), Description: ""},
	JT9:          {IsImportOnly: false, Comments: "", Key: "JT9", Submodes: submode.SubModeList{"JT9-1", "JT9-2", "JT9-5", "JT9-10", "JT9-30", "JT9A", "JT9B", "JT9C", "JT9D", "JT9E", "JT9E FAST", "JT9F", "JT9F FAST", "JT9G", "JT9G FAST", "JT9H", "JT9H FAST"}, Description: ""},
	MFSK:         {IsImportOnly: false, Comments: "", Key: "MFSK", Submodes: submode.SubModeList{"FSQCALL", "FST4", "FST4W", "FT4", "JS8", "JTMS", "MFSK4", "MFSK8", "MFSK11", "MFSK16", "MFSK22", "MFSK31", "MFSK32", "MFSK64", "MFSK64L", "MFSK128 MFSK128L", "Q65"}, Description: ""},
	MFSK16:       {IsImportOnly: true, Comments: "", Key: "MFSK16", Submodes: submode.SubModeList(nil), Description: ""},
	MFSK8:        {IsImportOnly: true, Comments: "", Key: "MFSK8", Submodes: submode.SubModeList(nil), Description: ""},
	MSK144:       {IsImportOnly: false, Comments: "", Key: "MSK144", Submodes: submode.SubModeList(nil), Description: ""},
	MT63:         {IsImportOnly: false, Comments: "", Key: "MT63", Submodes: submode.SubModeList(nil), Description: ""},
	MTONE:        {IsImportOnly: false, Comments: "", Key: "MTONE", Submodes: submode.SubModeList{"SCAMP_OO", "SCAMP_OO_SLW"}, Description: "Single modulated tone"},
	OLIVIA:       {IsImportOnly: false, Comments: "", Key: "OLIVIA", Submodes: submode.SubModeList{"OLIVIA 4/125", "OLIVIA 4/250", "OLIVIA 8/250", "OLIVIA 8/500", "OLIVIA 16/500", "OLIVIA 16/1000", "OLIVIA 32/1000"}, Description: ""},
	OPERA:        {IsImportOnly: false, Comments: "", Key: "OPERA", Submodes: submode.SubModeList{"OPERA-BEACON", "OPERA-QSO"}, Description: ""},
	PAC:          {IsImportOnly: false, Comments: "", Key: "PAC", Submodes: submode.SubModeList{"PAC2", "PAC3", "PAC4"}, Description: ""},
	PAC2:         {IsImportOnly: true, Comments: "", Key: "PAC2", Submodes: submode.SubModeList(nil), Description: ""},
	PAC3:         {IsImportOnly: true, Comments: "", Key: "PAC3", Submodes: submode.SubModeList(nil), Description: ""},
	PAX:          {IsImportOnly: false, Comments: "", Key: "PAX", Submodes: submode.SubModeList{"PAX2"}, Description: ""},
	PAX2:         {IsImportOnly: true, Comments: "", Key: "PAX2", Submodes: submode.SubModeList(nil), Description: ""},
	PCW:          {IsImportOnly: true, Comments: "", Key: "PCW", Submodes: submode.SubModeList(nil), Description: ""},
	PKT:          {IsImportOnly: false, Comments: "", Key: "PKT", Submodes: submode.SubModeList(nil), Description: ""},
	PSK:          {IsImportOnly: false, Comments: "", Key: "PSK", Submodes: submode.SubModeList{"8PSK125", "8PSK125F", "8PSK125FL", "8PSK250", "8PSK250F", "8PSK250FL", "8PSK500", "8PSK500F", "8PSK1000", "8PSK1000F", "8PSK1200F", "FSK31", "PSK10", "PSK31", "PSK63", "PSK63F", "PSK63RC4", "PSK63RC5", "PSK63RC10", "PSK63RC20", "PSK63RC32", "PSK125", "PSK125C12", "PSK125R", "PSK125RC10", "PSK125RC12", "PSK125RC16", "PSK125RC4", "PSK125RC5", "PSK250", "PSK250C6", "PSK250R", "PSK250RC2", "PSK250RC3", "PSK250RC5", "PSK250RC6", "PSK250RC7", "PSK500", "PSK500C2", "PSK500C4", "PSK500R", "PSK500RC2", "PSK500RC3", "PSK500RC4", "PSK800C2", "PSK800RC2", "PSK1000", "PSK1000C2", "PSK1000R", "PSK1000RC2", "PSKAM10", "PSKAM31", "PSKAM50", "PSKFEC31", "QPSK31", "QPSK63", "QPSK125", "QPSK250", "QPSK500", "SIM31"}, Description: ""},
	PSK10:        {IsImportOnly: true, Comments: "", Key: "PSK10", Submodes: submode.SubModeList(nil), Description: ""},
	PSK125:       {IsImportOnly: true, Comments: "", Key: "PSK125", Submodes: submode.SubModeList(nil), Description: ""},
	PSK2K:        {IsImportOnly: false, Comments: "", Key: "PSK2K", Submodes: submode.SubModeList(nil), Description: ""},
	PSK31:        {IsImportOnly: true, Comments: "", Key: "PSK31", Submodes: submode.SubModeList(nil), Description: ""},
	PSK63:        {IsImportOnly: true, Comments: "", Key: "PSK63", Submodes: submode.SubModeList(nil), Description: ""},
	PSK63F:       {IsImportOnly: true, Comments: "", Key: "PSK63F", Submodes: submode.SubModeList(nil), Description: ""},
	PSKAM10:      {IsImportOnly: true, Comments: "", Key: "PSKAM10", Submodes: submode.SubModeList(nil), Description: ""},
	PSKAM31:      {IsImportOnly: true, Comments: "", Key: "PSKAM31", Submodes: submode.SubModeList(nil), Description: ""},
	PSKAM50:      {IsImportOnly: true, Comments: "", Key: "PSKAM50", Submodes: submode.SubModeList(nil), Description: ""},
	PSKFEC31:     {IsImportOnly: true, Comments: "", Key: "PSKFEC31", Submodes: submode.SubModeList(nil), Description: ""},
	PSKHELL:      {IsImportOnly: true, Comments: "", Key: "PSKHELL", Submodes: submode.SubModeList(nil), Description: ""},
	Q15:          {IsImportOnly: false, Comments: "", Key: "Q15", Submodes: submode.SubModeList(nil), Description: ""},
	QPSK125:      {IsImportOnly: true, Comments: "", Key: "QPSK125", Submodes: submode.SubModeList(nil), Description: ""},
	QPSK31:       {IsImportOnly: true, Comments: "", Key: "QPSK31", Submodes: submode.SubModeList(nil), Description: ""},
	QPSK63:       {IsImportOnly: true, Comments: "", Key: "QPSK63", Submodes: submode.SubModeList(nil), Description: ""},
	QRA64:        {IsImportOnly: false, Comments: "", Key: "QRA64", Submodes: submode.SubModeList{"QRA64A", "QRA64B", "QRA64C", "QRA64D", "QRA64E"}, Description: ""},
	ROS:          {IsImportOnly: false, Comments: "", Key: "ROS", Submodes: submode.SubModeList{"ROS-EME", "ROS-HF", "ROS-MF"}, Description: ""},
	RTTY:         {IsImportOnly: false, Comments: "", Key: "RTTY", Submodes: submode.SubModeList{"ASCI"}, Description: ""},
	RTTYM:        {IsImportOnly: false, Comments: "", Key: "RTTYM", Submodes: submode.SubModeList(nil), Description: ""},
	SSB:          {IsImportOnly: false, Comments: "", Key: "SSB", Submodes: submode.SubModeList{"LSB", "USB"}, Description: ""},
	SSTV:         {IsImportOnly: false, Comments: "", Key: "SSTV", Submodes: submode.SubModeList(nil), Description: ""},
	T10:          {IsImportOnly: false, Comments: "", Key: "T10", Submodes: submode.SubModeList(nil), Description: "Tonal 10 digital mode with focus on sensitivity, band capacity and resistance to the HF Doppler frequency spread"},
	THOR:         {IsImportOnly: false, Comments: "", Key: "THOR", Submodes: submode.SubModeList{"THOR-M", "THOR4", "THOR5", "THOR8", "THOR11", "THOR16", "THOR22", "THOR25X4", "THOR50X1", "THOR50X2", "THOR100"}, Description: ""},
	THRB:         {IsImportOnly: false, Comments: "", Key: "THRB", Submodes: submode.SubModeList{"THRBX", "THRBX1", "THRBX2", "THRBX4", "THROB1", "THROB2", "THROB4"}, Description: ""},
	THRBX:        {IsImportOnly: true, Comments: "", Key: "THRBX", Submodes: submode.SubModeList(nil), Description: ""},
	TOR:          {IsImportOnly: false, Comments: "", Key: "TOR", Submodes: submode.SubModeList{"AMTORFEC", "GTOR", "NAVTEX", "SITORB"}, Description: ""},
	V4:           {IsImportOnly: false, Comments: "", Key: "V4", Submodes: submode.SubModeList(nil), Description: ""},
	VOI:          {IsImportOnly: false, Comments: "", Key: "VOI", Submodes: submode.SubModeList(nil), Description: ""},
	WINMOR:       {IsImportOnly: false, Comments: "", Key: "WINMOR", Submodes: submode.SubModeList(nil), Description: ""},
	WSPR:         {IsImportOnly: false, Comments: "", Key: "WSPR", Submodes: submode.SubModeList(nil), Description: ""},
}

// All Mode specifications including depreciated and import only.
var ModeListAll = []Spec{
	ModeMap[AM],
	ModeMap[AMTORFEC],
	ModeMap[ARDOP],
	ModeMap[ASCI],
	ModeMap[ATV],
	ModeMap[C4FM],
	ModeMap[CHIP],
	ModeMap[CHIP128],
	ModeMap[CHIP64],
	ModeMap[CLO],
	ModeMap[CONTESTI],
	ModeMap[CW],
	ModeMap[DIGITALVOICE],
	ModeMap[DOMINO],
	ModeMap[DOMINOF],
	ModeMap[DSTAR],
	ModeMap[DYNAMIC],
	ModeMap[FAX],
	ModeMap[FM],
	ModeMap[FMHELL],
	ModeMap[FSK],
	ModeMap[FSK31],
	ModeMap[FSK441],
	ModeMap[FT8],
	ModeMap[GTOR],
	ModeMap[HELL],
	ModeMap[HELL80],
	ModeMap[HFSK],
	ModeMap[ISCAT],
	ModeMap[JT4],
	ModeMap[JT44],
	ModeMap[JT4A],
	ModeMap[JT4B],
	ModeMap[JT4C],
	ModeMap[JT4D],
	ModeMap[JT4E],
	ModeMap[JT4F],
	ModeMap[JT4G],
	ModeMap[JT65],
	ModeMap[JT65A],
	ModeMap[JT65B],
	ModeMap[JT65C],
	ModeMap[JT6M],
	ModeMap[JT9],
	ModeMap[MFSK],
	ModeMap[MFSK16],
	ModeMap[MFSK8],
	ModeMap[MSK144],
	ModeMap[MT63],
	ModeMap[MTONE],
	ModeMap[OLIVIA],
	ModeMap[OPERA],
	ModeMap[PAC],
	ModeMap[PAC2],
	ModeMap[PAC3],
	ModeMap[PAX],
	ModeMap[PAX2],
	ModeMap[PCW],
	ModeMap[PKT],
	ModeMap[PSK],
	ModeMap[PSK10],
	ModeMap[PSK125],
	ModeMap[PSK2K],
	ModeMap[PSK31],
	ModeMap[PSK63],
	ModeMap[PSK63F],
	ModeMap[PSKAM10],
	ModeMap[PSKAM31],
	ModeMap[PSKAM50],
	ModeMap[PSKFEC31],
	ModeMap[PSKHELL],
	ModeMap[Q15],
	ModeMap[QPSK125],
	ModeMap[QPSK31],
	ModeMap[QPSK63],
	ModeMap[QRA64],
	ModeMap[ROS],
	ModeMap[RTTY],
	ModeMap[RTTYM],
	ModeMap[SSB],
	ModeMap[SSTV],
	ModeMap[T10],
	ModeMap[THOR],
	ModeMap[THRB],
	ModeMap[THRBX],
	ModeMap[TOR],
	ModeMap[V4],
	ModeMap[VOI],
	ModeMap[WINMOR],
	ModeMap[WSPR],
}

// All Mode specifications values that are NOT marked import-only.
var ModeListCurrent = []Spec{
	ModeMap[AM],
	ModeMap[ARDOP],
	ModeMap[ATV],
	ModeMap[CHIP],
	ModeMap[CLO],
	ModeMap[CONTESTI],
	ModeMap[CW],
	ModeMap[DIGITALVOICE],
	ModeMap[DOMINO],
	ModeMap[DYNAMIC],
	ModeMap[FAX],
	ModeMap[FM],
	ModeMap[FSK],
	ModeMap[FSK441],
	ModeMap[FT8],
	ModeMap[HELL],
	ModeMap[ISCAT],
	ModeMap[JT4],
	ModeMap[JT44],
	ModeMap[JT65],
	ModeMap[JT6M],
	ModeMap[JT9],
	ModeMap[MFSK],
	ModeMap[MSK144],
	ModeMap[MT63],
	ModeMap[MTONE],
	ModeMap[OLIVIA],
	ModeMap[OPERA],
	ModeMap[PAC],
	ModeMap[PAX],
	ModeMap[PKT],
	ModeMap[PSK],
	ModeMap[PSK2K],
	ModeMap[Q15],
	ModeMap[QRA64],
	ModeMap[ROS],
	ModeMap[RTTY],
	ModeMap[RTTYM],
	ModeMap[SSB],
	ModeMap[SSTV],
	ModeMap[T10],
	ModeMap[THOR],
	ModeMap[THRB],
	ModeMap[TOR],
	ModeMap[V4],
	ModeMap[VOI],
	ModeMap[WINMOR],
	ModeMap[WSPR],
}
