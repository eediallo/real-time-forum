package handler

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	if !NotFoundHandler(w, r) {
		return
	}

	data := PageData{
		CSSHomePage:        indexCSS,
		Logo:               logPath,
		HomePath:           homePagePath,
		LogoCSS:            cssLogoPath,
		GoIamge:            goImagePath,
		RustImage:          rustImagePath,
		JsImage:            jsImagPath,
		GolangOfficialPage: goOfficialPagePath,
		RustOfficialPage:   rustOfficialPagePath,
		JSOfficialPage:     jsOfficialPagePath,
		WS: wsPath,
	}
	RenderTemplate(w, "index", data)
}
