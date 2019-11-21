package main

import (
	"net/http"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/roundrobin"
	"github.com/vulcand/oxy/testutils"
	"github.com/vulcand/oxy/buffer"
)
  
func main() {
	url1 := "http://localhost:8000"
	// url2 string = ""  Add URLs for other servers
	// url3 string = ""


	fwd, _ := forward.New()
	lb, _ := roundrobin.New(fwd)

	lb.UpsertServer(testutils.ParseURI(url1))
	// lb.UpsertServer(url2)
	// lb.UpsertServer(url3)

	buffer, _ := buffer.New(lb, buffer.Retry(`IsNetworkError() && Attempts() < 2`))

	s := &http.Server {
		Addr: ":8080",
		Handler: buffer,
	}

	s.ListenAndServe()
}
