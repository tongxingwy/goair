package zeroconf

import (
	"net"
	"encoding/hex"
	"log"
	"fmt"
	"os/signal"
	"os"
	"syscall"
	"github.com/grandcat/zeroconf"
)

// Register RAOP and Airplay services in Bonjour/DNSSD.
func RegisterServices(servername string, raopPort int, airplayPort int) {
	hardwareAddr := getMacAddress()
	interfaces := getInterfaces()

	name := fmt.Sprintf("%s@%s", hex.EncodeToString(hardwareAddr), servername)

	log.Printf("registerServices _raop._tcp servername: %s", name)
	s_raop, err := zeroconf.Register(name, "_raop._tcp", "", raopPort,
		[]string{"am=AppleTV3,2",
				 "ch=2",
				 "cn=0,1,2,3",
				 "da=true",
				 "ek=1",
				 "et=0,3,5",
				 "ft=0x5A7FFFF7,0x1E",
				 "md=0,1,2",
				 "rmodel=IotstarTV",
				 "pk=257c4b520b1075423863ecb8f1c401f59a8e9fe82411d436b658ba070144fb67",
				 "pw=false",
				 "sf=0x44",
				 "sm=false",
				 "sr=44100",
				 "ss=16",
				 "sv=false",
				 "tp=TCP,UDP",
				 "txtvers=1",
				 "vn=65537",
				 "vs=220.68",

		},
		interfaces)
	if err != nil {
		log.Printf("Failed to register RAOP service: %s", err)
		return
	}
	defer s_raop.Shutdown()

	log.Printf("registerServices _airplay._tcp servername: %s", servername)
	s_airplay, err := zeroconf.Register(servername, "_airplay._tcp", "", airplayPort,
		[]string{"deviceid=" + hardwareAddr.String(),
				 "features=0x5A7FFFF7,0x1E",
				 "flags=0x44",
				 "model=AppleTV3,2",
				 "pi=5e66cf9b-0a39-4e0c-9d32-081a8ce63231",
				 "pk=257c4b520b1075423863ecb8f1c401f59a8e9fe82411d436b658ba070144fb67",
				 "rmodel=IotstarTV",
				 "rrv=1.0",
				 "rsv=1.0",
				 "srcvers=220.68",
				 "vv=2",
		},
		interfaces)

	if err != nil {
		log.Printf("Failed to register airplay service: %s", err)
		return
	}
	defer s_airplay.Shutdown()

	// Clean exit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sig:
		log.Println("Exit by user")
		s_raop.Shutdown()
		s_airplay.Shutdown()
		break
	}
	log.Printf("RAOP and Airplay services over ...")
}

// getMacAddress gets the mac address to bro2adcast our DNS services on.
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

func getInterfaces() []net.Interface {
	var interfaces []net.Interface
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, ifi := range ifaces {
		if (ifi.Flags&net.FlagUp) == 0 || (ifi.Flags&net.FlagLoopback) == 4 {
			continue
		}
		addrs, err := ifi.Addrs()
		if err != nil || len(addrs) == 0 {
			continue
		}
		if (ifi.Flags & net.FlagMulticast) > 0 {
			interfaces = append(interfaces, ifi)
		}
		break
	}

	return interfaces
}
