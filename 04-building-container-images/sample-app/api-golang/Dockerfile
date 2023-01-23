# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

CMD ["go", "run", "./main.go"]

# TODO use best practices: https://snyk.io/blog/containerizing-go-applications-with-docker/