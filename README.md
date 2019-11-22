# =telega=
HTTP proxy for sending messages to telegram group chats


## Motivation
As you probably know Telegram API is blocked by Russian authorities, meaning one cannot access ``https://api.telegram.org`` from within Russia.
This simple proxy can be run anywhere except Russia, on any cloud provider e.g. Heroku (free 🍺). 
Its primary use-case is sending group notifications from CI (e.g. Jenkins).

## Installation
```zsh
$ go build -ldflags="-s -w"
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
$ wget --header="Authorization: Basic bWFzbGljazoxMjM0NQ==" -q -O- --post-data="{\"text\":\"Yo, guys\"}" localhost:8080/send
```
