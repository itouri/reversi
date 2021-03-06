FROM ubuntu:latest
# install docker
RUN apt-get -y update \
    && apt-get install -y apt-transport-https ca-certificates curl software-properties-common \
    && curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
RUN apt-key fingerprint 0EBFCD88
RUN add-apt-repository \ 
    "deb [arch=amd64] https://download.docker.com/linux/ubuntu artful stable"
RUN apt-get update
RUN apt-get install -y docker-ce
RUN service docker start
# install docker-compose
RUN curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
RUN chmod +x /usr/local/bin/docker-compose
# install golang
RUN apt-get install -y golang-go
ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:$PATH
# install git
RUN apt-get install -y git
ADD . $GOPATH/src/github.com/itouri/reversi
# install glide
RUN go get -u github.com/Masterminds/glide/...
# compile each services
WORKDIR $GOPATH/src/github.com/itouri/reversi/api
RUN glide update
RUN CGO_ENABLED=0 go build -o /cigo/api
WORKDIR $GOPATH/src/github.com/itouri/reversi/websocket
RUN glide update
RUN CGO_ENABLED=0 go build -o /cigo/websocket
# docker-compose up
WORKDIR $GOPATH/src/github.com/itouri/reversi/ci/acceptance-test
#WORKDIR /home
CMD cp /cigo/api /ci \
    && cp /cigo/websocket /ci \
    && docker-compose up