# Mainflux Lite

[![License](https://img.shields.io/badge/license-Apache%20v2.0-blue.svg)](LICENSE)
[![Build Status](https://travis-ci.org/Mainflux/mainflux-lite.svg?branch=master)](https://travis-ci.org/Mainflux/mainflux-lite)
[![Go Report Card](https://goreportcard.com/badge/github.com/Mainflux/mainflux-lite)](https://goreportcard.com/report/github.com/Mainflux/mainflux-lite)
[![Join the chat at https://gitter.im/Mainflux/mainflux](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Mainflux Lite is the compact and simple single binary (monolith) server with Mainflux IoT functionalities.

It is very useful for development, prototyping and quick and simple deployments - in situation where full-blown Mainflux system (based on plethora of microservices) is not needed.

### Installation
#### TL;DR
```bash
go get github.com/mainflux/mainflux-lite
$GOBIN/mainflux-lite
```
#### Docker
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

#### Code
##### Prerequisite
If not set already, please set your `GOPATH` and `GOBIN` environment variables. For example:
```bash
mkdir -p ~/go
export GOPATH=~/go
export GOBIN=$GOPATH/bin
```

#### Get the code
Use [`go`](https://golang.org/cmd/go/) tool to "get" (i.e. fetch and build) `mainflux-lite` package:
```bash
go get github.com/mainflux/mainflux-lite
```

This will download the code to `$GOPATH/src/github.com/mainflux/mainflux-lite` directory,
and then compile it and install the binary in `$GOBIN` directory.

Now you can run the server:
```bash
$GOBIN/mainflux-lite
```

Please note that the binary `mainflux-lite` expects to find configuration file `config.yml` in
direcotry provided by `MAINFLUX_LITE_CONFIG_DIR` if this variable is set. Otherwise it looks for `config.yml`
in `$GOPATH/src/github.com/mainflux/mainflux-core-server`.

Note also that using `go get` is prefered than out-of-gopath code fetch by cloning the git repo like this:
```
git clone https://github.com/Mainflux/mainflux-core-server && cd mainflux-lite
go get
go build
MAINFLUX_CORE_SERVER_CONFIG_DIR=. ./mainflux-core-server
```
#### Dependencies
Mainflux Core Server is connected to `MongoDB` (and potentially `InfluxDB`) on southbound interface.

This is why to run Mainflux Core Server you have to have running:
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
MAINFLUX_LITE_CONFIG_DIR=. ./mainflux-core-server
```

Note that when running services in this way (weather they are installed in the localhost system or run and mapped on localhost ports) you will need to change [`config.yml`](config.yml) and replace `influx` and `mongo` hostnames by `localhost`

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
