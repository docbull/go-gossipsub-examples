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
