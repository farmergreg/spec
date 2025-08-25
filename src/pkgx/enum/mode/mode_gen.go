// GENERATED CODE; DO NOT EDIT
// ADIF: 3.1.6 Proposed

package mode

var (
	AM           Mode = "AM"           //
	AMTORFEC     Mode = "AMTORFEC"     // Deprecated:
	ARDOP        Mode = "ARDOP"        // Amateur Radio Digital Open Protocol
	ASCI         Mode = "ASCI"         // Deprecated:
	ATV          Mode = "ATV"          //
	C4FM         Mode = "C4FM"         // Deprecated: C4FM 4-level FSK Technology Imported QSOs with <MODE:4>C4FM> must be exported as: <MODE:12>DIGITALVOICE <SUBMODE:4>C4FM
	CHIP         Mode = "CHIP"         //
	CHIP128      Mode = "CHIP128"      // Deprecated:
	CHIP64       Mode = "CHIP64"       // Deprecated:
	CLO          Mode = "CLO"          //
	CONTESTI     Mode = "CONTESTI"     //
	CW           Mode = "CW"           //
	DIGITALVOICE Mode = "DIGITALVOICE" //
	DOMINO       Mode = "DOMINO"       //
	DOMINOF      Mode = "DOMINOF"      // Deprecated:
	DSTAR        Mode = "DSTAR"        // Deprecated: Digital Smart Technologies for Amateur Radio Imported QSOs with <MODE:5>DSTAR must be exported as: <MODE:12>DIGITALVOICE <SUBMODE:5>DSTAR
	DYNAMIC      Mode = "DYNAMIC"      //
	FAX          Mode = "FAX"          //
	FM           Mode = "FM"           //
	FMHELL       Mode = "FMHELL"       // Deprecated:
	FSK          Mode = "FSK"          // Frequency shift keying
	FSK31        Mode = "FSK31"        // Deprecated:
	FSK441       Mode = "FSK441"       //
	FT8          Mode = "FT8"          // Franke-Taylor design, 8-FSK modulation
	GTOR         Mode = "GTOR"         // Deprecated:
	HELL         Mode = "HELL"         //
	HELL80       Mode = "HELL80"       // Deprecated:
	HFSK         Mode = "HFSK"         // Deprecated:
	ISCAT        Mode = "ISCAT"        //
	JT4          Mode = "JT4"          //
	JT44         Mode = "JT44"         //
	JT4A         Mode = "JT4A"         // Deprecated:
	JT4B         Mode = "JT4B"         // Deprecated:
	JT4C         Mode = "JT4C"         // Deprecated:
	JT4D         Mode = "JT4D"         // Deprecated:
	JT4E         Mode = "JT4E"         // Deprecated:
	JT4F         Mode = "JT4F"         // Deprecated:
	JT4G         Mode = "JT4G"         // Deprecated:
	JT65         Mode = "JT65"         //
	JT65A        Mode = "JT65A"        // Deprecated:
	JT65B        Mode = "JT65B"        // Deprecated:
	JT65C        Mode = "JT65C"        // Deprecated:
	JT6M         Mode = "JT6M"         //
	JT9          Mode = "JT9"          //
	MFSK         Mode = "MFSK"         //
	MFSK16       Mode = "MFSK16"       // Deprecated:
	MFSK8        Mode = "MFSK8"        // Deprecated:
	MSK144       Mode = "MSK144"       //
	MT63         Mode = "MT63"         //
	MTONE        Mode = "MTONE"        // Single modulated tone
	OLIVIA       Mode = "OLIVIA"       //
	OPERA        Mode = "OPERA"        //
	PAC          Mode = "PAC"          //
	PAC2         Mode = "PAC2"         // Deprecated:
	PAC3         Mode = "PAC3"         // Deprecated:
	PAX          Mode = "PAX"          //
	PAX2         Mode = "PAX2"         // Deprecated:
	PCW          Mode = "PCW"          // Deprecated:
	PKT          Mode = "PKT"          //
	PSK          Mode = "PSK"          //
	PSK10        Mode = "PSK10"        // Deprecated:
	PSK125       Mode = "PSK125"       // Deprecated:
	PSK2K        Mode = "PSK2K"        //
	PSK31        Mode = "PSK31"        // Deprecated:
	PSK63        Mode = "PSK63"        // Deprecated:
	PSK63F       Mode = "PSK63F"       // Deprecated:
	PSKAM10      Mode = "PSKAM10"      // Deprecated:
	PSKAM31      Mode = "PSKAM31"      // Deprecated:
	PSKAM50      Mode = "PSKAM50"      // Deprecated:
	PSKFEC31     Mode = "PSKFEC31"     // Deprecated:
	PSKHELL      Mode = "PSKHELL"      // Deprecated:
	Q15          Mode = "Q15"          //
	QPSK125      Mode = "QPSK125"      // Deprecated:
	QPSK31       Mode = "QPSK31"       // Deprecated:
	QPSK63       Mode = "QPSK63"       // Deprecated:
	QRA64        Mode = "QRA64"        //
	ROS          Mode = "ROS"          //
	RTTY         Mode = "RTTY"         //
	RTTYM        Mode = "RTTYM"        //
	SSB          Mode = "SSB"          //
	SSTV         Mode = "SSTV"         //
	T10          Mode = "T10"          // Tonal 10 digital mode with focus on sensitivity, band capacity and resistance to the HF Doppler frequency spread
	THOR         Mode = "THOR"         //
	THRB         Mode = "THRB"         //
	THRBX        Mode = "THRBX"        // Deprecated:
	TOR          Mode = "TOR"          //
	V4           Mode = "V4"           //
	VOI          Mode = "VOI"          //
	WINMOR       Mode = "WINMOR"       //
	WSPR         Mode = "WSPR"         //
)

