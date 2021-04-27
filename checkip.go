package check

import (
	"encoding/json"
	"net/http"
)

const (
	path = "https://api.ipify.org?format=json"
)

type ipifyResponse struct {
	IP string `json:"ip"`
}

func IP(client *http.Client) (string, error) {
	var resp = &http.Response{}
	var err error

	if client == nil {
		resp, err = http.Get(path)
	} else {
		resp, err = client.Get(path)
	}

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	ipResp := ipifyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&ipResp); err != nil {
		return "", err
	}

	return ipResp.IP, nil
}
