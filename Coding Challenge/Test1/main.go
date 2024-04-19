package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func serveHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// enable 8000 port for egress and engress
// iptables -I INPUT -p tcp --dport 8000 -j ACCEPT
// iptables -I OUTPUT -p tcp --sport 8000 -j ACCEPT
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/authorizing-access-to-an-instance.html
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", serveHomePage)
	log.Fatal(http.ListenAndServe(":9000", router))
}
