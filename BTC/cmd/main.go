package main

import (
	"vlog/croe"
)

func main() {
	bc := croe.NewBlockChain()
	bc.SendData("send 1 btc")
	bc.SendData("send 1 eos")
	bc.SendData("send 1 btc")
	bc.SendData("send 1 eos")
	bc.SendData("send 1 btc")
	bc.SendData("send 1 eos")
	bc.SendData("send 1 btc")
	bc.SendData("send 1 eos")

	bc.Print()
}
