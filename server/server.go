package main

import (
	"fmt"
	"log"
	"net"

	coap "github.com/dustin/go-coap"
	"github.com/fatih/color"
)

func handlePing(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
	log.Printf("Got message in handlePing: path=%q: %#v from %v", m.Path(), m, a)
	if m.IsConfirmable() {
		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Content,
			MessageID: m.MessageID,
			Token:     m.Token,
			Payload:   []byte("pong"),
		}
		res.SetOption(coap.ContentFormat, coap.TextPlain)

		log.Printf("Transmitting from %#v", res)
		return res
	}
	return nil
}

func main() {
	mux := coap.NewServeMux()
	fmt.Println()
	color.Red("CoAP Server lite!")
	color.Blue("Avialable methods")
	color.Blue("___________________________________________________________")
	fmt.Println()
	color.Yellow("/ping")
	color.Yellow("When called, this method responds with a pong message")
	fmt.Println()
	color.Cyan("CoAP server up and running on port: 5683 ")
	fmt.Println()

	mux.Handle("/ping", coap.FuncHandler(handlePing))
	log.Fatal(coap.ListenAndServe("udp", ":5683", mux))
}
