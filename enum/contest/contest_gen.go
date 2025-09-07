// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package contest provides code and constants as defined in ADIF 3.1.6 (Proposed)
package contest

import "sync"

const (
	Contest_070_160M_SPRINT        Contest = "070-160M-SPRINT"        // 070-160M-SPRINT      = PODXS Great Pumpkin Sprint
	Contest_070_3_DAY              Contest = "070-3-DAY"              // 070-3-DAY            = PODXS Three Day Weekend
	Contest_070_31_FLAVORS         Contest = "070-31-FLAVORS"         // 070-31-FLAVORS       = PODXS 31 Flavors
	Contest_070_40M_SPRINT         Contest = "070-40M-SPRINT"         // 070-40M-SPRINT       = PODXS 40m Firecracker Sprint
	Contest_070_80M_SPRINT         Contest = "070-80M-SPRINT"         // 070-80M-SPRINT       = PODXS 80m Jay Hudak Memorial Sprint
	Contest_070_PSKFEST            Contest = "070-PSKFEST"            // 070-PSKFEST          = PODXS PSKFest
	Contest_070_ST_PATS_DAY        Contest = "070-ST-PATS-DAY"        // 070-ST-PATS-DAY      = PODXS St. Patricks Day
	Contest_070_VALENTINE_SPRINT   Contest = "070-VALENTINE-SPRINT"   // 070-VALENTINE-SPRINT = PODXS Valentine Sprint
	Contest_10_RTTY                Contest = "10-RTTY"                // 10-RTTY              = Ten-Meter RTTY Contest (2011 onwards)
	Contest_1010_OPEN_SEASON       Contest = "1010-OPEN-SEASON"       // 1010-OPEN-SEASON     = Open Season Ten Meter QSO Party
	Contest_7QP                    Contest = "7QP"                    // 7QP                  = 7th-Area QSO Party
	Contest_AL_QSO_PARTY           Contest = "AL-QSO-PARTY"           // AL-QSO-PARTY         = Alabama QSO Party
	Contest_ALL_ASIAN_DX_CW        Contest = "ALL-ASIAN-DX-CW"        // ALL-ASIAN-DX-CW      = JARL All Asian DX Contest (CW)
	Contest_ALL_ASIAN_DX_PHONE     Contest = "ALL-ASIAN-DX-PHONE"     // ALL-ASIAN-DX-PHONE   = JARL All Asian DX Contest (PHONE)
	Contest_ANARTS_RTTY            Contest = "ANARTS-RTTY"            // ANARTS-RTTY          = ANARTS WW RTTY
	Contest_ANATOLIAN_RTTY         Contest = "ANATOLIAN-RTTY"         // ANATOLIAN-RTTY       = Anatolian WW RTTY
	Contest_AP_SPRINT              Contest = "AP-SPRINT"              // AP-SPRINT            = Asia - Pacific Sprint
	Contest_AR_QSO_PARTY           Contest = "AR-QSO-PARTY"           // AR-QSO-PARTY         = Arkansas QSO Party
	Contest_ARI_DX                 Contest = "ARI-DX"                 // ARI-DX               = ARI DX Contest
	Contest_ARI_EME                Contest = "ARI-EME"                // ARI-EME              = ARI Italian EME Trophy
	Contest_ARI_IAC_13CM           Contest = "ARI-IAC-13CM"           // ARI-IAC-13CM         = ARI Italian Activity Contest (13cm+)
	Contest_ARI_IAC_23CM           Contest = "ARI-IAC-23CM"           // ARI-IAC-23CM         = ARI Italian Activity Contest (23cm)
	Contest_ARI_IAC_6M             Contest = "ARI-IAC-6M"             // ARI-IAC-6M           = ARI Italian Activity Contest (6m)
	Contest_ARI_IAC_UHF            Contest = "ARI-IAC-UHF"            // ARI-IAC-UHF          = ARI Italian Activity Contest (UHF)
	Contest_ARI_IAC_VHF            Contest = "ARI-IAC-VHF"            // ARI-IAC-VHF          = ARI Italian Activity Contest (VHF)
	Contest_ARRL_10                Contest = "ARRL-10"                // ARRL-10              = ARRL 10 Meter Contest
	Contest_ARRL_10_GHZ            Contest = "ARRL-10-GHZ"            // ARRL-10-GHZ          = ARRL 10 GHz and Up Contest
	Contest_ARRL_160               Contest = "ARRL-160"               // ARRL-160             = ARRL 160 Meter Contest
	Contest_ARRL_222               Contest = "ARRL-222"               // ARRL-222             = ARRL 222 MHz and Up Distance Contest
	Contest_ARRL_DIGI              Contest = "ARRL-DIGI"              // ARRL-DIGI            = ARRL International Digital Contest
	Contest_ARRL_DX_CW             Contest = "ARRL-DX-CW"             // ARRL-DX-CW           = ARRL International DX Contest (CW)
	Contest_ARRL_DX_SSB            Contest = "ARRL-DX-SSB"            // ARRL-DX-SSB          = ARRL International DX Contest (Phone)
	Contest_ARRL_EME               Contest = "ARRL-EME"               // ARRL-EME             = ARRL EME contest
	Contest_ARRL_FIELD_DAY         Contest = "ARRL-FIELD-DAY"         // ARRL-FIELD-DAY       = ARRL Field Day
	Contest_ARRL_RR_CW             Contest = "ARRL-RR-CW"             // ARRL-RR-CW           = ARRL Rookie Roundup (CW)
	Contest_ARRL_RR_RTTY           Contest = "ARRL-RR-RTTY"           // ARRL-RR-RTTY         = ARRL Rookie Roundup (RTTY)
	Contest_ARRL_RR_SSB            Contest = "ARRL-RR-SSB"            // ARRL-RR-SSB          = ARRL Rookie Roundup (Phone)
	Contest_ARRL_RTTY              Contest = "ARRL-RTTY"              // ARRL-RTTY            = ARRL RTTY Round-Up
	Contest_ARRL_SCR               Contest = "ARRL-SCR"               // ARRL-SCR             = ARRL School Club Roundup
	Contest_ARRL_SS_CW             Contest = "ARRL-SS-CW"             // ARRL-SS-CW           = ARRL November Sweepstakes (CW)
	Contest_ARRL_SS_SSB            Contest = "ARRL-SS-SSB"            // ARRL-SS-SSB          = ARRL November Sweepstakes (Phone)
	Contest_ARRL_UHF_AUG           Contest = "ARRL-UHF-AUG"           // ARRL-UHF-AUG         = ARRL August UHF Contest
	Contest_ARRL_VHF_JAN           Contest = "ARRL-VHF-JAN"           // ARRL-VHF-JAN         = ARRL January VHF Sweepstakes
	Contest_ARRL_VHF_JUN           Contest = "ARRL-VHF-JUN"           // ARRL-VHF-JUN         = ARRL June VHF QSO Party
	Contest_ARRL_VHF_SEP           Contest = "ARRL-VHF-SEP"           // ARRL-VHF-SEP         = ARRL September VHF QSO Party
	Contest_AZ_QSO_PARTY           Contest = "AZ-QSO-PARTY"           // AZ-QSO-PARTY         = Arizona QSO Party
	Contest_BANGGAI_DX             Contest = "BANGGAI-DX"             // BANGGAI-DX           = ORARI Banggai DX Contest
	Contest_BARTG_RTTY             Contest = "BARTG-RTTY"             // BARTG-RTTY           = BARTG Spring RTTY Contest
	Contest_BARTG_SPRINT           Contest = "BARTG-SPRINT"           // BARTG-SPRINT         = BARTG Sprint Contest
	Contest_BC_QSO_PARTY           Contest = "BC-QSO-PARTY"           // BC-QSO-PARTY         = British Columbia QSO Party
	Contest_BEKASI_MERDEKA_CONTEST Contest = "BEKASI-MERDEKA-CONTEST" // BEKASI-MERDEKA-CONTEST = ORARI Bekasi Merdeka Contest
	Contest_CA_QSO_PARTY           Contest = "CA-QSO-PARTY"           // CA-QSO-PARTY         = California QSO Party
	Contest_CIS_DX                 Contest = "CIS-DX"                 // CIS-DX               = CIS DX Contest
	Contest_CO_QSO_PARTY           Contest = "CO-QSO-PARTY"           // CO-QSO-PARTY         = Colorado QSO Party
	Contest_CQ_160_CW              Contest = "CQ-160-CW"              // CQ-160-CW            = CQ WW 160 Meter DX Contest (CW)
	Contest_CQ_160_SSB             Contest = "CQ-160-SSB"             // CQ-160-SSB           = CQ WW 160 Meter DX Contest (SSB)
	Contest_CQ_M                   Contest = "CQ-M"                   // CQ-M                 = CQ-M International DX Contest
	Contest_CQ_VHF                 Contest = "CQ-VHF"                 // CQ-VHF               = CQ World-Wide VHF Contest
	Contest_CQ_WPX_CW              Contest = "CQ-WPX-CW"              // CQ-WPX-CW            = CQ WW WPX Contest (CW)
	Contest_CQ_WPX_RTTY            Contest = "CQ-WPX-RTTY"            // CQ-WPX-RTTY          = CQ/RJ WW RTTY WPX Contest
	Contest_CQ_WPX_SSB             Contest = "CQ-WPX-SSB"             // CQ-WPX-SSB           = CQ WW WPX Contest (SSB)
	Contest_CQ_WW_CW               Contest = "CQ-WW-CW"               // CQ-WW-CW             = CQ WW DX Contest (CW)
	Contest_CQ_WW_RTTY             Contest = "CQ-WW-RTTY"             // CQ-WW-RTTY           = CQ/RJ WW RTTY DX Contest
	Contest_CQ_WW_SSB              Contest = "CQ-WW-SSB"              // CQ-WW-SSB            = CQ WW DX Contest (SSB)
	Contest_CT_QSO_PARTY           Contest = "CT-QSO-PARTY"           // CT-QSO-PARTY         = Connecticut QSO Party
	Contest_CVA_DX_CW              Contest = "CVA-DX-CW"              // CVA-DX-CW            = Concurso Verde e Amarelo DX CW Contest
	Contest_CVA_DX_SSB             Contest = "CVA-DX-SSB"             // CVA-DX-SSB           = Concurso Verde e Amarelo DX CW Contest
	Contest_CWOPS_CW_OPEN          Contest = "CWOPS-CW-OPEN"          // CWOPS-CW-OPEN        = CWops CW Open Competition
	Contest_CWOPS_CWT              Contest = "CWOPS-CWT"              // CWOPS-CWT            = CWops Mini-CWT Test
	Contest_DARC_10                Contest = "DARC-10"                // DARC-10              = DARC 10m Contest
	Contest_DARC_CWA               Contest = "DARC-CWA"               // DARC-CWA             = DARC CW Trainee Contest
	Contest_DARC_FT4               Contest = "DARC-FT4"               // DARC-FT4             = DARC FT4 Contest
	Contest_DARC_HELL              Contest = "DARC-HELL"              // DARC-HELL            = DARC Hell Contest
	Contest_DARC_MICROWAVE         Contest = "DARC-MICROWAVE"         // DARC-MICROWAVE       = DARC Microwave Contest
	Contest_DARC_TRAINEE           Contest = "DARC-TRAINEE"           // DARC-TRAINEE         = DARC Trainee Contest
	Contest_DARC_UKW_FIELD_DAY     Contest = "DARC-UKW-FIELD-DAY"     // DARC-UKW-FIELD-DAY   = DARC UKW Summer Contest
	Contest_DARC_UKW_SPRING        Contest = "DARC-UKW-SPRING"        // DARC-UKW-SPRING      = DARC UKW Spring Contest
	Contest_DARC_VHF_UHF_MICROWAVE Contest = "DARC-VHF-UHF-MICROWAVE" // DARC-VHF-UHF-MICROWAVE = DARC VHF-, UHF-, Microwave Contest (May)
	Contest_DARC_WAEDC_CW          Contest = "DARC-WAEDC-CW"          // DARC-WAEDC-CW        = WAE DX Contest (CW)
	Contest_DARC_WAEDC_RTTY        Contest = "DARC-WAEDC-RTTY"        // DARC-WAEDC-RTTY      = WAE DX Contest (RTTY)
	Contest_DARC_WAEDC_SSB         Contest = "DARC-WAEDC-SSB"         // DARC-WAEDC-SSB       = WAE DX Contest (SSB)
	Contest_DARC_WAG               Contest = "DARC-WAG"               // DARC-WAG             = DARC Worked All Germany
	Contest_DE_QSO_PARTY           Contest = "DE-QSO-PARTY"           // DE-QSO-PARTY         = Delaware QSO Party
	Contest_DL_DX_RTTY             Contest = "DL-DX-RTTY"             // DL-DX-RTTY           = DL-DX RTTY Contest
	Contest_DMC_RTTY               Contest = "DMC-RTTY"               // DMC-RTTY             = DMC RTTY Contest
	Contest_EA_CNCW                Contest = "EA-CNCW"                // EA-CNCW              = Concurso Nacional de Telegrafía
	Contest_EA_DME                 Contest = "EA-DME"                 // EA-DME               = Municipios Españoles
	Contest_EA_MAJESTAD_CW         Contest = "EA-MAJESTAD-CW"         // EA-MAJESTAD-CW       = His Majesty The King of Spain CW Contest (2022 and later)
	Contest_EA_MAJESTAD_SSB        Contest = "EA-MAJESTAD-SSB"        // EA-MAJESTAD-SSB      = His Majesty The King of Spain SSB Contest (2022 and later)
	Contest_EA_PSK63               Contest = "EA-PSK63"               // EA-PSK63             = EA PSK63
	Contest_EA_RTTY                Contest = "EA-RTTY"                // Deprecated: EA-RTTY              = Unión de Radioaficionados Españoles RTTY Contest
	Contest_EA_SMRE_CW             Contest = "EA-SMRE-CW"             // EA-SMRE-CW           = Su Majestad El Rey de España - CW (2021 and earlier)
	Contest_EA_SMRE_SSB            Contest = "EA-SMRE-SSB"            // EA-SMRE-SSB          = Su Majestad El Rey de España - SSB (2021 and earlier)
	Contest_EA_VHF_ATLANTIC        Contest = "EA-VHF-ATLANTIC"        // EA-VHF-ATLANTIC      = Atlántico V-UHF
	Contest_EA_VHF_COM             Contest = "EA-VHF-COM"             // EA-VHF-COM           = Combinado de V-UHF
	Contest_EA_VHF_COSTA_SOL       Contest = "EA-VHF-COSTA-SOL"       // EA-VHF-COSTA-SOL     = Costa del Sol V-UHF
	Contest_EA_VHF_EA              Contest = "EA-VHF-EA"              // EA-VHF-EA            = Nacional VHF
	Contest_EA_VHF_EA1RCS          Contest = "EA-VHF-EA1RCS"          // EA-VHF-EA1RCS        = Segovia EA1RCS V-UHF
	Contest_EA_VHF_QSL             Contest = "EA-VHF-QSL"             // EA-VHF-QSL           = QSL V-UHF & 50MHz
	Contest_EA_VHF_SADURNI         Contest = "EA-VHF-SADURNI"         // EA-VHF-SADURNI       = Sant Sadurni V-UHF
	Contest_EA_WW_RTTY             Contest = "EA-WW-RTTY"             // EA-WW-RTTY           = Unión de Radioaficionados Españoles RTTY Contest
	Contest_EASTER                 Contest = "EASTER"                 // EASTER               = DARC Easter Contest
	Contest_EPC_PSK63              Contest = "EPC-PSK63"              // EPC-PSK63            = PSK63 QSO Party
	Contest_EU_SPRINT              Contest = "EU SPRINT"              // EU SPRINT            = EU Sprint
	Contest_EU_HF                  Contest = "EU-HF"                  // EU-HF                = EU HF Championship
	Contest_EU_PSK_DX              Contest = "EU-PSK-DX"              // EU-PSK-DX            = EU PSK DX Contest
	Contest_EUCW160M               Contest = "EUCW160M"               // EUCW160M             = European CW Association 160m CW Party
	Contest_FALL_SPRINT            Contest = "FALL SPRINT"            // FALL SPRINT          = FISTS Fall Sprint
	Contest_FL_QSO_PARTY           Contest = "FL-QSO-PARTY"           // FL-QSO-PARTY         = Florida QSO Party
	Contest_GA_QSO_PARTY           Contest = "GA-QSO-PARTY"           // GA-QSO-PARTY         = Georgia QSO Party
	Contest_HA_DX                  Contest = "HA-DX"                  // HA-DX                = Hungarian DX Contest
	Contest_HELVETIA               Contest = "HELVETIA"               // HELVETIA             = Helvetia Contest
	Contest_HI_QSO_PARTY           Contest = "HI-QSO-PARTY"           // HI-QSO-PARTY         = Hawaiian QSO Party
	Contest_HOLYLAND               Contest = "HOLYLAND"               // HOLYLAND             = IARC Holyland Contest
	Contest_IA_QSO_PARTY           Contest = "IA-QSO-PARTY"           // IA-QSO-PARTY         = Iowa QSO Party
	Contest_IARU_FIELD_DAY         Contest = "IARU-FIELD-DAY"         // IARU-FIELD-DAY       = DARC IARU Region 1 Field Day
	Contest_IARU_HF                Contest = "IARU-HF"                // IARU-HF              = IARU HF World Championship
	Contest_ICWC_MST               Contest = "ICWC-MST"               // ICWC-MST             = ICWC Medium Speed Test
	Contest_ID_QSO_PARTY           Contest = "ID-QSO-PARTY"           // ID-QSO-PARTY         = Idaho QSO Party
	Contest_IL_QSO_PARTY           Contest = "IL QSO PARTY"           // IL QSO PARTY         = Illinois QSO Party
	Contest_IN_QSO_PARTY           Contest = "IN-QSO-PARTY"           // IN-QSO-PARTY         = Indiana QSO Party
	Contest_JARTS_WW_RTTY          Contest = "JARTS-WW-RTTY"          // JARTS-WW-RTTY        = JARTS WW RTTY
	Contest_JIDX_CW                Contest = "JIDX-CW"                // JIDX-CW              = Japan International DX Contest (CW)
	Contest_JIDX_SSB               Contest = "JIDX-SSB"               // JIDX-SSB             = Japan International DX Contest (SSB)
	Contest_JT_DX_RTTY             Contest = "JT-DX-RTTY"             // JT-DX-RTTY           = Mongolian RTTY DX Contest
	Contest_K1USN_SSO              Contest = "K1USN-SSO"              // K1USN-SSO            = K1USN Slow Speed Open
	Contest_K1USN_SST              Contest = "K1USN-SST"              // K1USN-SST            = K1USN Slow Speed Test
	Contest_KS_QSO_PARTY           Contest = "KS-QSO-PARTY"           // KS-QSO-PARTY         = Kansas QSO Party
	Contest_KY_QSO_PARTY           Contest = "KY-QSO-PARTY"           // KY-QSO-PARTY         = Kentucky QSO Party
	Contest_LA_QSO_PARTY           Contest = "LA-QSO-PARTY"           // LA-QSO-PARTY         = Louisiana QSO Party
	Contest_LDC_RTTY               Contest = "LDC-RTTY"               // LDC-RTTY             = DRCG Long Distance Contest (RTTY)
	Contest_LZ_DX                  Contest = "LZ DX"                  // LZ DX                = LZ DX Contest
	Contest_MAR_QSO_PARTY          Contest = "MAR-QSO-PARTY"          // MAR-QSO-PARTY        = Maritimes QSO Party
	Contest_MD_QSO_PARTY           Contest = "MD-QSO-PARTY"           // MD-QSO-PARTY         = Maryland QSO Party
	Contest_ME_QSO_PARTY           Contest = "ME-QSO-PARTY"           // ME-QSO-PARTY         = Maine QSO Party
	Contest_MI_QSO_PARTY           Contest = "MI-QSO-PARTY"           // MI-QSO-PARTY         = Michigan QSO Party
	Contest_MIDATLANTIC_QSO_PARTY  Contest = "MIDATLANTIC-QSO-PARTY"  // MIDATLANTIC-QSO-PARTY = Mid-Atlantic QSO Party
	Contest_MN_QSO_PARTY           Contest = "MN-QSO-PARTY"           // MN-QSO-PARTY         = Minnesota QSO Party
	Contest_MO_QSO_PARTY           Contest = "MO-QSO-PARTY"           // MO-QSO-PARTY         = Missouri QSO Party
	Contest_MS_QSO_PARTY           Contest = "MS-QSO-PARTY"           // MS-QSO-PARTY         = Mississippi QSO Party
	Contest_MT_QSO_PARTY           Contest = "MT-QSO-PARTY"           // MT-QSO-PARTY         = Montana QSO Party
	Contest_NA_SPRINT_CW           Contest = "NA-SPRINT-CW"           // NA-SPRINT-CW         = North America Sprint (CW)
	Contest_NA_SPRINT_RTTY         Contest = "NA-SPRINT-RTTY"         // NA-SPRINT-RTTY       = North America Sprint (RTTY)
	Contest_NA_SPRINT_SSB          Contest = "NA-SPRINT-SSB"          // NA-SPRINT-SSB        = North America Sprint (Phone)
	Contest_NAQP_CW                Contest = "NAQP-CW"                // NAQP-CW              = North America QSO Party (CW)
	Contest_NAQP_RTTY              Contest = "NAQP-RTTY"              // NAQP-RTTY            = North America QSO Party (RTTY)
	Contest_NAQP_SSB               Contest = "NAQP-SSB"               // NAQP-SSB             = North America QSO Party (Phone)
	Contest_NAVAL                  Contest = "NAVAL"                  // NAVAL                = International Naval Contest (INC)
	Contest_NC_QSO_PARTY           Contest = "NC-QSO-PARTY"           // NC-QSO-PARTY         = North Carolina QSO Party
	Contest_ND_QSO_PARTY           Contest = "ND-QSO-PARTY"           // ND-QSO-PARTY         = North Dakota QSO Party
	Contest_NE_QSO_PARTY           Contest = "NE-QSO-PARTY"           // NE-QSO-PARTY         = Nebraska QSO Party
	Contest_NEQP                   Contest = "NEQP"                   // NEQP                 = New England QSO Party
	Contest_NH_QSO_PARTY           Contest = "NH-QSO-PARTY"           // NH-QSO-PARTY         = New Hampshire QSO Party
	Contest_NJ_QSO_PARTY           Contest = "NJ-QSO-PARTY"           // NJ-QSO-PARTY         = New Jersey QSO Party
	Contest_NM_QSO_PARTY           Contest = "NM-QSO-PARTY"           // NM-QSO-PARTY         = New Mexico QSO Party
	Contest_NRAU_BALTIC_CW         Contest = "NRAU-BALTIC-CW"         // NRAU-BALTIC-CW       = NRAU-Baltic Contest (CW)
	Contest_NRAU_BALTIC_SSB        Contest = "NRAU-BALTIC-SSB"        // NRAU-BALTIC-SSB      = NRAU-Baltic Contest (SSB)
	Contest_NV_QSO_PARTY           Contest = "NV-QSO-PARTY"           // NV-QSO-PARTY         = Nevada QSO Party
	Contest_NY_QSO_PARTY           Contest = "NY-QSO-PARTY"           // NY-QSO-PARTY         = New York QSO Party
	Contest_OCEANIA_DX_CW          Contest = "OCEANIA-DX-CW"          // OCEANIA-DX-CW        = Oceania DX Contest (CW)
	Contest_OCEANIA_DX_SSB         Contest = "OCEANIA-DX-SSB"         // OCEANIA-DX-SSB       = Oceania DX Contest (SSB)
	Contest_OH_QSO_PARTY           Contest = "OH-QSO-PARTY"           // OH-QSO-PARTY         = Ohio QSO Party
	Contest_OK_DX_RTTY             Contest = "OK-DX-RTTY"             // OK-DX-RTTY           = Czech Radio Club OK DX Contest
	Contest_OK_OM_DX               Contest = "OK-OM-DX"               // OK-OM-DX             = Czech Radio Club OK-OM DX Contest
	Contest_OK_QSO_PARTY           Contest = "OK-QSO-PARTY"           // OK-QSO-PARTY         = Oklahoma QSO Party
	Contest_OMISS_QSO_PARTY        Contest = "OMISS-QSO-PARTY"        // OMISS-QSO-PARTY      = Old Man International Sideband Society QSO Party
	Contest_ON_QSO_PARTY           Contest = "ON-QSO-PARTY"           // ON-QSO-PARTY         = Ontario QSO Party
	Contest_OR_QSO_PARTY           Contest = "OR-QSO-PARTY"           // OR-QSO-PARTY         = Oregon QSO Party
	Contest_ORARI_DX               Contest = "ORARI-DX"               // ORARI-DX             = ORARI DX Contest
	Contest_PA_QSO_PARTY           Contest = "PA-QSO-PARTY"           // PA-QSO-PARTY         = Pennsylvania QSO Party
	Contest_PACC                   Contest = "PACC"                   // PACC                 = Dutch PACC Contest
	Contest_PCC                    Contest = "PCC"                    // PCC                  = PCCPro CW Contest
	Contest_PSK_DEATHMATCH         Contest = "PSK-DEATHMATCH"         // PSK-DEATHMATCH       = MDXA PSK DeathMatch (2005-2010)
	Contest_QC_QSO_PARTY           Contest = "QC-QSO-PARTY"           // QC-QSO-PARTY         = Quebec QSO Party
	Contest_RAC                    Contest = "RAC"                    // Deprecated: RAC                  = Canadian Amateur Radio Society Contest
	Contest_RAC_CANADA_DAY         Contest = "RAC-CANADA-DAY"         // RAC-CANADA-DAY       = RAC Canada Day Contest
	Contest_RAC_CANADA_WINTER      Contest = "RAC-CANADA-WINTER"      // RAC-CANADA-WINTER    = RAC Canada Winter Contest
	Contest_RDAC                   Contest = "RDAC"                   // RDAC                 = Russian District Award Contest
	Contest_RDXC                   Contest = "RDXC"                   // RDXC                 = Russian DX Contest
	Contest_REF_160M               Contest = "REF-160M"               // REF-160M             = Reseau des Emetteurs Francais 160m Contest
	Contest_REF_CW                 Contest = "REF-CW"                 // REF-CW               = Reseau des Emetteurs Francais Contest (CW)
	Contest_REF_SSB                Contest = "REF-SSB"                // REF-SSB              = Reseau des Emetteurs Francais Contest (SSB)
	Contest_REP_PORTUGAL_DAY_HF    Contest = "REP-PORTUGAL-DAY-HF"    // REP-PORTUGAL-DAY-HF  = Rede dos Emissores Portugueses Portugal Day HF Contest
	Contest_RI_QSO_PARTY           Contest = "RI-QSO-PARTY"           // RI-QSO-PARTY         = Rhode Island QSO Party
	Contest_RSGB_160               Contest = "RSGB-160"               // RSGB-160             = 1.8MHz Contest
	Contest_RSGB_21_28_CW          Contest = "RSGB-21/28-CW"          // RSGB-21/28-CW        = 21/28 MHz Contest (CW)
	Contest_RSGB_21_28_SSB         Contest = "RSGB-21/28-SSB"         // RSGB-21/28-SSB       = 21/28 MHz Contest (SSB)
	Contest_RSGB_80M_CC            Contest = "RSGB-80M-CC"            // RSGB-80M-CC          = 80m Club Championships
	Contest_RSGB_AFS_CW            Contest = "RSGB-AFS-CW"            // RSGB-AFS-CW          = Affiliated Societies Team Contest (CW)
	Contest_RSGB_AFS_SSB           Contest = "RSGB-AFS-SSB"           // RSGB-AFS-SSB         = Affiliated Societies Team Contest (SSB)
	Contest_RSGB_CLUB_CALLS        Contest = "RSGB-CLUB-CALLS"        // RSGB-CLUB-CALLS      = Club Calls
	Contest_RSGB_COMMONWEALTH      Contest = "RSGB-COMMONWEALTH"      // RSGB-COMMONWEALTH    = Commonwealth Contest
	Contest_RSGB_IOTA              Contest = "RSGB-IOTA"              // RSGB-IOTA            = IOTA Contest
	Contest_RSGB_LOW_POWER         Contest = "RSGB-LOW-POWER"         // RSGB-LOW-POWER       = Low Power Field Day
	Contest_RSGB_NFD               Contest = "RSGB-NFD"               // RSGB-NFD             = National Field Day
	Contest_RSGB_ROPOCO            Contest = "RSGB-ROPOCO"            // RSGB-ROPOCO          = RoPoCo
	Contest_RSGB_SSB_FD            Contest = "RSGB-SSB-FD"            // RSGB-SSB-FD          = SSB Field Day
	Contest_RUSSIAN_RTTY           Contest = "RUSSIAN-RTTY"           // RUSSIAN-RTTY         = Russian Radio RTTY Worldwide Contest
	Contest_SAC_CW                 Contest = "SAC-CW"                 // SAC-CW               = Scandinavian Activity Contest (CW)
	Contest_SAC_SSB                Contest = "SAC-SSB"                // SAC-SSB              = Scandinavian Activity Contest (SSB)
	Contest_SARTG_RTTY             Contest = "SARTG-RTTY"             // SARTG-RTTY           = SARTG WW RTTY
	Contest_SC_QSO_PARTY           Contest = "SC-QSO-PARTY"           // SC-QSO-PARTY         = South Carolina QSO Party
	Contest_SCC_RTTY               Contest = "SCC-RTTY"               // SCC-RTTY             = SCC RTTY Championship
	Contest_SD_QSO_PARTY           Contest = "SD-QSO-PARTY"           // SD-QSO-PARTY         = South Dakota QSO Party
	Contest_SHORTRY                Contest = "SHORTRY"                // SHORTRY              = DARC RTTY Short Contest
	Contest_SMP_AUG                Contest = "SMP-AUG"                // SMP-AUG              = SSA Portabeltest
	Contest_SMP_MAY                Contest = "SMP-MAY"                // SMP-MAY              = SSA Portabeltest
	Contest_SP_DX_RTTY             Contest = "SP-DX-RTTY"             // SP-DX-RTTY           = PRC SPDX Contest (RTTY)
	Contest_SPAR_WINTER_FD         Contest = "SPAR-WINTER-FD"         // SPAR-WINTER-FD       = SPAR Winter Field Day(2016 and earlier)
	Contest_SPDXCONTEST            Contest = "SPDXCONTEST"            // SPDXCONTEST          = SP DX Contest
	Contest_SPRING_SPRINT          Contest = "SPRING SPRINT"          // SPRING SPRINT        = FISTS Spring Sprint
	Contest_SR_MARATHON            Contest = "SR-MARATHON"            // SR-MARATHON          = Scottish-Russian Marathon
	Contest_STEW_PERRY             Contest = "STEW-PERRY"             // STEW-PERRY           = Stew Perry Topband Distance Challenge
	Contest_SUMMER_SPRINT          Contest = "SUMMER SPRINT"          // SUMMER SPRINT        = FISTS Summer Sprint
	Contest_TARA_GRID_DIP          Contest = "TARA-GRID-DIP"          // TARA-GRID-DIP        = TARA Grid Dip PSK-RTTY Shindig
	Contest_TARA_RTTY              Contest = "TARA-RTTY"              // TARA-RTTY            = TARA RTTY Mêlée
	Contest_TARA_RUMBLE            Contest = "TARA-RUMBLE"            // TARA-RUMBLE          = TARA Rumble PSK Contest
	Contest_TARA_SKIRMISH          Contest = "TARA-SKIRMISH"          // TARA-SKIRMISH        = TARA Skirmish Digital Prefix Contest
	Contest_TEN_RTTY               Contest = "TEN-RTTY"               // TEN-RTTY             = Ten-Meter RTTY Contest (before 2011)
	Contest_TMC_RTTY               Contest = "TMC-RTTY"               // TMC-RTTY             = The Makrothen Contest
	Contest_TN_QSO_PARTY           Contest = "TN-QSO-PARTY"           // TN-QSO-PARTY         = Tennessee QSO Party
	Contest_TX_QSO_PARTY           Contest = "TX-QSO-PARTY"           // TX-QSO-PARTY         = Texas QSO Party
	Contest_UBA_DX_CW              Contest = "UBA-DX-CW"              // UBA-DX-CW            = UBA Contest (CW)
	Contest_UBA_DX_SSB             Contest = "UBA-DX-SSB"             // UBA-DX-SSB           = UBA Contest (SSB)
	Contest_UK_DX_BPSK63           Contest = "UK-DX-BPSK63"           // UK-DX-BPSK63         = European PSK Club BPSK63 Contest
	Contest_UK_DX_RTTY             Contest = "UK-DX-RTTY"             // UK-DX-RTTY           = UK DX RTTY Contest
	Contest_UKR_CHAMP_RTTY         Contest = "UKR-CHAMP-RTTY"         // UKR-CHAMP-RTTY       = Open Ukraine RTTY Championship
	Contest_UKRAINIAN_DX           Contest = "UKRAINIAN DX"           // UKRAINIAN DX         = Ukrainian DX
	Contest_UKSMG_6M_MARATHON      Contest = "UKSMG-6M-MARATHON"      // UKSMG-6M-MARATHON    = UKSMG 6m Marathon
	Contest_UKSMG_SUMMER_ES        Contest = "UKSMG-SUMMER-ES"        // UKSMG-SUMMER-ES      = UKSMG Summer Es Contest
	Contest_URE_DX                 Contest = "URE-DX"                 // Deprecated: URE-DX               = Ukrainian DX Contest
	Contest_US_COUNTIES_QSO        Contest = "US-COUNTIES-QSO"        // US-COUNTIES-QSO      = Mobile Amateur Awards Club
	Contest_UT_QSO_PARTY           Contest = "UT-QSO-PARTY"           // UT-QSO-PARTY         = Utah QSO Party
	Contest_VA_QSO_PARTY           Contest = "VA-QSO-PARTY"           // VA-QSO-PARTY         = Virginia QSO Party
	Contest_VENEZ_IND_DAY          Contest = "VENEZ-IND-DAY"          // VENEZ-IND-DAY        = RCV Venezuelan Independence Day Contest
	Contest_VIRGINIA_QSO_PARTY     Contest = "VIRGINIA QSO PARTY"     // Deprecated: VIRGINIA QSO PARTY   = Virginia QSO Party
	Contest_VOLTA_RTTY             Contest = "VOLTA-RTTY"             // VOLTA-RTTY           = Alessandro Volta RTTY DX Contest
	Contest_VT_QSO_PARTY           Contest = "VT-QSO-PARTY"           // VT-QSO-PARTY         = Vermont QSO Party
	Contest_WA_QSO_PARTY           Contest = "WA-QSO-PARTY"           // WA-QSO-PARTY         = Washington QSO Party
	Contest_WFD                    Contest = "WFD"                    // WFD                  = Winter Field Day (2017 and later)
	Contest_WI_QSO_PARTY           Contest = "WI-QSO-PARTY"           // WI-QSO-PARTY         = Wisconsin QSO Party
	Contest_WIA_HARRY_ANGEL        Contest = "WIA-HARRY ANGEL"        // WIA-HARRY ANGEL      = WIA Harry Angel Memorial 80m Sprint
	Contest_WIA_JMMFD              Contest = "WIA-JMMFD"              // WIA-JMMFD            = WIA John Moyle Memorial Field Day
	Contest_WIA_OCDX               Contest = "WIA-OCDX"               // WIA-OCDX             = WIA Oceania DX (OCDX) Contest
	Contest_WIA_REMEMBRANCE        Contest = "WIA-REMEMBRANCE"        // WIA-REMEMBRANCE      = WIA Remembrance Day
	Contest_WIA_ROSS_HULL          Contest = "WIA-ROSS HULL"          // WIA-ROSS HULL        = WIA Ross Hull Memorial VHF/UHF Contest
	Contest_WIA_TRANS_TASMAN       Contest = "WIA-TRANS TASMAN"       // WIA-TRANS TASMAN     = WIA Trans Tasman Low Bands Challenge
	Contest_WIA_VHF_UHF_FD         Contest = "WIA-VHF/UHF FD"         // WIA-VHF/UHF FD       = WIA VHF UHF Field Days
	Contest_WIA_VK_SHIRES          Contest = "WIA-VK SHIRES"          // WIA-VK SHIRES        = WIA VK Shires
	Contest_WINTER_SPRINT          Contest = "WINTER SPRINT"          // WINTER SPRINT        = FISTS Winter Sprint
	Contest_WV_QSO_PARTY           Contest = "WV-QSO-PARTY"           // WV-QSO-PARTY         = West Virginia QSO Party
	Contest_WW_DIGI                Contest = "WW-DIGI"                // WW-DIGI              = World Wide Digi DX Contest
	Contest_WY_QSO_PARTY           Contest = "WY-QSO-PARTY"           // WY-QSO-PARTY         = Wyoming QSO Party
	Contest_XE_INTL_RTTY           Contest = "XE-INTL-RTTY"           // XE-INTL-RTTY         = Mexico International Contest (RTTY)
	Contest_YOHFDX                 Contest = "YOHFDX"                 // YOHFDX               = YODX HF contest
	Contest_YUDXC                  Contest = "YUDXC"                  // YUDXC                = YU DX Contest
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Contest specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "070-160M-SPRINT", Description: "PODXS Great Pumpkin Sprint"},
	{IsImportOnly: false, Key: "070-3-DAY", Description: "PODXS Three Day Weekend"},
	{IsImportOnly: false, Key: "070-31-FLAVORS", Description: "PODXS 31 Flavors"},
	{IsImportOnly: false, Key: "070-40M-SPRINT", Description: "PODXS 40m Firecracker Sprint"},
	{IsImportOnly: false, Key: "070-80M-SPRINT", Description: "PODXS 80m Jay Hudak Memorial Sprint"},
	{IsImportOnly: false, Key: "070-PSKFEST", Description: "PODXS PSKFest"},
	{IsImportOnly: false, Key: "070-ST-PATS-DAY", Description: "PODXS St. Patricks Day"},
	{IsImportOnly: false, Key: "070-VALENTINE-SPRINT", Description: "PODXS Valentine Sprint"},
	{IsImportOnly: false, Key: "10-RTTY", Description: "Ten-Meter RTTY Contest (2011 onwards)"},
	{IsImportOnly: false, Key: "1010-OPEN-SEASON", Description: "Open Season Ten Meter QSO Party"},
	{IsImportOnly: false, Key: "7QP", Description: "7th-Area QSO Party"},
	{IsImportOnly: false, Key: "AL-QSO-PARTY", Description: "Alabama QSO Party"},
	{IsImportOnly: false, Key: "ALL-ASIAN-DX-CW", Description: "JARL All Asian DX Contest (CW)"},
	{IsImportOnly: false, Key: "ALL-ASIAN-DX-PHONE", Description: "JARL All Asian DX Contest (PHONE)"},
	{IsImportOnly: false, Key: "ANARTS-RTTY", Description: "ANARTS WW RTTY"},
	{IsImportOnly: false, Key: "ANATOLIAN-RTTY", Description: "Anatolian WW RTTY"},
	{IsImportOnly: false, Key: "AP-SPRINT", Description: "Asia - Pacific Sprint"},
	{IsImportOnly: false, Key: "AR-QSO-PARTY", Description: "Arkansas QSO Party"},
	{IsImportOnly: false, Key: "ARI-DX", Description: "ARI DX Contest"},
	{IsImportOnly: false, Key: "ARI-EME", Description: "ARI Italian EME Trophy"},
	{IsImportOnly: false, Key: "ARI-IAC-13CM", Description: "ARI Italian Activity Contest (13cm+)"},
	{IsImportOnly: false, Key: "ARI-IAC-23CM", Description: "ARI Italian Activity Contest (23cm)"},
	{IsImportOnly: false, Key: "ARI-IAC-6M", Description: "ARI Italian Activity Contest (6m)"},
	{IsImportOnly: false, Key: "ARI-IAC-UHF", Description: "ARI Italian Activity Contest (UHF)"},
	{IsImportOnly: false, Key: "ARI-IAC-VHF", Description: "ARI Italian Activity Contest (VHF)"},
	{IsImportOnly: false, Key: "ARRL-10", Description: "ARRL 10 Meter Contest"},
	{IsImportOnly: false, Key: "ARRL-10-GHZ", Description: "ARRL 10 GHz and Up Contest"},
	{IsImportOnly: false, Key: "ARRL-160", Description: "ARRL 160 Meter Contest"},
	{IsImportOnly: false, Key: "ARRL-222", Description: "ARRL 222 MHz and Up Distance Contest"},
	{IsImportOnly: false, Key: "ARRL-DIGI", Description: "ARRL International Digital Contest"},
	{IsImportOnly: false, Key: "ARRL-DX-CW", Description: "ARRL International DX Contest (CW)"},
	{IsImportOnly: false, Key: "ARRL-DX-SSB", Description: "ARRL International DX Contest (Phone)"},
	{IsImportOnly: false, Key: "ARRL-EME", Description: "ARRL EME contest"},
	{IsImportOnly: false, Key: "ARRL-FIELD-DAY", Description: "ARRL Field Day"},
	{IsImportOnly: false, Key: "ARRL-RR-CW", Description: "ARRL Rookie Roundup (CW)"},
	{IsImportOnly: false, Key: "ARRL-RR-RTTY", Description: "ARRL Rookie Roundup (RTTY)"},
	{IsImportOnly: false, Key: "ARRL-RR-SSB", Description: "ARRL Rookie Roundup (Phone)"},
	{IsImportOnly: false, Key: "ARRL-RTTY", Description: "ARRL RTTY Round-Up"},
	{IsImportOnly: false, Key: "ARRL-SCR", Description: "ARRL School Club Roundup"},
	{IsImportOnly: false, Key: "ARRL-SS-CW", Description: "ARRL November Sweepstakes (CW)"},
	{IsImportOnly: false, Key: "ARRL-SS-SSB", Description: "ARRL November Sweepstakes (Phone)"},
	{IsImportOnly: false, Key: "ARRL-UHF-AUG", Description: "ARRL August UHF Contest"},
	{IsImportOnly: false, Key: "ARRL-VHF-JAN", Description: "ARRL January VHF Sweepstakes"},
	{IsImportOnly: false, Key: "ARRL-VHF-JUN", Description: "ARRL June VHF QSO Party"},
	{IsImportOnly: false, Key: "ARRL-VHF-SEP", Description: "ARRL September VHF QSO Party"},
	{IsImportOnly: false, Key: "AZ-QSO-PARTY", Description: "Arizona QSO Party"},
	{IsImportOnly: false, Key: "BANGGAI-DX", Description: "ORARI Banggai DX Contest"},
	{IsImportOnly: false, Key: "BARTG-RTTY", Description: "BARTG Spring RTTY Contest"},
	{IsImportOnly: false, Key: "BARTG-SPRINT", Description: "BARTG Sprint Contest"},
	{IsImportOnly: false, Key: "BC-QSO-PARTY", Description: "British Columbia QSO Party"},
	{IsImportOnly: false, Key: "BEKASI-MERDEKA-CONTEST", Description: "ORARI Bekasi Merdeka Contest"},
	{IsImportOnly: false, Key: "CA-QSO-PARTY", Description: "California QSO Party"},
	{IsImportOnly: false, Key: "CIS-DX", Description: "CIS DX Contest"},
	{IsImportOnly: false, Key: "CO-QSO-PARTY", Description: "Colorado QSO Party"},
	{IsImportOnly: false, Key: "CQ-160-CW", Description: "CQ WW 160 Meter DX Contest (CW)"},
	{IsImportOnly: false, Key: "CQ-160-SSB", Description: "CQ WW 160 Meter DX Contest (SSB)"},
	{IsImportOnly: false, Key: "CQ-M", Description: "CQ-M International DX Contest"},
	{IsImportOnly: false, Key: "CQ-VHF", Description: "CQ World-Wide VHF Contest"},
	{IsImportOnly: false, Key: "CQ-WPX-CW", Description: "CQ WW WPX Contest (CW)"},
	{IsImportOnly: false, Key: "CQ-WPX-RTTY", Description: "CQ/RJ WW RTTY WPX Contest"},
	{IsImportOnly: false, Key: "CQ-WPX-SSB", Description: "CQ WW WPX Contest (SSB)"},
	{IsImportOnly: false, Key: "CQ-WW-CW", Description: "CQ WW DX Contest (CW)"},
	{IsImportOnly: false, Key: "CQ-WW-RTTY", Description: "CQ/RJ WW RTTY DX Contest"},
	{IsImportOnly: false, Key: "CQ-WW-SSB", Description: "CQ WW DX Contest (SSB)"},
	{IsImportOnly: false, Key: "CT-QSO-PARTY", Description: "Connecticut QSO Party"},
	{IsImportOnly: false, Key: "CVA-DX-CW", Description: "Concurso Verde e Amarelo DX CW Contest"},
	{IsImportOnly: false, Key: "CVA-DX-SSB", Description: "Concurso Verde e Amarelo DX CW Contest"},
	{IsImportOnly: false, Key: "CWOPS-CW-OPEN", Description: "CWops CW Open Competition"},
	{IsImportOnly: false, Key: "CWOPS-CWT", Description: "CWops Mini-CWT Test"},
	{IsImportOnly: false, Key: "DARC-10", Description: "DARC 10m Contest"},
	{IsImportOnly: false, Key: "DARC-CWA", Description: "DARC CW Trainee Contest"},
	{IsImportOnly: false, Key: "DARC-FT4", Description: "DARC FT4 Contest"},
	{IsImportOnly: false, Key: "DARC-HELL", Description: "DARC Hell Contest"},
	{IsImportOnly: false, Key: "DARC-MICROWAVE", Description: "DARC Microwave Contest"},
	{IsImportOnly: false, Key: "DARC-TRAINEE", Description: "DARC Trainee Contest"},
	{IsImportOnly: false, Key: "DARC-UKW-FIELD-DAY", Description: "DARC UKW Summer Contest"},
	{IsImportOnly: false, Key: "DARC-UKW-SPRING", Description: "DARC UKW Spring Contest"},
	{IsImportOnly: false, Key: "DARC-VHF-UHF-MICROWAVE", Description: "DARC VHF-, UHF-, Microwave Contest (May)"},
	{IsImportOnly: false, Key: "DARC-WAEDC-CW", Description: "WAE DX Contest (CW)"},
	{IsImportOnly: false, Key: "DARC-WAEDC-RTTY", Description: "WAE DX Contest (RTTY)"},
	{IsImportOnly: false, Key: "DARC-WAEDC-SSB", Description: "WAE DX Contest (SSB)"},
	{IsImportOnly: false, Key: "DARC-WAG", Description: "DARC Worked All Germany"},
	{IsImportOnly: false, Key: "DE-QSO-PARTY", Description: "Delaware QSO Party"},
	{IsImportOnly: false, Key: "DL-DX-RTTY", Description: "DL-DX RTTY Contest"},
	{IsImportOnly: false, Key: "DMC-RTTY", Description: "DMC RTTY Contest"},
	{IsImportOnly: false, Key: "EA-CNCW", Description: "Concurso Nacional de Telegrafía"},
	{IsImportOnly: false, Key: "EA-DME", Description: "Municipios Españoles"},
	{IsImportOnly: false, Key: "EA-MAJESTAD-CW", Description: "His Majesty The King of Spain CW Contest (2022 and later)"},
	{IsImportOnly: false, Key: "EA-MAJESTAD-SSB", Description: "His Majesty The King of Spain SSB Contest (2022 and later)"},
	{IsImportOnly: false, Key: "EA-PSK63", Description: "EA PSK63"},
	{IsImportOnly: true, Key: "EA-RTTY", Description: "Unión de Radioaficionados Españoles RTTY Contest"},
	{IsImportOnly: false, Key: "EA-SMRE-CW", Description: "Su Majestad El Rey de España - CW (2021 and earlier)"},
	{IsImportOnly: false, Key: "EA-SMRE-SSB", Description: "Su Majestad El Rey de España - SSB (2021 and earlier)"},
	{IsImportOnly: false, Key: "EA-VHF-ATLANTIC", Description: "Atlántico V-UHF"},
	{IsImportOnly: false, Key: "EA-VHF-COM", Description: "Combinado de V-UHF"},
	{IsImportOnly: false, Key: "EA-VHF-COSTA-SOL", Description: "Costa del Sol V-UHF"},
	{IsImportOnly: false, Key: "EA-VHF-EA", Description: "Nacional VHF"},
	{IsImportOnly: false, Key: "EA-VHF-EA1RCS", Description: "Segovia EA1RCS V-UHF"},
	{IsImportOnly: false, Key: "EA-VHF-QSL", Description: "QSL V-UHF & 50MHz"},
	{IsImportOnly: false, Key: "EA-VHF-SADURNI", Description: "Sant Sadurni V-UHF"},
	{IsImportOnly: false, Key: "EA-WW-RTTY", Description: "Unión de Radioaficionados Españoles RTTY Contest"},
	{IsImportOnly: false, Key: "EASTER", Description: "DARC Easter Contest"},
	{IsImportOnly: false, Key: "EPC-PSK63", Description: "PSK63 QSO Party"},
	{IsImportOnly: false, Key: "EU SPRINT", Description: "EU Sprint"},
	{IsImportOnly: false, Key: "EU-HF", Description: "EU HF Championship"},
	{IsImportOnly: false, Key: "EU-PSK-DX", Description: "EU PSK DX Contest"},
	{IsImportOnly: false, Key: "EUCW160M", Description: "European CW Association 160m CW Party"},
	{IsImportOnly: false, Key: "FALL SPRINT", Description: "FISTS Fall Sprint"},
	{IsImportOnly: false, Key: "FL-QSO-PARTY", Description: "Florida QSO Party"},
	{IsImportOnly: false, Key: "GA-QSO-PARTY", Description: "Georgia QSO Party"},
	{IsImportOnly: false, Key: "HA-DX", Description: "Hungarian DX Contest"},
	{IsImportOnly: false, Key: "HELVETIA", Description: "Helvetia Contest"},
	{IsImportOnly: false, Key: "HI-QSO-PARTY", Description: "Hawaiian QSO Party"},
	{IsImportOnly: false, Key: "HOLYLAND", Description: "IARC Holyland Contest"},
	{IsImportOnly: false, Key: "IA-QSO-PARTY", Description: "Iowa QSO Party"},
	{IsImportOnly: false, Key: "IARU-FIELD-DAY", Description: "DARC IARU Region 1 Field Day"},
	{IsImportOnly: false, Key: "IARU-HF", Description: "IARU HF World Championship"},
	{IsImportOnly: false, Key: "ICWC-MST", Description: "ICWC Medium Speed Test"},
	{IsImportOnly: false, Key: "ID-QSO-PARTY", Description: "Idaho QSO Party"},
	{IsImportOnly: false, Key: "IL QSO PARTY", Description: "Illinois QSO Party"},
	{IsImportOnly: false, Key: "IN-QSO-PARTY", Description: "Indiana QSO Party"},
	{IsImportOnly: false, Key: "JARTS-WW-RTTY", Description: "JARTS WW RTTY"},
	{IsImportOnly: false, Key: "JIDX-CW", Description: "Japan International DX Contest (CW)"},
	{IsImportOnly: false, Key: "JIDX-SSB", Description: "Japan International DX Contest (SSB)"},
	{IsImportOnly: false, Key: "JT-DX-RTTY", Description: "Mongolian RTTY DX Contest"},
	{IsImportOnly: false, Key: "K1USN-SSO", Description: "K1USN Slow Speed Open"},
	{IsImportOnly: false, Key: "K1USN-SST", Description: "K1USN Slow Speed Test"},
	{IsImportOnly: false, Key: "KS-QSO-PARTY", Description: "Kansas QSO Party"},
	{IsImportOnly: false, Key: "KY-QSO-PARTY", Description: "Kentucky QSO Party"},
	{IsImportOnly: false, Key: "LA-QSO-PARTY", Description: "Louisiana QSO Party"},
	{IsImportOnly: false, Key: "LDC-RTTY", Description: "DRCG Long Distance Contest (RTTY)"},
	{IsImportOnly: false, Key: "LZ DX", Description: "LZ DX Contest"},
	{IsImportOnly: false, Key: "MAR-QSO-PARTY", Description: "Maritimes QSO Party"},
	{IsImportOnly: false, Key: "MD-QSO-PARTY", Description: "Maryland QSO Party"},
	{IsImportOnly: false, Key: "ME-QSO-PARTY", Description: "Maine QSO Party"},
	{IsImportOnly: false, Key: "MI-QSO-PARTY", Description: "Michigan QSO Party"},
	{IsImportOnly: false, Key: "MIDATLANTIC-QSO-PARTY", Description: "Mid-Atlantic QSO Party"},
	{IsImportOnly: false, Key: "MN-QSO-PARTY", Description: "Minnesota QSO Party"},
	{IsImportOnly: false, Key: "MO-QSO-PARTY", Description: "Missouri QSO Party"},
	{IsImportOnly: false, Key: "MS-QSO-PARTY", Description: "Mississippi QSO Party"},
	{IsImportOnly: false, Key: "MT-QSO-PARTY", Description: "Montana QSO Party"},
	{IsImportOnly: false, Key: "NA-SPRINT-CW", Description: "North America Sprint (CW)"},
	{IsImportOnly: false, Key: "NA-SPRINT-RTTY", Description: "North America Sprint (RTTY)"},
	{IsImportOnly: false, Key: "NA-SPRINT-SSB", Description: "North America Sprint (Phone)"},
	{IsImportOnly: false, Key: "NAQP-CW", Description: "North America QSO Party (CW)"},
	{IsImportOnly: false, Key: "NAQP-RTTY", Description: "North America QSO Party (RTTY)"},
	{IsImportOnly: false, Key: "NAQP-SSB", Description: "North America QSO Party (Phone)"},
	{IsImportOnly: false, Key: "NAVAL", Description: "International Naval Contest (INC)"},
	{IsImportOnly: false, Key: "NC-QSO-PARTY", Description: "North Carolina QSO Party"},
	{IsImportOnly: false, Key: "ND-QSO-PARTY", Description: "North Dakota QSO Party"},
	{IsImportOnly: false, Key: "NE-QSO-PARTY", Description: "Nebraska QSO Party"},
	{IsImportOnly: false, Key: "NEQP", Description: "New England QSO Party"},
	{IsImportOnly: false, Key: "NH-QSO-PARTY", Description: "New Hampshire QSO Party"},
	{IsImportOnly: false, Key: "NJ-QSO-PARTY", Description: "New Jersey QSO Party"},
	{IsImportOnly: false, Key: "NM-QSO-PARTY", Description: "New Mexico QSO Party"},
	{IsImportOnly: false, Key: "NRAU-BALTIC-CW", Description: "NRAU-Baltic Contest (CW)"},
	{IsImportOnly: false, Key: "NRAU-BALTIC-SSB", Description: "NRAU-Baltic Contest (SSB)"},
	{IsImportOnly: false, Key: "NV-QSO-PARTY", Description: "Nevada QSO Party"},
	{IsImportOnly: false, Key: "NY-QSO-PARTY", Description: "New York QSO Party"},
	{IsImportOnly: false, Key: "OCEANIA-DX-CW", Description: "Oceania DX Contest (CW)"},
	{IsImportOnly: false, Key: "OCEANIA-DX-SSB", Description: "Oceania DX Contest (SSB)"},
	{IsImportOnly: false, Key: "OH-QSO-PARTY", Description: "Ohio QSO Party"},
	{IsImportOnly: false, Key: "OK-DX-RTTY", Description: "Czech Radio Club OK DX Contest"},
	{IsImportOnly: false, Key: "OK-OM-DX", Description: "Czech Radio Club OK-OM DX Contest"},
	{IsImportOnly: false, Key: "OK-QSO-PARTY", Description: "Oklahoma QSO Party"},
	{IsImportOnly: false, Key: "OMISS-QSO-PARTY", Description: "Old Man International Sideband Society QSO Party"},
	{IsImportOnly: false, Key: "ON-QSO-PARTY", Description: "Ontario QSO Party"},
	{IsImportOnly: false, Key: "OR-QSO-PARTY", Description: "Oregon QSO Party"},
	{IsImportOnly: false, Key: "ORARI-DX", Description: "ORARI DX Contest"},
	{IsImportOnly: false, Key: "PA-QSO-PARTY", Description: "Pennsylvania QSO Party"},
	{IsImportOnly: false, Key: "PACC", Description: "Dutch PACC Contest"},
	{IsImportOnly: false, Key: "PCC", Description: "PCCPro CW Contest"},
	{IsImportOnly: false, Key: "PSK-DEATHMATCH", Description: "MDXA PSK DeathMatch (2005-2010)"},
	{IsImportOnly: false, Key: "QC-QSO-PARTY", Description: "Quebec QSO Party"},
	{IsImportOnly: true, Key: "RAC", Description: "Canadian Amateur Radio Society Contest"},
	{IsImportOnly: false, Key: "RAC-CANADA-DAY", Description: "RAC Canada Day Contest"},
	{IsImportOnly: false, Key: "RAC-CANADA-WINTER", Description: "RAC Canada Winter Contest"},
	{IsImportOnly: false, Key: "RDAC", Description: "Russian District Award Contest"},
	{IsImportOnly: false, Key: "RDXC", Description: "Russian DX Contest"},
	{IsImportOnly: false, Key: "REF-160M", Description: "Reseau des Emetteurs Francais 160m Contest"},
	{IsImportOnly: false, Key: "REF-CW", Description: "Reseau des Emetteurs Francais Contest (CW)"},
	{IsImportOnly: false, Key: "REF-SSB", Description: "Reseau des Emetteurs Francais Contest (SSB)"},
	{IsImportOnly: false, Key: "REP-PORTUGAL-DAY-HF", Description: "Rede dos Emissores Portugueses Portugal Day HF Contest"},
	{IsImportOnly: false, Key: "RI-QSO-PARTY", Description: "Rhode Island QSO Party"},
	{IsImportOnly: false, Key: "RSGB-160", Description: "1.8MHz Contest"},
	{IsImportOnly: false, Key: "RSGB-21/28-CW", Description: "21/28 MHz Contest (CW)"},
	{IsImportOnly: false, Key: "RSGB-21/28-SSB", Description: "21/28 MHz Contest (SSB)"},
	{IsImportOnly: false, Key: "RSGB-80M-CC", Description: "80m Club Championships"},
	{IsImportOnly: false, Key: "RSGB-AFS-CW", Description: "Affiliated Societies Team Contest (CW)"},
	{IsImportOnly: false, Key: "RSGB-AFS-SSB", Description: "Affiliated Societies Team Contest (SSB)"},
	{IsImportOnly: false, Key: "RSGB-CLUB-CALLS", Description: "Club Calls"},
	{IsImportOnly: false, Key: "RSGB-COMMONWEALTH", Description: "Commonwealth Contest"},
	{IsImportOnly: false, Key: "RSGB-IOTA", Description: "IOTA Contest"},
	{IsImportOnly: false, Key: "RSGB-LOW-POWER", Description: "Low Power Field Day"},
	{IsImportOnly: false, Key: "RSGB-NFD", Description: "National Field Day"},
	{IsImportOnly: false, Key: "RSGB-ROPOCO", Description: "RoPoCo"},
	{IsImportOnly: false, Key: "RSGB-SSB-FD", Description: "SSB Field Day"},
	{IsImportOnly: false, Key: "RUSSIAN-RTTY", Description: "Russian Radio RTTY Worldwide Contest"},
	{IsImportOnly: false, Key: "SAC-CW", Description: "Scandinavian Activity Contest (CW)"},
	{IsImportOnly: false, Key: "SAC-SSB", Description: "Scandinavian Activity Contest (SSB)"},
	{IsImportOnly: false, Key: "SARTG-RTTY", Description: "SARTG WW RTTY"},
	{IsImportOnly: false, Key: "SC-QSO-PARTY", Description: "South Carolina QSO Party"},
	{IsImportOnly: false, Key: "SCC-RTTY", Description: "SCC RTTY Championship"},
	{IsImportOnly: false, Key: "SD-QSO-PARTY", Description: "South Dakota QSO Party"},
	{IsImportOnly: false, Key: "SHORTRY", Description: "DARC RTTY Short Contest"},
	{IsImportOnly: false, Key: "SMP-AUG", Description: "SSA Portabeltest"},
	{IsImportOnly: false, Key: "SMP-MAY", Description: "SSA Portabeltest"},
	{IsImportOnly: false, Key: "SP-DX-RTTY", Description: "PRC SPDX Contest (RTTY)"},
	{IsImportOnly: false, Key: "SPAR-WINTER-FD", Description: "SPAR Winter Field Day(2016 and earlier)"},
	{IsImportOnly: false, Key: "SPDXCONTEST", Description: "SP DX Contest"},
	{IsImportOnly: false, Key: "SPRING SPRINT", Description: "FISTS Spring Sprint"},
	{IsImportOnly: false, Key: "SR-MARATHON", Description: "Scottish-Russian Marathon"},
	{IsImportOnly: false, Key: "STEW-PERRY", Description: "Stew Perry Topband Distance Challenge"},
	{IsImportOnly: false, Key: "SUMMER SPRINT", Description: "FISTS Summer Sprint"},
	{IsImportOnly: false, Key: "TARA-GRID-DIP", Description: "TARA Grid Dip PSK-RTTY Shindig"},
	{IsImportOnly: false, Key: "TARA-RTTY", Description: "TARA RTTY Mêlée"},
	{IsImportOnly: false, Key: "TARA-RUMBLE", Description: "TARA Rumble PSK Contest"},
	{IsImportOnly: false, Key: "TARA-SKIRMISH", Description: "TARA Skirmish Digital Prefix Contest"},
	{IsImportOnly: false, Key: "TEN-RTTY", Description: "Ten-Meter RTTY Contest (before 2011)"},
	{IsImportOnly: false, Key: "TMC-RTTY", Description: "The Makrothen Contest"},
	{IsImportOnly: false, Key: "TN-QSO-PARTY", Description: "Tennessee QSO Party"},
	{IsImportOnly: false, Key: "TX-QSO-PARTY", Description: "Texas QSO Party"},
	{IsImportOnly: false, Key: "UBA-DX-CW", Description: "UBA Contest (CW)"},
	{IsImportOnly: false, Key: "UBA-DX-SSB", Description: "UBA Contest (SSB)"},
	{IsImportOnly: false, Key: "UK-DX-BPSK63", Description: "European PSK Club BPSK63 Contest"},
	{IsImportOnly: false, Key: "UK-DX-RTTY", Description: "UK DX RTTY Contest"},
	{IsImportOnly: false, Key: "UKR-CHAMP-RTTY", Description: "Open Ukraine RTTY Championship"},
	{IsImportOnly: false, Key: "UKRAINIAN DX", Description: "Ukrainian DX"},
	{IsImportOnly: false, Key: "UKSMG-6M-MARATHON", Description: "UKSMG 6m Marathon"},
	{IsImportOnly: false, Key: "UKSMG-SUMMER-ES", Description: "UKSMG Summer Es Contest"},
	{IsImportOnly: true, Key: "URE-DX", Description: "Ukrainian DX Contest"},
	{IsImportOnly: false, Key: "US-COUNTIES-QSO", Description: "Mobile Amateur Awards Club"},
	{IsImportOnly: false, Key: "UT-QSO-PARTY", Description: "Utah QSO Party"},
	{IsImportOnly: false, Key: "VA-QSO-PARTY", Description: "Virginia QSO Party"},
	{IsImportOnly: false, Key: "VENEZ-IND-DAY", Description: "RCV Venezuelan Independence Day Contest"},
	{IsImportOnly: true, Key: "VIRGINIA QSO PARTY", Description: "Virginia QSO Party"},
	{IsImportOnly: false, Key: "VOLTA-RTTY", Description: "Alessandro Volta RTTY DX Contest"},
	{IsImportOnly: false, Key: "VT-QSO-PARTY", Description: "Vermont QSO Party"},
	{IsImportOnly: false, Key: "WA-QSO-PARTY", Description: "Washington QSO Party"},
	{IsImportOnly: false, Key: "WFD", Description: "Winter Field Day (2017 and later)"},
	{IsImportOnly: false, Key: "WI-QSO-PARTY", Description: "Wisconsin QSO Party"},
	{IsImportOnly: false, Key: "WIA-HARRY ANGEL", Description: "WIA Harry Angel Memorial 80m Sprint"},
	{IsImportOnly: false, Key: "WIA-JMMFD", Description: "WIA John Moyle Memorial Field Day"},
	{IsImportOnly: false, Key: "WIA-OCDX", Description: "WIA Oceania DX (OCDX) Contest"},
	{IsImportOnly: false, Key: "WIA-REMEMBRANCE", Description: "WIA Remembrance Day"},
	{IsImportOnly: false, Key: "WIA-ROSS HULL", Description: "WIA Ross Hull Memorial VHF/UHF Contest"},
	{IsImportOnly: false, Key: "WIA-TRANS TASMAN", Description: "WIA Trans Tasman Low Bands Challenge"},
	{IsImportOnly: false, Key: "WIA-VHF/UHF FD", Description: "WIA VHF UHF Field Days"},
	{IsImportOnly: false, Key: "WIA-VK SHIRES", Description: "WIA VK Shires"},
	{IsImportOnly: false, Key: "WINTER SPRINT", Description: "FISTS Winter Sprint"},
	{IsImportOnly: false, Key: "WV-QSO-PARTY", Description: "West Virginia QSO Party"},
	{IsImportOnly: false, Key: "WW-DIGI", Description: "World Wide Digi DX Contest"},
	{IsImportOnly: false, Key: "WY-QSO-PARTY", Description: "Wyoming QSO Party"},
	{IsImportOnly: false, Key: "XE-INTL-RTTY", Description: "Mexico International Contest (RTTY)"},
	{IsImportOnly: false, Key: "YOHFDX", Description: "YODX HF contest"},
	{IsImportOnly: false, Key: "YUDXC", Description: "YU DX Contest"},
}

