package zimuzu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

var (
	searchURL = "http://" + domain + "/search/api"
)

type searchResult struct {
	Status int             `json:"status"`
	Info   string          `json:"info"`
	Data   json.RawMessage `json:"data"`
}

type ResourceInfo struct {
	ID        string `json:"id"`
	ItemID    string `json:"itemid"`
	Title     string `json:"title"`
	Prefix    string `json:"prefix"`
	Suffix    string `json:"suffix"`
	Type      string `json:"type"`
	Channel   string `json:"channel"`
	Version   string `json:"version"`
	Character string `json:"character"`
	Pubtime   string `json:"pubtime"`
	Uptime    string `json:"uptime"`
	Poster    string `json:"poster"`
}

func (c *Ctx) Search(q string) ([]ResourceInfo, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s?keyword=%s", searchURL, url.QueryEscape(q)))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("search: http %d", resp.StatusCode)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result searchResult
	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, err
	}

	if result.Status != 1 {
		return nil, fmt.Errorf("search: %s", result.Info)
	}

	ret := make([]ResourceInfo, 0)
	if err := json.Unmarshal(result.Data, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}
