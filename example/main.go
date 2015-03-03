package main

import (
	"github.com/tongxingwy/goair"
	//"log"
)

type airplayHandler struct {
}

// Get this party started.
func main() {
	h := airplayHandler{}
	goair.Start("goair", &h) // set the display name of your server.
}

func (h *airplayHandler) ReceivedAudioPacket(c *goair.Client, data []byte, length int) {
	//log.Println("got an audio packet")
}

func (h *airplayHandler) SupportedMirrorFeatures() goair.MirrorFeatures {
	return goair.MirrorFeatures{Height: 1280, Width: 720, Overscanned: true, RefreshRate: 0.016666666666666666}
}
