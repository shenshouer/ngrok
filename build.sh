#!/usr/bin/env bash

pwd=`pwd`

mkdir $pwd/bin
echo "build ngrok client ..."
cd $pwd/main/ngrok && GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-s" -a -installsuffix cgo  -o $pwd/bin/ngrok
echo "build ngrokd server ..."
cd $pwd/main/ngrokd && GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-s" -a -installsuffix cgo  -o $pwd/bin/ngrokd
echo "copy the necessary assets ..."
cp -r $pwd/assets $pwd/bin/
cp $pwd/client.yaml $pwd/bin/
