##
## Build
##
FROM golang:1.18.2 AS build

WORKDIR /go/src/app
COPY . .

RUN go build -ldflags="-s" -o /clump 

##
## Containerize
##
FROM ubuntu:20.04

WORKDIR /
COPY --from=build /clump /clump

EXPOSE 8080

ENTRYPOINT ["/clump"]
