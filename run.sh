#!/bin/ash

 sslocal -s "$SERVER_IP" -p "$SERVER_PORT" -b "0.0.0.0" -l 1080 -k "${PASSWORD:-empty}" -t 600 -m aes-256-cfb &

 cd /go/src/github.com/dcb9/steamer

 steamer
