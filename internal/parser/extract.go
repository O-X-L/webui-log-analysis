package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/O-X-L/webui-log-analysis/internal/cnf"
)

func extract(line string, e cnf.Extractor) interface{} {
	if e.Multiple {
		// count matches and run extractor N times
		return extractSingle(line, e)
	} else {
		return extractSingle(line, e)
	}
}

func extractSingle(line string, e cnf.Extractor) interface{} {
	d := ""

	if e.Search != "" {
		d = findSearch(line, e.Search)
	} else if e.Regex != "" {
		d = findRegex(line, e.Regex)
	}

	if e.Cleanup.Remove != "" {
		d = cleanRemove(d, e.Cleanup.Remove)
	}

	return typeCasting(d, e)
}

func typeCasting(d string, e cnf.Extractor) interface{} {
	f := e.Fallback
	if e.Kind == "int" || e.Kind == "digit" || e.Kind == "number" {
		d2, err := typingInt(d)
		if err != nil {
			return d2
		} else if f != "" {
			d3, err := typingInt(f)
			if err == nil {
				return d3
			} else {
				return f
			}
		}
	} else if e.Kind == "float" {
		d2, err := typingFloat(d)
		if err != nil {
			return d2
		} else if f != "" {
			d3, err := typingFloat(f)
			if err == nil {
				return d3
			} else {
				return f
			}
		}
	} else if e.Kind == "time" {
		d2, err := time.Parse(e.TimeFormat, d)
		if err != nil {
			fmt.Printf("ERROR | unable to parse as time: '%v'\n", d)
			return d2
		}
	}
	return d
}

func findRegex(v string, pattern string) string {
	if cnf.DEBUG {
		fmt.Println("DEBUG | Regex |", pattern)
	}
	m := regexp.MustCompile(pattern).FindStringSubmatch(v)
	if len(m) >= 2 {
		return m[1]
	}
	return ""
}

func findSearch(v string, pattern string) string {
	return findRegex(
		v,
		".*"+strings.ReplaceAll(
			strings.ReplaceAll(
				regexp.QuoteMeta(pattern),
				" ",
				".*?",
			),
			cnf.EXTRACT_SEARCH_MATCH,
			"(.*?)",
		)+".*",
	)
}

func cleanRemove(d string, remove string) string {
	return strings.ReplaceAll(d, remove, "")
}

func typingInt(d string) (int, error) {
	if d == "" {
		return 0, errors.New("empty input")
	}

	dt, err := strconv.Atoi(d)
	if err != nil {
		fmt.Printf("ERROR | unable to parse as int: '%v'\n", d)
		return 0, err
	}
	return dt, nil
}

func typingFloat(d string) (float64, error) {
	if d == "" {
		return 0, errors.New("empty input")
	}

	dt, err := strconv.ParseFloat(d, 32)
	if err != nil {
		fmt.Printf("ERROR | unable to parse as float: '%v'\n", d)
		return 0, err
	}
	return dt, nil
}
