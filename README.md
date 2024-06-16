# GO-Chat

A real-time chat server and client written in GO.



https://github.com/unnxt30/GO-Chat/assets/105226341/96806b35-a421-46e3-8a83-03c1d2ce167a



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
