package webrtc

import (
	"fmt"
	"net/http"
	"testing"
)

func TestTurn(t *testing.T) {
	var listenAddress string
	listenAddress = "0.0.0.0:8080"
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Printf("HTTP server listen on http://%s\n", listenAddress)

	StartTurn("127.0.0.1", 13902, "localhost", "foo", "bar")

	panic(http.ListenAndServe(listenAddress, nil))
}