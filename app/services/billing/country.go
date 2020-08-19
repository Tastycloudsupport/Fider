package billing

import (
	"context"

	"github.com/getfider/fider/app"

	"github.com/getfider/fider/app/models/dto"
	"github.com/getfider/fider/app/models/query"
	"github.com/goenning/vat"
)

func getAllCountries(ctx context.Context, q *query.GetAllCountries) error {
	q.Result = allCountries
	return nil
}

func getCountryByCode(ctx context.Context, q *query.GetCountryByCode) error {
	for _, c := range allCountries {
		if c.Code == q.Code {
			q.Result = c
			return nil
		}
	}
	return app.ErrNotFound
}

var allCountries = []*dto.Country{
	&dto.Country{Code: "AF", Name: "Afghanistan", IsEU: vat.IsEU("AF")},
	&dto.Country{Code: "AX", Name: "Åland Islands", IsEU: vat.IsEU("AX")},
	&dto.Country{Code: "AL", Name: "Albania", IsEU: vat.IsEU("AL")},
	&dto.Country{Code: "DZ", Name: "Algeria", IsEU: vat.IsEU("DZ")},
	&dto.Country{Code: "AS", Name: "American Samoa", IsEU: vat.IsEU("AS")},
	&dto.Country{Code: "AD", Name: "Andorra", IsEU: vat.IsEU("AD")},
	&dto.Country{Code: "AO", Name: "Angola", IsEU: vat.IsEU("AO")},
	&dto.Country{Code: "AI", Name: "Anguilla", IsEU: vat.IsEU("AI")},
	&dto.Country{Code: "AQ", Name: "Antarctica", IsEU: vat.IsEU("AQ")},
	&dto.Country{Code: "AG", Name: "Antigua and Barbuda", IsEU: vat.IsEU("AG")},
	&dto.Country{Code: "AR", Name: "Argentina", IsEU: vat.IsEU("AR")},
	&dto.Country{Code: "AM", Name: "Armenia", IsEU: vat.IsEU("AM")},
	&dto.Country{Code: "AW", Name: "Aruba", IsEU: vat.IsEU("AW")},
	&dto.Country{Code: "AU", Name: "Australia", IsEU: vat.IsEU("AU")},
	&dto.Country{Code: "AT", Name: "Austria", IsEU: vat.IsEU("AT")},
	&dto.Country{Code: "AZ", Name: "Azerbaijan", IsEU: vat.IsEU("AZ")},
	&dto.Country{Code: "BS", Name: "Bahamas", IsEU: vat.IsEU("BS")},
	&dto.Country{Code: "BH", Name: "Bahrain", IsEU: vat.IsEU("BH")},
	&dto.Country{Code: "BD", Name: "Bangladesh", IsEU: vat.IsEU("BD")},
	&dto.Country{Code: "BB", Name: "Barbados", IsEU: vat.IsEU("BB")},
	&dto.Country{Code: "BY", Name: "Belarus", IsEU: vat.IsEU("BY")},
	&dto.Country{Code: "BE", Name: "Belgium", IsEU: vat.IsEU("BE")},
	&dto.Country{Code: "BZ", Name: "Belize", IsEU: vat.IsEU("BZ")},
	&dto.Country{Code: "BJ", Name: "Benin", IsEU: vat.IsEU("BJ")},
	&dto.Country{Code: "BM", Name: "Bermuda", IsEU: vat.IsEU("BM")},
	&dto.Country{Code: "BT", Name: "Bhutan", IsEU: vat.IsEU("BT")},
	&dto.Country{Code: "BO", Name: "Bolivia, Plurinational State of", IsEU: vat.IsEU("BO")},
	&dto.Country{Code: "BQ", Name: "Bonaire, Sint Eustatius and Saba", IsEU: vat.IsEU("BQ")},
	&dto.Country{Code: "BA", Name: "Bosnia and Herzegovina", IsEU: vat.IsEU("BA")},
	&dto.Country{Code: "BW", Name: "Botswana", IsEU: vat.IsEU("BW")},
	&dto.Country{Code: "BV", Name: "Bouvet Island", IsEU: vat.IsEU("BV")},
	&dto.Country{Code: "BR", Name: "Brazil", IsEU: vat.IsEU("BR")},
	&dto.Country{Code: "IO", Name: "British Indian Ocean Territory", IsEU: vat.IsEU("IO")},
	&dto.Country{Code: "BN", Name: "Brunei Darussalam", IsEU: vat.IsEU("BN")},
	&dto.Country{Code: "BG", Name: "Bulgaria", IsEU: vat.IsEU("BG")},
	&dto.Country{Code: "BF", Name: "Burkina Faso", IsEU: vat.IsEU("BF")},
	&dto.Country{Code: "BI", Name: "Burundi", IsEU: vat.IsEU("BI")},
	&dto.Country{Code: "KH", Name: "Cambodia", IsEU: vat.IsEU("KH")},
	&dto.Country{Code: "CM", Name: "Cameroon", IsEU: vat.IsEU("CM")},
	&dto.Country{Code: "CA", Name: "Canada", IsEU: vat.IsEU("CA")},
	&dto.Country{Code: "CV", Name: "Cape Verde", IsEU: vat.IsEU("CV")},
	&dto.Country{Code: "KY", Name: "Cayman Islands", IsEU: vat.IsEU("KY")},
	&dto.Country{Code: "CF", Name: "Central African Republic", IsEU: vat.IsEU("CF")},
	&dto.Country{Code: "TD", Name: "Chad", IsEU: vat.IsEU("TD")},
	&dto.Country{Code: "CL", Name: "Chile", IsEU: vat.IsEU("CL")},
	&dto.Country{Code: "CN", Name: "China", IsEU: vat.IsEU("CN")},
	&dto.Country{Code: "CX", Name: "Christmas Island", IsEU: vat.IsEU("CX")},
	&dto.Country{Code: "CC", Name: "Cocos (Keeling) Islands", IsEU: vat.IsEU("CC")},
	&dto.Country{Code: "CO", Name: "Colombia", IsEU: vat.IsEU("CO")},
	&dto.Country{Code: "KM", Name: "Comoros", IsEU: vat.IsEU("KM")},
	&dto.Country{Code: "CG", Name: "Congo", IsEU: vat.IsEU("CG")},
	&dto.Country{Code: "CD", Name: "Congo, the Democratic Republic of the", IsEU: vat.IsEU("CD")},
	&dto.Country{Code: "CK", Name: "Cook Islands", IsEU: vat.IsEU("CK")},
	&dto.Country{Code: "CR", Name: "Costa Rica", IsEU: vat.IsEU("CR")},
	&dto.Country{Code: "CI", Name: "Côte d'Ivoire", IsEU: vat.IsEU("CI")},
	&dto.Country{Code: "HR", Name: "Croatia", IsEU: vat.IsEU("HR")},
	&dto.Country{Code: "CU", Name: "Cuba", IsEU: vat.IsEU("CU")},
	&dto.Country{Code: "CW", Name: "Curaçao", IsEU: vat.IsEU("CW")},
	&dto.Country{Code: "CY", Name: "Cyprus", IsEU: vat.IsEU("CY")},
	&dto.Country{Code: "CZ", Name: "Czech Republic", IsEU: vat.IsEU("CZ")},
	&dto.Country{Code: "DK", Name: "Denmark", IsEU: vat.IsEU("DK")},
	&dto.Country{Code: "DJ", Name: "Djibouti", IsEU: vat.IsEU("DJ")},
	&dto.Country{Code: "DM", Name: "Dominica", IsEU: vat.IsEU("DM")},
	&dto.Country{Code: "DO", Name: "Dominican Republic", IsEU: vat.IsEU("DO")},
	&dto.Country{Code: "EC", Name: "Ecuador", IsEU: vat.IsEU("EC")},
	&dto.Country{Code: "EG", Name: "Egypt", IsEU: vat.IsEU("EG")},
	&dto.Country{Code: "SV", Name: "El Salvador", IsEU: vat.IsEU("SV")},
	&dto.Country{Code: "GQ", Name: "Equatorial Guinea", IsEU: vat.IsEU("GQ")},
	&dto.Country{Code: "ER", Name: "Eritrea", IsEU: vat.IsEU("ER")},
	&dto.Country{Code: "EE", Name: "Estonia", IsEU: vat.IsEU("EE")},
	&dto.Country{Code: "ET", Name: "Ethiopia", IsEU: vat.IsEU("ET")},
	&dto.Country{Code: "FK", Name: "Falkland Islands (Malvinas)", IsEU: vat.IsEU("FK")},
	&dto.Country{Code: "FO", Name: "Faroe Islands", IsEU: vat.IsEU("FO")},
	&dto.Country{Code: "FJ", Name: "Fiji", IsEU: vat.IsEU("FJ")},
	&dto.Country{Code: "FI", Name: "Finland", IsEU: vat.IsEU("FI")},
	&dto.Country{Code: "FR", Name: "France", IsEU: vat.IsEU("FR")},
	&dto.Country{Code: "GF", Name: "French Guiana", IsEU: vat.IsEU("GF")},
	&dto.Country{Code: "PF", Name: "French Polynesia", IsEU: vat.IsEU("PF")},
	&dto.Country{Code: "TF", Name: "French Southern Territories", IsEU: vat.IsEU("TF")},
	&dto.Country{Code: "GA", Name: "Gabon", IsEU: vat.IsEU("GA")},
	&dto.Country{Code: "GM", Name: "Gambia", IsEU: vat.IsEU("GM")},
	&dto.Country{Code: "GE", Name: "Georgia", IsEU: vat.IsEU("GE")},
	&dto.Country{Code: "DE", Name: "Germany", IsEU: vat.IsEU("DE")},
	&dto.Country{Code: "GH", Name: "Ghana", IsEU: vat.IsEU("GH")},
	&dto.Country{Code: "GI", Name: "Gibraltar", IsEU: vat.IsEU("GI")},
	&dto.Country{Code: "GR", Name: "Greece", IsEU: vat.IsEU("GR")},
	&dto.Country{Code: "GL", Name: "Greenland", IsEU: vat.IsEU("GL")},
	&dto.Country{Code: "GD", Name: "Grenada", IsEU: vat.IsEU("GD")},
	&dto.Country{Code: "GP", Name: "Guadeloupe", IsEU: vat.IsEU("GP")},
	&dto.Country{Code: "GU", Name: "Guam", IsEU: vat.IsEU("GU")},
	&dto.Country{Code: "GT", Name: "Guatemala", IsEU: vat.IsEU("GT")},
	&dto.Country{Code: "GG", Name: "Guernsey", IsEU: vat.IsEU("GG")},
	&dto.Country{Code: "GN", Name: "Guinea", IsEU: vat.IsEU("GN")},
	&dto.Country{Code: "GW", Name: "Guinea-Bissau", IsEU: vat.IsEU("GW")},
	&dto.Country{Code: "GY", Name: "Guyana", IsEU: vat.IsEU("GY")},
	&dto.Country{Code: "HT", Name: "Haiti", IsEU: vat.IsEU("HT")},
	&dto.Country{Code: "HM", Name: "Heard Island and McDonald Islands", IsEU: vat.IsEU("HM")},
	&dto.Country{Code: "VA", Name: "Holy See (Vatican City State)", IsEU: vat.IsEU("VA")},
	&dto.Country{Code: "HN", Name: "Honduras", IsEU: vat.IsEU("HN")},
	&dto.Country{Code: "HK", Name: "Hong Kong", IsEU: vat.IsEU("HK")},
	&dto.Country{Code: "HU", Name: "Hungary", IsEU: vat.IsEU("HU")},
	&dto.Country{Code: "IS", Name: "Iceland", IsEU: vat.IsEU("IS")},
	&dto.Country{Code: "IN", Name: "India", IsEU: vat.IsEU("IN")},
	&dto.Country{Code: "ID", Name: "Indonesia", IsEU: vat.IsEU("ID")},
	&dto.Country{Code: "IR", Name: "Iran, Islamic Republic of", IsEU: vat.IsEU("IR")},
	&dto.Country{Code: "IQ", Name: "Iraq", IsEU: vat.IsEU("IQ")},
	&dto.Country{Code: "IE", Name: "Ireland", IsEU: vat.IsEU("IE")},
	&dto.Country{Code: "IM", Name: "Isle of Man", IsEU: vat.IsEU("IM")},
	&dto.Country{Code: "IL", Name: "Israel", IsEU: vat.IsEU("IL")},
	&dto.Country{Code: "IT", Name: "Italy", IsEU: vat.IsEU("IT")},
	&dto.Country{Code: "JM", Name: "Jamaica", IsEU: vat.IsEU("JM")},
	&dto.Country{Code: "JP", Name: "Japan", IsEU: vat.IsEU("JP")},
	&dto.Country{Code: "JE", Name: "Jersey", IsEU: vat.IsEU("JE")},
	&dto.Country{Code: "JO", Name: "Jordan", IsEU: vat.IsEU("JO")},
	&dto.Country{Code: "KZ", Name: "Kazakhstan", IsEU: vat.IsEU("KZ")},
	&dto.Country{Code: "KE", Name: "Kenya", IsEU: vat.IsEU("KE")},
	&dto.Country{Code: "KI", Name: "Kiribati", IsEU: vat.IsEU("KI")},
	&dto.Country{Code: "KP", Name: "Korea, Democratic People's Republic of", IsEU: vat.IsEU("KP")},
	&dto.Country{Code: "KR", Name: "Korea, Republic of", IsEU: vat.IsEU("KR")},
	&dto.Country{Code: "KW", Name: "Kuwait", IsEU: vat.IsEU("KW")},
	&dto.Country{Code: "KG", Name: "Kyrgyzstan", IsEU: vat.IsEU("KG")},
	&dto.Country{Code: "LA", Name: "Lao People's Democratic Republic", IsEU: vat.IsEU("LA")},
	&dto.Country{Code: "LV", Name: "Latvia", IsEU: vat.IsEU("LV")},
	&dto.Country{Code: "LB", Name: "Lebanon", IsEU: vat.IsEU("LB")},
	&dto.Country{Code: "LS", Name: "Lesotho", IsEU: vat.IsEU("LS")},
	&dto.Country{Code: "LR", Name: "Liberia", IsEU: vat.IsEU("LR")},
	&dto.Country{Code: "LY", Name: "Libya", IsEU: vat.IsEU("LY")},
	&dto.Country{Code: "LI", Name: "Liechtenstein", IsEU: vat.IsEU("LI")},
	&dto.Country{Code: "LT", Name: "Lithuania", IsEU: vat.IsEU("LT")},
	&dto.Country{Code: "LU", Name: "Luxembourg", IsEU: vat.IsEU("LU")},
	&dto.Country{Code: "MO", Name: "Macao", IsEU: vat.IsEU("MO")},
	&dto.Country{Code: "MK", Name: "Macedonia, the former Yugoslav Republic of", IsEU: vat.IsEU("MK")},
	&dto.Country{Code: "MG", Name: "Madagascar", IsEU: vat.IsEU("MG")},
	&dto.Country{Code: "MW", Name: "Malawi", IsEU: vat.IsEU("MW")},
	&dto.Country{Code: "MY", Name: "Malaysia", IsEU: vat.IsEU("MY")},
	&dto.Country{Code: "MV", Name: "Maldives", IsEU: vat.IsEU("MV")},
	&dto.Country{Code: "ML", Name: "Mali", IsEU: vat.IsEU("ML")},
	&dto.Country{Code: "MT", Name: "Malta", IsEU: vat.IsEU("MT")},
	&dto.Country{Code: "MH", Name: "Marshall Islands", IsEU: vat.IsEU("MH")},
	&dto.Country{Code: "MQ", Name: "Martinique", IsEU: vat.IsEU("MQ")},
	&dto.Country{Code: "MR", Name: "Mauritania", IsEU: vat.IsEU("MR")},
	&dto.Country{Code: "MU", Name: "Mauritius", IsEU: vat.IsEU("MU")},
	&dto.Country{Code: "YT", Name: "Mayotte", IsEU: vat.IsEU("YT")},
	&dto.Country{Code: "MX", Name: "Mexico", IsEU: vat.IsEU("MX")},
	&dto.Country{Code: "FM", Name: "Micronesia, Federated States of", IsEU: vat.IsEU("FM")},
	&dto.Country{Code: "MD", Name: "Moldova, Republic of", IsEU: vat.IsEU("MD")},
	&dto.Country{Code: "MC", Name: "Monaco", IsEU: vat.IsEU("MC")},
	&dto.Country{Code: "MN", Name: "Mongolia", IsEU: vat.IsEU("MN")},
	&dto.Country{Code: "ME", Name: "Montenegro", IsEU: vat.IsEU("ME")},
	&dto.Country{Code: "MS", Name: "Montserrat", IsEU: vat.IsEU("MS")},
	&dto.Country{Code: "MA", Name: "Morocco", IsEU: vat.IsEU("MA")},
	&dto.Country{Code: "MZ", Name: "Mozambique", IsEU: vat.IsEU("MZ")},
	&dto.Country{Code: "MM", Name: "Myanmar", IsEU: vat.IsEU("MM")},
	&dto.Country{Code: "NA", Name: "Namibia", IsEU: vat.IsEU("NA")},
	&dto.Country{Code: "NR", Name: "Nauru", IsEU: vat.IsEU("NR")},
	&dto.Country{Code: "NP", Name: "Nepal", IsEU: vat.IsEU("NP")},
	&dto.Country{Code: "NL", Name: "Netherlands", IsEU: vat.IsEU("NL")},
	&dto.Country{Code: "NC", Name: "New Caledonia", IsEU: vat.IsEU("NC")},
	&dto.Country{Code: "NZ", Name: "New Zealand", IsEU: vat.IsEU("NZ")},
	&dto.Country{Code: "NI", Name: "Nicaragua", IsEU: vat.IsEU("NI")},
	&dto.Country{Code: "NE", Name: "Niger", IsEU: vat.IsEU("NE")},
	&dto.Country{Code: "NG", Name: "Nigeria", IsEU: vat.IsEU("NG")},
	&dto.Country{Code: "NU", Name: "Niue", IsEU: vat.IsEU("NU")},
	&dto.Country{Code: "NF", Name: "Norfolk Island", IsEU: vat.IsEU("NF")},
	&dto.Country{Code: "MP", Name: "Northern Mariana Islands", IsEU: vat.IsEU("MP")},
	&dto.Country{Code: "NO", Name: "Norway", IsEU: vat.IsEU("NO")},
	&dto.Country{Code: "OM", Name: "Oman", IsEU: vat.IsEU("OM")},
	&dto.Country{Code: "PK", Name: "Pakistan", IsEU: vat.IsEU("PK")},
	&dto.Country{Code: "PW", Name: "Palau", IsEU: vat.IsEU("PW")},
	&dto.Country{Code: "PS", Name: "Palestinian Territory, Occupied", IsEU: vat.IsEU("PS")},
	&dto.Country{Code: "PA", Name: "Panama", IsEU: vat.IsEU("PA")},
	&dto.Country{Code: "PG", Name: "Papua New Guinea", IsEU: vat.IsEU("PG")},
	&dto.Country{Code: "PY", Name: "Paraguay", IsEU: vat.IsEU("PY")},
	&dto.Country{Code: "PE", Name: "Peru", IsEU: vat.IsEU("PE")},
	&dto.Country{Code: "PH", Name: "Philippines", IsEU: vat.IsEU("PH")},
	&dto.Country{Code: "PN", Name: "Pitcairn", IsEU: vat.IsEU("PN")},
	&dto.Country{Code: "PL", Name: "Poland", IsEU: vat.IsEU("PL")},
	&dto.Country{Code: "PT", Name: "Portugal", IsEU: vat.IsEU("PT")},
	&dto.Country{Code: "PR", Name: "Puerto Rico", IsEU: vat.IsEU("PR")},
	&dto.Country{Code: "QA", Name: "Qatar", IsEU: vat.IsEU("QA")},
	&dto.Country{Code: "RE", Name: "Réunion", IsEU: vat.IsEU("RE")},
	&dto.Country{Code: "RO", Name: "Romania", IsEU: vat.IsEU("RO")},
	&dto.Country{Code: "RU", Name: "Russian Federation", IsEU: vat.IsEU("RU")},
	&dto.Country{Code: "RW", Name: "Rwanda", IsEU: vat.IsEU("RW")},
	&dto.Country{Code: "BL", Name: "Saint Barthélemy", IsEU: vat.IsEU("BL")},
	&dto.Country{Code: "SH", Name: "Saint Helena, Ascension and Tristan da Cunha", IsEU: vat.IsEU("SH")},
	&dto.Country{Code: "KN", Name: "Saint Kitts and Nevis", IsEU: vat.IsEU("KN")},
	&dto.Country{Code: "LC", Name: "Saint Lucia", IsEU: vat.IsEU("LC")},
	&dto.Country{Code: "MF", Name: "Saint Martin (French part)", IsEU: vat.IsEU("MF")},
	&dto.Country{Code: "PM", Name: "Saint Pierre and Miquelon", IsEU: vat.IsEU("PM")},
	&dto.Country{Code: "VC", Name: "Saint Vincent and the Grenadines", IsEU: vat.IsEU("VC")},
	&dto.Country{Code: "WS", Name: "Samoa", IsEU: vat.IsEU("WS")},
	&dto.Country{Code: "SM", Name: "San Marino", IsEU: vat.IsEU("SM")},
	&dto.Country{Code: "ST", Name: "Sao Tome and Principe", IsEU: vat.IsEU("ST")},
	&dto.Country{Code: "SA", Name: "Saudi Arabia", IsEU: vat.IsEU("SA")},
	&dto.Country{Code: "SN", Name: "Senegal", IsEU: vat.IsEU("SN")},
	&dto.Country{Code: "RS", Name: "Serbia", IsEU: vat.IsEU("RS")},
	&dto.Country{Code: "SC", Name: "Seychelles", IsEU: vat.IsEU("SC")},
	&dto.Country{Code: "SL", Name: "Sierra Leone", IsEU: vat.IsEU("SL")},
	&dto.Country{Code: "SG", Name: "Singapore", IsEU: vat.IsEU("SG")},
	&dto.Country{Code: "SX", Name: "Sint Maarten (Dutch part)", IsEU: vat.IsEU("SX")},
	&dto.Country{Code: "SK", Name: "Slovakia", IsEU: vat.IsEU("SK")},
	&dto.Country{Code: "SI", Name: "Slovenia", IsEU: vat.IsEU("SI")},
	&dto.Country{Code: "SB", Name: "Solomon Islands", IsEU: vat.IsEU("SB")},
	&dto.Country{Code: "SO", Name: "Somalia", IsEU: vat.IsEU("SO")},
	&dto.Country{Code: "ZA", Name: "South Africa", IsEU: vat.IsEU("ZA")},
	&dto.Country{Code: "GS", Name: "South Georgia and the South Sandwich Islands", IsEU: vat.IsEU("GS")},
	&dto.Country{Code: "SS", Name: "South Sudan", IsEU: vat.IsEU("SS")},
	&dto.Country{Code: "ES", Name: "Spain", IsEU: vat.IsEU("ES")},
	&dto.Country{Code: "LK", Name: "Sri Lanka", IsEU: vat.IsEU("LK")},
	&dto.Country{Code: "SD", Name: "Sudan", IsEU: vat.IsEU("SD")},
	&dto.Country{Code: "SR", Name: "Suriname", IsEU: vat.IsEU("SR")},
	&dto.Country{Code: "SJ", Name: "Svalbard and Jan Mayen", IsEU: vat.IsEU("SJ")},
	&dto.Country{Code: "SZ", Name: "Swaziland", IsEU: vat.IsEU("SZ")},
	&dto.Country{Code: "SE", Name: "Sweden", IsEU: vat.IsEU("SE")},
	&dto.Country{Code: "CH", Name: "Switzerland", IsEU: vat.IsEU("CH")},
	&dto.Country{Code: "SY", Name: "Syrian Arab Republic", IsEU: vat.IsEU("SY")},
	&dto.Country{Code: "TW", Name: "Taiwan, Province of China", IsEU: vat.IsEU("TW")},
	&dto.Country{Code: "TJ", Name: "Tajikistan", IsEU: vat.IsEU("TJ")},
	&dto.Country{Code: "TZ", Name: "Tanzania, United Republic of", IsEU: vat.IsEU("TZ")},
	&dto.Country{Code: "TH", Name: "Thailand", IsEU: vat.IsEU("TH")},
	&dto.Country{Code: "TL", Name: "Timor-Leste", IsEU: vat.IsEU("TL")},
	&dto.Country{Code: "TG", Name: "Togo", IsEU: vat.IsEU("TG")},
	&dto.Country{Code: "TK", Name: "Tokelau", IsEU: vat.IsEU("TK")},
	&dto.Country{Code: "TO", Name: "Tonga", IsEU: vat.IsEU("TO")},
	&dto.Country{Code: "TT", Name: "Trinidad and Tobago", IsEU: vat.IsEU("TT")},
	&dto.Country{Code: "TN", Name: "Tunisia", IsEU: vat.IsEU("TN")},
	&dto.Country{Code: "TR", Name: "Turkey", IsEU: vat.IsEU("TR")},
	&dto.Country{Code: "TM", Name: "Turkmenistan", IsEU: vat.IsEU("TM")},
	&dto.Country{Code: "TC", Name: "Turks and Caicos Islands", IsEU: vat.IsEU("TC")},
	&dto.Country{Code: "TV", Name: "Tuvalu", IsEU: vat.IsEU("TV")},
	&dto.Country{Code: "UG", Name: "Uganda", IsEU: vat.IsEU("UG")},
	&dto.Country{Code: "UA", Name: "Ukraine", IsEU: vat.IsEU("UA")},
	&dto.Country{Code: "AE", Name: "United Arab Emirates", IsEU: vat.IsEU("AE")},
	&dto.Country{Code: "GB", Name: "United Kingdom", IsEU: vat.IsEU("GB")},
	&dto.Country{Code: "US", Name: "United States", IsEU: vat.IsEU("US")},
	&dto.Country{Code: "UM", Name: "United States Minor Outlying Islands", IsEU: vat.IsEU("UM")},
	&dto.Country{Code: "UY", Name: "Uruguay", IsEU: vat.IsEU("UY")},
	&dto.Country{Code: "UZ", Name: "Uzbekistan", IsEU: vat.IsEU("UZ")},
	&dto.Country{Code: "VU", Name: "Vanuatu", IsEU: vat.IsEU("VU")},
	&dto.Country{Code: "VE", Name: "Venezuela, Bolivarian Republic of", IsEU: vat.IsEU("VE")},
	&dto.Country{Code: "VN", Name: "Viet Nam", IsEU: vat.IsEU("VN")},
	&dto.Country{Code: "VG", Name: "Virgin Islands, British", IsEU: vat.IsEU("VG")},
	&dto.Country{Code: "VI", Name: "Virgin Islands, U.S.", IsEU: vat.IsEU("VI")},
	&dto.Country{Code: "WF", Name: "Wallis and Futuna", IsEU: vat.IsEU("WF")},
	&dto.Country{Code: "EH", Name: "Western Sahara", IsEU: vat.IsEU("EH")},
	&dto.Country{Code: "YE", Name: "Yemen", IsEU: vat.IsEU("YE")},
	&dto.Country{Code: "ZM", Name: "Zambia", IsEU: vat.IsEU("ZM")},
	&dto.Country{Code: "ZW", Name: "Zimbabwe", IsEU: vat.IsEU("ZW")},
}