// IsValid returns true if the Mode exists in the ADIF specification. This includes Import Only and Deleted values.
func (value Mode) IsValid() bool {
	switch value {
	case AM:
		return true
	case AMTORFEC:
		return true
	case ARDOP:
		return true
	case ASCI:
		return true
	case ATV:
		return true
	case C4FM:
		return true
	case CHIP:
		return true
	case CHIP128:
		return true
	case CHIP64:
		return true
	case CLO:
		return true
	case CONTESTI:
		return true
	case CW:
		return true
	case DIGITALVOICE:
		return true
	case DOMINO:
		return true
	case DOMINOF:
		return true
	case DSTAR:
		return true
	case DYNAMIC:
		return true
	case FAX:
		return true
	case FM:
		return true
	case FMHELL:
		return true
	case FSK:
		return true
	case FSK31:
		return true
	case FSK441:
		return true
	case FT8:
		return true
	case GTOR:
		return true
	case HELL:
		return true
	case HELL80:
		return true
	case HFSK:
		return true
	case ISCAT:
		return true
	case JT4:
		return true
	case JT44:
		return true
	case JT4A:
		return true
	case JT4B:
		return true
	case JT4C:
		return true
	case JT4D:
		return true
	case JT4E:
		return true
	case JT4F:
		return true
	case JT4G:
		return true
	case JT65:
		return true
	case JT65A:
		return true
	case JT65B:
		return true
	case JT65C:
		return true
	case JT6M:
		return true
	case JT9:
		return true
	case MFSK:
		return true
	case MFSK16:
		return true
	case MFSK8:
		return true
	case MSK144:
		return true
	case MT63:
		return true
	case MTONE:
		return true
	case OLIVIA:
		return true
	case OPERA:
		return true
	case PAC:
		return true
	case PAC2:
		return true
	case PAC3:
		return true
	case PAX:
		return true
	case PAX2:
		return true
	case PCW:
		return true
	case PKT:
		return true
	case PSK:
		return true
	case PSK10:
		return true
	case PSK125:
		return true
	case PSK2K:
		return true
	case PSK31:
		return true
	case PSK63:
		return true
	case PSK63F:
		return true
	case PSKAM10:
		return true
	case PSKAM31:
		return true
	case PSKAM50:
		return true
	case PSKFEC31:
		return true
	case PSKHELL:
		return true
	case Q15:
		return true
	case QPSK125:
		return true
	case QPSK31:
		return true
	case QPSK63:
		return true
	case QRA64:
		return true
	case ROS:
		return true
	case RTTY:
		return true
	case RTTYM:
		return true
	case SSB:
		return true
	case SSTV:
		return true
	case T10:
		return true
	case THOR:
		return true
	case THRB:
		return true
	case THRBX:
		return true
	case TOR:
		return true
	case V4:
		return true
	case VOI:
		return true
	case WINMOR:
		return true
	case WSPR:
		return true
	}
	return false
}
