package pkg

import (
	"net/http"
	"time"
)

type ArcDps struct {
	Latest  string    `json:"latest"`
	History []History `json:"history"`
}

// main function for arcdps
func modArcdps() {
	return
}

func latest_arcdps() (time.Time, error) {
	url := "https://www.deltaconnected.com/arcdps/"

	// make head request to url, parse out last-modified header
	client := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return time.Time{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return time.Time{}, err
	}
	// parse out time from last-modified header
	t, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
