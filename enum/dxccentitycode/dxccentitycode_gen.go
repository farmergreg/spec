// DO NOT EDIT; GENERATED CODE
// run `go generate ./...` from the project root to rebuild this file.

// Package dxccentitycode provides code and constants as defined in ADIF 3.1.6 (Proposed)
package dxccentitycode

import "sync"

const (
	NONE                                     DXCCEntityCode = 0   // 0 = None (the contacted station is known to not be within a DXCC entity)
	CANADA                                   DXCCEntityCode = 1   // 1 = CANADA
	ABU_AIL_IS_DELETED                       DXCCEntityCode = 2   // 2 = ABU AIL IS. (DELETED)
	AFGHANISTAN                              DXCCEntityCode = 3   // 3 = AFGHANISTAN
	AGALEGA_AND_ST_BRANDON_IS                DXCCEntityCode = 4   // 4 = AGALEGA & ST. BRANDON IS.
	ALAND_IS                                 DXCCEntityCode = 5   // 5 = ALAND IS.
	ALASKA                                   DXCCEntityCode = 6   // 6 = ALASKA
	ALBANIA                                  DXCCEntityCode = 7   // 7 = ALBANIA
	ALDABRA_DELETED                          DXCCEntityCode = 8   // 8 = ALDABRA (DELETED)
	AMERICAN_SAMOA                           DXCCEntityCode = 9   // 9 = AMERICAN SAMOA
	AMSTERDAM_AND_ST_PAUL_IS                 DXCCEntityCode = 10  // 10 = AMSTERDAM & ST. PAUL IS.
	ANDAMAN_AND_NICOBAR_IS                   DXCCEntityCode = 11  // 11 = ANDAMAN & NICOBAR IS.
	ANGUILLA                                 DXCCEntityCode = 12  // 12 = ANGUILLA
	ANTARCTICA                               DXCCEntityCode = 13  // 13 = ANTARCTICA
	ARMENIA                                  DXCCEntityCode = 14  // 14 = ARMENIA
	ASIATIC_RUSSIA                           DXCCEntityCode = 15  // 15 = ASIATIC RUSSIA
	NEW_ZEALAND_SUBANTARCTIC_ISLANDS         DXCCEntityCode = 16  // 16 = NEW ZEALAND SUBANTARCTIC ISLANDS
	AVES_I                                   DXCCEntityCode = 17  // 17 = AVES I.
	AZERBAIJAN                               DXCCEntityCode = 18  // 18 = AZERBAIJAN
	BAJO_NUEVO_DELETED                       DXCCEntityCode = 19  // 19 = BAJO NUEVO (DELETED)
	BAKER_AND_HOWLAND_IS                     DXCCEntityCode = 20  // 20 = BAKER & HOWLAND IS.
	BALEARIC_IS                              DXCCEntityCode = 21  // 21 = BALEARIC IS.
	PALAU                                    DXCCEntityCode = 22  // 22 = PALAU
	BLENHEIM_REEF_DELETED                    DXCCEntityCode = 23  // 23 = BLENHEIM REEF (DELETED)
	BOUVET                                   DXCCEntityCode = 24  // 24 = BOUVET
	BRITISH_NORTH_BORNEO_DELETED             DXCCEntityCode = 25  // 25 = BRITISH NORTH BORNEO (DELETED)
	BRITISH_SOMALILAND_DELETED               DXCCEntityCode = 26  // 26 = BRITISH SOMALILAND (DELETED)
	BELARUS                                  DXCCEntityCode = 27  // 27 = BELARUS
	CANAL_ZONE_DELETED                       DXCCEntityCode = 28  // 28 = CANAL ZONE (DELETED)
	CANARY_IS                                DXCCEntityCode = 29  // 29 = CANARY IS.
	CELEBE_AND_MOLUCCA_IS_DELETED            DXCCEntityCode = 30  // 30 = CELEBE & MOLUCCA IS. (DELETED)
	C_KIRIBATI_BRITISH_PHOENIX_IS            DXCCEntityCode = 31  // 31 = C. KIRIBATI (BRITISH PHOENIX IS.)
	CEUTA_AND_MELILLA                        DXCCEntityCode = 32  // 32 = CEUTA & MELILLA
	CHAGOS_IS                                DXCCEntityCode = 33  // 33 = CHAGOS IS.
	CHATHAM_IS                               DXCCEntityCode = 34  // 34 = CHATHAM IS.
	CHRISTMAS_I                              DXCCEntityCode = 35  // 35 = CHRISTMAS I.
	CLIPPERTON_I                             DXCCEntityCode = 36  // 36 = CLIPPERTON I.
	COCOS_I                                  DXCCEntityCode = 37  // 37 = COCOS I.
	COCOS_KEELING_IS                         DXCCEntityCode = 38  // 38 = COCOS (KEELING) IS.
	COMOROS_DELETED                          DXCCEntityCode = 39  // 39 = COMOROS (DELETED)
	CRETE                                    DXCCEntityCode = 40  // 40 = CRETE
	CROZET_I                                 DXCCEntityCode = 41  // 41 = CROZET I.
	DAMAO_DIU_DELETED                        DXCCEntityCode = 42  // 42 = DAMAO, DIU (DELETED)
	DESECHEO_I                               DXCCEntityCode = 43  // 43 = DESECHEO I.
	DESROCHES_DELETED                        DXCCEntityCode = 44  // 44 = DESROCHES (DELETED)
	DODECANESE                               DXCCEntityCode = 45  // 45 = DODECANESE
	EAST_MALAYSIA                            DXCCEntityCode = 46  // 46 = EAST MALAYSIA
	EASTER_I                                 DXCCEntityCode = 47  // 47 = EASTER I.
	E_KIRIBATI_LINE_IS                       DXCCEntityCode = 48  // 48 = E. KIRIBATI (LINE IS.)
	EQUATORIAL_GUINEA                        DXCCEntityCode = 49  // 49 = EQUATORIAL GUINEA
	MEXICO                                   DXCCEntityCode = 50  // 50 = MEXICO
	ERITREA                                  DXCCEntityCode = 51  // 51 = ERITREA
	ESTONIA                                  DXCCEntityCode = 52  // 52 = ESTONIA
	ETHIOPIA                                 DXCCEntityCode = 53  // 53 = ETHIOPIA
	EUROPEAN_RUSSIA                          DXCCEntityCode = 54  // 54 = EUROPEAN RUSSIA
	FARQUHAR_DELETED                         DXCCEntityCode = 55  // 55 = FARQUHAR (DELETED)
	FERNANDO_DE_NORONHA                      DXCCEntityCode = 56  // 56 = FERNANDO DE NORONHA
	FRENCH_EQUATORIAL_AFRICA_DELETED         DXCCEntityCode = 57  // 57 = FRENCH EQUATORIAL AFRICA (DELETED)
	FRENCH_INDO_CHINA_DELETED                DXCCEntityCode = 58  // 58 = FRENCH INDO-CHINA (DELETED)
	FRENCH_WEST_AFRICA_DELETED               DXCCEntityCode = 59  // 59 = FRENCH WEST AFRICA (DELETED)
	BAHAMAS                                  DXCCEntityCode = 60  // 60 = BAHAMAS
	FRANZ_JOSEF_LAND                         DXCCEntityCode = 61  // 61 = FRANZ JOSEF LAND
	BARBADOS                                 DXCCEntityCode = 62  // 62 = BARBADOS
	FRENCH_GUIANA                            DXCCEntityCode = 63  // 63 = FRENCH GUIANA
	BERMUDA                                  DXCCEntityCode = 64  // 64 = BERMUDA
	BRITISH_VIRGIN_IS                        DXCCEntityCode = 65  // 65 = BRITISH VIRGIN IS.
	BELIZE                                   DXCCEntityCode = 66  // 66 = BELIZE
	FRENCH_INDIA_DELETED                     DXCCEntityCode = 67  // 67 = FRENCH INDIA (DELETED)
	KUWAIT_SAUDI_ARABIA_NEUTRAL_ZONE_DELETED DXCCEntityCode = 68  // 68 = KUWAIT/SAUDI ARABIA NEUTRAL ZONE (DELETED)
	CAYMAN_IS                                DXCCEntityCode = 69  // 69 = CAYMAN IS.
	CUBA                                     DXCCEntityCode = 70  // 70 = CUBA
	GALAPAGOS_IS                             DXCCEntityCode = 71  // 71 = GALAPAGOS IS.
	DOMINICAN_REPUBLIC                       DXCCEntityCode = 72  // 72 = DOMINICAN REPUBLIC
	EL_SALVADOR                              DXCCEntityCode = 74  // 74 = EL SALVADOR
	GEORGIA                                  DXCCEntityCode = 75  // 75 = GEORGIA
	GUATEMALA                                DXCCEntityCode = 76  // 76 = GUATEMALA
	GRENADA                                  DXCCEntityCode = 77  // 77 = GRENADA
	HAITI                                    DXCCEntityCode = 78  // 78 = HAITI
	GUADELOUPE                               DXCCEntityCode = 79  // 79 = GUADELOUPE
	HONDURAS                                 DXCCEntityCode = 80  // 80 = HONDURAS
	GERMANY_DELETED                          DXCCEntityCode = 81  // 81 = GERMANY (DELETED)
	JAMAICA                                  DXCCEntityCode = 82  // 82 = JAMAICA
	MARTINIQUE                               DXCCEntityCode = 84  // 84 = MARTINIQUE
	BONAIRE_CURACAO_DELETED                  DXCCEntityCode = 85  // 85 = BONAIRE, CURACAO (DELETED)
	NICARAGUA                                DXCCEntityCode = 86  // 86 = NICARAGUA
	PANAMA                                   DXCCEntityCode = 88  // 88 = PANAMA
	TURKS_AND_CAICOS_IS                      DXCCEntityCode = 89  // 89 = TURKS & CAICOS IS.
	TRINIDAD_AND_TOBAGO                      DXCCEntityCode = 90  // 90 = TRINIDAD & TOBAGO
	ARUBA                                    DXCCEntityCode = 91  // 91 = ARUBA
	GEYSER_REEF_DELETED                      DXCCEntityCode = 93  // 93 = GEYSER REEF (DELETED)
	ANTIGUA_AND_BARBUDA                      DXCCEntityCode = 94  // 94 = ANTIGUA & BARBUDA
	DOMINICA                                 DXCCEntityCode = 95  // 95 = DOMINICA
	MONTSERRAT                               DXCCEntityCode = 96  // 96 = MONTSERRAT
	ST_LUCIA                                 DXCCEntityCode = 97  // 97 = ST. LUCIA
	ST_VINCENT                               DXCCEntityCode = 98  // 98 = ST. VINCENT
	GLORIOSO_IS                              DXCCEntityCode = 99  // 99 = GLORIOSO IS.
	ARGENTINA                                DXCCEntityCode = 100 // 100 = ARGENTINA
	GOA_DELETED                              DXCCEntityCode = 101 // 101 = GOA (DELETED)
	GOLD_COAST_TOGOLAND_DELETED              DXCCEntityCode = 102 // 102 = GOLD COAST, TOGOLAND (DELETED)
	GUAM                                     DXCCEntityCode = 103 // 103 = GUAM
	BOLIVIA                                  DXCCEntityCode = 104 // 104 = BOLIVIA
	GUANTANAMO_BAY                           DXCCEntityCode = 105 // 105 = GUANTANAMO BAY
	GUERNSEY                                 DXCCEntityCode = 106 // 106 = GUERNSEY
	GUINEA                                   DXCCEntityCode = 107 // 107 = GUINEA
	BRAZIL                                   DXCCEntityCode = 108 // 108 = BRAZIL
	GUINEA_BISSAU                            DXCCEntityCode = 109 // 109 = GUINEA-BISSAU
	HAWAII                                   DXCCEntityCode = 110 // 110 = HAWAII
	HEARD_I                                  DXCCEntityCode = 111 // 111 = HEARD I.
	CHILE                                    DXCCEntityCode = 112 // 112 = CHILE
	IFNI_DELETED                             DXCCEntityCode = 113 // 113 = IFNI (DELETED)
	ISLE_OF_MAN                              DXCCEntityCode = 114 // 114 = ISLE OF MAN
	ITALIAN_SOMALILAND_DELETED               DXCCEntityCode = 115 // 115 = ITALIAN SOMALILAND (DELETED)
	COLOMBIA                                 DXCCEntityCode = 116 // 116 = COLOMBIA
	ITU_HQ                                   DXCCEntityCode = 117 // 117 = ITU HQ
	JAN_MAYEN                                DXCCEntityCode = 118 // 118 = JAN MAYEN
	JAVA_DELETED                             DXCCEntityCode = 119 // 119 = JAVA (DELETED)
	ECUADOR                                  DXCCEntityCode = 120 // 120 = ECUADOR
	JERSEY                                   DXCCEntityCode = 122 // 122 = JERSEY
	JOHNSTON_I                               DXCCEntityCode = 123 // 123 = JOHNSTON I.
	JUAN_DE_NOVA_EUROPA                      DXCCEntityCode = 124 // 124 = JUAN DE NOVA, EUROPA
	JUAN_FERNANDEZ_IS                        DXCCEntityCode = 125 // 125 = JUAN FERNANDEZ IS.
	KALININGRAD                              DXCCEntityCode = 126 // 126 = KALININGRAD
	KAMARAN_IS_DELETED                       DXCCEntityCode = 127 // 127 = KAMARAN IS. (DELETED)
	KARELO_FINNISH_REPUBLIC_DELETED          DXCCEntityCode = 128 // 128 = KARELO-FINNISH REPUBLIC (DELETED)
	GUYANA                                   DXCCEntityCode = 129 // 129 = GUYANA
	KAZAKHSTAN                               DXCCEntityCode = 130 // 130 = KAZAKHSTAN
	KERGUELEN_IS                             DXCCEntityCode = 131 // 131 = KERGUELEN IS.
	PARAGUAY                                 DXCCEntityCode = 132 // 132 = PARAGUAY
	KERMADEC_IS                              DXCCEntityCode = 133 // 133 = KERMADEC IS.
	KINGMAN_REEF_DELETED                     DXCCEntityCode = 134 // 134 = KINGMAN REEF (DELETED)
	KYRGYZSTAN                               DXCCEntityCode = 135 // 135 = KYRGYZSTAN
	PERU                                     DXCCEntityCode = 136 // 136 = PERU
	REPUBLIC_OF_KOREA                        DXCCEntityCode = 137 // 137 = REPUBLIC OF KOREA
	KURE_I                                   DXCCEntityCode = 138 // 138 = KURE I.
	KURIA_MURIA_I_DELETED                    DXCCEntityCode = 139 // 139 = KURIA MURIA I. (DELETED)
	SURINAME                                 DXCCEntityCode = 140 // 140 = SURINAME
	FALKLAND_IS                              DXCCEntityCode = 141 // 141 = FALKLAND IS.
	LAKSHADWEEP_IS                           DXCCEntityCode = 142 // 142 = LAKSHADWEEP IS.
	LAOS                                     DXCCEntityCode = 143 // 143 = LAOS
	URUGUAY                                  DXCCEntityCode = 144 // 144 = URUGUAY
	LATVIA                                   DXCCEntityCode = 145 // 145 = LATVIA
	LITHUANIA                                DXCCEntityCode = 146 // 146 = LITHUANIA
	LORD_HOWE_I                              DXCCEntityCode = 147 // 147 = LORD HOWE I.
	VENEZUELA                                DXCCEntityCode = 148 // 148 = VENEZUELA
	AZORES                                   DXCCEntityCode = 149 // 149 = AZORES
	AUSTRALIA                                DXCCEntityCode = 150 // 150 = AUSTRALIA
	MALYJ_VYSOTSKIJ_I_DELETED                DXCCEntityCode = 151 // 151 = MALYJ VYSOTSKIJ I. (DELETED)
	MACAO                                    DXCCEntityCode = 152 // 152 = MACAO
	MACQUARIE_I                              DXCCEntityCode = 153 // 153 = MACQUARIE I.
	YEMEN_ARAB_REPUBLIC_DELETED              DXCCEntityCode = 154 // 154 = YEMEN ARAB REPUBLIC (DELETED)
	MALAYA_DELETED                           DXCCEntityCode = 155 // 155 = MALAYA (DELETED)
	NAURU                                    DXCCEntityCode = 157 // 157 = NAURU
	VANUATU                                  DXCCEntityCode = 158 // 158 = VANUATU
	MALDIVES                                 DXCCEntityCode = 159 // 159 = MALDIVES
	TONGA                                    DXCCEntityCode = 160 // 160 = TONGA
	MALPELO_I                                DXCCEntityCode = 161 // 161 = MALPELO I.
	NEW_CALEDONIA                            DXCCEntityCode = 162 // 162 = NEW CALEDONIA
	PAPUA_NEW_GUINEA                         DXCCEntityCode = 163 // 163 = PAPUA NEW GUINEA
	MANCHURIA_DELETED                        DXCCEntityCode = 164 // 164 = MANCHURIA (DELETED)
	MAURITIUS                                DXCCEntityCode = 165 // 165 = MAURITIUS
	MARIANA_IS                               DXCCEntityCode = 166 // 166 = MARIANA IS.
	MARKET_REEF                              DXCCEntityCode = 167 // 167 = MARKET REEF
	MARSHALL_IS                              DXCCEntityCode = 168 // 168 = MARSHALL IS.
	MAYOTTE                                  DXCCEntityCode = 169 // 169 = MAYOTTE
	NEW_ZEALAND                              DXCCEntityCode = 170 // 170 = NEW ZEALAND
	MELLISH_REEF                             DXCCEntityCode = 171 // 171 = MELLISH REEF
	PITCAIRN_I                               DXCCEntityCode = 172 // 172 = PITCAIRN I.
	MICRONESIA                               DXCCEntityCode = 173 // 173 = MICRONESIA
	MIDWAY_I                                 DXCCEntityCode = 174 // 174 = MIDWAY I.
	FRENCH_POLYNESIA                         DXCCEntityCode = 175 // 175 = FRENCH POLYNESIA
	FIJI                                     DXCCEntityCode = 176 // 176 = FIJI
	MINAMI_TORISHIMA                         DXCCEntityCode = 177 // 177 = MINAMI TORISHIMA
	MINERVA_REEF_DELETED                     DXCCEntityCode = 178 // 178 = MINERVA REEF (DELETED)
	MOLDOVA                                  DXCCEntityCode = 179 // 179 = MOLDOVA
	MOUNT_ATHOS                              DXCCEntityCode = 180 // 180 = MOUNT ATHOS
	MOZAMBIQUE                               DXCCEntityCode = 181 // 181 = MOZAMBIQUE
	NAVASSA_I                                DXCCEntityCode = 182 // 182 = NAVASSA I.
	NETHERLANDS_BORNEO_DELETED               DXCCEntityCode = 183 // 183 = NETHERLANDS BORNEO (DELETED)
	NETHERLANDS_NEW_GUINEA_DELETED           DXCCEntityCode = 184 // 184 = NETHERLANDS NEW GUINEA (DELETED)
	SOLOMON_IS                               DXCCEntityCode = 185 // 185 = SOLOMON IS.
	NEWFOUNDLAND_LABRADOR_DELETED            DXCCEntityCode = 186 // 186 = NEWFOUNDLAND, LABRADOR (DELETED)
	NIGER                                    DXCCEntityCode = 187 // 187 = NIGER
	NIUE                                     DXCCEntityCode = 188 // 188 = NIUE
	NORFOLK_I                                DXCCEntityCode = 189 // 189 = NORFOLK I.
	SAMOA                                    DXCCEntityCode = 190 // 190 = SAMOA
	NORTH_COOK_IS                            DXCCEntityCode = 191 // 191 = NORTH COOK IS.
	OGASAWARA                                DXCCEntityCode = 192 // 192 = OGASAWARA
	OKINAWA_RYUKYU_IS_DELETED                DXCCEntityCode = 193 // 193 = OKINAWA (RYUKYU IS.) (DELETED)
	OKINO_TORI_SHIMA_DELETED                 DXCCEntityCode = 194 // 194 = OKINO TORI-SHIMA (DELETED)
	ANNOBON_I                                DXCCEntityCode = 195 // 195 = ANNOBON I.
	PALESTINE_DELETED                        DXCCEntityCode = 196 // 196 = PALESTINE (DELETED)
	PALMYRA_AND_JARVIS_IS                    DXCCEntityCode = 197 // 197 = PALMYRA & JARVIS IS.
	PAPUA_TERRITORY_DELETED                  DXCCEntityCode = 198 // 198 = PAPUA TERRITORY (DELETED)
	PETER_1_I                                DXCCEntityCode = 199 // 199 = PETER 1 I.
	PORTUGUESE_TIMOR_DELETED                 DXCCEntityCode = 200 // 200 = PORTUGUESE TIMOR (DELETED)
	PRINCE_EDWARD_AND_MARION_IS              DXCCEntityCode = 201 // 201 = PRINCE EDWARD & MARION IS.
	PUERTO_RICO                              DXCCEntityCode = 202 // 202 = PUERTO RICO
	ANDORRA                                  DXCCEntityCode = 203 // 203 = ANDORRA
	REVILLAGIGEDO                            DXCCEntityCode = 204 // 204 = REVILLAGIGEDO
	ASCENSION_I                              DXCCEntityCode = 205 // 205 = ASCENSION I.
	AUSTRIA                                  DXCCEntityCode = 206 // 206 = AUSTRIA
	RODRIGUES_I                              DXCCEntityCode = 207 // 207 = RODRIGUES I.
	RUANDA_URUNDI_DELETED                    DXCCEntityCode = 208 // 208 = RUANDA-URUNDI (DELETED)
	BELGIUM                                  DXCCEntityCode = 209 // 209 = BELGIUM
	SAAR_DELETED                             DXCCEntityCode = 210 // 210 = SAAR (DELETED)
	SABLE_I                                  DXCCEntityCode = 211 // 211 = SABLE I.
	BULGARIA                                 DXCCEntityCode = 212 // 212 = BULGARIA
	SAINT_MARTIN                             DXCCEntityCode = 213 // 213 = SAINT MARTIN
	CORSICA                                  DXCCEntityCode = 214 // 214 = CORSICA
	CYPRUS                                   DXCCEntityCode = 215 // 215 = CYPRUS
	SAN_ANDRES_AND_PROVIDENCIA               DXCCEntityCode = 216 // 216 = SAN ANDRES & PROVIDENCIA
	SAN_FELIX_AND_SAN_AMBROSIO               DXCCEntityCode = 217 // 217 = SAN FELIX & SAN AMBROSIO
	CZECHOSLOVAKIA_DELETED                   DXCCEntityCode = 218 // 218 = CZECHOSLOVAKIA (DELETED)
	SAO_TOME_AND_PRINCIPE                    DXCCEntityCode = 219 // 219 = SAO TOME & PRINCIPE
	SARAWAK_DELETED                          DXCCEntityCode = 220 // 220 = SARAWAK (DELETED)
	DENMARK                                  DXCCEntityCode = 221 // 221 = DENMARK
	FAROE_IS                                 DXCCEntityCode = 222 // 222 = FAROE IS.
	ENGLAND                                  DXCCEntityCode = 223 // 223 = ENGLAND
	FINLAND                                  DXCCEntityCode = 224 // 224 = FINLAND
	SARDINIA                                 DXCCEntityCode = 225 // 225 = SARDINIA
	SAUDI_ARABIA_IRAQ_NEUTRAL_ZONE_DELETED   DXCCEntityCode = 226 // 226 = SAUDI ARABIA/IRAQ NEUTRAL ZONE (DELETED)
	FRANCE                                   DXCCEntityCode = 227 // 227 = FRANCE
	SERRANA_BANK_AND_RONCADOR_CAY_DELETED    DXCCEntityCode = 228 // 228 = SERRANA BANK & RONCADOR CAY (DELETED)
	GERMAN_DEMOCRATIC_REPUBLIC_DELETED       DXCCEntityCode = 229 // 229 = GERMAN DEMOCRATIC REPUBLIC (DELETED)
	FEDERAL_REPUBLIC_OF_GERMANY              DXCCEntityCode = 230 // 230 = FEDERAL REPUBLIC OF GERMANY
	SIKKIM_DELETED                           DXCCEntityCode = 231 // 231 = SIKKIM (DELETED)
	SOMALIA                                  DXCCEntityCode = 232 // 232 = SOMALIA
	GIBRALTAR                                DXCCEntityCode = 233 // 233 = GIBRALTAR
	SOUTH_COOK_IS                            DXCCEntityCode = 234 // 234 = SOUTH COOK IS.
	SOUTH_GEORGIA_I                          DXCCEntityCode = 235 // 235 = SOUTH GEORGIA I.
	GREECE                                   DXCCEntityCode = 236 // 236 = GREECE
	GREENLAND                                DXCCEntityCode = 237 // 237 = GREENLAND
	SOUTH_ORKNEY_IS                          DXCCEntityCode = 238 // 238 = SOUTH ORKNEY IS.
	HUNGARY                                  DXCCEntityCode = 239 // 239 = HUNGARY
	SOUTH_SANDWICH_IS                        DXCCEntityCode = 240 // 240 = SOUTH SANDWICH IS.
	SOUTH_SHETLAND_IS                        DXCCEntityCode = 241 // 241 = SOUTH SHETLAND IS.
	ICELAND                                  DXCCEntityCode = 242 // 242 = ICELAND
	PEOPLES_DEMOCRATIC_REP_OF_YEMEN_DELETED  DXCCEntityCode = 243 // 243 = PEOPLE'S DEMOCRATIC REP. OF YEMEN (DELETED)
	SOUTHERN_SUDAN_DELETED                   DXCCEntityCode = 244 // 244 = SOUTHERN SUDAN (DELETED)
	IRELAND                                  DXCCEntityCode = 245 // 245 = IRELAND
	SOVEREIGN_MILITARY_ORDER_OF_MALTA        DXCCEntityCode = 246 // 246 = SOVEREIGN MILITARY ORDER OF MALTA
	SPRATLY_IS                               DXCCEntityCode = 247 // 247 = SPRATLY IS.
	ITALY                                    DXCCEntityCode = 248 // 248 = ITALY
	ST_KITTS_AND_NEVIS                       DXCCEntityCode = 249 // 249 = ST. KITTS & NEVIS
	ST_HELENA                                DXCCEntityCode = 250 // 250 = ST. HELENA
	LIECHTENSTEIN                            DXCCEntityCode = 251 // 251 = LIECHTENSTEIN
	ST_PAUL_I                                DXCCEntityCode = 252 // 252 = ST. PAUL I.
	ST_PETER_AND_ST_PAUL_ROCKS               DXCCEntityCode = 253 // 253 = ST. PETER & ST. PAUL ROCKS
	LUXEMBOURG                               DXCCEntityCode = 254 // 254 = LUXEMBOURG
	ST_MAARTEN_SABA_ST_EUSTATIUS_DELETED     DXCCEntityCode = 255 // 255 = ST. MAARTEN, SABA, ST. EUSTATIUS (DELETED)
	MADEIRA_IS                               DXCCEntityCode = 256 // 256 = MADEIRA IS.
	MALTA                                    DXCCEntityCode = 257 // 257 = MALTA
	SUMATRA_DELETED                          DXCCEntityCode = 258 // 258 = SUMATRA (DELETED)
	SVALBARD                                 DXCCEntityCode = 259 // 259 = SVALBARD
	MONACO                                   DXCCEntityCode = 260 // 260 = MONACO
	SWAN_IS_DELETED                          DXCCEntityCode = 261 // 261 = SWAN IS. (DELETED)
	TAJIKISTAN                               DXCCEntityCode = 262 // 262 = TAJIKISTAN
	NETHERLANDS                              DXCCEntityCode = 263 // 263 = NETHERLANDS
	TANGIER_DELETED                          DXCCEntityCode = 264 // 264 = TANGIER (DELETED)
	NORTHERN_IRELAND                         DXCCEntityCode = 265 // 265 = NORTHERN IRELAND
	NORWAY                                   DXCCEntityCode = 266 // 266 = NORWAY
	TERRITORY_OF_NEW_GUINEA_DELETED          DXCCEntityCode = 267 // 267 = TERRITORY OF NEW GUINEA (DELETED)
	TIBET_DELETED                            DXCCEntityCode = 268 // 268 = TIBET (DELETED)
	POLAND                                   DXCCEntityCode = 269 // 269 = POLAND
	TOKELAU_IS                               DXCCEntityCode = 270 // 270 = TOKELAU IS.
	TRIESTE_DELETED                          DXCCEntityCode = 271 // 271 = TRIESTE (DELETED)
	PORTUGAL                                 DXCCEntityCode = 272 // 272 = PORTUGAL
	TRINDADE_AND_MARTIM_VAZ_IS               DXCCEntityCode = 273 // 273 = TRINDADE & MARTIM VAZ IS.
	TRISTAN_DA_CUNHA_AND_GOUGH_I             DXCCEntityCode = 274 // 274 = TRISTAN DA CUNHA & GOUGH I.
	ROMANIA                                  DXCCEntityCode = 275 // 275 = ROMANIA
	TROMELIN_I                               DXCCEntityCode = 276 // 276 = TROMELIN I.
	ST_PIERRE_AND_MIQUELON                   DXCCEntityCode = 277 // 277 = ST. PIERRE & MIQUELON
	SAN_MARINO                               DXCCEntityCode = 278 // 278 = SAN MARINO
	SCOTLAND                                 DXCCEntityCode = 279 // 279 = SCOTLAND
	TURKMENISTAN                             DXCCEntityCode = 280 // 280 = TURKMENISTAN
	SPAIN                                    DXCCEntityCode = 281 // 281 = SPAIN
	TUVALU                                   DXCCEntityCode = 282 // 282 = TUVALU
	UK_SOVEREIGN_BASE_AREAS_ON_CYPRUS        DXCCEntityCode = 283 // 283 = UK SOVEREIGN BASE AREAS ON CYPRUS
	SWEDEN                                   DXCCEntityCode = 284 // 284 = SWEDEN
	VIRGIN_IS                                DXCCEntityCode = 285 // 285 = VIRGIN IS.
	UGANDA                                   DXCCEntityCode = 286 // 286 = UGANDA
	SWITZERLAND                              DXCCEntityCode = 287 // 287 = SWITZERLAND
	UKRAINE                                  DXCCEntityCode = 288 // 288 = UKRAINE
	UNITED_NATIONS_HQ                        DXCCEntityCode = 289 // 289 = UNITED NATIONS HQ
	UNITED_STATES_OF_AMERICA                 DXCCEntityCode = 291 // 291 = UNITED STATES OF AMERICA
	UZBEKISTAN                               DXCCEntityCode = 292 // 292 = UZBEKISTAN
	VIET_NAM                                 DXCCEntityCode = 293 // 293 = VIET NAM
	WALES                                    DXCCEntityCode = 294 // 294 = WALES
	VATICAN                                  DXCCEntityCode = 295 // 295 = VATICAN
	SERBIA                                   DXCCEntityCode = 296 // 296 = SERBIA
	WAKE_I                                   DXCCEntityCode = 297 // 297 = WAKE I.
	WALLIS_AND_FUTUNA_IS                     DXCCEntityCode = 298 // 298 = WALLIS & FUTUNA IS.
	WEST_MALAYSIA                            DXCCEntityCode = 299 // 299 = WEST MALAYSIA
	W_KIRIBATI_GILBERT_IS                    DXCCEntityCode = 301 // 301 = W. KIRIBATI (GILBERT IS. )
	WESTERN_SAHARA                           DXCCEntityCode = 302 // 302 = WESTERN SAHARA
	WILLIS_I                                 DXCCEntityCode = 303 // 303 = WILLIS I.
	BAHRAIN                                  DXCCEntityCode = 304 // 304 = BAHRAIN
	BANGLADESH                               DXCCEntityCode = 305 // 305 = BANGLADESH
	BHUTAN                                   DXCCEntityCode = 306 // 306 = BHUTAN
	ZANZIBAR_DELETED                         DXCCEntityCode = 307 // 307 = ZANZIBAR (DELETED)
	COSTA_RICA                               DXCCEntityCode = 308 // 308 = COSTA RICA
	MYANMAR                                  DXCCEntityCode = 309 // 309 = MYANMAR
	CAMBODIA                                 DXCCEntityCode = 312 // 312 = CAMBODIA
	SRI_LANKA                                DXCCEntityCode = 315 // 315 = SRI LANKA
	CHINA                                    DXCCEntityCode = 318 // 318 = CHINA
	HONG_KONG                                DXCCEntityCode = 321 // 321 = HONG KONG
	INDIA                                    DXCCEntityCode = 324 // 324 = INDIA
	INDONESIA                                DXCCEntityCode = 327 // 327 = INDONESIA
	IRAN                                     DXCCEntityCode = 330 // 330 = IRAN
	IRAQ                                     DXCCEntityCode = 333 // 333 = IRAQ
	ISRAEL                                   DXCCEntityCode = 336 // 336 = ISRAEL
	JAPAN                                    DXCCEntityCode = 339 // 339 = JAPAN
	JORDAN                                   DXCCEntityCode = 342 // 342 = JORDAN
	DEMOCRATIC_PEOPLES_REP_OF_KOREA          DXCCEntityCode = 344 // 344 = DEMOCRATIC PEOPLE'S REP. OF KOREA
	BRUNEI_DARUSSALAM                        DXCCEntityCode = 345 // 345 = BRUNEI DARUSSALAM
	KUWAIT                                   DXCCEntityCode = 348 // 348 = KUWAIT
	LEBANON                                  DXCCEntityCode = 354 // 354 = LEBANON
	MONGOLIA                                 DXCCEntityCode = 363 // 363 = MONGOLIA
	NEPAL                                    DXCCEntityCode = 369 // 369 = NEPAL
	OMAN                                     DXCCEntityCode = 370 // 370 = OMAN
	PAKISTAN                                 DXCCEntityCode = 372 // 372 = PAKISTAN
	PHILIPPINES                              DXCCEntityCode = 375 // 375 = PHILIPPINES
	QATAR                                    DXCCEntityCode = 376 // 376 = QATAR
	SAUDI_ARABIA                             DXCCEntityCode = 378 // 378 = SAUDI ARABIA
	SEYCHELLES                               DXCCEntityCode = 379 // 379 = SEYCHELLES
	SINGAPORE                                DXCCEntityCode = 381 // 381 = SINGAPORE
	DJIBOUTI                                 DXCCEntityCode = 382 // 382 = DJIBOUTI
	SYRIA                                    DXCCEntityCode = 384 // 384 = SYRIA
	TAIWAN                                   DXCCEntityCode = 386 // 386 = TAIWAN
	THAILAND                                 DXCCEntityCode = 387 // 387 = THAILAND
	TURKEY                                   DXCCEntityCode = 390 // 390 = TURKEY
	UNITED_ARAB_EMIRATES                     DXCCEntityCode = 391 // 391 = UNITED ARAB EMIRATES
	ALGERIA                                  DXCCEntityCode = 400 // 400 = ALGERIA
	ANGOLA                                   DXCCEntityCode = 401 // 401 = ANGOLA
	BOTSWANA                                 DXCCEntityCode = 402 // 402 = BOTSWANA
	BURUNDI                                  DXCCEntityCode = 404 // 404 = BURUNDI
	CAMEROON                                 DXCCEntityCode = 406 // 406 = CAMEROON
	CENTRAL_AFRICA                           DXCCEntityCode = 408 // 408 = CENTRAL AFRICA
	CAPE_VERDE                               DXCCEntityCode = 409 // 409 = CAPE VERDE
	CHAD                                     DXCCEntityCode = 410 // 410 = CHAD
	COMOROS                                  DXCCEntityCode = 411 // 411 = COMOROS
	REPUBLIC_OF_THE_CONGO                    DXCCEntityCode = 412 // 412 = REPUBLIC OF THE CONGO
	DEMOCRATIC_REPUBLIC_OF_THE_CONGO         DXCCEntityCode = 414 // 414 = DEMOCRATIC REPUBLIC OF THE CONGO
	BENIN                                    DXCCEntityCode = 416 // 416 = BENIN
	GABON                                    DXCCEntityCode = 420 // 420 = GABON
	THE_GAMBIA                               DXCCEntityCode = 422 // 422 = THE GAMBIA
	GHANA                                    DXCCEntityCode = 424 // 424 = GHANA
	COTE_DIVOIRE                             DXCCEntityCode = 428 // 428 = COTE D'IVOIRE
	KENYA                                    DXCCEntityCode = 430 // 430 = KENYA
	LESOTHO                                  DXCCEntityCode = 432 // 432 = LESOTHO
	LIBERIA                                  DXCCEntityCode = 434 // 434 = LIBERIA
	LIBYA                                    DXCCEntityCode = 436 // 436 = LIBYA
	MADAGASCAR                               DXCCEntityCode = 438 // 438 = MADAGASCAR
	MALAWI                                   DXCCEntityCode = 440 // 440 = MALAWI
	MALI                                     DXCCEntityCode = 442 // 442 = MALI
	MAURITANIA                               DXCCEntityCode = 444 // 444 = MAURITANIA
	MOROCCO                                  DXCCEntityCode = 446 // 446 = MOROCCO
	NIGERIA                                  DXCCEntityCode = 450 // 450 = NIGERIA
	ZIMBABWE                                 DXCCEntityCode = 452 // 452 = ZIMBABWE
	REUNION_I                                DXCCEntityCode = 453 // 453 = REUNION I.
	RWANDA                                   DXCCEntityCode = 454 // 454 = RWANDA
	SENEGAL                                  DXCCEntityCode = 456 // 456 = SENEGAL
	SIERRA_LEONE                             DXCCEntityCode = 458 // 458 = SIERRA LEONE
	ROTUMA_I                                 DXCCEntityCode = 460 // 460 = ROTUMA I.
	REPUBLIC_OF_SOUTH_AFRICA                 DXCCEntityCode = 462 // 462 = REPUBLIC OF SOUTH AFRICA
	NAMIBIA                                  DXCCEntityCode = 464 // 464 = NAMIBIA
	SUDAN                                    DXCCEntityCode = 466 // 466 = SUDAN
	KINGDOM_OF_ESWATINI                      DXCCEntityCode = 468 // 468 = KINGDOM OF ESWATINI
	TANZANIA                                 DXCCEntityCode = 470 // 470 = TANZANIA
	TUNISIA                                  DXCCEntityCode = 474 // 474 = TUNISIA
	EGYPT                                    DXCCEntityCode = 478 // 478 = EGYPT
	BURKINA_FASO                             DXCCEntityCode = 480 // 480 = BURKINA FASO
	ZAMBIA                                   DXCCEntityCode = 482 // 482 = ZAMBIA
	TOGO                                     DXCCEntityCode = 483 // 483 = TOGO
	WALVIS_BAY_DELETED                       DXCCEntityCode = 488 // 488 = WALVIS BAY (DELETED)
	CONWAY_REEF                              DXCCEntityCode = 489 // 489 = CONWAY REEF
	BANABA_I_OCEAN_I                         DXCCEntityCode = 490 // 490 = BANABA I. (OCEAN I.)
	YEMEN                                    DXCCEntityCode = 492 // 492 = YEMEN
	PENGUIN_IS_DELETED                       DXCCEntityCode = 493 // 493 = PENGUIN IS. (DELETED)
	CROATIA                                  DXCCEntityCode = 497 // 497 = CROATIA
	SLOVENIA                                 DXCCEntityCode = 499 // 499 = SLOVENIA
	BOSNIA_HERZEGOVINA                       DXCCEntityCode = 501 // 501 = BOSNIA-HERZEGOVINA
	NORTH_MACEDONIA_REPUBLIC_OF              DXCCEntityCode = 502 // 502 = NORTH MACEDONIA (REPUBLIC OF)
	CZECH_REPUBLIC                           DXCCEntityCode = 503 // 503 = CZECH REPUBLIC
	SLOVAK_REPUBLIC                          DXCCEntityCode = 504 // 504 = SLOVAK REPUBLIC
	PRATAS_I                                 DXCCEntityCode = 505 // 505 = PRATAS I.
	SCARBOROUGH_REEF                         DXCCEntityCode = 506 // 506 = SCARBOROUGH REEF
	TEMOTU_PROVINCE                          DXCCEntityCode = 507 // 507 = TEMOTU PROVINCE
	AUSTRAL_I                                DXCCEntityCode = 508 // 508 = AUSTRAL I.
	MARQUESAS_IS                             DXCCEntityCode = 509 // 509 = MARQUESAS IS.
	PALESTINE                                DXCCEntityCode = 510 // 510 = PALESTINE
	TIMOR_LESTE                              DXCCEntityCode = 511 // 511 = TIMOR-LESTE
	CHESTERFIELD_IS                          DXCCEntityCode = 512 // 512 = CHESTERFIELD IS.
	DUCIE_I                                  DXCCEntityCode = 513 // 513 = DUCIE I.
	MONTENEGRO                               DXCCEntityCode = 514 // 514 = MONTENEGRO
	SWAINS_I                                 DXCCEntityCode = 515 // 515 = SWAINS I.
	SAINT_BARTHELEMY                         DXCCEntityCode = 516 // 516 = SAINT BARTHELEMY
	CURACAO                                  DXCCEntityCode = 517 // 517 = CURACAO
	SINT_MAARTEN                             DXCCEntityCode = 518 // 518 = SINT MAARTEN
	SABA_AND_ST_EUSTATIUS                    DXCCEntityCode = 519 // 519 = SABA & ST. EUSTATIUS
	BONAIRE                                  DXCCEntityCode = 520 // 520 = BONAIRE
	SOUTH_SUDAN_REPUBLIC_OF                  DXCCEntityCode = 521 // 521 = SOUTH SUDAN (REPUBLIC OF)
	REPUBLIC_OF_KOSOVO                       DXCCEntityCode = 522 // 522 = REPUBLIC OF KOSOVO
)

