package redirect

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathsToURL map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		dest, ok := pathsToURL[path]
		if ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r) // ie. call the next handler. passing request down the chain.
	}

}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	var data []pathURL
	err := yaml.Unmarshal(yml, &data)

	// unmarshal will parse yaml bytes and fills data with Go structs.
	// data is a slice of pathURL , so the YAML must match that structure. ie.
	// each yaml item must contain fields that map to path -> Path and url -> URL.
	// if it doesnt match, then unmarshal ret an err.

	if err != nil {
		return nil, err
	}

	m := make(map[string]string)

	for _, d := range data {
		m[d.Path] = d.URL
	}

	return MapHandler(m, fallback), nil
}
