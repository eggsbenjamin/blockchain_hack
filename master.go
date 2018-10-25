package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
