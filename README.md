# Mainflux Lite

[![License](https://img.shields.io/badge/license-Apache%20v2.0-blue.svg)](LICENSE)
[![Build Status](https://travis-ci.org/Mainflux/mainflux-lite.svg?branch=master)](https://travis-ci.org/Mainflux/mainflux-lite)
[![Go Report Card](https://goreportcard.com/badge/github.com/Mainflux/mainflux-lite)](https://goreportcard.com/report/github.com/Mainflux/mainflux-lite)
[![Join the chat at https://gitter.im/Mainflux/mainflux](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Mainflux Lite is the compact and simple single binary (monolith) server with Mainflux IoT functionalities.

It is very useful for development, prototyping and quick and simple deployments - in situation where full-blown Mainflux system (based on plethora of microservices) is not needed.

### Installation
Mainflux Lite uses [MongoDB](https://www.mongodb.com/) and [InfluxDB](https://influxdata.com/), so insure that these are installed on your system.

Installing Mainflux Lite is trivial `go get`:
```bash
go get github.com/mainflux/mainflux-lite
$GOBIN/mainflux-lite
```

### Docker
```bash
# Influx prerequisite
docker pull influxdb
docker run --name influx -it influxdb
# MongoDB prerequisite
docker pull mongo
docker run --name mongo -it mongo
# Mainflux Lite
docker pull mainflux/mainflux-lite
docker run --name mainflux-lite -p 7070:7070 --link=mongo:mongo --link=influx:influx \
        -it mainflux/mainflux-lite
```

### Dependencies
Mainflux Lite server is connected to `MongoDB` (and potentially `InfluxDB`) on southbound interface.

This is why to run Mainflux Lite server you have to have running:
- [MongoDB](https://github.com/mongodb/mongo)
- [InfluxDB](https://github.com/influxdata/influxdb)

Installation and start of these services depends the operating system running on host (e.g. for Debian you can use `apt-get` to fetch and install these), so you must follow the installation instructions for each of the project.

However, each of these projects provides an official Docker image which can be pulled from DockerHub and started in a separate container (each in separate terminal if not detached):
```bash
docker run -p 27017:27017 -it mongo
docker run -p 8086:8086 -it influxdb
```
Now you can run `mainflux-lite`:
```bash
./mainflux-lite ./config/config.yml
```

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
