package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	ID        string    `json:"id"`
	ParentID  string    `json:"parent_id"`
	Data      []byte    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func Hash(block Block) string {
	h := sha256.New()
	h.Write(
		bytes.Join(
			[][]byte{
				[]byte(block.ID),
				[]byte(block.ParentID),
				block.Data,
				[]byte(block.Timestamp.String()),
			},
			[]byte{},
		),
	)
	return fmt.Sprintf("%x", h.Sum(nil))
}

type Chain []Block

type chainHolder struct {
	chain Chain
}

func (c *chainHolder) Add(block Block) error {
	c.chain = append(c.chain, block)
	return nil
}

func NewChain(chain *Chain) *chainHolder {
	return &chainHolder{
		chain: *chain,
	}
}
