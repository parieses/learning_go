package croe

import (
	"fmt"
	"log"
)

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.ApendBlock(&newBlock)
}
func NewBlockChain() *BlockChain {
	genesisBlock := GenerateGenesisBlock()
	blockchain :=BlockChain{}
	blockchain.ApendBlock(&genesisBlock)
	return  &blockchain
}
func (bc *BlockChain) ApendBlock(newBlock *Block) {
	if len(bc.Blocks) ==0  {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("异常")
	}
}
func isValid(newBlock Block, oldBlock Block) bool {

	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks{
		fmt.Print(*block)
		fmt.Print("\n")
	}

}
