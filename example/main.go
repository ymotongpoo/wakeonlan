package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ymotongpoo/wakeonlan"
)

var pass = flag.String("p", "", "password string")

func usage() {
	fmt.Println("usage: wol [-p password] macaddr")
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
	}

	mac := strings.TrimSpace(flag.Arg(0))
	p, err := wakeonlan.NewMagicPacket(mac, []byte(*pass))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sending Magic Packet to %v\n", mac)
	err = p.Broadcast()
	if err != nil {
		log.Fatal(err)
	}
}
