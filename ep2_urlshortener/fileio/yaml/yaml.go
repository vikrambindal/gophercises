package fileio

import (
	"gopkg.in/yaml.v2"

	"ep2_urlshortener/model"
)

type YamlHandler struct {
}

func (y YamlHandler) Parse(yml []byte) ([]model.PathToUrl, error) {

	pathToUrl := []model.PathToUrl{}

	err := yaml.Unmarshal(yml, &pathToUrl)
	if err != nil {
		return nil, err
	}

	return pathToUrl, nil
}
