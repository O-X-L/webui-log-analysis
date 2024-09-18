package parser

import (
	"fmt"
	"net"
	"reflect"
	"strings"

	"github.com/O-X-L/webui-log-analysis/internal/cnf"
	"github.com/oschwald/maxminddb-golang"
)

func lookupBase(ip net.IP, dataStructure interface{}, dbFile string) (interface{}, error) {
	db, err := maxminddb.Open(dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Lookup(ip, &dataStructure)
	if err != nil {
		return nil, err
	}
	return dataStructure, nil
}

func ipInfoCountry(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_COUNTRY, cnf.Config.GeoIP.IPInfoCountry)
}

func ipInfoLocation(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_CITY, cnf.Config.GeoIP.IPInfoLocation)
}

func ipInfoAsn(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_ASN, cnf.Config.GeoIP.IPInfoASN)
}

/*
func ipInfoPrivacy(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_PRIVACY, cnf.Config.GeoIP.I)
}
*/

func maxMindCountry(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.MAXMIND_COUNTRY, cnf.Config.GeoIP.MaxMindCountry)
}

func maxMindCity(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.MAXMIND_CITY, cnf.Config.GeoIP.MaxMindCity)
}

func maxMindAsn(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.MAXMIND_ASN, cnf.Config.GeoIP.MaxMindASN)
}

/*
func maxMindPrivacy(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.MAXMIND_PRIVACY, cnf.DB_PRIVACY)
}
*/

var GEOIP_LOOKUPS = map[string]interface{}{
	"ipinfo_country":  ipInfoCountry,
	"ipinfo_asn":      ipInfoAsn,
	"ipinfo_city":     ipInfoLocation,
	"maxmind_country": maxMindCountry,
	"maxmind_asn":     maxMindAsn,
	"maxmind_city":    maxMindCity,
}

var GEOIP_LOOKUP_DEFAULT_ATTRIBUTES = map[string]string{
	"ipinfo_country":  "country",
	"ipinfo_asn":      "asn",
	"ipinfo_city":     "city",
	"maxmind_country": "country.iso_code",
	"maxmind_asn":     "asn",
	"maxmind_city":    "city.iso_code",
}

func getMapValue(dataStructure interface{}, name string) interface{} {
	return reflect.Indirect(
		reflect.ValueOf(&dataStructure),
	).Elem().Interface().(map[string]interface{})[name]
}

// v = possible IP, k = db key, a = attribute
func geoipLookup(v string, k string, a string) interface{} {
	if v == "" {
		return ""
	}

	if a == "" {
		a = GEOIP_LOOKUP_DEFAULT_ATTRIBUTES[k]
	}

	ip := net.ParseIP(v)
	if ip == nil {
		fmt.Printf("ERROR | GeoIP lookup failed for: '%v' (invalid IP)\n", v)
		return ""
	}

	f := GEOIP_LOOKUPS[k]
	if f == nil {
		fmt.Printf("ERROR | GeoIP database not a valid choice: '%v'\n", k)
		return ""
	}

	d, err := f.(func(net.IP) (interface{}, error))(ip)
	if err != nil {
		fmt.Printf("ERROR | GeoIP lookup failed for: '%v'\n", v)
		return ""
	}

	if !strings.Contains(a, ".") {
		return getMapValue(d, a)
	}

	d2 := d
	for _, as := range strings.Split(a, ".") {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf(
					"ERROR | GeoIP lookup failed for: '%v' (invalid attribute '%v' not found in '%+v')\n",
					v, a, d,
				)
			}
		}()
		d2 = getMapValue(d2, as)
		if d2 == nil {
			fmt.Printf(
				"ERROR | GeoIP lookup failed for: '%v' (invalid attribute '%v' not found in '%+v')\n",
				v, as, d,
			)
			return ""
		}
	}
	return d2
}
