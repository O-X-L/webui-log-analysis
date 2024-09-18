package cnf

import "net"

const GEOIP_DB_TYPE_IPINFO uint8 = 1

// IPInfo schema: https://github.com/ipinfo/sample-database/
var IPINFO_COUNTRY struct {
	StartIp       net.IP `maxminddb:"start_ip"`
	EndIp         net.IP `maxminddb:"end_ip"`
	Country       string `maxminddb:"country"`
	CountryName   string `maxminddb:"country_name"`
	Continent     string `maxminddb:"continent"`
	ContinentName string `maxminddb:"continent_name"`
}

const IPINFO_COUNTRY_DEFAULT_ATTR = "Country"

var IPINFO_ASN struct {
	StartIp net.IP `maxminddb:"start_ip"`
	EndIp   net.IP `maxminddb:"end_ip"`
	ASN     string `maxminddb:"asn"`
	Name    string `maxminddb:"name"`
	Domain  string `maxminddb:"domain"`
}

const IPINFO_ASN_DEFAULT_ATTR = "ASN"

var IPINFO_ASN_EXT struct {
	StartIp net.IP `maxminddb:"start_ip"`
	EndIp   net.IP `maxminddb:"end_ip"`
	JoinKey net.IP `maxminddb:"join_key"`
	ASN     string `maxminddb:"asn"`
	Name    string `maxminddb:"name"`
	Domain  string `maxminddb:"domain"`
	Type    string `maxminddb:"type"`
	Country string `maxminddb:"country"`
}

var IPINFO_COUNTRY_ASN struct {
	StartIp       net.IP `maxminddb:"start_ip"`
	EndIp         net.IP `maxminddb:"end_ip"`
	Country       string `maxminddb:"country"`
	CountryName   string `maxminddb:"country_name"`
	Continent     string `maxminddb:"continent"`
	ContinentName string `maxminddb:"continent_name"`
	ASN           string `maxminddb:"asn"`
	ASName        string `maxminddb:"as_name"`
	ASDomain      string `maxminddb:"as_domain"`
}

var IPINFO_PRIVACY struct {
	StartIp net.IP `maxminddb:"start_ip"`
	EndIp   net.IP `maxminddb:"end_ip"`
	JoinKey net.IP `maxminddb:"join_key"`
	Hosting bool   `maxminddb:"hosting"`
	Proxy   bool   `maxminddb:"proxy"`
	Tor     bool   `maxminddb:"tor"`
	Vpn     bool   `maxminddb:"vpn"`
	Relay   bool   `maxminddb:"relay"`
	Service string `maxminddb:"service"`
}

const IPINFO_PRIVACY_DEFAULT_ATTR = "City"

var IPINFO_CITY struct {
	StartIp    net.IP  `maxminddb:"start_ip"`
	EndIp      net.IP  `maxminddb:"end_ip"`
	JoinKey    net.IP  `maxminddb:"join_key"`
	City       string  `maxminddb:"city"`
	Region     string  `maxminddb:"region"`
	Country    string  `maxminddb:"country"`
	Latitude   float32 `maxminddb:"latitude"`
	Longitude  float32 `maxminddb:"longitude"`
	PostalCode string  `maxminddb:"postal_code"`
	Timezone   string  `maxminddb:"timezone"`
}

const IPINFO_CITY_DEFAULT_ATTR = "City"

// todo: https://github.com/ipinfo/sample-database/tree/main/IP%20to%20Company
// todo: https://github.com/ipinfo/sample-database/tree/main/IP%20to%20Mobile%20Carrier
// todo: https://github.com/ipinfo/sample-database/tree/main/IP%20Geolocation%20Extended
// todo: https://github.com/ipinfo/sample-database/tree/main/Privacy%20Detection%20Extended
// todo: https://github.com/ipinfo/sample-database/tree/main/Abuse%20Contact
// todo: https://github.com/ipinfo/sample-database/tree/main/Hosted%20Domains
