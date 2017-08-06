package goair

import (
	"github.com/tongxingwy/goair/bonjour"
	"log"
)

type airServerInterface interface {
	ReceivedAudioPacket(c *Client, data []byte, length int)
	SupportedMirrorFeatures() MirrorFeatures
}

type airServer struct {
	clients  map[string]*Client //the connected clients. Key names are based on the client's IP
	delegate airServerInterface
}

//Start the airplay server. The delegate will contain an interface of stuff to deal with (like audio/video streams, volume controls, etc)
func Start(serverName string, delegate airServerInterface) {
	s := airServer{clients: make(map[string]*Client)}
	s.delegate = delegate
	// Start broadcasting available services in DNSSD.
	go bonjour.RegisterServices(serverName, raopPort, airplayPort)

	// Start the Remote Audio Protocol Server.
	go s.startRAOPServer()

	// Start the Airplay Server.
	s.startAirplay()

	log.Println("RAOP and Airplay services stoped  ...")

}
