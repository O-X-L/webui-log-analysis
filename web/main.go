package web

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/O-X-L/webui-log-analysis/web/api"
	"github.com/O-X-L/webui-log-analysis/web/ws"
)

//go:embed all:static/*
var statics embed.FS

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(statics, "static")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}

func Server(listen string) {
	// backend
	http.HandleFunc("/api/test", api.Test)
	http.HandleFunc("/ws/test", ws.Test)

	// frontend
	http.Handle("/", http.FileServer(getFileSystem()))

	fmt.Println("Listening on http://" + listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}
