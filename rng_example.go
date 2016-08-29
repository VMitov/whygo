package main

import (
	"encoding/hex"
	"fmt"
	"io"

	"github.com/miracl/gomiracl"
	_ "github.com/miracl/gomiracl/go"
)

func main() {
	impl, _ := miracl.GetImpl(miracl.GO)
	seed, _ := hex.DecodeString("9e8b4178790cd57a5761c4a6f164ba72")
	rng := miracl.NewSeededRand(impl, seed)

	randomNum := make([]byte, 12)
	io.ReadFull(rng, randomNum)

	fmt.Println(randomNum)
}
