package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Handler serves files as HTTP response.
type Handler struct {
	debug *log.Logger
	file  http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.debug.Println(r.Method, r.URL.String())
	h.file.ServeHTTP(w, r)

	return
}

func run() error {
	host := flag.String("host", "localhost", "listening host")
	port := flag.String("port", "8080", "listening port")
	verbose := flag.Bool("verbose", false, "enable verbose outputs")

	flag.Parse()

	if len(flag.Args()) < 1 {
		return fmt.Errorf("specify working directory")
	}
	if host == nil {
		return fmt.Errorf("specify --host flag")
	}
	if port == nil {
		return fmt.Errorf("specify --port flag")
	}
	if verbose == nil {
		return fmt.Errorf("--verbose flag must be true or false")
	}

	debug := log.New(ioutil.Discard, "", 0)

	if *verbose {
		debug = log.New(os.Stdout, "debug: ", 0)
	}

	debug.Printf("listening on port %s\n", *port)

	h := &Handler{
		debug: debug,
		file:  http.FileServer(http.Dir(flag.Args()[0])),
	}

	return http.ListenAndServe(fmt.Sprintf("%s:%s", *host, *port), h)
}

func main() {
	if err := run(); err != nil {
		log.New(os.Stderr, "error: ", 0).Fatal(err)
	}
}
