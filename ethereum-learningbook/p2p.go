package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rlp"
)

const messageId = 0

type Message string

func MyProtocol() p2p.Protocol {
	return p2p.Protocol{
		Name:    "MyProtocol",
		Version: 1,
		Length:  1,
		Run:     msgHandler,
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	nodekey, _ := crypto.GenerateKey()
	//fmt.Printf("the generated nodekey is %v \n", nodekey)
	config := p2p.Config{}
	config.MaxPeers = 10
	config.Name = "my mac node"
	config.ListenAddr = ":30300"
	config.PrivateKey = nodekey
	config.Protocols = []p2p.Protocol{MyProtocol()}

	srv := p2p.Server{
		Config: config,
	}
	if err := srv.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("the server is running ... \n")

	go func() {
		Msg := p2p.Msg{}
		Msg.Code = 404
		var data interface{}
		data = "processing"
		size, r, err := rlp.EncodeToReader(data)
		if err != nil {
			fmt.Printf("panic, err is %s \n", err)
		}
		Msg.Size = uint32(size)
		Msg.Payload = r
		ws, _ := p2p.MsgPipe()
		fmt.Printf("panic, err before write messge is %v and %v\n", ws, Msg)
		err = ws.WriteMsg(Msg)
		if err != nil {
			fmt.Printf("panic, err in write messge is %s \n", err)
		}
		//fmt.Printf("panic, err before write messge is %v and %v\n", ws, Msg)
		msg, err := ws.ReadMsg()
		if err != nil {
			fmt.Printf("panic, err is %s \n", err)
		}
		fmt.Printf("return %v \n", msg)
		wg.Done()
	}()

	wg.Wait()
	//select {}
}

func msgHandler(peer *p2p.Peer, ws p2p.MsgReadWriter) error {
	for {
		msg, err := ws.ReadMsg()
		if err != nil {
			return err
		}

		var myMessage Message
		err = msg.Decode(&myMessage)
		if err != nil {
			// handle decode error
			continue
		}

		switch myMessage {
		case "foo":
			err := p2p.SendItems(ws, messageId, "bar")
			//err := peer.SendItems(ws, messageId, "bar")
			if err != nil {
				return err
			}
		default:
			fmt.Println("recv:", myMessage)
		}
	}

	return nil
}
