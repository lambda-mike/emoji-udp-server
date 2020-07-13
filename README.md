# Emoji UDP Server Project

## Docker

### Build image

`docker build -t emoji/udp-server:latest .`

### Run container

`docker run -it --rm -p 54321:54321/udp --name emoji-udp-server emoji/udp-server:latest`

## Dev

IMPORTANT: `cd` into `src` folder and then run `go` cmds

### Build

`go build`

### Test

`go test ./...`

Test verboselly:

`go test -v ./...`

Test specific area:

`go test -v ./... -run Build`

### Run

`EMOJI_PORT=54321 ./emoji-udp-server`

`EMOJI_PORT=54321 ./emoji-udp-server -n 3 -s ','`

### Debug

You can use client e2e testing app located in `src/client/client.go` to test the server manually.

Adjust the port value in the source file if needed - it should match the value set in environment variable.

`go run client/client.go`

Type you cmd and hit enter.

Example commands:

- `3 :ok:`

- `:metrics:`
