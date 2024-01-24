FROM golang:1.19-bullseye AS build

WORKDIR /app 

COPY . .

# Copy only required insall files (better caching)
COPY go.mod go.sum ./

RUN go mod download

# Compile application during build rather than at runtime
# Add flags to statically link binary
RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o api-golang

CMD ["go", "run", "./main.go"]