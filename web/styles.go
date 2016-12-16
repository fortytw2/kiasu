package web

import (
	"net/http"

	"github.com/fortytw2/httpkit"
)

//go:generate bash -c "lessc --clean-css less/base.less dist/hydrocarbon.min.css"
//go:generate go-bindata -pkg web -o styles_generated.go -ignore .gitkeep dist/

// Stylesheet writes the stylesheet
func Stylesheet(w http.ResponseWriter, r *http.Request) error {
	buf, err := Asset("dist/hydrocarbon.min.css")
	if err != nil {
		return httpkit.Wrap(err, 404)
	}
	w.Header().Set("Content-Type", "text/css")

	_, err = w.Write(buf)
	if err != nil {
		panic(err)
	}

	return nil
}
