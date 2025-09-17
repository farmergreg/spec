// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package contest provides code and constants as defined in ADIF 3.1.6 (Released)
package contest

import "sync"

const (
	CONTEST_070_160M_SPRINT        Contest = "070-160m-sprint"        // 070-160m-sprint      = PODXS Great Pumpkin Sprint
	CONTEST_070_3_DAY              Contest = "070-3-day"              // 070-3-day            = PODXS Three Day Weekend
	CONTEST_070_31_FLAVORS         Contest = "070-31-flavors"         // 070-31-flavors       = PODXS 31 Flavors
	CONTEST_070_40M_SPRINT         Contest = "070-40m-sprint"         // 070-40m-sprint       = PODXS 40m Firecracker Sprint
	CONTEST_070_80M_SPRINT         Contest = "070-80m-sprint"         // 070-80m-sprint       = PODXS 80m Jay Hudak Memorial Sprint
	CONTEST_070_PSKFEST            Contest = "070-pskfest"            // 070-pskfest          = PODXS PSKFest
	CONTEST_070_ST_PATS_DAY        Contest = "070-st-pats-day"        // 070-st-pats-day      = PODXS St. Patricks Day
	CONTEST_070_VALENTINE_SPRINT   Contest = "070-valentine-sprint"   // 070-valentine-sprint = PODXS Valentine Sprint
	CONTEST_10_RTTY                Contest = "10-rtty"                // 10-rtty              = Ten-Meter RTTY Contest (2011 onwards)
	CONTEST_1010_OPEN_SEASON       Contest = "1010-open-season"       // 1010-open-season     = Open Season Ten Meter QSO Party
	CONTEST_7QP                    Contest = "7qp"                    // 7qp                  = 7th-Area QSO Party
	CONTEST_AL_QSO_PARTY           Contest = "al-qso-party"           // al-qso-party         = Alabama QSO Party
	CONTEST_ALL_ASIAN_DX_CW        Contest = "all-asian-dx-cw"        // all-asian-dx-cw      = JARL All Asian DX Contest (CW)
	CONTEST_ALL_ASIAN_DX_PHONE     Contest = "all-asian-dx-phone"     // all-asian-dx-phone   = JARL All Asian DX Contest (PHONE)
	CONTEST_ANARTS_RTTY            Contest = "anarts-rtty"            // anarts-rtty          = ANARTS WW RTTY
	CONTEST_ANATOLIAN_RTTY         Contest = "anatolian-rtty"         // anatolian-rtty       = Anatolian WW RTTY
	CONTEST_AP_SPRINT              Contest = "ap-sprint"              // ap-sprint            = Asia - Pacific Sprint
	CONTEST_AR_QSO_PARTY           Contest = "ar-qso-party"           // ar-qso-party         = Arkansas QSO Party
	CONTEST_ARI_DX                 Contest = "ari-dx"                 // ari-dx               = ARI DX Contest
	CONTEST_ARI_EME                Contest = "ari-eme"                // ari-eme              = ARI Italian EME Trophy
	CONTEST_ARI_IAC_13CM           Contest = "ari-iac-13cm"           // ari-iac-13cm         = ARI Italian Activity Contest (13cm+)
	CONTEST_ARI_IAC_23CM           Contest = "ari-iac-23cm"           // ari-iac-23cm         = ARI Italian Activity Contest (23cm)
	CONTEST_ARI_IAC_6M             Contest = "ari-iac-6m"             // ari-iac-6m           = ARI Italian Activity Contest (6m)
	CONTEST_ARI_IAC_UHF            Contest = "ari-iac-uhf"            // ari-iac-uhf          = ARI Italian Activity Contest (UHF)
	CONTEST_ARI_IAC_VHF            Contest = "ari-iac-vhf"            // ari-iac-vhf          = ARI Italian Activity Contest (VHF)
	CONTEST_ARRL_10                Contest = "arrl-10"                // arrl-10              = ARRL 10 Meter Contest
	CONTEST_ARRL_10_GHZ            Contest = "arrl-10-ghz"            // arrl-10-ghz          = ARRL 10 GHz and Up Contest
	CONTEST_ARRL_160               Contest = "arrl-160"               // arrl-160             = ARRL 160 Meter Contest
	CONTEST_ARRL_222               Contest = "arrl-222"               // arrl-222             = ARRL 222 MHz and Up Distance Contest
	CONTEST_ARRL_DIGI              Contest = "arrl-digi"              // arrl-digi            = ARRL International Digital Contest
	CONTEST_ARRL_DX_CW             Contest = "arrl-dx-cw"             // arrl-dx-cw           = ARRL International DX Contest (CW)
	CONTEST_ARRL_DX_SSB            Contest = "arrl-dx-ssb"            // arrl-dx-ssb          = ARRL International DX Contest (Phone)
	CONTEST_ARRL_EME               Contest = "arrl-eme"               // arrl-eme             = ARRL EME contest
	CONTEST_ARRL_FIELD_DAY         Contest = "arrl-field-day"         // arrl-field-day       = ARRL Field Day
	CONTEST_ARRL_RR_CW             Contest = "arrl-rr-cw"             // arrl-rr-cw           = ARRL Rookie Roundup (CW)
	CONTEST_ARRL_RR_RTTY           Contest = "arrl-rr-rtty"           // arrl-rr-rtty         = ARRL Rookie Roundup (RTTY)
	CONTEST_ARRL_RR_SSB            Contest = "arrl-rr-ssb"            // arrl-rr-ssb          = ARRL Rookie Roundup (Phone)
	CONTEST_ARRL_RTTY              Contest = "arrl-rtty"              // arrl-rtty            = ARRL RTTY Round-Up
	CONTEST_ARRL_SCR               Contest = "arrl-scr"               // arrl-scr             = ARRL School Club Roundup
	CONTEST_ARRL_SS_CW             Contest = "arrl-ss-cw"             // arrl-ss-cw           = ARRL November Sweepstakes (CW)
	CONTEST_ARRL_SS_SSB            Contest = "arrl-ss-ssb"            // arrl-ss-ssb          = ARRL November Sweepstakes (Phone)
	CONTEST_ARRL_UHF_AUG           Contest = "arrl-uhf-aug"           // arrl-uhf-aug         = ARRL August UHF Contest
	CONTEST_ARRL_VHF_JAN           Contest = "arrl-vhf-jan"           // arrl-vhf-jan         = ARRL January VHF Sweepstakes
	CONTEST_ARRL_VHF_JUN           Contest = "arrl-vhf-jun"           // arrl-vhf-jun         = ARRL June VHF QSO Party
	CONTEST_ARRL_VHF_SEP           Contest = "arrl-vhf-sep"           // arrl-vhf-sep         = ARRL September VHF QSO Party
	CONTEST_AZ_QSO_PARTY           Contest = "az-qso-party"           // az-qso-party         = Arizona QSO Party
	CONTEST_BANGGAI_DX             Contest = "banggai-dx"             // banggai-dx           = ORARI Banggai DX Contest
	CONTEST_BARTG_RTTY             Contest = "bartg-rtty"             // bartg-rtty           = BARTG Spring RTTY Contest
	CONTEST_BARTG_SPRINT           Contest = "bartg-sprint"           // bartg-sprint         = BARTG Sprint Contest
	CONTEST_BC_QSO_PARTY           Contest = "bc-qso-party"           // bc-qso-party         = British Columbia QSO Party
	CONTEST_BEKASI_MERDEKA_CONTEST Contest = "bekasi-merdeka-contest" // bekasi-merdeka-contest = ORARI Bekasi Merdeka Contest
	CONTEST_CA_QSO_PARTY           Contest = "ca-qso-party"           // ca-qso-party         = California QSO Party
	CONTEST_CIS_DX                 Contest = "cis-dx"                 // cis-dx               = CIS DX Contest
	CONTEST_CO_QSO_PARTY           Contest = "co-qso-party"           // co-qso-party         = Colorado QSO Party
	CONTEST_CQ_160_CW              Contest = "cq-160-cw"              // cq-160-cw            = CQ WW 160 Meter DX Contest (CW)
	CONTEST_CQ_160_SSB             Contest = "cq-160-ssb"             // cq-160-ssb           = CQ WW 160 Meter DX Contest (SSB)
	CONTEST_CQ_M                   Contest = "cq-m"                   // cq-m                 = CQ-M International DX Contest
	CONTEST_CQ_VHF                 Contest = "cq-vhf"                 // cq-vhf               = CQ World-Wide VHF Contest
	CONTEST_CQ_WPX_CW              Contest = "cq-wpx-cw"              // cq-wpx-cw            = CQ WW WPX Contest (CW)
	CONTEST_CQ_WPX_RTTY            Contest = "cq-wpx-rtty"            // cq-wpx-rtty          = CQ/RJ WW RTTY WPX Contest
	CONTEST_CQ_WPX_SSB             Contest = "cq-wpx-ssb"             // cq-wpx-ssb           = CQ WW WPX Contest (SSB)
	CONTEST_CQ_WW_CW               Contest = "cq-ww-cw"               // cq-ww-cw             = CQ WW DX Contest (CW)
	CONTEST_CQ_WW_RTTY             Contest = "cq-ww-rtty"             // cq-ww-rtty           = CQ/RJ WW RTTY DX Contest
	CONTEST_CQ_WW_SSB              Contest = "cq-ww-ssb"              // cq-ww-ssb            = CQ WW DX Contest (SSB)
	CONTEST_CT_QSO_PARTY           Contest = "ct-qso-party"           // ct-qso-party         = Connecticut QSO Party
	CONTEST_CVA_DX_CW              Contest = "cva-dx-cw"              // cva-dx-cw            = Concurso Verde e Amarelo DX CW Contest
	CONTEST_CVA_DX_SSB             Contest = "cva-dx-ssb"             // cva-dx-ssb           = Concurso Verde e Amarelo DX CW Contest
	CONTEST_CWOPS_CW_OPEN          Contest = "cwops-cw-open"          // cwops-cw-open        = CWops CW Open Competition
	CONTEST_CWOPS_CWT              Contest = "cwops-cwt"              // cwops-cwt            = CWops Mini-CWT Test
	CONTEST_DARC_10                Contest = "darc-10"                // darc-10              = DARC 10m Contest
	CONTEST_DARC_CWA               Contest = "darc-cwa"               // darc-cwa             = DARC CW Trainee Contest
	CONTEST_DARC_FT4               Contest = "darc-ft4"               // darc-ft4             = DARC FT4 Contest
	CONTEST_DARC_HELL              Contest = "darc-hell"              // darc-hell            = DARC Hell Contest
	CONTEST_DARC_MICROWAVE         Contest = "darc-microwave"         // darc-microwave       = DARC Microwave Contest
	CONTEST_DARC_TRAINEE           Contest = "darc-trainee"           // darc-trainee         = DARC Trainee Contest
	CONTEST_DARC_UKW_FIELD_DAY     Contest = "darc-ukw-field-day"     // darc-ukw-field-day   = DARC UKW Summer Contest
	CONTEST_DARC_UKW_SPRING        Contest = "darc-ukw-spring"        // darc-ukw-spring      = DARC UKW Spring Contest
	CONTEST_DARC_VHF_UHF_MICROWAVE Contest = "darc-vhf-uhf-microwave" // darc-vhf-uhf-microwave = DARC VHF-, UHF-, Microwave Contest (May)
	CONTEST_DARC_WAEDC_CW          Contest = "darc-waedc-cw"          // darc-waedc-cw        = WAE DX Contest (CW)
	CONTEST_DARC_WAEDC_RTTY        Contest = "darc-waedc-rtty"        // darc-waedc-rtty      = WAE DX Contest (RTTY)
	CONTEST_DARC_WAEDC_SSB         Contest = "darc-waedc-ssb"         // darc-waedc-ssb       = WAE DX Contest (SSB)
	CONTEST_DARC_WAG               Contest = "darc-wag"               // darc-wag             = DARC Worked All Germany
	CONTEST_DE_QSO_PARTY           Contest = "de-qso-party"           // de-qso-party         = Delaware QSO Party
	CONTEST_DL_DX_RTTY             Contest = "dl-dx-rtty"             // dl-dx-rtty           = DL-DX RTTY Contest
	CONTEST_DMC_RTTY               Contest = "dmc-rtty"               // dmc-rtty             = DMC RTTY Contest
	CONTEST_EA_CNCW                Contest = "ea-cncw"                // ea-cncw              = Concurso Nacional de Telegrafía
	CONTEST_EA_DME                 Contest = "ea-dme"                 // ea-dme               = Municipios Españoles
	CONTEST_EA_MAJESTAD_CW         Contest = "ea-majestad-cw"         // ea-majestad-cw       = His Majesty The King of Spain CW Contest (2022 and later)
	CONTEST_EA_MAJESTAD_SSB        Contest = "ea-majestad-ssb"        // ea-majestad-ssb      = His Majesty The King of Spain SSB Contest (2022 and later)
	CONTEST_EA_PSK63               Contest = "ea-psk63"               // ea-psk63             = EA PSK63
	CONTEST_EA_RTTY                Contest = "ea-rtty"                // Deprecated: ea-rtty              = Unión de Radioaficionados Españoles RTTY Contest
	CONTEST_EA_SMRE_CW             Contest = "ea-smre-cw"             // ea-smre-cw           = Su Majestad El Rey de España - CW (2021 and earlier)
	CONTEST_EA_SMRE_SSB            Contest = "ea-smre-ssb"            // ea-smre-ssb          = Su Majestad El Rey de España - SSB (2021 and earlier)
	CONTEST_EA_VHF_ATLANTIC        Contest = "ea-vhf-atlantic"        // ea-vhf-atlantic      = Atlántico V-UHF
	CONTEST_EA_VHF_COM             Contest = "ea-vhf-com"             // ea-vhf-com           = Combinado de V-UHF
	CONTEST_EA_VHF_COSTA_SOL       Contest = "ea-vhf-costa-sol"       // ea-vhf-costa-sol     = Costa del Sol V-UHF
	CONTEST_EA_VHF_EA              Contest = "ea-vhf-ea"              // ea-vhf-ea            = Nacional VHF
	CONTEST_EA_VHF_EA1RCS          Contest = "ea-vhf-ea1rcs"          // ea-vhf-ea1rcs        = Segovia EA1RCS V-UHF
	CONTEST_EA_VHF_QSL             Contest = "ea-vhf-qsl"             // ea-vhf-qsl           = QSL V-UHF & 50MHz
	CONTEST_EA_VHF_SADURNI         Contest = "ea-vhf-sadurni"         // ea-vhf-sadurni       = Sant Sadurni V-UHF
	CONTEST_EA_WW_RTTY             Contest = "ea-ww-rtty"             // ea-ww-rtty           = Unión de Radioaficionados Españoles RTTY Contest
	CONTEST_EASTER                 Contest = "easter"                 // easter               = DARC Easter Contest
	CONTEST_EPC_PSK63              Contest = "epc-psk63"              // epc-psk63            = PSK63 QSO Party
	CONTEST_EU_SPRINT              Contest = "eu sprint"              // eu sprint            = EU Sprint
	CONTEST_EU_HF                  Contest = "eu-hf"                  // eu-hf                = EU HF Championship
	CONTEST_EU_PSK_DX              Contest = "eu-psk-dx"              // eu-psk-dx            = EU PSK DX Contest
	CONTEST_EUCW160M               Contest = "eucw160m"               // eucw160m             = European CW Association 160m CW Party
	CONTEST_FALL_SPRINT            Contest = "fall sprint"            // fall sprint          = FISTS Fall Sprint
	CONTEST_FL_QSO_PARTY           Contest = "fl-qso-party"           // fl-qso-party         = Florida QSO Party
	CONTEST_GA_QSO_PARTY           Contest = "ga-qso-party"           // ga-qso-party         = Georgia QSO Party
	CONTEST_HA_DX                  Contest = "ha-dx"                  // ha-dx                = Hungarian DX Contest
	CONTEST_HELVETIA               Contest = "helvetia"               // helvetia             = Helvetia Contest
	CONTEST_HI_QSO_PARTY           Contest = "hi-qso-party"           // hi-qso-party         = Hawaiian QSO Party
	CONTEST_HOLYLAND               Contest = "holyland"               // holyland             = IARC Holyland Contest
	CONTEST_IA_QSO_PARTY           Contest = "ia-qso-party"           // ia-qso-party         = Iowa QSO Party
	CONTEST_IARU_FIELD_DAY         Contest = "iaru-field-day"         // iaru-field-day       = DARC IARU Region 1 Field Day
	CONTEST_IARU_HF                Contest = "iaru-hf"                // iaru-hf              = IARU HF World Championship
	CONTEST_ICWC_MST               Contest = "icwc-mst"               // icwc-mst             = ICWC Medium Speed Test
	CONTEST_ID_QSO_PARTY           Contest = "id-qso-party"           // id-qso-party         = Idaho QSO Party
	CONTEST_IL_QSO_PARTY           Contest = "il qso party"           // il qso party         = Illinois QSO Party
	CONTEST_IN_QSO_PARTY           Contest = "in-qso-party"           // in-qso-party         = Indiana QSO Party
	CONTEST_JARTS_WW_RTTY          Contest = "jarts-ww-rtty"          // jarts-ww-rtty        = JARTS WW RTTY
	CONTEST_JIDX_CW                Contest = "jidx-cw"                // jidx-cw              = Japan International DX Contest (CW)
	CONTEST_JIDX_SSB               Contest = "jidx-ssb"               // jidx-ssb             = Japan International DX Contest (SSB)
	CONTEST_JT_DX_RTTY             Contest = "jt-dx-rtty"             // jt-dx-rtty           = Mongolian RTTY DX Contest
	CONTEST_K1USN_SSO              Contest = "k1usn-sso"              // k1usn-sso            = K1USN Slow Speed Open
	CONTEST_K1USN_SST              Contest = "k1usn-sst"              // k1usn-sst            = K1USN Slow Speed Test
	CONTEST_KS_QSO_PARTY           Contest = "ks-qso-party"           // ks-qso-party         = Kansas QSO Party
	CONTEST_KY_QSO_PARTY           Contest = "ky-qso-party"           // ky-qso-party         = Kentucky QSO Party
	CONTEST_LA_QSO_PARTY           Contest = "la-qso-party"           // la-qso-party         = Louisiana QSO Party
	CONTEST_LDC_RTTY               Contest = "ldc-rtty"               // ldc-rtty             = DRCG Long Distance Contest (RTTY)
	CONTEST_LZ_DX                  Contest = "lz dx"                  // lz dx                = LZ DX Contest
	CONTEST_MAR_QSO_PARTY          Contest = "mar-qso-party"          // mar-qso-party        = Maritimes QSO Party
	CONTEST_MD_QSO_PARTY           Contest = "md-qso-party"           // md-qso-party         = Maryland QSO Party
	CONTEST_ME_QSO_PARTY           Contest = "me-qso-party"           // me-qso-party         = Maine QSO Party
	CONTEST_MI_QSO_PARTY           Contest = "mi-qso-party"           // mi-qso-party         = Michigan QSO Party
	CONTEST_MIDATLANTIC_QSO_PARTY  Contest = "midatlantic-qso-party"  // midatlantic-qso-party = Mid-Atlantic QSO Party
	CONTEST_MN_QSO_PARTY           Contest = "mn-qso-party"           // mn-qso-party         = Minnesota QSO Party
	CONTEST_MO_QSO_PARTY           Contest = "mo-qso-party"           // mo-qso-party         = Missouri QSO Party
	CONTEST_MS_QSO_PARTY           Contest = "ms-qso-party"           // ms-qso-party         = Mississippi QSO Party
	CONTEST_MT_QSO_PARTY           Contest = "mt-qso-party"           // mt-qso-party         = Montana QSO Party
	CONTEST_NA_SPRINT_CW           Contest = "na-sprint-cw"           // na-sprint-cw         = North America Sprint (CW)
	CONTEST_NA_SPRINT_RTTY         Contest = "na-sprint-rtty"         // na-sprint-rtty       = North America Sprint (RTTY)
	CONTEST_NA_SPRINT_SSB          Contest = "na-sprint-ssb"          // na-sprint-ssb        = North America Sprint (Phone)
	CONTEST_NAQP_CW                Contest = "naqp-cw"                // naqp-cw              = North America QSO Party (CW)
	CONTEST_NAQP_RTTY              Contest = "naqp-rtty"              // naqp-rtty            = North America QSO Party (RTTY)
	CONTEST_NAQP_SSB               Contest = "naqp-ssb"               // naqp-ssb             = North America QSO Party (Phone)
	CONTEST_NAVAL                  Contest = "naval"                  // naval                = International Naval Contest (INC)
	CONTEST_NC_QSO_PARTY           Contest = "nc-qso-party"           // nc-qso-party         = North Carolina QSO Party
	CONTEST_ND_QSO_PARTY           Contest = "nd-qso-party"           // nd-qso-party         = North Dakota QSO Party
	CONTEST_NE_QSO_PARTY           Contest = "ne-qso-party"           // ne-qso-party         = Nebraska QSO Party
	CONTEST_NEQP                   Contest = "neqp"                   // neqp                 = New England QSO Party
	CONTEST_NH_QSO_PARTY           Contest = "nh-qso-party"           // nh-qso-party         = New Hampshire QSO Party
	CONTEST_NJ_QSO_PARTY           Contest = "nj-qso-party"           // nj-qso-party         = New Jersey QSO Party
	CONTEST_NM_QSO_PARTY           Contest = "nm-qso-party"           // nm-qso-party         = New Mexico QSO Party
	CONTEST_NRAU_BALTIC_CW         Contest = "nrau-baltic-cw"         // nrau-baltic-cw       = NRAU-Baltic Contest (CW)
	CONTEST_NRAU_BALTIC_SSB        Contest = "nrau-baltic-ssb"        // nrau-baltic-ssb      = NRAU-Baltic Contest (SSB)
	CONTEST_NV_QSO_PARTY           Contest = "nv-qso-party"           // nv-qso-party         = Nevada QSO Party
	CONTEST_NY_QSO_PARTY           Contest = "ny-qso-party"           // ny-qso-party         = New York QSO Party
	CONTEST_OCEANIA_DX_CW          Contest = "oceania-dx-cw"          // oceania-dx-cw        = Oceania DX Contest (CW)
	CONTEST_OCEANIA_DX_SSB         Contest = "oceania-dx-ssb"         // oceania-dx-ssb       = Oceania DX Contest (SSB)
	CONTEST_OH_QSO_PARTY           Contest = "oh-qso-party"           // oh-qso-party         = Ohio QSO Party
	CONTEST_OK_DX_RTTY             Contest = "ok-dx-rtty"             // ok-dx-rtty           = Czech Radio Club OK DX Contest
	CONTEST_OK_OM_DX               Contest = "ok-om-dx"               // ok-om-dx             = Czech Radio Club OK-OM DX Contest
	CONTEST_OK_QSO_PARTY           Contest = "ok-qso-party"           // ok-qso-party         = Oklahoma QSO Party
	CONTEST_OMISS_QSO_PARTY        Contest = "omiss-qso-party"        // omiss-qso-party      = Old Man International Sideband Society QSO Party
	CONTEST_ON_QSO_PARTY           Contest = "on-qso-party"           // on-qso-party         = Ontario QSO Party
	CONTEST_OR_QSO_PARTY           Contest = "or-qso-party"           // or-qso-party         = Oregon QSO Party
	CONTEST_ORARI_DX               Contest = "orari-dx"               // orari-dx             = ORARI DX Contest
	CONTEST_PA_QSO_PARTY           Contest = "pa-qso-party"           // pa-qso-party         = Pennsylvania QSO Party
	CONTEST_PACC                   Contest = "pacc"                   // pacc                 = Dutch PACC Contest
	CONTEST_PCC                    Contest = "pcc"                    // pcc                  = PCCPro CW Contest
	CONTEST_PSK_DEATHMATCH         Contest = "psk-deathmatch"         // psk-deathmatch       = MDXA PSK DeathMatch (2005-2010)
	CONTEST_QC_QSO_PARTY           Contest = "qc-qso-party"           // qc-qso-party         = Quebec QSO Party
	CONTEST_RAC                    Contest = "rac"                    // Deprecated: rac                  = Canadian Amateur Radio Society Contest
	CONTEST_RAC_CANADA_DAY         Contest = "rac-canada-day"         // rac-canada-day       = RAC Canada Day Contest
	CONTEST_RAC_CANADA_WINTER      Contest = "rac-canada-winter"      // rac-canada-winter    = RAC Canada Winter Contest
	CONTEST_RDAC                   Contest = "rdac"                   // rdac                 = Russian District Award Contest
	CONTEST_RDXC                   Contest = "rdxc"                   // rdxc                 = Russian DX Contest
	CONTEST_REF_160M               Contest = "ref-160m"               // ref-160m             = Reseau des Emetteurs Francais 160m Contest
	CONTEST_REF_CW                 Contest = "ref-cw"                 // ref-cw               = Reseau des Emetteurs Francais Contest (CW)
	CONTEST_REF_SSB                Contest = "ref-ssb"                // ref-ssb              = Reseau des Emetteurs Francais Contest (SSB)
	CONTEST_REP_PORTUGAL_DAY_HF    Contest = "rep-portugal-day-hf"    // rep-portugal-day-hf  = Rede dos Emissores Portugueses Portugal Day HF Contest
	CONTEST_RI_QSO_PARTY           Contest = "ri-qso-party"           // ri-qso-party         = Rhode Island QSO Party
	CONTEST_RSGB_160               Contest = "rsgb-160"               // rsgb-160             = 1.8MHz Contest
	CONTEST_RSGB_21_28_CW          Contest = "rsgb-21/28-cw"          // rsgb-21/28-cw        = 21/28 MHz Contest (CW)
	CONTEST_RSGB_21_28_SSB         Contest = "rsgb-21/28-ssb"         // rsgb-21/28-ssb       = 21/28 MHz Contest (SSB)
	CONTEST_RSGB_80M_CC            Contest = "rsgb-80m-cc"            // rsgb-80m-cc          = 80m Club Championships
	CONTEST_RSGB_AFS_CW            Contest = "rsgb-afs-cw"            // rsgb-afs-cw          = Affiliated Societies Team Contest (CW)
	CONTEST_RSGB_AFS_SSB           Contest = "rsgb-afs-ssb"           // rsgb-afs-ssb         = Affiliated Societies Team Contest (SSB)
	CONTEST_RSGB_CLUB_CALLS        Contest = "rsgb-club-calls"        // rsgb-club-calls      = Club Calls
	CONTEST_RSGB_COMMONWEALTH      Contest = "rsgb-commonwealth"      // rsgb-commonwealth    = Commonwealth Contest
	CONTEST_RSGB_IOTA              Contest = "rsgb-iota"              // rsgb-iota            = IOTA Contest
	CONTEST_RSGB_LOW_POWER         Contest = "rsgb-low-power"         // rsgb-low-power       = Low Power Field Day
	CONTEST_RSGB_NFD               Contest = "rsgb-nfd"               // rsgb-nfd             = National Field Day
	CONTEST_RSGB_ROPOCO            Contest = "rsgb-ropoco"            // rsgb-ropoco          = RoPoCo
	CONTEST_RSGB_SSB_FD            Contest = "rsgb-ssb-fd"            // rsgb-ssb-fd          = SSB Field Day
	CONTEST_RUSSIAN_RTTY           Contest = "russian-rtty"           // russian-rtty         = Russian Radio RTTY Worldwide Contest
	CONTEST_SAC_CW                 Contest = "sac-cw"                 // sac-cw               = Scandinavian Activity Contest (CW)
	CONTEST_SAC_SSB                Contest = "sac-ssb"                // sac-ssb              = Scandinavian Activity Contest (SSB)
	CONTEST_SARTG_RTTY             Contest = "sartg-rtty"             // sartg-rtty           = SARTG WW RTTY
	CONTEST_SC_QSO_PARTY           Contest = "sc-qso-party"           // sc-qso-party         = South Carolina QSO Party
	CONTEST_SCC_RTTY               Contest = "scc-rtty"               // scc-rtty             = SCC RTTY Championship
	CONTEST_SD_QSO_PARTY           Contest = "sd-qso-party"           // sd-qso-party         = South Dakota QSO Party
	CONTEST_SHORTRY                Contest = "shortry"                // shortry              = DARC RTTY Short Contest
	CONTEST_SMP_AUG                Contest = "smp-aug"                // smp-aug              = SSA Portabeltest
	CONTEST_SMP_MAY                Contest = "smp-may"                // smp-may              = SSA Portabeltest
	CONTEST_SP_DX_RTTY             Contest = "sp-dx-rtty"             // sp-dx-rtty           = PRC SPDX Contest (RTTY)
	CONTEST_SPAR_WINTER_FD         Contest = "spar-winter-fd"         // spar-winter-fd       = SPAR Winter Field Day(2016 and earlier)
	CONTEST_SPDXCONTEST            Contest = "spdxcontest"            // spdxcontest          = SP DX Contest
	CONTEST_SPRING_SPRINT          Contest = "spring sprint"          // spring sprint        = FISTS Spring Sprint
	CONTEST_SR_MARATHON            Contest = "sr-marathon"            // sr-marathon          = Scottish-Russian Marathon
	CONTEST_STEW_PERRY             Contest = "stew-perry"             // stew-perry           = Stew Perry Topband Distance Challenge
	CONTEST_SUMMER_SPRINT          Contest = "summer sprint"          // summer sprint        = FISTS Summer Sprint
	CONTEST_TARA_GRID_DIP          Contest = "tara-grid-dip"          // tara-grid-dip        = TARA Grid Dip PSK-RTTY Shindig
	CONTEST_TARA_RTTY              Contest = "tara-rtty"              // tara-rtty            = TARA RTTY Mêlée
	CONTEST_TARA_RUMBLE            Contest = "tara-rumble"            // tara-rumble          = TARA Rumble PSK Contest
	CONTEST_TARA_SKIRMISH          Contest = "tara-skirmish"          // tara-skirmish        = TARA Skirmish Digital Prefix Contest
	CONTEST_TEN_RTTY               Contest = "ten-rtty"               // ten-rtty             = Ten-Meter RTTY Contest (before 2011)
	CONTEST_TMC_RTTY               Contest = "tmc-rtty"               // tmc-rtty             = The Makrothen Contest
	CONTEST_TN_QSO_PARTY           Contest = "tn-qso-party"           // tn-qso-party         = Tennessee QSO Party
	CONTEST_TX_QSO_PARTY           Contest = "tx-qso-party"           // tx-qso-party         = Texas QSO Party
	CONTEST_UBA_DX_CW              Contest = "uba-dx-cw"              // uba-dx-cw            = UBA Contest (CW)
	CONTEST_UBA_DX_SSB             Contest = "uba-dx-ssb"             // uba-dx-ssb           = UBA Contest (SSB)
	CONTEST_UK_DX_BPSK63           Contest = "uk-dx-bpsk63"           // uk-dx-bpsk63         = European PSK Club BPSK63 Contest
	CONTEST_UK_DX_RTTY             Contest = "uk-dx-rtty"             // uk-dx-rtty           = UK DX RTTY Contest
	CONTEST_UKR_CHAMP_RTTY         Contest = "ukr-champ-rtty"         // ukr-champ-rtty       = Open Ukraine RTTY Championship
	CONTEST_UKRAINIAN_DX           Contest = "ukrainian dx"           // ukrainian dx         = Ukrainian DX
	CONTEST_UKSMG_6M_MARATHON      Contest = "uksmg-6m-marathon"      // uksmg-6m-marathon    = UKSMG 6m Marathon
	CONTEST_UKSMG_SUMMER_ES        Contest = "uksmg-summer-es"        // uksmg-summer-es      = UKSMG Summer Es Contest
	CONTEST_URE_DX                 Contest = "ure-dx"                 // Deprecated: ure-dx               = Ukrainian DX Contest
	CONTEST_US_COUNTIES_QSO        Contest = "us-counties-qso"        // us-counties-qso      = Mobile Amateur Awards Club
	CONTEST_UT_QSO_PARTY           Contest = "ut-qso-party"           // ut-qso-party         = Utah QSO Party
	CONTEST_VA_QSO_PARTY           Contest = "va-qso-party"           // va-qso-party         = Virginia QSO Party
	CONTEST_VENEZ_IND_DAY          Contest = "venez-ind-day"          // venez-ind-day        = RCV Venezuelan Independence Day Contest
	CONTEST_VIRGINIA_QSO_PARTY     Contest = "virginia qso party"     // Deprecated: virginia qso party   = Virginia QSO Party
	CONTEST_VOLTA_RTTY             Contest = "volta-rtty"             // volta-rtty           = Alessandro Volta RTTY DX Contest
	CONTEST_VT_QSO_PARTY           Contest = "vt-qso-party"           // vt-qso-party         = Vermont QSO Party
	CONTEST_WA_QSO_PARTY           Contest = "wa-qso-party"           // wa-qso-party         = Washington QSO Party
	CONTEST_WFD                    Contest = "wfd"                    // wfd                  = Winter Field Day (2017 and later)
	CONTEST_WI_QSO_PARTY           Contest = "wi-qso-party"           // wi-qso-party         = Wisconsin QSO Party
	CONTEST_WIA_HARRY_ANGEL        Contest = "wia-harry angel"        // wia-harry angel      = WIA Harry Angel Memorial 80m Sprint
	CONTEST_WIA_JMMFD              Contest = "wia-jmmfd"              // wia-jmmfd            = WIA John Moyle Memorial Field Day
	CONTEST_WIA_OCDX               Contest = "wia-ocdx"               // wia-ocdx             = WIA Oceania DX (OCDX) Contest
	CONTEST_WIA_REMEMBRANCE        Contest = "wia-remembrance"        // wia-remembrance      = WIA Remembrance Day
	CONTEST_WIA_ROSS_HULL          Contest = "wia-ross hull"          // wia-ross hull        = WIA Ross Hull Memorial VHF/UHF Contest
	CONTEST_WIA_TRANS_TASMAN       Contest = "wia-trans tasman"       // wia-trans tasman     = WIA Trans Tasman Low Bands Challenge
	CONTEST_WIA_VHF_UHF_FD         Contest = "wia-vhf/uhf fd"         // wia-vhf/uhf fd       = WIA VHF UHF Field Days
	CONTEST_WIA_VK_SHIRES          Contest = "wia-vk shires"          // wia-vk shires        = WIA VK Shires
	CONTEST_WINTER_SPRINT          Contest = "winter sprint"          // winter sprint        = FISTS Winter Sprint
	CONTEST_WV_QSO_PARTY           Contest = "wv-qso-party"           // wv-qso-party         = West Virginia QSO Party
	CONTEST_WW_DIGI                Contest = "ww-digi"                // ww-digi              = World Wide Digi DX Contest
	CONTEST_WY_QSO_PARTY           Contest = "wy-qso-party"           // wy-qso-party         = Wyoming QSO Party
	CONTEST_XE_INTL_RTTY           Contest = "xe-intl-rtty"           // xe-intl-rtty         = Mexico International Contest (RTTY)
	CONTEST_YOHFDX                 Contest = "yohfdx"                 // yohfdx               = YODX HF contest
	CONTEST_YUDXC                  Contest = "yudxc"                  // yudxc                = YU DX Contest
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known Contest specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: "070-160m-sprint", Description: "PODXS Great Pumpkin Sprint"},
	{IsImportOnly: false, Key: "070-3-day", Description: "PODXS Three Day Weekend"},
	{IsImportOnly: false, Key: "070-31-flavors", Description: "PODXS 31 Flavors"},
	{IsImportOnly: false, Key: "070-40m-sprint", Description: "PODXS 40m Firecracker Sprint"},
	{IsImportOnly: false, Key: "070-80m-sprint", Description: "PODXS 80m Jay Hudak Memorial Sprint"},
	{IsImportOnly: false, Key: "070-pskfest", Description: "PODXS PSKFest"},
	{IsImportOnly: false, Key: "070-st-pats-day", Description: "PODXS St. Patricks Day"},
	{IsImportOnly: false, Key: "070-valentine-sprint", Description: "PODXS Valentine Sprint"},
	{IsImportOnly: false, Key: "10-rtty", Description: "Ten-Meter RTTY Contest (2011 onwards)"},
	{IsImportOnly: false, Key: "1010-open-season", Description: "Open Season Ten Meter QSO Party"},
	{IsImportOnly: false, Key: "7qp", Description: "7th-Area QSO Party"},
	{IsImportOnly: false, Key: "al-qso-party", Description: "Alabama QSO Party"},
	{IsImportOnly: false, Key: "all-asian-dx-cw", Description: "JARL All Asian DX Contest (CW)"},
	{IsImportOnly: false, Key: "all-asian-dx-phone", Description: "JARL All Asian DX Contest (PHONE)"},
	{IsImportOnly: false, Key: "anarts-rtty", Description: "ANARTS WW RTTY"},
	{IsImportOnly: false, Key: "anatolian-rtty", Description: "Anatolian WW RTTY"},
	{IsImportOnly: false, Key: "ap-sprint", Description: "Asia - Pacific Sprint"},
	{IsImportOnly: false, Key: "ar-qso-party", Description: "Arkansas QSO Party"},
	{IsImportOnly: false, Key: "ari-dx", Description: "ARI DX Contest"},
	{IsImportOnly: false, Key: "ari-eme", Description: "ARI Italian EME Trophy"},
	{IsImportOnly: false, Key: "ari-iac-13cm", Description: "ARI Italian Activity Contest (13cm+)"},
	{IsImportOnly: false, Key: "ari-iac-23cm", Description: "ARI Italian Activity Contest (23cm)"},
	{IsImportOnly: false, Key: "ari-iac-6m", Description: "ARI Italian Activity Contest (6m)"},
	{IsImportOnly: false, Key: "ari-iac-uhf", Description: "ARI Italian Activity Contest (UHF)"},
	{IsImportOnly: false, Key: "ari-iac-vhf", Description: "ARI Italian Activity Contest (VHF)"},
	{IsImportOnly: false, Key: "arrl-10", Description: "ARRL 10 Meter Contest"},
	{IsImportOnly: false, Key: "arrl-10-ghz", Description: "ARRL 10 GHz and Up Contest"},
	{IsImportOnly: false, Key: "arrl-160", Description: "ARRL 160 Meter Contest"},
	{IsImportOnly: false, Key: "arrl-222", Description: "ARRL 222 MHz and Up Distance Contest"},
	{IsImportOnly: false, Key: "arrl-digi", Description: "ARRL International Digital Contest"},
	{IsImportOnly: false, Key: "arrl-dx-cw", Description: "ARRL International DX Contest (CW)"},
	{IsImportOnly: false, Key: "arrl-dx-ssb", Description: "ARRL International DX Contest (Phone)"},
	{IsImportOnly: false, Key: "arrl-eme", Description: "ARRL EME contest"},
	{IsImportOnly: false, Key: "arrl-field-day", Description: "ARRL Field Day"},
	{IsImportOnly: false, Key: "arrl-rr-cw", Description: "ARRL Rookie Roundup (CW)"},
	{IsImportOnly: false, Key: "arrl-rr-rtty", Description: "ARRL Rookie Roundup (RTTY)"},
	{IsImportOnly: false, Key: "arrl-rr-ssb", Description: "ARRL Rookie Roundup (Phone)"},
	{IsImportOnly: false, Key: "arrl-rtty", Description: "ARRL RTTY Round-Up"},
	{IsImportOnly: false, Key: "arrl-scr", Description: "ARRL School Club Roundup"},
	{IsImportOnly: false, Key: "arrl-ss-cw", Description: "ARRL November Sweepstakes (CW)"},
	{IsImportOnly: false, Key: "arrl-ss-ssb", Description: "ARRL November Sweepstakes (Phone)"},
	{IsImportOnly: false, Key: "arrl-uhf-aug", Description: "ARRL August UHF Contest"},
	{IsImportOnly: false, Key: "arrl-vhf-jan", Description: "ARRL January VHF Sweepstakes"},
	{IsImportOnly: false, Key: "arrl-vhf-jun", Description: "ARRL June VHF QSO Party"},
	{IsImportOnly: false, Key: "arrl-vhf-sep", Description: "ARRL September VHF QSO Party"},
	{IsImportOnly: false, Key: "az-qso-party", Description: "Arizona QSO Party"},
	{IsImportOnly: false, Key: "banggai-dx", Description: "ORARI Banggai DX Contest"},
	{IsImportOnly: false, Key: "bartg-rtty", Description: "BARTG Spring RTTY Contest"},
	{IsImportOnly: false, Key: "bartg-sprint", Description: "BARTG Sprint Contest"},
	{IsImportOnly: false, Key: "bc-qso-party", Description: "British Columbia QSO Party"},
	{IsImportOnly: false, Key: "bekasi-merdeka-contest", Description: "ORARI Bekasi Merdeka Contest"},
	{IsImportOnly: false, Key: "ca-qso-party", Description: "California QSO Party"},
	{IsImportOnly: false, Key: "cis-dx", Description: "CIS DX Contest"},
	{IsImportOnly: false, Key: "co-qso-party", Description: "Colorado QSO Party"},
	{IsImportOnly: false, Key: "cq-160-cw", Description: "CQ WW 160 Meter DX Contest (CW)"},
	{IsImportOnly: false, Key: "cq-160-ssb", Description: "CQ WW 160 Meter DX Contest (SSB)"},
	{IsImportOnly: false, Key: "cq-m", Description: "CQ-M International DX Contest"},
	{IsImportOnly: false, Key: "cq-vhf", Description: "CQ World-Wide VHF Contest"},
	{IsImportOnly: false, Key: "cq-wpx-cw", Description: "CQ WW WPX Contest (CW)"},
	{IsImportOnly: false, Key: "cq-wpx-rtty", Description: "CQ/RJ WW RTTY WPX Contest"},
	{IsImportOnly: false, Key: "cq-wpx-ssb", Description: "CQ WW WPX Contest (SSB)"},
	{IsImportOnly: false, Key: "cq-ww-cw", Description: "CQ WW DX Contest (CW)"},
	{IsImportOnly: false, Key: "cq-ww-rtty", Description: "CQ/RJ WW RTTY DX Contest"},
	{IsImportOnly: false, Key: "cq-ww-ssb", Description: "CQ WW DX Contest (SSB)"},
	{IsImportOnly: false, Key: "ct-qso-party", Description: "Connecticut QSO Party"},
	{IsImportOnly: false, Key: "cva-dx-cw", Description: "Concurso Verde e Amarelo DX CW Contest"},
	{IsImportOnly: false, Key: "cva-dx-ssb", Description: "Concurso Verde e Amarelo DX CW Contest"},
	{IsImportOnly: false, Key: "cwops-cw-open", Description: "CWops CW Open Competition"},
	{IsImportOnly: false, Key: "cwops-cwt", Description: "CWops Mini-CWT Test"},
	{IsImportOnly: false, Key: "darc-10", Description: "DARC 10m Contest"},
	{IsImportOnly: false, Key: "darc-cwa", Description: "DARC CW Trainee Contest"},
	{IsImportOnly: false, Key: "darc-ft4", Description: "DARC FT4 Contest"},
	{IsImportOnly: false, Key: "darc-hell", Description: "DARC Hell Contest"},
	{IsImportOnly: false, Key: "darc-microwave", Description: "DARC Microwave Contest"},
	{IsImportOnly: false, Key: "darc-trainee", Description: "DARC Trainee Contest"},
	{IsImportOnly: false, Key: "darc-ukw-field-day", Description: "DARC UKW Summer Contest"},
	{IsImportOnly: false, Key: "darc-ukw-spring", Description: "DARC UKW Spring Contest"},
	{IsImportOnly: false, Key: "darc-vhf-uhf-microwave", Description: "DARC VHF-, UHF-, Microwave Contest (May)"},
	{IsImportOnly: false, Key: "darc-waedc-cw", Description: "WAE DX Contest (CW)"},
	{IsImportOnly: false, Key: "darc-waedc-rtty", Description: "WAE DX Contest (RTTY)"},
	{IsImportOnly: false, Key: "darc-waedc-ssb", Description: "WAE DX Contest (SSB)"},
	{IsImportOnly: false, Key: "darc-wag", Description: "DARC Worked All Germany"},
	{IsImportOnly: false, Key: "de-qso-party", Description: "Delaware QSO Party"},
	{IsImportOnly: false, Key: "dl-dx-rtty", Description: "DL-DX RTTY Contest"},
	{IsImportOnly: false, Key: "dmc-rtty", Description: "DMC RTTY Contest"},
	{IsImportOnly: false, Key: "ea-cncw", Description: "Concurso Nacional de Telegrafía"},
	{IsImportOnly: false, Key: "ea-dme", Description: "Municipios Españoles"},
	{IsImportOnly: false, Key: "ea-majestad-cw", Description: "His Majesty The King of Spain CW Contest (2022 and later)"},
	{IsImportOnly: false, Key: "ea-majestad-ssb", Description: "His Majesty The King of Spain SSB Contest (2022 and later)"},
	{IsImportOnly: false, Key: "ea-psk63", Description: "EA PSK63"},
	{IsImportOnly: true, Key: "ea-rtty", Description: "Unión de Radioaficionados Españoles RTTY Contest"},
	{IsImportOnly: false, Key: "ea-smre-cw", Description: "Su Majestad El Rey de España - CW (2021 and earlier)"},
	{IsImportOnly: false, Key: "ea-smre-ssb", Description: "Su Majestad El Rey de España - SSB (2021 and earlier)"},
	{IsImportOnly: false, Key: "ea-vhf-atlantic", Description: "Atlántico V-UHF"},
	{IsImportOnly: false, Key: "ea-vhf-com", Description: "Combinado de V-UHF"},
	{IsImportOnly: false, Key: "ea-vhf-costa-sol", Description: "Costa del Sol V-UHF"},
	{IsImportOnly: false, Key: "ea-vhf-ea", Description: "Nacional VHF"},
	{IsImportOnly: false, Key: "ea-vhf-ea1rcs", Description: "Segovia EA1RCS V-UHF"},
	{IsImportOnly: false, Key: "ea-vhf-qsl", Description: "QSL V-UHF & 50MHz"},
	{IsImportOnly: false, Key: "ea-vhf-sadurni", Description: "Sant Sadurni V-UHF"},
	{IsImportOnly: false, Key: "ea-ww-rtty", Description: "Unión de Radioaficionados Españoles RTTY Contest"},
	{IsImportOnly: false, Key: "easter", Description: "DARC Easter Contest"},
	{IsImportOnly: false, Key: "epc-psk63", Description: "PSK63 QSO Party"},
	{IsImportOnly: false, Key: "eu sprint", Description: "EU Sprint"},
	{IsImportOnly: false, Key: "eu-hf", Description: "EU HF Championship"},
	{IsImportOnly: false, Key: "eu-psk-dx", Description: "EU PSK DX Contest"},
	{IsImportOnly: false, Key: "eucw160m", Description: "European CW Association 160m CW Party"},
	{IsImportOnly: false, Key: "fall sprint", Description: "FISTS Fall Sprint"},
	{IsImportOnly: false, Key: "fl-qso-party", Description: "Florida QSO Party"},
	{IsImportOnly: false, Key: "ga-qso-party", Description: "Georgia QSO Party"},
	{IsImportOnly: false, Key: "ha-dx", Description: "Hungarian DX Contest"},
	{IsImportOnly: false, Key: "helvetia", Description: "Helvetia Contest"},
	{IsImportOnly: false, Key: "hi-qso-party", Description: "Hawaiian QSO Party"},
	{IsImportOnly: false, Key: "holyland", Description: "IARC Holyland Contest"},
	{IsImportOnly: false, Key: "ia-qso-party", Description: "Iowa QSO Party"},
	{IsImportOnly: false, Key: "iaru-field-day", Description: "DARC IARU Region 1 Field Day"},
	{IsImportOnly: false, Key: "iaru-hf", Description: "IARU HF World Championship"},
	{IsImportOnly: false, Key: "icwc-mst", Description: "ICWC Medium Speed Test"},
	{IsImportOnly: false, Key: "id-qso-party", Description: "Idaho QSO Party"},
	{IsImportOnly: false, Key: "il qso party", Description: "Illinois QSO Party"},
	{IsImportOnly: false, Key: "in-qso-party", Description: "Indiana QSO Party"},
	{IsImportOnly: false, Key: "jarts-ww-rtty", Description: "JARTS WW RTTY"},
	{IsImportOnly: false, Key: "jidx-cw", Description: "Japan International DX Contest (CW)"},
	{IsImportOnly: false, Key: "jidx-ssb", Description: "Japan International DX Contest (SSB)"},
	{IsImportOnly: false, Key: "jt-dx-rtty", Description: "Mongolian RTTY DX Contest"},
	{IsImportOnly: false, Key: "k1usn-sso", Description: "K1USN Slow Speed Open"},
	{IsImportOnly: false, Key: "k1usn-sst", Description: "K1USN Slow Speed Test"},
	{IsImportOnly: false, Key: "ks-qso-party", Description: "Kansas QSO Party"},
	{IsImportOnly: false, Key: "ky-qso-party", Description: "Kentucky QSO Party"},
	{IsImportOnly: false, Key: "la-qso-party", Description: "Louisiana QSO Party"},
	{IsImportOnly: false, Key: "ldc-rtty", Description: "DRCG Long Distance Contest (RTTY)"},
	{IsImportOnly: false, Key: "lz dx", Description: "LZ DX Contest"},
	{IsImportOnly: false, Key: "mar-qso-party", Description: "Maritimes QSO Party"},
	{IsImportOnly: false, Key: "md-qso-party", Description: "Maryland QSO Party"},
	{IsImportOnly: false, Key: "me-qso-party", Description: "Maine QSO Party"},
	{IsImportOnly: false, Key: "mi-qso-party", Description: "Michigan QSO Party"},
	{IsImportOnly: false, Key: "midatlantic-qso-party", Description: "Mid-Atlantic QSO Party"},
	{IsImportOnly: false, Key: "mn-qso-party", Description: "Minnesota QSO Party"},
	{IsImportOnly: false, Key: "mo-qso-party", Description: "Missouri QSO Party"},
	{IsImportOnly: false, Key: "ms-qso-party", Description: "Mississippi QSO Party"},
	{IsImportOnly: false, Key: "mt-qso-party", Description: "Montana QSO Party"},
	{IsImportOnly: false, Key: "na-sprint-cw", Description: "North America Sprint (CW)"},
	{IsImportOnly: false, Key: "na-sprint-rtty", Description: "North America Sprint (RTTY)"},
	{IsImportOnly: false, Key: "na-sprint-ssb", Description: "North America Sprint (Phone)"},
	{IsImportOnly: false, Key: "naqp-cw", Description: "North America QSO Party (CW)"},
	{IsImportOnly: false, Key: "naqp-rtty", Description: "North America QSO Party (RTTY)"},
	{IsImportOnly: false, Key: "naqp-ssb", Description: "North America QSO Party (Phone)"},
	{IsImportOnly: false, Key: "naval", Description: "International Naval Contest (INC)"},
	{IsImportOnly: false, Key: "nc-qso-party", Description: "North Carolina QSO Party"},
	{IsImportOnly: false, Key: "nd-qso-party", Description: "North Dakota QSO Party"},
	{IsImportOnly: false, Key: "ne-qso-party", Description: "Nebraska QSO Party"},
	{IsImportOnly: false, Key: "neqp", Description: "New England QSO Party"},
	{IsImportOnly: false, Key: "nh-qso-party", Description: "New Hampshire QSO Party"},
	{IsImportOnly: false, Key: "nj-qso-party", Description: "New Jersey QSO Party"},
	{IsImportOnly: false, Key: "nm-qso-party", Description: "New Mexico QSO Party"},
	{IsImportOnly: false, Key: "nrau-baltic-cw", Description: "NRAU-Baltic Contest (CW)"},
	{IsImportOnly: false, Key: "nrau-baltic-ssb", Description: "NRAU-Baltic Contest (SSB)"},
	{IsImportOnly: false, Key: "nv-qso-party", Description: "Nevada QSO Party"},
	{IsImportOnly: false, Key: "ny-qso-party", Description: "New York QSO Party"},
	{IsImportOnly: false, Key: "oceania-dx-cw", Description: "Oceania DX Contest (CW)"},
	{IsImportOnly: false, Key: "oceania-dx-ssb", Description: "Oceania DX Contest (SSB)"},
	{IsImportOnly: false, Key: "oh-qso-party", Description: "Ohio QSO Party"},
	{IsImportOnly: false, Key: "ok-dx-rtty", Description: "Czech Radio Club OK DX Contest"},
	{IsImportOnly: false, Key: "ok-om-dx", Description: "Czech Radio Club OK-OM DX Contest"},
	{IsImportOnly: false, Key: "ok-qso-party", Description: "Oklahoma QSO Party"},
	{IsImportOnly: false, Key: "omiss-qso-party", Description: "Old Man International Sideband Society QSO Party"},
	{IsImportOnly: false, Key: "on-qso-party", Description: "Ontario QSO Party"},
	{IsImportOnly: false, Key: "or-qso-party", Description: "Oregon QSO Party"},
	{IsImportOnly: false, Key: "orari-dx", Description: "ORARI DX Contest"},
	{IsImportOnly: false, Key: "pa-qso-party", Description: "Pennsylvania QSO Party"},
	{IsImportOnly: false, Key: "pacc", Description: "Dutch PACC Contest"},
	{IsImportOnly: false, Key: "pcc", Description: "PCCPro CW Contest"},
	{IsImportOnly: false, Key: "psk-deathmatch", Description: "MDXA PSK DeathMatch (2005-2010)"},
	{IsImportOnly: false, Key: "qc-qso-party", Description: "Quebec QSO Party"},
	{IsImportOnly: true, Key: "rac", Description: "Canadian Amateur Radio Society Contest"},
	{IsImportOnly: false, Key: "rac-canada-day", Description: "RAC Canada Day Contest"},
	{IsImportOnly: false, Key: "rac-canada-winter", Description: "RAC Canada Winter Contest"},
	{IsImportOnly: false, Key: "rdac", Description: "Russian District Award Contest"},
	{IsImportOnly: false, Key: "rdxc", Description: "Russian DX Contest"},
	{IsImportOnly: false, Key: "ref-160m", Description: "Reseau des Emetteurs Francais 160m Contest"},
	{IsImportOnly: false, Key: "ref-cw", Description: "Reseau des Emetteurs Francais Contest (CW)"},
	{IsImportOnly: false, Key: "ref-ssb", Description: "Reseau des Emetteurs Francais Contest (SSB)"},
	{IsImportOnly: false, Key: "rep-portugal-day-hf", Description: "Rede dos Emissores Portugueses Portugal Day HF Contest"},
	{IsImportOnly: false, Key: "ri-qso-party", Description: "Rhode Island QSO Party"},
	{IsImportOnly: false, Key: "rsgb-160", Description: "1.8MHz Contest"},
	{IsImportOnly: false, Key: "rsgb-21/28-cw", Description: "21/28 MHz Contest (CW)"},
	{IsImportOnly: false, Key: "rsgb-21/28-ssb", Description: "21/28 MHz Contest (SSB)"},
	{IsImportOnly: false, Key: "rsgb-80m-cc", Description: "80m Club Championships"},
	{IsImportOnly: false, Key: "rsgb-afs-cw", Description: "Affiliated Societies Team Contest (CW)"},
	{IsImportOnly: false, Key: "rsgb-afs-ssb", Description: "Affiliated Societies Team Contest (SSB)"},
	{IsImportOnly: false, Key: "rsgb-club-calls", Description: "Club Calls"},
	{IsImportOnly: false, Key: "rsgb-commonwealth", Description: "Commonwealth Contest"},
	{IsImportOnly: false, Key: "rsgb-iota", Description: "IOTA Contest"},
	{IsImportOnly: false, Key: "rsgb-low-power", Description: "Low Power Field Day"},
	{IsImportOnly: false, Key: "rsgb-nfd", Description: "National Field Day"},
	{IsImportOnly: false, Key: "rsgb-ropoco", Description: "RoPoCo"},
	{IsImportOnly: false, Key: "rsgb-ssb-fd", Description: "SSB Field Day"},
	{IsImportOnly: false, Key: "russian-rtty", Description: "Russian Radio RTTY Worldwide Contest"},
	{IsImportOnly: false, Key: "sac-cw", Description: "Scandinavian Activity Contest (CW)"},
	{IsImportOnly: false, Key: "sac-ssb", Description: "Scandinavian Activity Contest (SSB)"},
	{IsImportOnly: false, Key: "sartg-rtty", Description: "SARTG WW RTTY"},
	{IsImportOnly: false, Key: "sc-qso-party", Description: "South Carolina QSO Party"},
	{IsImportOnly: false, Key: "scc-rtty", Description: "SCC RTTY Championship"},
	{IsImportOnly: false, Key: "sd-qso-party", Description: "South Dakota QSO Party"},
	{IsImportOnly: false, Key: "shortry", Description: "DARC RTTY Short Contest"},
	{IsImportOnly: false, Key: "smp-aug", Description: "SSA Portabeltest"},
	{IsImportOnly: false, Key: "smp-may", Description: "SSA Portabeltest"},
	{IsImportOnly: false, Key: "sp-dx-rtty", Description: "PRC SPDX Contest (RTTY)"},
	{IsImportOnly: false, Key: "spar-winter-fd", Description: "SPAR Winter Field Day(2016 and earlier)"},
	{IsImportOnly: false, Key: "spdxcontest", Description: "SP DX Contest"},
	{IsImportOnly: false, Key: "spring sprint", Description: "FISTS Spring Sprint"},
	{IsImportOnly: false, Key: "sr-marathon", Description: "Scottish-Russian Marathon"},
	{IsImportOnly: false, Key: "stew-perry", Description: "Stew Perry Topband Distance Challenge"},
	{IsImportOnly: false, Key: "summer sprint", Description: "FISTS Summer Sprint"},
	{IsImportOnly: false, Key: "tara-grid-dip", Description: "TARA Grid Dip PSK-RTTY Shindig"},
	{IsImportOnly: false, Key: "tara-rtty", Description: "TARA RTTY Mêlée"},
	{IsImportOnly: false, Key: "tara-rumble", Description: "TARA Rumble PSK Contest"},
	{IsImportOnly: false, Key: "tara-skirmish", Description: "TARA Skirmish Digital Prefix Contest"},
	{IsImportOnly: false, Key: "ten-rtty", Description: "Ten-Meter RTTY Contest (before 2011)"},
	{IsImportOnly: false, Key: "tmc-rtty", Description: "The Makrothen Contest"},
	{IsImportOnly: false, Key: "tn-qso-party", Description: "Tennessee QSO Party"},
	{IsImportOnly: false, Key: "tx-qso-party", Description: "Texas QSO Party"},
	{IsImportOnly: false, Key: "uba-dx-cw", Description: "UBA Contest (CW)"},
	{IsImportOnly: false, Key: "uba-dx-ssb", Description: "UBA Contest (SSB)"},
	{IsImportOnly: false, Key: "uk-dx-bpsk63", Description: "European PSK Club BPSK63 Contest"},
	{IsImportOnly: false, Key: "uk-dx-rtty", Description: "UK DX RTTY Contest"},
	{IsImportOnly: false, Key: "ukr-champ-rtty", Description: "Open Ukraine RTTY Championship"},
	{IsImportOnly: false, Key: "ukrainian dx", Description: "Ukrainian DX"},
	{IsImportOnly: false, Key: "uksmg-6m-marathon", Description: "UKSMG 6m Marathon"},
	{IsImportOnly: false, Key: "uksmg-summer-es", Description: "UKSMG Summer Es Contest"},
	{IsImportOnly: true, Key: "ure-dx", Description: "Ukrainian DX Contest"},
	{IsImportOnly: false, Key: "us-counties-qso", Description: "Mobile Amateur Awards Club"},
	{IsImportOnly: false, Key: "ut-qso-party", Description: "Utah QSO Party"},
	{IsImportOnly: false, Key: "va-qso-party", Description: "Virginia QSO Party"},
	{IsImportOnly: false, Key: "venez-ind-day", Description: "RCV Venezuelan Independence Day Contest"},
	{IsImportOnly: true, Key: "virginia qso party", Description: "Virginia QSO Party"},
	{IsImportOnly: false, Key: "volta-rtty", Description: "Alessandro Volta RTTY DX Contest"},
	{IsImportOnly: false, Key: "vt-qso-party", Description: "Vermont QSO Party"},
	{IsImportOnly: false, Key: "wa-qso-party", Description: "Washington QSO Party"},
	{IsImportOnly: false, Key: "wfd", Description: "Winter Field Day (2017 and later)"},
	{IsImportOnly: false, Key: "wi-qso-party", Description: "Wisconsin QSO Party"},
	{IsImportOnly: false, Key: "wia-harry angel", Description: "WIA Harry Angel Memorial 80m Sprint"},
	{IsImportOnly: false, Key: "wia-jmmfd", Description: "WIA John Moyle Memorial Field Day"},
	{IsImportOnly: false, Key: "wia-ocdx", Description: "WIA Oceania DX (OCDX) Contest"},
	{IsImportOnly: false, Key: "wia-remembrance", Description: "WIA Remembrance Day"},
	{IsImportOnly: false, Key: "wia-ross hull", Description: "WIA Ross Hull Memorial VHF/UHF Contest"},
	{IsImportOnly: false, Key: "wia-trans tasman", Description: "WIA Trans Tasman Low Bands Challenge"},
	{IsImportOnly: false, Key: "wia-vhf/uhf fd", Description: "WIA VHF UHF Field Days"},
	{IsImportOnly: false, Key: "wia-vk shires", Description: "WIA VK Shires"},
	{IsImportOnly: false, Key: "winter sprint", Description: "FISTS Winter Sprint"},
	{IsImportOnly: false, Key: "wv-qso-party", Description: "West Virginia QSO Party"},
	{IsImportOnly: false, Key: "ww-digi", Description: "World Wide Digi DX Contest"},
	{IsImportOnly: false, Key: "wy-qso-party", Description: "Wyoming QSO Party"},
	{IsImportOnly: false, Key: "xe-intl-rtty", Description: "Mexico International Contest (RTTY)"},
	{IsImportOnly: false, Key: "yohfdx", Description: "YODX HF contest"},
	{IsImportOnly: false, Key: "yudxc", Description: "YU DX Contest"},
}

