// Binary fileserver is a file server example.
package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":8080", "address of the webserver")
	root = flag.String("root", "/var/www", "root directory")
)

func main() {
	flag.Parse()
	log.Fatal(http.ListenAndServe(
		*addr,
		http.FileServer(http.Dir(*root)),
	))
}
