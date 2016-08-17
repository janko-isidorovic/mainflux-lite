###
# Mainflux Dockerfile
###

FROM golang:alpine
MAINTAINER Mainflux

ENV MAINFLUX_CORE_SERVER_CONFIG_DIR=/config/lite

###
# Install
###

RUN apk update && apk add git && rm -rf /var/cache/apk/*

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/mainflux/mainflux-lite

RUN mkdir -p /config/lite
COPY config.yml /config/lite

# Get and install the dependencies
RUN go get github.com/mainflux/mainflux-lite

###
# Run main command from entrypoint and parameters in CMD[]
###
CMD [""]

# Run mainflux command by default when the container starts.
ENTRYPOINT /go/bin/mainflux-lite

