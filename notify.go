package main

import (
    "github.com/btcsuite/btcrpcclient"
    "github.com/btcsuite/btcutil"
    "io/ioutil"
    "log"
    "path/filepath"
    "time"
)


func main() {
    // Setup handlers for blockconnected and blockdisconnected
    // notifications.
    ntfnHandlers := btcrpcclient.NotificationHandlers{
       
    }

    // Load the certificate for the TLS connection which is automatically
    // generated by btcd when it starts the RPC server and doesn't already
    // have one.
    btcdHomeDir := btcutil.AppDataDir("btcd", false)
    certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
    if err != nil {
        log.Fatal(err)
    }

    // Create a new RPC client using websockets.
    connCfg := &btcrpcclient.ConnConfig{
        Host:         "localhost:8334",
        Endpoint:     "ws",
        User:         "alexjavabraz",
        Pass:         "deltasp@5k",
        Certificates: certs,
    }
    client, err := btcrpcclient.New(connCfg, &ntfnHandlers)
    if err != nil {
        log.Fatal(err)
    }

    // Register for blockconnected and blockdisconneted notifications.
    if err := client.NotifyBlocks(); err != nil {
        client.Shutdown()
        log.Fatal(err)
    }

    // For this example, gracefully shutdown the client after 10 seconds.
    // Ordinarily when to shutdown the client is highly application
    // specific.
    log.Println("Client shutdown in 10 seconds...")
    time.AfterFunc(time.Second*10, func() {
        log.Println("Client shutting down...")
        client.Shutdown()
        log.Println("Client shutdown complete.")
    })

    // Wait until the client either shuts down gracefully (or the user
    // terminates the process with Ctrl+C).
    client.WaitForShutdown()
}