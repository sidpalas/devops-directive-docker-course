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

At this point, we already have a `docker compose` file which allows us to specify our application using a single command. There are a few limitations with docker compose that make it less than ideal to run a production application with docker compose directly:

- No support for secrets (we are passing our DB credentials as environment variables)
- No support for zero downtime re-deployments

Luckily, docker swarm does provide those things and setting up a single node cluster can be done with one command `docker swarm init`. Making just a few modifications to the docker compose file (adding `deploy` configurations, passing sensitive info as secrets and reading those data as files within the applications) it is ready to deploy.

1) Create a virtual machine with your favorite cloud provider. Make sure to set up the firewall to listen on ports 80, 443, and 22.
2) Use the script at https://get.docker.com/ to install docker engine. If managing docker as a non-root user, you may need to follow these steps: https://docs.docker.com/engine/install/linux-postinstall/ after installing.
3) Set the `DOCKER_HOST` environment variable in the Makefile to `USERNAME@IP_ADDRESS` of your virtual machine (this will allow your local docker client to use the remote docker daemon!)
4) Build and push the container images to a registry
5) Populate the secrets by running `make create-secrets`
6) Deploy the application by running `make swarm-deploy-stack` (uses the `docker stack deploy` command under the hood)
7) Set up a DNS A record to route traffic to your VM (or access using the IP address)

## Kubernetes

I was planning to only include the Railway + Swarm examples, but figured a course about containers wouldn't be complete without at least mentioning Kubernetes, the most popular container orchestrator today. Like Docker Swarm, Kubernetes is designed to schedule and run your containers, but has more maturity when it comes to cloud provider support.

I created the necessary resource yaml files in `./kubernetes` to deploy the application. You will notice that it is somewhat more verbose than the swarm specification.

1) Create a kubernetes cluster with your favorite cloud provider.
2) Set up kubectl to connect to the cluster (using the cloud provider instructions)
3) Install Traefik ingress controller by running `make install-traefik` (uses this helm chart https://github.com/traefik/traefik-helm-chart)
4) Install Postgres by running `make install-postgres` (uses this helm chart https://github.com/bitnami/charts/tree/main/bitnami/postgresql)
5) Deploy the application by running `make deploy-app`
6) Set up a DNS A record to route traffic to the IP address of the load balancer that traefik provisions
