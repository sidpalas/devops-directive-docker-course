FROM golang:1.19-buster as build

WORKDIR /app

RUN useradd -u 1001 madu

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o api-golang 

###
FROM scratch

COPY --from=build /app/api-golang api-golang

COPY --from=build /etc/passwd /etc/paaswd

COPY --from=build /app/api-golang api-golang

USER madu

ENV GIN-MODE release

EXPOSE 8080

CMD [ "/api-golang" ]