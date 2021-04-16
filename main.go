package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"go.yhsif.com/vanity"
)

func main() {
	mappings := []vanity.Mapping{
		{
			Path:        "/aula-assistant",
			URL:         "https://github.com/arnested/aula-assistant",
			Description: "aula-assistant",
			HideInIndex: false,
		},
		{
			Path:        "/dsupdate",
			URL:         "https://github.com/arnested/dsupdate",
			Description: "dsupdate",
			HideInIndex: false,
		},
		{
			Path:        "/healthy",
			URL:         "https://github.com/arnested/healthy",
			Description: "healthy",
			HideInIndex: false,
		},
		{
			Path:        "/triagebot",
			URL:         "https://github.com/arnested/triagebot",
			Description: "triagebot",
			HideInIndex: false,
		},
		{
			Path:        "/vanity-server",
			URL:         "https://github.com/arnested/vanity-server",
			Description: "vanity-server",
			HideInIndex: false,
		},
	}

	args := vanity.Args{
		Config: vanity.Config{
			Prefix:   "arnested.dk/go",
			Mappings: mappings,
		},
		NoIndex: false,
	}

	addr := ":0"
	if value, ok := os.LookupEnv("ADDR"); ok {
		addr = value
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Using port: %d", listener.Addr().(*net.TCPAddr).Port)

	handler := vanity.Handler(args)

	http.HandleFunc("/", handler)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
}
