package mac

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"github.com/tongxingwy/dnssd"
)

// Register RAOP and Airplay services in Bonjour/DNSSD.
func RegisterServices(servername string, raopPort int, airplayPort int) {
	hardwareAddr := getMacAddress()

	name := fmt.Sprintf("%s@%s", hex.EncodeToString(hardwareAddr), servername)
	log.Printf("registerServices _raop._tcp servername: %s", name)
	op := dnssd.NewRegisterOp(name, "_raop._tcp", raopPort, registerServiceCallbackFunc)

	op.SetTXTPair("am", "AppleTV3,2")
	op.SetTXTPair("ch", "2")
	op.SetTXTPair("cn", "0,1,2,3")
	op.SetTXTPair("da", "true")
	op.SetTXTPair("ek", "1")
	op.SetTXTPair("et", "0,3,5")
	op.SetTXTPair("ft", "0x5A7FFFF7,0x1E")
	op.SetTXTPair("md", "0,1,2")
	op.SetTXTPair("rmodel", "IotstarTV")
	op.SetTXTPair("pk", "257c4b520b1075423863ecb8f1c401f59a8e9fe82411d436b658ba070144fb67")
	op.SetTXTPair("pw", "false")
	op.SetTXTPair("sf", "0x44")
	op.SetTXTPair("sm", "false")
	op.SetTXTPair("sr", "44100")
	op.SetTXTPair("ss", "16")
	op.SetTXTPair("sv", "false")
	op.SetTXTPair("tp", "TCP,UDP")
	op.SetTXTPair("txtvers", "1")
	op.SetTXTPair("vn", "65537")
	op.SetTXTPair("vs", "220.68")

	//op.SetTXTPair("rhd", "4.1.3")
	//op.SetTXTPair("vv", "1")
	//op.SetTXTPair("vn", "65537")
	//op.SetTXTPair("sm", "false")
	//op.SetTXTPair("ek", "1")
	//op.SetTXTPair("sf", "0x4")

	err := op.Start()
	if err != nil {
		log.Printf("Failed to register RAOP service: %s", err)
		return
	}

	airplayOp := dnssd.NewRegisterOp(servername, "_airplay._tcp", airplayPort, registerServiceCallbackFunc)
	log.Printf("registerServices _airplay._tcp servername: %s", servername)
	airplayOp.SetTXTPair("deviceid", hardwareAddr.String())
	//mask := 0x100029ff //0x00C0
	//features := fmt.Sprintf("0x%x", mask)
	airplayOp.SetTXTPair("features", "0x5A7FFFF7,0x1E")
	airplayOp.SetTXTPair("flags", "0x44")
	airplayOp.SetTXTPair("model", "AppleTV3,2")
	airplayOp.SetTXTPair("pi", "5e66cf9b-0a39-4e0c-9d32-081a8ce63231")
	airplayOp.SetTXTPair("pk", "257c4b520b1075423863ecb8f1c401f59a8e9fe82411d436b658ba070144fb67")
	airplayOp.SetTXTPair("rmodel", "IotstarTV")
	airplayOp.SetTXTPair("rrv", "1.0")
	airplayOp.SetTXTPair("rsv", "1.0")
	airplayOp.SetTXTPair("srcvers", "220.68")
	airplayOp.SetTXTPair("vv", "2")
	err = airplayOp.Start()
	if err != nil {
		log.Printf("Failed to register airplay service: %s", err)
		return
	}
}

// Throw away callback func.
func registerServiceCallbackFunc(op *dnssd.RegisterOp, err error, add bool, name, serviceType, domain string) {
	if err != nil {
		log.Printf("registerServiceCallbackFunc error: %s", err)
		log.Printf("registerServiceCallbackFunc name:", name)
	}
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
