package main

import (
	"fmt"

	"github.com/leonhfr/go-blockchain/blockchain"
)

func main() {
	bc := blockchain.New()

	bc.AddBlock("Send some stuff")
	bc.AddBlock("Send some more stuff")

	fmt.Println(bc.String())
}
