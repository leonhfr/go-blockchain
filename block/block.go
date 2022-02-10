package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block represents a single blockchain block
type Block struct {
	Timestamp int64  // Creation time of the block
	Prev      []byte // Hash of the previous block
	Hash      []byte // Hash of the current block
	Data      []byte // Information store in the block
}

// New creates a Block
func New(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)}
	block.SetHash()
	return block
}

// NewGenesis creates the genesis Block
func NewGenesis() *Block {
	return New("Genesis block", []byte{})
}

// SetHash computes and sets the block hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.Prev, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func (b *Block) String() string {
	return fmt.Sprintf("Prev: %x\nData: %s\nHash: %x", b.Prev, b.Data, b.Hash)
}
