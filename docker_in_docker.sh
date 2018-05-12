#!/bin/sh
# $1 = image名

docker run -it -v /var/run/docker.sock:/var/run/docker.sock -v /tmp/ci:/ci $1 $2