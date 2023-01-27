Docker Desktop: https://docs.docker.com/get-docker/

Docker Engine: https://get.docker.com/ 

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