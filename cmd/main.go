package main

import (
	"github.com/Oppodelldog/sigdump/internal/sig"
	"log"
	"os"
)

func main() {
	var sigCh = make(chan os.Signal)

	sig.RegisterSignals(sigCh)

	for s := range sigCh {
		log.Printf("received signal '%v'", s)
	}

	log.Println("signal channel closed")
}
