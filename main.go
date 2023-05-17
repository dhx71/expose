package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	cname     = flag.String("cname", "www.example.com", "hostname to reach the listening address:port")
	token     = flag.String("token", "", "token to authenticate client")
	localPort = flag.Int("localport", 8080, "port where trafic to <cname> is forwarded to: 127.0.0.1:<localport>. Always sending to 127.0.0.1")
	server    = flag.String("server", "", "Must be https://<cname>/path of the server instance")
	dyndns2   = flag.String("dyndns2", "", "Url of the DNS update request. Will be called only when IP changes. The actual IP is append to the request.\nEx.: https://username:password@domains.google.com/nic/update?hostname=subdomain.yourdomain.com&myip=")
)

func main() {
	flag.Parse()
	if len(*server) > 0 {
		// this is a client
		log.Printf("connecting to %s...\n", *server)
		startClient()
		return
	}
	if len(*dyndns2) > 0 {
		go continuouslyUpdateIP()
	}
	startServer()
	fmt.Println("salut")
}
