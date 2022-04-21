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
