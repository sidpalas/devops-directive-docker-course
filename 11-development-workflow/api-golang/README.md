Remote debugging setup (vscode `launch.json`):

```json
    {
      "name": "Docker: Attach to Golang",
      "type": "go",
      "debugAdapter": "dlv-dap",
      "mode": "remote",
      "request": "attach",
      "port": 4000,
      "remotePath": "/app",
      "substitutePath": [
        {
          "from": "${workspaceFolder}/docker-course/devops-directive-docker-course/05-example-web-application/api-golang",
          "to": "/app"
        }
      ]
    }
```