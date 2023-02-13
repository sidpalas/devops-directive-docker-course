# Deploying containers

One of the major benefits of containerization is that it provides a standard interface that others can use to design their systems. Because of this, TONS of different options are available for deploying containers to the cloud.

Within AWS alone, Corey Quinn (from the Duckbill Group) noted there are 17 unique ways to run containers (https://www.lastweekinaws.com/blog/the-17-ways-to-run-containers-on-aws/). That was in 2021... the number has probably gone up by now!

We will be deploying in 3 different ways:

## Railway.app

Railway is a relatively new infrastructure company focused on making it as easy as possible for users to deploy applications. They offer automated deployments from GitHub, an intuitive user interface, and even created a technology called Nixpacks (https://nixpacks.com/) which enable you to create a container image automagically (without the need for a Dockerfile). For some situations the nixpacks work seamlessly (for the golang api it worked), for others it can be necessary to specify your own Dockerfile so that you have full control.

Unfortunately there is no way (in Feb 2023) to specify a Docker build context other than the location of the Dockerfile. Also, their build environment didn't support using `--mount=type=cache`. Because of this I created a `Dockerfile.railway` in the `06-building-container-images` directory. To prevent this cluttering the repo, I only included this on the `railway` git branch.

```
git fetch && git checkout railway # fetch and checkout the railway branch
```

The deployment configuration is specified via `railway.toml` files, also included on the `railway` git branch.

A specific `nginx-railway.conf` is also used in order to specify public domains of the apis. Within docker compose we use Docker's internal DNS on our bridge network, but Railway doesn't (yet) support private networking so we use the public domains.

The only other consideration is setting the `PORT` environment variable for each service (80 for nginx, 3000 for node, and 8080 for golang).

When we provision a PostgresDB within the project, Railway automatically sets the necessary `DATABASE_URL` environment variable for the other services.

## Docker swarm



## Kubernetes
