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
# install docker-compose
RUN curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
RUN chmod +x /usr/local/bin/docker-compose
CMD /bin/bash