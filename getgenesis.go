package main 

import (
	"github.com/btcsuite/btcrpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)


func main(){

	btcdHomeDir := btcutil.AppDataDir("btcd", false)
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))

	if err != nil{
		log.Fatal(err)
	}

	connCfg := &btcrpcclient.ConnConfig{
		Host: "localhost:18334",
		Endpoint: "ws",
		User: "alexjavabraz",
		Pass: "deltasp@5k",
		Certificates: certs,
	}

	client, err := btcrpcclient.New(connCfg, nil)

	if err != nil{
		log.Fatal(err)
	}

	defer client.Shutdown()

	genesisHashStr := "000000008d2e3835e7ab612f4eabcef194c22c658e06265272c76dcfb2db5049"

	blockHash, err := chainhash.NewHashFromStr(genesisHashStr)

	if err != nil{
		log.Fatal(err)
	}

	block, err := client.GetBlockVerbose(blockHash)

	if err != nil{
		log.Fatal(err)
	}

	log.Printf("Hash: %v\n", block.Hash)

	log.Printf("Hash: %v\n", block.Hash)
    log.Printf("Previous Block: %v\n", block.PreviousHash)
    log.Printf("Next Block: %v\n", block.NextHash)
    log.Printf("Merkle root: %v\n", block.MerkleRoot)
    log.Printf("Timestamp: %v\n", time.Unix(block.Time, 0).UTC())
    log.Printf("Confirmations: %v\n", block.Confirmations)
    log.Printf("Difficulty: %f\n", block.Difficulty)
    log.Printf("Size (in bytes): %v\n", block.Size)
    log.Printf("Num transactions: %v\n", len(block.Tx))


}