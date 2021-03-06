package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type httpSource struct {
	url  string
	name string
}

func (a *httpSource) getHitokoto() (string, error) {
	resp, err := http.Get(a.url)
	if resp.StatusCode != 200 || err != nil {
		return "", err
	}
	var hitokoto hitokoto
	err = json.NewDecoder(resp.Body).Decode(&hitokoto)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("「%s」 - %s", hitokoto.Hitokoto, hitokoto.From), nil
}

func (a *httpSource) getName() string {
	return a.name
}

func addHTTPSource(name, url string) error {
	if nameGetSource(name) != nil {
		return errors.New("Source " + name + " has existed")
	}
	sourceList = append(sourceList, &httpSource{
		name: name,
		url:  url,
	})
	return nil
}
