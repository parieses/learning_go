package main

import (
	"encoding/json"
	"io"
	"net/http"
	"BTC/croe"
)
var blockchain *croe.BlockChain

func main() {
	blockchain = croe.NewBlockChain()
	run()
}
func run()  {
	http.HandleFunc("/blockchain/get",blockchainGet)
	http.HandleFunc("/blockchain/write",blockchainWrite)
	http.ListenAndServe(":8080",nil)

}
func blockchainGet(w http.ResponseWriter,r *http.Request)  {
	bytes , error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w,error.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))

}
func blockchainWrite(w http.ResponseWriter,r *http.Request)  {
	blockdata := r.URL.Query().Get("data")
	blockchain.SendData(blockdata)
	blockchainGet(w,r)
}

