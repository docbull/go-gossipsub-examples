# 🌒  🌓  🌔  🌕  🌖  🌗  🌘 

# Libp2p GossipSub Examples

Libp2p GossipSub-based examples for getting skilled using GossipSub. You can run some example apps and open the source codes to understand how it works.

## Examples

- [Chat App](#chat-app)
- [File Sharing](#file-sharing)

## Chat App

GossipSub chatting application is based on [go-libp2p pubsub example](https://github.com/libp2p/go-libp2p/tree/master/examples/pubsub).

You can chat with another peers in the same LAN and topic (P2P network group) by running this simple chat app.

Users can set own nickname by nick flag (--nickname=NICKNAME) and room name by room flag (--room=ROOMNAME). If you didn't set any names, your nickname would be $USER-RANDOM_TEXT and room name would be test by default.

Run chat app like this:

```go
go run . --nick=docbull --room=ChatApp
```

And run another chat app user in a new terminal:

```go
go run . --nick=watson --room=ChatApp
```

Enter any message in the terminal. The message would be sent to other peers using GossipSub. If you want to leave the chatting room, just enter /quit command.

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

In this example, you don't need to setup your name, it only works by default. On the other hand, network group name would be set by network flag (--network==NETWORK_NAME).

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
QmS... send a file: text.txt
```
