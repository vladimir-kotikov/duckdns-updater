duckdns: *.go go.mod
	@go build -o duckdns

image:
	@docker build .
