package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
	"glorieux.io/adapter"
)

const applicationName = "docker-console"

// Version is the application version
// Set at build time
var Version string

func main() {
	showVersion := flag.Bool("version", false, fmt.Sprintf("Show %s version.", applicationName))
	openNewTab := flag.Bool("open", true, "Opens or not a new browser tab when launching.")
	port := flag.String("port", "4242", fmt.Sprintf("%s HTTP port.", applicationName))
	flag.Parse()
	if *showVersion {
		fmt.Println(Version)
		return
	}

	server, err := NewServer()
	if err != nil {
		log.Fatal(err)
	}

	if *openNewTab {
		err = open.Run("http://localhost:4242")
		if err != nil {
			log.Errorf("Can't launch browser at http://localhost:%s", *port)
		}
	} else {
		log.Infof("Application available at http://localhost:%s", *port)
	}

	http.ListenAndServe(":"+*port, adapter.Adapt(
		server,
		adapter.Timing(),
		handlers.CompressHandler,
	))
}
