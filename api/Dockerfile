FROM golang:1.8.0-alpine

# Expose HTTP port and set necessary environment variables
EXPOSE 12345

# copy source code into the $GOPATH and switch to that directory
ADD . ${GOPATH}/src/github.com/itouri/reversi/api/

WORKDIR ${GOPATH}/src/github.com/itouri/reversi/api

#TODO --no-cacheをつけないとapk addできない
RUN apk add --no-cache git

RUN go get -u github.com/Masterminds/glide/...
RUN glide update

# the default command runs the service in the foreground
# CMD ["go-echo"]

# compile source code and copy into $PATH
RUN go install

# CMD go run server.go
CMD go run server.go
