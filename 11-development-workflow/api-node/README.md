Remote debugging setup (vscode `launch.json`):

```json
    {
      "name": "Docker: Attach to Node",
      "type": "node",
      "request": "attach",
      "localRoot": "${workspaceFolder}/docker-course/devops-directive-docker-course/05-example-web-application/api-node",
      "remoteRoot": "/usr/src/app",
      "port": 9229
    },
```