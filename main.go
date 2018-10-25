package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(Hash(Block{"123", "456", []byte("test data"), time.Now()}))

	NewMaster()
}
