package ledger

import (
	"fmt"
	"io"
	"net/http"
)

// Ledger is the interface to all ledger API calls
type Ledger struct {
	endpoint  string
	authToken string
}

// NewLedger returns instance of a Ledger
func NewLedger(endpoint string, authToken string) Ledger {
	return Ledger{endpoint: endpoint, authToken: authToken}
}

// GetEndpoint returns the enpoint of the ledger
func (l *Ledger) GetEndpoint() string {
	return l.endpoint
}

// DoRequest creates a new request to Ledger with necessary headers set
func (l *Ledger) DoRequest(method, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, l.endpoint+url, body)
	req.Header.Set("Content-Type", "application/json")
	if l.authToken != "" {
		req.Header.Add("Authorization", l.authToken)
	}
	resp, err := client.Do(req)
	if err != nil || (resp != nil && resp.StatusCode > 399) {
		return resp, fmt.Errorf("qledger: %s %s received HTTP status %s", req.Method, req.URL.String(), resp.Status)
	}
	return resp, err
}
