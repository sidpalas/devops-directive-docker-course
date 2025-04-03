[Home](../README.md) | [History and Motivation](../01-history-and-motivation/README.md)
| [Technology Overview](../02-technology-overview/README.md)
| [Installation and Set Up](../03-installation-and-set-up/README.md)
| [Using 3rd Party Containers](../04-using-3rd-party-containers/README.md)
| [Example Web Application](../05-example-web-application/README.md)
| [Building Container Images](../06-building-container-images/README.md)
| [Container Registries](../07-container-registries/README.md)
| [Running Containers](../08-running-containers/README.md)
| [Container Security](../09-container-security/README.md)
| [Interacting with Docker Objects](../10-interacting-with-docker-objects/README.md)
| [Development Workflows](../11-development-workflow/README.md)
| [Deploying Containers](../12-deploying-containers/README.md)

---

# Container Security

There are two main considerations when it comes to container security (1) the contents of your container image and (2) the security of the execution configuration and environment.

## Image Security

*“What vulnerabilities exist in your image that an attacker could exploit?”*

- Keep attack surface area as small as possible:
  - Use minimal base images (multi-stage builds are a key enabler)
  - Don’t install things you don’t need (don’t install dev deps)
- Scan images!
- Use users with minimal permissions
- Keep sensitive info out of images
- Sign and verify images
- Use fixed image tags, either:
  - Pin major.minor (allows patch fixes to be integrated)
  - Pin specific image hash

## Runtime Security

*If an attacker successfully compromises a container, what can they do? How difficult will it be to move laterally?*

### Docker daemon (dockerd)
  - Start with --userns-remap option(https://docs.docker.com/engine/security/userns-remap/)

### Individual containers:
- Use read only filesystem if writes are not needed
- --cap-drop=all, then --cap-add anything you need
- Limit cpu and memory --cpus=“0.5” --memory 1024m
- Use --security-opt
  - seccomp profiles (https://docs.docker.com/engine/security/seccomp/)
  - apparmor profiles (https://docs.docker.com/engine/security/apparmor/)