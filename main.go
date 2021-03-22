package main

import (
	"log"
	"net/http"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/router"
)

func main() {
	rtr := router.New()

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
