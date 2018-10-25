package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const localhost = "http://localhost:8080/chain"

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

	NewMaster()
}

func AddBlock(c http.Client, b Block) error {
	buff := bytes.NewReader{}
	enc := json.NewEncoder(buff)
	enc.Encode(b)
	bJSON, _ := json.Marshal(b)
	req, err := http.NewRequest(http.MethodPost, localhost, buff)
	resp, err := c.Do(req)
	if resp.StatusCode != 200 {
		return fmt.Errorf("received %d instead of 200", resp.StatusCode)
	}
	return err
}

func GetChain(c http.Client) (Chain, error) {
	resp, err := http.Get(localhost)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("received %d instead of 200", resp.StatusCode)
	}
	dec := json.NewDecoder(resp.Body)
	var chain Chain
	err = dec.Decode(chain)
	return chain, err
}
