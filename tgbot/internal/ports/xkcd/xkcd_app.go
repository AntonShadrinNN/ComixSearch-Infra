package xkcd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type XkcdComixSearcher struct {
	ctx      context.Context
	client   *http.Client
	endpoint string
}

func (cs XkcdComixSearcher) GetSuitableComices(ctx context.Context, keywords string, limit int) (map[string]string, error) {
	params := map[string]string{
		"keywords": keywords,
	}
	log.Println("Find comix")
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s?limit=%d", cs.endpoint, limit)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := cs.client.Do(req)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	var comixResp ComixApiResponse
	err = dec.Decode(&comixResp)
	if err != nil {
		return nil, err
	}

	return comixResp.Comices, nil
}

func NewSearcher(ctx context.Context, endpoint string) XkcdComixSearcher {
	return XkcdComixSearcher{
		ctx:      ctx,
		client:   http.DefaultClient,
		endpoint: endpoint,
	}
}
