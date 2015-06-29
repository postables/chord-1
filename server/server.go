package main

import (
	"flag"
	"fmt"
	"github.com/cbocovic/chord"
)

func main() {

	//set up flags
	addressPtr := flag.String("addr", "127.0.0.1:8888", "the port you will listen on for incomming messages")
	joinPtr := flag.String("join", "", "an address of a server in the Chord network to join to")

	flag.Parse()
	me := new(chord.ChordNode)

	//join node to network or start a new network
	if *joinPtr == "" {
		me = chord.Create(*addressPtr)
	} else {
		me = chord.Join(*addressPtr, *joinPtr)
	}
	fmt.Printf("My address is: %s.\n", *addressPtr)
	me.Maintain()

}
