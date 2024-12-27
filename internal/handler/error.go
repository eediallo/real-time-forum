package handler

import (
	"net/http"
)

func ErrorPageHandler(w http.ResponseWriter, errMsg string, element interface{}) {
	// Define the template
	data := struct {
		Msg     string
		Url     string
		Element interface{}
	}{
		Msg:     errMsg,
		Url:     homePagePath,
		Element: element,
	}
	RenderTemplate(w, "error", data)
}
