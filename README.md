# =telega=
HTTP proxy for sending messages to telegram group chats


## Motivation
As you probably know Telegram API was blocked by Russian authorities a while ago, meaning one cannot access ``https://api.telegram.org`` from within Russia.

The solution is to run a proxy outside Russia. In fact there are many proxies out there (primarily SOCKS).
Their main disadvantage is that these proxies come and go, and you simply don't have control over this process. If you need a stable connection, you will eventually run your own server.

This simple HTTP proxy can be run on any cloud provider e.g. Heroku (free 🍺). 
Its primary use-case is sending group notifications from CI (e.g. Jenkins).

## Features
* Written in Go :heart:
* Lightweight: ~2.3 MB zipped
* Cloud-native friendly: Docker + k8s
* Secure: Basic authentication

## Installation
```zsh
$ go test
$ go build -ldflags="-s -w"
$ go build -ldflags="-s -w" && upx telega
```

## Usage
```zsh
$ export TOKEN=1234567890abcdef
$ export CHAT_ID=-12345
$ ./telega
Starting server on port 8080 ...

$ curl -s -X POST localhost:8080/send --data "{\"text\": \"Hello world\"}"
$ http POST :8080/send <<< '{"text": "Hi folks!"}'
$ wget -q -O- --post-data="{\"text\":\"Yo, guys\"}" localhost:8080/send
```

Basic authentication:
```zsh
$ export TOKEN=1234567890abcdef
$ export CHAT_ID=-12345
$ export USERNAME=maslick
$ export PASSWORD=12345
$ export PORT=4000
$ ./telega
Starting server on port 4000 ...

$ curl -s -H "Authorization: Basic bWFzbGljazoxMjM0NQ==" -X POST localhost:4000/send --data "{\"text\": \"Hello world\"}"
$ http -a maslick:12345 POST  :4000/send <<< '{"text": "Hi folks!"}'
$ wget --header="Authorization: Basic bWFzbGljazoxMjM0NQ==" -q -O- --post-data="{\"text\":\"Yo, guys\"}" localhost:4000/send
```

## Docker
```zsh
$ docker build -t maslick/telega .
$ docker run -d \
   -e TOKEN=1234567890abcdef \
   -e CHAT_ID=-12345 \
   -p 8081:8080 \
   maslick/telega

$ docker run -d \
   -e TOKEN=1234567890abcdef \
   -e CHAT_ID=-12345 \
   -e USERNAME=maslick \
   -e PASSWORD=12345 \
   -p 8082:8080 \
   maslick/telega

$ http POST `docker-machine ip default`:8081/send <<< '{"text": "Hi folks!"}'
$ http -a maslick:12345 POST `docker-machine ip default`:8082/send <<< '{"text": "Hi folks!"}'
```

## K8s
```zsh
$ kubectl apply -f k8s
$ kubectl set env deploy telega \
   TOKEN=1234567890abcdef \
   CHAT_ID=-12345 \
   USERNAME=maslick \
   PASSWORD=12345
```
