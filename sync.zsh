#!/bin/zsh

set -a
source .env
set +a

GOARCH=$GOARCH GOOS=$GOOS go build -o $APP_NAME

rsync -avz -e "sudo ssh -i $SSH_KEY_PATH" ./$APP_NAME root@$TARGET_HOST:$TARGET_PATH
rsync -avz -e "sudo ssh -i /Users/exiltrip/.ssh/id_ed25519" ./app root@192.168.0.54:~/api_go/