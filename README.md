<img src = "https://user-images.githubusercontent.com/59289320/164983610-4a7c91ad-08c2-4aed-8e78-00f8e2d35829.png" width="50%">

# Libp2p GossipSub Examples

Libp2p GossipSub-based examples for getting skilled using GossipSub. You can run some example apps and open the source codes to understand how it works.

## Examples

- [Chat App](#chat-app)
- [File Sharing](#file-sharing)
- [Multi-host Chat App](#multi-host-chat-app)

## Chat App

GossipSub chatting application is based on [go-libp2p pubsub example](https://github.com/libp2p/go-libp2p/tree/master/examples/pubsub).

You can chat with another peers in the same LAN and topic (P2P network group) by running this simple chat app.

Users can set own nickname by nick flag `--nickname=NICKNAME` and room name by room flag `--room=ROOMNAME`. If you didn't set any names, your nickname would be $USER-RANDOM_TEXT and room name would be test by default.

Run chat app like this:

```go
go run . --nick=docbull --room=ChatApp
```

And run another chat app user in a new terminal:

```go
go run . --nick=watson --room=ChatApp
```

Enter any message in the terminal. The message would be sent to other peers using GossipSub. If you want to leave the chatting room, just enter `/quit` command.

Output A:
```console
--------------------------
Room: ChatApp
Your name: docbull
--------------------------
hi, there!
```

Output B:
```console
--------------------------
Room: ChatApp
Your name: watson
--------------------------
docbull : hi, there!
/quit
```

## File Sharing

File sharing example transfer a file that you entered in the terminal. A receiver prints who sent the file and stores on own directory.

In this example, you don't need to setup your name, it only works by default. On the other hand, network group name would be set by network flag `--network==NETWORK_NAME`.

Run file sharing example like this:

```go
go run . --network=FileSharing
```

Output:
```console
--------------------------
Network Group: FileSharing
Your name: QmS...
--------------------------
```

Open a new terminal in local or in the same LAN nodes:

```go
go run . --network=FileSharing
```

Now enter a name of file that you want to share.

Output A:
```console
--------------------------
Network Group: FileSharing
Your name: QmS...
--------------------------
text.txt
```

Output B:
```console
--------------------------
Network Group: FileSharing
Your name: QmP...
--------------------------
QmS... sent a file: text.txt
```

## Multi-host Chat App

This multi-host chat app is based on [go-libp2p example](https://github.com/libp2p/go-libp2p/tree/master/examples/ipfs-camp-2019).

The simple chat app example above is only working on the same subnet peers like an intranet. That means, the messages cannot go outside and inside neither.

This example lets you experience running chat app that available communicate with peers outside of subnet.

Running peer options:
- `--port`: Configures peer's listening port number
- `--mode`: If you want to run your node as a bootstrap node, set this flag as bootstrap; `--mode=bootstrap`
- `--bootstrap`: Decides connecting bootstrap peer using bootstrap peer's multiaddrs

Run bootstrap peer like this:

```go
go run . --mode=bootstrap --port=4001
```

Output:

```console
Listening on /ip4/BOOTSTRAP_IP/tcp/4001
Listening on /ip4/127.0.0.1/tcp/4001
Peer ID: QmS...
Copy and paste this multiaddrs for joining chat app in another peer: /ip4/BOOTSTRAP_IP/tcp/4001/p2p/QmS...
```


If you have bootstrap peer, lets run common peer to join in the chat:

```go
go run . --bootstrap=/ip4/BOOTSTRAP_IP/tcp/4001/p2p/QmS...
```

Output:

```console
Listening on /ip4/IP_ADDRESS/tcp/35021
Listening on /ip4/127.0.0.1/tcp/35021
Listening on /ip4/IP_ADDRESS/tcp/44693/ws
Listening on /ip4/127.0.0.1/tcp/44693/ws
Peer ID: QmT...
Connected to QmS...
```

Now, enter any message in the terminal, then it would be disseminated to all peers in the chat using GossipSub.
