# Building Container Images

## General Process

Dockerfiles generally have steps that are similar to those you would use to get your application running on a server.

1) Start with an Operating System
2) Install the language runtime
3) Install any application dependencies
4) Set up the execution environment
5) Run the application

***Note:** We can often jump right to #3 by choosing a base image that has the OS and language runtime preinstalled.*

## Writing good Dockerfiles:

For each of the components of the example application I have included a series of Dockerfiles (`Dockerfile.0` -> `Dockerfile.N`) starting with the most simple naive approach, and improving them with each step.

Types of improvments:
1) **Pinning a specific base image:** By specifying an image tag, you can avoid nasty surprises where the base image
2) **Choosing a smaller base image:** There are often a variety of base images we can choose from. Choosing a smaller base image will usually reduce the size of your final image.
3) **Choosing a more secure base image:** Like image size, we should consider the number of vulnerabilities in our base images and the attack surface area. Chaingaurd publishes a number of hardened images (https://www.chainguard.dev/chainguard-images).
4) **Specifying a working directory:** Many languages have a convention for how/where applications should be installed. Adhering to that convention will make it easier for developers to work with the container.
5) **Consider layer cache to improve build times:** By undersanding the layered nature of container filesytems and choosing when to copy particular files we can make better use of the Docker caching system.
6) **Use COPY —link where appropriate:** The `--link` option was added to the `COPY` command in march 2022. It allows you to improve cache behavior in certain situations by copying files into an independent image layer not dependent on its predecessors.
7) **Use a non-root user within the container:** While containers can utilize a user namespace to differentiate between root inside the container and root on the host, this feature won't always be leveraged and by using a non-root user we improve the default safety of the container.
8) **Specify the environment correctly:** Only install production dependencies for a production image, and specify any necessary environment variables to configure the language runtime accordingly.
9) **Avoid assumptions:** Using commands like `EXPOSE <PORT>` make it clear to users how the image is intended to be used and avoids the need for them to make assumptions.
10) **Use multi-stage builds where sensible:** For some situations, multi-stage builds can vastly reduce the size of the final image and improve build times. Learn about and use multi-stage builds where appropriate.

All of these techniques are leveraged across the example applications in this repo.

## Additional Features

There are some additional features of Dockerfiles that are not shown in the example applications but are worth knowing about. These are highlighted in `Dockerfile.sample` and the corresponding build / run commands in the `Makefile`

1) **Parser directives:**
2) **ARG:**
3) **Parser directives:**
4) **Mounting secrets:**
5) **ENTRYPOINT + CMD:**

## Beyond the scope of this course:

1) **buildx (multi-architecture images):** You can use a feature called `buildx` to create images for multiple architectures from a single Dockerfile. This video goes into depth on that topic: https://www.youtube.com/watch?v=hWSHtHasJUI