package main

import (
	"github.com/gin-gonic/gin"
)

func NewMaster() {
	// rest api

	chain := Chain{}

	rootBlock := Block{
		ID: "1234",
	}

	chain.Add(rootBlock)

	id := Hash(rootBlock)
	secondBlock := Block{
		ID:       id,
		ParentID: rootBlock.ID,
	}

	chain.Add(secondBlock)

	r := gin.Default()

	r.GET("/chain", func(c *gin.Context) {
		c.JSON(200, chain)
	})

	r.POST("/chain", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
