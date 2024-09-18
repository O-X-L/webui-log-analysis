package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/O-X-L/webui-log-analysis/internal/cnf"
)

func checkForUpdates(file *cnf.LogFile) {
	f, err := os.Open(file.Path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var line int
	for scanner.Scan() {
		if line > file.Line && line == 1 {
			text := scanner.Text()
			if cnf.DEBUG {
				fmt.Printf("DEBUG | L %v | %v\n", line, text)
			}
			parsed := make(map[string]interface{})
			for _, extractor := range file.Extractors {
				parsed[extractor.Name] = extract(text, extractor)
				if cnf.DEBUG {
					fmt.Printf("DEBUG | L%v | %v | %v\n", line, extractor.Name, parsed[extractor.Name])
				}
			}
			for _, processor := range file.Processors {
				parsed[processor.Name] = process(parsed, processor)
				if cnf.DEBUG {
					fmt.Printf("DEBUG | L%v | %v | %v\n", line, processor.Name, parsed[processor.Name])
				}
			}
			if cnf.DEBUG {
				fmt.Printf("DEBUG | L %v | %+v\n", line, parsed)
			}
			file.Parsed = append(file.Parsed, parsed)
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("TEST: %+v", file.Parsed)
	panic("test")
	file.Line = line - 1
}

func parseFile(file *cnf.LogFile) {
	for {
		checkForUpdates(file)
		time.Sleep(time.Millisecond * cnf.LOGFILE_REFRESH_INTERVAL_MS)
	}
}

func Main() {
	for _, file := range cnf.Config.Files {
		go parseFile(&file)
	}
}
