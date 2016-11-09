#!/bin/ash

sslocal -s "$SS_HOST" -p "$SS_PORT" -b "0.0.0.0" -l 1080 -k "${SS_PASS:-empty}" -t 600 -m aes-256-cfb &

steamer
