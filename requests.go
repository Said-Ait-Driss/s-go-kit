// the primary messaging pattern is RPC, this means that every method in our interface
// should be modeled as a client-service interaction. The requesting program is a client,
// and the service is the server. This allows us to specify the parameters and return types of each method

package main

type uppercaseRequest struct {
	S string `json:"s"`
}

type countRequest struct {
	S string `json:"s"`
}
