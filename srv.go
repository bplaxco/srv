package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	address := flag.String("b", "", "Bind to an address")
	folder := flag.String("f", "./", "Set the folder")
	port := flag.String("p", "1337", "Set the port")
	flag.Parse()

	fmt.Println("type -h to view avalible options")
	fmt.Println("Serving " + *folder + " at " + *address + ":" + *port)

	http.Handle("/", http.FileServer(http.Dir(*folder)))
	http.ListenAndServe(*address+":"+*port, nil)
}
