package croe

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

//区块结构体
type Block struct {
	Index         int64  //区块编号
	Timestamp     int64  //区块时间戳
	PrevBlockHash string //上个区块的hash值
	Hash          string //当前区块hash
	Data          string //区块数据
}
func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInByte := sha256.Sum256([] byte(blockData))
	return hex.EncodeToString(hashInByte[:])
}
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return  newBlock
}
func GenerateGenesisBlock() Block{
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock,"Genesis Block")
}
