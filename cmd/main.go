package main

import (
	"context"
	"github.com/yinebebt/priceestimation/initiator"
)

// main is application entrance point
func main() {
	initiator.Initiator(context.Background())
}
