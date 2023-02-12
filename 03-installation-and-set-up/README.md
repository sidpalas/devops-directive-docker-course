# Installation and Set Up

Docker Desktop: https://docs.docker.com/get-docker/

Docker Engine: https://get.docker.com/ 

***Note:*** See [02-technology-overview](../02-technology-overview/README.md) for a description of the difference between Docker Desktop and Docker Engine. If you are installing on your development system, you will most likely want Docker Desktop.

---

## Configuring Docker Desktop

The default settings are likely fine for getting started, but if you begin to run more intensive applications, you may want to adjust the resources available to Docker. This can be done within the settings panel in the GUI.

![](./readme-assets/docker-desktop-config.jpg)

---

## Running Your First Containers

Hello World:
```
docker run docker/whalesay cowsay "Hey Team! 👋"
```

Run Postgres:
```
docker run \
  --env POSTGRES_PASSWORD=foobarbaz \
  --publish 5432:5432 \
  postgres:15.1-alpine
```