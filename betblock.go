
package main

import(
    "github.com/btcsuite/btcrpcclient"
    "github.com/btcsuite/btcutil"
    "io/ioutil"
    "log"
    "path/filepath"
    "fmt"
)


func main(){

    btcdHomeDir := btcutil.AppDataDir("btcd", false)
    certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))

    fmt.Print(btcdHomeDir)
    fmt.Print("/")
     fmt.Println("rpc.cert")
    fmt.Println("Procurando o certificado")

    if err != nil {
        log.Fatal(err)
    }

    connCfg := &btcrpcclient.ConnConfig{
        Host: "127.0.0.1:18334",
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

        blockCount, err := client.GetBlockCount()

    if err != nil{
        log.Fatal(err)
    }

    log.Printf("Block count: %d", blockCount)



}
