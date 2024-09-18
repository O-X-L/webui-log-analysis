package parser

import (
	"fmt"

	"github.com/O-X-L/webui-log-analysis/internal/cnf"
)

func process(parsed map[string]interface{}, p cnf.Processor) (d string) {
	if p.From == "" || parsed[p.From] == nil {
		return ""
	}

	v := fmt.Sprintf("%v", parsed[p.From])
	if p.Search != "" {
		d = findSearch(v, p.Search)
	} else if p.Regex != "" {
		d = findRegex(v, p.Regex)
	} else if p.GeoIP.DB != "" {
		d = fmt.Sprintf(
			"%v", geoipLookup(v, p.GeoIP.DB, p.GeoIP.Attribute),
		)
	}

	if p.Cleanup.Remove != "" {
		d = cleanRemove(d, p.Cleanup.Remove)
	}

	return
}
