# GO-Chat

A real-time chat server and client written in GO.



### Techniques Utilised:
- `net` package of GO Standard Library
- `channels` & `goroutines` to maintain concurrency in communication between multiple clients and server
- `tcp` network protocol


### How to Run
- Run `go run ./server` in a terminal window
- Run `go run ./client` in seperate terminal windows, to create multiple users.
- Communicate between the two clients by sending messages.
- Kill the server after completion `ctrl-c`

### Pre-requisites
- [Install](https://go.dev/doc/install) GO  