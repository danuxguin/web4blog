#!/bin/bash

rm -f ./webserver
CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -v ./webserver.go

if [ -a webserver ] ; then
	echo "==>build Server successed!"
else
	echo "==>build Server failed!"
	exit -1
fi