// lookupMap contains all known Contest specifications.
var lookupMap = map[Contest]*Spec{
	CONTEST_070_160M_SPRINT:        &lookupList[0],
	CONTEST_070_3_DAY:              &lookupList[1],
	CONTEST_070_31_FLAVORS:         &lookupList[2],
	CONTEST_070_40M_SPRINT:         &lookupList[3],
	CONTEST_070_80M_SPRINT:         &lookupList[4],
	CONTEST_070_PSKFEST:            &lookupList[5],
	CONTEST_070_ST_PATS_DAY:        &lookupList[6],
	CONTEST_070_VALENTINE_SPRINT:   &lookupList[7],
	CONTEST_10_RTTY:                &lookupList[8],
	CONTEST_1010_OPEN_SEASON:       &lookupList[9],
	CONTEST_7QP:                    &lookupList[10],
	CONTEST_AL_QSO_PARTY:           &lookupList[11],
	CONTEST_ALL_ASIAN_DX_CW:        &lookupList[12],
	CONTEST_ALL_ASIAN_DX_PHONE:     &lookupList[13],
	CONTEST_ANARTS_RTTY:            &lookupList[14],
	CONTEST_ANATOLIAN_RTTY:         &lookupList[15],
	CONTEST_AP_SPRINT:              &lookupList[16],
	CONTEST_AR_QSO_PARTY:           &lookupList[17],
	CONTEST_ARI_DX:                 &lookupList[18],
	CONTEST_ARI_EME:                &lookupList[19],
	CONTEST_ARI_IAC_13CM:           &lookupList[20],
	CONTEST_ARI_IAC_23CM:           &lookupList[21],
	CONTEST_ARI_IAC_6M:             &lookupList[22],
	CONTEST_ARI_IAC_UHF:            &lookupList[23],
	CONTEST_ARI_IAC_VHF:            &lookupList[24],
	CONTEST_ARRL_10:                &lookupList[25],
	CONTEST_ARRL_10_GHZ:            &lookupList[26],
	CONTEST_ARRL_160:               &lookupList[27],
	CONTEST_ARRL_222:               &lookupList[28],
	CONTEST_ARRL_DIGI:              &lookupList[29],
	CONTEST_ARRL_DX_CW:             &lookupList[30],
	CONTEST_ARRL_DX_SSB:            &lookupList[31],
	CONTEST_ARRL_EME:               &lookupList[32],
	CONTEST_ARRL_FIELD_DAY:         &lookupList[33],
	CONTEST_ARRL_RR_CW:             &lookupList[34],
	CONTEST_ARRL_RR_RTTY:           &lookupList[35],
	CONTEST_ARRL_RR_SSB:            &lookupList[36],
	CONTEST_ARRL_RTTY:              &lookupList[37],
	CONTEST_ARRL_SCR:               &lookupList[38],
	CONTEST_ARRL_SS_CW:             &lookupList[39],
	CONTEST_ARRL_SS_SSB:            &lookupList[40],
	CONTEST_ARRL_UHF_AUG:           &lookupList[41],
	CONTEST_ARRL_VHF_JAN:           &lookupList[42],
	CONTEST_ARRL_VHF_JUN:           &lookupList[43],
	CONTEST_ARRL_VHF_SEP:           &lookupList[44],
	CONTEST_AZ_QSO_PARTY:           &lookupList[45],
	CONTEST_BANGGAI_DX:             &lookupList[46],
	CONTEST_BARTG_RTTY:             &lookupList[47],
	CONTEST_BARTG_SPRINT:           &lookupList[48],
	CONTEST_BC_QSO_PARTY:           &lookupList[49],
	CONTEST_BEKASI_MERDEKA_CONTEST: &lookupList[50],
	CONTEST_CA_QSO_PARTY:           &lookupList[51],
	CONTEST_CIS_DX:                 &lookupList[52],
	CONTEST_CO_QSO_PARTY:           &lookupList[53],
	CONTEST_CQ_160_CW:              &lookupList[54],
	CONTEST_CQ_160_SSB:             &lookupList[55],
	CONTEST_CQ_M:                   &lookupList[56],
	CONTEST_CQ_VHF:                 &lookupList[57],
	CONTEST_CQ_WPX_CW:              &lookupList[58],
	CONTEST_CQ_WPX_RTTY:            &lookupList[59],
	CONTEST_CQ_WPX_SSB:             &lookupList[60],
	CONTEST_CQ_WW_CW:               &lookupList[61],
	CONTEST_CQ_WW_RTTY:             &lookupList[62],
	CONTEST_CQ_WW_SSB:              &lookupList[63],
	CONTEST_CT_QSO_PARTY:           &lookupList[64],
	CONTEST_CVA_DX_CW:              &lookupList[65],
	CONTEST_CVA_DX_SSB:             &lookupList[66],
	CONTEST_CWOPS_CW_OPEN:          &lookupList[67],
	CONTEST_CWOPS_CWT:              &lookupList[68],
	CONTEST_DARC_10:                &lookupList[69],
	CONTEST_DARC_CWA:               &lookupList[70],
	CONTEST_DARC_FT4:               &lookupList[71],
	CONTEST_DARC_HELL:              &lookupList[72],
	CONTEST_DARC_MICROWAVE:         &lookupList[73],
	CONTEST_DARC_TRAINEE:           &lookupList[74],
	CONTEST_DARC_UKW_FIELD_DAY:     &lookupList[75],
	CONTEST_DARC_UKW_SPRING:        &lookupList[76],
	CONTEST_DARC_VHF_UHF_MICROWAVE: &lookupList[77],
	CONTEST_DARC_WAEDC_CW:          &lookupList[78],
	CONTEST_DARC_WAEDC_RTTY:        &lookupList[79],
	CONTEST_DARC_WAEDC_SSB:         &lookupList[80],
	CONTEST_DARC_WAG:               &lookupList[81],
	CONTEST_DE_QSO_PARTY:           &lookupList[82],
	CONTEST_DL_DX_RTTY:             &lookupList[83],
	CONTEST_DMC_RTTY:               &lookupList[84],
	CONTEST_EA_CNCW:                &lookupList[85],
	CONTEST_EA_DME:                 &lookupList[86],
	CONTEST_EA_MAJESTAD_CW:         &lookupList[87],
	CONTEST_EA_MAJESTAD_SSB:        &lookupList[88],
	CONTEST_EA_PSK63:               &lookupList[89],
	CONTEST_EA_RTTY:                &lookupList[90],
	CONTEST_EA_SMRE_CW:             &lookupList[91],
	CONTEST_EA_SMRE_SSB:            &lookupList[92],
	CONTEST_EA_VHF_ATLANTIC:        &lookupList[93],
	CONTEST_EA_VHF_COM:             &lookupList[94],
	CONTEST_EA_VHF_COSTA_SOL:       &lookupList[95],
	CONTEST_EA_VHF_EA:              &lookupList[96],
	CONTEST_EA_VHF_EA1RCS:          &lookupList[97],
	CONTEST_EA_VHF_QSL:             &lookupList[98],
	CONTEST_EA_VHF_SADURNI:         &lookupList[99],
	CONTEST_EA_WW_RTTY:             &lookupList[100],
	CONTEST_EASTER:                 &lookupList[101],
	CONTEST_EPC_PSK63:              &lookupList[102],
	CONTEST_EU_SPRINT:              &lookupList[103],
	CONTEST_EU_HF:                  &lookupList[104],
	CONTEST_EU_PSK_DX:              &lookupList[105],
	CONTEST_EUCW160M:               &lookupList[106],
	CONTEST_FALL_SPRINT:            &lookupList[107],
	CONTEST_FL_QSO_PARTY:           &lookupList[108],
	CONTEST_GA_QSO_PARTY:           &lookupList[109],
	CONTEST_HA_DX:                  &lookupList[110],
	CONTEST_HELVETIA:               &lookupList[111],
	CONTEST_HI_QSO_PARTY:           &lookupList[112],
	CONTEST_HOLYLAND:               &lookupList[113],
	CONTEST_IA_QSO_PARTY:           &lookupList[114],
	CONTEST_IARU_FIELD_DAY:         &lookupList[115],
	CONTEST_IARU_HF:                &lookupList[116],
	CONTEST_ICWC_MST:               &lookupList[117],
	CONTEST_ID_QSO_PARTY:           &lookupList[118],
	CONTEST_IL_QSO_PARTY:           &lookupList[119],
	CONTEST_IN_QSO_PARTY:           &lookupList[120],
	CONTEST_JARTS_WW_RTTY:          &lookupList[121],
	CONTEST_JIDX_CW:                &lookupList[122],
	CONTEST_JIDX_SSB:               &lookupList[123],
	CONTEST_JT_DX_RTTY:             &lookupList[124],
	CONTEST_K1USN_SSO:              &lookupList[125],
	CONTEST_K1USN_SST:              &lookupList[126],
	CONTEST_KS_QSO_PARTY:           &lookupList[127],
	CONTEST_KY_QSO_PARTY:           &lookupList[128],
	CONTEST_LA_QSO_PARTY:           &lookupList[129],
	CONTEST_LDC_RTTY:               &lookupList[130],
	CONTEST_LZ_DX:                  &lookupList[131],
	CONTEST_MAR_QSO_PARTY:          &lookupList[132],
	CONTEST_MD_QSO_PARTY:           &lookupList[133],
	CONTEST_ME_QSO_PARTY:           &lookupList[134],
	CONTEST_MI_QSO_PARTY:           &lookupList[135],
	CONTEST_MIDATLANTIC_QSO_PARTY:  &lookupList[136],
	CONTEST_MN_QSO_PARTY:           &lookupList[137],
	CONTEST_MO_QSO_PARTY:           &lookupList[138],
	CONTEST_MS_QSO_PARTY:           &lookupList[139],
	CONTEST_MT_QSO_PARTY:           &lookupList[140],
	CONTEST_NA_SPRINT_CW:           &lookupList[141],
	CONTEST_NA_SPRINT_RTTY:         &lookupList[142],
	CONTEST_NA_SPRINT_SSB:          &lookupList[143],
	CONTEST_NAQP_CW:                &lookupList[144],
	CONTEST_NAQP_RTTY:              &lookupList[145],
	CONTEST_NAQP_SSB:               &lookupList[146],
	CONTEST_NAVAL:                  &lookupList[147],
	CONTEST_NC_QSO_PARTY:           &lookupList[148],
	CONTEST_ND_QSO_PARTY:           &lookupList[149],
	CONTEST_NE_QSO_PARTY:           &lookupList[150],
	CONTEST_NEQP:                   &lookupList[151],
	CONTEST_NH_QSO_PARTY:           &lookupList[152],
	CONTEST_NJ_QSO_PARTY:           &lookupList[153],
	CONTEST_NM_QSO_PARTY:           &lookupList[154],
	CONTEST_NRAU_BALTIC_CW:         &lookupList[155],
	CONTEST_NRAU_BALTIC_SSB:        &lookupList[156],
	CONTEST_NV_QSO_PARTY:           &lookupList[157],
	CONTEST_NY_QSO_PARTY:           &lookupList[158],
	CONTEST_OCEANIA_DX_CW:          &lookupList[159],
	CONTEST_OCEANIA_DX_SSB:         &lookupList[160],
	CONTEST_OH_QSO_PARTY:           &lookupList[161],
	CONTEST_OK_DX_RTTY:             &lookupList[162],
	CONTEST_OK_OM_DX:               &lookupList[163],
	CONTEST_OK_QSO_PARTY:           &lookupList[164],
	CONTEST_OMISS_QSO_PARTY:        &lookupList[165],
	CONTEST_ON_QSO_PARTY:           &lookupList[166],
	CONTEST_OR_QSO_PARTY:           &lookupList[167],
	CONTEST_ORARI_DX:               &lookupList[168],
	CONTEST_PA_QSO_PARTY:           &lookupList[169],
	CONTEST_PACC:                   &lookupList[170],
	CONTEST_PCC:                    &lookupList[171],
	CONTEST_PSK_DEATHMATCH:         &lookupList[172],
	CONTEST_QC_QSO_PARTY:           &lookupList[173],
	CONTEST_RAC:                    &lookupList[174],
	CONTEST_RAC_CANADA_DAY:         &lookupList[175],
	CONTEST_RAC_CANADA_WINTER:      &lookupList[176],
	CONTEST_RDAC:                   &lookupList[177],
	CONTEST_RDXC:                   &lookupList[178],
	CONTEST_REF_160M:               &lookupList[179],
	CONTEST_REF_CW:                 &lookupList[180],
	CONTEST_REF_SSB:                &lookupList[181],
	CONTEST_REP_PORTUGAL_DAY_HF:    &lookupList[182],
	CONTEST_RI_QSO_PARTY:           &lookupList[183],
	CONTEST_RSGB_160:               &lookupList[184],
	CONTEST_RSGB_21_28_CW:          &lookupList[185],
	CONTEST_RSGB_21_28_SSB:         &lookupList[186],
	CONTEST_RSGB_80M_CC:            &lookupList[187],
	CONTEST_RSGB_AFS_CW:            &lookupList[188],
	CONTEST_RSGB_AFS_SSB:           &lookupList[189],
	CONTEST_RSGB_CLUB_CALLS:        &lookupList[190],
	CONTEST_RSGB_COMMONWEALTH:      &lookupList[191],
	CONTEST_RSGB_IOTA:              &lookupList[192],
	CONTEST_RSGB_LOW_POWER:         &lookupList[193],
	CONTEST_RSGB_NFD:               &lookupList[194],
	CONTEST_RSGB_ROPOCO:            &lookupList[195],
	CONTEST_RSGB_SSB_FD:            &lookupList[196],
	CONTEST_RUSSIAN_RTTY:           &lookupList[197],
	CONTEST_SAC_CW:                 &lookupList[198],
	CONTEST_SAC_SSB:                &lookupList[199],
	CONTEST_SARTG_RTTY:             &lookupList[200],
	CONTEST_SC_QSO_PARTY:           &lookupList[201],
	CONTEST_SCC_RTTY:               &lookupList[202],
	CONTEST_SD_QSO_PARTY:           &lookupList[203],
	CONTEST_SHORTRY:                &lookupList[204],
	CONTEST_SMP_AUG:                &lookupList[205],
	CONTEST_SMP_MAY:                &lookupList[206],
	CONTEST_SP_DX_RTTY:             &lookupList[207],
	CONTEST_SPAR_WINTER_FD:         &lookupList[208],
	CONTEST_SPDXCONTEST:            &lookupList[209],
	CONTEST_SPRING_SPRINT:          &lookupList[210],
	CONTEST_SR_MARATHON:            &lookupList[211],
	CONTEST_STEW_PERRY:             &lookupList[212],
	CONTEST_SUMMER_SPRINT:          &lookupList[213],
	CONTEST_TARA_GRID_DIP:          &lookupList[214],
	CONTEST_TARA_RTTY:              &lookupList[215],
	CONTEST_TARA_RUMBLE:            &lookupList[216],
	CONTEST_TARA_SKIRMISH:          &lookupList[217],
	CONTEST_TEN_RTTY:               &lookupList[218],
	CONTEST_TMC_RTTY:               &lookupList[219],
	CONTEST_TN_QSO_PARTY:           &lookupList[220],
	CONTEST_TX_QSO_PARTY:           &lookupList[221],
	CONTEST_UBA_DX_CW:              &lookupList[222],
	CONTEST_UBA_DX_SSB:             &lookupList[223],
	CONTEST_UK_DX_BPSK63:           &lookupList[224],
	CONTEST_UK_DX_RTTY:             &lookupList[225],
	CONTEST_UKR_CHAMP_RTTY:         &lookupList[226],
	CONTEST_UKRAINIAN_DX:           &lookupList[227],
	CONTEST_UKSMG_6M_MARATHON:      &lookupList[228],
	CONTEST_UKSMG_SUMMER_ES:        &lookupList[229],
	CONTEST_URE_DX:                 &lookupList[230],
	CONTEST_US_COUNTIES_QSO:        &lookupList[231],
	CONTEST_UT_QSO_PARTY:           &lookupList[232],
	CONTEST_VA_QSO_PARTY:           &lookupList[233],
	CONTEST_VENEZ_IND_DAY:          &lookupList[234],
	CONTEST_VIRGINIA_QSO_PARTY:     &lookupList[235],
	CONTEST_VOLTA_RTTY:             &lookupList[236],
	CONTEST_VT_QSO_PARTY:           &lookupList[237],
	CONTEST_WA_QSO_PARTY:           &lookupList[238],
	CONTEST_WFD:                    &lookupList[239],
	CONTEST_WI_QSO_PARTY:           &lookupList[240],
	CONTEST_WIA_HARRY_ANGEL:        &lookupList[241],
	CONTEST_WIA_JMMFD:              &lookupList[242],
	CONTEST_WIA_OCDX:               &lookupList[243],
	CONTEST_WIA_REMEMBRANCE:        &lookupList[244],
	CONTEST_WIA_ROSS_HULL:          &lookupList[245],
	CONTEST_WIA_TRANS_TASMAN:       &lookupList[246],
	CONTEST_WIA_VHF_UHF_FD:         &lookupList[247],
	CONTEST_WIA_VK_SHIRES:          &lookupList[248],
	CONTEST_WINTER_SPRINT:          &lookupList[249],
	CONTEST_WV_QSO_PARTY:           &lookupList[250],
	CONTEST_WW_DIGI:                &lookupList[251],
	CONTEST_WY_QSO_PARTY:           &lookupList[252],
	CONTEST_XE_INTL_RTTY:           &lookupList[253],
	CONTEST_YOHFDX:                 &lookupList[254],
	CONTEST_YUDXC:                  &lookupList[255],
}

// Lookup returns the specification for the provided Contest.
// ADIF 3.1.6
func Lookup(c Contest) (Spec, bool) {
	spec, ok := lookupMap[c]
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
