package fileio

import (
	"encoding/json"

	"ep2_urlshortener/model"
)

type JsonHandler struct {
}

func (jh JsonHandler) Parse(data []byte) ([]model.PathToUrl, error) {

	pathToUrls := []model.PathToUrl{}
	err := json.Unmarshal(data, &pathToUrls)

	if err != nil {
		return nil, err
	}

	return pathToUrls, nil
}
