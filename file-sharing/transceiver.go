package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const BufSize = 128

// File structure includes filename, file data, and 
// sender's peer ID
type File struct {
	FileName   string
	Data       []byte
	SenderPeer string
}

type FileTransceiver struct {
	ReceivedFile chan *File

	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	networkGroup string
	peerID       peer.ID
}

// JoinNetwork let the peers join in the topic that was set using
// network group name
func JoinNetwork(ctx context.Context, ps *pubsub.PubSub, peerID peer.ID, ng string) (*FileTransceiver, error) {
	topic, err := ps.Join(topicName(ng))
	if err != nil {
		return nil, err
	}

	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	ft := &FileTransceiver{
		ReceivedFile: make(chan *File, BufSize),
		ctx:          ctx,
		ps:           ps,
		topic:        topic,
		sub:          sub,
		peerID:       peerID,
		networkGroup: ng,
	}

	go ft.readLoop()
	return ft, nil
}

// readLoop reading files as loop, and it unmarshalling received
// data that marshalled as JSON format, and then it pushes the 
// raw data into ReceivedFile channel
func (ft *FileTransceiver) readLoop() {
	for {
		d, err := ft.sub.Next(ft.ctx)
		if err != nil {
			close(ft.ReceivedFile)
			return
		}
		if d.ReceivedFrom == ft.peerID {
			continue
		}
		cm := new(File)
		err = json.Unmarshal(d.Data, cm)
		if err != nil {
			continue
		}
		ft.ReceivedFile <- cm
	}
}

// Run waiting users' command and listening file transmission from 
// other peers
func (ft *FileTransceiver) Run() error {
	go ft.handleEvents()

	fmt.Println("--------------------------")
	fmt.Println("Network Group:", ft.networkGroup)
	fmt.Println("Your ID:", ft.peerID)
	fmt.Println("--------------------------")

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fileName := scanner.Text()
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err, "Please try it again.")
			continue
		}

		if len(data) == 0 {
			continue
		}
		if scanner.Text() == "/quit" {
			// when you enter /quit into the chat, it closes chat app.
			break
		} else {
			err := ft.PublishWithFileName(fileName, data)
			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}

// PublishWithFileName publishes File structure into topic, i.e. network group
func (ft *FileTransceiver) PublishWithFileName(fileName string, file []byte) error {
	m := File{
		FileName:   fileName,
		Data:       file,
		SenderPeer: ft.peerID.Pretty(),
	}
	fileBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return ft.topic.Publish(ft.ctx, fileBytes)
}

// handleEvents handles receiving events when you receive
// a file from other peers
func (ft *FileTransceiver) handleEvents() {
	for {
		select {
		case m := <-ft.ReceivedFile:
			ft.handleReceivedFile(m)
		case <-ft.ctx.Done():
			return
		}
	}
}

// handleReceivedFile shows who sent the file and what is the file name, 
// and store it into the directory that you run this example
func (ft *FileTransceiver) handleReceivedFile(receivedFile *File) {
	fmt.Println(receivedFile.SenderPeer, "sent a file:", receivedFile.FileName)
	file, err := os.Create(receivedFile.FileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(receivedFile.Data)
	if err != nil {
		panic(err)
	}
}

// topicName returns topic name that includes network group
// name, if you want to set randomized topic name, you should
// modify this function. e.g., using DES key value + network 
// group name
func topicName(ng string) string {
	return "Network Group" + ng
}
