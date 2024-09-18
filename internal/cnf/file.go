package cnf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type StringCleanup struct {
	Remove string `yaml:"remove"`
}

type Extractor struct {
	Name       string        `yaml:"name"`
	Search     string        `yaml:"search"`
	Regex      string        `yaml:"regex"`
	Cleanup    StringCleanup `yaml:"clean"`
	Kind       string        `yaml:"type"`
	TimeFormat string        `yaml:"time_format"`
	Fallback   string        `yaml:"fallback"`
	Multiple   bool          `yaml:"multiple"`
}

type ProcessorGeoIP struct {
	DB        string `yaml:"db"`        // ipinfo_asn / ipinfo_country
	Attribute string `yaml:"attribute"` // defaults (asn/country)
}

type Processor struct {
	Name    string         `yaml:"name"`
	From    string         `yaml:"from"`
	GeoIP   ProcessorGeoIP `yaml:"geoip"`
	Search  string         `yaml:"search"`
	Regex   string         `yaml:"regex"`
	Cleanup StringCleanup  `yaml:"clean"`
}

type LogFile struct {
	Path       string      `yaml:"path"`
	Prefix     string      `yaml:"prefix"`
	Extractors []Extractor `yaml:"extract"`
	Processors []Processor `yaml:"process"`
	// runtime
	Line   int
	Parsed []map[string]interface{}
}

type GeoIPDatabaseFiles struct {
	IPInfoASN      string `yaml:"ipinfo_asn"`
	IPInfoCountry  string `yaml:"ipinfo_country"`
	IPInfoLocation string `yaml:"ipinfo_location"`
	MaxMindASN     string `yaml:"maxmind_asn"`
	MaxMindCountry string `yaml:"maxmind_country"`
	MaxMindCity    string `yaml:"maxmind_city"`
}

type ConfigFile struct {
	Files []LogFile          `yaml:"files"`
	GeoIP GeoIPDatabaseFiles `yaml:"geoip"`
}

var Config ConfigFile

func CheckGeoIPFile(file string) {
	if file == "" {
		return
	}
	if _, err := os.Stat(file); err != nil {
		panic(fmt.Sprintf(
			"GeoIP file %v is configured but does not exist", file,
		))
	}
}

func ValidateConfig() {
	CheckGeoIPFile(Config.GeoIP.IPInfoASN)
	CheckGeoIPFile(Config.GeoIP.IPInfoCountry)
	CheckGeoIPFile(Config.GeoIP.MaxMindASN)
	CheckGeoIPFile(Config.GeoIP.MaxMindCountry)
}

func LoadConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic("failed to read config file")
	}

	yaml.Unmarshal(data, &Config)
	ValidateConfig()
}
