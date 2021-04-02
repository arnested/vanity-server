package main

import (
	"fmt"
	"net/http"

	"go.yhsif.com/vanity"
)

func main() {
	mappings := []vanity.Mapping{
		vanity.Mapping{
			Path:        "/dsupdate",
			URL:         "https://github.com/arnested/dsupdate",
			Description: "dsupdate",
			HideInIndex: false,
		},
		vanity.Mapping{
			Path:        "/triagebot",
			URL:         "https://github.com/arnested/triagebot",
			Description: "triagebot",
			HideInIndex: false,
		},
		vanity.Mapping{
			Path:        "/healthy",
			URL:         "https://github.com/arnested/healthy",
			Description: "healthy",
			HideInIndex: false,
		},
		vanity.Mapping{
			Path:        "/vanity-server",
			URL:         "https://github.com/arnested/vanity-server",
			Description: "vanity-server",
			HideInIndex: false,
		},
	}

	addr := ":80"
	args := vanity.Args{
		Config: vanity.Config{
			Prefix:   "arnested.dk/go",
			Mappings: mappings,
		},
		NoIndex: false,
	}

	handler := vanity.Handler(args)

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Print(err)
	}
}
