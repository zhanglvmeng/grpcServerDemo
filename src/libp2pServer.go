package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
)

func main3()  {
	// deng pei's code
	//host, _ := libp2p.New(context.Background())
	//ma, _ := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/9001/p2p/1111")
	//pi, _ := peer.AddrInfoFromP2pAddr(ma)
	//host.Peerstore().AddAddr(pi.ID, ma, time.Second * 10)
	//stm,_:= host.NewStream(context.Background(), pi.ID, "/node/0.0.2")
	//stm.Write([]byte("hello world"))

	// demo 1 -- begin
	//// The context governs the lifetime of the libp2p node
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	//// To construct a simple host with all the default settings, just use `New`
	//h, err := libp2p.New(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("Hello World, my hosts ID is %s\n", h.ID())
	// demo 1 -- end

	// demo 2 -- begin
	// Set your own keypair
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h2, err := libp2p.New(ctx,
		// Use your own created keypair
		libp2p.Identity(priv),

		// Set your own listen address
		// The config takes an array of addresses, specify as many as you want.
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/9000"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello World, my second hosts ID is %s\n", h2.ID())

	// demo 2 -- end

}
