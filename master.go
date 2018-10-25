package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func NewMaster() {
	// rest api

	chain := Chain{}

	ch := NewChain(&chain)

	rootBlock := Block{
		ID: "1234",
	}

	ch.Add(rootBlock)

	id := Hash(rootBlock)
	secondBlock := Block{
		ID:       id,
		ParentID: rootBlock.ID,
	}

	ch.Add(secondBlock)

	fmt.Printf("%+v\n", ch)

	r := gin.Default()

	r.GET("/chain", func(c *gin.Context) {
		c.JSON(200, ch.chain)
	})

	r.POST("/chain", func(c *gin.Context) {
		blk := []Block{}
		data, _ := ioutil.ReadAll(c.Request.Body)
		json.Unmarshal(data, &blk)

		for _, b := range blk {
			bid := Hash(b)
			tipId := ch.chain[len(ch.chain)-1].ID

			if bid == tipId {
				ch.Add(b)
			} else {
				c.Status(http.StatusBadRequest)
				return
			}
		}

		c.Status(http.StatusOK)
	})
	r.Run()
}
