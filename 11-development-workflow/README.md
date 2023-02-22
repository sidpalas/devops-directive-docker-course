# Development Workflow

## Development Environment

Because we are running our application within containers, we need a way to quickly iterate and make changes to them. Some of our tactics in `06-building-container-images` help here (e.g. protecting the layer cache) so that images build quickly, but we can do better.

We want our development environment to have the following attributes:

1) **Easy/simple to set up:** Using docker compose, we can define the entire environment with a single yaml file. To get started, team members can issue a single command `make compose-up-build` or `make compose-up-build-debug` depending if they want to run the debugger or not.

2) **Ability to iterate without rebuilding the container image:** In order to avoid having to rebuild the container image with every single change, we can use a bind mount to mount the code from our host into the container filesystem. For example:

```yml
      - type: bind
        source: ../05-example-web-application/api-node/
        target: /usr/src/app/
```

3) **Automatic reloading of the application:** 
   - <ins>*React Client:*</ins> We are using Vite for the react client which handles this handles this automatically
   - <ins>*Node API:*</ins> We added nodemon as a development dependency and specify the Docker CMD to use it
   - <ins>*Golang API:*</ins> We added a utility called `air` (https://github.com/cosmtrek/air) within `Dockerfile.dev` which watches for changes and rebuild the app automatically.

4) **Use a debugger:**
   - <ins>*React Client:*</ins> For a react app, you can use the browser developer tools + extensions to debug. I did include `react-query-devtools` to help debug react query specific things. It is also viewed from within the browser.
   - <ins>*Node API:*</ins> To enable debugging for a NodeJS application we can run the app with the `--inspect` flag. The debug session can then be accessed via a websocket on port `9229`. The additional considerations in this case are to specify that the debugger listen for requests from 0.0.0.0 (any) and to publish port `9229` from the container to localhost.
   - <ins>*Golang API:*</ins> To enable remote debugging for a golang application I installed a tool called delve (https://github.com/go-delve/delve) within `./api-golang/Dockerfile.dev`. We then override the command used to run the container to use this tool (see: `docker-compose-debug.yml`)
     
      ---

      These modifications to the configuration (overridden commands + port publishing) are specified in `docker-compose-debug.yml`. By passing both `docker-compose-dev.yml` AND `docker-compose-debug.yml` to the `docker compose up` command (See: `make compose-up-debug-build`) Docker combines the two files, taking the config from the latter and overlaying it onto the former.

      Both `./api-golang/README.md` and `./api-node/README.md` show a launch.json configuration you can use to connnect to these remote debuggers using VSCode. The key setting is `substitutePath` such that you can set breakpoints on your local system that get recognized within the container.

5) **Executing tests:** We also need the ability to execute our test suites within containers. Again, we can create a custom `docker-compose-test.yml` overlay which modifies the container commands to execute our tests. To build the api images and execute their tests, you can execute `make run-tests` which will use the `test` compose file along with the `dev` compose file to do so.

## Continuous Integration

See `.github/workflows/image-ci.yml` for a basic GitHub Action workflow that builds, scans, tags, and pushes a container image.

It leverages a few publicly available actions from the marketplace:
1) https://github.com/marketplace/actions/docker-metadata-action (generates tags for the container images)
2) https://github.com/marketplace/actions/docker-login (logs into DockerHub)
3) https://github.com/marketplace/actions/build-and-push-docker-images (builds and pushes the images)
4) https://github.com/marketplace/actions/aqua-security-trivy (scans the images for vulnerabilities)

If you want to build out more advanced CI workflows I recommend looking at Bret Fisher's `Automation with Docker for CI/CD Workflows` repo (https://github.com/BretFisher/docker-cicd-automation). It has many great examples of the types of things you might want to do with Docker in a CI/CD pipeline!