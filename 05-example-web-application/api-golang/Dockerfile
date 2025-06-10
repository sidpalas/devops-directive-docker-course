  FROM golang:1.19-buster AS build                                                          
                                                                                               
     WORKDIR /app                                                                              
     COPY go.sum go.mod ./                                                                     
     RUN --mount=type=cache,target=/go/pkg/mod \                                               
         --mount=type=cache,target=/root/.cache/go-build \                                     
         go mod download                                                                       
     COPY healthcheck/healthcheck.go ./healthcheck/                                            
     RUN go build \                                                                            
         -ldflags="-linkmode external -extldflags -static" \                                   
         -tags netgo \                                                                         
         -o healthcheck \                                                                      
         ./healthcheck/healthcheck.go                                                          
     COPY api-golang.go ./                                                                     
     RUN go build \                                                                            
         -ldflags="-linkmode external -extldflags -static" \                                   
         -tags netgo \                                                                         
         -o api-golang \                                                                       
         ./api-golang.go                                                                       
                                                                                               
     # Stage 2: Final Image                                                                    
     FROM scratch                                                                              
     ENV GIN_MODE=release                                                                      
     COPY --from=build /etc/passwd /etc/passwd                                                 
     COPY --from=build /app/healthcheck/healthcheck healthcheck                                
     COPY --from=build /app/api-golang api-golang                                              
                                                                                               
     USER nonroot                                                                              
     EXPOSE 8080                                                                               
     CMD ["./api-golang"]