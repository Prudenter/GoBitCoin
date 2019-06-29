#!/usr/bin/env bash
rm blockchain
rm *.db

go build -o blockchain *.go