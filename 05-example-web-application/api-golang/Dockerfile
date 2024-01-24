FROM golang:1.19-bullseye AS build

WORKDIR /app 

COPY . .

RUN go mod download

# Compile application during build rather than at runtime
RUN go build -o api-golang

CMD ["go", "run", "./main.go"]