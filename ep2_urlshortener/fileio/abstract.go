package fileio

import (
	"ep2_urlshortener/model"
)

type FileIO interface {
	Parse(data []byte) ([]model.PathToUrl, error)
}

func Build(pathToUrls []model.PathToUrl) map[string]string {

	pathToUrlMap := make(map[string]string)

	for _, pathToUrl := range pathToUrls {
		pathToUrlMap[pathToUrl.Path] = pathToUrl.Url
	}

	return pathToUrlMap
}
