FROM golang:alpine AS build
COPY *.go go.mod /go/
RUN go build

FROM alpine
COPY --from=build /go/duckdns duckdns
ENTRYPOINT [ "/duckdns" ]
