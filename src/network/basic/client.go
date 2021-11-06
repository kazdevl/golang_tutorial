package basic

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func GetRequest(url string) (Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Response{}, err
	}
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	var body Response
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return Response{}, err
	}
	return body, nil
}
