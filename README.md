# =telega=
HTTP proxy for sending messages to telegram chats

## Installation
```zsh
go build -ldflags="-s -w"
```

## Usage
```zsh
$ export TOKEN=1234567890abcdef
$ export CHAT_ID=-12345
$ ./telega
Starting server on port 8080 ...

$ curl -s -X POST localhost:8080/send --data "{\"text\": \"Hello world\"}"
$ http POST :8080/send <<< '{"text": "Hi folks!"}'
```