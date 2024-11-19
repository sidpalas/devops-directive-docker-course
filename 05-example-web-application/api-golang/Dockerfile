FROM golang:1.19-buster AS build
RUN useradd -u 1001 nonroot
WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

COPY . .

RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o api-golang

#################
FROM scratch
ENV GIN_MODE=release
EXPOSE 8080
WORKDIR /
COPY --from=build /etc/passwd /etc/passwd
COPY --from=Build /app/api-golang api-golang
USER nonroot
CMD ["./api-golang"]