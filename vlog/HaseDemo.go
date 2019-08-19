package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func main() {
	getHash("123")//2019/07/13 15:00:14 123,a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3
	getHash("123")//2019/07/13 15:00:14 123,a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3
	getHash("223")//2019/07/13 15:00:14 223,56f4da26ed956730309fa1488611ee0f13b0ac95ebb1bc9b5d210e31ff70e79c
}
func getHash(str string) string {
	hashInNyBytes := sha256.Sum256([] byte (str))
	//fmt.Print(hashInNyBytes)
	hashInNyStr :=hex.EncodeToString(hashInNyBytes[:])
	log.Printf("%s,%s",str,hashInNyStr)
	return  hashInNyStr
}
