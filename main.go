package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const localhost = "http://localhost:8080/chain"

func main() {
	fmt.Println(Hash(Block{"123", "456", []byte("test data"), time.Now()}))

	NewMaster()
}

func AddBlock(c http.Client, b []Block) error {
	bJSON, _ := json.Marshal(b)
	req, err := http.NewRequest(http.MethodPost, localhost, bytes.NewReader(bJSON))
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
