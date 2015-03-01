package goair

import (
	"encoding/hex"
	"fmt"
	"github.com/andrewtj/dnssd"
	"log"
	"net"
)

// Register RAOP and Airplay services in Bonjour/DNSSD.
func registerServices(servername string) {
	hardwareAddr := getMacAddress()

	name := fmt.Sprintf("%s@%s", hex.EncodeToString(hardwareAddr), servername)
	log.Printf("registerServices _raop._tcp servername: %s", name)
	op := dnssd.NewRegisterOp(name, "_raop._tcp", raopPort, registerServiceCallbackFunc)

	op.SetTXTPair("txtvers", "1")
	op.SetTXTPair("ch", "2")
	op.SetTXTPair("cn", "0,1,2,3")
	op.SetTXTPair("et", "0,1")
	op.SetTXTPair("sv", "false")
	op.SetTXTPair("da", "true")
	op.SetTXTPair("ft", "0x100029ff")
	op.SetTXTPair("sr", "44100")
	op.SetTXTPair("ss", "16")
	op.SetTXTPair("pw", "false")
	op.SetTXTPair("vn", "3")
	op.SetTXTPair("tp", "UDP")
	op.SetTXTPair("md", "0,1,2")
	op.SetTXTPair("vs", "150.33")
	op.SetTXTPair("rhd", "4.1.3")
	op.SetTXTPair("vv", "1")
	op.SetTXTPair("vn", "65537")
	op.SetTXTPair("sm", "false")
	op.SetTXTPair("ek", "1")
	op.SetTXTPair("am", "AppleTV3,1")
	op.SetTXTPair("sf","0x4")
	err := op.Start()
	if err != nil {
		log.Printf("Failed to register RAOP service: %s", err)
		return
	}

	airplayOp := dnssd.NewRegisterOp(servername, "_airplay._tcp", airplayPort, registerServiceCallbackFunc)

	airplayOp.SetTXTPair("deviceid", hardwareAddr.String())
	mask := 0x100029ff //0x00C0
	features := fmt.Sprintf("0x%x", mask)
	airplayOp.SetTXTPair("features", features)
	airplayOp.SetTXTPair("flags","0x4")
	airplayOp.SetTXTPair("model", "AppleTV3,1")
	airplayOp.SetTXTPair("srcvers","150.33")
	airplayOp.SetTXTPair("vv","1")
	err = airplayOp.Start()
	if err != nil {
		log.Printf("Failed to register airplay service: %s", err)
		return
	}
}

// Throw away callback func.
func registerServiceCallbackFunc(op *dnssd.RegisterOp, err error, add bool, name, serviceType, domain string) {
	log.Printf("registerServiceCallbackFunc error: %s", err)
	log.Printf("registerServiceCallbackFunc name:", name)
}

// getMacAddress gets the mac address to broadcast our DNS services on.
func getMacAddress() net.HardwareAddr {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
	}

	for _, inter := range interfaces {
		if inter.HardwareAddr != nil && len(inter.HardwareAddr) > 0 && inter.Flags&net.FlagLoopback == 0 && inter.Flags&net.FlagUp != 0 && inter.Flags&net.FlagMulticast != 0 && inter.Flags&net.FlagBroadcast != 0 {
			addrs, _ := inter.Addrs()
			for _, addr := range addrs {
				if addr.String() != "" {
					return inter.HardwareAddr
				}
			}
		}
	}
	log.Println("WARNING: didn't find mac address, using default one")
	return []byte{0x48, 0x5d, 0x60, 0x7c, 0xee, 0x22} //default because we couldn't find the real one
}
