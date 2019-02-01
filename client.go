package main

import (
	"flag"
	"log"

	coap "github.com/dustin/go-coap"
)

func request(serveraddress string, action string) {

	req := coap.Message{
		Type:      coap.Confirmable,
		Code:      coap.GET,
		MessageID: 12345,
		Payload:   []byte(""),
	}

	path := action

	req.SetOption(coap.ETag, "weetag")
	req.SetOption(coap.MaxAge, 3)
	req.SetPathString(path)

	c, err := coap.Dial("udp", serveraddress+":5683")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	rv, err := c.Send(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	if rv != nil {
		log.Printf("Response payload: %s", rv.Payload)
	}
}

func main() {
	server := flag.String("server", "localhost", "-server 192.168.7.2")
	action := flag.String("action", "/ping", "action /ping")
	flag.Parse()
	request(*server, *action)

}
