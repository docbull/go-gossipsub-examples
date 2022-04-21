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