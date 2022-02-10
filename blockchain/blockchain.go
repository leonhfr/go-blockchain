package blockchain

import (
	"strings"

	"github.com/leonhfr/go-blockchain/block"
)

// Blockchain represents the list of Blocks
type Blockchain struct {
	blocks []*block.Block ``
}

// New creates a new Blockchain
func New() *Blockchain {
	return &Blockchain{
		[]*block.Block{block.NewGenesis()},
	}
}

// AddBlock adds a Block to the Blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := block.New(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) String() string {
	var blocks []string
	for _, b := range bc.blocks {
		blocks = append(blocks, b.String())
	}
	return strings.Join(blocks, "\n\n")
}
