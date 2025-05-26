package main

import (
	"log"
	"net/http"

	handler "github.com/adwinugroho/go-vercel-wedding-invitation/api"
	"github.com/common-nighthawk/go-figure"
)

// Local test
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	myFigure := figure.NewFigure("Echo..", "", true)
	myFigure.Print()
	http.HandleFunc("/", handler.Handler)
	log.Println("Listening on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
