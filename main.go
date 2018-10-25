package main

import "time"

type Block struct {
	ID        string
	ParentID  string
	Data      []byte
	Timestamp time.Time
}

func Hash(block Block) string {
	return ""
}

type Chain []Block

func (c Chain) Add(block Block) error {
	return nil
}

func main() {

}
