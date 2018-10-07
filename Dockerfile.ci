# Heads-Up: This Dockerfile is *exclusively* used for CI. It is referenced by
# Jenkinsfile and should not be used by any other means.

FROM golang:1-alpine

USER root

RUN apk add --no-cache curl git make 

WORKDIR /go/src/app

RUN curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose && chmod +x /usr/local/bin/docker-compose