// lookupMap contains all known Contest specifications.
var lookupMap = map[Contest]*Spec{
	Contest_070_160M_SPRINT:        &lookupList[0],
	Contest_070_3_DAY:              &lookupList[1],
	Contest_070_31_FLAVORS:         &lookupList[2],
	Contest_070_40M_SPRINT:         &lookupList[3],
	Contest_070_80M_SPRINT:         &lookupList[4],
	Contest_070_PSKFEST:            &lookupList[5],
	Contest_070_ST_PATS_DAY:        &lookupList[6],
	Contest_070_VALENTINE_SPRINT:   &lookupList[7],
	Contest_10_RTTY:                &lookupList[8],
	Contest_1010_OPEN_SEASON:       &lookupList[9],
	Contest_7QP:                    &lookupList[10],
	Contest_AL_QSO_PARTY:           &lookupList[11],
	Contest_ALL_ASIAN_DX_CW:        &lookupList[12],
	Contest_ALL_ASIAN_DX_PHONE:     &lookupList[13],
	Contest_ANARTS_RTTY:            &lookupList[14],
	Contest_ANATOLIAN_RTTY:         &lookupList[15],
	Contest_AP_SPRINT:              &lookupList[16],
	Contest_AR_QSO_PARTY:           &lookupList[17],
	Contest_ARI_DX:                 &lookupList[18],
	Contest_ARI_EME:                &lookupList[19],
	Contest_ARI_IAC_13CM:           &lookupList[20],
	Contest_ARI_IAC_23CM:           &lookupList[21],
	Contest_ARI_IAC_6M:             &lookupList[22],
	Contest_ARI_IAC_UHF:            &lookupList[23],
	Contest_ARI_IAC_VHF:            &lookupList[24],
	Contest_ARRL_10:                &lookupList[25],
	Contest_ARRL_10_GHZ:            &lookupList[26],
	Contest_ARRL_160:               &lookupList[27],
	Contest_ARRL_222:               &lookupList[28],
	Contest_ARRL_DIGI:              &lookupList[29],
	Contest_ARRL_DX_CW:             &lookupList[30],
	Contest_ARRL_DX_SSB:            &lookupList[31],
	Contest_ARRL_EME:               &lookupList[32],
	Contest_ARRL_FIELD_DAY:         &lookupList[33],
	Contest_ARRL_RR_CW:             &lookupList[34],
	Contest_ARRL_RR_RTTY:           &lookupList[35],
	Contest_ARRL_RR_SSB:            &lookupList[36],
	Contest_ARRL_RTTY:              &lookupList[37],
	Contest_ARRL_SCR:               &lookupList[38],
	Contest_ARRL_SS_CW:             &lookupList[39],
	Contest_ARRL_SS_SSB:            &lookupList[40],
	Contest_ARRL_UHF_AUG:           &lookupList[41],
	Contest_ARRL_VHF_JAN:           &lookupList[42],
	Contest_ARRL_VHF_JUN:           &lookupList[43],
	Contest_ARRL_VHF_SEP:           &lookupList[44],
	Contest_AZ_QSO_PARTY:           &lookupList[45],
	Contest_BANGGAI_DX:             &lookupList[46],
	Contest_BARTG_RTTY:             &lookupList[47],
	Contest_BARTG_SPRINT:           &lookupList[48],
	Contest_BC_QSO_PARTY:           &lookupList[49],
	Contest_BEKASI_MERDEKA_CONTEST: &lookupList[50],
	Contest_CA_QSO_PARTY:           &lookupList[51],
	Contest_CIS_DX:                 &lookupList[52],
	Contest_CO_QSO_PARTY:           &lookupList[53],
	Contest_CQ_160_CW:              &lookupList[54],
	Contest_CQ_160_SSB:             &lookupList[55],
	Contest_CQ_M:                   &lookupList[56],
	Contest_CQ_VHF:                 &lookupList[57],
	Contest_CQ_WPX_CW:              &lookupList[58],
	Contest_CQ_WPX_RTTY:            &lookupList[59],
	Contest_CQ_WPX_SSB:             &lookupList[60],
	Contest_CQ_WW_CW:               &lookupList[61],
	Contest_CQ_WW_RTTY:             &lookupList[62],
	Contest_CQ_WW_SSB:              &lookupList[63],
	Contest_CT_QSO_PARTY:           &lookupList[64],
	Contest_CVA_DX_CW:              &lookupList[65],
	Contest_CVA_DX_SSB:             &lookupList[66],
	Contest_CWOPS_CW_OPEN:          &lookupList[67],
	Contest_CWOPS_CWT:              &lookupList[68],
	Contest_DARC_10:                &lookupList[69],
	Contest_DARC_CWA:               &lookupList[70],
	Contest_DARC_FT4:               &lookupList[71],
	Contest_DARC_HELL:              &lookupList[72],
	Contest_DARC_MICROWAVE:         &lookupList[73],
	Contest_DARC_TRAINEE:           &lookupList[74],
	Contest_DARC_UKW_FIELD_DAY:     &lookupList[75],
	Contest_DARC_UKW_SPRING:        &lookupList[76],
	Contest_DARC_VHF_UHF_MICROWAVE: &lookupList[77],
	Contest_DARC_WAEDC_CW:          &lookupList[78],
	Contest_DARC_WAEDC_RTTY:        &lookupList[79],
	Contest_DARC_WAEDC_SSB:         &lookupList[80],
	Contest_DARC_WAG:               &lookupList[81],
	Contest_DE_QSO_PARTY:           &lookupList[82],
	Contest_DL_DX_RTTY:             &lookupList[83],
	Contest_DMC_RTTY:               &lookupList[84],
	Contest_EA_CNCW:                &lookupList[85],
	Contest_EA_DME:                 &lookupList[86],
	Contest_EA_MAJESTAD_CW:         &lookupList[87],
	Contest_EA_MAJESTAD_SSB:        &lookupList[88],
	Contest_EA_PSK63:               &lookupList[89],
	Contest_EA_RTTY:                &lookupList[90],
	Contest_EA_SMRE_CW:             &lookupList[91],
	Contest_EA_SMRE_SSB:            &lookupList[92],
	Contest_EA_VHF_ATLANTIC:        &lookupList[93],
	Contest_EA_VHF_COM:             &lookupList[94],
	Contest_EA_VHF_COSTA_SOL:       &lookupList[95],
	Contest_EA_VHF_EA:              &lookupList[96],
	Contest_EA_VHF_EA1RCS:          &lookupList[97],
	Contest_EA_VHF_QSL:             &lookupList[98],
	Contest_EA_VHF_SADURNI:         &lookupList[99],
	Contest_EA_WW_RTTY:             &lookupList[100],
	Contest_EASTER:                 &lookupList[101],
	Contest_EPC_PSK63:              &lookupList[102],
	Contest_EU_SPRINT:              &lookupList[103],
	Contest_EU_HF:                  &lookupList[104],
	Contest_EU_PSK_DX:              &lookupList[105],
	Contest_EUCW160M:               &lookupList[106],
	Contest_FALL_SPRINT:            &lookupList[107],
	Contest_FL_QSO_PARTY:           &lookupList[108],
	Contest_GA_QSO_PARTY:           &lookupList[109],
	Contest_HA_DX:                  &lookupList[110],
	Contest_HELVETIA:               &lookupList[111],
	Contest_HI_QSO_PARTY:           &lookupList[112],
	Contest_HOLYLAND:               &lookupList[113],
	Contest_IA_QSO_PARTY:           &lookupList[114],
	Contest_IARU_FIELD_DAY:         &lookupList[115],
	Contest_IARU_HF:                &lookupList[116],
	Contest_ICWC_MST:               &lookupList[117],
	Contest_ID_QSO_PARTY:           &lookupList[118],
	Contest_IL_QSO_PARTY:           &lookupList[119],
	Contest_IN_QSO_PARTY:           &lookupList[120],
	Contest_JARTS_WW_RTTY:          &lookupList[121],
	Contest_JIDX_CW:                &lookupList[122],
	Contest_JIDX_SSB:               &lookupList[123],
	Contest_JT_DX_RTTY:             &lookupList[124],
	Contest_K1USN_SSO:              &lookupList[125],
	Contest_K1USN_SST:              &lookupList[126],
	Contest_KS_QSO_PARTY:           &lookupList[127],
	Contest_KY_QSO_PARTY:           &lookupList[128],
	Contest_LA_QSO_PARTY:           &lookupList[129],
	Contest_LDC_RTTY:               &lookupList[130],
	Contest_LZ_DX:                  &lookupList[131],
	Contest_MAR_QSO_PARTY:          &lookupList[132],
	Contest_MD_QSO_PARTY:           &lookupList[133],
	Contest_ME_QSO_PARTY:           &lookupList[134],
	Contest_MI_QSO_PARTY:           &lookupList[135],
	Contest_MIDATLANTIC_QSO_PARTY:  &lookupList[136],
	Contest_MN_QSO_PARTY:           &lookupList[137],
	Contest_MO_QSO_PARTY:           &lookupList[138],
	Contest_MS_QSO_PARTY:           &lookupList[139],
	Contest_MT_QSO_PARTY:           &lookupList[140],
	Contest_NA_SPRINT_CW:           &lookupList[141],
	Contest_NA_SPRINT_RTTY:         &lookupList[142],
	Contest_NA_SPRINT_SSB:          &lookupList[143],
	Contest_NAQP_CW:                &lookupList[144],
	Contest_NAQP_RTTY:              &lookupList[145],
	Contest_NAQP_SSB:               &lookupList[146],
	Contest_NAVAL:                  &lookupList[147],
	Contest_NC_QSO_PARTY:           &lookupList[148],
	Contest_ND_QSO_PARTY:           &lookupList[149],
	Contest_NE_QSO_PARTY:           &lookupList[150],
	Contest_NEQP:                   &lookupList[151],
	Contest_NH_QSO_PARTY:           &lookupList[152],
	Contest_NJ_QSO_PARTY:           &lookupList[153],
	Contest_NM_QSO_PARTY:           &lookupList[154],
	Contest_NRAU_BALTIC_CW:         &lookupList[155],
	Contest_NRAU_BALTIC_SSB:        &lookupList[156],
	Contest_NV_QSO_PARTY:           &lookupList[157],
	Contest_NY_QSO_PARTY:           &lookupList[158],
	Contest_OCEANIA_DX_CW:          &lookupList[159],
	Contest_OCEANIA_DX_SSB:         &lookupList[160],
	Contest_OH_QSO_PARTY:           &lookupList[161],
	Contest_OK_DX_RTTY:             &lookupList[162],
	Contest_OK_OM_DX:               &lookupList[163],
	Contest_OK_QSO_PARTY:           &lookupList[164],
	Contest_OMISS_QSO_PARTY:        &lookupList[165],
	Contest_ON_QSO_PARTY:           &lookupList[166],
	Contest_OR_QSO_PARTY:           &lookupList[167],
	Contest_ORARI_DX:               &lookupList[168],
	Contest_PA_QSO_PARTY:           &lookupList[169],
	Contest_PACC:                   &lookupList[170],
	Contest_PCC:                    &lookupList[171],
	Contest_PSK_DEATHMATCH:         &lookupList[172],
	Contest_QC_QSO_PARTY:           &lookupList[173],
	Contest_RAC:                    &lookupList[174],
	Contest_RAC_CANADA_DAY:         &lookupList[175],
	Contest_RAC_CANADA_WINTER:      &lookupList[176],
	Contest_RDAC:                   &lookupList[177],
	Contest_RDXC:                   &lookupList[178],
	Contest_REF_160M:               &lookupList[179],
	Contest_REF_CW:                 &lookupList[180],
	Contest_REF_SSB:                &lookupList[181],
	Contest_REP_PORTUGAL_DAY_HF:    &lookupList[182],
	Contest_RI_QSO_PARTY:           &lookupList[183],
	Contest_RSGB_160:               &lookupList[184],
	Contest_RSGB_21_28_CW:          &lookupList[185],
	Contest_RSGB_21_28_SSB:         &lookupList[186],
	Contest_RSGB_80M_CC:            &lookupList[187],
	Contest_RSGB_AFS_CW:            &lookupList[188],
	Contest_RSGB_AFS_SSB:           &lookupList[189],
	Contest_RSGB_CLUB_CALLS:        &lookupList[190],
	Contest_RSGB_COMMONWEALTH:      &lookupList[191],
	Contest_RSGB_IOTA:              &lookupList[192],
	Contest_RSGB_LOW_POWER:         &lookupList[193],
	Contest_RSGB_NFD:               &lookupList[194],
	Contest_RSGB_ROPOCO:            &lookupList[195],
	Contest_RSGB_SSB_FD:            &lookupList[196],
	Contest_RUSSIAN_RTTY:           &lookupList[197],
	Contest_SAC_CW:                 &lookupList[198],
	Contest_SAC_SSB:                &lookupList[199],
	Contest_SARTG_RTTY:             &lookupList[200],
	Contest_SC_QSO_PARTY:           &lookupList[201],
	Contest_SCC_RTTY:               &lookupList[202],
	Contest_SD_QSO_PARTY:           &lookupList[203],
	Contest_SHORTRY:                &lookupList[204],
	Contest_SMP_AUG:                &lookupList[205],
	Contest_SMP_MAY:                &lookupList[206],
	Contest_SP_DX_RTTY:             &lookupList[207],
	Contest_SPAR_WINTER_FD:         &lookupList[208],
	Contest_SPDXCONTEST:            &lookupList[209],
	Contest_SPRING_SPRINT:          &lookupList[210],
	Contest_SR_MARATHON:            &lookupList[211],
	Contest_STEW_PERRY:             &lookupList[212],
	Contest_SUMMER_SPRINT:          &lookupList[213],
	Contest_TARA_GRID_DIP:          &lookupList[214],
	Contest_TARA_RTTY:              &lookupList[215],
	Contest_TARA_RUMBLE:            &lookupList[216],
	Contest_TARA_SKIRMISH:          &lookupList[217],
	Contest_TEN_RTTY:               &lookupList[218],
	Contest_TMC_RTTY:               &lookupList[219],
	Contest_TN_QSO_PARTY:           &lookupList[220],
	Contest_TX_QSO_PARTY:           &lookupList[221],
	Contest_UBA_DX_CW:              &lookupList[222],
	Contest_UBA_DX_SSB:             &lookupList[223],
	Contest_UK_DX_BPSK63:           &lookupList[224],
	Contest_UK_DX_RTTY:             &lookupList[225],
	Contest_UKR_CHAMP_RTTY:         &lookupList[226],
	Contest_UKRAINIAN_DX:           &lookupList[227],
	Contest_UKSMG_6M_MARATHON:      &lookupList[228],
	Contest_UKSMG_SUMMER_ES:        &lookupList[229],
	Contest_URE_DX:                 &lookupList[230],
	Contest_US_COUNTIES_QSO:        &lookupList[231],
	Contest_UT_QSO_PARTY:           &lookupList[232],
	Contest_VA_QSO_PARTY:           &lookupList[233],
	Contest_VENEZ_IND_DAY:          &lookupList[234],
	Contest_VIRGINIA_QSO_PARTY:     &lookupList[235],
	Contest_VOLTA_RTTY:             &lookupList[236],
	Contest_VT_QSO_PARTY:           &lookupList[237],
	Contest_WA_QSO_PARTY:           &lookupList[238],
	Contest_WFD:                    &lookupList[239],
	Contest_WI_QSO_PARTY:           &lookupList[240],
	Contest_WIA_HARRY_ANGEL:        &lookupList[241],
	Contest_WIA_JMMFD:              &lookupList[242],
	Contest_WIA_OCDX:               &lookupList[243],
	Contest_WIA_REMEMBRANCE:        &lookupList[244],
	Contest_WIA_ROSS_HULL:          &lookupList[245],
	Contest_WIA_TRANS_TASMAN:       &lookupList[246],
	Contest_WIA_VHF_UHF_FD:         &lookupList[247],
	Contest_WIA_VK_SHIRES:          &lookupList[248],
	Contest_WINTER_SPRINT:          &lookupList[249],
	Contest_WV_QSO_PARTY:           &lookupList[250],
	Contest_WW_DIGI:                &lookupList[251],
	Contest_WY_QSO_PARTY:           &lookupList[252],
	Contest_XE_INTL_RTTY:           &lookupList[253],
	Contest_YOHFDX:                 &lookupList[254],
	Contest_YUDXC:                  &lookupList[255],
}

// Lookup returns the specification for the provided Contest.
// ADIF 3.1.6
func Lookup(contest Contest) (Spec, bool) {
	spec, ok := lookupMap[contest]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all Contest specifications that match the provided filter function.
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

// List returns all Contest specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns Contest specifications.
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
