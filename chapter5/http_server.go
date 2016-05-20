/**
go http server
 */
package main

import (
	"net/http"
	"log"
)

func main() {
	s := & http.Server{
		Addr:	":8080",
	}
	log.Fatal(s.ListenAndServe())
}
