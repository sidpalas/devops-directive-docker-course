# Pin specific version for stability
# Use debian for easier build utilities
FROM golang:1.19-bullseye AS build

WORKDIR /app 

# Copy only files required to install dependencies (better layer caching)
COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Compile application during build rather than at runtime
RUN go build -o api-golang

CMD ["./api-golang"]
