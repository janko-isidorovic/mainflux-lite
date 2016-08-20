# Docker
Insure that you have [Docker](https://www.docker.com/) installed on your host.

Now you can fetch all the images and start the containers:

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

Note that `mainflux-lite` container is linked
to other two servers (running in their proper containers) and also exposes port 7070 on which it listens by binding it to `localhost`.
This means that `mainflux-lite` server can be reached now on `http://localhost:7070` even though it euns in it's own container (and not on localhost directly).

