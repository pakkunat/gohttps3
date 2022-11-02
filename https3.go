// The executable program need in main package.
package main

import (
	// for Fprintf.
	"fmt"
	// for ResponseWriter, Request, HandleFunc, ListenAndServe.
	"github.com/lucas-clemente/quic-go/http3"
	"net/http"
	//"os"
	//"crypto/tls"
	"log"
)

// Handler function.
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

// main function.
func main() {
	// Subscribe handler.
	//http.HandleFunc("/", handler)
	// Listen for port 443.
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handler))

	//w := os.Stdout

	//server := http3.Server {
	//	Addr: "0.0.0.0:443",
	//	TLSConfig: &tls.Config {
	//		MinVersion: tls.VersionTLS13,
	//		MaxVersion: tls.VersionTLS13,
	//		KeyLogWriter: w,
	//	},
	//	Handler: mux,
	//}
	err := http3.ListenAndServeQUIC("127.0.0.1:443", "/etc/pki/tls/certs/test.crt.pem", "/etc/pki/tls/private/privkey-nopass.pem", mux)
	if err != nil {
		log.Fatal(err)
	}
}
