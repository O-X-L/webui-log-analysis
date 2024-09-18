package main

import (
	"flag"
	"fmt"

	"github.com/O-X-L/webui-log-analysis/internal/cnf"
	"github.com/O-X-L/webui-log-analysis/internal/parser"
	"github.com/O-X-L/webui-log-analysis/web"
)

func welcome() {
	fmt.Println("Log-Analysis WebUI")
	fmt.Printf("Version: %v\n", cnf.VERSION)
	fmt.Printf("by OXL IT Services (License: GPLv3)\n\n")
}

func main() {
	var listenPort uint
	var listenAddr string
	var configFile string

	flag.UintVar(&listenPort, "p", 8000, "Port to listen on")
	flag.StringVar(&listenAddr, "l", "127.0.0.1", "Address to listen on")
	flag.StringVar(&configFile, "c", "config.yml", "Path to the config-file")
	flag.BoolVar(&cnf.DEBUG, "d", false, "Debug mode")
	flag.Parse()

	welcome()
	cnf.LoadConfig(configFile)

	if cnf.DEBUG {
		fmt.Printf("DEBUG | Config\n%+v\n\n", cnf.Config)
	}

	go parser.Main()
	web.Server(fmt.Sprintf("%v:%v", listenAddr, listenPort))
}
