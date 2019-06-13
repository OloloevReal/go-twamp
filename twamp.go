package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"

	slog "github.com/OloloevReal/go-simple-log"
	"golang.org/x/net/ipv4"
)

const Version = "0.0.1"

func main() {
	slog.Printf("Starting TWAMP, version %s", Version)
	defer slog.Println("TWAMP closed")

	slog.SetOptions(slog.SetDebug)

	mode := flag.String("mode", "reflector", "set -mode=reflector")
	ip := flag.String("ip", "127.0.0.1", "set -ip=127.0.01")
	port := flag.Int("port", 20001, "set -port=20001")
	_ = mode
	_ = ip
	_ = port
	flag.Parse()
	slog.Printf("Args mode: %s", *mode)
	slog.Printf("Args ip: %s", *ip)
	slog.Printf("Args port: %d", *port)

	switch *mode {
	case "reflector":
		if err := Reflector(*ip, *port); err != nil {
			slog.Printf("[ERROR] failed to start reflector, %s", err)
		}
	case "client":
	default:
		slog.Fatalf("[FATAL] Undefined mode: %s", *mode)
	}

}

func Reflector(ip string, port int) error {
	pcon, err := net.ListenPacket("ip4:udp", fmt.Sprintf("%s", ip))
	if err != nil {
		return err
	}
	defer pcon.Close()

	rawc, err := ipv4.NewRawConn(pcon)
	if err != nil {
		return err
	}
	defer rawc.Close()

	b := make([]byte, 2048)
	for {
		h, p, _, err := rawc.ReadFrom(b)
		if err != nil {
			slog.Printf("[DEBUG] failed to read from socket, %s", err)
		}

		if h != nil {
			//slog.Printf("[DEBUG] Header Src: %s, TOS: %d, TTL: %d\tPayload Len: %d", h.Src.String(), h.TOS, h.TTL, len(p))
			slog.Println("[DEBUG] ", h.String())
			if h.Protocol != 17 {
				continue
			}
			dstPort := binary.BigEndian.Uint16(p[2:4])
			if int(dstPort) != port {
				continue
			}
			slog.Printf("[DEBUG] Dst port: %d", dstPort)
			l := binary.BigEndian.Uint16(p[4:6])
			slog.Printf("[DEBUG] UDP Len: %d", l)
			slog.Printf("[DEBUG] TWAMP: % x", p[8:len(p)])
		} else {
			slog.Printf("[DEBUG] Header is nil\tPayload Len: %d", len(p))
		}

	}
	return nil
}
