package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const localhost = "http://localhost:8080/chain"

func main() {
	fmt.Println(Hash(Block{"123", "456", []byte("test data"), time.Now()}))

	typ := os.Getenv("TYPE")
	if typ != "MASTER" && typ != "CLIENT" {
		panic("TYPE MUST BE MASTER OR CLIENT!!!!!!!")
	}

	if typ == "MASTER" {
		NewMaster()
	} else {
		for {
			chain, err := GetChain(*http.DefaultClient)
			if err != nil {
				fmt.Println("GetChain err: ", err)
				continue
			}

			parent := chain[len(chain)-1]
			err = AddBlock(*http.DefaultClient, []Block{
				{
					ID:        Hash(parent),
					ParentID:  parent.ID,
					Data:      []byte("abooshk!!!"),
					Timestamp: time.Now(),
				},
			})
			if err != nil {
				fmt.Println("AddBlock err: ", err)
				continue
			}

			time.Sleep(time.Second)
		}
	}

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
	err = dec.Decode(&chain)
	return chain, err
}
