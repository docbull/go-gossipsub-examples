package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "pubsub-chat-example"

func main() {
	// nickFlag indicates user's nickname, you can set this name by --nick=NICKNAME
	nickFlag := flag.String("nick", "", "nickname to use in chat")
	// roomFlag indicates chat room name, you can set this name by --room=ROOMNAME
	roomFlag := flag.String("room", "test", "name of chat room to join")
	flag.Parse()

	ctx := context.Background()

	host, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}

	ps, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		panic(err)
	}

	// setup local mDNS discovery
	if err := setupDiscovery(host); err != nil {
		panic(err)
	}

	// nick is your nickname for this chat.
	nick := *nickFlag
	if len(nick) == 0 {
		nick = defaultNick(host.ID())
	}

	room := *roomFlag

	chat, err := JoinChat(ctx, ps, host.ID(), nick, room)
	if err != nil {
		panic(err)
	}

	err = chat.Run()
	if err != nil {
		panic(err)
	}
}

// defaultNick returns default nickname that you didn't set any
// nickname when you start this chat example.
func defaultNick(p peer.ID) string {
	return fmt.Sprintf("%s-%s", os.Getenv("USER"), shortID(p))
}

func shortID(p peer.ID) string {
	pretty := p.Pretty()
	return pretty[len(pretty)-8:]
}

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	_ = n.h.Connect(context.Background(), pi)
}

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func setupDiscovery(h host.Host) error {
	// setup mDNS discovery to find local peers
	s := mdns.NewMdnsService(h, DiscoveryServiceTag, &discoveryNotifee{h: h})
	return s.Start()
}
