package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func makeRefreshURL(domains, token string) string {
	query := url.Values{
		"domains": {domains},
		"token":   {token},
		"verbose": {"true"},
	}.Encode()

	return fmt.Sprintf("https://www.duckdns.org/update?%v", query)
}

func tryRefreshIP(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed to refresh IP address: ", err.Error())
		return
	}

	if resp.StatusCode != 200 {
		log.Println("Failed to refresh IP address, HTTP error: ", resp.Status)
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response: ", err)
		return
	}

	strResp := string(buf)
	// Ignore NOCHANGE responses to avoid flooding console and logs
	if !strings.Contains(strResp, "NOCHANGE") {
		log.Println(strings.ReplaceAll(string(buf), "\n", " "))
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	domains := os.Getenv("DUCKDNS_DOMAINS")
	token := os.Getenv("DUCKDNS_TOKEN")

	if domains == "" || token == "" {
		log.Fatalln("Configuration is empty")
	}

	url := makeRefreshURL(domains, token)
	log.Println("Using url", url)
	tryRefreshIP(url)

	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for range ticker.C {
			tryRefreshIP(url)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	<-sigint
	ticker.Stop()
}
