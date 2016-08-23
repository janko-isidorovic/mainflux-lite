# Docker
Insure that you have [Docker](https://www.docker.com/) installed on your host as well as [Docker Compose](https://docs.docker.com/compose/).

It is very easy to start Docker composition (Mainflux Lite depends on MongoDB and InfluxDB and has to be started in composition with them):
```
wget https://raw.githubusercontent.com/Mainflux/mainflux-lite/master/docker-compose.yml
docker-compose up
```

Otherwise, you can fetch all the images one by one and start the containers manually:
```bash
# MongoDB prerequisite
docker pull mongo
docker run --name mongo -it mongo
# Mainflux Lite
docker pull mainflux/mainflux-lite
docker run --name mainflux-lite -p 7070:7070 --link=mongo:mongo -it mainflux/mainflux-lite
```

Note that `mainflux-lite` container is linked
to other two servers (running in their proper containers) and also exposes port 7070 on which it listens by binding it to `localhost`.
This means that `mainflux-lite` server can be reached now on `http://localhost:7070` even though it euns in it's own container (and not on localhost directly).

