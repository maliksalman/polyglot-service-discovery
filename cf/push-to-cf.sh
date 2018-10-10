#!/bin/bash

pushd ../services
GOOS=linux go build frontend.go 
GOOS=linux go build color.go 
popd

cf target > /dev/null 2> /dev/null
if [[ $? -ne 0 ]]; then
	echo "ERROR: not logged into CloudFoundry"
	exit 1
fi


