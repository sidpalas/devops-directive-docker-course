# Pin specific version for stability
# Use separate stage for building image
# Use debian for easier build utilities
FROM golang:1.19-bullseye AS build

WORKDIR /app 

# Copy only required insall files (better caching)
COPY go.mod go.sum ./

# Use cache mount to speed up install of existing dependencies
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

COPY . .

# Compile application during build rather than at runtime
# Add flags to statically link binary
RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o api-golang

# Use separate stage to reduce image size
FROM scratch

# Set gin mode
ENV GIN_MODE release

WORKDIR /

# Copy password file
COPY --from=build /etc/passwd /etc/passwd

# Copy the binary from build stage
COPY --from=build /app/api-golang api-golang

USER nonroot

# Expose port
EXPOSE 8080

CMD ["/api-golang"]