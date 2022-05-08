package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	res := ""
	if len(os.Args) > 1 {
		keyAsString := os.Args[1]
		res = gen2(keyAsString)
	} else {
		res = gen1()
	}
	fmt.Printf("Your key: %s \n", res)
}

func gen1() (finkey string) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	res := hex.EncodeToString(bytes)
	return res
}
func gen2(key string) (finkey string) {
	keyAsBytes := []byte(key)
	res := hex.EncodeToString(keyAsBytes)
	return res
}
