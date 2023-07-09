# Simple script for send message from telegrab bot to user

Not a best-practice, not a professional work. I created this util because i was need to send some message from linux machine from console by cron or when user connecting by ssh to a server.

## Requirements

1. [GoLang] Compiler (script was built with Go 1.20)
2. Existing telegram bot and it's API token
3. User chat identifier

## Building

If you want to build your own binary you can use next command:

```sh
    $ go build send-bot-message.go
```

Also ready-to-use binary located in bin/ folder of this repo. Just download it and follow next steps

## Usage

```bash

# Clone this repo to a local/remote machine
$ git clone https://github.com/hinotora/message-sender.git

# Cd to a folde and give exec permission to a script binary
$ cd message-sender/bin
$ chmod +x

# Execute script with required parameters
$ send-bot-message --token=api_telegram_bot_token --chat=chat_id_integer --text="Hello world!"
```

[GoLang]: <https://go.dev/dl/>
