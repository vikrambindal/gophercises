package ep2_urlshortener

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	fileio "ep2_urlshortener/fileio"
	json "ep2_urlshortener/fileio/json"
	yaml "ep2_urlshortener/fileio/yaml"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redirectUri := pathsToUrls[r.URL.Path]
		if redirectUri != "" {
			http.Redirect(w, r, redirectUri, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	})
}

func IOHandler(fileName *string, fallback http.Handler) (http.HandlerFunc, error) {
	data, err := ioutil.ReadFile(*fileName)
	if err != nil {
		return fallback.ServeHTTP, err
	}

	fileExtension := strings.SplitAfter(filepath.Ext(*fileName), ".")[1]

	var handler fileio.FileIO

	switch fileExtension {
	case "json":
		handler = json.JsonHandler{}
	case "yaml":
		fallthrough
	case "yml":
		fallthrough
	default:
		handler = yaml.YamlHandler{}
	}

	return process([]byte(data), fallback, handler)
}

func process(filedata []byte, fallback http.Handler, fileHandler fileio.FileIO) (http.HandlerFunc, error) {

	parsedData, err := fileHandler.Parse(filedata)
	if err != nil {
		return nil, err
	}
	pathMap := fileio.Build(parsedData)

	return MapHandler(pathMap, fallback), nil
}
