# Running Containers (with Docker)

Documentation: https://docs.docker.com/engine/reference/run/

Options everyone should know:
```
-d
--entrypoint
--env, -e, --env-file
--init
--interactive, -i
--mount, --volume, -v
--name
--network, --net
--platform
--publish, -p
--restart
--rm
--tty, -t
```

Less commonly used options, but worth knowing about:

```bash
--cap-add, --cap-drop
--cgroup-parent
--cpu-shares
--cpuset-cpus (pin execution to specific CPU cores)
--device-cgroup-rule,
--device-read-bps, --device-read-iops, --device-write-bps, --device-write-iops
--gpus (NVIDIA Only)
--health-cmd, --health-interval, --health-retries, --health-start-period, --health-timeout
--memory , -m
--pid, --pids-limit
--privileged
--read-only
--security-opt
--userns
```

## Example web app

### individual docker run commands

See Makefile:
```bash
make docker-build-all
make docker-run-all
```

- Uses the default docker bridge network
- Uses `--link` to enable easy host name for network connections
- Publishing ports useful to connect to each service individually from host, but only necessary to connect to the frontend
- Named containers make it easier to reference (e.g. with link), but does require removing them to avoid naming conflict
- Restart policy allows docker to restart the container (for example if database weren't up yet causing one of the api servers to crash)

### docker compose

See Makefile:
```bash
make compose-build
make compose-up
```

Using docker compose allows encoding all of the logic from the `docker build` and `docker run` commands into a single file. Docker compose also manages naming of the container images and containers, attaching to logs from all the containers at runtime, etc...
