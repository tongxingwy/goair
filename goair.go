package goair

import (
	"github.com/tongxingwy/goair/bonjour"
	"log"
	"github.com/tongxingwy/goair/crypto/curve25519"
)

type airServerInterface interface {
	ReceivedAudioPacket(c *Client, data []byte, length int)
	SupportedMirrorFeatures() MirrorFeatures
}

type airServer struct {
	clients    map[string]*Client //the connected clients. Key names are based on the client's IP
	delegate   airServerInterface
	publicKey  [32]byte
	privateKey [32]byte
}

//Start the airplay server. The delegate will contain an interface of stuff to deal with (like audio/video streams, volume controls, etc)
func Start(serverName string, delegate airServerInterface) {
	var raopPort = 16000
	var mirroringPort = 17000
	priKey := curve25519.GeneratePrivateKey()
	pubKey := curve25519.PublicKey(priKey)
	s := airServer{
		clients:    make(map[string]*Client),
		publicKey:  pubKey,
		privateKey: priKey,
	}
	s.delegate = delegate
	// Start broadcasting available services in DNSSD.
	go bonjour.RegisterServices(serverName, raopPort, mirroringPort)

	// Start the Remote Audio Protocol Server.
	go s.startRAOPServer(raopPort)

	// Start the Airplay Server.
	s.startAirplay(mirroringPort)

	log.Println("RAOP and Airplay services stoped  ...")

}