var (
	listActive     []Spec    // listActive is a cached copy of the active specifications (those not marked as import-only).
	listActiveOnce sync.Once // listActive is lazy loaded instead of utilizing an init() function. This allows the compiler to remove unused data / variables.
)

// lookupList contains all known DXCCEntityCode specifications.
var lookupList = []Spec{
	{IsImportOnly: false, Key: 0, EntityName: "None (the contacted station is known to not be within a DXCC entity)", IsDeleted: false},
	{IsImportOnly: false, Key: 1, EntityName: "CANADA", IsDeleted: false},
	{IsImportOnly: false, Key: 2, EntityName: "ABU AIL IS.", IsDeleted: true},
	{IsImportOnly: false, Key: 3, EntityName: "AFGHANISTAN", IsDeleted: false},
	{IsImportOnly: false, Key: 4, EntityName: "AGALEGA & ST. BRANDON IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 5, EntityName: "ALAND IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 6, EntityName: "ALASKA", IsDeleted: false},
	{IsImportOnly: false, Key: 7, EntityName: "ALBANIA", IsDeleted: false},
	{IsImportOnly: false, Key: 8, EntityName: "ALDABRA", IsDeleted: true},
	{IsImportOnly: false, Key: 9, EntityName: "AMERICAN SAMOA", IsDeleted: false},
	{IsImportOnly: false, Key: 10, EntityName: "AMSTERDAM & ST. PAUL IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 11, EntityName: "ANDAMAN & NICOBAR IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 12, EntityName: "ANGUILLA", IsDeleted: false},
	{IsImportOnly: false, Key: 13, EntityName: "ANTARCTICA", IsDeleted: false},
	{IsImportOnly: false, Key: 14, EntityName: "ARMENIA", IsDeleted: false},
	{IsImportOnly: false, Key: 15, EntityName: "ASIATIC RUSSIA", IsDeleted: false},
	{IsImportOnly: false, Key: 16, EntityName: "NEW ZEALAND SUBANTARCTIC ISLANDS", IsDeleted: false},
	{IsImportOnly: false, Key: 17, EntityName: "AVES I.", IsDeleted: false},
	{IsImportOnly: false, Key: 18, EntityName: "AZERBAIJAN", IsDeleted: false},
	{IsImportOnly: false, Key: 19, EntityName: "BAJO NUEVO", IsDeleted: true},
	{IsImportOnly: false, Key: 20, EntityName: "BAKER & HOWLAND IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 21, EntityName: "BALEARIC IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 22, EntityName: "PALAU", IsDeleted: false},
	{IsImportOnly: false, Key: 23, EntityName: "BLENHEIM REEF", IsDeleted: true},
	{IsImportOnly: false, Key: 24, EntityName: "BOUVET", IsDeleted: false},
	{IsImportOnly: false, Key: 25, EntityName: "BRITISH NORTH BORNEO", IsDeleted: true},
	{IsImportOnly: false, Key: 26, EntityName: "BRITISH SOMALILAND", IsDeleted: true},
	{IsImportOnly: false, Key: 27, EntityName: "BELARUS", IsDeleted: false},
	{IsImportOnly: false, Key: 28, EntityName: "CANAL ZONE", IsDeleted: true},
	{IsImportOnly: false, Key: 29, EntityName: "CANARY IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 30, EntityName: "CELEBE & MOLUCCA IS.", IsDeleted: true},
	{IsImportOnly: false, Key: 31, EntityName: "C. KIRIBATI (BRITISH PHOENIX IS.)", IsDeleted: false},
	{IsImportOnly: false, Key: 32, EntityName: "CEUTA & MELILLA", IsDeleted: false},
	{IsImportOnly: false, Key: 33, EntityName: "CHAGOS IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 34, EntityName: "CHATHAM IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 35, EntityName: "CHRISTMAS I.", IsDeleted: false},
	{IsImportOnly: false, Key: 36, EntityName: "CLIPPERTON I.", IsDeleted: false},
	{IsImportOnly: false, Key: 37, EntityName: "COCOS I.", IsDeleted: false},
	{IsImportOnly: false, Key: 38, EntityName: "COCOS (KEELING) IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 39, EntityName: "COMOROS", IsDeleted: true},
	{IsImportOnly: false, Key: 40, EntityName: "CRETE", IsDeleted: false},
	{IsImportOnly: false, Key: 41, EntityName: "CROZET I.", IsDeleted: false},
	{IsImportOnly: false, Key: 42, EntityName: "DAMAO, DIU", IsDeleted: true},
	{IsImportOnly: false, Key: 43, EntityName: "DESECHEO I.", IsDeleted: false},
	{IsImportOnly: false, Key: 44, EntityName: "DESROCHES", IsDeleted: true},
	{IsImportOnly: false, Key: 45, EntityName: "DODECANESE", IsDeleted: false},
	{IsImportOnly: false, Key: 46, EntityName: "EAST MALAYSIA", IsDeleted: false},
	{IsImportOnly: false, Key: 47, EntityName: "EASTER I.", IsDeleted: false},
	{IsImportOnly: false, Key: 48, EntityName: "E. KIRIBATI (LINE IS.)", IsDeleted: false},
	{IsImportOnly: false, Key: 49, EntityName: "EQUATORIAL GUINEA", IsDeleted: false},
	{IsImportOnly: false, Key: 50, EntityName: "MEXICO", IsDeleted: false},
	{IsImportOnly: false, Key: 51, EntityName: "ERITREA", IsDeleted: false},
	{IsImportOnly: false, Key: 52, EntityName: "ESTONIA", IsDeleted: false},
	{IsImportOnly: false, Key: 53, EntityName: "ETHIOPIA", IsDeleted: false},
	{IsImportOnly: false, Key: 54, EntityName: "EUROPEAN RUSSIA", IsDeleted: false},
	{IsImportOnly: false, Key: 55, EntityName: "FARQUHAR", IsDeleted: true},
	{IsImportOnly: false, Key: 56, EntityName: "FERNANDO DE NORONHA", IsDeleted: false},
	{IsImportOnly: false, Key: 57, EntityName: "FRENCH EQUATORIAL AFRICA", IsDeleted: true},
	{IsImportOnly: false, Key: 58, EntityName: "FRENCH INDO-CHINA", IsDeleted: true},
	{IsImportOnly: false, Key: 59, EntityName: "FRENCH WEST AFRICA", IsDeleted: true},
	{IsImportOnly: false, Key: 60, EntityName: "BAHAMAS", IsDeleted: false},
	{IsImportOnly: false, Key: 61, EntityName: "FRANZ JOSEF LAND", IsDeleted: false},
	{IsImportOnly: false, Key: 62, EntityName: "BARBADOS", IsDeleted: false},
	{IsImportOnly: false, Key: 63, EntityName: "FRENCH GUIANA", IsDeleted: false},
	{IsImportOnly: false, Key: 64, EntityName: "BERMUDA", IsDeleted: false},
	{IsImportOnly: false, Key: 65, EntityName: "BRITISH VIRGIN IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 66, EntityName: "BELIZE", IsDeleted: false},
	{IsImportOnly: false, Key: 67, EntityName: "FRENCH INDIA", IsDeleted: true},
	{IsImportOnly: false, Key: 68, EntityName: "KUWAIT/SAUDI ARABIA NEUTRAL ZONE", IsDeleted: true},
	{IsImportOnly: false, Key: 69, EntityName: "CAYMAN IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 70, EntityName: "CUBA", IsDeleted: false},
	{IsImportOnly: false, Key: 71, EntityName: "GALAPAGOS IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 72, EntityName: "DOMINICAN REPUBLIC", IsDeleted: false},
	{IsImportOnly: false, Key: 74, EntityName: "EL SALVADOR", IsDeleted: false},
	{IsImportOnly: false, Key: 75, EntityName: "GEORGIA", IsDeleted: false},
	{IsImportOnly: false, Key: 76, EntityName: "GUATEMALA", IsDeleted: false},
	{IsImportOnly: false, Key: 77, EntityName: "GRENADA", IsDeleted: false},
	{IsImportOnly: false, Key: 78, EntityName: "HAITI", IsDeleted: false},
	{IsImportOnly: false, Key: 79, EntityName: "GUADELOUPE", IsDeleted: false},
	{IsImportOnly: false, Key: 80, EntityName: "HONDURAS", IsDeleted: false},
	{IsImportOnly: false, Key: 81, EntityName: "GERMANY", IsDeleted: true},
	{IsImportOnly: false, Key: 82, EntityName: "JAMAICA", IsDeleted: false},
	{IsImportOnly: false, Key: 84, EntityName: "MARTINIQUE", IsDeleted: false},
	{IsImportOnly: false, Key: 85, EntityName: "BONAIRE, CURACAO", IsDeleted: true},
	{IsImportOnly: false, Key: 86, EntityName: "NICARAGUA", IsDeleted: false},
	{IsImportOnly: false, Key: 88, EntityName: "PANAMA", IsDeleted: false},
	{IsImportOnly: false, Key: 89, EntityName: "TURKS & CAICOS IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 90, EntityName: "TRINIDAD & TOBAGO", IsDeleted: false},
	{IsImportOnly: false, Key: 91, EntityName: "ARUBA", IsDeleted: false},
	{IsImportOnly: false, Key: 93, EntityName: "GEYSER REEF", IsDeleted: true},
	{IsImportOnly: false, Key: 94, EntityName: "ANTIGUA & BARBUDA", IsDeleted: false},
	{IsImportOnly: false, Key: 95, EntityName: "DOMINICA", IsDeleted: false},
	{IsImportOnly: false, Key: 96, EntityName: "MONTSERRAT", IsDeleted: false},
	{IsImportOnly: false, Key: 97, EntityName: "ST. LUCIA", IsDeleted: false},
	{IsImportOnly: false, Key: 98, EntityName: "ST. VINCENT", IsDeleted: false},
	{IsImportOnly: false, Key: 99, EntityName: "GLORIOSO IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 100, EntityName: "ARGENTINA", IsDeleted: false},
	{IsImportOnly: false, Key: 101, EntityName: "GOA", IsDeleted: true},
	{IsImportOnly: false, Key: 102, EntityName: "GOLD COAST, TOGOLAND", IsDeleted: true},
	{IsImportOnly: false, Key: 103, EntityName: "GUAM", IsDeleted: false},
	{IsImportOnly: false, Key: 104, EntityName: "BOLIVIA", IsDeleted: false},
	{IsImportOnly: false, Key: 105, EntityName: "GUANTANAMO BAY", IsDeleted: false},
	{IsImportOnly: false, Key: 106, EntityName: "GUERNSEY", IsDeleted: false},
	{IsImportOnly: false, Key: 107, EntityName: "GUINEA", IsDeleted: false},
	{IsImportOnly: false, Key: 108, EntityName: "BRAZIL", IsDeleted: false},
	{IsImportOnly: false, Key: 109, EntityName: "GUINEA-BISSAU", IsDeleted: false},
	{IsImportOnly: false, Key: 110, EntityName: "HAWAII", IsDeleted: false},
	{IsImportOnly: false, Key: 111, EntityName: "HEARD I.", IsDeleted: false},
	{IsImportOnly: false, Key: 112, EntityName: "CHILE", IsDeleted: false},
	{IsImportOnly: false, Key: 113, EntityName: "IFNI", IsDeleted: true},
	{IsImportOnly: false, Key: 114, EntityName: "ISLE OF MAN", IsDeleted: false},
	{IsImportOnly: false, Key: 115, EntityName: "ITALIAN SOMALILAND", IsDeleted: true},
	{IsImportOnly: false, Key: 116, EntityName: "COLOMBIA", IsDeleted: false},
	{IsImportOnly: false, Key: 117, EntityName: "ITU HQ", IsDeleted: false},
	{IsImportOnly: false, Key: 118, EntityName: "JAN MAYEN", IsDeleted: false},
	{IsImportOnly: false, Key: 119, EntityName: "JAVA", IsDeleted: true},
	{IsImportOnly: false, Key: 120, EntityName: "ECUADOR", IsDeleted: false},
	{IsImportOnly: false, Key: 122, EntityName: "JERSEY", IsDeleted: false},
	{IsImportOnly: false, Key: 123, EntityName: "JOHNSTON I.", IsDeleted: false},
	{IsImportOnly: false, Key: 124, EntityName: "JUAN DE NOVA, EUROPA", IsDeleted: false},
	{IsImportOnly: false, Key: 125, EntityName: "JUAN FERNANDEZ IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 126, EntityName: "KALININGRAD", IsDeleted: false},
	{IsImportOnly: false, Key: 127, EntityName: "KAMARAN IS.", IsDeleted: true},
	{IsImportOnly: false, Key: 128, EntityName: "KARELO-FINNISH REPUBLIC", IsDeleted: true},
	{IsImportOnly: false, Key: 129, EntityName: "GUYANA", IsDeleted: false},
	{IsImportOnly: false, Key: 130, EntityName: "KAZAKHSTAN", IsDeleted: false},
	{IsImportOnly: false, Key: 131, EntityName: "KERGUELEN IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 132, EntityName: "PARAGUAY", IsDeleted: false},
	{IsImportOnly: false, Key: 133, EntityName: "KERMADEC IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 134, EntityName: "KINGMAN REEF", IsDeleted: true},
	{IsImportOnly: false, Key: 135, EntityName: "KYRGYZSTAN", IsDeleted: false},
	{IsImportOnly: false, Key: 136, EntityName: "PERU", IsDeleted: false},
	{IsImportOnly: false, Key: 137, EntityName: "REPUBLIC OF KOREA", IsDeleted: false},
	{IsImportOnly: false, Key: 138, EntityName: "KURE I.", IsDeleted: false},
	{IsImportOnly: false, Key: 139, EntityName: "KURIA MURIA I.", IsDeleted: true},
	{IsImportOnly: false, Key: 140, EntityName: "SURINAME", IsDeleted: false},
	{IsImportOnly: false, Key: 141, EntityName: "FALKLAND IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 142, EntityName: "LAKSHADWEEP IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 143, EntityName: "LAOS", IsDeleted: false},
	{IsImportOnly: false, Key: 144, EntityName: "URUGUAY", IsDeleted: false},
	{IsImportOnly: false, Key: 145, EntityName: "LATVIA", IsDeleted: false},
	{IsImportOnly: false, Key: 146, EntityName: "LITHUANIA", IsDeleted: false},
	{IsImportOnly: false, Key: 147, EntityName: "LORD HOWE I.", IsDeleted: false},
	{IsImportOnly: false, Key: 148, EntityName: "VENEZUELA", IsDeleted: false},
	{IsImportOnly: false, Key: 149, EntityName: "AZORES", IsDeleted: false},
	{IsImportOnly: false, Key: 150, EntityName: "AUSTRALIA", IsDeleted: false},
	{IsImportOnly: false, Key: 151, EntityName: "MALYJ VYSOTSKIJ I.", IsDeleted: true},
	{IsImportOnly: false, Key: 152, EntityName: "MACAO", IsDeleted: false},
	{IsImportOnly: false, Key: 153, EntityName: "MACQUARIE I.", IsDeleted: false},
	{IsImportOnly: false, Key: 154, EntityName: "YEMEN ARAB REPUBLIC", IsDeleted: true},
	{IsImportOnly: false, Key: 155, EntityName: "MALAYA", IsDeleted: true},
	{IsImportOnly: false, Key: 157, EntityName: "NAURU", IsDeleted: false},
	{IsImportOnly: false, Key: 158, EntityName: "VANUATU", IsDeleted: false},
	{IsImportOnly: false, Key: 159, EntityName: "MALDIVES", IsDeleted: false},
	{IsImportOnly: false, Key: 160, EntityName: "TONGA", IsDeleted: false},
	{IsImportOnly: false, Key: 161, EntityName: "MALPELO I.", IsDeleted: false},
	{IsImportOnly: false, Key: 162, EntityName: "NEW CALEDONIA", IsDeleted: false},
	{IsImportOnly: false, Key: 163, EntityName: "PAPUA NEW GUINEA", IsDeleted: false},
	{IsImportOnly: false, Key: 164, EntityName: "MANCHURIA", IsDeleted: true},
	{IsImportOnly: false, Key: 165, EntityName: "MAURITIUS", IsDeleted: false},
	{IsImportOnly: false, Key: 166, EntityName: "MARIANA IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 167, EntityName: "MARKET REEF", IsDeleted: false},
	{IsImportOnly: false, Key: 168, EntityName: "MARSHALL IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 169, EntityName: "MAYOTTE", IsDeleted: false},
	{IsImportOnly: false, Key: 170, EntityName: "NEW ZEALAND", IsDeleted: false},
	{IsImportOnly: false, Key: 171, EntityName: "MELLISH REEF", IsDeleted: false},
	{IsImportOnly: false, Key: 172, EntityName: "PITCAIRN I.", IsDeleted: false},
	{IsImportOnly: false, Key: 173, EntityName: "MICRONESIA", IsDeleted: false},
	{IsImportOnly: false, Key: 174, EntityName: "MIDWAY I.", IsDeleted: false},
	{IsImportOnly: false, Key: 175, EntityName: "FRENCH POLYNESIA", IsDeleted: false},
	{IsImportOnly: false, Key: 176, EntityName: "FIJI", IsDeleted: false},
	{IsImportOnly: false, Key: 177, EntityName: "MINAMI TORISHIMA", IsDeleted: false},
	{IsImportOnly: false, Key: 178, EntityName: "MINERVA REEF", IsDeleted: true},
	{IsImportOnly: false, Key: 179, EntityName: "MOLDOVA", IsDeleted: false},
	{IsImportOnly: false, Key: 180, EntityName: "MOUNT ATHOS", IsDeleted: false},
	{IsImportOnly: false, Key: 181, EntityName: "MOZAMBIQUE", IsDeleted: false},
	{IsImportOnly: false, Key: 182, EntityName: "NAVASSA I.", IsDeleted: false},
	{IsImportOnly: false, Key: 183, EntityName: "NETHERLANDS BORNEO", IsDeleted: true},
	{IsImportOnly: false, Key: 184, EntityName: "NETHERLANDS NEW GUINEA", IsDeleted: true},
	{IsImportOnly: false, Key: 185, EntityName: "SOLOMON IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 186, EntityName: "NEWFOUNDLAND, LABRADOR", IsDeleted: true},
	{IsImportOnly: false, Key: 187, EntityName: "NIGER", IsDeleted: false},
	{IsImportOnly: false, Key: 188, EntityName: "NIUE", IsDeleted: false},
	{IsImportOnly: false, Key: 189, EntityName: "NORFOLK I.", IsDeleted: false},
	{IsImportOnly: false, Key: 190, EntityName: "SAMOA", IsDeleted: false},
	{IsImportOnly: false, Key: 191, EntityName: "NORTH COOK IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 192, EntityName: "OGASAWARA", IsDeleted: false},
	{IsImportOnly: false, Key: 193, EntityName: "OKINAWA (RYUKYU IS.)", IsDeleted: true},
	{IsImportOnly: false, Key: 194, EntityName: "OKINO TORI-SHIMA", IsDeleted: true},
	{IsImportOnly: false, Key: 195, EntityName: "ANNOBON I.", IsDeleted: false},
	{IsImportOnly: false, Key: 196, EntityName: "PALESTINE", IsDeleted: true},
	{IsImportOnly: false, Key: 197, EntityName: "PALMYRA & JARVIS IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 198, EntityName: "PAPUA TERRITORY", IsDeleted: true},
	{IsImportOnly: false, Key: 199, EntityName: "PETER 1 I.", IsDeleted: false},
	{IsImportOnly: false, Key: 200, EntityName: "PORTUGUESE TIMOR", IsDeleted: true},
	{IsImportOnly: false, Key: 201, EntityName: "PRINCE EDWARD & MARION IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 202, EntityName: "PUERTO RICO", IsDeleted: false},
	{IsImportOnly: false, Key: 203, EntityName: "ANDORRA", IsDeleted: false},
	{IsImportOnly: false, Key: 204, EntityName: "REVILLAGIGEDO", IsDeleted: false},
	{IsImportOnly: false, Key: 205, EntityName: "ASCENSION I.", IsDeleted: false},
	{IsImportOnly: false, Key: 206, EntityName: "AUSTRIA", IsDeleted: false},
	{IsImportOnly: false, Key: 207, EntityName: "RODRIGUES I.", IsDeleted: false},
	{IsImportOnly: false, Key: 208, EntityName: "RUANDA-URUNDI", IsDeleted: true},
	{IsImportOnly: false, Key: 209, EntityName: "BELGIUM", IsDeleted: false},
	{IsImportOnly: false, Key: 210, EntityName: "SAAR", IsDeleted: true},
	{IsImportOnly: false, Key: 211, EntityName: "SABLE I.", IsDeleted: false},
	{IsImportOnly: false, Key: 212, EntityName: "BULGARIA", IsDeleted: false},
	{IsImportOnly: false, Key: 213, EntityName: "SAINT MARTIN", IsDeleted: false},
	{IsImportOnly: false, Key: 214, EntityName: "CORSICA", IsDeleted: false},
	{IsImportOnly: false, Key: 215, EntityName: "CYPRUS", IsDeleted: false},
	{IsImportOnly: false, Key: 216, EntityName: "SAN ANDRES & PROVIDENCIA", IsDeleted: false},
	{IsImportOnly: false, Key: 217, EntityName: "SAN FELIX & SAN AMBROSIO", IsDeleted: false},
	{IsImportOnly: false, Key: 218, EntityName: "CZECHOSLOVAKIA", IsDeleted: true},
	{IsImportOnly: false, Key: 219, EntityName: "SAO TOME & PRINCIPE", IsDeleted: false},
	{IsImportOnly: false, Key: 220, EntityName: "SARAWAK", IsDeleted: true},
	{IsImportOnly: false, Key: 221, EntityName: "DENMARK", IsDeleted: false},
	{IsImportOnly: false, Key: 222, EntityName: "FAROE IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 223, EntityName: "ENGLAND", IsDeleted: false},
	{IsImportOnly: false, Key: 224, EntityName: "FINLAND", IsDeleted: false},
	{IsImportOnly: false, Key: 225, EntityName: "SARDINIA", IsDeleted: false},
	{IsImportOnly: false, Key: 226, EntityName: "SAUDI ARABIA/IRAQ NEUTRAL ZONE", IsDeleted: true},
	{IsImportOnly: false, Key: 227, EntityName: "FRANCE", IsDeleted: false},
	{IsImportOnly: false, Key: 228, EntityName: "SERRANA BANK & RONCADOR CAY", IsDeleted: true},
	{IsImportOnly: false, Key: 229, EntityName: "GERMAN DEMOCRATIC REPUBLIC", IsDeleted: true},
	{IsImportOnly: false, Key: 230, EntityName: "FEDERAL REPUBLIC OF GERMANY", IsDeleted: false},
	{IsImportOnly: false, Key: 231, EntityName: "SIKKIM", IsDeleted: true},
	{IsImportOnly: false, Key: 232, EntityName: "SOMALIA", IsDeleted: false},
	{IsImportOnly: false, Key: 233, EntityName: "GIBRALTAR", IsDeleted: false},
	{IsImportOnly: false, Key: 234, EntityName: "SOUTH COOK IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 235, EntityName: "SOUTH GEORGIA I.", IsDeleted: false},
	{IsImportOnly: false, Key: 236, EntityName: "GREECE", IsDeleted: false},
	{IsImportOnly: false, Key: 237, EntityName: "GREENLAND", IsDeleted: false},
	{IsImportOnly: false, Key: 238, EntityName: "SOUTH ORKNEY IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 239, EntityName: "HUNGARY", IsDeleted: false},
	{IsImportOnly: false, Key: 240, EntityName: "SOUTH SANDWICH IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 241, EntityName: "SOUTH SHETLAND IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 242, EntityName: "ICELAND", IsDeleted: false},
	{IsImportOnly: false, Key: 243, EntityName: "PEOPLE'S DEMOCRATIC REP. OF YEMEN", IsDeleted: true},
	{IsImportOnly: false, Key: 244, EntityName: "SOUTHERN SUDAN", IsDeleted: true},
	{IsImportOnly: false, Key: 245, EntityName: "IRELAND", IsDeleted: false},
	{IsImportOnly: false, Key: 246, EntityName: "SOVEREIGN MILITARY ORDER OF MALTA", IsDeleted: false},
	{IsImportOnly: false, Key: 247, EntityName: "SPRATLY IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 248, EntityName: "ITALY", IsDeleted: false},
	{IsImportOnly: false, Key: 249, EntityName: "ST. KITTS & NEVIS", IsDeleted: false},
	{IsImportOnly: false, Key: 250, EntityName: "ST. HELENA", IsDeleted: false},
	{IsImportOnly: false, Key: 251, EntityName: "LIECHTENSTEIN", IsDeleted: false},
	{IsImportOnly: false, Key: 252, EntityName: "ST. PAUL I.", IsDeleted: false},
	{IsImportOnly: false, Key: 253, EntityName: "ST. PETER & ST. PAUL ROCKS", IsDeleted: false},
	{IsImportOnly: false, Key: 254, EntityName: "LUXEMBOURG", IsDeleted: false},
	{IsImportOnly: false, Key: 255, EntityName: "ST. MAARTEN, SABA, ST. EUSTATIUS", IsDeleted: true},
	{IsImportOnly: false, Key: 256, EntityName: "MADEIRA IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 257, EntityName: "MALTA", IsDeleted: false},
	{IsImportOnly: false, Key: 258, EntityName: "SUMATRA", IsDeleted: true},
	{IsImportOnly: false, Key: 259, EntityName: "SVALBARD", IsDeleted: false},
	{IsImportOnly: false, Key: 260, EntityName: "MONACO", IsDeleted: false},
	{IsImportOnly: false, Key: 261, EntityName: "SWAN IS.", IsDeleted: true},
	{IsImportOnly: false, Key: 262, EntityName: "TAJIKISTAN", IsDeleted: false},
	{IsImportOnly: false, Key: 263, EntityName: "NETHERLANDS", IsDeleted: false},
	{IsImportOnly: false, Key: 264, EntityName: "TANGIER", IsDeleted: true},
	{IsImportOnly: false, Key: 265, EntityName: "NORTHERN IRELAND", IsDeleted: false},
	{IsImportOnly: false, Key: 266, EntityName: "NORWAY", IsDeleted: false},
	{IsImportOnly: false, Key: 267, EntityName: "TERRITORY OF NEW GUINEA", IsDeleted: true},
	{IsImportOnly: false, Key: 268, EntityName: "TIBET", IsDeleted: true},
	{IsImportOnly: false, Key: 269, EntityName: "POLAND", IsDeleted: false},
	{IsImportOnly: false, Key: 270, EntityName: "TOKELAU IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 271, EntityName: "TRIESTE", IsDeleted: true},
	{IsImportOnly: false, Key: 272, EntityName: "PORTUGAL", IsDeleted: false},
	{IsImportOnly: false, Key: 273, EntityName: "TRINDADE & MARTIM VAZ IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 274, EntityName: "TRISTAN DA CUNHA & GOUGH I.", IsDeleted: false},
	{IsImportOnly: false, Key: 275, EntityName: "ROMANIA", IsDeleted: false},
	{IsImportOnly: false, Key: 276, EntityName: "TROMELIN I.", IsDeleted: false},
	{IsImportOnly: false, Key: 277, EntityName: "ST. PIERRE & MIQUELON", IsDeleted: false},
	{IsImportOnly: false, Key: 278, EntityName: "SAN MARINO", IsDeleted: false},
	{IsImportOnly: false, Key: 279, EntityName: "SCOTLAND", IsDeleted: false},
	{IsImportOnly: false, Key: 280, EntityName: "TURKMENISTAN", IsDeleted: false},
	{IsImportOnly: false, Key: 281, EntityName: "SPAIN", IsDeleted: false},
	{IsImportOnly: false, Key: 282, EntityName: "TUVALU", IsDeleted: false},
	{IsImportOnly: false, Key: 283, EntityName: "UK SOVEREIGN BASE AREAS ON CYPRUS", IsDeleted: false},
	{IsImportOnly: false, Key: 284, EntityName: "SWEDEN", IsDeleted: false},
	{IsImportOnly: false, Key: 285, EntityName: "VIRGIN IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 286, EntityName: "UGANDA", IsDeleted: false},
	{IsImportOnly: false, Key: 287, EntityName: "SWITZERLAND", IsDeleted: false},
	{IsImportOnly: false, Key: 288, EntityName: "UKRAINE", IsDeleted: false},
	{IsImportOnly: false, Key: 289, EntityName: "UNITED NATIONS HQ", IsDeleted: false},
	{IsImportOnly: false, Key: 291, EntityName: "UNITED STATES OF AMERICA", IsDeleted: false},
	{IsImportOnly: false, Key: 292, EntityName: "UZBEKISTAN", IsDeleted: false},
	{IsImportOnly: false, Key: 293, EntityName: "VIET NAM", IsDeleted: false},
	{IsImportOnly: false, Key: 294, EntityName: "WALES", IsDeleted: false},
	{IsImportOnly: false, Key: 295, EntityName: "VATICAN", IsDeleted: false},
	{IsImportOnly: false, Key: 296, EntityName: "SERBIA", IsDeleted: false},
	{IsImportOnly: false, Key: 297, EntityName: "WAKE I.", IsDeleted: false},
	{IsImportOnly: false, Key: 298, EntityName: "WALLIS & FUTUNA IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 299, EntityName: "WEST MALAYSIA", IsDeleted: false},
	{IsImportOnly: false, Key: 301, EntityName: "W. KIRIBATI (GILBERT IS. )", IsDeleted: false},
	{IsImportOnly: false, Key: 302, EntityName: "WESTERN SAHARA", IsDeleted: false},
	{IsImportOnly: false, Key: 303, EntityName: "WILLIS I.", IsDeleted: false},
	{IsImportOnly: false, Key: 304, EntityName: "BAHRAIN", IsDeleted: false},
	{IsImportOnly: false, Key: 305, EntityName: "BANGLADESH", IsDeleted: false},
	{IsImportOnly: false, Key: 306, EntityName: "BHUTAN", IsDeleted: false},
	{IsImportOnly: false, Key: 307, EntityName: "ZANZIBAR", IsDeleted: true},
	{IsImportOnly: false, Key: 308, EntityName: "COSTA RICA", IsDeleted: false},
	{IsImportOnly: false, Key: 309, EntityName: "MYANMAR", IsDeleted: false},
	{IsImportOnly: false, Key: 312, EntityName: "CAMBODIA", IsDeleted: false},
	{IsImportOnly: false, Key: 315, EntityName: "SRI LANKA", IsDeleted: false},
	{IsImportOnly: false, Key: 318, EntityName: "CHINA", IsDeleted: false},
	{IsImportOnly: false, Key: 321, EntityName: "HONG KONG", IsDeleted: false},
	{IsImportOnly: false, Key: 324, EntityName: "INDIA", IsDeleted: false},
	{IsImportOnly: false, Key: 327, EntityName: "INDONESIA", IsDeleted: false},
	{IsImportOnly: false, Key: 330, EntityName: "IRAN", IsDeleted: false},
	{IsImportOnly: false, Key: 333, EntityName: "IRAQ", IsDeleted: false},
	{IsImportOnly: false, Key: 336, EntityName: "ISRAEL", IsDeleted: false},
	{IsImportOnly: false, Key: 339, EntityName: "JAPAN", IsDeleted: false},
	{IsImportOnly: false, Key: 342, EntityName: "JORDAN", IsDeleted: false},
	{IsImportOnly: false, Key: 344, EntityName: "DEMOCRATIC PEOPLE'S REP. OF KOREA", IsDeleted: false},
	{IsImportOnly: false, Key: 345, EntityName: "BRUNEI DARUSSALAM", IsDeleted: false},
	{IsImportOnly: false, Key: 348, EntityName: "KUWAIT", IsDeleted: false},
	{IsImportOnly: false, Key: 354, EntityName: "LEBANON", IsDeleted: false},
	{IsImportOnly: false, Key: 363, EntityName: "MONGOLIA", IsDeleted: false},
	{IsImportOnly: false, Key: 369, EntityName: "NEPAL", IsDeleted: false},
	{IsImportOnly: false, Key: 370, EntityName: "OMAN", IsDeleted: false},
	{IsImportOnly: false, Key: 372, EntityName: "PAKISTAN", IsDeleted: false},
	{IsImportOnly: false, Key: 375, EntityName: "PHILIPPINES", IsDeleted: false},
	{IsImportOnly: false, Key: 376, EntityName: "QATAR", IsDeleted: false},
	{IsImportOnly: false, Key: 378, EntityName: "SAUDI ARABIA", IsDeleted: false},
	{IsImportOnly: false, Key: 379, EntityName: "SEYCHELLES", IsDeleted: false},
	{IsImportOnly: false, Key: 381, EntityName: "SINGAPORE", IsDeleted: false},
	{IsImportOnly: false, Key: 382, EntityName: "DJIBOUTI", IsDeleted: false},
	{IsImportOnly: false, Key: 384, EntityName: "SYRIA", IsDeleted: false},
	{IsImportOnly: false, Key: 386, EntityName: "TAIWAN", IsDeleted: false},
	{IsImportOnly: false, Key: 387, EntityName: "THAILAND", IsDeleted: false},
	{IsImportOnly: false, Key: 390, EntityName: "TURKEY", IsDeleted: false},
	{IsImportOnly: false, Key: 391, EntityName: "UNITED ARAB EMIRATES", IsDeleted: false},
	{IsImportOnly: false, Key: 400, EntityName: "ALGERIA", IsDeleted: false},
	{IsImportOnly: false, Key: 401, EntityName: "ANGOLA", IsDeleted: false},
	{IsImportOnly: false, Key: 402, EntityName: "BOTSWANA", IsDeleted: false},
	{IsImportOnly: false, Key: 404, EntityName: "BURUNDI", IsDeleted: false},
	{IsImportOnly: false, Key: 406, EntityName: "CAMEROON", IsDeleted: false},
	{IsImportOnly: false, Key: 408, EntityName: "CENTRAL AFRICA", IsDeleted: false},
	{IsImportOnly: false, Key: 409, EntityName: "CAPE VERDE", IsDeleted: false},
	{IsImportOnly: false, Key: 410, EntityName: "CHAD", IsDeleted: false},
	{IsImportOnly: false, Key: 411, EntityName: "COMOROS", IsDeleted: false},
	{IsImportOnly: false, Key: 412, EntityName: "REPUBLIC OF THE CONGO", IsDeleted: false},
	{IsImportOnly: false, Key: 414, EntityName: "DEMOCRATIC REPUBLIC OF THE CONGO", IsDeleted: false},
	{IsImportOnly: false, Key: 416, EntityName: "BENIN", IsDeleted: false},
	{IsImportOnly: false, Key: 420, EntityName: "GABON", IsDeleted: false},
	{IsImportOnly: false, Key: 422, EntityName: "THE GAMBIA", IsDeleted: false},
	{IsImportOnly: false, Key: 424, EntityName: "GHANA", IsDeleted: false},
	{IsImportOnly: false, Key: 428, EntityName: "COTE D'IVOIRE", IsDeleted: false},
	{IsImportOnly: false, Key: 430, EntityName: "KENYA", IsDeleted: false},
	{IsImportOnly: false, Key: 432, EntityName: "LESOTHO", IsDeleted: false},
	{IsImportOnly: false, Key: 434, EntityName: "LIBERIA", IsDeleted: false},
	{IsImportOnly: false, Key: 436, EntityName: "LIBYA", IsDeleted: false},
	{IsImportOnly: false, Key: 438, EntityName: "MADAGASCAR", IsDeleted: false},
	{IsImportOnly: false, Key: 440, EntityName: "MALAWI", IsDeleted: false},
	{IsImportOnly: false, Key: 442, EntityName: "MALI", IsDeleted: false},
	{IsImportOnly: false, Key: 444, EntityName: "MAURITANIA", IsDeleted: false},
	{IsImportOnly: false, Key: 446, EntityName: "MOROCCO", IsDeleted: false},
	{IsImportOnly: false, Key: 450, EntityName: "NIGERIA", IsDeleted: false},
	{IsImportOnly: false, Key: 452, EntityName: "ZIMBABWE", IsDeleted: false},
	{IsImportOnly: false, Key: 453, EntityName: "REUNION I.", IsDeleted: false},
	{IsImportOnly: false, Key: 454, EntityName: "RWANDA", IsDeleted: false},
	{IsImportOnly: false, Key: 456, EntityName: "SENEGAL", IsDeleted: false},
	{IsImportOnly: false, Key: 458, EntityName: "SIERRA LEONE", IsDeleted: false},
	{IsImportOnly: false, Key: 460, EntityName: "ROTUMA I.", IsDeleted: false},
	{IsImportOnly: false, Key: 462, EntityName: "REPUBLIC OF SOUTH AFRICA", IsDeleted: false},
	{IsImportOnly: false, Key: 464, EntityName: "NAMIBIA", IsDeleted: false},
	{IsImportOnly: false, Key: 466, EntityName: "SUDAN", IsDeleted: false},
	{IsImportOnly: false, Key: 468, EntityName: "KINGDOM OF ESWATINI", IsDeleted: false},
	{IsImportOnly: false, Key: 470, EntityName: "TANZANIA", IsDeleted: false},
	{IsImportOnly: false, Key: 474, EntityName: "TUNISIA", IsDeleted: false},
	{IsImportOnly: false, Key: 478, EntityName: "EGYPT", IsDeleted: false},
	{IsImportOnly: false, Key: 480, EntityName: "BURKINA FASO", IsDeleted: false},
	{IsImportOnly: false, Key: 482, EntityName: "ZAMBIA", IsDeleted: false},
	{IsImportOnly: false, Key: 483, EntityName: "TOGO", IsDeleted: false},
	{IsImportOnly: false, Key: 488, EntityName: "WALVIS BAY", IsDeleted: true},
	{IsImportOnly: false, Key: 489, EntityName: "CONWAY REEF", IsDeleted: false},
	{IsImportOnly: false, Key: 490, EntityName: "BANABA I. (OCEAN I.)", IsDeleted: false},
	{IsImportOnly: false, Key: 492, EntityName: "YEMEN", IsDeleted: false},
	{IsImportOnly: false, Key: 493, EntityName: "PENGUIN IS.", IsDeleted: true},
	{IsImportOnly: false, Key: 497, EntityName: "CROATIA", IsDeleted: false},
	{IsImportOnly: false, Key: 499, EntityName: "SLOVENIA", IsDeleted: false},
	{IsImportOnly: false, Key: 501, EntityName: "BOSNIA-HERZEGOVINA", IsDeleted: false},
	{IsImportOnly: false, Key: 502, EntityName: "NORTH MACEDONIA (REPUBLIC OF)", IsDeleted: false},
	{IsImportOnly: false, Key: 503, EntityName: "CZECH REPUBLIC", IsDeleted: false},
	{IsImportOnly: false, Key: 504, EntityName: "SLOVAK REPUBLIC", IsDeleted: false},
	{IsImportOnly: false, Key: 505, EntityName: "PRATAS I.", IsDeleted: false},
	{IsImportOnly: false, Key: 506, EntityName: "SCARBOROUGH REEF", IsDeleted: false},
	{IsImportOnly: false, Key: 507, EntityName: "TEMOTU PROVINCE", IsDeleted: false},
	{IsImportOnly: false, Key: 508, EntityName: "AUSTRAL I.", IsDeleted: false},
	{IsImportOnly: false, Key: 509, EntityName: "MARQUESAS IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 510, EntityName: "PALESTINE", IsDeleted: false},
	{IsImportOnly: false, Key: 511, EntityName: "TIMOR-LESTE", IsDeleted: false},
	{IsImportOnly: false, Key: 512, EntityName: "CHESTERFIELD IS.", IsDeleted: false},
	{IsImportOnly: false, Key: 513, EntityName: "DUCIE I.", IsDeleted: false},
	{IsImportOnly: false, Key: 514, EntityName: "MONTENEGRO", IsDeleted: false},
	{IsImportOnly: false, Key: 515, EntityName: "SWAINS I.", IsDeleted: false},
	{IsImportOnly: false, Key: 516, EntityName: "SAINT BARTHELEMY", IsDeleted: false},
	{IsImportOnly: false, Key: 517, EntityName: "CURACAO", IsDeleted: false},
	{IsImportOnly: false, Key: 518, EntityName: "SINT MAARTEN", IsDeleted: false},
	{IsImportOnly: false, Key: 519, EntityName: "SABA & ST. EUSTATIUS", IsDeleted: false},
	{IsImportOnly: false, Key: 520, EntityName: "BONAIRE", IsDeleted: false},
	{IsImportOnly: false, Key: 521, EntityName: "SOUTH SUDAN (REPUBLIC OF)", IsDeleted: false},
	{IsImportOnly: false, Key: 522, EntityName: "REPUBLIC OF KOSOVO", IsDeleted: false},
}

// lookupMap contains all known DXCCEntityCode specifications.
var lookupMap = map[DXCCEntityCode]*Spec{
	NONE:                                     &lookupList[0],
	CANADA:                                   &lookupList[1],
	ABU_AIL_IS_DELETED:                       &lookupList[2],
	AFGHANISTAN:                              &lookupList[3],
	AGALEGA_AND_ST_BRANDON_IS:                &lookupList[4],
	ALAND_IS:                                 &lookupList[5],
	ALASKA:                                   &lookupList[6],
	ALBANIA:                                  &lookupList[7],
	ALDABRA_DELETED:                          &lookupList[8],
	AMERICAN_SAMOA:                           &lookupList[9],
	AMSTERDAM_AND_ST_PAUL_IS:                 &lookupList[10],
	ANDAMAN_AND_NICOBAR_IS:                   &lookupList[11],
	ANGUILLA:                                 &lookupList[12],
	ANTARCTICA:                               &lookupList[13],
	ARMENIA:                                  &lookupList[14],
	ASIATIC_RUSSIA:                           &lookupList[15],
	NEW_ZEALAND_SUBANTARCTIC_ISLANDS:         &lookupList[16],
	AVES_I:                                   &lookupList[17],
	AZERBAIJAN:                               &lookupList[18],
	BAJO_NUEVO_DELETED:                       &lookupList[19],
	BAKER_AND_HOWLAND_IS:                     &lookupList[20],
	BALEARIC_IS:                              &lookupList[21],
	PALAU:                                    &lookupList[22],
	BLENHEIM_REEF_DELETED:                    &lookupList[23],
	BOUVET:                                   &lookupList[24],
	BRITISH_NORTH_BORNEO_DELETED:             &lookupList[25],
	BRITISH_SOMALILAND_DELETED:               &lookupList[26],
	BELARUS:                                  &lookupList[27],
	CANAL_ZONE_DELETED:                       &lookupList[28],
	CANARY_IS:                                &lookupList[29],
	CELEBE_AND_MOLUCCA_IS_DELETED:            &lookupList[30],
	C_KIRIBATI_BRITISH_PHOENIX_IS:            &lookupList[31],
	CEUTA_AND_MELILLA:                        &lookupList[32],
	CHAGOS_IS:                                &lookupList[33],
	CHATHAM_IS:                               &lookupList[34],
	CHRISTMAS_I:                              &lookupList[35],
	CLIPPERTON_I:                             &lookupList[36],
	COCOS_I:                                  &lookupList[37],
	COCOS_KEELING_IS:                         &lookupList[38],
	COMOROS_DELETED:                          &lookupList[39],
	CRETE:                                    &lookupList[40],
	CROZET_I:                                 &lookupList[41],
	DAMAO_DIU_DELETED:                        &lookupList[42],
	DESECHEO_I:                               &lookupList[43],
	DESROCHES_DELETED:                        &lookupList[44],
	DODECANESE:                               &lookupList[45],
	EAST_MALAYSIA:                            &lookupList[46],
	EASTER_I:                                 &lookupList[47],
	E_KIRIBATI_LINE_IS:                       &lookupList[48],
	EQUATORIAL_GUINEA:                        &lookupList[49],
	MEXICO:                                   &lookupList[50],
	ERITREA:                                  &lookupList[51],
	ESTONIA:                                  &lookupList[52],
	ETHIOPIA:                                 &lookupList[53],
	EUROPEAN_RUSSIA:                          &lookupList[54],
	FARQUHAR_DELETED:                         &lookupList[55],
	FERNANDO_DE_NORONHA:                      &lookupList[56],
	FRENCH_EQUATORIAL_AFRICA_DELETED:         &lookupList[57],
	FRENCH_INDO_CHINA_DELETED:                &lookupList[58],
	FRENCH_WEST_AFRICA_DELETED:               &lookupList[59],
	BAHAMAS:                                  &lookupList[60],
	FRANZ_JOSEF_LAND:                         &lookupList[61],
	BARBADOS:                                 &lookupList[62],
	FRENCH_GUIANA:                            &lookupList[63],
	BERMUDA:                                  &lookupList[64],
	BRITISH_VIRGIN_IS:                        &lookupList[65],
	BELIZE:                                   &lookupList[66],
	FRENCH_INDIA_DELETED:                     &lookupList[67],
	KUWAIT_SAUDI_ARABIA_NEUTRAL_ZONE_DELETED: &lookupList[68],
	CAYMAN_IS:                                &lookupList[69],
	CUBA:                                     &lookupList[70],
	GALAPAGOS_IS:                             &lookupList[71],
	DOMINICAN_REPUBLIC:                       &lookupList[72],
	EL_SALVADOR:                              &lookupList[73],
	GEORGIA:                                  &lookupList[74],
	GUATEMALA:                                &lookupList[75],
	GRENADA:                                  &lookupList[76],
	HAITI:                                    &lookupList[77],
	GUADELOUPE:                               &lookupList[78],
	HONDURAS:                                 &lookupList[79],
	GERMANY_DELETED:                          &lookupList[80],
	JAMAICA:                                  &lookupList[81],
	MARTINIQUE:                               &lookupList[82],
	BONAIRE_CURACAO_DELETED:                  &lookupList[83],
	NICARAGUA:                                &lookupList[84],
	PANAMA:                                   &lookupList[85],
	TURKS_AND_CAICOS_IS:                      &lookupList[86],
	TRINIDAD_AND_TOBAGO:                      &lookupList[87],
	ARUBA:                                    &lookupList[88],
	GEYSER_REEF_DELETED:                      &lookupList[89],
	ANTIGUA_AND_BARBUDA:                      &lookupList[90],
	DOMINICA:                                 &lookupList[91],
	MONTSERRAT:                               &lookupList[92],
	ST_LUCIA:                                 &lookupList[93],
	ST_VINCENT:                               &lookupList[94],
	GLORIOSO_IS:                              &lookupList[95],
	ARGENTINA:                                &lookupList[96],
	GOA_DELETED:                              &lookupList[97],
	GOLD_COAST_TOGOLAND_DELETED:              &lookupList[98],
	GUAM:                                     &lookupList[99],
	BOLIVIA:                                  &lookupList[100],
	GUANTANAMO_BAY:                           &lookupList[101],
	GUERNSEY:                                 &lookupList[102],
	GUINEA:                                   &lookupList[103],
	BRAZIL:                                   &lookupList[104],
	GUINEA_BISSAU:                            &lookupList[105],
	HAWAII:                                   &lookupList[106],
	HEARD_I:                                  &lookupList[107],
	CHILE:                                    &lookupList[108],
	IFNI_DELETED:                             &lookupList[109],
	ISLE_OF_MAN:                              &lookupList[110],
	ITALIAN_SOMALILAND_DELETED:               &lookupList[111],
	COLOMBIA:                                 &lookupList[112],
	ITU_HQ:                                   &lookupList[113],
	JAN_MAYEN:                                &lookupList[114],
	JAVA_DELETED:                             &lookupList[115],
	ECUADOR:                                  &lookupList[116],
	JERSEY:                                   &lookupList[117],
	JOHNSTON_I:                               &lookupList[118],
	JUAN_DE_NOVA_EUROPA:                      &lookupList[119],
	JUAN_FERNANDEZ_IS:                        &lookupList[120],
	KALININGRAD:                              &lookupList[121],
	KAMARAN_IS_DELETED:                       &lookupList[122],
	KARELO_FINNISH_REPUBLIC_DELETED:          &lookupList[123],
	GUYANA:                                   &lookupList[124],
	KAZAKHSTAN:                               &lookupList[125],
	KERGUELEN_IS:                             &lookupList[126],
	PARAGUAY:                                 &lookupList[127],
	KERMADEC_IS:                              &lookupList[128],
	KINGMAN_REEF_DELETED:                     &lookupList[129],
	KYRGYZSTAN:                               &lookupList[130],
	PERU:                                     &lookupList[131],
	REPUBLIC_OF_KOREA:                        &lookupList[132],
	KURE_I:                                   &lookupList[133],
	KURIA_MURIA_I_DELETED:                    &lookupList[134],
	SURINAME:                                 &lookupList[135],
	FALKLAND_IS:                              &lookupList[136],
	LAKSHADWEEP_IS:                           &lookupList[137],
	LAOS:                                     &lookupList[138],
	URUGUAY:                                  &lookupList[139],
	LATVIA:                                   &lookupList[140],
	LITHUANIA:                                &lookupList[141],
	LORD_HOWE_I:                              &lookupList[142],
	VENEZUELA:                                &lookupList[143],
	AZORES:                                   &lookupList[144],
	AUSTRALIA:                                &lookupList[145],
	MALYJ_VYSOTSKIJ_I_DELETED:                &lookupList[146],
	MACAO:                                    &lookupList[147],
	MACQUARIE_I:                              &lookupList[148],
	YEMEN_ARAB_REPUBLIC_DELETED:              &lookupList[149],
	MALAYA_DELETED:                           &lookupList[150],
	NAURU:                                    &lookupList[151],
	VANUATU:                                  &lookupList[152],
	MALDIVES:                                 &lookupList[153],
	TONGA:                                    &lookupList[154],
	MALPELO_I:                                &lookupList[155],
	NEW_CALEDONIA:                            &lookupList[156],
	PAPUA_NEW_GUINEA:                         &lookupList[157],
	MANCHURIA_DELETED:                        &lookupList[158],
	MAURITIUS:                                &lookupList[159],
	MARIANA_IS:                               &lookupList[160],
	MARKET_REEF:                              &lookupList[161],
	MARSHALL_IS:                              &lookupList[162],
	MAYOTTE:                                  &lookupList[163],
	NEW_ZEALAND:                              &lookupList[164],
	MELLISH_REEF:                             &lookupList[165],
	PITCAIRN_I:                               &lookupList[166],
	MICRONESIA:                               &lookupList[167],
	MIDWAY_I:                                 &lookupList[168],
	FRENCH_POLYNESIA:                         &lookupList[169],
	FIJI:                                     &lookupList[170],
	MINAMI_TORISHIMA:                         &lookupList[171],
	MINERVA_REEF_DELETED:                     &lookupList[172],
	MOLDOVA:                                  &lookupList[173],
	MOUNT_ATHOS:                              &lookupList[174],
	MOZAMBIQUE:                               &lookupList[175],
	NAVASSA_I:                                &lookupList[176],
	NETHERLANDS_BORNEO_DELETED:               &lookupList[177],
	NETHERLANDS_NEW_GUINEA_DELETED:           &lookupList[178],
	SOLOMON_IS:                               &lookupList[179],
	NEWFOUNDLAND_LABRADOR_DELETED:            &lookupList[180],
	NIGER:                                    &lookupList[181],
	NIUE:                                     &lookupList[182],
	NORFOLK_I:                                &lookupList[183],
	SAMOA:                                    &lookupList[184],
	NORTH_COOK_IS:                            &lookupList[185],
	OGASAWARA:                                &lookupList[186],
	OKINAWA_RYUKYU_IS_DELETED:                &lookupList[187],
	OKINO_TORI_SHIMA_DELETED:                 &lookupList[188],
	ANNOBON_I:                                &lookupList[189],
	PALESTINE_DELETED:                        &lookupList[190],
	PALMYRA_AND_JARVIS_IS:                    &lookupList[191],
	PAPUA_TERRITORY_DELETED:                  &lookupList[192],
	PETER_1_I:                                &lookupList[193],
	PORTUGUESE_TIMOR_DELETED:                 &lookupList[194],
	PRINCE_EDWARD_AND_MARION_IS:              &lookupList[195],
	PUERTO_RICO:                              &lookupList[196],
	ANDORRA:                                  &lookupList[197],
	REVILLAGIGEDO:                            &lookupList[198],
	ASCENSION_I:                              &lookupList[199],
	AUSTRIA:                                  &lookupList[200],
	RODRIGUES_I:                              &lookupList[201],
	RUANDA_URUNDI_DELETED:                    &lookupList[202],
	BELGIUM:                                  &lookupList[203],
	SAAR_DELETED:                             &lookupList[204],
	SABLE_I:                                  &lookupList[205],
	BULGARIA:                                 &lookupList[206],
	SAINT_MARTIN:                             &lookupList[207],
	CORSICA:                                  &lookupList[208],
	CYPRUS:                                   &lookupList[209],
	SAN_ANDRES_AND_PROVIDENCIA:               &lookupList[210],
	SAN_FELIX_AND_SAN_AMBROSIO:               &lookupList[211],
	CZECHOSLOVAKIA_DELETED:                   &lookupList[212],
	SAO_TOME_AND_PRINCIPE:                    &lookupList[213],
	SARAWAK_DELETED:                          &lookupList[214],
	DENMARK:                                  &lookupList[215],
	FAROE_IS:                                 &lookupList[216],
	ENGLAND:                                  &lookupList[217],
	FINLAND:                                  &lookupList[218],
	SARDINIA:                                 &lookupList[219],
	SAUDI_ARABIA_IRAQ_NEUTRAL_ZONE_DELETED:   &lookupList[220],
	FRANCE:                                   &lookupList[221],
	SERRANA_BANK_AND_RONCADOR_CAY_DELETED:    &lookupList[222],
	GERMAN_DEMOCRATIC_REPUBLIC_DELETED:       &lookupList[223],
	FEDERAL_REPUBLIC_OF_GERMANY:              &lookupList[224],
	SIKKIM_DELETED:                           &lookupList[225],
	SOMALIA:                                  &lookupList[226],
	GIBRALTAR:                                &lookupList[227],
	SOUTH_COOK_IS:                            &lookupList[228],
	SOUTH_GEORGIA_I:                          &lookupList[229],
	GREECE:                                   &lookupList[230],
	GREENLAND:                                &lookupList[231],
	SOUTH_ORKNEY_IS:                          &lookupList[232],
	HUNGARY:                                  &lookupList[233],
	SOUTH_SANDWICH_IS:                        &lookupList[234],
	SOUTH_SHETLAND_IS:                        &lookupList[235],
	ICELAND:                                  &lookupList[236],
	PEOPLES_DEMOCRATIC_REP_OF_YEMEN_DELETED:  &lookupList[237],
	SOUTHERN_SUDAN_DELETED:                   &lookupList[238],
	IRELAND:                                  &lookupList[239],
	SOVEREIGN_MILITARY_ORDER_OF_MALTA:        &lookupList[240],
	SPRATLY_IS:                               &lookupList[241],
	ITALY:                                    &lookupList[242],
	ST_KITTS_AND_NEVIS:                       &lookupList[243],
	ST_HELENA:                                &lookupList[244],
	LIECHTENSTEIN:                            &lookupList[245],
	ST_PAUL_I:                                &lookupList[246],
	ST_PETER_AND_ST_PAUL_ROCKS:               &lookupList[247],
	LUXEMBOURG:                               &lookupList[248],
	ST_MAARTEN_SABA_ST_EUSTATIUS_DELETED:     &lookupList[249],
	MADEIRA_IS:                               &lookupList[250],
	MALTA:                                    &lookupList[251],
	SUMATRA_DELETED:                          &lookupList[252],
	SVALBARD:                                 &lookupList[253],
	MONACO:                                   &lookupList[254],
	SWAN_IS_DELETED:                          &lookupList[255],
	TAJIKISTAN:                               &lookupList[256],
	NETHERLANDS:                              &lookupList[257],
	TANGIER_DELETED:                          &lookupList[258],
	NORTHERN_IRELAND:                         &lookupList[259],
	NORWAY:                                   &lookupList[260],
	TERRITORY_OF_NEW_GUINEA_DELETED:          &lookupList[261],
	TIBET_DELETED:                            &lookupList[262],
	POLAND:                                   &lookupList[263],
	TOKELAU_IS:                               &lookupList[264],
	TRIESTE_DELETED:                          &lookupList[265],
	PORTUGAL:                                 &lookupList[266],
	TRINDADE_AND_MARTIM_VAZ_IS:               &lookupList[267],
	TRISTAN_DA_CUNHA_AND_GOUGH_I:             &lookupList[268],
	ROMANIA:                                  &lookupList[269],
	TROMELIN_I:                               &lookupList[270],
	ST_PIERRE_AND_MIQUELON:                   &lookupList[271],
	SAN_MARINO:                               &lookupList[272],
	SCOTLAND:                                 &lookupList[273],
	TURKMENISTAN:                             &lookupList[274],
	SPAIN:                                    &lookupList[275],
	TUVALU:                                   &lookupList[276],
	UK_SOVEREIGN_BASE_AREAS_ON_CYPRUS:        &lookupList[277],
	SWEDEN:                                   &lookupList[278],
	VIRGIN_IS:                                &lookupList[279],
	UGANDA:                                   &lookupList[280],
	SWITZERLAND:                              &lookupList[281],
	UKRAINE:                                  &lookupList[282],
	UNITED_NATIONS_HQ:                        &lookupList[283],
	UNITED_STATES_OF_AMERICA:                 &lookupList[284],
	UZBEKISTAN:                               &lookupList[285],
	VIET_NAM:                                 &lookupList[286],
	WALES:                                    &lookupList[287],
	VATICAN:                                  &lookupList[288],
	SERBIA:                                   &lookupList[289],
	WAKE_I:                                   &lookupList[290],
	WALLIS_AND_FUTUNA_IS:                     &lookupList[291],
	WEST_MALAYSIA:                            &lookupList[292],
	W_KIRIBATI_GILBERT_IS:                    &lookupList[293],
	WESTERN_SAHARA:                           &lookupList[294],
	WILLIS_I:                                 &lookupList[295],
	BAHRAIN:                                  &lookupList[296],
	BANGLADESH:                               &lookupList[297],
	BHUTAN:                                   &lookupList[298],
	ZANZIBAR_DELETED:                         &lookupList[299],
	COSTA_RICA:                               &lookupList[300],
	MYANMAR:                                  &lookupList[301],
	CAMBODIA:                                 &lookupList[302],
	SRI_LANKA:                                &lookupList[303],
	CHINA:                                    &lookupList[304],
	HONG_KONG:                                &lookupList[305],
	INDIA:                                    &lookupList[306],
	INDONESIA:                                &lookupList[307],
	IRAN:                                     &lookupList[308],
	IRAQ:                                     &lookupList[309],
	ISRAEL:                                   &lookupList[310],
	JAPAN:                                    &lookupList[311],
	JORDAN:                                   &lookupList[312],
	DEMOCRATIC_PEOPLES_REP_OF_KOREA:          &lookupList[313],
	BRUNEI_DARUSSALAM:                        &lookupList[314],
	KUWAIT:                                   &lookupList[315],
	LEBANON:                                  &lookupList[316],
	MONGOLIA:                                 &lookupList[317],
	NEPAL:                                    &lookupList[318],
	OMAN:                                     &lookupList[319],
	PAKISTAN:                                 &lookupList[320],
	PHILIPPINES:                              &lookupList[321],
	QATAR:                                    &lookupList[322],
	SAUDI_ARABIA:                             &lookupList[323],
	SEYCHELLES:                               &lookupList[324],
	SINGAPORE:                                &lookupList[325],
	DJIBOUTI:                                 &lookupList[326],
	SYRIA:                                    &lookupList[327],
	TAIWAN:                                   &lookupList[328],
	THAILAND:                                 &lookupList[329],
	TURKEY:                                   &lookupList[330],
	UNITED_ARAB_EMIRATES:                     &lookupList[331],
	ALGERIA:                                  &lookupList[332],
	ANGOLA:                                   &lookupList[333],
	BOTSWANA:                                 &lookupList[334],
	BURUNDI:                                  &lookupList[335],
	CAMEROON:                                 &lookupList[336],
	CENTRAL_AFRICA:                           &lookupList[337],
	CAPE_VERDE:                               &lookupList[338],
	CHAD:                                     &lookupList[339],
	COMOROS:                                  &lookupList[340],
	REPUBLIC_OF_THE_CONGO:                    &lookupList[341],
	DEMOCRATIC_REPUBLIC_OF_THE_CONGO:         &lookupList[342],
	BENIN:                                    &lookupList[343],
	GABON:                                    &lookupList[344],
	THE_GAMBIA:                               &lookupList[345],
	GHANA:                                    &lookupList[346],
	COTE_DIVOIRE:                             &lookupList[347],
	KENYA:                                    &lookupList[348],
	LESOTHO:                                  &lookupList[349],
	LIBERIA:                                  &lookupList[350],
	LIBYA:                                    &lookupList[351],
	MADAGASCAR:                               &lookupList[352],
	MALAWI:                                   &lookupList[353],
	MALI:                                     &lookupList[354],
	MAURITANIA:                               &lookupList[355],
	MOROCCO:                                  &lookupList[356],
	NIGERIA:                                  &lookupList[357],
	ZIMBABWE:                                 &lookupList[358],
	REUNION_I:                                &lookupList[359],
	RWANDA:                                   &lookupList[360],
	SENEGAL:                                  &lookupList[361],
	SIERRA_LEONE:                             &lookupList[362],
	ROTUMA_I:                                 &lookupList[363],
	REPUBLIC_OF_SOUTH_AFRICA:                 &lookupList[364],
	NAMIBIA:                                  &lookupList[365],
	SUDAN:                                    &lookupList[366],
	KINGDOM_OF_ESWATINI:                      &lookupList[367],
	TANZANIA:                                 &lookupList[368],
	TUNISIA:                                  &lookupList[369],
	EGYPT:                                    &lookupList[370],
	BURKINA_FASO:                             &lookupList[371],
	ZAMBIA:                                   &lookupList[372],
	TOGO:                                     &lookupList[373],
	WALVIS_BAY_DELETED:                       &lookupList[374],
	CONWAY_REEF:                              &lookupList[375],
	BANABA_I_OCEAN_I:                         &lookupList[376],
	YEMEN:                                    &lookupList[377],
	PENGUIN_IS_DELETED:                       &lookupList[378],
	CROATIA:                                  &lookupList[379],
	SLOVENIA:                                 &lookupList[380],
	BOSNIA_HERZEGOVINA:                       &lookupList[381],
	NORTH_MACEDONIA_REPUBLIC_OF:              &lookupList[382],
	CZECH_REPUBLIC:                           &lookupList[383],
	SLOVAK_REPUBLIC:                          &lookupList[384],
	PRATAS_I:                                 &lookupList[385],
	SCARBOROUGH_REEF:                         &lookupList[386],
	TEMOTU_PROVINCE:                          &lookupList[387],
	AUSTRAL_I:                                &lookupList[388],
	MARQUESAS_IS:                             &lookupList[389],
	PALESTINE:                                &lookupList[390],
	TIMOR_LESTE:                              &lookupList[391],
	CHESTERFIELD_IS:                          &lookupList[392],
	DUCIE_I:                                  &lookupList[393],
	MONTENEGRO:                               &lookupList[394],
	SWAINS_I:                                 &lookupList[395],
	SAINT_BARTHELEMY:                         &lookupList[396],
	CURACAO:                                  &lookupList[397],
	SINT_MAARTEN:                             &lookupList[398],
	SABA_AND_ST_EUSTATIUS:                    &lookupList[399],
	BONAIRE:                                  &lookupList[400],
	SOUTH_SUDAN_REPUBLIC_OF:                  &lookupList[401],
	REPUBLIC_OF_KOSOVO:                       &lookupList[402],
}

// Lookup returns the specification for the provided DXCCEntityCode.
// ADIF 3.1.6
func Lookup(d DXCCEntityCode) (Spec, bool) {
	spec, ok := lookupMap[d]
	if !ok {
		return Spec{}, false
	}
	return *spec, true
}

// LookupByFilter returns all DXCCEntityCode specifications that match the provided filter function.
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

// List returns all DXCCEntityCode specifications.
// This list includes those marked import-only.
// ADIF 3.1.6
func List() []Spec {
	list := make([]Spec, len(lookupList))
	copy(list, lookupList)
	return list
}

// ListActive returns DXCCEntityCode specifications.
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
