# Mainflux Lite

[![License](https://img.shields.io/badge/license-Apache%20v2.0-blue.svg)](LICENSE)
[![Build Status](https://travis-ci.org/Mainflux/mainflux-lite.svg?branch=master)](https://travis-ci.org/Mainflux/mainflux-lite)
[![Go Report Card](https://goreportcard.com/badge/github.com/Mainflux/mainflux-lite)](https://goreportcard.com/report/github.com/Mainflux/mainflux-lite)
[![Join the chat at https://gitter.im/Mainflux/mainflux](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Mainflux Lite is the compact and simple single binary (monolith) server with Mainflux IoT functionalities.

It is very useful for development, prototyping and quick and simple deployments - in situation where full-blown Mainflux system (based on plethora of microservices) is not needed.

### Installation
Mainflux Lite uses [MongoDB](https://www.mongodb.com/) and [InfluxDB](https://influxdata.com/), so insure that these are installed on your system (more info [here](https://github.com/Mainflux/mainflux-lite/blob/master/doc/dependencies.md)).

Installing Mainflux Lite is trivial [`go get`](https://golang.org/cmd/go/):
```bash
go get github.com/mainflux/mainflux-lite
$GOBIN/mainflux-lite
```

If you are new to Go, more information about setting-up environment and fetching Mainflux Lite code can be found [here](https://github.com/Mainflux/mainflux-lite/blob/master/doc/install.md).

### Docker
Running Mainflux Lite in a Docker container is equally trivial:
```
wget https://raw.githubusercontent.com/Mainflux/mainflux-lite/master/docker-compose.yml
docker-compose up
```
You will of course have to have [Docker Compose](https://docs.docker.com/compose/) installed on your machine.

For more information about running Mainflux Lite in Docker take a look [here](https://github.com/Mainflux/mainflux-lite/blob/master/doc/docker.md).

### Documentation
Development documentation can be found on our [Mainflux GitHub Wiki](https://github.com/Mainflux/mainflux/wiki).

Swagger-generated API reference can be foud at [http://mainflux.com/apidoc](http://mainflux.com/apidoc).

### Community
#### Mailing lists
- [mainflux-dev](https://groups.google.com/forum/#!forum/mainflux-dev) - developers related. This is discussion about development of Mainflux IoT cloud itself.
- [mainflux-user](https://groups.google.com/forum/#!forum/mainflux-user) - general discussion and support. If you do not participate in development of Mainflux cloud infrastructure, this is probably what you're looking for.

#### IRC
[Mainflux Gitter](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

#### Twitter
[@mainflux](https://twitter.com/mainflux)

### License
[Apache License, version 2.0](LICENSE)
