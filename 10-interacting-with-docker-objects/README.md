# Interacting with Containers and Other Docker Objects

Familiarize yourself with the docker command line!

You should:
1) Use the documentation here: https://docs.docker.com/engine/reference/commandline/cli/
2) Use the `--help` flag (e.g. `docker build --help`) to get more info about each command.

## Images

`docker image COMMAND`:
```
  build       Build an image from a Dockerfile (`docker build` is the same as `docker image build`)
  history     Show the history of an image
  import      Import the contents from a tarball to create a filesystem image
  inspect     Display detailed information on one or more images
  load        Load an image from a tar archive or STDIN
  ls          List images
  prune       Remove unused images
  pull        Pull an image or a repository from a registry
  push        Push an image or a repository to a registry
  rm          Remove one or more images
  save        Save one or more images to a tar archive (streamed to STDOUT by default)
  tag         Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE
```

### Scanning Images

Not a `docker image` subcommand, but still something you do with images:

```
docker scan IMAGE
```

***Note:*** You can also use a 3rd party scanner such as Trivy (https://github.com/aquasecurity/trivy)

### Signing Images

Another protection against software supply chain attacks is the ability to uniquely sign specific image tags to ensure an image was created by the entity who signed it.

```
docker trust sign IMAGE:TAG
docker trust inspect --pretty IMAGE:TAG
```

## Containers

`docker container COMMAND`:

```
  attach      Attach local standard input, output, and error streams to a running container
  commit      Create a new image from a container's changes
  cp          Copy files/folders between a container and the local filesystem
  create      Create a new container
  diff        Inspect changes to files or directories on a container's filesystem
  exec        Run a command in a running container
  export      Export a container's filesystem as a tar archive
  inspect     Display detailed information on one or more containers
  kill        Kill one or more running containers
  logs        Fetch the logs of a container
  ls          List containers
  pause       Pause all processes within one or more containers
  port        List port mappings or a specific mapping for the container
  prune       Remove all stopped containers
  rename      Rename a container
  restart     Restart one or more containers
  rm          Remove one or more containers
  run         Run a command in a new container
  start       Start one or more stopped containers
  stats       Display a live stream of container(s) resource usage statistics
  stop        Stop one or more running containers
  top         Display the running processes of a container
  unpause     Unpause all processes within one or more containers
  update      Update configuration of one or more containers
  wait        Block until one or more containers stop, then print their exit codes
```

## Volumes

`docker volume COMMAND`:
```
  create      Create a volume
  inspect     Display detailed information on one or more volumes
  ls          List volumes
  prune       Remove all unused local volumes
  rm          Remove one or more volumes
```

## Networks

`docker network COMMAND`:
```
  connect     Connect a container to a network
  create      Create a network
  disconnect  Disconnect a container from a network
  inspect     Display detailed information on one or more networks
  ls          List networks
  prune       Remove all unused networks
  rm          Remove one or more networks
```
