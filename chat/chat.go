package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const ChatRoomBufSize = 128

type Chat struct {
	Messages chan *ChatMessage

	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	roomName string
	self     peer.ID
	nick     string
}

type ChatMessage struct {
	Message    string
	SenderID   string
	SenderNick string
}

func JoinChat(ctx context.Context, ps *pubsub.PubSub, selfID peer.ID, nickname string, roomName string) (*Chat, error) {
	topic, err := ps.Join(topicName(roomName))
	if err != nil {
		return nil, err
	}

	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	chat := &Chat{
		Messages: make(chan *ChatMessage, ChatRoomBufSize),
		ctx:      ctx,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		self:     selfID,
		nick:     nickname,
		roomName: roomName,
	}

	go chat.readLoop()
	return chat, nil
}

func (chat *Chat) readLoop() {
	for {
		msg, err := chat.sub.Next(chat.ctx)
		if err != nil {
			close(chat.Messages)
			return
		}
		if msg.ReceivedFrom == chat.self {
			continue
		}
		cm := new(ChatMessage)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}
		chat.Messages <- cm
	}
}

func (chat *Chat) Run() error {
	go chat.handleEvents()

	fmt.Println("--------------------------")
	fmt.Println("Room:", chat.roomName)
	fmt.Println("Your name:", chat.nick)
	fmt.Println("--------------------------")

	for {
		var msg string
		//fmt.Print(chat.nick, "> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		msg = scanner.Text()
		if len(msg) == 0 {
			continue
		}
		if msg == "/quit" {
			// when you enter /quit into the chat, it closes chat app.
			break
		} else {
			err := chat.Publish(msg)
			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}

func (chat *Chat) Publish(message string) error {
	m := ChatMessage{
		Message:    message,
		SenderID:   chat.self.Pretty(),
		SenderNick: chat.nick,
	}
	msgBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return chat.topic.Publish(chat.ctx, msgBytes)
}

func (chat *Chat) handleEvents() {
	for {
		select {
		case m := <-chat.Messages:
			fmt.Println(m.SenderNick, ":", m.Message)
		case <-chat.ctx.Done():
			return
		}
	}
}

func topicName(roomName string) string {
	return "chat-room" + roomName
}
