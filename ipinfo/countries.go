package ipinfo

type CountryDetails struct {
	Name string `json:"name"`
	IsEU bool   `json:"isEU"`
}

// GetCountryName gets the full name of a country from its code, e.g.
// "PK" -> "Pakistan".
func GetCountryName(country string) string {
	return countriesMap[country].Name
}

// IsEU takes the country code and returns `true`
// if the country is a member of the EU, e.g. "SE" -> true
func IsEU(country string) bool {
	return countriesMap[country].IsEU
}

var countriesMap = map[string]CountryDetails{
	"BD": {"Bangladesh", false},
	"BE": {"Belgium", true},
	"BF": {"Burkina Faso", false},
	"BG": {"Bulgaria", true},
	"BA": {"Bosnia and Herzegovina", false},
	"BB": {"Barbados", false},
	"WF": {"Wallis and Futuna", false},
	"BL": {"Saint Barthelemy", false},
	"BM": {"Bermuda", false},
	"BN": {"Brunei", false},
	"BO": {"Bolivia", false},
	"BH": {"Bahrain", false},
	"BI": {"Burundi", false},
	"BJ": {"Benin", false},
	"BT": {"Bhutan", false},
	"JM": {"Jamaica", false},
	"BV": {"Bouvet Island", false},
	"BW": {"Botswana", false},
	"WS": {"Samoa", false},
	"BQ": {"Bonaire, Saint Eustatius and Saba ", false},
	"BR": {"Brazil", false},
	"BS": {"Bahamas", false},
	"JE": {"Jersey", false},
	"BY": {"Belarus", false},
	"BZ": {"Belize", false},
	"RU": {"Russia", false},
	"RW": {"Rwanda", false},
	"RS": {"Serbia", false},
	"TL": {"East Timor", false},
	"RE": {"Reunion", false},
	"TM": {"Turkmenistan", false},
	"TJ": {"Tajikistan", false},
	"RO": {"Romania", true},
	"TK": {"Tokelau", false},
	"GW": {"Guinea-Bissau", false},
	"GU": {"Guam", false},
	"GT": {"Guatemala", false},
	"GS": {"South Georgia and the South Sandwich Islands", false},
	"GR": {"Greece", true},
	"GQ": {"Equatorial Guinea", false},
	"GP": {"Guadeloupe", false},
	"JP": {"Japan", false},
	"GY": {"Guyana", false},
	"GG": {"Guernsey", false},
	"GF": {"French Guiana", false},
	"GE": {"Georgia", false},
	"GD": {"Grenada", false},
	"GB": {"United Kingdom", false},
	"GA": {"Gabon", false},
	"SV": {"El Salvador", false},
	"GN": {"Guinea", false},
	"GM": {"Gambia", false},
	"GL": {"Greenland", false},
	"GI": {"Gibraltar", false},
	"GH": {"Ghana", false},
	"OM": {"Oman", false},
	"TN": {"Tunisia", false},
	"JO": {"Jordan", false},
	"HR": {"Croatia", true},
	"HT": {"Haiti", false},
	"HU": {"Hungary", true},
	"HK": {"Hong Kong", false},
	"HN": {"Honduras", false},
	"HM": {"Heard Island and McDonald Islands", false},
	"VE": {"Venezuela", false},
	"PR": {"Puerto Rico", false},
	"PS": {"Palestinian Territory", false},
	"PW": {"Palau", false},
	"PT": {"Portugal", true},
	"SJ": {"Svalbard and Jan Mayen", false},
	"PY": {"Paraguay", false},
	"IQ": {"Iraq", false},
	"PA": {"Panama", false},
	"PF": {"French Polynesia", false},
	"PG": {"Papua New Guinea", false},
	"PE": {"Peru", false},
	"PK": {"Pakistan", false},
	"PH": {"Philippines", false},
	"PN": {"Pitcairn", false},
	"PL": {"Poland", true},
	"PM": {"Saint Pierre and Miquelon", false},
	"ZM": {"Zambia", false},
	"EH": {"Western Sahara", false},
	"EE": {"Estonia", true},
	"EG": {"Egypt", false},
	"ZA": {"South Africa", false},
	"EC": {"Ecuador", false},
	"IT": {"Italy", true},
	"VN": {"Vietnam", false},
	"SB": {"Solomon Islands", false},
	"ET": {"Ethiopia", false},
	"SO": {"Somalia", false},
	"ZW": {"Zimbabwe", false},
	"SA": {"Saudi Arabia", false},
	"ES": {"Spain", true},
	"ER": {"Eritrea", false},
	"ME": {"Montenegro", false},
	"MD": {"Moldova", false},
	"MG": {"Madagascar", false},
	"MF": {"Saint Martin", false},
	"MA": {"Morocco", false},
	"MC": {"Monaco", false},
	"UZ": {"Uzbekistan", false},
	"MM": {"Myanmar", false},
	"ML": {"Mali", false},
	"MO": {"Macao", false},
	"MN": {"Mongolia", false},
	"MH": {"Marshall Islands", false},
	"MK": {"Macedonia", false},
	"MU": {"Mauritius", false},
	"MT": {"Malta", true},
	"MW": {"Malawi", false},
	"MV": {"Maldives", false},
	"MQ": {"Martinique", false},
	"MP": {"Northern Mariana Islands", false},
	"MS": {"Montserrat", false},
	"MR": {"Mauritania", false},
	"IM": {"Isle of Man", false},
	"UG": {"Uganda", false},
	"TZ": {"Tanzania", false},
	"MY": {"Malaysia", false},
	"MX": {"Mexico", false},
	"IL": {"Israel", false},
	"FR": {"France", true},
	"IO": {"British Indian Ocean Territory", false},
	"SH": {"Saint Helena", false},
	"FI": {"Finland", true},
	"FJ": {"Fiji", false},
	"FK": {"Falkland Islands", false},
	"FM": {"Micronesia", false},
	"FO": {"Faroe Islands", false},
	"NI": {"Nicaragua", false},
	"NL": {"Netherlands", true},
	"NO": {"Norway", false},
	"NA": {"Namibia", false},
	"VU": {"Vanuatu", false},
	"NC": {"New Caledonia", false},
	"NE": {"Niger", false},
	"NF": {"Norfolk Island", false},
	"NG": {"Nigeria", false},
	"NZ": {"New Zealand", false},
	"NP": {"Nepal", false},
	"NR": {"Nauru", false},
	"NU": {"Niue", false},
	"CK": {"Cook Islands", false},
	"XK": {"Kosovo", false},
	"CI": {"Ivory Coast", false},
	"CH": {"Switzerland", false},
	"CO": {"Colombia", false},
	"CN": {"China", false},
	"CM": {"Cameroon", false},
	"CL": {"Chile", false},
	"CC": {"Cocos Islands", false},
	"CA": {"Canada", false},
	"CG": {"Republic of the Congo", false},
	"CF": {"Central African Republic", false},
	"CD": {"Democratic Republic of the Congo", false},
	"CZ": {"Czech Republic", true},
	"CY": {"Cyprus", true},
	"CX": {"Christmas Island", false},
	"CR": {"Costa Rica", false},
	"CW": {"Curacao", false},
	"CV": {"Cape Verde", false},
	"CU": {"Cuba", false},
	"SZ": {"Swaziland", false},
	"SY": {"Syria", false},
	"SX": {"Sint Maarten", false},
	"KG": {"Kyrgyzstan", false},
	"KE": {"Kenya", false},
	"SS": {"South Sudan", false},
	"SR": {"Suriname", false},
	"KI": {"Kiribati", false},
	"KH": {"Cambodia", false},
	"KN": {"Saint Kitts and Nevis", false},
	"KM": {"Comoros", false},
	"ST": {"Sao Tome and Principe", false},
	"SK": {"Slovakia", true},
	"KR": {"South Korea", false},
	"SI": {"Slovenia", true},
	"KP": {"North Korea", false},
	"KW": {"Kuwait", false},
	"SN": {"Senegal", false},
	"SM": {"San Marino", false},
	"SL": {"Sierra Leone", false},
	"SC": {"Seychelles", false},
	"KZ": {"Kazakhstan", false},
	"KY": {"Cayman Islands", false},
	"SG": {"Singapore", false},
	"SE": {"Sweden", true},
	"SD": {"Sudan", false},
	"DO": {"Dominican Republic", false},
	"DM": {"Dominica", false},
	"DJ": {"Djibouti", false},
	"DK": {"Denmark", true},
	"VG": {"British Virgin Islands", false},
	"DE": {"Germany", true},
	"YE": {"Yemen", false},
	"DZ": {"Algeria", false},
	"US": {"United States", false},
	"UY": {"Uruguay", false},
	"YT": {"Mayotte", false},
	"UM": {"United States Minor Outlying Islands", false},
	"LB": {"Lebanon", false},
	"LC": {"Saint Lucia", false},
	"LA": {"Laos", false},
	"TV": {"Tuvalu", false},
	"TW": {"Taiwan", false},
	"TT": {"Trinidad and Tobago", false},
	"TR": {"Turkey", false},
	"LK": {"Sri Lanka", false},
	"LI": {"Liechtenstein", false},
	"LV": {"Latvia", true},
	"TO": {"Tonga", false},
	"LT": {"Lithuania", true},
	"LU": {"Luxembourg", true},
	"LR": {"Liberia", false},
	"LS": {"Lesotho", false},
	"TH": {"Thailand", false},
	"TF": {"French Southern Territories", false},
	"TG": {"Togo", false},
	"TD": {"Chad", false},
	"TC": {"Turks and Caicos Islands", false},
	"LY": {"Libya", false},
	"VA": {"Vatican", false},
	"VC": {"Saint Vincent and the Grenadines", false},
	"AE": {"United Arab Emirates", false},
	"AD": {"Andorra", false},
	"AG": {"Antigua and Barbuda", false},
	"AF": {"Afghanistan", false},
	"AI": {"Anguilla", false},
	"VI": {"U.S. Virgin Islands", false},
	"IS": {"Iceland", false},
	"IR": {"Iran", false},
	"AM": {"Armenia", false},
	"AL": {"Albania", false},
	"AO": {"Angola", false},
	"AQ": {"Antarctica", false},
	"AS": {"American Samoa", false},
	"AR": {"Argentina", false},
	"AU": {"Australia", false},
	"AT": {"Austria", true},
	"AW": {"Aruba", false},
	"IN": {"India", false},
	"AX": {"Aland Islands", false},
	"AZ": {"Azerbaijan", false},
	"IE": {"Ireland", true},
	"ID": {"Indonesia", false},
	"UA": {"Ukraine", false},
	"QA": {"Qatar", false},
	"MZ": {"Mozambique", false},
}
