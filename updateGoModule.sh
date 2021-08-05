#!/bin/bash
UsernameModuleName="0187773933/ADBWrapper"
GitHash=$(curl -s 'https://api.github.com/repos/$UsernameModuleName/git/refs/heads/master' | jq -r '.object.sha')
go get -u "github.com/$UsernameModuleName/@$GitHash"