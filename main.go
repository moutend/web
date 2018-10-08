package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Handler struct {
	Logger     *log.Logger
	FileServer http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Logger.Println(r.Method, r.URL.String())
	h.FileServer.ServeHTTP(w, r)

	return
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("shutting down")
	}
	return
}

func run(args []string) (err error) {
	if len(args) < 2 {
		return fmt.Errorf("specify root dir")
	}
	rootDir := args[1]
	fmt.Println("listening on port 4000")
	h := &Handler{
		Logger:     log.New(os.Stdout, "", 0),
		FileServer: http.FileServer(http.Dir(rootDir)),
	}
	return http.ListenAndServe(":4000", h)
}
