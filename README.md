# Golang-based Libp2p GossipSub Examples

## Examples

- [Chat App](#chat-app)

## Chat App

GossipSub chatting application is based on [go-libp2p pubsub example](https://github.com/libp2p/go-libp2p/tree/master/examples/pubsub).

You can chat with another peers in the same topic (P2P network group) by running this simple chat app.

Users can set own nickname by nick flag (--nickname=NICKNAME) and room name by room flag (--room=ROOMNAME). If you didn't set any names, your nickname would be inlab-RANDOM_TEXT and room name would be test by default.

Run this app like this:

```go
go run . --nick=docbull --room=ChatApp
```

Output:
```console
--------------------------
Room: ChatApp
Your name: docbull
--------------------------
```