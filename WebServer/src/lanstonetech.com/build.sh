#!/bin/bash

rm -f ./webserver
go build -v ./webserver.go

if [ -a webserver ] ; then
	echo "==>build Server successed!"
else
	echo "==>build Server failed!"
	exit -1
fi
