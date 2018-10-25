package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	ID        string
	ParentID  string
	Data      []byte
	Timestamp time.Time
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

func (c Chain) Add(block Block) error {
	return nil
}

func main() {
	fmt.Println(Hash(Block{"123", "456", []byte("test data"), time.Now()}))
}
