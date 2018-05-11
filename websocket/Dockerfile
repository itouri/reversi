FROM golang:1.8.0-alpine

# Expose HTTP port and set necessary environment variables
EXPOSE 23456

#TODO 汚い
# copy source code into the $GOPATH and switch to that directory
ADD ./handlers ${GOPATH}/src/github.com/itouri/reversi/websocket/handlers
ADD ./ws ${GOPATH}/src/github.com/itouri/reversi/websocket/ws
ADD ./util ${GOPATH}/src/github.com/itouri/reversi/websocket/util
ADD . ${GOPATH}/src/github.com/itouri/reversi/websocket/

WORKDIR ${GOPATH}/src/github.com/itouri/reversi/websocket

#TODO --no-cacheをつけないとapk addできない
RUN apk add --no-cache git

RUN go get -u github.com/Masterminds/glide/...
RUN glide update

# the default command runs the service in the foreground
# CMD ["go-echo"]

# compile source code and copy into $PATH
RUN go install

# CMD go run server.go
CMD ["/bin/sh"]
