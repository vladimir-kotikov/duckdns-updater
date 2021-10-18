FROM golang:alpine AS build
WORKDIR /go/src/duckdns
COPY *.go go.mod ./
RUN go build

FROM alpine
COPY --from=build /go/src/duckdns/duckdns ./
ENTRYPOINT [ "/duckdns" ]
