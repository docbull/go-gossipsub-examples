package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

	inputCh chan string
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

	cr := &Chat{
		Messages: make(chan *ChatMessage, ChatRoomBufSize),
		ctx:      ctx,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		self:     selfID,
		nick:     nickname,
		roomName: roomName,
	}

	go cr.readLoop()
	return cr, nil
}

func (cr *Chat) readLoop() {
	for {
		msg, err := cr.sub.Next(cr.ctx)
		if err != nil {
			close(cr.Messages)
			return
		}
		if msg.ReceivedFrom == cr.self {
			continue
		}
		cm := new(ChatMessage)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}
		cr.Messages <- cm
	}
}

func (cr *Chat) Run() error {
	go cr.handleEvents()

	return nil
}

func (cr *Chat) Publish(message string) error {
	m := ChatMessage{
		Message:    message,
		SenderID:   cr.self.Pretty(),
		SenderNick: cr.nick,
	}
	msgBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return cr.topic.Publish(cr.ctx, msgBytes)
}

func (cr *Chat) handleEvents() {
	for {
		select {
		case input := <-cr.inputCh:
			err := cr.Publish(input)
			if err != nil {
				printErr("publish error: %s", err)
			}
		case m := <-cr.Messages:
			prompt := withColor("green", fmt.Sprintf("<%s>:", cr.nick))
			fmt.Println(prompt, m)
		case <-cr.ctx.Done():
			return
		}
	}
}

func topicName(roomName string) string {
	return "chat-room" + roomName
}

// withColor wraps a string with color tags for display in the messages text box.
func withColor(color, msg string) string {
	return fmt.Sprintf("[%s]%s[-]", color, msg)
}
