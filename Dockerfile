FROM golang:alpine AS build
COPY *.go /go
RUN go build -o duckdns

FROM scratch
COPY --from=build /go/duckdns duckdns
ENTRYPOINT [ "duckdns" ]
