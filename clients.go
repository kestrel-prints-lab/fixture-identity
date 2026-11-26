// Generated from testenv/topology/topology.yaml by testenv/topology/generate.js. Do not edit by hand.
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Clients makes the outbound HTTP calls of identity.
type Clients struct {
	http *http.Client
}

// NewClients reads the target base URLs from the environment.
func NewClients() *Clients {
	return &Clients{
		http: &http.Client{Timeout: 5 * time.Second},
	}
}

func decode(resp *http.Response, err error) map[string]any {
	out := map[string]any{}
	if err != nil {
		return out
	}
	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&out)
	return out
}
